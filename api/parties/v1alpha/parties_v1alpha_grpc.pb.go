// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/parties/v1alpha/parties_v1alpha.proto

package parties

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

// PartiesServiceClient is the client API for PartiesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartiesServiceClient interface {
	// List all parties for an organization. The organization is derived from the calling user.
	ListParties(ctx context.Context, in *ListPartiesRequest, opts ...grpc.CallOption) (*ListPartiesResponse, error)
	// Get a party by its id
	GetParty(ctx context.Context, in *GetPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error)
	// Create or update a party. To create a party, leave its id empty. To update a party, set its id.
	UpsertParty(ctx context.Context, in *UpsertPartyRequest, opts ...grpc.CallOption) (*UpsertPartyResponse, error)
}

type partiesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartiesServiceClient(cc grpc.ClientConnInterface) PartiesServiceClient {
	return &partiesServiceClient{cc}
}

func (c *partiesServiceClient) ListParties(ctx context.Context, in *ListPartiesRequest, opts ...grpc.CallOption) (*ListPartiesResponse, error) {
	out := new(ListPartiesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.parties.v1alpha.PartiesService/ListParties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partiesServiceClient) GetParty(ctx context.Context, in *GetPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error) {
	out := new(GetPartyResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.parties.v1alpha.PartiesService/GetParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partiesServiceClient) UpsertParty(ctx context.Context, in *UpsertPartyRequest, opts ...grpc.CallOption) (*UpsertPartyResponse, error) {
	out := new(UpsertPartyResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.parties.v1alpha.PartiesService/UpsertParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartiesServiceServer is the server API for PartiesService service.
// All implementations should embed UnimplementedPartiesServiceServer
// for forward compatibility
type PartiesServiceServer interface {
	// List all parties for an organization. The organization is derived from the calling user.
	ListParties(context.Context, *ListPartiesRequest) (*ListPartiesResponse, error)
	// Get a party by its id
	GetParty(context.Context, *GetPartyRequest) (*GetPartyResponse, error)
	// Create or update a party. To create a party, leave its id empty. To update a party, set its id.
	UpsertParty(context.Context, *UpsertPartyRequest) (*UpsertPartyResponse, error)
}

// UnimplementedPartiesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPartiesServiceServer struct {
}

func (UnimplementedPartiesServiceServer) ListParties(context.Context, *ListPartiesRequest) (*ListPartiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListParties not implemented")
}
func (UnimplementedPartiesServiceServer) GetParty(context.Context, *GetPartyRequest) (*GetPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParty not implemented")
}
func (UnimplementedPartiesServiceServer) UpsertParty(context.Context, *UpsertPartyRequest) (*UpsertPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertParty not implemented")
}

// UnsafePartiesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartiesServiceServer will
// result in compilation errors.
type UnsafePartiesServiceServer interface {
	mustEmbedUnimplementedPartiesServiceServer()
}

func RegisterPartiesServiceServer(s grpc.ServiceRegistrar, srv PartiesServiceServer) {
	s.RegisterService(&PartiesService_ServiceDesc, srv)
}

func _PartiesService_ListParties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPartiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartiesServiceServer).ListParties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.parties.v1alpha.PartiesService/ListParties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartiesServiceServer).ListParties(ctx, req.(*ListPartiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartiesService_GetParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartiesServiceServer).GetParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.parties.v1alpha.PartiesService/GetParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartiesServiceServer).GetParty(ctx, req.(*GetPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartiesService_UpsertParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartiesServiceServer).UpsertParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.parties.v1alpha.PartiesService/UpsertParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartiesServiceServer).UpsertParty(ctx, req.(*UpsertPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartiesService_ServiceDesc is the grpc.ServiceDesc for PartiesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartiesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.parties.v1alpha.PartiesService",
	HandlerType: (*PartiesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListParties",
			Handler:    _PartiesService_ListParties_Handler,
		},
		{
			MethodName: "GetParty",
			Handler:    _PartiesService_GetParty_Handler,
		},
		{
			MethodName: "UpsertParty",
			Handler:    _PartiesService_UpsertParty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/parties/v1alpha/parties_v1alpha.proto",
}
