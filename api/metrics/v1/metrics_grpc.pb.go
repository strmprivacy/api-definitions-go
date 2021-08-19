// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package metrics

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

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsServiceClient interface {
	GetSystemState(ctx context.Context, in *GetSystemStateRequest, opts ...grpc.CallOption) (*GetSystemStateResponse, error)
	GetLatency(ctx context.Context, in *GetLatencyRequest, opts ...grpc.CallOption) (*GetLatencyResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) GetSystemState(ctx context.Context, in *GetSystemStateRequest, opts ...grpc.CallOption) (*GetSystemStateResponse, error) {
	out := new(GetSystemStateResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.metrics.v1.MetricsService/GetSystemState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetLatency(ctx context.Context, in *GetLatencyRequest, opts ...grpc.CallOption) (*GetLatencyResponse, error) {
	out := new(GetLatencyResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.metrics.v1.MetricsService/GetLatency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
// All implementations must embed UnimplementedMetricsServiceServer
// for forward compatibility
type MetricsServiceServer interface {
	GetSystemState(context.Context, *GetSystemStateRequest) (*GetSystemStateResponse, error)
	GetLatency(context.Context, *GetLatencyRequest) (*GetLatencyResponse, error)
	mustEmbedUnimplementedMetricsServiceServer()
}

// UnimplementedMetricsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (UnimplementedMetricsServiceServer) GetSystemState(context.Context, *GetSystemStateRequest) (*GetSystemStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemState not implemented")
}
func (UnimplementedMetricsServiceServer) GetLatency(context.Context, *GetLatencyRequest) (*GetLatencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatency not implemented")
}
func (UnimplementedMetricsServiceServer) mustEmbedUnimplementedMetricsServiceServer() {}

// UnsafeMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServiceServer will
// result in compilation errors.
type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	s.RegisterService(&MetricsService_ServiceDesc, srv)
}

func _MetricsService_GetSystemState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetSystemState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.metrics.v1.MetricsService/GetSystemState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetSystemState(ctx, req.(*GetSystemStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetLatency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetLatency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.metrics.v1.MetricsService/GetLatency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetLatency(ctx, req.(*GetLatencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsService_ServiceDesc is the grpc.ServiceDesc for MetricsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streammachine.api.metrics.v1.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystemState",
			Handler:    _MetricsService_GetSystemState_Handler,
		},
		{
			MethodName: "GetLatency",
			Handler:    _MetricsService_GetLatency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streammachine/api/metrics/v1/metrics.proto",
}
