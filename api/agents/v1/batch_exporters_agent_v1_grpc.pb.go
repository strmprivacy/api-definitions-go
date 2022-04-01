// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: strmprivacy/api/agents/v1/batch_exporters_agent_v1.proto

package agents

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

// BatchExportersAgentServiceClient is the client API for BatchExportersAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BatchExportersAgentServiceClient interface {
	GetDesiredBatchExporters(ctx context.Context, in *GetDesiredBatchExportersRequest, opts ...grpc.CallOption) (*GetDesiredBatchExportersResponse, error)
}

type batchExportersAgentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBatchExportersAgentServiceClient(cc grpc.ClientConnInterface) BatchExportersAgentServiceClient {
	return &batchExportersAgentServiceClient{cc}
}

func (c *batchExportersAgentServiceClient) GetDesiredBatchExporters(ctx context.Context, in *GetDesiredBatchExportersRequest, opts ...grpc.CallOption) (*GetDesiredBatchExportersResponse, error) {
	out := new(GetDesiredBatchExportersResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.agents.v1.BatchExportersAgentService/GetDesiredBatchExporters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatchExportersAgentServiceServer is the server API for BatchExportersAgentService service.
// All implementations must embed UnimplementedBatchExportersAgentServiceServer
// for forward compatibility
type BatchExportersAgentServiceServer interface {
	GetDesiredBatchExporters(context.Context, *GetDesiredBatchExportersRequest) (*GetDesiredBatchExportersResponse, error)
	mustEmbedUnimplementedBatchExportersAgentServiceServer()
}

// UnimplementedBatchExportersAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBatchExportersAgentServiceServer struct {
}

func (UnimplementedBatchExportersAgentServiceServer) GetDesiredBatchExporters(context.Context, *GetDesiredBatchExportersRequest) (*GetDesiredBatchExportersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDesiredBatchExporters not implemented")
}
func (UnimplementedBatchExportersAgentServiceServer) mustEmbedUnimplementedBatchExportersAgentServiceServer() {
}

// UnsafeBatchExportersAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BatchExportersAgentServiceServer will
// result in compilation errors.
type UnsafeBatchExportersAgentServiceServer interface {
	mustEmbedUnimplementedBatchExportersAgentServiceServer()
}

func RegisterBatchExportersAgentServiceServer(s grpc.ServiceRegistrar, srv BatchExportersAgentServiceServer) {
	s.RegisterService(&BatchExportersAgentService_ServiceDesc, srv)
}

func _BatchExportersAgentService_GetDesiredBatchExporters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDesiredBatchExportersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchExportersAgentServiceServer).GetDesiredBatchExporters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.agents.v1.BatchExportersAgentService/GetDesiredBatchExporters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchExportersAgentServiceServer).GetDesiredBatchExporters(ctx, req.(*GetDesiredBatchExportersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BatchExportersAgentService_ServiceDesc is the grpc.ServiceDesc for BatchExportersAgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BatchExportersAgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.agents.v1.BatchExportersAgentService",
	HandlerType: (*BatchExportersAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDesiredBatchExporters",
			Handler:    _BatchExportersAgentService_GetDesiredBatchExporters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/agents/v1/batch_exporters_agent_v1.proto",
}