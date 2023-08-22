// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/categories/v1alpha/categories_v1alpha.proto

package categories

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

// CategoriesServiceClient is the client API for CategoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoriesServiceClient interface {
	ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error)
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error)
	// Create or update a category.
	UpsertCategory(ctx context.Context, in *UpsertCategoryRequest, opts ...grpc.CallOption) (*UpsertCategoryResponse, error)
}

type categoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoriesServiceClient(cc grpc.ClientConnInterface) CategoriesServiceClient {
	return &categoriesServiceClient{cc}
}

func (c *categoriesServiceClient) ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error) {
	out := new(ListCategoriesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.categories.v1alpha.CategoriesService/ListCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error) {
	out := new(GetCategoryResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.categories.v1alpha.CategoriesService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoriesServiceClient) UpsertCategory(ctx context.Context, in *UpsertCategoryRequest, opts ...grpc.CallOption) (*UpsertCategoryResponse, error) {
	out := new(UpsertCategoryResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.categories.v1alpha.CategoriesService/UpsertCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoriesServiceServer is the server API for CategoriesService service.
// All implementations should embed UnimplementedCategoriesServiceServer
// for forward compatibility
type CategoriesServiceServer interface {
	ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error)
	GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error)
	// Create or update a category.
	UpsertCategory(context.Context, *UpsertCategoryRequest) (*UpsertCategoryResponse, error)
}

// UnimplementedCategoriesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCategoriesServiceServer struct {
}

func (UnimplementedCategoriesServiceServer) ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedCategoriesServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedCategoriesServiceServer) UpsertCategory(context.Context, *UpsertCategoryRequest) (*UpsertCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertCategory not implemented")
}

// UnsafeCategoriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoriesServiceServer will
// result in compilation errors.
type UnsafeCategoriesServiceServer interface {
	mustEmbedUnimplementedCategoriesServiceServer()
}

func RegisterCategoriesServiceServer(s grpc.ServiceRegistrar, srv CategoriesServiceServer) {
	s.RegisterService(&CategoriesService_ServiceDesc, srv)
}

func _CategoriesService_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.categories.v1alpha.CategoriesService/ListCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).ListCategories(ctx, req.(*ListCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.categories.v1alpha.CategoriesService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoriesService_UpsertCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoriesServiceServer).UpsertCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.categories.v1alpha.CategoriesService/UpsertCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoriesServiceServer).UpsertCategory(ctx, req.(*UpsertCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoriesService_ServiceDesc is the grpc.ServiceDesc for CategoriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.categories.v1alpha.CategoriesService",
	HandlerType: (*CategoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCategories",
			Handler:    _CategoriesService_ListCategories_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _CategoriesService_GetCategory_Handler,
		},
		{
			MethodName: "UpsertCategory",
			Handler:    _CategoriesService_UpsertCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/categories/v1alpha/categories_v1alpha.proto",
}