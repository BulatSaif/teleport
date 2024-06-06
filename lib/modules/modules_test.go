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

package modules_test

import (
	"context"
	"os"
	"testing"

	"github.com/gravitational/teleport/api/client/proto"
	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/api/constants"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/auth"
	"github.com/gravitational/teleport/lib/modules"
)

func TestMain(m *testing.M) {
	modules.SetInsecureTestMode(true)
	os.Exit(m.Run())
}

func TestOSSModules(t *testing.T) {
	require.False(t, modules.GetModules().IsBoringBinary())
	require.Equal(t, modules.BuildOSS, modules.GetModules().BuildType())
}

func TestValidateAuthPreferenceOnCloud(t *testing.T) {
	ctx := context.Background()
	testServer, err := auth.NewTestAuthServer(auth.TestAuthServerConfig{
		Dir: t.TempDir(),
	})
	require.NoError(t, err)

	modules.SetTestModules(t, &modules.TestModules{
		TestBuildType: modules.BuildEnterprise,
		TestFeatures: modules.Features{
			Cloud: true,
		},
	})

	authPref, err := testServer.AuthServer.UpsertAuthPreference(ctx, types.DefaultAuthPreference())
	require.NoError(t, err)

	authPref.SetSecondFactor(constants.SecondFactorOff)
	_, err = testServer.AuthServer.UpdateAuthPreference(ctx, authPref)
	require.EqualError(t, err, modules.ErrCannotDisableSecondFactor.Error())
}

func TestValidateSessionRecordingConfigOnCloud(t *testing.T) {
	ctx := context.Background()

	testServer, err := auth.NewTestAuthServer(auth.TestAuthServerConfig{
		Dir: t.TempDir(),
	})
	require.NoError(t, err)

	modules.SetTestModules(t, &modules.TestModules{
		TestBuildType: modules.BuildEnterprise,
		TestFeatures: modules.Features{
			Cloud: true,
		},
	})

	recConfig := types.DefaultSessionRecordingConfig()
	_, err = testServer.AuthServer.UpsertSessionRecordingConfig(ctx, recConfig)
	require.NoError(t, err)

	recConfig.SetMode(types.RecordAtProxy)
	_, err = testServer.AuthServer.UpsertSessionRecordingConfig(ctx, recConfig)
	require.EqualError(t, err, "cannot set proxy recording mode on Cloud")
}

func TestFeatures_ToProto(t *testing.T) {
	expected := &proto.Features{
		Assist:                  false,
		CustomTheme:             "dark",
		ProductType:             1,
		SupportType:             1,
		AccessControls:          true,
		AccessGraph:             true,
		AdvancedAccessWorkflows: true,
		App:                     true,
		AutomaticUpgrades:       true,
		Cloud:                   true,
		DB:                      true,
		Desktop:                 true,
		ExternalAuditStorage:    true,
		FeatureHiding:           true,
		HSM:                     true,
		IdentityGovernance:      true,
		IsStripeManaged:         true,
		IsUsageBased:            true,
		JoinActiveSessions:      true,
		Kubernetes:              true,
		MobileDeviceManagement:  true,
		OIDC:                    true,
		Plugins:                 true,
		Questionnaire:           true,
		RecoveryCodes:           true,
		SAML:                    true,

		AccessList:       &proto.AccessListFeature{CreateLimit: 111},
		AccessMonitoring: &proto.AccessMonitoringFeature{Enabled: true, MaxReportRangeLimit: 2113},
		AccessRequests:   &proto.AccessRequestsFeature{MonthlyRequestLimit: 39},
		DeviceTrust:      &proto.DeviceTrustFeature{Enabled: true, DevicesUsageLimit: 103},
		Policy:           &proto.PolicyFeature{Enabled: true},
	}

	f := modules.Features{
		CustomTheme:             "dark",
		ProductType:             1,
		SupportType:             1,
		AccessControls:          true,
		AccessGraph:             true,
		AdvancedAccessWorkflows: true,
		Assist:                  true,
		AutomaticUpgrades:       true,
		Cloud:                   true,
		IsStripeManaged:         true,
		IsUsageBasedBilling:     true,
		Plugins:                 true,
		Questionnaire:           true,
		RecoveryCodes:           true,
		Entitlements: map[teleport.EntitlementKind]modules.EntitlementInfo{
			teleport.AccessLists:            {Enabled: true, Limit: 111},
			teleport.AccessMonitoring:       {Enabled: true, Limit: 2113},
			teleport.AccessRequests:         {Enabled: true, Limit: 39},
			teleport.App:                    {Enabled: true, Limit: 3},
			teleport.CloudAuditLogRetention: {Enabled: true, Limit: 3},
			teleport.DB:                     {Enabled: true, Limit: 3},
			teleport.Desktop:                {Enabled: true, Limit: 3},
			teleport.DeviceTrust:            {Enabled: true, Limit: 103},
			teleport.ExternalAuditStorage:   {Enabled: true, Limit: 3},
			teleport.FeatureHiding:          {Enabled: true, Limit: 3},
			teleport.HSM:                    {Enabled: true, Limit: 3},
			teleport.Identity:               {Enabled: true, Limit: 3},
			teleport.JoinActiveSessions:     {Enabled: true, Limit: 3},
			teleport.K8s:                    {Enabled: true, Limit: 3},
			teleport.MobileDeviceManagement: {Enabled: true, Limit: 3},
			teleport.OIDC:                   {Enabled: true, Limit: 3},
			teleport.OktaSCIM:               {Enabled: true, Limit: 3},
			teleport.OktaUserSync:           {Enabled: true, Limit: 3},
			teleport.Policy:                 {Enabled: true, Limit: 3},
			teleport.SAML:                   {Enabled: true, Limit: 3},
			teleport.SessionLocks:           {Enabled: true, Limit: 3},
			teleport.UpsellAlert:            {Enabled: true, Limit: 3},
			teleport.UsageReporting:         {Enabled: true, Limit: 3},
		},
	}

	actual := f.ToProto()
	require.Equal(t, expected, actual)
}

func TestFeatures_GetEntitlement(t *testing.T) {
	f := modules.Features{
		Entitlements: map[teleport.EntitlementKind]modules.EntitlementInfo{
			teleport.AccessLists: {Enabled: true, Limit: 111},
			teleport.K8s:         {Enabled: false},
			teleport.SAML:        {},
		},
	}

	actual := f.GetEntitlement(teleport.AccessLists)
	require.Equal(t, modules.EntitlementInfo{Enabled: true, Limit: 111}, actual)

	actual = f.GetEntitlement(teleport.K8s)
	require.Equal(t, modules.EntitlementInfo{Enabled: false}, actual)

	actual = f.GetEntitlement(teleport.SAML)
	require.Equal(t, modules.EntitlementInfo{}, actual)

	actual = f.GetEntitlement(teleport.UsageReporting)
	require.Equal(t, modules.EntitlementInfo{}, actual)
}
