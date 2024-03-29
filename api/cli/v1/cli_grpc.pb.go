// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/cli/v1/cli.proto

package cli

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

// CliServiceClient is the client API for CliService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CliServiceClient interface {
	GetRelease(ctx context.Context, in *GetReleaseRequest, opts ...grpc.CallOption) (*GetReleaseResponse, error)
}

type cliServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCliServiceClient(cc grpc.ClientConnInterface) CliServiceClient {
	return &cliServiceClient{cc}
}

func (c *cliServiceClient) GetRelease(ctx context.Context, in *GetReleaseRequest, opts ...grpc.CallOption) (*GetReleaseResponse, error) {
	out := new(GetReleaseResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.cli.v1.CliService/GetRelease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CliServiceServer is the server API for CliService service.
// All implementations should embed UnimplementedCliServiceServer
// for forward compatibility
type CliServiceServer interface {
	GetRelease(context.Context, *GetReleaseRequest) (*GetReleaseResponse, error)
}

// UnimplementedCliServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCliServiceServer struct {
}

func (UnimplementedCliServiceServer) GetRelease(context.Context, *GetReleaseRequest) (*GetReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelease not implemented")
}

// UnsafeCliServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CliServiceServer will
// result in compilation errors.
type UnsafeCliServiceServer interface {
	mustEmbedUnimplementedCliServiceServer()
}

func RegisterCliServiceServer(s grpc.ServiceRegistrar, srv CliServiceServer) {
	s.RegisterService(&CliService_ServiceDesc, srv)
}

func _CliService_GetRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).GetRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.cli.v1.CliService/GetRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).GetRelease(ctx, req.(*GetReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CliService_ServiceDesc is the grpc.ServiceDesc for CliService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CliService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.cli.v1.CliService",
	HandlerType: (*CliServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRelease",
			Handler:    _CliService_GetRelease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/cli/v1/cli.proto",
}
