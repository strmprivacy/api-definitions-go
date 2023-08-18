// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/ropa/v1alpha/ropa_v1alpha.proto

package ropa

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

// RopaServiceClient is the client API for RopaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RopaServiceClient interface {
	// This will return all (latest versions of) the records.
	GetRopa(ctx context.Context, in *GetRopaRequest, opts ...grpc.CallOption) (*GetRopaResponse, error)
	// Create a new record or a new version of an existing record.
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error)
	// Get a record by id or data contract ref.
	GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*GetRecordResponse, error)
}

type ropaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRopaServiceClient(cc grpc.ClientConnInterface) RopaServiceClient {
	return &ropaServiceClient{cc}
}

func (c *ropaServiceClient) GetRopa(ctx context.Context, in *GetRopaRequest, opts ...grpc.CallOption) (*GetRopaResponse, error) {
	out := new(GetRopaResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.ropa.v1alpha.RopaService/GetRopa", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ropaServiceClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error) {
	out := new(CreateRecordResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.ropa.v1alpha.RopaService/CreateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ropaServiceClient) GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*GetRecordResponse, error) {
	out := new(GetRecordResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.ropa.v1alpha.RopaService/GetRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RopaServiceServer is the server API for RopaService service.
// All implementations should embed UnimplementedRopaServiceServer
// for forward compatibility
type RopaServiceServer interface {
	// This will return all (latest versions of) the records.
	GetRopa(context.Context, *GetRopaRequest) (*GetRopaResponse, error)
	// Create a new record or a new version of an existing record.
	CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error)
	// Get a record by id or data contract ref.
	GetRecord(context.Context, *GetRecordRequest) (*GetRecordResponse, error)
}

// UnimplementedRopaServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRopaServiceServer struct {
}

func (UnimplementedRopaServiceServer) GetRopa(context.Context, *GetRopaRequest) (*GetRopaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRopa not implemented")
}
func (UnimplementedRopaServiceServer) CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedRopaServiceServer) GetRecord(context.Context, *GetRecordRequest) (*GetRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}

// UnsafeRopaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RopaServiceServer will
// result in compilation errors.
type UnsafeRopaServiceServer interface {
	mustEmbedUnimplementedRopaServiceServer()
}

func RegisterRopaServiceServer(s grpc.ServiceRegistrar, srv RopaServiceServer) {
	s.RegisterService(&RopaService_ServiceDesc, srv)
}

func _RopaService_GetRopa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRopaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RopaServiceServer).GetRopa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.ropa.v1alpha.RopaService/GetRopa",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RopaServiceServer).GetRopa(ctx, req.(*GetRopaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RopaService_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RopaServiceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.ropa.v1alpha.RopaService/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RopaServiceServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RopaService_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RopaServiceServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.ropa.v1alpha.RopaService/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RopaServiceServer).GetRecord(ctx, req.(*GetRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RopaService_ServiceDesc is the grpc.ServiceDesc for RopaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RopaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.ropa.v1alpha.RopaService",
	HandlerType: (*RopaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRopa",
			Handler:    _RopaService_GetRopa_Handler,
		},
		{
			MethodName: "CreateRecord",
			Handler:    _RopaService_CreateRecord_Handler,
		},
		{
			MethodName: "GetRecord",
			Handler:    _RopaService_GetRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/ropa/v1alpha/ropa_v1alpha.proto",
}
