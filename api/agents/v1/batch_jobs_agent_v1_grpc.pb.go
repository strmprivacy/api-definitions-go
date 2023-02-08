// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/agents/v1/batch_jobs_agent_v1.proto

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

// BatchJobsAgentServiceClient is the client API for BatchJobsAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BatchJobsAgentServiceClient interface {
	GetDesiredBatchJobs(ctx context.Context, in *GetDesiredBatchJobsRequest, opts ...grpc.CallOption) (*GetDesiredBatchJobsResponse, error)
}

type batchJobsAgentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBatchJobsAgentServiceClient(cc grpc.ClientConnInterface) BatchJobsAgentServiceClient {
	return &batchJobsAgentServiceClient{cc}
}

func (c *batchJobsAgentServiceClient) GetDesiredBatchJobs(ctx context.Context, in *GetDesiredBatchJobsRequest, opts ...grpc.CallOption) (*GetDesiredBatchJobsResponse, error) {
	out := new(GetDesiredBatchJobsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.agents.v1.BatchJobsAgentService/GetDesiredBatchJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatchJobsAgentServiceServer is the server API for BatchJobsAgentService service.
// All implementations should embed UnimplementedBatchJobsAgentServiceServer
// for forward compatibility
type BatchJobsAgentServiceServer interface {
	GetDesiredBatchJobs(context.Context, *GetDesiredBatchJobsRequest) (*GetDesiredBatchJobsResponse, error)
}

// UnimplementedBatchJobsAgentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBatchJobsAgentServiceServer struct {
}

func (UnimplementedBatchJobsAgentServiceServer) GetDesiredBatchJobs(context.Context, *GetDesiredBatchJobsRequest) (*GetDesiredBatchJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDesiredBatchJobs not implemented")
}

// UnsafeBatchJobsAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BatchJobsAgentServiceServer will
// result in compilation errors.
type UnsafeBatchJobsAgentServiceServer interface {
	mustEmbedUnimplementedBatchJobsAgentServiceServer()
}

func RegisterBatchJobsAgentServiceServer(s grpc.ServiceRegistrar, srv BatchJobsAgentServiceServer) {
	s.RegisterService(&BatchJobsAgentService_ServiceDesc, srv)
}

func _BatchJobsAgentService_GetDesiredBatchJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDesiredBatchJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobsAgentServiceServer).GetDesiredBatchJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.agents.v1.BatchJobsAgentService/GetDesiredBatchJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobsAgentServiceServer).GetDesiredBatchJobs(ctx, req.(*GetDesiredBatchJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BatchJobsAgentService_ServiceDesc is the grpc.ServiceDesc for BatchJobsAgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BatchJobsAgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.agents.v1.BatchJobsAgentService",
	HandlerType: (*BatchJobsAgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDesiredBatchJobs",
			Handler:    _BatchJobsAgentService_GetDesiredBatchJobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/agents/v1/batch_jobs_agent_v1.proto",
}
