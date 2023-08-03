// (-- api-linter: core::0133::request-id-field=disabled
//     aip.dev/not-precedent: We need to do this because we typically use entities that contain
//     references inside them, and not an entity id directly in the request. --)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: strmprivacy/api/purpose_mapping/v1/purpose_mapping_v1.proto

package purpose_mapping

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/strmprivacy/api-definitions-go/v3/api/entities/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListPurposeMappingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPurposeMappingsRequest) Reset() {
	*x = ListPurposeMappingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPurposeMappingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPurposeMappingsRequest) ProtoMessage() {}

func (x *ListPurposeMappingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPurposeMappingsRequest.ProtoReflect.Descriptor instead.
func (*ListPurposeMappingsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescGZIP(), []int{0}
}

type ListPurposeMappingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PurposeMappings []*v1.PurposeMapping `protobuf:"bytes,1,rep,name=purpose_mappings,json=purposeMappings,proto3" json:"purpose_mappings,omitempty"`
}

func (x *ListPurposeMappingsResponse) Reset() {
	*x = ListPurposeMappingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPurposeMappingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPurposeMappingsResponse) ProtoMessage() {}

func (x *ListPurposeMappingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPurposeMappingsResponse.ProtoReflect.Descriptor instead.
func (*ListPurposeMappingsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListPurposeMappingsResponse) GetPurposeMappings() []*v1.PurposeMapping {
	if x != nil {
		return x.PurposeMappings
	}
	return nil
}

type CreatePurposeMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PurposeMapping *v1.PurposeMapping `protobuf:"bytes,1,opt,name=purpose_mapping,json=purposeMapping,proto3" json:"purpose_mapping,omitempty"`
}

func (x *CreatePurposeMappingRequest) Reset() {
	*x = CreatePurposeMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePurposeMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePurposeMappingRequest) ProtoMessage() {}

func (x *CreatePurposeMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePurposeMappingRequest.ProtoReflect.Descriptor instead.
func (*CreatePurposeMappingRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePurposeMappingRequest) GetPurposeMapping() *v1.PurposeMapping {
	if x != nil {
		return x.PurposeMapping
	}
	return nil
}

type CreatePurposeMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PurposeMapping *v1.PurposeMapping `protobuf:"bytes,1,opt,name=purpose_mapping,json=purposeMapping,proto3" json:"purpose_mapping,omitempty"`
}

func (x *CreatePurposeMappingResponse) Reset() {
	*x = CreatePurposeMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePurposeMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePurposeMappingResponse) ProtoMessage() {}

func (x *CreatePurposeMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePurposeMappingResponse.ProtoReflect.Descriptor instead.
func (*CreatePurposeMappingResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePurposeMappingResponse) GetPurposeMapping() *v1.PurposeMapping {
	if x != nil {
		return x.PurposeMapping
	}
	return nil
}

type GetPurposeMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *GetPurposeMappingRequest) Reset() {
	*x = GetPurposeMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPurposeMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPurposeMappingRequest) ProtoMessage() {}

func (x *GetPurposeMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPurposeMappingRequest.ProtoReflect.Descriptor instead.
func (*GetPurposeMappingRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescGZIP(), []int{4}
}

func (x *GetPurposeMappingRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type GetPurposeMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PurposeMapping *v1.PurposeMapping `protobuf:"bytes,1,opt,name=purpose_mapping,json=purposeMapping,proto3" json:"purpose_mapping,omitempty"`
}

func (x *GetPurposeMappingResponse) Reset() {
	*x = GetPurposeMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPurposeMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPurposeMappingResponse) ProtoMessage() {}

func (x *GetPurposeMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPurposeMappingResponse.ProtoReflect.Descriptor instead.
func (*GetPurposeMappingResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescGZIP(), []int{5}
}

func (x *GetPurposeMappingResponse) GetPurposeMapping() *v1.PurposeMapping {
	if x != nil {
		return x.PurposeMapping
	}
	return nil
}

var File_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x75, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x10, 0x70, 0x75, 0x72, 0x70, 0x6f,
	0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0f,
	0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x22,
	0x7d, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5e,
	0x0a, 0x0f, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0e,
	0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x79,
	0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59,
	0x0a, 0x0f, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x70, 0x75, 0x72, 0x70, 0x6f,
	0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x39, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x22, 0x76, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x75, 0x72, 0x70, 0x6f,
	0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x59, 0x0a, 0x0f, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x70, 0x75,
	0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x32, 0xdf, 0x03, 0x0a,
	0x15, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3e,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x90, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x72, 0x70,
	0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72,
	0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75,
	0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x7e,
	0x0a, 0x25, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x75, 0x72, 0x70,
	0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescData = file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDesc
)

func file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDescData
}

var file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_goTypes = []interface{}{
	(*ListPurposeMappingsRequest)(nil),   // 0: strmprivacy.api.purpose_mapping.v1.ListPurposeMappingsRequest
	(*ListPurposeMappingsResponse)(nil),  // 1: strmprivacy.api.purpose_mapping.v1.ListPurposeMappingsResponse
	(*CreatePurposeMappingRequest)(nil),  // 2: strmprivacy.api.purpose_mapping.v1.CreatePurposeMappingRequest
	(*CreatePurposeMappingResponse)(nil), // 3: strmprivacy.api.purpose_mapping.v1.CreatePurposeMappingResponse
	(*GetPurposeMappingRequest)(nil),     // 4: strmprivacy.api.purpose_mapping.v1.GetPurposeMappingRequest
	(*GetPurposeMappingResponse)(nil),    // 5: strmprivacy.api.purpose_mapping.v1.GetPurposeMappingResponse
	(*v1.PurposeMapping)(nil),            // 6: strmprivacy.api.entities.v1.PurposeMapping
}
var file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_depIdxs = []int32{
	6, // 0: strmprivacy.api.purpose_mapping.v1.ListPurposeMappingsResponse.purpose_mappings:type_name -> strmprivacy.api.entities.v1.PurposeMapping
	6, // 1: strmprivacy.api.purpose_mapping.v1.CreatePurposeMappingRequest.purpose_mapping:type_name -> strmprivacy.api.entities.v1.PurposeMapping
	6, // 2: strmprivacy.api.purpose_mapping.v1.CreatePurposeMappingResponse.purpose_mapping:type_name -> strmprivacy.api.entities.v1.PurposeMapping
	6, // 3: strmprivacy.api.purpose_mapping.v1.GetPurposeMappingResponse.purpose_mapping:type_name -> strmprivacy.api.entities.v1.PurposeMapping
	0, // 4: strmprivacy.api.purpose_mapping.v1.PurposeMappingService.ListPurposeMappings:input_type -> strmprivacy.api.purpose_mapping.v1.ListPurposeMappingsRequest
	4, // 5: strmprivacy.api.purpose_mapping.v1.PurposeMappingService.GetPurposeMapping:input_type -> strmprivacy.api.purpose_mapping.v1.GetPurposeMappingRequest
	2, // 6: strmprivacy.api.purpose_mapping.v1.PurposeMappingService.CreatePurposeMapping:input_type -> strmprivacy.api.purpose_mapping.v1.CreatePurposeMappingRequest
	1, // 7: strmprivacy.api.purpose_mapping.v1.PurposeMappingService.ListPurposeMappings:output_type -> strmprivacy.api.purpose_mapping.v1.ListPurposeMappingsResponse
	5, // 8: strmprivacy.api.purpose_mapping.v1.PurposeMappingService.GetPurposeMapping:output_type -> strmprivacy.api.purpose_mapping.v1.GetPurposeMappingResponse
	3, // 9: strmprivacy.api.purpose_mapping.v1.PurposeMappingService.CreatePurposeMapping:output_type -> strmprivacy.api.purpose_mapping.v1.CreatePurposeMappingResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_init() }
func file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_init() {
	if File_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPurposeMappingsRequest); i {
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
		file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPurposeMappingsResponse); i {
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
		file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePurposeMappingRequest); i {
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
		file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePurposeMappingResponse); i {
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
		file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPurposeMappingRequest); i {
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
		file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPurposeMappingResponse); i {
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
			RawDescriptor: file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto = out.File
	file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_rawDesc = nil
	file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_goTypes = nil
	file_strmprivacy_api_purpose_mapping_v1_purpose_mapping_v1_proto_depIdxs = nil
}
