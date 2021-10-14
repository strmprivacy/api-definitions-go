// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.15.8
// source: streammachine/api/streams/v1/streams_v1.proto

package streams

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

// StreamsServiceClient is the client API for StreamsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamsServiceClient interface {
	ListStreams(ctx context.Context, in *ListStreamsRequest, opts ...grpc.CallOption) (*ListStreamsResponse, error)
	GetStream(ctx context.Context, in *GetStreamRequest, opts ...grpc.CallOption) (*GetStreamResponse, error)
	DeleteStream(ctx context.Context, in *DeleteStreamRequest, opts ...grpc.CallOption) (*DeleteStreamResponse, error)
	CreateStream(ctx context.Context, in *CreateStreamRequest, opts ...grpc.CallOption) (*CreateStreamResponse, error)
}

type streamsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamsServiceClient(cc grpc.ClientConnInterface) StreamsServiceClient {
	return &streamsServiceClient{cc}
}

func (c *streamsServiceClient) ListStreams(ctx context.Context, in *ListStreamsRequest, opts ...grpc.CallOption) (*ListStreamsResponse, error) {
	out := new(ListStreamsResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.streams.v1.StreamsService/ListStreams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamsServiceClient) GetStream(ctx context.Context, in *GetStreamRequest, opts ...grpc.CallOption) (*GetStreamResponse, error) {
	out := new(GetStreamResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.streams.v1.StreamsService/GetStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamsServiceClient) DeleteStream(ctx context.Context, in *DeleteStreamRequest, opts ...grpc.CallOption) (*DeleteStreamResponse, error) {
	out := new(DeleteStreamResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.streams.v1.StreamsService/DeleteStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamsServiceClient) CreateStream(ctx context.Context, in *CreateStreamRequest, opts ...grpc.CallOption) (*CreateStreamResponse, error) {
	out := new(CreateStreamResponse)
	err := c.cc.Invoke(ctx, "/streammachine.api.streams.v1.StreamsService/CreateStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamsServiceServer is the server API for StreamsService service.
// All implementations must embed UnimplementedStreamsServiceServer
// for forward compatibility
type StreamsServiceServer interface {
	ListStreams(context.Context, *ListStreamsRequest) (*ListStreamsResponse, error)
	GetStream(context.Context, *GetStreamRequest) (*GetStreamResponse, error)
	DeleteStream(context.Context, *DeleteStreamRequest) (*DeleteStreamResponse, error)
	CreateStream(context.Context, *CreateStreamRequest) (*CreateStreamResponse, error)
	mustEmbedUnimplementedStreamsServiceServer()
}

// UnimplementedStreamsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamsServiceServer struct {
}

func (UnimplementedStreamsServiceServer) ListStreams(context.Context, *ListStreamsRequest) (*ListStreamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStreams not implemented")
}
func (UnimplementedStreamsServiceServer) GetStream(context.Context, *GetStreamRequest) (*GetStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (UnimplementedStreamsServiceServer) DeleteStream(context.Context, *DeleteStreamRequest) (*DeleteStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStream not implemented")
}
func (UnimplementedStreamsServiceServer) CreateStream(context.Context, *CreateStreamRequest) (*CreateStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStream not implemented")
}
func (UnimplementedStreamsServiceServer) mustEmbedUnimplementedStreamsServiceServer() {}

// UnsafeStreamsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamsServiceServer will
// result in compilation errors.
type UnsafeStreamsServiceServer interface {
	mustEmbedUnimplementedStreamsServiceServer()
}

func RegisterStreamsServiceServer(s grpc.ServiceRegistrar, srv StreamsServiceServer) {
	s.RegisterService(&StreamsService_ServiceDesc, srv)
}

func _StreamsService_ListStreams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStreamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamsServiceServer).ListStreams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.streams.v1.StreamsService/ListStreams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamsServiceServer).ListStreams(ctx, req.(*ListStreamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamsService_GetStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamsServiceServer).GetStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.streams.v1.StreamsService/GetStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamsServiceServer).GetStream(ctx, req.(*GetStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamsService_DeleteStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamsServiceServer).DeleteStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.streams.v1.StreamsService/DeleteStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamsServiceServer).DeleteStream(ctx, req.(*DeleteStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamsService_CreateStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamsServiceServer).CreateStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/streammachine.api.streams.v1.StreamsService/CreateStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamsServiceServer).CreateStream(ctx, req.(*CreateStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamsService_ServiceDesc is the grpc.ServiceDesc for StreamsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streammachine.api.streams.v1.StreamsService",
	HandlerType: (*StreamsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListStreams",
			Handler:    _StreamsService_ListStreams_Handler,
		},
		{
			MethodName: "GetStream",
			Handler:    _StreamsService_GetStream_Handler,
		},
		{
			MethodName: "DeleteStream",
			Handler:    _StreamsService_DeleteStream_Handler,
		},
		{
			MethodName: "CreateStream",
			Handler:    _StreamsService_CreateStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streammachine/api/streams/v1/streams_v1.proto",
}
