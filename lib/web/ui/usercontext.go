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

package ui

import (
	"github.com/gravitational/teleport/api/client/proto"
	"github.com/gravitational/teleport/api/constants"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/services"
)

type accessStrategy struct {
	// Type determines how a user should access teleport resources.
	// ie: does the user require a request to access resources?
	Type types.RequestStrategy `json:"type"`
	// Prompt is the optional dialog shown to user,
	// when the access strategy type requires a reason.
	Prompt string `json:"prompt"`
}

// AccessCapabilities defines allowable access request rules defined in a user's roles.
type AccessCapabilities struct {
	// RequestableRoles is a list of roles that the user can select when requesting access.
	RequestableRoles []string `json:"requestableRoles"`
	// SuggestedReviewers is a list of reviewers that the user can select when creating a request.
	SuggestedReviewers []string `json:"suggestedReviewers"`
}

type authType string

const (
	authLocal authType = "local"
	authSSO   authType = "sso"
)

// UserContext describes user settings and access to various resources.
type UserContext struct {
	// AuthType is auth method of this user.
	AuthType authType `json:"authType"`
	// Name is this user name.
	Name string `json:"userName"`
	// ACL contains user access control list.
	ACL services.UserACL `json:"userAcl"`
	// Cluster contains cluster detail for this user's context.
	Cluster *Cluster `json:"cluster"`
	// AccessStrategy describes how a user should access teleport resources.
	AccessStrategy accessStrategy `json:"accessStrategy"`
	// AccessCapabilities defines allowable access request rules defined in a user's roles.
	AccessCapabilities AccessCapabilities `json:"accessCapabilities"`
	// ConsumedAccessRequestID is the request ID of the access request from which the assumed role was
	// obtained
	ConsumedAccessRequestID string `json:"accessRequestId,omitempty"`
	// AllowedSearchAsRoles is the SearchAsRoles the user has access to for creating access requests.
	AllowedSearchAsRoles []string `json:"allowedSearchAsRoles"`
	// PasswordState specifies whether the user has a password set or not.
	PasswordSate types.PasswordState `json:"passwordState"`
	// SSOContext contains information regarding the SSO session, if this user is logged in via an auth connector.
	SSOContext SSOContext `json:"ssoContext,omitempty"`
}

// SSOContext contains information regarding the SSO session, if this user is logged in via an auth connector.
type SSOContext struct {
	// ConnectorType is the type of the SSO connector, either "saml", "oidc", or "github".
	ConnectorType string `json:"connectorType,omitempty"`
	// SAMLSingleLogoutURL is the SAML Single log-out URL to initiate SAML SLO, if configured.
	SAMLSingleLogoutURL string `json:"samlSingleLogoutUrl,omitempty"`
}

func getAccessStrategy(roleset services.RoleSet) accessStrategy {
	strategy := types.RequestStrategyOptional
	prompt := ""

	for _, role := range roleset {
		options := role.GetOptions()

		if options.RequestAccess == types.RequestStrategyReason {
			strategy = types.RequestStrategyReason
			prompt = options.RequestPrompt
			break
		}

		if options.RequestAccess == types.RequestStrategyAlways {
			strategy = types.RequestStrategyAlways
		}
	}

	return accessStrategy{
		Type:   strategy,
		Prompt: prompt,
	}
}

// NewUserContext returns user context
func NewUserContext(user types.User, userRoles services.RoleSet, features proto.Features, desktopRecordingEnabled, accessMonitoringEnabled bool) (*UserContext, error) {
	acl := services.NewUserACL(user, userRoles, features, desktopRecordingEnabled, accessMonitoringEnabled)
	accessStrategy := getAccessStrategy(userRoles)

	// Default user context.
	userContext := &UserContext{
		Name:           user.GetName(),
		ACL:            acl,
		AuthType:       authLocal,
		AccessStrategy: accessStrategy,
		PasswordSate:   user.GetPasswordState(),
	}

	if len(user.GetOIDCIdentities()) > 0 {
		userContext.AuthType = authSSO
		userContext.SSOContext = SSOContext{
			ConnectorType: constants.OIDC,
		}
	}

	if len(user.GetGithubIdentities()) > 0 {
		userContext.AuthType = authSSO
		userContext.SSOContext = SSOContext{
			ConnectorType: constants.Github,
		}
	}

	if len(user.GetSAMLIdentities()) > 0 {
		userContext.AuthType = authSSO
		userContext.SSOContext = SSOContext{
			ConnectorType:       constants.SAML,
			SAMLSingleLogoutURL: user.GetSAMLIdentities()[0].SAMLSingleLogoutURL,
		}
	}

	return userContext, nil
}
