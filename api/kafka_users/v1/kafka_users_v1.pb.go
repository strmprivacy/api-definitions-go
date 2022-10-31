// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: strmprivacy/api/kafka_users/v1/kafka_users_v1.proto

package kafka_users

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

type ListKafkaUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (-- api-linter: core::0122::name-suffix=disabled
	//     aip.dev/not-precedent: ref is a better name than Google suggests ¯\_(ツ)_/¯ --)
	Ref *v1.KafkaExporterRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *ListKafkaUsersRequest) Reset() {
	*x = ListKafkaUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKafkaUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKafkaUsersRequest) ProtoMessage() {}

func (x *ListKafkaUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKafkaUsersRequest.ProtoReflect.Descriptor instead.
func (*ListKafkaUsersRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescGZIP(), []int{0}
}

func (x *ListKafkaUsersRequest) GetRef() *v1.KafkaExporterRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type ListKafkaUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KafkaUsers []*v1.KafkaUser `protobuf:"bytes,1,rep,name=kafka_users,json=kafkaUsers,proto3" json:"kafka_users,omitempty"`
}

func (x *ListKafkaUsersResponse) Reset() {
	*x = ListKafkaUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKafkaUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKafkaUsersResponse) ProtoMessage() {}

func (x *ListKafkaUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKafkaUsersResponse.ProtoReflect.Descriptor instead.
func (*ListKafkaUsersResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListKafkaUsersResponse) GetKafkaUsers() []*v1.KafkaUser {
	if x != nil {
		return x.KafkaUsers
	}
	return nil
}

type DeleteKafkaUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.KafkaUserRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *DeleteKafkaUserRequest) Reset() {
	*x = DeleteKafkaUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKafkaUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKafkaUserRequest) ProtoMessage() {}

func (x *DeleteKafkaUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKafkaUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteKafkaUserRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteKafkaUserRequest) GetRef() *v1.KafkaUserRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type DeleteKafkaUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteKafkaUserResponse) Reset() {
	*x = DeleteKafkaUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKafkaUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKafkaUserResponse) ProtoMessage() {}

func (x *DeleteKafkaUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKafkaUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteKafkaUserResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescGZIP(), []int{3}
}

type CreateKafkaUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KafkaUser *v1.KafkaUser `protobuf:"bytes,1,opt,name=kafka_user,json=kafkaUser,proto3" json:"kafka_user,omitempty"`
}

func (x *CreateKafkaUserRequest) Reset() {
	*x = CreateKafkaUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKafkaUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKafkaUserRequest) ProtoMessage() {}

func (x *CreateKafkaUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKafkaUserRequest.ProtoReflect.Descriptor instead.
func (*CreateKafkaUserRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescGZIP(), []int{4}
}

func (x *CreateKafkaUserRequest) GetKafkaUser() *v1.KafkaUser {
	if x != nil {
		return x.KafkaUser
	}
	return nil
}

type CreateKafkaUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KafkaUser *v1.KafkaUser `protobuf:"bytes,1,opt,name=kafka_user,json=kafkaUser,proto3" json:"kafka_user,omitempty"`
}

func (x *CreateKafkaUserResponse) Reset() {
	*x = CreateKafkaUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKafkaUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKafkaUserResponse) ProtoMessage() {}

func (x *CreateKafkaUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKafkaUserResponse.ProtoReflect.Descriptor instead.
func (*CreateKafkaUserResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CreateKafkaUserResponse) GetKafkaUser() *v1.KafkaUser {
	if x != nil {
		return x.KafkaUser
	}
	return nil
}

type GetKafkaUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.KafkaUserRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *GetKafkaUserRequest) Reset() {
	*x = GetKafkaUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKafkaUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaUserRequest) ProtoMessage() {}

func (x *GetKafkaUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaUserRequest.ProtoReflect.Descriptor instead.
func (*GetKafkaUserRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetKafkaUserRequest) GetRef() *v1.KafkaUserRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type GetKafkaUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KafkaUser *v1.KafkaUser `protobuf:"bytes,1,opt,name=kafka_user,json=kafkaUser,proto3" json:"kafka_user,omitempty"`
}

func (x *GetKafkaUserResponse) Reset() {
	*x = GetKafkaUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKafkaUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaUserResponse) ProtoMessage() {}

func (x *GetKafkaUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaUserResponse.ProtoReflect.Descriptor instead.
func (*GetKafkaUserResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescGZIP(), []int{7}
}

func (x *GetKafkaUserResponse) GetKafkaUser() *v1.KafkaUser {
	if x != nil {
		return x.KafkaUser
	}
	return nil
}

var File_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDesc = []byte{
	0x0a, 0x33, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44,
	0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x03, 0x72, 0x65, 0x66, 0x22, 0x61, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x0b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0a, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x5a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03,
	0x72, 0x65, 0x66, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5f,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x22,
	0x60, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x57, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x5d, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x32, 0x99, 0x04, 0x0a, 0x11, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x7f, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x35, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x79, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x33, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x36, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x82, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x36, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x72, 0x0a, 0x21, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b,
	0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescData = file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDesc
)

func file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDescData
}

var file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_goTypes = []interface{}{
	(*ListKafkaUsersRequest)(nil),   // 0: strmprivacy.api.kafka_users.v1.ListKafkaUsersRequest
	(*ListKafkaUsersResponse)(nil),  // 1: strmprivacy.api.kafka_users.v1.ListKafkaUsersResponse
	(*DeleteKafkaUserRequest)(nil),  // 2: strmprivacy.api.kafka_users.v1.DeleteKafkaUserRequest
	(*DeleteKafkaUserResponse)(nil), // 3: strmprivacy.api.kafka_users.v1.DeleteKafkaUserResponse
	(*CreateKafkaUserRequest)(nil),  // 4: strmprivacy.api.kafka_users.v1.CreateKafkaUserRequest
	(*CreateKafkaUserResponse)(nil), // 5: strmprivacy.api.kafka_users.v1.CreateKafkaUserResponse
	(*GetKafkaUserRequest)(nil),     // 6: strmprivacy.api.kafka_users.v1.GetKafkaUserRequest
	(*GetKafkaUserResponse)(nil),    // 7: strmprivacy.api.kafka_users.v1.GetKafkaUserResponse
	(*v1.KafkaExporterRef)(nil),     // 8: strmprivacy.api.entities.v1.KafkaExporterRef
	(*v1.KafkaUser)(nil),            // 9: strmprivacy.api.entities.v1.KafkaUser
	(*v1.KafkaUserRef)(nil),         // 10: strmprivacy.api.entities.v1.KafkaUserRef
}
var file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_depIdxs = []int32{
	8,  // 0: strmprivacy.api.kafka_users.v1.ListKafkaUsersRequest.ref:type_name -> strmprivacy.api.entities.v1.KafkaExporterRef
	9,  // 1: strmprivacy.api.kafka_users.v1.ListKafkaUsersResponse.kafka_users:type_name -> strmprivacy.api.entities.v1.KafkaUser
	10, // 2: strmprivacy.api.kafka_users.v1.DeleteKafkaUserRequest.ref:type_name -> strmprivacy.api.entities.v1.KafkaUserRef
	9,  // 3: strmprivacy.api.kafka_users.v1.CreateKafkaUserRequest.kafka_user:type_name -> strmprivacy.api.entities.v1.KafkaUser
	9,  // 4: strmprivacy.api.kafka_users.v1.CreateKafkaUserResponse.kafka_user:type_name -> strmprivacy.api.entities.v1.KafkaUser
	10, // 5: strmprivacy.api.kafka_users.v1.GetKafkaUserRequest.ref:type_name -> strmprivacy.api.entities.v1.KafkaUserRef
	9,  // 6: strmprivacy.api.kafka_users.v1.GetKafkaUserResponse.kafka_user:type_name -> strmprivacy.api.entities.v1.KafkaUser
	0,  // 7: strmprivacy.api.kafka_users.v1.KafkaUsersService.ListKafkaUsers:input_type -> strmprivacy.api.kafka_users.v1.ListKafkaUsersRequest
	6,  // 8: strmprivacy.api.kafka_users.v1.KafkaUsersService.GetKafkaUser:input_type -> strmprivacy.api.kafka_users.v1.GetKafkaUserRequest
	2,  // 9: strmprivacy.api.kafka_users.v1.KafkaUsersService.DeleteKafkaUser:input_type -> strmprivacy.api.kafka_users.v1.DeleteKafkaUserRequest
	4,  // 10: strmprivacy.api.kafka_users.v1.KafkaUsersService.CreateKafkaUser:input_type -> strmprivacy.api.kafka_users.v1.CreateKafkaUserRequest
	1,  // 11: strmprivacy.api.kafka_users.v1.KafkaUsersService.ListKafkaUsers:output_type -> strmprivacy.api.kafka_users.v1.ListKafkaUsersResponse
	7,  // 12: strmprivacy.api.kafka_users.v1.KafkaUsersService.GetKafkaUser:output_type -> strmprivacy.api.kafka_users.v1.GetKafkaUserResponse
	3,  // 13: strmprivacy.api.kafka_users.v1.KafkaUsersService.DeleteKafkaUser:output_type -> strmprivacy.api.kafka_users.v1.DeleteKafkaUserResponse
	5,  // 14: strmprivacy.api.kafka_users.v1.KafkaUsersService.CreateKafkaUser:output_type -> strmprivacy.api.kafka_users.v1.CreateKafkaUserResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_init() }
func file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_init() {
	if File_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKafkaUsersRequest); i {
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
		file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKafkaUsersResponse); i {
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
		file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKafkaUserRequest); i {
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
		file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKafkaUserResponse); i {
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
		file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKafkaUserRequest); i {
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
		file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKafkaUserResponse); i {
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
		file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKafkaUserRequest); i {
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
		file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKafkaUserResponse); i {
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
			RawDescriptor: file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto = out.File
	file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_rawDesc = nil
	file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_goTypes = nil
	file_strmprivacy_api_kafka_users_v1_kafka_users_v1_proto_depIdxs = nil
}
