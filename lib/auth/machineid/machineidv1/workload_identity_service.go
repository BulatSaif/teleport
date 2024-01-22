/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package machineidv1

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"net"
	"net/url"
	"time"

	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"

	pb "github.com/gravitational/teleport/api/gen/proto/go/teleport/machineid/v1"
	"github.com/gravitational/teleport/api/types"
	apievents "github.com/gravitational/teleport/api/types/events"
	"github.com/gravitational/teleport/lib/authz"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/teleport/lib/tlsca"
	usagereporter "github.com/gravitational/teleport/lib/usagereporter/teleport"
)

const spiffeScheme = "spiffe"

// WorkloadIdentityServiceConfig holds configuration options for
// the WorkloadIdentity gRPC service.
type WorkloadIdentityServiceConfig struct {
	Authorizer authz.Authorizer
	Cache      WorkloadIdentityCacher
	Backend    Backend
	Logger     logrus.FieldLogger
	Emitter    apievents.Emitter
	Reporter   usagereporter.UsageReporter
	Clock      clockwork.Clock
}

type WorkloadIdentityCacher interface {
	GetCertAuthority(ctx context.Context, id types.CertAuthID, loadKeys bool) (types.CertAuthority, error)
	GetClusterName(opts ...services.MarshalOption) (types.ClusterName, error)
}

type KeyStorer interface {
	GetTLSCertAndSigner(ctx context.Context, ca types.CertAuthority) ([]byte, crypto.Signer, error)
}

// NewWorkloadIdentityService returns a new instance of the
// WorkloadIdentityService.
func NewWorkloadIdentityService(
	cfg WorkloadIdentityServiceConfig,
) (*WorkloadIdentityService, error) {
	switch {
	case cfg.Cache == nil:
		return nil, trace.BadParameter("cache service is required")
	case cfg.Backend == nil:
		return nil, trace.BadParameter("backend service is required")
	case cfg.Authorizer == nil:
		return nil, trace.BadParameter("authorizer is required")
	case cfg.Emitter == nil:
		return nil, trace.BadParameter("emitter is required")
	case cfg.Reporter == nil:
		return nil, trace.BadParameter("reporter is required")
	}

	if cfg.Logger == nil {
		cfg.Logger = logrus.WithField(trace.Component, "bot.service")
	}
	if cfg.Clock == nil {
		cfg.Clock = clockwork.NewRealClock()
	}

	return &WorkloadIdentityService{
		logger:     cfg.Logger,
		authorizer: cfg.Authorizer,
		cache:      cfg.Cache,
		backend:    cfg.Backend,
		emitter:    cfg.Emitter,
		reporter:   cfg.Reporter,
		clock:      cfg.Clock,
	}, nil
}

// WorkloadIdentityService implements the teleport.machineid.v1.WorkloadIdentity
// RPC service.
type WorkloadIdentityService struct {
	pb.UnimplementedWorkloadIdentityServiceServer

	cache      WorkloadIdentityCacher
	backend    Backend
	authorizer authz.Authorizer
	keyStorer  KeyStorer
	logger     logrus.FieldLogger
	emitter    apievents.Emitter
	reporter   usagereporter.UsageReporter
	clock      clockwork.Clock
}

func (wis *WorkloadIdentityService) signX509SVID(
	ctx context.Context, req *pb.SVIDRequest, clusterName string, ca *tlsca.CertAuthority,
) (*pb.SVIDResponse, error) {
	// TODO: Authn/authz
	// TODO: Ensure they can issue the IPs, SANs and SPIFFE ID
	// TODO: Validate req.SpiffeIDPath for any potential weirdness

	spiffeID := &url.URL{
		Scheme: spiffeScheme,
		Host:   clusterName,
		Path:   req.SpiffeIdPath,
	}

	// Sign certificate

	ipSans := []net.IP{}
	for _, stringIP := range req.IpSans {
		ipSans = append(ipSans, net.ParseIP(stringIP))
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	template := &x509.Certificate{
		SerialNumber: serialNumber,
		// NotBefore is one minute in the past to prevent "Not yet valid" errors on
		// time skewed clusters.
		NotBefore: wis.clock.Now().UTC().Add(-1 * time.Minute),
		// TODO: Source TTL from req to the configured limit in rbac???
		NotAfter: wis.clock.Now().UTC().Add(time.Hour),
		// SPEC(X509-SVID) 4.3. Key Usage:
		// - Leaf SVIDs MUST NOT set keyCertSign or cRLSign.
		// - Leaf SVIDs MUST set digitalSignature
		// - They MAY set keyEncipherment and/or keyAgreement;
		KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageKeyAgreement,
		// SPEC(X509-SVID) 4.4. Extended Key Usage:
		// - Leaf SVIDs SHOULD include this extension, and it MAY be marked as critical.
		// - When included, fields id-kp-serverAuth and id-kp-clientAuth MUST be set.
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth,
		},
		// SPEC(X509-SVID) 4.1. Basic Constraints:
		// - leaf certificates MUST set the cA field to false
		BasicConstraintsValid: true,
		IsCA:                  false,

		// SPEC(X509-SVID) 2. SPIFFE ID:
		// - The corresponding SPIFFE ID is set as a URI type in the Subject Alternative Name extension
		// - An X.509 SVID MUST contain exactly one URI SAN, and by extension, exactly one SPIFFE ID.
		// - An X.509 SVID MAY contain any number of other SAN field types, including DNS SANs.
		URIs:        []*url.URL{spiffeID},
		DNSNames:    req.DnsSans,
		IPAddresses: ipSans,
	}

	certBytes, err := x509.CreateCertificate(
		rand.Reader, template, ca.Cert, req.PublicKey, ca.Signer,
	)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	pemBytes := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certBytes})

	// TODO: Audit and analytics event

	return &pb.SVIDResponse{
		SpiffeId:    spiffeID.String(),
		Hint:        req.Hint,
		Certificate: pemBytes,
	}, nil
}

func (wis *WorkloadIdentityService) SignX509SVIDs(ctx context.Context, req *pb.SignX509SVIDsRequest) (*pb.SignX509SVIDsResponse, error) {
	if len(req.Svids) == 0 {
		return nil, trace.BadParameter("svids: must be non-empty")
	}

	// Fetch info that will be needed for all SPIFFE SVIDs requested
	clusterName, err := wis.cache.GetClusterName()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	ca, err := wis.cache.GetCertAuthority(ctx, types.CertAuthID{
		Type:       types.SPIFFECA,
		DomainName: clusterName.GetClusterName(),
	}, true)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	tlsCert, tlsSigner, err := wis.keyStorer.GetTLSCertAndSigner(ctx, ca)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	tlsCA, err := tlsca.FromCertAndSigner(tlsCert, tlsSigner)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	res := &pb.SignX509SVIDsResponse{}
	for i, svidReq := range req.Svids {
		svidRes, err := wis.signX509SVID(
			ctx, svidReq, clusterName.GetClusterName(), tlsCA,
		)
		if err != nil {
			return nil, trace.Wrap(err, "signing svid %d", i)
		}
		res.Svids = append(res.Svids, svidRes)
	}

	return res, nil
}
