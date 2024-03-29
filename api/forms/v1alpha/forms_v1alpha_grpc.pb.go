// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/forms/v1alpha/forms_v1alpha.proto

package forms

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

// FormsServiceClient is the client API for FormsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FormsServiceClient interface {
	// CRUD on Form Templates
	// returns all the form templates for the organization of the caller.
	// currently also includes the hard-coded templates
	ListFormTemplates(ctx context.Context, in *ListFormTemplatesRequest, opts ...grpc.CallOption) (*ListFormTemplatesResponse, error)
	GetFormTemplate(ctx context.Context, in *GetFormTemplateRequest, opts ...grpc.CallOption) (*GetFormTemplateResponse, error)
	UpsertFormTemplate(ctx context.Context, in *UpsertFormTemplateRequest, opts ...grpc.CallOption) (*UpsertFormTemplateResponse, error)
	DeleteFormTemplate(ctx context.Context, in *DeleteFormTemplateRequest, opts ...grpc.CallOption) (*DeleteFormTemplateResponse, error)
}

type formsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFormsServiceClient(cc grpc.ClientConnInterface) FormsServiceClient {
	return &formsServiceClient{cc}
}

func (c *formsServiceClient) ListFormTemplates(ctx context.Context, in *ListFormTemplatesRequest, opts ...grpc.CallOption) (*ListFormTemplatesResponse, error) {
	out := new(ListFormTemplatesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.forms.v1alpha.FormsService/ListFormTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formsServiceClient) GetFormTemplate(ctx context.Context, in *GetFormTemplateRequest, opts ...grpc.CallOption) (*GetFormTemplateResponse, error) {
	out := new(GetFormTemplateResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.forms.v1alpha.FormsService/GetFormTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formsServiceClient) UpsertFormTemplate(ctx context.Context, in *UpsertFormTemplateRequest, opts ...grpc.CallOption) (*UpsertFormTemplateResponse, error) {
	out := new(UpsertFormTemplateResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.forms.v1alpha.FormsService/UpsertFormTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formsServiceClient) DeleteFormTemplate(ctx context.Context, in *DeleteFormTemplateRequest, opts ...grpc.CallOption) (*DeleteFormTemplateResponse, error) {
	out := new(DeleteFormTemplateResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.forms.v1alpha.FormsService/DeleteFormTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FormsServiceServer is the server API for FormsService service.
// All implementations should embed UnimplementedFormsServiceServer
// for forward compatibility
type FormsServiceServer interface {
	// CRUD on Form Templates
	// returns all the form templates for the organization of the caller.
	// currently also includes the hard-coded templates
	ListFormTemplates(context.Context, *ListFormTemplatesRequest) (*ListFormTemplatesResponse, error)
	GetFormTemplate(context.Context, *GetFormTemplateRequest) (*GetFormTemplateResponse, error)
	UpsertFormTemplate(context.Context, *UpsertFormTemplateRequest) (*UpsertFormTemplateResponse, error)
	DeleteFormTemplate(context.Context, *DeleteFormTemplateRequest) (*DeleteFormTemplateResponse, error)
}

// UnimplementedFormsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFormsServiceServer struct {
}

func (UnimplementedFormsServiceServer) ListFormTemplates(context.Context, *ListFormTemplatesRequest) (*ListFormTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFormTemplates not implemented")
}
func (UnimplementedFormsServiceServer) GetFormTemplate(context.Context, *GetFormTemplateRequest) (*GetFormTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFormTemplate not implemented")
}
func (UnimplementedFormsServiceServer) UpsertFormTemplate(context.Context, *UpsertFormTemplateRequest) (*UpsertFormTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertFormTemplate not implemented")
}
func (UnimplementedFormsServiceServer) DeleteFormTemplate(context.Context, *DeleteFormTemplateRequest) (*DeleteFormTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFormTemplate not implemented")
}

// UnsafeFormsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FormsServiceServer will
// result in compilation errors.
type UnsafeFormsServiceServer interface {
	mustEmbedUnimplementedFormsServiceServer()
}

func RegisterFormsServiceServer(s grpc.ServiceRegistrar, srv FormsServiceServer) {
	s.RegisterService(&FormsService_ServiceDesc, srv)
}

func _FormsService_ListFormTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFormTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormsServiceServer).ListFormTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.forms.v1alpha.FormsService/ListFormTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormsServiceServer).ListFormTemplates(ctx, req.(*ListFormTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormsService_GetFormTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFormTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormsServiceServer).GetFormTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.forms.v1alpha.FormsService/GetFormTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormsServiceServer).GetFormTemplate(ctx, req.(*GetFormTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormsService_UpsertFormTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertFormTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormsServiceServer).UpsertFormTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.forms.v1alpha.FormsService/UpsertFormTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormsServiceServer).UpsertFormTemplate(ctx, req.(*UpsertFormTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FormsService_DeleteFormTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFormTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FormsServiceServer).DeleteFormTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.forms.v1alpha.FormsService/DeleteFormTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FormsServiceServer).DeleteFormTemplate(ctx, req.(*DeleteFormTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FormsService_ServiceDesc is the grpc.ServiceDesc for FormsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FormsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.forms.v1alpha.FormsService",
	HandlerType: (*FormsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFormTemplates",
			Handler:    _FormsService_ListFormTemplates_Handler,
		},
		{
			MethodName: "GetFormTemplate",
			Handler:    _FormsService_GetFormTemplate_Handler,
		},
		{
			MethodName: "UpsertFormTemplate",
			Handler:    _FormsService_UpsertFormTemplate_Handler,
		},
		{
			MethodName: "DeleteFormTemplate",
			Handler:    _FormsService_DeleteFormTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/forms/v1alpha/forms_v1alpha.proto",
}
