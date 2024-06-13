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

package local

import (
	"context"

	machineidv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/machineid/v1"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/api/utils"
	"github.com/gravitational/teleport/lib/backend"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/teleport/lib/services/local/generic"
	"github.com/gravitational/trace"
)

const (
	botInstancePrefix = "bot_instance"
)

type BotInstanceService struct {
	service *generic.ServiceWrapper[*machineidv1.BotInstance]
}

func NewBotInstanceService(backend backend.Backend) (*BotInstanceService, error) {
	service, err := generic.NewServiceWrapper(backend,
		types.KindBotInstance,
		botInstancePrefix,
		services.MarshalBotInstance,
		services.UnmarshalBotInstance)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return &BotInstanceService{service: service}, nil
}

func (b *BotInstanceService) CreateBotInstance(ctx context.Context, instance *machineidv1.BotInstance) (*machineidv1.BotInstance, error) {
	if err := services.ValidateBotInstance(instance); err != nil {
		return nil, trace.Wrap(err)
	}

	if instance.Metadata != nil {
		return nil, trace.BadParameter("metadata should be nil; initial parameters should be set in status")
	}

	instance.Kind = types.KindBotInstance
	instance.Version = types.V1

	// TODO: should bot name + id be part of the spec?
	serviceWithPrefix := b.service.WithPrefix(instance.Status.BotName)
	created, err := serviceWithPrefix.CreateResource(ctx, instance)
	return created, trace.Wrap(err)
}

func (b *BotInstanceService) GetBotInstance(ctx context.Context, botName, instanceID string) (*machineidv1.BotInstance, error) {
	serviceWithPrefix := b.service.WithPrefix(botName)
	instance, err := serviceWithPrefix.GetResource(ctx, instanceID)
	return instance, trace.Wrap(err)
}

func (b *BotInstanceService) ListBotInstances(ctx context.Context, botName string, pageSize int, lastKey string) ([]*machineidv1.BotInstance, string, error) {
	serviceWithPrefix := b.service.WithPrefix(botName)
	r, nextToken, err := serviceWithPrefix.ListResources(ctx, pageSize, lastKey)
	return r, nextToken, trace.Wrap(err)
}

func (b *BotInstanceService) DeleteBotInstance(ctx context.Context, botName, instanceID string) error {
	serviceWithPrefix := b.service.WithPrefix(botName)
	return trace.Wrap(serviceWithPrefix.DeleteResource(ctx, instanceID))
}

// PatchBotInstance uses the supplied function to patch the bot instance
// matching the given (botName, instanceID) key and persists the patched
// resource. It will make multiple attempts if a `CompareFailed` error is
// raised.
func (b *BotInstanceService) PatchBotInstance(
	ctx context.Context,
	botName, instanceID string,
	updateFn func(*machineidv1.BotInstance) (*machineidv1.BotInstance, error),
) (*machineidv1.BotInstance, error) {
	const iterLimit = 3

	for i := 0; i < iterLimit; i++ {
		existing, err := b.GetBotInstance(ctx, botName, instanceID)
		if err != nil {
			return nil, trace.Wrap(err)
		}

		updated, err := updateFn(utils.CloneProtoMsg(existing))
		if err != nil {
			return nil, trace.Wrap(err)
		}

		switch {
		case updated.GetMetadata().GetName() != existing.GetMetadata().GetName():
			return nil, trace.BadParameter("metadata.name: cannot be patched")
		case updated.GetMetadata().GetRevision() != existing.GetMetadata().GetRevision():
			return nil, trace.BadParameter("metadata.revision: cannot be patched")
		}

		lease, err := b.service.ConditionalUpdateResource(ctx, updated)
		if err != nil {
			if trace.IsCompareFailed(err) {
				continue
			}

			return nil, trace.Wrap(err)
		}

		updated.GetMetadata().Revision = lease.GetMetadata().Revision
		return updated, nil
	}

	return nil, trace.CompareFailed("failed to update bot instance within %v iterations", iterLimit)
}
