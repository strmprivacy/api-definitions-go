// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.15.8
// source: strmprivacy/api/account/v1/account_v1.proto

package account

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	GetAccountDetails(ctx context.Context, in *GetAccountDetailsRequest, opts ...grpc.CallOption) (*GetAccountDetailsResponse, error)
	GetLegacyBillingId(ctx context.Context, in *GetLegacyBillingIdRequest, opts ...grpc.CallOption) (*GetLegacyBillingIdResponse, error)
	CreateAccountHandle(ctx context.Context, in *CreateAccountHandleRequest, opts ...grpc.CallOption) (*CreateAccountHandleResponse, error)
	InitializeCheckout(ctx context.Context, in *InitializeCheckoutRequest, opts ...grpc.CallOption) (*InitializeCheckoutResponse, error)
	GetCheckoutStatus(ctx context.Context, in *GetCheckoutStatusRequest, opts ...grpc.CallOption) (*GetCheckoutStatusResponse, error)
	InitializeCustomerPortal(ctx context.Context, in *InitializeCustomerPortalRequest, opts ...grpc.CallOption) (*InitializeCustomerPortalResponse, error)
	// (-- api-linter: core::0134::synonyms=disabled
	//     aip.dev/not-precedent: We're not updating a Checkout here. --)
	SetCheckoutCancelled(ctx context.Context, in *SetCheckoutCancelledRequest, opts ...grpc.CallOption) (*SetCheckoutCancelledResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) GetAccountDetails(ctx context.Context, in *GetAccountDetailsRequest, opts ...grpc.CallOption) (*GetAccountDetailsResponse, error) {
	out := new(GetAccountDetailsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.account.v1.AccountService/GetAccountDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetLegacyBillingId(ctx context.Context, in *GetLegacyBillingIdRequest, opts ...grpc.CallOption) (*GetLegacyBillingIdResponse, error) {
	out := new(GetLegacyBillingIdResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.account.v1.AccountService/GetLegacyBillingId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CreateAccountHandle(ctx context.Context, in *CreateAccountHandleRequest, opts ...grpc.CallOption) (*CreateAccountHandleResponse, error) {
	out := new(CreateAccountHandleResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.account.v1.AccountService/CreateAccountHandle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) InitializeCheckout(ctx context.Context, in *InitializeCheckoutRequest, opts ...grpc.CallOption) (*InitializeCheckoutResponse, error) {
	out := new(InitializeCheckoutResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.account.v1.AccountService/InitializeCheckout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetCheckoutStatus(ctx context.Context, in *GetCheckoutStatusRequest, opts ...grpc.CallOption) (*GetCheckoutStatusResponse, error) {
	out := new(GetCheckoutStatusResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.account.v1.AccountService/GetCheckoutStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) InitializeCustomerPortal(ctx context.Context, in *InitializeCustomerPortalRequest, opts ...grpc.CallOption) (*InitializeCustomerPortalResponse, error) {
	out := new(InitializeCustomerPortalResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.account.v1.AccountService/InitializeCustomerPortal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SetCheckoutCancelled(ctx context.Context, in *SetCheckoutCancelledRequest, opts ...grpc.CallOption) (*SetCheckoutCancelledResponse, error) {
	out := new(SetCheckoutCancelledResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.account.v1.AccountService/SetCheckoutCancelled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	GetAccountDetails(context.Context, *GetAccountDetailsRequest) (*GetAccountDetailsResponse, error)
	GetLegacyBillingId(context.Context, *GetLegacyBillingIdRequest) (*GetLegacyBillingIdResponse, error)
	CreateAccountHandle(context.Context, *CreateAccountHandleRequest) (*CreateAccountHandleResponse, error)
	InitializeCheckout(context.Context, *InitializeCheckoutRequest) (*InitializeCheckoutResponse, error)
	GetCheckoutStatus(context.Context, *GetCheckoutStatusRequest) (*GetCheckoutStatusResponse, error)
	InitializeCustomerPortal(context.Context, *InitializeCustomerPortalRequest) (*InitializeCustomerPortalResponse, error)
	// (-- api-linter: core::0134::synonyms=disabled
	//     aip.dev/not-precedent: We're not updating a Checkout here. --)
	SetCheckoutCancelled(context.Context, *SetCheckoutCancelledRequest) (*SetCheckoutCancelledResponse, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) GetAccountDetails(context.Context, *GetAccountDetailsRequest) (*GetAccountDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountDetails not implemented")
}
func (UnimplementedAccountServiceServer) GetLegacyBillingId(context.Context, *GetLegacyBillingIdRequest) (*GetLegacyBillingIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLegacyBillingId not implemented")
}
func (UnimplementedAccountServiceServer) CreateAccountHandle(context.Context, *CreateAccountHandleRequest) (*CreateAccountHandleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountHandle not implemented")
}
func (UnimplementedAccountServiceServer) InitializeCheckout(context.Context, *InitializeCheckoutRequest) (*InitializeCheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeCheckout not implemented")
}
func (UnimplementedAccountServiceServer) GetCheckoutStatus(context.Context, *GetCheckoutStatusRequest) (*GetCheckoutStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckoutStatus not implemented")
}
func (UnimplementedAccountServiceServer) InitializeCustomerPortal(context.Context, *InitializeCustomerPortalRequest) (*InitializeCustomerPortalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeCustomerPortal not implemented")
}
func (UnimplementedAccountServiceServer) SetCheckoutCancelled(context.Context, *SetCheckoutCancelledRequest) (*SetCheckoutCancelledResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCheckoutCancelled not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_GetAccountDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.account.v1.AccountService/GetAccountDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountDetails(ctx, req.(*GetAccountDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetLegacyBillingId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLegacyBillingIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetLegacyBillingId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.account.v1.AccountService/GetLegacyBillingId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetLegacyBillingId(ctx, req.(*GetLegacyBillingIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CreateAccountHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountHandleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccountHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.account.v1.AccountService/CreateAccountHandle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccountHandle(ctx, req.(*CreateAccountHandleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_InitializeCheckout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeCheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).InitializeCheckout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.account.v1.AccountService/InitializeCheckout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).InitializeCheckout(ctx, req.(*InitializeCheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetCheckoutStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckoutStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetCheckoutStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.account.v1.AccountService/GetCheckoutStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetCheckoutStatus(ctx, req.(*GetCheckoutStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_InitializeCustomerPortal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeCustomerPortalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).InitializeCustomerPortal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.account.v1.AccountService/InitializeCustomerPortal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).InitializeCustomerPortal(ctx, req.(*InitializeCustomerPortalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SetCheckoutCancelled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCheckoutCancelledRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SetCheckoutCancelled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.account.v1.AccountService/SetCheckoutCancelled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SetCheckoutCancelled(ctx, req.(*SetCheckoutCancelledRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.account.v1.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountDetails",
			Handler:    _AccountService_GetAccountDetails_Handler,
		},
		{
			MethodName: "GetLegacyBillingId",
			Handler:    _AccountService_GetLegacyBillingId_Handler,
		},
		{
			MethodName: "CreateAccountHandle",
			Handler:    _AccountService_CreateAccountHandle_Handler,
		},
		{
			MethodName: "InitializeCheckout",
			Handler:    _AccountService_InitializeCheckout_Handler,
		},
		{
			MethodName: "GetCheckoutStatus",
			Handler:    _AccountService_GetCheckoutStatus_Handler,
		},
		{
			MethodName: "InitializeCustomerPortal",
			Handler:    _AccountService_InitializeCustomerPortal_Handler,
		},
		{
			MethodName: "SetCheckoutCancelled",
			Handler:    _AccountService_SetCheckoutCancelled_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/account/v1/account_v1.proto",
}
