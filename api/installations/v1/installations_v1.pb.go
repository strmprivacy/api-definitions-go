// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: strmprivacy/api/installations/v1/installations_v1.proto

package installations

import (
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

type CreateInstallationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installation *Installation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
}

func (x *CreateInstallationRequest) Reset() {
	*x = CreateInstallationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstallationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstallationRequest) ProtoMessage() {}

func (x *CreateInstallationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstallationRequest.ProtoReflect.Descriptor instead.
func (*CreateInstallationRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInstallationRequest) GetInstallation() *Installation {
	if x != nil {
		return x.Installation
	}
	return nil
}

type CreateInstallationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installation *Installation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
}

func (x *CreateInstallationResponse) Reset() {
	*x = CreateInstallationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInstallationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInstallationResponse) ProtoMessage() {}

func (x *CreateInstallationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInstallationResponse.ProtoReflect.Descriptor instead.
func (*CreateInstallationResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{1}
}

func (x *CreateInstallationResponse) GetInstallation() *Installation {
	if x != nil {
		return x.Installation
	}
	return nil
}

type DeleteInstallationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInstallationRequest) Reset() {
	*x = DeleteInstallationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInstallationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInstallationRequest) ProtoMessage() {}

func (x *DeleteInstallationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInstallationRequest.ProtoReflect.Descriptor instead.
func (*DeleteInstallationRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteInstallationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteInstallationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteInstallationResponse) Reset() {
	*x = DeleteInstallationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInstallationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInstallationResponse) ProtoMessage() {}

func (x *DeleteInstallationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInstallationResponse.ProtoReflect.Descriptor instead.
func (*DeleteInstallationResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{3}
}

type GetInstallationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetInstallationRequest) Reset() {
	*x = GetInstallationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstallationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstallationRequest) ProtoMessage() {}

func (x *GetInstallationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstallationRequest.ProtoReflect.Descriptor instead.
func (*GetInstallationRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{4}
}

func (x *GetInstallationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetInstallationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installation *Installation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
}

func (x *GetInstallationResponse) Reset() {
	*x = GetInstallationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstallationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstallationResponse) ProtoMessage() {}

func (x *GetInstallationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstallationResponse.ProtoReflect.Descriptor instead.
func (*GetInstallationResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{5}
}

func (x *GetInstallationResponse) GetInstallation() *Installation {
	if x != nil {
		return x.Installation
	}
	return nil
}

type ListInstallationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListInstallationsRequest) Reset() {
	*x = ListInstallationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstallationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstallationsRequest) ProtoMessage() {}

func (x *ListInstallationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstallationsRequest.ProtoReflect.Descriptor instead.
func (*ListInstallationsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{6}
}

type ListInstallationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Installations []*Installation `protobuf:"bytes,1,rep,name=installations,proto3" json:"installations,omitempty"`
}

func (x *ListInstallationsResponse) Reset() {
	*x = ListInstallationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstallationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstallationsResponse) ProtoMessage() {}

func (x *ListInstallationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstallationsResponse.ProtoReflect.Descriptor instead.
func (*ListInstallationsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{7}
}

func (x *ListInstallationsResponse) GetInstallations() []*Installation {
	if x != nil {
		return x.Installations
	}
	return nil
}

var File_strmprivacy_api_installations_v1_installations_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_installations_v1_installations_v1_proto_rawDesc = []byte{
	0x0a, 0x37, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x74, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a,
	0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x1c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x76, 0x0a, 0x19,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0xd2, 0x04, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8f, 0x01,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x8f, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x86, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x39, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x3a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x78, 0x0a, 0x23, 0x69, 0x6f, 0x2e,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescData = file_strmprivacy_api_installations_v1_installations_v1_proto_rawDesc
)

func file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescData
}

var file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_strmprivacy_api_installations_v1_installations_v1_proto_goTypes = []interface{}{
	(*CreateInstallationRequest)(nil),  // 0: strmprivacy.api.installations.v1.CreateInstallationRequest
	(*CreateInstallationResponse)(nil), // 1: strmprivacy.api.installations.v1.CreateInstallationResponse
	(*DeleteInstallationRequest)(nil),  // 2: strmprivacy.api.installations.v1.DeleteInstallationRequest
	(*DeleteInstallationResponse)(nil), // 3: strmprivacy.api.installations.v1.DeleteInstallationResponse
	(*GetInstallationRequest)(nil),     // 4: strmprivacy.api.installations.v1.GetInstallationRequest
	(*GetInstallationResponse)(nil),    // 5: strmprivacy.api.installations.v1.GetInstallationResponse
	(*ListInstallationsRequest)(nil),   // 6: strmprivacy.api.installations.v1.ListInstallationsRequest
	(*ListInstallationsResponse)(nil),  // 7: strmprivacy.api.installations.v1.ListInstallationsResponse
	(*Installation)(nil),               // 8: strmprivacy.api.installations.v1.Installation
}
var file_strmprivacy_api_installations_v1_installations_v1_proto_depIdxs = []int32{
	8, // 0: strmprivacy.api.installations.v1.CreateInstallationRequest.installation:type_name -> strmprivacy.api.installations.v1.Installation
	8, // 1: strmprivacy.api.installations.v1.CreateInstallationResponse.installation:type_name -> strmprivacy.api.installations.v1.Installation
	8, // 2: strmprivacy.api.installations.v1.GetInstallationResponse.installation:type_name -> strmprivacy.api.installations.v1.Installation
	8, // 3: strmprivacy.api.installations.v1.ListInstallationsResponse.installations:type_name -> strmprivacy.api.installations.v1.Installation
	0, // 4: strmprivacy.api.installations.v1.InstallationsService.CreateInstallation:input_type -> strmprivacy.api.installations.v1.CreateInstallationRequest
	2, // 5: strmprivacy.api.installations.v1.InstallationsService.DeleteInstallation:input_type -> strmprivacy.api.installations.v1.DeleteInstallationRequest
	4, // 6: strmprivacy.api.installations.v1.InstallationsService.GetInstallation:input_type -> strmprivacy.api.installations.v1.GetInstallationRequest
	6, // 7: strmprivacy.api.installations.v1.InstallationsService.ListInstallations:input_type -> strmprivacy.api.installations.v1.ListInstallationsRequest
	1, // 8: strmprivacy.api.installations.v1.InstallationsService.CreateInstallation:output_type -> strmprivacy.api.installations.v1.CreateInstallationResponse
	3, // 9: strmprivacy.api.installations.v1.InstallationsService.DeleteInstallation:output_type -> strmprivacy.api.installations.v1.DeleteInstallationResponse
	5, // 10: strmprivacy.api.installations.v1.InstallationsService.GetInstallation:output_type -> strmprivacy.api.installations.v1.GetInstallationResponse
	7, // 11: strmprivacy.api.installations.v1.InstallationsService.ListInstallations:output_type -> strmprivacy.api.installations.v1.ListInstallationsResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_installations_v1_installations_v1_proto_init() }
func file_strmprivacy_api_installations_v1_installations_v1_proto_init() {
	if File_strmprivacy_api_installations_v1_installations_v1_proto != nil {
		return
	}
	file_strmprivacy_api_installations_v1_entities_v1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstallationRequest); i {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInstallationResponse); i {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInstallationRequest); i {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInstallationResponse); i {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInstallationRequest); i {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInstallationResponse); i {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstallationsRequest); i {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstallationsResponse); i {
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
			RawDescriptor: file_strmprivacy_api_installations_v1_installations_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_installations_v1_installations_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_installations_v1_installations_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_installations_v1_installations_v1_proto = out.File
	file_strmprivacy_api_installations_v1_installations_v1_proto_rawDesc = nil
	file_strmprivacy_api_installations_v1_installations_v1_proto_goTypes = nil
	file_strmprivacy_api_installations_v1_installations_v1_proto_depIdxs = nil
}
