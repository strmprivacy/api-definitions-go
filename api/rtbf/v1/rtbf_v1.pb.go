// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: strmprivacy/api/rtbf/v1/rtbf_v1.proto

package rtbf

import (
	_ "github.com/strmprivacy/api-definitions-go/v2/api/entities/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetKeyLinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RtbfId string `protobuf:"bytes,1,opt,name=rtbf_id,json=rtbfId,proto3" json:"rtbf_id,omitempty"`
}

func (x *GetKeyLinksRequest) Reset() {
	*x = GetKeyLinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyLinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyLinksRequest) ProtoMessage() {}

func (x *GetKeyLinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyLinksRequest.ProtoReflect.Descriptor instead.
func (*GetKeyLinksRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescGZIP(), []int{0}
}

func (x *GetKeyLinksRequest) GetRtbfId() string {
	if x != nil {
		return x.RtbfId
	}
	return ""
}

type GetKeyLinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RtbfId   string   `protobuf:"bytes,1,opt,name=rtbf_id,json=rtbfId,proto3" json:"rtbf_id,omitempty"`
	KeyLinks []string `protobuf:"bytes,2,rep,name=key_links,json=keyLinks,proto3" json:"key_links,omitempty"`
}

func (x *GetKeyLinksResponse) Reset() {
	*x = GetKeyLinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKeyLinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKeyLinksResponse) ProtoMessage() {}

func (x *GetKeyLinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKeyLinksResponse.ProtoReflect.Descriptor instead.
func (*GetKeyLinksResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetKeyLinksResponse) GetRtbfId() string {
	if x != nil {
		return x.RtbfId
	}
	return ""
}

func (x *GetKeyLinksResponse) GetKeyLinks() []string {
	if x != nil {
		return x.KeyLinks
	}
	return nil
}

type GetRtbfsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RtbfId string `protobuf:"bytes,1,opt,name=rtbf_id,json=rtbfId,proto3" json:"rtbf_id,omitempty"`
}

func (x *GetRtbfsRequest) Reset() {
	*x = GetRtbfsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRtbfsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRtbfsRequest) ProtoMessage() {}

func (x *GetRtbfsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRtbfsRequest.ProtoReflect.Descriptor instead.
func (*GetRtbfsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescGZIP(), []int{2}
}

func (x *GetRtbfsRequest) GetRtbfId() string {
	if x != nil {
		return x.RtbfId
	}
	return ""
}

type GetRtbfsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RtbfId string   `protobuf:"bytes,1,opt,name=rtbf_id,json=rtbfId,proto3" json:"rtbf_id,omitempty"`
	Rtbf   []string `protobuf:"bytes,2,rep,name=rtbf,proto3" json:"rtbf,omitempty"`
}

func (x *GetRtbfsResponse) Reset() {
	*x = GetRtbfsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRtbfsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRtbfsResponse) ProtoMessage() {}

func (x *GetRtbfsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRtbfsResponse.ProtoReflect.Descriptor instead.
func (*GetRtbfsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescGZIP(), []int{3}
}

func (x *GetRtbfsResponse) GetRtbfId() string {
	if x != nil {
		return x.RtbfId
	}
	return ""
}

func (x *GetRtbfsResponse) GetRtbf() []string {
	if x != nil {
		return x.Rtbf
	}
	return nil
}

type AddKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RtbfId        string `protobuf:"bytes,1,opt,name=rtbf_id,json=rtbfId,proto3" json:"rtbf_id,omitempty"`
	KeyLink       string `protobuf:"bytes,2,opt,name=key_link,json=keyLink,proto3" json:"key_link,omitempty"`
	RtbfLink      string `protobuf:"bytes,3,opt,name=rtbf_link,json=rtbfLink,proto3" json:"rtbf_link,omitempty"`
	EncryptionKey string `protobuf:"bytes,4,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty"`
}

func (x *AddKeyRequest) Reset() {
	*x = AddKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKeyRequest) ProtoMessage() {}

func (x *AddKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKeyRequest.ProtoReflect.Descriptor instead.
func (*AddKeyRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescGZIP(), []int{4}
}

func (x *AddKeyRequest) GetRtbfId() string {
	if x != nil {
		return x.RtbfId
	}
	return ""
}

func (x *AddKeyRequest) GetKeyLink() string {
	if x != nil {
		return x.KeyLink
	}
	return ""
}

func (x *AddKeyRequest) GetRtbfLink() string {
	if x != nil {
		return x.RtbfLink
	}
	return ""
}

func (x *AddKeyRequest) GetEncryptionKey() string {
	if x != nil {
		return x.EncryptionKey
	}
	return ""
}

type AddKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddKeyResponse) Reset() {
	*x = AddKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddKeyResponse) ProtoMessage() {}

func (x *AddKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddKeyResponse.ProtoReflect.Descriptor instead.
func (*AddKeyResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescGZIP(), []int{5}
}

var File_strmprivacy_api_rtbf_v1_rtbf_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x74, 0x62, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x74, 0x62, 0x66, 0x5f, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x74, 0x62, 0x66, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x74, 0x62, 0x66,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x74, 0x62, 0x66, 0x49,
	0x64, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x74, 0x62, 0x66,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x74, 0x62, 0x66, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x2a,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x74, 0x62, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x74, 0x62, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x74, 0x62, 0x66, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x52, 0x74, 0x62, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x74, 0x62, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x74, 0x62, 0x66, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x74, 0x62, 0x66, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x72, 0x74, 0x62, 0x66, 0x22, 0x96, 0x01, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x07, 0x72, 0x74, 0x62, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x06, 0x72, 0x74, 0x62, 0x66, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x6b,
	0x65, 0x79, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x09, 0x72,
	0x74, 0x62, 0x66, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x08, 0x72, 0x74, 0x62, 0x66, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x25, 0x0a,
	0x0e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb3, 0x02, 0x0a, 0x0b, 0x52, 0x74, 0x62, 0x66, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x2b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x74, 0x62, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x74, 0x62, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4b, 0x65, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x74, 0x62, 0x66, 0x73, 0x12, 0x28, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x74, 0x62, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x74, 0x62, 0x66, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x74, 0x62, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x74, 0x62, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x59, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x74,
	0x62, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x74, 0x62, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5d, 0x0a, 0x1a,
	0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x74, 0x62, 0x66, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x74, 0x62, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x74, 0x62, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescData = file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDesc
)

func file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDescData
}

var file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_goTypes = []interface{}{
	(*GetKeyLinksRequest)(nil),  // 0: strmprivacy.api.rtbf.v1.GetKeyLinksRequest
	(*GetKeyLinksResponse)(nil), // 1: strmprivacy.api.rtbf.v1.GetKeyLinksResponse
	(*GetRtbfsRequest)(nil),     // 2: strmprivacy.api.rtbf.v1.GetRtbfsRequest
	(*GetRtbfsResponse)(nil),    // 3: strmprivacy.api.rtbf.v1.GetRtbfsResponse
	(*AddKeyRequest)(nil),       // 4: strmprivacy.api.rtbf.v1.AddKeyRequest
	(*AddKeyResponse)(nil),      // 5: strmprivacy.api.rtbf.v1.AddKeyResponse
}
var file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_depIdxs = []int32{
	0, // 0: strmprivacy.api.rtbf.v1.RtbfService.GetKeyLinks:input_type -> strmprivacy.api.rtbf.v1.GetKeyLinksRequest
	2, // 1: strmprivacy.api.rtbf.v1.RtbfService.GetRtbfs:input_type -> strmprivacy.api.rtbf.v1.GetRtbfsRequest
	4, // 2: strmprivacy.api.rtbf.v1.RtbfService.AddKey:input_type -> strmprivacy.api.rtbf.v1.AddKeyRequest
	1, // 3: strmprivacy.api.rtbf.v1.RtbfService.GetKeyLinks:output_type -> strmprivacy.api.rtbf.v1.GetKeyLinksResponse
	3, // 4: strmprivacy.api.rtbf.v1.RtbfService.GetRtbfs:output_type -> strmprivacy.api.rtbf.v1.GetRtbfsResponse
	5, // 5: strmprivacy.api.rtbf.v1.RtbfService.AddKey:output_type -> strmprivacy.api.rtbf.v1.AddKeyResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_init() }
func file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_init() {
	if File_strmprivacy_api_rtbf_v1_rtbf_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyLinksRequest); i {
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
		file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKeyLinksResponse); i {
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
		file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRtbfsRequest); i {
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
		file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRtbfsResponse); i {
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
		file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKeyRequest); i {
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
		file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddKeyResponse); i {
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
			RawDescriptor: file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_rtbf_v1_rtbf_v1_proto = out.File
	file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_rawDesc = nil
	file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_goTypes = nil
	file_strmprivacy_api_rtbf_v1_rtbf_v1_proto_depIdxs = nil
}
