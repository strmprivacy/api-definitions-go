// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/data_policies/v1alpha/data_policies_v1alpha.proto

package data_policies

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

// DataPolicyServiceClient is the client API for DataPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataPolicyServiceClient interface {
	ListDataPolicies(ctx context.Context, in *ListDataPoliciesRequest, opts ...grpc.CallOption) (*ListDataPoliciesResponse, error)
	UpsertDataPolicy(ctx context.Context, in *UpsertDataPolicyRequest, opts ...grpc.CallOption) (*UpsertDataPolicyResponse, error)
	// returns latest policy for an id
	GetDataPolicy(ctx context.Context, in *GetDataPolicyRequest, opts ...grpc.CallOption) (*GetDataPolicyResponse, error)
	// return configured processing platforms in DPS
	ListProcessingPlatforms(ctx context.Context, in *ListProcessingPlatformsRequest, opts ...grpc.CallOption) (*ListProcessingPlatformsResponse, error)
	// return table names as known by the platform
	ListProcessingPlatformTables(ctx context.Context, in *ListProcessingPlatformTablesRequest, opts ...grpc.CallOption) (*ListProcessingPlatformTablesResponse, error)
	// return groups as known by the platform
	ListProcessingPlatformGroups(ctx context.Context, in *ListProcessingPlatformGroupsRequest, opts ...grpc.CallOption) (*ListProcessingPlatformGroupsResponse, error)
	// return a data-policy without rules sets as built from the table description on the platform
	GetProcessingPlatformBarePolicy(ctx context.Context, in *GetProcessingPlatformBarePolicyRequest, opts ...grpc.CallOption) (*GetProcessingPlatformBarePolicyResponse, error)
	ListCatalogs(ctx context.Context, in *ListCatalogsRequest, opts ...grpc.CallOption) (*ListCatalogsResponse, error)
	ListDatabases(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error)
	ListSchemas(ctx context.Context, in *ListSchemasRequest, opts ...grpc.CallOption) (*ListSchemasResponse, error)
	ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error)
}

type dataPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataPolicyServiceClient(cc grpc.ClientConnInterface) DataPolicyServiceClient {
	return &dataPolicyServiceClient{cc}
}

func (c *dataPolicyServiceClient) ListDataPolicies(ctx context.Context, in *ListDataPoliciesRequest, opts ...grpc.CallOption) (*ListDataPoliciesResponse, error) {
	out := new(ListDataPoliciesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListDataPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) UpsertDataPolicy(ctx context.Context, in *UpsertDataPolicyRequest, opts ...grpc.CallOption) (*UpsertDataPolicyResponse, error) {
	out := new(UpsertDataPolicyResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/UpsertDataPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) GetDataPolicy(ctx context.Context, in *GetDataPolicyRequest, opts ...grpc.CallOption) (*GetDataPolicyResponse, error) {
	out := new(GetDataPolicyResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/GetDataPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) ListProcessingPlatforms(ctx context.Context, in *ListProcessingPlatformsRequest, opts ...grpc.CallOption) (*ListProcessingPlatformsResponse, error) {
	out := new(ListProcessingPlatformsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListProcessingPlatforms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) ListProcessingPlatformTables(ctx context.Context, in *ListProcessingPlatformTablesRequest, opts ...grpc.CallOption) (*ListProcessingPlatformTablesResponse, error) {
	out := new(ListProcessingPlatformTablesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListProcessingPlatformTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) ListProcessingPlatformGroups(ctx context.Context, in *ListProcessingPlatformGroupsRequest, opts ...grpc.CallOption) (*ListProcessingPlatformGroupsResponse, error) {
	out := new(ListProcessingPlatformGroupsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListProcessingPlatformGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) GetProcessingPlatformBarePolicy(ctx context.Context, in *GetProcessingPlatformBarePolicyRequest, opts ...grpc.CallOption) (*GetProcessingPlatformBarePolicyResponse, error) {
	out := new(GetProcessingPlatformBarePolicyResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/GetProcessingPlatformBarePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) ListCatalogs(ctx context.Context, in *ListCatalogsRequest, opts ...grpc.CallOption) (*ListCatalogsResponse, error) {
	out := new(ListCatalogsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListCatalogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) ListDatabases(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error) {
	out := new(ListDatabasesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListDatabases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) ListSchemas(ctx context.Context, in *ListSchemasRequest, opts ...grpc.CallOption) (*ListSchemasResponse, error) {
	out := new(ListSchemasResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListSchemas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataPolicyServiceClient) ListTables(ctx context.Context, in *ListTablesRequest, opts ...grpc.CallOption) (*ListTablesResponse, error) {
	out := new(ListTablesResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataPolicyServiceServer is the server API for DataPolicyService service.
// All implementations should embed UnimplementedDataPolicyServiceServer
// for forward compatibility
type DataPolicyServiceServer interface {
	ListDataPolicies(context.Context, *ListDataPoliciesRequest) (*ListDataPoliciesResponse, error)
	UpsertDataPolicy(context.Context, *UpsertDataPolicyRequest) (*UpsertDataPolicyResponse, error)
	// returns latest policy for an id
	GetDataPolicy(context.Context, *GetDataPolicyRequest) (*GetDataPolicyResponse, error)
	// return configured processing platforms in DPS
	ListProcessingPlatforms(context.Context, *ListProcessingPlatformsRequest) (*ListProcessingPlatformsResponse, error)
	// return table names as known by the platform
	ListProcessingPlatformTables(context.Context, *ListProcessingPlatformTablesRequest) (*ListProcessingPlatformTablesResponse, error)
	// return groups as known by the platform
	ListProcessingPlatformGroups(context.Context, *ListProcessingPlatformGroupsRequest) (*ListProcessingPlatformGroupsResponse, error)
	// return a data-policy without rules sets as built from the table description on the platform
	GetProcessingPlatformBarePolicy(context.Context, *GetProcessingPlatformBarePolicyRequest) (*GetProcessingPlatformBarePolicyResponse, error)
	ListCatalogs(context.Context, *ListCatalogsRequest) (*ListCatalogsResponse, error)
	ListDatabases(context.Context, *ListDatabasesRequest) (*ListDatabasesResponse, error)
	ListSchemas(context.Context, *ListSchemasRequest) (*ListSchemasResponse, error)
	ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error)
}

// UnimplementedDataPolicyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDataPolicyServiceServer struct {
}

func (UnimplementedDataPolicyServiceServer) ListDataPolicies(context.Context, *ListDataPoliciesRequest) (*ListDataPoliciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataPolicies not implemented")
}
func (UnimplementedDataPolicyServiceServer) UpsertDataPolicy(context.Context, *UpsertDataPolicyRequest) (*UpsertDataPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDataPolicy not implemented")
}
func (UnimplementedDataPolicyServiceServer) GetDataPolicy(context.Context, *GetDataPolicyRequest) (*GetDataPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataPolicy not implemented")
}
func (UnimplementedDataPolicyServiceServer) ListProcessingPlatforms(context.Context, *ListProcessingPlatformsRequest) (*ListProcessingPlatformsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProcessingPlatforms not implemented")
}
func (UnimplementedDataPolicyServiceServer) ListProcessingPlatformTables(context.Context, *ListProcessingPlatformTablesRequest) (*ListProcessingPlatformTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProcessingPlatformTables not implemented")
}
func (UnimplementedDataPolicyServiceServer) ListProcessingPlatformGroups(context.Context, *ListProcessingPlatformGroupsRequest) (*ListProcessingPlatformGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProcessingPlatformGroups not implemented")
}
func (UnimplementedDataPolicyServiceServer) GetProcessingPlatformBarePolicy(context.Context, *GetProcessingPlatformBarePolicyRequest) (*GetProcessingPlatformBarePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProcessingPlatformBarePolicy not implemented")
}
func (UnimplementedDataPolicyServiceServer) ListCatalogs(context.Context, *ListCatalogsRequest) (*ListCatalogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCatalogs not implemented")
}
func (UnimplementedDataPolicyServiceServer) ListDatabases(context.Context, *ListDatabasesRequest) (*ListDatabasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatabases not implemented")
}
func (UnimplementedDataPolicyServiceServer) ListSchemas(context.Context, *ListSchemasRequest) (*ListSchemasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchemas not implemented")
}
func (UnimplementedDataPolicyServiceServer) ListTables(context.Context, *ListTablesRequest) (*ListTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTables not implemented")
}

// UnsafeDataPolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataPolicyServiceServer will
// result in compilation errors.
type UnsafeDataPolicyServiceServer interface {
	mustEmbedUnimplementedDataPolicyServiceServer()
}

func RegisterDataPolicyServiceServer(s grpc.ServiceRegistrar, srv DataPolicyServiceServer) {
	s.RegisterService(&DataPolicyService_ServiceDesc, srv)
}

func _DataPolicyService_ListDataPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).ListDataPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListDataPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).ListDataPolicies(ctx, req.(*ListDataPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_UpsertDataPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDataPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).UpsertDataPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/UpsertDataPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).UpsertDataPolicy(ctx, req.(*UpsertDataPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_GetDataPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).GetDataPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/GetDataPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).GetDataPolicy(ctx, req.(*GetDataPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_ListProcessingPlatforms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProcessingPlatformsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).ListProcessingPlatforms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListProcessingPlatforms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).ListProcessingPlatforms(ctx, req.(*ListProcessingPlatformsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_ListProcessingPlatformTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProcessingPlatformTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).ListProcessingPlatformTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListProcessingPlatformTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).ListProcessingPlatformTables(ctx, req.(*ListProcessingPlatformTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_ListProcessingPlatformGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProcessingPlatformGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).ListProcessingPlatformGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListProcessingPlatformGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).ListProcessingPlatformGroups(ctx, req.(*ListProcessingPlatformGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_GetProcessingPlatformBarePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProcessingPlatformBarePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).GetProcessingPlatformBarePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/GetProcessingPlatformBarePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).GetProcessingPlatformBarePolicy(ctx, req.(*GetProcessingPlatformBarePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_ListCatalogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCatalogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).ListCatalogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListCatalogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).ListCatalogs(ctx, req.(*ListCatalogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_ListDatabases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).ListDatabases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListDatabases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).ListDatabases(ctx, req.(*ListDatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_ListSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSchemasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).ListSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListSchemas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).ListSchemas(ctx, req.(*ListSchemasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataPolicyService_ListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPolicyServiceServer).ListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_policies.v1alpha.DataPolicyService/ListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPolicyServiceServer).ListTables(ctx, req.(*ListTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataPolicyService_ServiceDesc is the grpc.ServiceDesc for DataPolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataPolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.data_policies.v1alpha.DataPolicyService",
	HandlerType: (*DataPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDataPolicies",
			Handler:    _DataPolicyService_ListDataPolicies_Handler,
		},
		{
			MethodName: "UpsertDataPolicy",
			Handler:    _DataPolicyService_UpsertDataPolicy_Handler,
		},
		{
			MethodName: "GetDataPolicy",
			Handler:    _DataPolicyService_GetDataPolicy_Handler,
		},
		{
			MethodName: "ListProcessingPlatforms",
			Handler:    _DataPolicyService_ListProcessingPlatforms_Handler,
		},
		{
			MethodName: "ListProcessingPlatformTables",
			Handler:    _DataPolicyService_ListProcessingPlatformTables_Handler,
		},
		{
			MethodName: "ListProcessingPlatformGroups",
			Handler:    _DataPolicyService_ListProcessingPlatformGroups_Handler,
		},
		{
			MethodName: "GetProcessingPlatformBarePolicy",
			Handler:    _DataPolicyService_GetProcessingPlatformBarePolicy_Handler,
		},
		{
			MethodName: "ListCatalogs",
			Handler:    _DataPolicyService_ListCatalogs_Handler,
		},
		{
			MethodName: "ListDatabases",
			Handler:    _DataPolicyService_ListDatabases_Handler,
		},
		{
			MethodName: "ListSchemas",
			Handler:    _DataPolicyService_ListSchemas_Handler,
		},
		{
			MethodName: "ListTables",
			Handler:    _DataPolicyService_ListTables_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/data_policies/v1alpha/data_policies_v1alpha.proto",
}