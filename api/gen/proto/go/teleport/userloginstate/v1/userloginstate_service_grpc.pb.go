// Copyright 2023 Gravitational, Inc
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
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: teleport/userloginstate/v1/userloginstate_service.proto

package userloginstatev1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserLoginStateService_GetUserLoginStates_FullMethodName       = "/teleport.userloginstate.v1.UserLoginStateService/GetUserLoginStates"
	UserLoginStateService_GetUserLoginState_FullMethodName        = "/teleport.userloginstate.v1.UserLoginStateService/GetUserLoginState"
	UserLoginStateService_UpsertUserLoginState_FullMethodName     = "/teleport.userloginstate.v1.UserLoginStateService/UpsertUserLoginState"
	UserLoginStateService_DeleteUserLoginState_FullMethodName     = "/teleport.userloginstate.v1.UserLoginStateService/DeleteUserLoginState"
	UserLoginStateService_DeleteAllUserLoginStates_FullMethodName = "/teleport.userloginstate.v1.UserLoginStateService/DeleteAllUserLoginStates"
)

// UserLoginStateServiceClient is the client API for UserLoginStateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// UserLoginStateService provides CRUD methods for user login state resources.
type UserLoginStateServiceClient interface {
	// GetUserLoginStates returns a list of all user login states.
	GetUserLoginStates(ctx context.Context, in *GetUserLoginStatesRequest, opts ...grpc.CallOption) (*GetUserLoginStatesResponse, error)
	// GetUserLoginState returns the specified user login state resource.
	GetUserLoginState(ctx context.Context, in *GetUserLoginStateRequest, opts ...grpc.CallOption) (*UserLoginState, error)
	// UpsertUserLoginState creates or updates a user login state resource.
	UpsertUserLoginState(ctx context.Context, in *UpsertUserLoginStateRequest, opts ...grpc.CallOption) (*UserLoginState, error)
	// DeleteUserLoginState hard deletes the specified user login state resource.
	DeleteUserLoginState(ctx context.Context, in *DeleteUserLoginStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteAllUserLoginStates hard deletes all user login states.
	DeleteAllUserLoginStates(ctx context.Context, in *DeleteAllUserLoginStatesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userLoginStateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLoginStateServiceClient(cc grpc.ClientConnInterface) UserLoginStateServiceClient {
	return &userLoginStateServiceClient{cc}
}

func (c *userLoginStateServiceClient) GetUserLoginStates(ctx context.Context, in *GetUserLoginStatesRequest, opts ...grpc.CallOption) (*GetUserLoginStatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserLoginStatesResponse)
	err := c.cc.Invoke(ctx, UserLoginStateService_GetUserLoginStates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginStateServiceClient) GetUserLoginState(ctx context.Context, in *GetUserLoginStateRequest, opts ...grpc.CallOption) (*UserLoginState, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginState)
	err := c.cc.Invoke(ctx, UserLoginStateService_GetUserLoginState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginStateServiceClient) UpsertUserLoginState(ctx context.Context, in *UpsertUserLoginStateRequest, opts ...grpc.CallOption) (*UserLoginState, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserLoginState)
	err := c.cc.Invoke(ctx, UserLoginStateService_UpsertUserLoginState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginStateServiceClient) DeleteUserLoginState(ctx context.Context, in *DeleteUserLoginStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserLoginStateService_DeleteUserLoginState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userLoginStateServiceClient) DeleteAllUserLoginStates(ctx context.Context, in *DeleteAllUserLoginStatesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserLoginStateService_DeleteAllUserLoginStates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginStateServiceServer is the server API for UserLoginStateService service.
// All implementations must embed UnimplementedUserLoginStateServiceServer
// for forward compatibility
//
// UserLoginStateService provides CRUD methods for user login state resources.
type UserLoginStateServiceServer interface {
	// GetUserLoginStates returns a list of all user login states.
	GetUserLoginStates(context.Context, *GetUserLoginStatesRequest) (*GetUserLoginStatesResponse, error)
	// GetUserLoginState returns the specified user login state resource.
	GetUserLoginState(context.Context, *GetUserLoginStateRequest) (*UserLoginState, error)
	// UpsertUserLoginState creates or updates a user login state resource.
	UpsertUserLoginState(context.Context, *UpsertUserLoginStateRequest) (*UserLoginState, error)
	// DeleteUserLoginState hard deletes the specified user login state resource.
	DeleteUserLoginState(context.Context, *DeleteUserLoginStateRequest) (*emptypb.Empty, error)
	// DeleteAllUserLoginStates hard deletes all user login states.
	DeleteAllUserLoginStates(context.Context, *DeleteAllUserLoginStatesRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserLoginStateServiceServer()
}

// UnimplementedUserLoginStateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserLoginStateServiceServer struct {
}

func (UnimplementedUserLoginStateServiceServer) GetUserLoginStates(context.Context, *GetUserLoginStatesRequest) (*GetUserLoginStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLoginStates not implemented")
}
func (UnimplementedUserLoginStateServiceServer) GetUserLoginState(context.Context, *GetUserLoginStateRequest) (*UserLoginState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLoginState not implemented")
}
func (UnimplementedUserLoginStateServiceServer) UpsertUserLoginState(context.Context, *UpsertUserLoginStateRequest) (*UserLoginState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertUserLoginState not implemented")
}
func (UnimplementedUserLoginStateServiceServer) DeleteUserLoginState(context.Context, *DeleteUserLoginStateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserLoginState not implemented")
}
func (UnimplementedUserLoginStateServiceServer) DeleteAllUserLoginStates(context.Context, *DeleteAllUserLoginStatesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllUserLoginStates not implemented")
}
func (UnimplementedUserLoginStateServiceServer) mustEmbedUnimplementedUserLoginStateServiceServer() {}

// UnsafeUserLoginStateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLoginStateServiceServer will
// result in compilation errors.
type UnsafeUserLoginStateServiceServer interface {
	mustEmbedUnimplementedUserLoginStateServiceServer()
}

func RegisterUserLoginStateServiceServer(s grpc.ServiceRegistrar, srv UserLoginStateServiceServer) {
	s.RegisterService(&UserLoginStateService_ServiceDesc, srv)
}

func _UserLoginStateService_GetUserLoginStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserLoginStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginStateServiceServer).GetUserLoginStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginStateService_GetUserLoginStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginStateServiceServer).GetUserLoginStates(ctx, req.(*GetUserLoginStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginStateService_GetUserLoginState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserLoginStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginStateServiceServer).GetUserLoginState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginStateService_GetUserLoginState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginStateServiceServer).GetUserLoginState(ctx, req.(*GetUserLoginStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginStateService_UpsertUserLoginState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserLoginStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginStateServiceServer).UpsertUserLoginState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginStateService_UpsertUserLoginState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginStateServiceServer).UpsertUserLoginState(ctx, req.(*UpsertUserLoginStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginStateService_DeleteUserLoginState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserLoginStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginStateServiceServer).DeleteUserLoginState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginStateService_DeleteUserLoginState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginStateServiceServer).DeleteUserLoginState(ctx, req.(*DeleteUserLoginStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserLoginStateService_DeleteAllUserLoginStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllUserLoginStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginStateServiceServer).DeleteAllUserLoginStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserLoginStateService_DeleteAllUserLoginStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginStateServiceServer).DeleteAllUserLoginStates(ctx, req.(*DeleteAllUserLoginStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLoginStateService_ServiceDesc is the grpc.ServiceDesc for UserLoginStateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLoginStateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.userloginstate.v1.UserLoginStateService",
	HandlerType: (*UserLoginStateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserLoginStates",
			Handler:    _UserLoginStateService_GetUserLoginStates_Handler,
		},
		{
			MethodName: "GetUserLoginState",
			Handler:    _UserLoginStateService_GetUserLoginState_Handler,
		},
		{
			MethodName: "UpsertUserLoginState",
			Handler:    _UserLoginStateService_UpsertUserLoginState_Handler,
		},
		{
			MethodName: "DeleteUserLoginState",
			Handler:    _UserLoginStateService_DeleteUserLoginState_Handler,
		},
		{
			MethodName: "DeleteAllUserLoginStates",
			Handler:    _UserLoginStateService_DeleteAllUserLoginStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/userloginstate/v1/userloginstate_service.proto",
}
