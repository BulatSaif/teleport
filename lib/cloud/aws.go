/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
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

package cloud

import (
	"context"
	"log/slog"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/gravitational/trace"

	cloudaws "github.com/gravitational/teleport/lib/cloud/aws"
	"github.com/gravitational/teleport/lib/modules"
	"github.com/gravitational/teleport/lib/utils"
)

// AWSIntegrationConfigV2Provider defines a function that creates an [awsconfig.Session] from a Region and an Integration.
// This is used to generate aws sessions for clients that must use an Integration instead of ambient credentials.
type AWSIntegrationConfigV2Provider func(ctx context.Context, region string, integration string) (*aws.Config, error)

// AWSV2Clients is an interface for providing AWS API clients.
type AWSClientsV2 interface {
	// GetAWSConfigV2 returns AWS session for the specified region, optionally
	// assuming AWS IAM Roles.
	GetAWSConfigV2(ctx context.Context, region string, opts ...AWSOptionsFn) (*aws.Config, error)
}

type awsClientsV2 struct {
	// sessionsCache is a cache of AWS sessions, where the cache key is
	// an instance of awsSessionCacheKey.
	sessionsCache *utils.FnCache
	// integrationConfigProviderFn is a AWS Session Generator that uses an Integration to generate an AWS Session.
	// TODO provide function to populate this. Currently it is always nil.
	integrationConfigProviderFn AWSIntegrationConfigV2Provider
	// stsEndpointResolver resolves endpoint for STS calls.
	stsEndpointResolver sts.EndpointResolverV2
}

func newAWSClientsV2() (*awsClientsV2, error) {
	sessionsCache, err := utils.NewFnCache(utils.FnCacheConfig{
		TTL: 15 * time.Minute,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &awsClientsV2{
		sessionsCache: sessionsCache,
		stsEndpointResolver: stsEndpointResolver{
			defaultResolver: sts.NewDefaultEndpointResolverV2(),
		},
	}, nil
}

func (c *awsClientsV2) GetAWSConfigV2(ctx context.Context, region string, opts ...AWSOptionsFn) (*aws.Config, error) {
	var options awsOptions
	for _, opt := range opts {
		opt(&options)
	}
	var err error
	if options.baseConfigV2 == nil {
		options.baseConfigV2, err = c.getAWSConfigV2ForRegion(ctx, region, options)
		if err != nil {
			return nil, trace.Wrap(err)
		}
	}
	if options.assumeRoleARN == "" {
		return options.baseConfigV2, nil
	}
	return c.getAWSConfigV2ForRole(ctx, region, options)
}

// getAWSConfigV2ForRegion returns AWS config for the specified region.
func (c *awsClientsV2) getAWSConfigV2ForRegion(ctx context.Context, region string, opts awsOptions) (*aws.Config, error) {
	if err := opts.checkAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	cacheKey := awsSessionCacheKey{
		region:      region,
		integration: opts.integration,
	}

	config, err := utils.FnCacheGet(ctx, c.sessionsCache, cacheKey, func(ctx context.Context) (*aws.Config, error) {
		if opts.credentialsSource == credentialsSourceIntegration {
			if c.integrationConfigProviderFn == nil {
				return nil, trace.BadParameter("missing aws integration config provider")
			}

			slog.DebugContext(ctx, "Initializing AWS config for integration.", "region", region, "integration", opts.integration)
			config, err := c.integrationConfigProviderFn(ctx, region, opts.integration)
			return config, trace.Wrap(err)
		}

		slog.DebugContext(ctx, "Initializing AWS config using ambient credentials.", "region", region)
		config, err := awsAmbientConfigV2Provider(ctx, region, nil /*credProvider*/)
		return config, trace.Wrap(err)
	})
	// TODO handle opts.customRetryer and opts.maxRetries
	return config, trace.Wrap(err)
}

// getAWSConfigV2ForRole returns AWS session for the specified region and role.
func (c *awsClientsV2) getAWSConfigV2ForRole(ctx context.Context, region string, options awsOptions) (*aws.Config, error) {
	if err := options.checkAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	if options.baseConfigV2 == nil {
		return nil, trace.BadParameter("missing base config")
	}

	cacheKey := awsSessionCacheKey{
		region:      region,
		integration: options.integration,
		roleARN:     options.assumeRoleARN,
		externalID:  options.assumeRoleExternalID,
	}
	return utils.FnCacheGet(ctx, c.sessionsCache, cacheKey, func(ctx context.Context) (*aws.Config, error) {
		stsClient := sts.NewFromConfig(*options.baseConfigV2, sts.WithEndpointResolverV2(c.stsEndpointResolver))
		provider := stscreds.NewAssumeRoleProvider(stsClient, options.assumeRoleARN, func(o *stscreds.AssumeRoleOptions) {
			if options.assumeRoleExternalID != "" {
				o.ExternalID = aws.String(options.assumeRoleExternalID)
			}
		})

		if _, err := provider.Retrieve(ctx); err != nil {
			err := cloudaws.ConvertSTSError(err)
			return nil, trace.Wrap(err)
		}

		config, err := awsAmbientConfigV2Provider(ctx, region, provider)
		return config, trace.Wrap(err)
	})
}

type stsEndpointResolver struct {
	defaultResolver sts.EndpointResolverV2
}

func (r stsEndpointResolver) ResolveEndpoint(ctx context.Context, params sts.EndpointParameters) (smithyendpoints.Endpoint, error) {
	// Use global endpoints when region is empty.
	if aws.ToString(params.Region) == "" {
		// A region is still required for passing basic validation and selecting
		// partition.
		params.Region = aws.String("us-east-1")
		params.UseGlobalEndpoint = aws.Bool(true)
		slog.DebugContext(ctx, "No region provided when resolving STS endpoint. Falling back to global endpoint.")
	}
	endpoint, err := r.defaultResolver.ResolveEndpoint(ctx, params)
	return endpoint, trace.Wrap(err)
}

// awsAmbientConfigV2Provider loads a new session using the environment variables.
func awsAmbientConfigV2Provider(ctx context.Context, region string, credProvider aws.CredentialsProvider) (*aws.Config, error) {
	opts := []func(*awsconfig.LoadOptions) error{
		awsConfigFipsOption(),
	}
	if region != "" {
		opts = append(opts, awsconfig.WithRegion(region))
	}
	if credProvider != nil {
		opts = append(opts, awsconfig.WithCredentialsProvider(credProvider))
	}
	config, err := awsconfig.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &config, nil
}

func awsConfigFipsOption() awsconfig.LoadOptionsFunc {
	if modules.GetModules().IsBoringBinary() {
		return awsconfig.WithUseFIPSEndpoint(aws.FIPSEndpointStateEnabled)
	}
	return awsconfig.WithUseFIPSEndpoint(aws.FIPSEndpointStateUnset)
}
