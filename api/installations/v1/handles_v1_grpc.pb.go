// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: strmprivacy/api/installations/v1/handles_v1.proto

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

// InstallationHandlesServiceClient is the client API for InstallationHandlesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstallationHandlesServiceClient interface {
	ListInstallationHandles(ctx context.Context, in *ListInstallationHandlesRequest, opts ...grpc.CallOption) (*ListInstallationHandlesResponse, error)
}

type installationHandlesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstallationHandlesServiceClient(cc grpc.ClientConnInterface) InstallationHandlesServiceClient {
	return &installationHandlesServiceClient{cc}
}

func (c *installationHandlesServiceClient) ListInstallationHandles(ctx context.Context, in *ListInstallationHandlesRequest, opts ...grpc.CallOption) (*ListInstallationHandlesResponse, error) {
	out := new(ListInstallationHandlesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.installations.v1.InstallationHandlesService/ListInstallationHandles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstallationHandlesServiceServer is the server API for InstallationHandlesService service.
// All implementations must embed UnimplementedInstallationHandlesServiceServer
// for forward compatibility
type InstallationHandlesServiceServer interface {
	ListInstallationHandles(context.Context, *ListInstallationHandlesRequest) (*ListInstallationHandlesResponse, error)
	mustEmbedUnimplementedInstallationHandlesServiceServer()
}

// UnimplementedInstallationHandlesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInstallationHandlesServiceServer struct {
}

func (UnimplementedInstallationHandlesServiceServer) ListInstallationHandles(context.Context, *ListInstallationHandlesRequest) (*ListInstallationHandlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstallationHandles not implemented")
}
func (UnimplementedInstallationHandlesServiceServer) mustEmbedUnimplementedInstallationHandlesServiceServer() {
}

// UnsafeInstallationHandlesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstallationHandlesServiceServer will
// result in compilation errors.
type UnsafeInstallationHandlesServiceServer interface {
	mustEmbedUnimplementedInstallationHandlesServiceServer()
}

func RegisterInstallationHandlesServiceServer(s grpc.ServiceRegistrar, srv InstallationHandlesServiceServer) {
	s.RegisterService(&InstallationHandlesService_ServiceDesc, srv)
}

func _InstallationHandlesService_ListInstallationHandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstallationHandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstallationHandlesServiceServer).ListInstallationHandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.installations.v1.InstallationHandlesService/ListInstallationHandles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstallationHandlesServiceServer).ListInstallationHandles(ctx, req.(*ListInstallationHandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InstallationHandlesService_ServiceDesc is the grpc.ServiceDesc for InstallationHandlesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstallationHandlesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.installations.v1.InstallationHandlesService",
	HandlerType: (*InstallationHandlesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInstallationHandles",
			Handler:    _InstallationHandlesService_ListInstallationHandles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/installations/v1/handles_v1.proto",
}