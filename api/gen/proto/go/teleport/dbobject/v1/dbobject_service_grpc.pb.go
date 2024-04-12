// Copyright 2024 Gravitational, Inc
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/dbobject/v1/dbobject_service.proto

package dbobjectv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DatabaseObjectService_GetDatabaseObject_FullMethodName    = "/teleport.dbobject.v1.DatabaseObjectService/GetDatabaseObject"
	DatabaseObjectService_ListDatabaseObjects_FullMethodName  = "/teleport.dbobject.v1.DatabaseObjectService/ListDatabaseObjects"
	DatabaseObjectService_CreateDatabaseObject_FullMethodName = "/teleport.dbobject.v1.DatabaseObjectService/CreateDatabaseObject"
	DatabaseObjectService_UpdateDatabaseObject_FullMethodName = "/teleport.dbobject.v1.DatabaseObjectService/UpdateDatabaseObject"
	DatabaseObjectService_UpsertDatabaseObject_FullMethodName = "/teleport.dbobject.v1.DatabaseObjectService/UpsertDatabaseObject"
	DatabaseObjectService_DeleteDatabaseObject_FullMethodName = "/teleport.dbobject.v1.DatabaseObjectService/DeleteDatabaseObject"
)

// DatabaseObjectServiceClient is the client API for DatabaseObjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseObjectServiceClient interface {
	// GetDatabaseObject is used to query a database object resource by its name.
	//
	// This will return a NotFound error if the specified database object does not exist.
	GetDatabaseObject(ctx context.Context, in *GetDatabaseObjectRequest, opts ...grpc.CallOption) (*DatabaseObject, error)
	// ListDatabaseObjects is used to query database objects.
	//
	// Follows the pagination semantics of
	// https://cloud.google.com/apis/design/standard_methods#list.
	ListDatabaseObjects(ctx context.Context, in *ListDatabaseObjectsRequest, opts ...grpc.CallOption) (*ListDatabaseObjectsResponse, error)
	// CreateDatabaseObject is used to create a database object.
	//
	// This will return an error if a database object by that name already exists.
	CreateDatabaseObject(ctx context.Context, in *CreateDatabaseObjectRequest, opts ...grpc.CallOption) (*DatabaseObject, error)
	// UpdateDatabaseObject is used to modify an existing database object.
	UpdateDatabaseObject(ctx context.Context, in *UpdateDatabaseObjectRequest, opts ...grpc.CallOption) (*DatabaseObject, error)
	// UpsertDatabaseObject is used to create or replace an existing database object.
	//
	// Prefer using CreateDatabaseObject and UpdateDatabaseObject.
	UpsertDatabaseObject(ctx context.Context, in *UpsertDatabaseObjectRequest, opts ...grpc.CallOption) (*DatabaseObject, error)
	// DeleteDatabaseObject is used to delete a specific database object.
	//
	// This will return a NotFound error if the specified database object does not exist.
	DeleteDatabaseObject(ctx context.Context, in *DeleteDatabaseObjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type databaseObjectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseObjectServiceClient(cc grpc.ClientConnInterface) DatabaseObjectServiceClient {
	return &databaseObjectServiceClient{cc}
}

func (c *databaseObjectServiceClient) GetDatabaseObject(ctx context.Context, in *GetDatabaseObjectRequest, opts ...grpc.CallOption) (*DatabaseObject, error) {
	out := new(DatabaseObject)
	err := c.cc.Invoke(ctx, DatabaseObjectService_GetDatabaseObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseObjectServiceClient) ListDatabaseObjects(ctx context.Context, in *ListDatabaseObjectsRequest, opts ...grpc.CallOption) (*ListDatabaseObjectsResponse, error) {
	out := new(ListDatabaseObjectsResponse)
	err := c.cc.Invoke(ctx, DatabaseObjectService_ListDatabaseObjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseObjectServiceClient) CreateDatabaseObject(ctx context.Context, in *CreateDatabaseObjectRequest, opts ...grpc.CallOption) (*DatabaseObject, error) {
	out := new(DatabaseObject)
	err := c.cc.Invoke(ctx, DatabaseObjectService_CreateDatabaseObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseObjectServiceClient) UpdateDatabaseObject(ctx context.Context, in *UpdateDatabaseObjectRequest, opts ...grpc.CallOption) (*DatabaseObject, error) {
	out := new(DatabaseObject)
	err := c.cc.Invoke(ctx, DatabaseObjectService_UpdateDatabaseObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseObjectServiceClient) UpsertDatabaseObject(ctx context.Context, in *UpsertDatabaseObjectRequest, opts ...grpc.CallOption) (*DatabaseObject, error) {
	out := new(DatabaseObject)
	err := c.cc.Invoke(ctx, DatabaseObjectService_UpsertDatabaseObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseObjectServiceClient) DeleteDatabaseObject(ctx context.Context, in *DeleteDatabaseObjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DatabaseObjectService_DeleteDatabaseObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseObjectServiceServer is the server API for DatabaseObjectService service.
// All implementations must embed UnimplementedDatabaseObjectServiceServer
// for forward compatibility
type DatabaseObjectServiceServer interface {
	// GetDatabaseObject is used to query a database object resource by its name.
	//
	// This will return a NotFound error if the specified database object does not exist.
	GetDatabaseObject(context.Context, *GetDatabaseObjectRequest) (*DatabaseObject, error)
	// ListDatabaseObjects is used to query database objects.
	//
	// Follows the pagination semantics of
	// https://cloud.google.com/apis/design/standard_methods#list.
	ListDatabaseObjects(context.Context, *ListDatabaseObjectsRequest) (*ListDatabaseObjectsResponse, error)
	// CreateDatabaseObject is used to create a database object.
	//
	// This will return an error if a database object by that name already exists.
	CreateDatabaseObject(context.Context, *CreateDatabaseObjectRequest) (*DatabaseObject, error)
	// UpdateDatabaseObject is used to modify an existing database object.
	UpdateDatabaseObject(context.Context, *UpdateDatabaseObjectRequest) (*DatabaseObject, error)
	// UpsertDatabaseObject is used to create or replace an existing database object.
	//
	// Prefer using CreateDatabaseObject and UpdateDatabaseObject.
	UpsertDatabaseObject(context.Context, *UpsertDatabaseObjectRequest) (*DatabaseObject, error)
	// DeleteDatabaseObject is used to delete a specific database object.
	//
	// This will return a NotFound error if the specified database object does not exist.
	DeleteDatabaseObject(context.Context, *DeleteDatabaseObjectRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDatabaseObjectServiceServer()
}

// UnimplementedDatabaseObjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseObjectServiceServer struct {
}

func (UnimplementedDatabaseObjectServiceServer) GetDatabaseObject(context.Context, *GetDatabaseObjectRequest) (*DatabaseObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatabaseObject not implemented")
}
func (UnimplementedDatabaseObjectServiceServer) ListDatabaseObjects(context.Context, *ListDatabaseObjectsRequest) (*ListDatabaseObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatabaseObjects not implemented")
}
func (UnimplementedDatabaseObjectServiceServer) CreateDatabaseObject(context.Context, *CreateDatabaseObjectRequest) (*DatabaseObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatabaseObject not implemented")
}
func (UnimplementedDatabaseObjectServiceServer) UpdateDatabaseObject(context.Context, *UpdateDatabaseObjectRequest) (*DatabaseObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDatabaseObject not implemented")
}
func (UnimplementedDatabaseObjectServiceServer) UpsertDatabaseObject(context.Context, *UpsertDatabaseObjectRequest) (*DatabaseObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDatabaseObject not implemented")
}
func (UnimplementedDatabaseObjectServiceServer) DeleteDatabaseObject(context.Context, *DeleteDatabaseObjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatabaseObject not implemented")
}
func (UnimplementedDatabaseObjectServiceServer) mustEmbedUnimplementedDatabaseObjectServiceServer() {}

// UnsafeDatabaseObjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseObjectServiceServer will
// result in compilation errors.
type UnsafeDatabaseObjectServiceServer interface {
	mustEmbedUnimplementedDatabaseObjectServiceServer()
}

func RegisterDatabaseObjectServiceServer(s grpc.ServiceRegistrar, srv DatabaseObjectServiceServer) {
	s.RegisterService(&DatabaseObjectService_ServiceDesc, srv)
}

func _DatabaseObjectService_GetDatabaseObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabaseObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseObjectServiceServer).GetDatabaseObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseObjectService_GetDatabaseObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseObjectServiceServer).GetDatabaseObject(ctx, req.(*GetDatabaseObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseObjectService_ListDatabaseObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatabaseObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseObjectServiceServer).ListDatabaseObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseObjectService_ListDatabaseObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseObjectServiceServer).ListDatabaseObjects(ctx, req.(*ListDatabaseObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseObjectService_CreateDatabaseObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatabaseObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseObjectServiceServer).CreateDatabaseObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseObjectService_CreateDatabaseObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseObjectServiceServer).CreateDatabaseObject(ctx, req.(*CreateDatabaseObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseObjectService_UpdateDatabaseObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatabaseObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseObjectServiceServer).UpdateDatabaseObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseObjectService_UpdateDatabaseObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseObjectServiceServer).UpdateDatabaseObject(ctx, req.(*UpdateDatabaseObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseObjectService_UpsertDatabaseObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDatabaseObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseObjectServiceServer).UpsertDatabaseObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseObjectService_UpsertDatabaseObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseObjectServiceServer).UpsertDatabaseObject(ctx, req.(*UpsertDatabaseObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseObjectService_DeleteDatabaseObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatabaseObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseObjectServiceServer).DeleteDatabaseObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseObjectService_DeleteDatabaseObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseObjectServiceServer).DeleteDatabaseObject(ctx, req.(*DeleteDatabaseObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseObjectService_ServiceDesc is the grpc.ServiceDesc for DatabaseObjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseObjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.dbobject.v1.DatabaseObjectService",
	HandlerType: (*DatabaseObjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDatabaseObject",
			Handler:    _DatabaseObjectService_GetDatabaseObject_Handler,
		},
		{
			MethodName: "ListDatabaseObjects",
			Handler:    _DatabaseObjectService_ListDatabaseObjects_Handler,
		},
		{
			MethodName: "CreateDatabaseObject",
			Handler:    _DatabaseObjectService_CreateDatabaseObject_Handler,
		},
		{
			MethodName: "UpdateDatabaseObject",
			Handler:    _DatabaseObjectService_UpdateDatabaseObject_Handler,
		},
		{
			MethodName: "UpsertDatabaseObject",
			Handler:    _DatabaseObjectService_UpsertDatabaseObject_Handler,
		},
		{
			MethodName: "DeleteDatabaseObject",
			Handler:    _DatabaseObjectService_DeleteDatabaseObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/dbobject/v1/dbobject_service.proto",
}