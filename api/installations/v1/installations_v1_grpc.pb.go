// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: strmprivacy/api/installations/v1/installations_v1.proto

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

// ComponentStatusServiceClient is the client API for ComponentStatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComponentStatusServiceClient interface {
	UpdateComponentStatus(ctx context.Context, in *UpdateComponentStatusRequest, opts ...grpc.CallOption) (*UpdateComponentStatusResponse, error)
}

type componentStatusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComponentStatusServiceClient(cc grpc.ClientConnInterface) ComponentStatusServiceClient {
	return &componentStatusServiceClient{cc}
}

func (c *componentStatusServiceClient) UpdateComponentStatus(ctx context.Context, in *UpdateComponentStatusRequest, opts ...grpc.CallOption) (*UpdateComponentStatusResponse, error) {
	out := new(UpdateComponentStatusResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.installations.v1.ComponentStatusService/UpdateComponentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentStatusServiceServer is the server API for ComponentStatusService service.
// All implementations must embed UnimplementedComponentStatusServiceServer
// for forward compatibility
type ComponentStatusServiceServer interface {
	UpdateComponentStatus(context.Context, *UpdateComponentStatusRequest) (*UpdateComponentStatusResponse, error)
	mustEmbedUnimplementedComponentStatusServiceServer()
}

// UnimplementedComponentStatusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComponentStatusServiceServer struct {
}

func (UnimplementedComponentStatusServiceServer) UpdateComponentStatus(context.Context, *UpdateComponentStatusRequest) (*UpdateComponentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComponentStatus not implemented")
}
func (UnimplementedComponentStatusServiceServer) mustEmbedUnimplementedComponentStatusServiceServer() {
}

// UnsafeComponentStatusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComponentStatusServiceServer will
// result in compilation errors.
type UnsafeComponentStatusServiceServer interface {
	mustEmbedUnimplementedComponentStatusServiceServer()
}

func RegisterComponentStatusServiceServer(s grpc.ServiceRegistrar, srv ComponentStatusServiceServer) {
	s.RegisterService(&ComponentStatusService_ServiceDesc, srv)
}

func _ComponentStatusService_UpdateComponentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateComponentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentStatusServiceServer).UpdateComponentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.installations.v1.ComponentStatusService/UpdateComponentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentStatusServiceServer).UpdateComponentStatus(ctx, req.(*UpdateComponentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComponentStatusService_ServiceDesc is the grpc.ServiceDesc for ComponentStatusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComponentStatusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.installations.v1.ComponentStatusService",
	HandlerType: (*ComponentStatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateComponentStatus",
			Handler:    _ComponentStatusService_UpdateComponentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/installations/v1/installations_v1.proto",
}
