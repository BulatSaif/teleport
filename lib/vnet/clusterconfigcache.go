// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package vnet

import (
	"cmp"
	"context"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
	"golang.org/x/sync/singleflight"
)

type getClusterClientFunc = func(ctx context.Context, profileName, leafClusterName string) (ClusterClient, error)

type clusterConfig struct {
	// clusterName is the name of the cluster as reported by Ping.
	clusterName string
	// proxyPublicAddr is the public address of the proxy as reported by Ping, with any ports removed, this is
	// just the hostname. This is often but not always identical to the clusterName. For root clusters this
	// will be the same as the profile name (the profile is named after the proxy public addr).
	proxyPublicAddr string
	// dnsZones is the list of DNS zones that are valid for this cluster, this includes proxyPublicAddr *and*
	// any configured custom DNS zones for the cluster.
	dnsZones []string
	// ipv4CIDRRange is the CIDR range that IPv4 addresses should be assigned from for apps in this cluster.
	ipv4CIDRRange string
	// expires is the time at which this information should be considered stale and refetched. Stale data may
	// be used if a subsequent fetch fails.
	expires time.Time
}

func (e *clusterConfig) stale(clock clockwork.Clock) bool {
	return clock.Now().After(e.expires)
}

// clusterConfigCache is a read-through cache for cluster VnetConfigs. Cached entries go stale after 5
// minutes, after which they will be re-fetched on the next read.
//
// If a read from the cluster fails but there is a stale cache entry present, this prefers to return the stale
// cached entry. This is desirable in cases where the profile for a cluster expires during VNet operation,
// it's better to use the stale custom DNS zones than to remove all DNS configuration for that cluster.
type clusterConfigCache struct {
	flightGroup singleflight.Group
	clock       clockwork.Clock
	getClient   getClusterClientFunc
	cache       map[string]*clusterConfig
	mu          sync.RWMutex
}

func newClusterConfigCache(getClient getClusterClientFunc, clock clockwork.Clock) *clusterConfigCache {
	return &clusterConfigCache{
		clock:     clock,
		getClient: getClient,
		cache:     make(map[string]*clusterConfig),
	}
}

func (c *clusterConfigCache) getClusterConfig(ctx context.Context, clusterClient ClusterClient) (*clusterConfig, error) {
	k := clusterClient.ClusterName()

	// Use a singleflight.Group to avoid concurrent requests for the same cluster VnetConfig.
	result, err, _ := c.flightGroup.Do(k, func() (any, error) {
		// Check the cache inside flightGroup.Do to avoid the chance of immediate repeat calls to the cluster.
		c.mu.RLock()
		existingCacheEntry, existingCacheEntryFound := c.cache[k]
		c.mu.RUnlock()
		if existingCacheEntryFound && !existingCacheEntry.stale(c.clock) {
			return existingCacheEntry, nil
		}

		clusterConfig, err := c.getClusterConfigUncached(ctx, clusterClient)
		if err != nil {
			// It's better to return a stale cached VnetConfig than an error. The profile probably expired and
			// we want to keep functioning until a relogin. We don't expect the VnetConfig to change very
			// often.
			if existingCacheEntryFound {
				return existingCacheEntry, nil
			}
			return nil, trace.Wrap(err)
		}

		c.mu.Lock()
		c.cache[k] = clusterConfig
		c.mu.Unlock()

		return clusterConfig, nil
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return result.(*clusterConfig), nil
}

func (c *clusterConfigCache) getClusterConfigUncached(ctx context.Context, clusterClient ClusterClient) (*clusterConfig, error) {
	pingResp, err := clusterClient.CurrentCluster().Ping(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	clusterName := pingResp.ClusterName
	proxyPublicAddr := pingResp.ProxyPublicAddr
	if strings.Contains(proxyPublicAddr, ":") {
		proxyPublicAddr, _, err = net.SplitHostPort(pingResp.ProxyPublicAddr)
		if err != nil {
			return nil, trace.Wrap(err, "parsing proxy public addr")
		}
	}

	dnsZones := []string{proxyPublicAddr}
	ipv4CIDRRange := defaultIPv4CIDRRange

	vnetConfig, err := clusterClient.CurrentCluster().GetVnetConfig(ctx)
	if trace.IsNotFound(err) || trace.IsNotImplemented(err) {
		// Use the defaults set above, nothing to do here.
	} else if err != nil {
		return nil, trace.Wrap(err)
	} else {
		for _, zone := range vnetConfig.GetSpec().GetCustomDnsZones() {
			dnsZones = append(dnsZones, zone.GetSuffix())
		}
		ipv4CIDRRange = cmp.Or(vnetConfig.GetSpec().GetIpv4CidrRange(), defaultIPv4CIDRRange)
	}

	return &clusterConfig{
		clusterName:     clusterName,
		proxyPublicAddr: proxyPublicAddr,
		dnsZones:        dnsZones,
		ipv4CIDRRange:   ipv4CIDRRange,
		expires:         c.clock.Now().Add(5 * time.Minute),
	}, nil
}
