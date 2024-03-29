// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/installations/v1/installed_components_v1.proto

package installations

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InstalledComponentsServiceClient is the client API for InstalledComponentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstalledComponentsServiceClient interface {
	// Authentication goes through the installations realm; a strm-installation-id is required in the metadata.
	UpdateInstalledComponentState(ctx context.Context, in *UpdateInstalledComponentStateRequest, opts ...grpc.CallOption) (*UpdateInstalledComponentStateResponse, error)
	// Authentication goes through the users realm; a strm-external-user-id is required in the metadata.
	GetInstalledComponent(ctx context.Context, in *GetInstalledComponentRequest, opts ...grpc.CallOption) (*GetInstalledComponentResponse, error)
	// Authentication goes through the users realm; a strm-external-user-id is required in the metadata.
	ListInstalledComponents(ctx context.Context, in *ListInstalledComponentsRequest, opts ...grpc.CallOption) (*ListInstalledComponentsResponse, error)
	// Authentication goes through the users realm; a strm-external-user-id is required in the metadata. Similar to list,
	// but only lists the state of the instances in the last x minutes
	ListInstalledComponentsCurrentStates(ctx context.Context, in *ListInstalledComponentsCurrentStatesRequest, opts ...grpc.CallOption) (*ListInstalledComponentsCurrentStatesResponse, error)
}

type installedComponentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstalledComponentsServiceClient(cc grpc.ClientConnInterface) InstalledComponentsServiceClient {
	return &installedComponentsServiceClient{cc}
}

func (c *installedComponentsServiceClient) UpdateInstalledComponentState(ctx context.Context, in *UpdateInstalledComponentStateRequest, opts ...grpc.CallOption) (*UpdateInstalledComponentStateResponse, error) {
	out := new(UpdateInstalledComponentStateResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.installations.v1.InstalledComponentsService/UpdateInstalledComponentState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installedComponentsServiceClient) GetInstalledComponent(ctx context.Context, in *GetInstalledComponentRequest, opts ...grpc.CallOption) (*GetInstalledComponentResponse, error) {
	out := new(GetInstalledComponentResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.installations.v1.InstalledComponentsService/GetInstalledComponent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installedComponentsServiceClient) ListInstalledComponents(ctx context.Context, in *ListInstalledComponentsRequest, opts ...grpc.CallOption) (*ListInstalledComponentsResponse, error) {
	out := new(ListInstalledComponentsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.installations.v1.InstalledComponentsService/ListInstalledComponents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *installedComponentsServiceClient) ListInstalledComponentsCurrentStates(ctx context.Context, in *ListInstalledComponentsCurrentStatesRequest, opts ...grpc.CallOption) (*ListInstalledComponentsCurrentStatesResponse, error) {
	out := new(ListInstalledComponentsCurrentStatesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.installations.v1.InstalledComponentsService/ListInstalledComponentsCurrentStates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstalledComponentsServiceServer is the server API for InstalledComponentsService service.
// All implementations should embed UnimplementedInstalledComponentsServiceServer
// for forward compatibility
type InstalledComponentsServiceServer interface {
	// Authentication goes through the installations realm; a strm-installation-id is required in the metadata.
	UpdateInstalledComponentState(context.Context, *UpdateInstalledComponentStateRequest) (*UpdateInstalledComponentStateResponse, error)
	// Authentication goes through the users realm; a strm-external-user-id is required in the metadata.
	GetInstalledComponent(context.Context, *GetInstalledComponentRequest) (*GetInstalledComponentResponse, error)
	// Authentication goes through the users realm; a strm-external-user-id is required in the metadata.
	ListInstalledComponents(context.Context, *ListInstalledComponentsRequest) (*ListInstalledComponentsResponse, error)
	// Authentication goes through the users realm; a strm-external-user-id is required in the metadata. Similar to list,
	// but only lists the state of the instances in the last x minutes
	ListInstalledComponentsCurrentStates(context.Context, *ListInstalledComponentsCurrentStatesRequest) (*ListInstalledComponentsCurrentStatesResponse, error)
}

// UnimplementedInstalledComponentsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInstalledComponentsServiceServer struct {
}

func (UnimplementedInstalledComponentsServiceServer) UpdateInstalledComponentState(context.Context, *UpdateInstalledComponentStateRequest) (*UpdateInstalledComponentStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstalledComponentState not implemented")
}
func (UnimplementedInstalledComponentsServiceServer) GetInstalledComponent(context.Context, *GetInstalledComponentRequest) (*GetInstalledComponentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstalledComponent not implemented")
}
func (UnimplementedInstalledComponentsServiceServer) ListInstalledComponents(context.Context, *ListInstalledComponentsRequest) (*ListInstalledComponentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstalledComponents not implemented")
}
func (UnimplementedInstalledComponentsServiceServer) ListInstalledComponentsCurrentStates(context.Context, *ListInstalledComponentsCurrentStatesRequest) (*ListInstalledComponentsCurrentStatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstalledComponentsCurrentStates not implemented")
}

// UnsafeInstalledComponentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstalledComponentsServiceServer will
// result in compilation errors.
type UnsafeInstalledComponentsServiceServer interface {
	mustEmbedUnimplementedInstalledComponentsServiceServer()
}

func RegisterInstalledComponentsServiceServer(s grpc.ServiceRegistrar, srv InstalledComponentsServiceServer) {
	s.RegisterService(&InstalledComponentsService_ServiceDesc, srv)
}

func _InstalledComponentsService_UpdateInstalledComponentState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstalledComponentStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstalledComponentsServiceServer).UpdateInstalledComponentState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.installations.v1.InstalledComponentsService/UpdateInstalledComponentState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstalledComponentsServiceServer).UpdateInstalledComponentState(ctx, req.(*UpdateInstalledComponentStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstalledComponentsService_GetInstalledComponent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstalledComponentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstalledComponentsServiceServer).GetInstalledComponent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.installations.v1.InstalledComponentsService/GetInstalledComponent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstalledComponentsServiceServer).GetInstalledComponent(ctx, req.(*GetInstalledComponentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstalledComponentsService_ListInstalledComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstalledComponentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstalledComponentsServiceServer).ListInstalledComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.installations.v1.InstalledComponentsService/ListInstalledComponents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstalledComponentsServiceServer).ListInstalledComponents(ctx, req.(*ListInstalledComponentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstalledComponentsService_ListInstalledComponentsCurrentStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstalledComponentsCurrentStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstalledComponentsServiceServer).ListInstalledComponentsCurrentStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.installations.v1.InstalledComponentsService/ListInstalledComponentsCurrentStates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstalledComponentsServiceServer).ListInstalledComponentsCurrentStates(ctx, req.(*ListInstalledComponentsCurrentStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InstalledComponentsService_ServiceDesc is the grpc.ServiceDesc for InstalledComponentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstalledComponentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.installations.v1.InstalledComponentsService",
	HandlerType: (*InstalledComponentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateInstalledComponentState",
			Handler:    _InstalledComponentsService_UpdateInstalledComponentState_Handler,
		},
		{
			MethodName: "GetInstalledComponent",
			Handler:    _InstalledComponentsService_GetInstalledComponent_Handler,
		},
		{
			MethodName: "ListInstalledComponents",
			Handler:    _InstalledComponentsService_ListInstalledComponents_Handler,
		},
		{
			MethodName: "ListInstalledComponentsCurrentStates",
			Handler:    _InstalledComponentsService_ListInstalledComponentsCurrentStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/installations/v1/installed_components_v1.proto",
}
