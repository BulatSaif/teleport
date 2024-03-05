/* eslint-disable */
// @generated by protobuf-ts 2.9.3 with parameter long_type_number,eslint_disable,add_pb_suffix,ts_nocheck,server_generic
// @generated from protobuf file "teleport/devicetrust/v1/devicetrust_service.proto" (package "teleport.devicetrust.v1", syntax proto3)
// tslint:disable
// @ts-nocheck
//
// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
import { DevicesUsage } from "./usage_pb";
import { GetDevicesUsageRequest } from "./devicetrust_service_pb";
import { SyncInventoryResponse } from "./devicetrust_service_pb";
import { SyncInventoryRequest } from "./devicetrust_service_pb";
import { AuthenticateDeviceResponse } from "./devicetrust_service_pb";
import { AuthenticateDeviceRequest } from "./devicetrust_service_pb";
import { RpcInputStream } from "@protobuf-ts/runtime-rpc";
import { RpcOutputStream } from "@protobuf-ts/runtime-rpc";
import { EnrollDeviceResponse } from "./devicetrust_service_pb";
import { EnrollDeviceRequest } from "./devicetrust_service_pb";
import { DeviceEnrollToken } from "./device_enroll_token_pb";
import { CreateDeviceEnrollTokenRequest } from "./devicetrust_service_pb";
import { BulkCreateDevicesResponse } from "./devicetrust_service_pb";
import { BulkCreateDevicesRequest } from "./devicetrust_service_pb";
import { ListDevicesResponse } from "./devicetrust_service_pb";
import { ListDevicesRequest } from "./devicetrust_service_pb";
import { GetDeviceRequest } from "./devicetrust_service_pb";
import { FindDevicesResponse } from "./devicetrust_service_pb";
import { FindDevicesRequest } from "./devicetrust_service_pb";
import { Empty } from "../../../google/protobuf/empty_pb";
import { DeleteDeviceRequest } from "./devicetrust_service_pb";
import { UpsertDeviceRequest } from "./devicetrust_service_pb";
import { UpdateDeviceRequest } from "./devicetrust_service_pb";
import { Device } from "./device_pb";
import { CreateDeviceRequest } from "./devicetrust_service_pb";
import { ServerCallContext } from "@protobuf-ts/runtime-rpc";
/**
 * DeviceTrustService provides methods to manage, enroll and authenticate
 * trusted devices.
 *
 * A trusted device is a device that is registered and enrolled with Teleport,
 * thus allowing the system to provide some guarantees about its provenance and
 * state.
 *
 * Managing devices requires the corresponding CRUD "device" permission.
 * Additionally, creating enrollment tokens requires the "create_enroll_token"
 * permission and enrolling devices requires the "enroll" permission. See
 * CreateDevice, CreateDeviceEnrollToken and EnrollDevice for reference.
 *
 * An authenticated, trusted device allows its user to perform device-aware
 * actions. Such actions include accessing an SSH node, managing sensitive
 * resources via `tctl`, etc. The enforcement mode is defined via cluster-wide
 * and/or per-role toggles. Device authentication is automatic for enrolled
 * devices communicating with Enterprise clusters. See AuthenticateDevice for
 * reference.
 *
 * Device Trust is a Teleport Enterprise feature. Open Source Teleport clusters
 * treat all Device RPCs as unimplemented (which, in fact, they are for OSS.)
 *
 * @generated from protobuf service teleport.devicetrust.v1.DeviceTrustService
 */
export interface IDeviceTrustService<T = ServerCallContext> {
    /**
     * CreateDevice creates a device, effectively registering it on Teleport.
     * Devices need to be registered before they can be enrolled.
     *
     * It is possible to create both a Device and a DeviceEnrollToken in a
     * single invocation, see CreateDeviceRequest.create_enroll_token.
     *
     * @generated from protobuf rpc: CreateDevice(teleport.devicetrust.v1.CreateDeviceRequest) returns (teleport.devicetrust.v1.Device);
     */
    createDevice(request: CreateDeviceRequest, context: T): Promise<Device>;
    /**
     * UpdateDevice is a masked device update.
     *
     * Only certain fields may be updated, see Device for details.
     *
     * @generated from protobuf rpc: UpdateDevice(teleport.devicetrust.v1.UpdateDeviceRequest) returns (teleport.devicetrust.v1.Device);
     */
    updateDevice(request: UpdateDeviceRequest, context: T): Promise<Device>;
    /**
     * UpsertDevice creates or updates a device.
     *
     * UpsertDevice attempts a write of all mutable fields on updates, therefore
     * reading a fresh copy of the device is recommended. Update semantics still
     * apply.
     *
     * @generated from protobuf rpc: UpsertDevice(teleport.devicetrust.v1.UpsertDeviceRequest) returns (teleport.devicetrust.v1.Device);
     */
    upsertDevice(request: UpsertDeviceRequest, context: T): Promise<Device>;
    /**
     * DeleteDevice hard-deletes a device, removing it and all collected data
     * history from the system.
     *
     * Prefer locking the device instead (see the `tctl lock` command). Deleting a
     * device doesn't invalidate existing device certificates, but does prevent
     * new device authentication ceremonies from occurring.
     *
     * Use with caution.
     *
     * @generated from protobuf rpc: DeleteDevice(teleport.devicetrust.v1.DeleteDeviceRequest) returns (google.protobuf.Empty);
     */
    deleteDevice(request: DeleteDeviceRequest, context: T): Promise<Empty>;
    /**
     * FindDevices retrieves devices by device ID and/or asset tag.
     *
     * It provides an in-between search between fetching a device by ID and
     * listing all devices.
     *
     * ID matches are guaranteed to be present in the response.
     *
     * @generated from protobuf rpc: FindDevices(teleport.devicetrust.v1.FindDevicesRequest) returns (teleport.devicetrust.v1.FindDevicesResponse);
     */
    findDevices(request: FindDevicesRequest, context: T): Promise<FindDevicesResponse>;
    /**
     * GetDevice retrieves a device by ID.
     *
     * @generated from protobuf rpc: GetDevice(teleport.devicetrust.v1.GetDeviceRequest) returns (teleport.devicetrust.v1.Device);
     */
    getDevice(request: GetDeviceRequest, context: T): Promise<Device>;
    /**
     * ListDevices lists all registered devices.
     *
     * @generated from protobuf rpc: ListDevices(teleport.devicetrust.v1.ListDevicesRequest) returns (teleport.devicetrust.v1.ListDevicesResponse);
     */
    listDevices(request: ListDevicesRequest, context: T): Promise<ListDevicesResponse>;
    /**
     * BulkCreateDevices is a bulk variant of CreateDevice.
     *
     * Unlike CreateDevice, it does not support creation of enrollment tokens, as
     * it is meant for bulk inventory registration.
     *
     * @generated from protobuf rpc: BulkCreateDevices(teleport.devicetrust.v1.BulkCreateDevicesRequest) returns (teleport.devicetrust.v1.BulkCreateDevicesResponse);
     */
    bulkCreateDevices(request: BulkCreateDevicesRequest, context: T): Promise<BulkCreateDevicesResponse>;
    /**
     * CreateDeviceEnrollToken creates a DeviceEnrollToken for a Device.
     * An enrollment token is required for the enrollment ceremony. See
     * EnrollDevice.
     *
     * @generated from protobuf rpc: CreateDeviceEnrollToken(teleport.devicetrust.v1.CreateDeviceEnrollTokenRequest) returns (teleport.devicetrust.v1.DeviceEnrollToken);
     */
    createDeviceEnrollToken(request: CreateDeviceEnrollTokenRequest, context: T): Promise<DeviceEnrollToken>;
    /**
     * EnrollDevice performs the device enrollment ceremony.
     *
     * Enrollment requires a previously-registered Device and a DeviceEnrollToken,
     * see CreateDevice and CreateDeviceEnrollToken.
     *
     * An enrolled device is allowed, via AuthenticateDevice, to acquire
     * certificates containing device extensions, thus gaining access to
     * device-aware actions.
     *
     * macOS enrollment flow:
     * -> EnrollDeviceInit (client)
     * <- MacOSEnrollChallenge (server)
     * -> MacOSEnrollChallengeResponse
     * <- EnrollDeviceSuccess
     *
     * TPM enrollment flow:
     * -> EnrollDeviceInit (client)
     * <- TPMEnrollChallenge (server)
     * -> TPMEnrollChallengeResponse
     * <- EnrollDeviceSuccess
     *
     * @generated from protobuf rpc: EnrollDevice(stream teleport.devicetrust.v1.EnrollDeviceRequest) returns (stream teleport.devicetrust.v1.EnrollDeviceResponse);
     */
    enrollDevice(requests: RpcOutputStream<EnrollDeviceRequest>, responses: RpcInputStream<EnrollDeviceResponse>, context: T): Promise<void>;
    /**
     * AuthenticateDevice performs the device authentication ceremony.
     *
     * Device authentication exchanges existing user certificates without device
     * extensions for certificates augmented with device extensions. The new
     * certificates allow the user to perform device-aware actions.
     *
     * Only registered and enrolled devices may perform device authentication.
     *
     * @generated from protobuf rpc: AuthenticateDevice(stream teleport.devicetrust.v1.AuthenticateDeviceRequest) returns (stream teleport.devicetrust.v1.AuthenticateDeviceResponse);
     */
    authenticateDevice(requests: RpcOutputStream<AuthenticateDeviceRequest>, responses: RpcInputStream<AuthenticateDeviceResponse>, context: T): Promise<void>;
    /**
     * Syncs device inventory from a source exterior to Teleport, for example an
     * MDM.
     * Allows both partial and full syncs; for the latter, devices missing from
     * the external inventory are handled as specified.
     * Authorized either by a valid MDM service certificate or the appropriate
     * "device" permissions (create/update/delete).
     *
     * @generated from protobuf rpc: SyncInventory(stream teleport.devicetrust.v1.SyncInventoryRequest) returns (stream teleport.devicetrust.v1.SyncInventoryResponse);
     */
    syncInventory(requests: RpcOutputStream<SyncInventoryRequest>, responses: RpcInputStream<SyncInventoryResponse>, context: T): Promise<void>;
    /**
     * Superseded by ResourceUsageService.GetUsage.
     *
     * @deprecated
     * @generated from protobuf rpc: GetDevicesUsage(teleport.devicetrust.v1.GetDevicesUsageRequest) returns (teleport.devicetrust.v1.DevicesUsage);
     */
    getDevicesUsage(request: GetDevicesUsageRequest, context: T): Promise<DevicesUsage>;
}