// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/data_contracts/v1/data_contracts_v1.proto

package data_contracts

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

// DataContractsServiceClient is the client API for DataContractsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataContractsServiceClient interface {
	ListDataContracts(ctx context.Context, in *ListDataContractsRequest, opts ...grpc.CallOption) (*ListDataContractsResponse, error)
	GetDataContract(ctx context.Context, in *GetDataContractRequest, opts ...grpc.CallOption) (*GetDataContractResponse, error)
	CreateDataContract(ctx context.Context, in *CreateDataContractRequest, opts ...grpc.CallOption) (*CreateDataContractResponse, error)
	UpdateDataContract(ctx context.Context, in *UpdateDataContractRequest, opts ...grpc.CallOption) (*UpdateDataContractResponse, error)
	ReviewDataContract(ctx context.Context, in *ReviewDataContractRequest, opts ...grpc.CallOption) (*ReviewDataContractResponse, error)
	ActivateDataContract(ctx context.Context, in *ActivateDataContractRequest, opts ...grpc.CallOption) (*ActivateDataContractResponse, error)
	DeleteDataContract(ctx context.Context, in *DeleteDataContractRequest, opts ...grpc.CallOption) (*DeleteDataContractResponse, error)
	ArchiveDataContract(ctx context.Context, in *ArchiveDataContractRequest, opts ...grpc.CallOption) (*ArchiveDataContractResponse, error)
	ApproveDataContract(ctx context.Context, in *ApproveDataContractRequest, opts ...grpc.CallOption) (*ApproveDataContractResponse, error)
	GetDataContractSchemaCode(ctx context.Context, in *GetDataContractSchemaCodeRequest, opts ...grpc.CallOption) (*GetDataContractSchemaCodeResponse, error)
	GetDataContractSchemaDefinition(ctx context.Context, in *GetDataContractSchemaDefinitionRequest, opts ...grpc.CallOption) (*GetDataContractSchemaDefinitionResponse, error)
	ValidateDataContractsMaskedFields(ctx context.Context, in *ValidateDataContractsMaskedFieldsRequest, opts ...grpc.CallOption) (*ValidateDataContractsMaskedFieldsResponse, error)
	SimpleSchemaDryRun(ctx context.Context, in *SimpleSchemaDryRunRequest, opts ...grpc.CallOption) (*SimpleSchemaDryRunResponse, error)
}

type dataContractsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataContractsServiceClient(cc grpc.ClientConnInterface) DataContractsServiceClient {
	return &dataContractsServiceClient{cc}
}

func (c *dataContractsServiceClient) ListDataContracts(ctx context.Context, in *ListDataContractsRequest, opts ...grpc.CallOption) (*ListDataContractsResponse, error) {
	out := new(ListDataContractsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/ListDataContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) GetDataContract(ctx context.Context, in *GetDataContractRequest, opts ...grpc.CallOption) (*GetDataContractResponse, error) {
	out := new(GetDataContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/GetDataContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) CreateDataContract(ctx context.Context, in *CreateDataContractRequest, opts ...grpc.CallOption) (*CreateDataContractResponse, error) {
	out := new(CreateDataContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/CreateDataContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) UpdateDataContract(ctx context.Context, in *UpdateDataContractRequest, opts ...grpc.CallOption) (*UpdateDataContractResponse, error) {
	out := new(UpdateDataContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/UpdateDataContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) ReviewDataContract(ctx context.Context, in *ReviewDataContractRequest, opts ...grpc.CallOption) (*ReviewDataContractResponse, error) {
	out := new(ReviewDataContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/ReviewDataContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) ActivateDataContract(ctx context.Context, in *ActivateDataContractRequest, opts ...grpc.CallOption) (*ActivateDataContractResponse, error) {
	out := new(ActivateDataContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/ActivateDataContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) DeleteDataContract(ctx context.Context, in *DeleteDataContractRequest, opts ...grpc.CallOption) (*DeleteDataContractResponse, error) {
	out := new(DeleteDataContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/DeleteDataContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) ArchiveDataContract(ctx context.Context, in *ArchiveDataContractRequest, opts ...grpc.CallOption) (*ArchiveDataContractResponse, error) {
	out := new(ArchiveDataContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/ArchiveDataContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) ApproveDataContract(ctx context.Context, in *ApproveDataContractRequest, opts ...grpc.CallOption) (*ApproveDataContractResponse, error) {
	out := new(ApproveDataContractResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/ApproveDataContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) GetDataContractSchemaCode(ctx context.Context, in *GetDataContractSchemaCodeRequest, opts ...grpc.CallOption) (*GetDataContractSchemaCodeResponse, error) {
	out := new(GetDataContractSchemaCodeResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/GetDataContractSchemaCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) GetDataContractSchemaDefinition(ctx context.Context, in *GetDataContractSchemaDefinitionRequest, opts ...grpc.CallOption) (*GetDataContractSchemaDefinitionResponse, error) {
	out := new(GetDataContractSchemaDefinitionResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/GetDataContractSchemaDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) ValidateDataContractsMaskedFields(ctx context.Context, in *ValidateDataContractsMaskedFieldsRequest, opts ...grpc.CallOption) (*ValidateDataContractsMaskedFieldsResponse, error) {
	out := new(ValidateDataContractsMaskedFieldsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/ValidateDataContractsMaskedFields", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataContractsServiceClient) SimpleSchemaDryRun(ctx context.Context, in *SimpleSchemaDryRunRequest, opts ...grpc.CallOption) (*SimpleSchemaDryRunResponse, error) {
	out := new(SimpleSchemaDryRunResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.data_contracts.v1.DataContractsService/SimpleSchemaDryRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataContractsServiceServer is the server API for DataContractsService service.
// All implementations should embed UnimplementedDataContractsServiceServer
// for forward compatibility
type DataContractsServiceServer interface {
	ListDataContracts(context.Context, *ListDataContractsRequest) (*ListDataContractsResponse, error)
	GetDataContract(context.Context, *GetDataContractRequest) (*GetDataContractResponse, error)
	CreateDataContract(context.Context, *CreateDataContractRequest) (*CreateDataContractResponse, error)
	UpdateDataContract(context.Context, *UpdateDataContractRequest) (*UpdateDataContractResponse, error)
	ReviewDataContract(context.Context, *ReviewDataContractRequest) (*ReviewDataContractResponse, error)
	ActivateDataContract(context.Context, *ActivateDataContractRequest) (*ActivateDataContractResponse, error)
	DeleteDataContract(context.Context, *DeleteDataContractRequest) (*DeleteDataContractResponse, error)
	ArchiveDataContract(context.Context, *ArchiveDataContractRequest) (*ArchiveDataContractResponse, error)
	ApproveDataContract(context.Context, *ApproveDataContractRequest) (*ApproveDataContractResponse, error)
	GetDataContractSchemaCode(context.Context, *GetDataContractSchemaCodeRequest) (*GetDataContractSchemaCodeResponse, error)
	GetDataContractSchemaDefinition(context.Context, *GetDataContractSchemaDefinitionRequest) (*GetDataContractSchemaDefinitionResponse, error)
	ValidateDataContractsMaskedFields(context.Context, *ValidateDataContractsMaskedFieldsRequest) (*ValidateDataContractsMaskedFieldsResponse, error)
	SimpleSchemaDryRun(context.Context, *SimpleSchemaDryRunRequest) (*SimpleSchemaDryRunResponse, error)
}

// UnimplementedDataContractsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDataContractsServiceServer struct {
}

func (UnimplementedDataContractsServiceServer) ListDataContracts(context.Context, *ListDataContractsRequest) (*ListDataContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDataContracts not implemented")
}
func (UnimplementedDataContractsServiceServer) GetDataContract(context.Context, *GetDataContractRequest) (*GetDataContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataContract not implemented")
}
func (UnimplementedDataContractsServiceServer) CreateDataContract(context.Context, *CreateDataContractRequest) (*CreateDataContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDataContract not implemented")
}
func (UnimplementedDataContractsServiceServer) UpdateDataContract(context.Context, *UpdateDataContractRequest) (*UpdateDataContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDataContract not implemented")
}
func (UnimplementedDataContractsServiceServer) ReviewDataContract(context.Context, *ReviewDataContractRequest) (*ReviewDataContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewDataContract not implemented")
}
func (UnimplementedDataContractsServiceServer) ActivateDataContract(context.Context, *ActivateDataContractRequest) (*ActivateDataContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateDataContract not implemented")
}
func (UnimplementedDataContractsServiceServer) DeleteDataContract(context.Context, *DeleteDataContractRequest) (*DeleteDataContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDataContract not implemented")
}
func (UnimplementedDataContractsServiceServer) ArchiveDataContract(context.Context, *ArchiveDataContractRequest) (*ArchiveDataContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArchiveDataContract not implemented")
}
func (UnimplementedDataContractsServiceServer) ApproveDataContract(context.Context, *ApproveDataContractRequest) (*ApproveDataContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveDataContract not implemented")
}
func (UnimplementedDataContractsServiceServer) GetDataContractSchemaCode(context.Context, *GetDataContractSchemaCodeRequest) (*GetDataContractSchemaCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataContractSchemaCode not implemented")
}
func (UnimplementedDataContractsServiceServer) GetDataContractSchemaDefinition(context.Context, *GetDataContractSchemaDefinitionRequest) (*GetDataContractSchemaDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataContractSchemaDefinition not implemented")
}
func (UnimplementedDataContractsServiceServer) ValidateDataContractsMaskedFields(context.Context, *ValidateDataContractsMaskedFieldsRequest) (*ValidateDataContractsMaskedFieldsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateDataContractsMaskedFields not implemented")
}
func (UnimplementedDataContractsServiceServer) SimpleSchemaDryRun(context.Context, *SimpleSchemaDryRunRequest) (*SimpleSchemaDryRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimpleSchemaDryRun not implemented")
}

// UnsafeDataContractsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataContractsServiceServer will
// result in compilation errors.
type UnsafeDataContractsServiceServer interface {
	mustEmbedUnimplementedDataContractsServiceServer()
}

func RegisterDataContractsServiceServer(s grpc.ServiceRegistrar, srv DataContractsServiceServer) {
	s.RegisterService(&DataContractsService_ServiceDesc, srv)
}

func _DataContractsService_ListDataContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataContractsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).ListDataContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/ListDataContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).ListDataContracts(ctx, req.(*ListDataContractsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_GetDataContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).GetDataContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/GetDataContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).GetDataContract(ctx, req.(*GetDataContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_CreateDataContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDataContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).CreateDataContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/CreateDataContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).CreateDataContract(ctx, req.(*CreateDataContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_UpdateDataContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDataContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).UpdateDataContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/UpdateDataContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).UpdateDataContract(ctx, req.(*UpdateDataContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_ReviewDataContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewDataContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).ReviewDataContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/ReviewDataContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).ReviewDataContract(ctx, req.(*ReviewDataContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_ActivateDataContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateDataContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).ActivateDataContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/ActivateDataContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).ActivateDataContract(ctx, req.(*ActivateDataContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_DeleteDataContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDataContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).DeleteDataContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/DeleteDataContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).DeleteDataContract(ctx, req.(*DeleteDataContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_ArchiveDataContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArchiveDataContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).ArchiveDataContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/ArchiveDataContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).ArchiveDataContract(ctx, req.(*ArchiveDataContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_ApproveDataContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveDataContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).ApproveDataContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/ApproveDataContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).ApproveDataContract(ctx, req.(*ApproveDataContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_GetDataContractSchemaCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataContractSchemaCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).GetDataContractSchemaCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/GetDataContractSchemaCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).GetDataContractSchemaCode(ctx, req.(*GetDataContractSchemaCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_GetDataContractSchemaDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataContractSchemaDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).GetDataContractSchemaDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/GetDataContractSchemaDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).GetDataContractSchemaDefinition(ctx, req.(*GetDataContractSchemaDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_ValidateDataContractsMaskedFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateDataContractsMaskedFieldsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).ValidateDataContractsMaskedFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/ValidateDataContractsMaskedFields",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).ValidateDataContractsMaskedFields(ctx, req.(*ValidateDataContractsMaskedFieldsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataContractsService_SimpleSchemaDryRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleSchemaDryRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataContractsServiceServer).SimpleSchemaDryRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.data_contracts.v1.DataContractsService/SimpleSchemaDryRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataContractsServiceServer).SimpleSchemaDryRun(ctx, req.(*SimpleSchemaDryRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataContractsService_ServiceDesc is the grpc.ServiceDesc for DataContractsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataContractsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.data_contracts.v1.DataContractsService",
	HandlerType: (*DataContractsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDataContracts",
			Handler:    _DataContractsService_ListDataContracts_Handler,
		},
		{
			MethodName: "GetDataContract",
			Handler:    _DataContractsService_GetDataContract_Handler,
		},
		{
			MethodName: "CreateDataContract",
			Handler:    _DataContractsService_CreateDataContract_Handler,
		},
		{
			MethodName: "UpdateDataContract",
			Handler:    _DataContractsService_UpdateDataContract_Handler,
		},
		{
			MethodName: "ReviewDataContract",
			Handler:    _DataContractsService_ReviewDataContract_Handler,
		},
		{
			MethodName: "ActivateDataContract",
			Handler:    _DataContractsService_ActivateDataContract_Handler,
		},
		{
			MethodName: "DeleteDataContract",
			Handler:    _DataContractsService_DeleteDataContract_Handler,
		},
		{
			MethodName: "ArchiveDataContract",
			Handler:    _DataContractsService_ArchiveDataContract_Handler,
		},
		{
			MethodName: "ApproveDataContract",
			Handler:    _DataContractsService_ApproveDataContract_Handler,
		},
		{
			MethodName: "GetDataContractSchemaCode",
			Handler:    _DataContractsService_GetDataContractSchemaCode_Handler,
		},
		{
			MethodName: "GetDataContractSchemaDefinition",
			Handler:    _DataContractsService_GetDataContractSchemaDefinition_Handler,
		},
		{
			MethodName: "ValidateDataContractsMaskedFields",
			Handler:    _DataContractsService_ValidateDataContractsMaskedFields_Handler,
		},
		{
			MethodName: "SimpleSchemaDryRun",
			Handler:    _DataContractsService_SimpleSchemaDryRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "strmprivacy/api/data_contracts/v1/data_contracts_v1.proto",
}
