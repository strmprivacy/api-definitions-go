// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.8
// source: strmprivacy/api/batch_exporters/v1/batch_exporters_v1.proto

package batch_exporters

import (
	v1 "github.com/strmprivacy/api-definitions-go/api/entities/v1"
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

type ListBatchExportersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
}

func (x *ListBatchExportersRequest) Reset() {
	*x = ListBatchExportersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBatchExportersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBatchExportersRequest) ProtoMessage() {}

func (x *ListBatchExportersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBatchExportersRequest.ProtoReflect.Descriptor instead.
func (*ListBatchExportersRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescGZIP(), []int{0}
}

func (x *ListBatchExportersRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

type ListBatchExportersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchExporters []*v1.BatchExporter `protobuf:"bytes,1,rep,name=batch_exporters,json=batchExporters,proto3" json:"batch_exporters,omitempty"`
}

func (x *ListBatchExportersResponse) Reset() {
	*x = ListBatchExportersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBatchExportersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBatchExportersResponse) ProtoMessage() {}

func (x *ListBatchExportersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBatchExportersResponse.ProtoReflect.Descriptor instead.
func (*ListBatchExportersResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListBatchExportersResponse) GetBatchExporters() []*v1.BatchExporter {
	if x != nil {
		return x.BatchExporters
	}
	return nil
}

type DeleteBatchExporterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.BatchExporterRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *DeleteBatchExporterRequest) Reset() {
	*x = DeleteBatchExporterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBatchExporterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBatchExporterRequest) ProtoMessage() {}

func (x *DeleteBatchExporterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBatchExporterRequest.ProtoReflect.Descriptor instead.
func (*DeleteBatchExporterRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBatchExporterRequest) GetRef() *v1.BatchExporterRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type DeleteBatchExporterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBatchExporterResponse) Reset() {
	*x = DeleteBatchExporterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBatchExporterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBatchExporterResponse) ProtoMessage() {}

func (x *DeleteBatchExporterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBatchExporterResponse.ProtoReflect.Descriptor instead.
func (*DeleteBatchExporterResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescGZIP(), []int{3}
}

type CreateBatchExporterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchExporter *v1.BatchExporter `protobuf:"bytes,1,opt,name=batch_exporter,json=batchExporter,proto3" json:"batch_exporter,omitempty"`
}

func (x *CreateBatchExporterRequest) Reset() {
	*x = CreateBatchExporterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchExporterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchExporterRequest) ProtoMessage() {}

func (x *CreateBatchExporterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchExporterRequest.ProtoReflect.Descriptor instead.
func (*CreateBatchExporterRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBatchExporterRequest) GetBatchExporter() *v1.BatchExporter {
	if x != nil {
		return x.BatchExporter
	}
	return nil
}

type CreateBatchExporterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchExporter *v1.BatchExporter `protobuf:"bytes,1,opt,name=batch_exporter,json=batchExporter,proto3" json:"batch_exporter,omitempty"`
}

func (x *CreateBatchExporterResponse) Reset() {
	*x = CreateBatchExporterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchExporterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchExporterResponse) ProtoMessage() {}

func (x *CreateBatchExporterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchExporterResponse.ProtoReflect.Descriptor instead.
func (*CreateBatchExporterResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBatchExporterResponse) GetBatchExporter() *v1.BatchExporter {
	if x != nil {
		return x.BatchExporter
	}
	return nil
}

type GetBatchExporterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.BatchExporterRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *GetBatchExporterRequest) Reset() {
	*x = GetBatchExporterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchExporterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchExporterRequest) ProtoMessage() {}

func (x *GetBatchExporterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchExporterRequest.ProtoReflect.Descriptor instead.
func (*GetBatchExporterRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetBatchExporterRequest) GetRef() *v1.BatchExporterRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type GetBatchExporterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchExporter *v1.BatchExporter `protobuf:"bytes,1,opt,name=batch_exporter,json=batchExporter,proto3" json:"batch_exporter,omitempty"`
}

func (x *GetBatchExporterResponse) Reset() {
	*x = GetBatchExporterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchExporterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchExporterResponse) ProtoMessage() {}

func (x *GetBatchExporterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchExporterResponse.ProtoReflect.Descriptor instead.
func (*GetBatchExporterResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescGZIP(), []int{7}
}

func (x *GetBatchExporterResponse) GetBatchExporter() *v1.BatchExporter {
	if x != nil {
		return x.BatchExporter
	}
	return nil
}

var File_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3a, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x71, 0x0a,
	0x1a, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x52, 0x0e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73,
	0x22, 0x62, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44,
	0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x03, 0x72, 0x65, 0x66, 0x22, 0x1d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x6f, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x51, 0x0a, 0x0e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x0d, 0x62, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x0d, 0x62, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x22, 0x5f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x44, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x6d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x0d, 0x62, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x32, 0xef, 0x04, 0x0a, 0x15, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x93, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3d, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x3e,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x96, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x3e, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x7b, 0x0a, 0x25, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescData = file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDesc
)

func file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDescData
}

var file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_goTypes = []interface{}{
	(*ListBatchExportersRequest)(nil),   // 0: strmprivacy.api.batch_exporters.v1.ListBatchExportersRequest
	(*ListBatchExportersResponse)(nil),  // 1: strmprivacy.api.batch_exporters.v1.ListBatchExportersResponse
	(*DeleteBatchExporterRequest)(nil),  // 2: strmprivacy.api.batch_exporters.v1.DeleteBatchExporterRequest
	(*DeleteBatchExporterResponse)(nil), // 3: strmprivacy.api.batch_exporters.v1.DeleteBatchExporterResponse
	(*CreateBatchExporterRequest)(nil),  // 4: strmprivacy.api.batch_exporters.v1.CreateBatchExporterRequest
	(*CreateBatchExporterResponse)(nil), // 5: strmprivacy.api.batch_exporters.v1.CreateBatchExporterResponse
	(*GetBatchExporterRequest)(nil),     // 6: strmprivacy.api.batch_exporters.v1.GetBatchExporterRequest
	(*GetBatchExporterResponse)(nil),    // 7: strmprivacy.api.batch_exporters.v1.GetBatchExporterResponse
	(*v1.BatchExporter)(nil),            // 8: strmprivacy.api.entities.v1.BatchExporter
	(*v1.BatchExporterRef)(nil),         // 9: strmprivacy.api.entities.v1.BatchExporterRef
}
var file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_depIdxs = []int32{
	8,  // 0: strmprivacy.api.batch_exporters.v1.ListBatchExportersResponse.batch_exporters:type_name -> strmprivacy.api.entities.v1.BatchExporter
	9,  // 1: strmprivacy.api.batch_exporters.v1.DeleteBatchExporterRequest.ref:type_name -> strmprivacy.api.entities.v1.BatchExporterRef
	8,  // 2: strmprivacy.api.batch_exporters.v1.CreateBatchExporterRequest.batch_exporter:type_name -> strmprivacy.api.entities.v1.BatchExporter
	8,  // 3: strmprivacy.api.batch_exporters.v1.CreateBatchExporterResponse.batch_exporter:type_name -> strmprivacy.api.entities.v1.BatchExporter
	9,  // 4: strmprivacy.api.batch_exporters.v1.GetBatchExporterRequest.ref:type_name -> strmprivacy.api.entities.v1.BatchExporterRef
	8,  // 5: strmprivacy.api.batch_exporters.v1.GetBatchExporterResponse.batch_exporter:type_name -> strmprivacy.api.entities.v1.BatchExporter
	0,  // 6: strmprivacy.api.batch_exporters.v1.BatchExportersService.ListBatchExporters:input_type -> strmprivacy.api.batch_exporters.v1.ListBatchExportersRequest
	6,  // 7: strmprivacy.api.batch_exporters.v1.BatchExportersService.GetBatchExporter:input_type -> strmprivacy.api.batch_exporters.v1.GetBatchExporterRequest
	2,  // 8: strmprivacy.api.batch_exporters.v1.BatchExportersService.DeleteBatchExporter:input_type -> strmprivacy.api.batch_exporters.v1.DeleteBatchExporterRequest
	4,  // 9: strmprivacy.api.batch_exporters.v1.BatchExportersService.CreateBatchExporter:input_type -> strmprivacy.api.batch_exporters.v1.CreateBatchExporterRequest
	1,  // 10: strmprivacy.api.batch_exporters.v1.BatchExportersService.ListBatchExporters:output_type -> strmprivacy.api.batch_exporters.v1.ListBatchExportersResponse
	7,  // 11: strmprivacy.api.batch_exporters.v1.BatchExportersService.GetBatchExporter:output_type -> strmprivacy.api.batch_exporters.v1.GetBatchExporterResponse
	3,  // 12: strmprivacy.api.batch_exporters.v1.BatchExportersService.DeleteBatchExporter:output_type -> strmprivacy.api.batch_exporters.v1.DeleteBatchExporterResponse
	5,  // 13: strmprivacy.api.batch_exporters.v1.BatchExportersService.CreateBatchExporter:output_type -> strmprivacy.api.batch_exporters.v1.CreateBatchExporterResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_init() }
func file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_init() {
	if File_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBatchExportersRequest); i {
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
		file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBatchExportersResponse); i {
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
		file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBatchExporterRequest); i {
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
		file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBatchExporterResponse); i {
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
		file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchExporterRequest); i {
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
		file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchExporterResponse); i {
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
		file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchExporterRequest); i {
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
		file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchExporterResponse); i {
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
			RawDescriptor: file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto = out.File
	file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_rawDesc = nil
	file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_goTypes = nil
	file_strmprivacy_api_batch_exporters_v1_batch_exporters_v1_proto_depIdxs = nil
}
