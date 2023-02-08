// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/handles/v1/handles_v1.proto

package handles

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

// HandlesServiceClient is the client API for HandlesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandlesServiceClient interface {
	ListInstallationHandles(ctx context.Context, in *ListInstallationHandlesRequest, opts ...grpc.CallOption) (*ListInstallationHandlesResponse, error)
}

type handlesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHandlesServiceClient(cc grpc.ClientConnInterface) HandlesServiceClient {
	return &handlesServiceClient{cc}
}

func (c *handlesServiceClient) ListInstallationHandles(ctx context.Context, in *ListInstallationHandlesRequest, opts ...grpc.CallOption) (*ListInstallationHandlesResponse, error) {
	out := new(ListInstallationHandlesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.handles.v1.HandlesService/ListInstallationHandles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HandlesServiceServer is the server API for HandlesService service.
// All implementations should embed UnimplementedHandlesServiceServer
// for forward compatibility
type HandlesServiceServer interface {
	ListInstallationHandles(context.Context, *ListInstallationHandlesRequest) (*ListInstallationHandlesResponse, error)
}

// UnimplementedHandlesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedHandlesServiceServer struct {
}

func (UnimplementedHandlesServiceServer) ListInstallationHandles(context.Context, *ListInstallationHandlesRequest) (*ListInstallationHandlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstallationHandles not implemented")
}

// UnsafeHandlesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandlesServiceServer will
// result in compilation errors.
type UnsafeHandlesServiceServer interface {
	mustEmbedUnimplementedHandlesServiceServer()
}

func RegisterHandlesServiceServer(s grpc.ServiceRegistrar, srv HandlesServiceServer) {
	s.RegisterService(&HandlesService_ServiceDesc, srv)
}

func _HandlesService_ListInstallationHandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstallationHandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlesServiceServer).ListInstallationHandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.handles.v1.HandlesService/ListInstallationHandles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlesServiceServer).ListInstallationHandles(ctx, req.(*ListInstallationHandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HandlesService_ServiceDesc is the grpc.ServiceDesc for HandlesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HandlesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.handles.v1.HandlesService",
	HandlerType: (*HandlesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListInstallationHandles",
			Handler:    _HandlesService_ListInstallationHandles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/handles/v1/handles_v1.proto",
}
