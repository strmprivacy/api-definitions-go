// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.2
// source: streammachine/api/schemas/v1/schemas_v1.proto

package schemas

import (
	v1 "github.com/streammachineio/api-definitions-go/api/entities/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListSchemasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	Filter    string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListSchemasRequest) Reset() {
	*x = ListSchemasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchemasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchemasRequest) ProtoMessage() {}

func (x *ListSchemasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchemasRequest.ProtoReflect.Descriptor instead.
func (*ListSchemasRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescGZIP(), []int{0}
}

func (x *ListSchemasRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *ListSchemasRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// (-- api-linter: core::0158::response-plural-first-field=disabled
//     aip.dev/not-precedent: Somehow this rule is not correctly checked here. --)
type ListSchemasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schemas []*v1.Schema `protobuf:"bytes,1,rep,name=schemas,proto3" json:"schemas,omitempty"`
}

func (x *ListSchemasResponse) Reset() {
	*x = ListSchemasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchemasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchemasResponse) ProtoMessage() {}

func (x *ListSchemasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchemasResponse.ProtoReflect.Descriptor instead.
func (*ListSchemasResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListSchemasResponse) GetSchemas() []*v1.Schema {
	if x != nil {
		return x.Schemas
	}
	return nil
}

type CreateSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string     `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	Schema    *v1.Schema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *CreateSchemaRequest) Reset() {
	*x = CreateSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchemaRequest) ProtoMessage() {}

func (x *CreateSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSchemaRequest.ProtoReflect.Descriptor instead.
func (*CreateSchemaRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSchemaRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *CreateSchemaRequest) GetSchema() *v1.Schema {
	if x != nil {
		return x.Schema
	}
	return nil
}

type CreateSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema *v1.Schema `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *CreateSchemaResponse) Reset() {
	*x = CreateSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchemaResponse) ProtoMessage() {}

func (x *CreateSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSchemaResponse.ProtoReflect.Descriptor instead.
func (*CreateSchemaResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSchemaResponse) GetSchema() *v1.Schema {
	if x != nil {
		return x.Schema
	}
	return nil
}

type GetSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId  string              `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	Ref        *v1.SchemaRef       `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	ClusterRef *v1.KafkaClusterRef `protobuf:"bytes,3,opt,name=cluster_ref,json=clusterRef,proto3" json:"cluster_ref,omitempty"`
}

func (x *GetSchemaRequest) Reset() {
	*x = GetSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaRequest) ProtoMessage() {}

func (x *GetSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaRequest.ProtoReflect.Descriptor instead.
func (*GetSchemaRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescGZIP(), []int{4}
}

func (x *GetSchemaRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *GetSchemaRequest) GetRef() *v1.SchemaRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *GetSchemaRequest) GetClusterRef() *v1.KafkaClusterRef {
	if x != nil {
		return x.ClusterRef
	}
	return nil
}

type GetSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema      *v1.Schema `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	ConfluentId int32      `protobuf:"varint,2,opt,name=confluent_id,json=confluentId,proto3" json:"confluent_id,omitempty"`
}

func (x *GetSchemaResponse) Reset() {
	*x = GetSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaResponse) ProtoMessage() {}

func (x *GetSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaResponse.ProtoReflect.Descriptor instead.
func (*GetSchemaResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescGZIP(), []int{5}
}

func (x *GetSchemaResponse) GetSchema() *v1.Schema {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *GetSchemaResponse) GetConfluentId() int32 {
	if x != nil {
		return x.ConfluentId
	}
	return 0
}

type GetSchemaCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string        `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	Ref       *v1.SchemaRef `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	Type      v1.SchemaType `protobuf:"varint,3,opt,name=type,proto3,enum=streammachine.api.entities.v1.SchemaType" json:"type,omitempty"`
}

func (x *GetSchemaCodeRequest) Reset() {
	*x = GetSchemaCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchemaCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaCodeRequest) ProtoMessage() {}

func (x *GetSchemaCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaCodeRequest.ProtoReflect.Descriptor instead.
func (*GetSchemaCodeRequest) Descriptor() ([]byte, []int) {
	return file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetSchemaCodeRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *GetSchemaCodeRequest) GetRef() *v1.SchemaRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *GetSchemaCodeRequest) GetType() v1.SchemaType {
	if x != nil {
		return x.Type
	}
	return v1.SchemaType_SCHEMA_TYPE_UNSPECIFIED
}

type GetSchemaCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	DataSize int64  `protobuf:"varint,2,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
	Data     []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSchemaCodeResponse) Reset() {
	*x = GetSchemaCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchemaCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaCodeResponse) ProtoMessage() {}

func (x *GetSchemaCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaCodeResponse.ProtoReflect.Descriptor instead.
func (*GetSchemaCodeResponse) Descriptor() ([]byte, []int) {
	return file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescGZIP(), []int{7}
}

func (x *GetSchemaCodeResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *GetSchemaCodeResponse) GetDataSize() int64 {
	if x != nil {
		return x.DataSize
	}
	return 0
}

func (x *GetSchemaCodeResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_streammachine_api_schemas_v1_schemas_v1_proto protoreflect.FileDescriptor

var file_streammachine_api_schemas_v1_schemas_v1_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x07, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x22, 0x73, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x55, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22,
	0xbe, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x66, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12,
	0x4f, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x66, 0x52, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66,
	0x22, 0x75, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x75, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66,
	0x6c, 0x75, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x3a, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x66, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x3d, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x64, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xe3, 0x03, 0x0a, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x12, 0x30, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x78, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x2e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x69, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescOnce sync.Once
	file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescData = file_streammachine_api_schemas_v1_schemas_v1_proto_rawDesc
)

func file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescGZIP() []byte {
	file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescOnce.Do(func() {
		file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescData)
	})
	return file_streammachine_api_schemas_v1_schemas_v1_proto_rawDescData
}

var file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_streammachine_api_schemas_v1_schemas_v1_proto_goTypes = []interface{}{
	(*ListSchemasRequest)(nil),    // 0: streammachine.api.schemas.v1.ListSchemasRequest
	(*ListSchemasResponse)(nil),   // 1: streammachine.api.schemas.v1.ListSchemasResponse
	(*CreateSchemaRequest)(nil),   // 2: streammachine.api.schemas.v1.CreateSchemaRequest
	(*CreateSchemaResponse)(nil),  // 3: streammachine.api.schemas.v1.CreateSchemaResponse
	(*GetSchemaRequest)(nil),      // 4: streammachine.api.schemas.v1.GetSchemaRequest
	(*GetSchemaResponse)(nil),     // 5: streammachine.api.schemas.v1.GetSchemaResponse
	(*GetSchemaCodeRequest)(nil),  // 6: streammachine.api.schemas.v1.GetSchemaCodeRequest
	(*GetSchemaCodeResponse)(nil), // 7: streammachine.api.schemas.v1.GetSchemaCodeResponse
	(*v1.Schema)(nil),             // 8: streammachine.api.entities.v1.Schema
	(*v1.SchemaRef)(nil),          // 9: streammachine.api.entities.v1.SchemaRef
	(*v1.KafkaClusterRef)(nil),    // 10: streammachine.api.entities.v1.KafkaClusterRef
	(v1.SchemaType)(0),            // 11: streammachine.api.entities.v1.SchemaType
}
var file_streammachine_api_schemas_v1_schemas_v1_proto_depIdxs = []int32{
	8,  // 0: streammachine.api.schemas.v1.ListSchemasResponse.schemas:type_name -> streammachine.api.entities.v1.Schema
	8,  // 1: streammachine.api.schemas.v1.CreateSchemaRequest.schema:type_name -> streammachine.api.entities.v1.Schema
	8,  // 2: streammachine.api.schemas.v1.CreateSchemaResponse.schema:type_name -> streammachine.api.entities.v1.Schema
	9,  // 3: streammachine.api.schemas.v1.GetSchemaRequest.ref:type_name -> streammachine.api.entities.v1.SchemaRef
	10, // 4: streammachine.api.schemas.v1.GetSchemaRequest.cluster_ref:type_name -> streammachine.api.entities.v1.KafkaClusterRef
	8,  // 5: streammachine.api.schemas.v1.GetSchemaResponse.schema:type_name -> streammachine.api.entities.v1.Schema
	9,  // 6: streammachine.api.schemas.v1.GetSchemaCodeRequest.ref:type_name -> streammachine.api.entities.v1.SchemaRef
	11, // 7: streammachine.api.schemas.v1.GetSchemaCodeRequest.type:type_name -> streammachine.api.entities.v1.SchemaType
	0,  // 8: streammachine.api.schemas.v1.SchemasService.ListSchemas:input_type -> streammachine.api.schemas.v1.ListSchemasRequest
	4,  // 9: streammachine.api.schemas.v1.SchemasService.GetSchema:input_type -> streammachine.api.schemas.v1.GetSchemaRequest
	2,  // 10: streammachine.api.schemas.v1.SchemasService.CreateSchema:input_type -> streammachine.api.schemas.v1.CreateSchemaRequest
	6,  // 11: streammachine.api.schemas.v1.SchemasService.GetSchemaCode:input_type -> streammachine.api.schemas.v1.GetSchemaCodeRequest
	1,  // 12: streammachine.api.schemas.v1.SchemasService.ListSchemas:output_type -> streammachine.api.schemas.v1.ListSchemasResponse
	5,  // 13: streammachine.api.schemas.v1.SchemasService.GetSchema:output_type -> streammachine.api.schemas.v1.GetSchemaResponse
	3,  // 14: streammachine.api.schemas.v1.SchemasService.CreateSchema:output_type -> streammachine.api.schemas.v1.CreateSchemaResponse
	7,  // 15: streammachine.api.schemas.v1.SchemasService.GetSchemaCode:output_type -> streammachine.api.schemas.v1.GetSchemaCodeResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_streammachine_api_schemas_v1_schemas_v1_proto_init() }
func file_streammachine_api_schemas_v1_schemas_v1_proto_init() {
	if File_streammachine_api_schemas_v1_schemas_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchemasRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchemasResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSchemaRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSchemaResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchemaRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchemaResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchemaCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchemaCodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_streammachine_api_schemas_v1_schemas_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_streammachine_api_schemas_v1_schemas_v1_proto_goTypes,
		DependencyIndexes: file_streammachine_api_schemas_v1_schemas_v1_proto_depIdxs,
		MessageInfos:      file_streammachine_api_schemas_v1_schemas_v1_proto_msgTypes,
	}.Build()
	File_streammachine_api_schemas_v1_schemas_v1_proto = out.File
	file_streammachine_api_schemas_v1_schemas_v1_proto_rawDesc = nil
	file_streammachine_api_schemas_v1_schemas_v1_proto_goTypes = nil
	file_streammachine_api_schemas_v1_schemas_v1_proto_depIdxs = nil
}
