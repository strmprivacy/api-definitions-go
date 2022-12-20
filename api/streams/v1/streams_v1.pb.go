// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: strmprivacy/api/streams/v1/streams_v1.proto

package streams

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

type ListStreamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	BillingId string `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	Recursive bool   `protobuf:"varint,2,opt,name=recursive,proto3" json:"recursive,omitempty"`
	ProjectId string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ListStreamsRequest) Reset() {
	*x = ListStreamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStreamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStreamsRequest) ProtoMessage() {}

func (x *ListStreamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStreamsRequest.ProtoReflect.Descriptor instead.
func (*ListStreamsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *ListStreamsRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *ListStreamsRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

func (x *ListStreamsRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ListStreamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streams []*v1.StreamTree `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty"`
}

func (x *ListStreamsResponse) Reset() {
	*x = ListStreamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStreamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStreamsResponse) ProtoMessage() {}

func (x *ListStreamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStreamsResponse.ProtoReflect.Descriptor instead.
func (*ListStreamsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListStreamsResponse) GetStreams() []*v1.StreamTree {
	if x != nil {
		return x.Streams
	}
	return nil
}

type ListExtendedStreamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientIds []string `protobuf:"bytes,1,rep,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
}

func (x *ListExtendedStreamsRequest) Reset() {
	*x = ListExtendedStreamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExtendedStreamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExtendedStreamsRequest) ProtoMessage() {}

func (x *ListExtendedStreamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExtendedStreamsRequest.ProtoReflect.Descriptor instead.
func (*ListExtendedStreamsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{2}
}

func (x *ListExtendedStreamsRequest) GetClientIds() []string {
	if x != nil {
		return x.ClientIds
	}
	return nil
}

type ListExtendedStreamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtendedStreams []*v1.ExtendedStream `protobuf:"bytes,1,rep,name=extended_streams,json=extendedStreams,proto3" json:"extended_streams,omitempty"`
}

func (x *ListExtendedStreamsResponse) Reset() {
	*x = ListExtendedStreamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExtendedStreamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExtendedStreamsResponse) ProtoMessage() {}

func (x *ListExtendedStreamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExtendedStreamsResponse.ProtoReflect.Descriptor instead.
func (*ListExtendedStreamsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{3}
}

func (x *ListExtendedStreamsResponse) GetExtendedStreams() []*v1.ExtendedStream {
	if x != nil {
		return x.ExtendedStreams
	}
	return nil
}

type DeleteStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.StreamRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// (-- api-linter: core::0135::request-unknown-fields=disabled
	//     aip.dev/not-precedent: We really need this field. --)
	Recursive bool `protobuf:"varint,2,opt,name=recursive,proto3" json:"recursive,omitempty"`
}

func (x *DeleteStreamRequest) Reset() {
	*x = DeleteStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStreamRequest) ProtoMessage() {}

func (x *DeleteStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStreamRequest.ProtoReflect.Descriptor instead.
func (*DeleteStreamRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteStreamRequest) GetRef() *v1.StreamRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *DeleteStreamRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

type DeleteStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteStreamResponse) Reset() {
	*x = DeleteStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStreamResponse) ProtoMessage() {}

func (x *DeleteStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStreamResponse.ProtoReflect.Descriptor instead.
func (*DeleteStreamResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{5}
}

type CreateStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream *v1.Stream `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *CreateStreamRequest) Reset() {
	*x = CreateStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStreamRequest) ProtoMessage() {}

func (x *CreateStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStreamRequest.ProtoReflect.Descriptor instead.
func (*CreateStreamRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{6}
}

func (x *CreateStreamRequest) GetStream() *v1.Stream {
	if x != nil {
		return x.Stream
	}
	return nil
}

type CreateStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream *v1.Stream `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *CreateStreamResponse) Reset() {
	*x = CreateStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStreamResponse) ProtoMessage() {}

func (x *CreateStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStreamResponse.ProtoReflect.Descriptor instead.
func (*CreateStreamResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{7}
}

func (x *CreateStreamResponse) GetStream() *v1.Stream {
	if x != nil {
		return x.Stream
	}
	return nil
}

type GetStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref       *v1.StreamRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Recursive bool          `protobuf:"varint,2,opt,name=recursive,proto3" json:"recursive,omitempty"`
}

func (x *GetStreamRequest) Reset() {
	*x = GetStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamRequest) ProtoMessage() {}

func (x *GetStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamRequest.ProtoReflect.Descriptor instead.
func (*GetStreamRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{8}
}

func (x *GetStreamRequest) GetRef() *v1.StreamRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *GetStreamRequest) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

type GetStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamTree *v1.StreamTree `protobuf:"bytes,1,opt,name=stream_tree,json=streamTree,proto3" json:"stream_tree,omitempty"`
}

func (x *GetStreamResponse) Reset() {
	*x = GetStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamResponse) ProtoMessage() {}

func (x *GetStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamResponse.ProtoReflect.Descriptor instead.
func (*GetStreamResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP(), []int{9}
}

func (x *GetStreamResponse) GetStreamTree() *v1.StreamTree {
	if x != nil {
		return x.StreamTree
	}
	return nil
}

var File_strmprivacy_api_streams_v1_streams_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_streams_v1_streams_v1_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22,
	0x58, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x65, 0x65,
	0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x3b, 0x0a, 0x1a, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22, 0x75, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x0f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x77, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03,
	0x72, 0x65, 0x66, 0x12, 0x21, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x57,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x53, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x74, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3d, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12,
	0x21, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x22, 0x5d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x54, 0x72, 0x65, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x72, 0x65,
	0x65, 0x32, 0xd9, 0x04, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x12, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x36, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2c, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2f, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x66, 0x0a,
	0x1d, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescData = file_strmprivacy_api_streams_v1_streams_v1_proto_rawDesc
)

func file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_streams_v1_streams_v1_proto_rawDescData
}

var file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_strmprivacy_api_streams_v1_streams_v1_proto_goTypes = []interface{}{
	(*ListStreamsRequest)(nil),          // 0: strmprivacy.api.streams.v1.ListStreamsRequest
	(*ListStreamsResponse)(nil),         // 1: strmprivacy.api.streams.v1.ListStreamsResponse
	(*ListExtendedStreamsRequest)(nil),  // 2: strmprivacy.api.streams.v1.ListExtendedStreamsRequest
	(*ListExtendedStreamsResponse)(nil), // 3: strmprivacy.api.streams.v1.ListExtendedStreamsResponse
	(*DeleteStreamRequest)(nil),         // 4: strmprivacy.api.streams.v1.DeleteStreamRequest
	(*DeleteStreamResponse)(nil),        // 5: strmprivacy.api.streams.v1.DeleteStreamResponse
	(*CreateStreamRequest)(nil),         // 6: strmprivacy.api.streams.v1.CreateStreamRequest
	(*CreateStreamResponse)(nil),        // 7: strmprivacy.api.streams.v1.CreateStreamResponse
	(*GetStreamRequest)(nil),            // 8: strmprivacy.api.streams.v1.GetStreamRequest
	(*GetStreamResponse)(nil),           // 9: strmprivacy.api.streams.v1.GetStreamResponse
	(*v1.StreamTree)(nil),               // 10: strmprivacy.api.entities.v1.StreamTree
	(*v1.ExtendedStream)(nil),           // 11: strmprivacy.api.entities.v1.ExtendedStream
	(*v1.StreamRef)(nil),                // 12: strmprivacy.api.entities.v1.StreamRef
	(*v1.Stream)(nil),                   // 13: strmprivacy.api.entities.v1.Stream
}
var file_strmprivacy_api_streams_v1_streams_v1_proto_depIdxs = []int32{
	10, // 0: strmprivacy.api.streams.v1.ListStreamsResponse.streams:type_name -> strmprivacy.api.entities.v1.StreamTree
	11, // 1: strmprivacy.api.streams.v1.ListExtendedStreamsResponse.extended_streams:type_name -> strmprivacy.api.entities.v1.ExtendedStream
	12, // 2: strmprivacy.api.streams.v1.DeleteStreamRequest.ref:type_name -> strmprivacy.api.entities.v1.StreamRef
	13, // 3: strmprivacy.api.streams.v1.CreateStreamRequest.stream:type_name -> strmprivacy.api.entities.v1.Stream
	13, // 4: strmprivacy.api.streams.v1.CreateStreamResponse.stream:type_name -> strmprivacy.api.entities.v1.Stream
	12, // 5: strmprivacy.api.streams.v1.GetStreamRequest.ref:type_name -> strmprivacy.api.entities.v1.StreamRef
	10, // 6: strmprivacy.api.streams.v1.GetStreamResponse.stream_tree:type_name -> strmprivacy.api.entities.v1.StreamTree
	0,  // 7: strmprivacy.api.streams.v1.StreamsService.ListStreams:input_type -> strmprivacy.api.streams.v1.ListStreamsRequest
	2,  // 8: strmprivacy.api.streams.v1.StreamsService.ListExtendedStreams:input_type -> strmprivacy.api.streams.v1.ListExtendedStreamsRequest
	8,  // 9: strmprivacy.api.streams.v1.StreamsService.GetStream:input_type -> strmprivacy.api.streams.v1.GetStreamRequest
	4,  // 10: strmprivacy.api.streams.v1.StreamsService.DeleteStream:input_type -> strmprivacy.api.streams.v1.DeleteStreamRequest
	6,  // 11: strmprivacy.api.streams.v1.StreamsService.CreateStream:input_type -> strmprivacy.api.streams.v1.CreateStreamRequest
	1,  // 12: strmprivacy.api.streams.v1.StreamsService.ListStreams:output_type -> strmprivacy.api.streams.v1.ListStreamsResponse
	3,  // 13: strmprivacy.api.streams.v1.StreamsService.ListExtendedStreams:output_type -> strmprivacy.api.streams.v1.ListExtendedStreamsResponse
	9,  // 14: strmprivacy.api.streams.v1.StreamsService.GetStream:output_type -> strmprivacy.api.streams.v1.GetStreamResponse
	5,  // 15: strmprivacy.api.streams.v1.StreamsService.DeleteStream:output_type -> strmprivacy.api.streams.v1.DeleteStreamResponse
	7,  // 16: strmprivacy.api.streams.v1.StreamsService.CreateStream:output_type -> strmprivacy.api.streams.v1.CreateStreamResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_streams_v1_streams_v1_proto_init() }
func file_strmprivacy_api_streams_v1_streams_v1_proto_init() {
	if File_strmprivacy_api_streams_v1_streams_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStreamsRequest); i {
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
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStreamsResponse); i {
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
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExtendedStreamsRequest); i {
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
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExtendedStreamsResponse); i {
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
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStreamRequest); i {
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
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStreamResponse); i {
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
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStreamRequest); i {
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
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStreamResponse); i {
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
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamRequest); i {
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
		file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamResponse); i {
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
			RawDescriptor: file_strmprivacy_api_streams_v1_streams_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_streams_v1_streams_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_streams_v1_streams_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_streams_v1_streams_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_streams_v1_streams_v1_proto = out.File
	file_strmprivacy_api_streams_v1_streams_v1_proto_rawDesc = nil
	file_strmprivacy_api_streams_v1_streams_v1_proto_goTypes = nil
	file_strmprivacy_api_streams_v1_streams_v1_proto_depIdxs = nil
}
