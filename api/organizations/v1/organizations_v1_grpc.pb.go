// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.18.1
// source: strmprivacy/api/organizations/v1/organizations_v1.proto

package organizations

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

// OrganizationsServiceClient is the client API for OrganizationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrganizationsServiceClient interface {
	InviteUsers(ctx context.Context, in *InviteUsersRequest, opts ...grpc.CallOption) (*InviteUsersResponse, error)
	UpdateUserRoles(ctx context.Context, in *UpdateUserRolesRequest, opts ...grpc.CallOption) (*UpdateUserRolesResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	ListOrganizationMembers(ctx context.Context, in *ListOrganizationMembersRequest, opts ...grpc.CallOption) (*ListOrganizationMembersResponse, error)
}

type organizationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrganizationsServiceClient(cc grpc.ClientConnInterface) OrganizationsServiceClient {
	return &organizationsServiceClient{cc}
}

func (c *organizationsServiceClient) InviteUsers(ctx context.Context, in *InviteUsersRequest, opts ...grpc.CallOption) (*InviteUsersResponse, error) {
	out := new(InviteUsersResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.organizations.v1.OrganizationsService/InviteUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsServiceClient) UpdateUserRoles(ctx context.Context, in *UpdateUserRolesRequest, opts ...grpc.CallOption) (*UpdateUserRolesResponse, error) {
	out := new(UpdateUserRolesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.organizations.v1.OrganizationsService/UpdateUserRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.organizations.v1.OrganizationsService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationsServiceClient) ListOrganizationMembers(ctx context.Context, in *ListOrganizationMembersRequest, opts ...grpc.CallOption) (*ListOrganizationMembersResponse, error) {
	out := new(ListOrganizationMembersResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.organizations.v1.OrganizationsService/ListOrganizationMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationsServiceServer is the server API for OrganizationsService service.
// All implementations must embed UnimplementedOrganizationsServiceServer
// for forward compatibility
type OrganizationsServiceServer interface {
	InviteUsers(context.Context, *InviteUsersRequest) (*InviteUsersResponse, error)
	UpdateUserRoles(context.Context, *UpdateUserRolesRequest) (*UpdateUserRolesResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	ListOrganizationMembers(context.Context, *ListOrganizationMembersRequest) (*ListOrganizationMembersResponse, error)
	mustEmbedUnimplementedOrganizationsServiceServer()
}

// UnimplementedOrganizationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrganizationsServiceServer struct {
}

func (UnimplementedOrganizationsServiceServer) InviteUsers(context.Context, *InviteUsersRequest) (*InviteUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteUsers not implemented")
}
func (UnimplementedOrganizationsServiceServer) UpdateUserRoles(context.Context, *UpdateUserRolesRequest) (*UpdateUserRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRoles not implemented")
}
func (UnimplementedOrganizationsServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedOrganizationsServiceServer) ListOrganizationMembers(context.Context, *ListOrganizationMembersRequest) (*ListOrganizationMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrganizationMembers not implemented")
}
func (UnimplementedOrganizationsServiceServer) mustEmbedUnimplementedOrganizationsServiceServer() {}

// UnsafeOrganizationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrganizationsServiceServer will
// result in compilation errors.
type UnsafeOrganizationsServiceServer interface {
	mustEmbedUnimplementedOrganizationsServiceServer()
}

func RegisterOrganizationsServiceServer(s grpc.ServiceRegistrar, srv OrganizationsServiceServer) {
	s.RegisterService(&OrganizationsService_ServiceDesc, srv)
}

func _OrganizationsService_InviteUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServiceServer).InviteUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.organizations.v1.OrganizationsService/InviteUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServiceServer).InviteUsers(ctx, req.(*InviteUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationsService_UpdateUserRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServiceServer).UpdateUserRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.organizations.v1.OrganizationsService/UpdateUserRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServiceServer).UpdateUserRoles(ctx, req.(*UpdateUserRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationsService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.organizations.v1.OrganizationsService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrganizationsService_ListOrganizationMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrganizationMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationsServiceServer).ListOrganizationMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.organizations.v1.OrganizationsService/ListOrganizationMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationsServiceServer).ListOrganizationMembers(ctx, req.(*ListOrganizationMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrganizationsService_ServiceDesc is the grpc.ServiceDesc for OrganizationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrganizationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.organizations.v1.OrganizationsService",
	HandlerType: (*OrganizationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InviteUsers",
			Handler:    _OrganizationsService_InviteUsers_Handler,
		},
		{
			MethodName: "UpdateUserRoles",
			Handler:    _OrganizationsService_UpdateUserRoles_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _OrganizationsService_GetUser_Handler,
		},
		{
			MethodName: "ListOrganizationMembers",
			Handler:    _OrganizationsService_ListOrganizationMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/organizations/v1/organizations_v1.proto",
}
