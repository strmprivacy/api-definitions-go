// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: strmprivacy/api/consent_levels/v1/consent_levels_v1.proto

package consent_levels

import (
	v1 "github.com/strmprivacy/api-definitions-go/v2/api/entities/v1"
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

type ListConsentLevelMappingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	BillingId string `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ListConsentLevelMappingsRequest) Reset() {
	*x = ListConsentLevelMappingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConsentLevelMappingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConsentLevelMappingsRequest) ProtoMessage() {}

func (x *ListConsentLevelMappingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConsentLevelMappingsRequest.ProtoReflect.Descriptor instead.
func (*ListConsentLevelMappingsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *ListConsentLevelMappingsRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *ListConsentLevelMappingsRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ListConsentLevelMappingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsentLevelMappings []*v1.ConsentLevelMapping `protobuf:"bytes,1,rep,name=consent_level_mappings,json=consentLevelMappings,proto3" json:"consent_level_mappings,omitempty"`
}

func (x *ListConsentLevelMappingsResponse) Reset() {
	*x = ListConsentLevelMappingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConsentLevelMappingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConsentLevelMappingsResponse) ProtoMessage() {}

func (x *ListConsentLevelMappingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConsentLevelMappingsResponse.ProtoReflect.Descriptor instead.
func (*ListConsentLevelMappingsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListConsentLevelMappingsResponse) GetConsentLevelMappings() []*v1.ConsentLevelMapping {
	if x != nil {
		return x.ConsentLevelMappings
	}
	return nil
}

// (-- api-linter: core::0135::request-name-required=disabled
//     aip.dev/not-precedent: A consent-level has no name, but a level. --)
type DeleteConsentLevelMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (-- api-linter: core::0135::request-unknown-fields=disabled
	//     aip.dev/not-precedent: A consent-level has no name, but a level. --)
	Ref *v1.ConsentLevelMappingRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *DeleteConsentLevelMappingRequest) Reset() {
	*x = DeleteConsentLevelMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConsentLevelMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConsentLevelMappingRequest) ProtoMessage() {}

func (x *DeleteConsentLevelMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConsentLevelMappingRequest.ProtoReflect.Descriptor instead.
func (*DeleteConsentLevelMappingRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteConsentLevelMappingRequest) GetRef() *v1.ConsentLevelMappingRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type DeleteConsentLevelMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteConsentLevelMappingResponse) Reset() {
	*x = DeleteConsentLevelMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConsentLevelMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConsentLevelMappingResponse) ProtoMessage() {}

func (x *DeleteConsentLevelMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConsentLevelMappingResponse.ProtoReflect.Descriptor instead.
func (*DeleteConsentLevelMappingResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescGZIP(), []int{3}
}

type CreateConsentLevelMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsentLevelMapping *v1.ConsentLevelMapping `protobuf:"bytes,1,opt,name=consent_level_mapping,json=consentLevelMapping,proto3" json:"consent_level_mapping,omitempty"`
}

func (x *CreateConsentLevelMappingRequest) Reset() {
	*x = CreateConsentLevelMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsentLevelMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsentLevelMappingRequest) ProtoMessage() {}

func (x *CreateConsentLevelMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsentLevelMappingRequest.ProtoReflect.Descriptor instead.
func (*CreateConsentLevelMappingRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescGZIP(), []int{4}
}

func (x *CreateConsentLevelMappingRequest) GetConsentLevelMapping() *v1.ConsentLevelMapping {
	if x != nil {
		return x.ConsentLevelMapping
	}
	return nil
}

type CreateConsentLevelMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsentLevelMapping *v1.ConsentLevelMapping `protobuf:"bytes,1,opt,name=consent_level_mapping,json=consentLevelMapping,proto3" json:"consent_level_mapping,omitempty"`
}

func (x *CreateConsentLevelMappingResponse) Reset() {
	*x = CreateConsentLevelMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsentLevelMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsentLevelMappingResponse) ProtoMessage() {}

func (x *CreateConsentLevelMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsentLevelMappingResponse.ProtoReflect.Descriptor instead.
func (*CreateConsentLevelMappingResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CreateConsentLevelMappingResponse) GetConsentLevelMapping() *v1.ConsentLevelMapping {
	if x != nil {
		return x.ConsentLevelMapping
	}
	return nil
}

type GetConsentLevelMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.ConsentLevelMappingRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *GetConsentLevelMappingRequest) Reset() {
	*x = GetConsentLevelMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsentLevelMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsentLevelMappingRequest) ProtoMessage() {}

func (x *GetConsentLevelMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsentLevelMappingRequest.ProtoReflect.Descriptor instead.
func (*GetConsentLevelMappingRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetConsentLevelMappingRequest) GetRef() *v1.ConsentLevelMappingRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type GetConsentLevelMappingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsentLevelMapping *v1.ConsentLevelMapping `protobuf:"bytes,1,opt,name=consent_level_mapping,json=consentLevelMapping,proto3" json:"consent_level_mapping,omitempty"`
}

func (x *GetConsentLevelMappingResponse) Reset() {
	*x = GetConsentLevelMappingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsentLevelMappingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsentLevelMappingResponse) ProtoMessage() {}

func (x *GetConsentLevelMappingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsentLevelMappingResponse.ProtoReflect.Descriptor instead.
func (*GetConsentLevelMappingResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescGZIP(), []int{7}
}

func (x *GetConsentLevelMappingResponse) GetConsentLevelMapping() *v1.ConsentLevelMapping {
	if x != nil {
		return x.ConsentLevelMapping
	}
	return nil
}

var File_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDesc = []byte{
	0x0a, 0x39, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68,
	0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x20, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a,
	0x16, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x14, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x6e, 0x0a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x03, 0x72, 0x65, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x23, 0x0a, 0x21, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x20, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x64, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x89, 0x01, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x15, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x13, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x22, 0x6b, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4a, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x86,
	0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x64, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x32, 0xb5, 0x05, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x42, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9d, 0x01,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa6, 0x01,
	0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x43, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x44, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa6, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x12, 0x43, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x7b, 0x0a, 0x24, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescData = file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDesc
)

func file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDescData
}

var file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_goTypes = []interface{}{
	(*ListConsentLevelMappingsRequest)(nil),   // 0: strmprivacy.api.consent_levels.v1.ListConsentLevelMappingsRequest
	(*ListConsentLevelMappingsResponse)(nil),  // 1: strmprivacy.api.consent_levels.v1.ListConsentLevelMappingsResponse
	(*DeleteConsentLevelMappingRequest)(nil),  // 2: strmprivacy.api.consent_levels.v1.DeleteConsentLevelMappingRequest
	(*DeleteConsentLevelMappingResponse)(nil), // 3: strmprivacy.api.consent_levels.v1.DeleteConsentLevelMappingResponse
	(*CreateConsentLevelMappingRequest)(nil),  // 4: strmprivacy.api.consent_levels.v1.CreateConsentLevelMappingRequest
	(*CreateConsentLevelMappingResponse)(nil), // 5: strmprivacy.api.consent_levels.v1.CreateConsentLevelMappingResponse
	(*GetConsentLevelMappingRequest)(nil),     // 6: strmprivacy.api.consent_levels.v1.GetConsentLevelMappingRequest
	(*GetConsentLevelMappingResponse)(nil),    // 7: strmprivacy.api.consent_levels.v1.GetConsentLevelMappingResponse
	(*v1.ConsentLevelMapping)(nil),            // 8: strmprivacy.api.entities.v1.ConsentLevelMapping
	(*v1.ConsentLevelMappingRef)(nil),         // 9: strmprivacy.api.entities.v1.ConsentLevelMappingRef
}
var file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_depIdxs = []int32{
	8,  // 0: strmprivacy.api.consent_levels.v1.ListConsentLevelMappingsResponse.consent_level_mappings:type_name -> strmprivacy.api.entities.v1.ConsentLevelMapping
	9,  // 1: strmprivacy.api.consent_levels.v1.DeleteConsentLevelMappingRequest.ref:type_name -> strmprivacy.api.entities.v1.ConsentLevelMappingRef
	8,  // 2: strmprivacy.api.consent_levels.v1.CreateConsentLevelMappingRequest.consent_level_mapping:type_name -> strmprivacy.api.entities.v1.ConsentLevelMapping
	8,  // 3: strmprivacy.api.consent_levels.v1.CreateConsentLevelMappingResponse.consent_level_mapping:type_name -> strmprivacy.api.entities.v1.ConsentLevelMapping
	9,  // 4: strmprivacy.api.consent_levels.v1.GetConsentLevelMappingRequest.ref:type_name -> strmprivacy.api.entities.v1.ConsentLevelMappingRef
	8,  // 5: strmprivacy.api.consent_levels.v1.GetConsentLevelMappingResponse.consent_level_mapping:type_name -> strmprivacy.api.entities.v1.ConsentLevelMapping
	0,  // 6: strmprivacy.api.consent_levels.v1.ConsentLevelMappingsService.ListConsentLevelMappings:input_type -> strmprivacy.api.consent_levels.v1.ListConsentLevelMappingsRequest
	6,  // 7: strmprivacy.api.consent_levels.v1.ConsentLevelMappingsService.GetConsentLevelMapping:input_type -> strmprivacy.api.consent_levels.v1.GetConsentLevelMappingRequest
	2,  // 8: strmprivacy.api.consent_levels.v1.ConsentLevelMappingsService.DeleteConsentLevelMapping:input_type -> strmprivacy.api.consent_levels.v1.DeleteConsentLevelMappingRequest
	4,  // 9: strmprivacy.api.consent_levels.v1.ConsentLevelMappingsService.CreateConsentLevelMapping:input_type -> strmprivacy.api.consent_levels.v1.CreateConsentLevelMappingRequest
	1,  // 10: strmprivacy.api.consent_levels.v1.ConsentLevelMappingsService.ListConsentLevelMappings:output_type -> strmprivacy.api.consent_levels.v1.ListConsentLevelMappingsResponse
	7,  // 11: strmprivacy.api.consent_levels.v1.ConsentLevelMappingsService.GetConsentLevelMapping:output_type -> strmprivacy.api.consent_levels.v1.GetConsentLevelMappingResponse
	3,  // 12: strmprivacy.api.consent_levels.v1.ConsentLevelMappingsService.DeleteConsentLevelMapping:output_type -> strmprivacy.api.consent_levels.v1.DeleteConsentLevelMappingResponse
	5,  // 13: strmprivacy.api.consent_levels.v1.ConsentLevelMappingsService.CreateConsentLevelMapping:output_type -> strmprivacy.api.consent_levels.v1.CreateConsentLevelMappingResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_init() }
func file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_init() {
	if File_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConsentLevelMappingsRequest); i {
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
		file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConsentLevelMappingsResponse); i {
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
		file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConsentLevelMappingRequest); i {
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
		file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConsentLevelMappingResponse); i {
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
		file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsentLevelMappingRequest); i {
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
		file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsentLevelMappingResponse); i {
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
		file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsentLevelMappingRequest); i {
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
		file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsentLevelMappingResponse); i {
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
			RawDescriptor: file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto = out.File
	file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_rawDesc = nil
	file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_goTypes = nil
	file_strmprivacy_api_consent_levels_v1_consent_levels_v1_proto_depIdxs = nil
}
