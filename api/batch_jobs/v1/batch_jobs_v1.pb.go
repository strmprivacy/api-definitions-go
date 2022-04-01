// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: strmprivacy/api/batch_jobs/v1/batch_jobs_v1.proto

package batch_jobs

import (
	v1 "github.com/strmprivacy/api-definitions-go/v2/api/entities/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBatchJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.BatchJobRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *GetBatchJobRequest) Reset() {
	*x = GetBatchJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchJobRequest) ProtoMessage() {}

func (x *GetBatchJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchJobRequest.ProtoReflect.Descriptor instead.
func (*GetBatchJobRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{0}
}

func (x *GetBatchJobRequest) GetRef() *v1.BatchJobRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type GetBatchJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchJob *v1.BatchJob `protobuf:"bytes,1,opt,name=batch_job,json=batchJob,proto3" json:"batch_job,omitempty"`
}

func (x *GetBatchJobResponse) Reset() {
	*x = GetBatchJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchJobResponse) ProtoMessage() {}

func (x *GetBatchJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchJobResponse.ProtoReflect.Descriptor instead.
func (*GetBatchJobResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetBatchJobResponse) GetBatchJob() *v1.BatchJob {
	if x != nil {
		return x.BatchJob
	}
	return nil
}

type ListBatchJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
}

func (x *ListBatchJobsRequest) Reset() {
	*x = ListBatchJobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBatchJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBatchJobsRequest) ProtoMessage() {}

func (x *ListBatchJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBatchJobsRequest.ProtoReflect.Descriptor instead.
func (*ListBatchJobsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{2}
}

func (x *ListBatchJobsRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

type ListBatchJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchJobs []*v1.BatchJob `protobuf:"bytes,1,rep,name=batch_jobs,json=batchJobs,proto3" json:"batch_jobs,omitempty"`
}

func (x *ListBatchJobsResponse) Reset() {
	*x = ListBatchJobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBatchJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBatchJobsResponse) ProtoMessage() {}

func (x *ListBatchJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBatchJobsResponse.ProtoReflect.Descriptor instead.
func (*ListBatchJobsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{3}
}

func (x *ListBatchJobsResponse) GetBatchJobs() []*v1.BatchJob {
	if x != nil {
		return x.BatchJobs
	}
	return nil
}

type CreateBatchJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchJob *v1.BatchJob `protobuf:"bytes,1,opt,name=batch_job,json=batchJob,proto3" json:"batch_job,omitempty"`
}

func (x *CreateBatchJobRequest) Reset() {
	*x = CreateBatchJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchJobRequest) ProtoMessage() {}

func (x *CreateBatchJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchJobRequest.ProtoReflect.Descriptor instead.
func (*CreateBatchJobRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBatchJobRequest) GetBatchJob() *v1.BatchJob {
	if x != nil {
		return x.BatchJob
	}
	return nil
}

type CreateBatchJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchJob *v1.BatchJob `protobuf:"bytes,1,opt,name=batch_job,json=batchJob,proto3" json:"batch_job,omitempty"`
}

func (x *CreateBatchJobResponse) Reset() {
	*x = CreateBatchJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBatchJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBatchJobResponse) ProtoMessage() {}

func (x *CreateBatchJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBatchJobResponse.ProtoReflect.Descriptor instead.
func (*CreateBatchJobResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBatchJobResponse) GetBatchJob() *v1.BatchJob {
	if x != nil {
		return x.BatchJob
	}
	return nil
}

type DeleteBatchJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref *v1.BatchJobRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *DeleteBatchJobRequest) Reset() {
	*x = DeleteBatchJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBatchJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBatchJobRequest) ProtoMessage() {}

func (x *DeleteBatchJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBatchJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteBatchJobRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBatchJobRequest) GetRef() *v1.BatchJobRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type DeleteBatchJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBatchJobResponse) Reset() {
	*x = DeleteBatchJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBatchJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBatchJobResponse) ProtoMessage() {}

func (x *DeleteBatchJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBatchJobResponse.ProtoReflect.Descriptor instead.
func (*DeleteBatchJobResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{7}
}

// (-- api-linter: core::0134::request-mask-required=disabled
//     aip.dev/not-precedent: This RPC only updates one field. --)
type UpdateBatchJobStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref           *v1.BatchJobRef   `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	BatchJobState *v1.BatchJobState `protobuf:"bytes,2,opt,name=batch_job_state,json=batchJobState,proto3" json:"batch_job_state,omitempty"`
}

func (x *UpdateBatchJobStateRequest) Reset() {
	*x = UpdateBatchJobStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBatchJobStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBatchJobStateRequest) ProtoMessage() {}

func (x *UpdateBatchJobStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBatchJobStateRequest.ProtoReflect.Descriptor instead.
func (*UpdateBatchJobStateRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateBatchJobStateRequest) GetRef() *v1.BatchJobRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *UpdateBatchJobStateRequest) GetBatchJobState() *v1.BatchJobState {
	if x != nil {
		return x.BatchJobState
	}
	return nil
}

type UpdateBatchJobStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBatchJobStateResponse) Reset() {
	*x = UpdateBatchJobStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBatchJobStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBatchJobStateResponse) ProtoMessage() {}

func (x *UpdateBatchJobStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBatchJobStateResponse.ProtoReflect.Descriptor instead.
func (*UpdateBatchJobStateResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP(), []int{9}
}

var File_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDesc = []byte{
	0x0a, 0x31, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x03, 0x72, 0x65, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x66,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x5e, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x22, 0x3a, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x4a, 0x6f, 0x62, 0x73, 0x22, 0x60, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47,
	0x0a, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x22, 0x61, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x22, 0x58, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x03, 0x72, 0x65, 0x66, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb6,
	0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a,
	0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x57,
	0x0a, 0x0f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x62, 0x61, 0x74, 0x63, 0x68, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x91, 0x05, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4a, 0x6f, 0x62, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x12, 0x31, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x7a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f,
	0x62, 0x73, 0x12, 0x33, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x12,
	0x34, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f,
	0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7d, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x12, 0x34,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x39, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x0a, 0x20, 0x69, 0x6f,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescData = file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDesc
)

func file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDescData
}

var file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_goTypes = []interface{}{
	(*GetBatchJobRequest)(nil),          // 0: strmprivacy.api.batch_jobs.v1.GetBatchJobRequest
	(*GetBatchJobResponse)(nil),         // 1: strmprivacy.api.batch_jobs.v1.GetBatchJobResponse
	(*ListBatchJobsRequest)(nil),        // 2: strmprivacy.api.batch_jobs.v1.ListBatchJobsRequest
	(*ListBatchJobsResponse)(nil),       // 3: strmprivacy.api.batch_jobs.v1.ListBatchJobsResponse
	(*CreateBatchJobRequest)(nil),       // 4: strmprivacy.api.batch_jobs.v1.CreateBatchJobRequest
	(*CreateBatchJobResponse)(nil),      // 5: strmprivacy.api.batch_jobs.v1.CreateBatchJobResponse
	(*DeleteBatchJobRequest)(nil),       // 6: strmprivacy.api.batch_jobs.v1.DeleteBatchJobRequest
	(*DeleteBatchJobResponse)(nil),      // 7: strmprivacy.api.batch_jobs.v1.DeleteBatchJobResponse
	(*UpdateBatchJobStateRequest)(nil),  // 8: strmprivacy.api.batch_jobs.v1.UpdateBatchJobStateRequest
	(*UpdateBatchJobStateResponse)(nil), // 9: strmprivacy.api.batch_jobs.v1.UpdateBatchJobStateResponse
	(*v1.BatchJobRef)(nil),              // 10: strmprivacy.api.entities.v1.BatchJobRef
	(*v1.BatchJob)(nil),                 // 11: strmprivacy.api.entities.v1.BatchJob
	(*v1.BatchJobState)(nil),            // 12: strmprivacy.api.entities.v1.BatchJobState
}
var file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_depIdxs = []int32{
	10, // 0: strmprivacy.api.batch_jobs.v1.GetBatchJobRequest.ref:type_name -> strmprivacy.api.entities.v1.BatchJobRef
	11, // 1: strmprivacy.api.batch_jobs.v1.GetBatchJobResponse.batch_job:type_name -> strmprivacy.api.entities.v1.BatchJob
	11, // 2: strmprivacy.api.batch_jobs.v1.ListBatchJobsResponse.batch_jobs:type_name -> strmprivacy.api.entities.v1.BatchJob
	11, // 3: strmprivacy.api.batch_jobs.v1.CreateBatchJobRequest.batch_job:type_name -> strmprivacy.api.entities.v1.BatchJob
	11, // 4: strmprivacy.api.batch_jobs.v1.CreateBatchJobResponse.batch_job:type_name -> strmprivacy.api.entities.v1.BatchJob
	10, // 5: strmprivacy.api.batch_jobs.v1.DeleteBatchJobRequest.ref:type_name -> strmprivacy.api.entities.v1.BatchJobRef
	10, // 6: strmprivacy.api.batch_jobs.v1.UpdateBatchJobStateRequest.ref:type_name -> strmprivacy.api.entities.v1.BatchJobRef
	12, // 7: strmprivacy.api.batch_jobs.v1.UpdateBatchJobStateRequest.batch_job_state:type_name -> strmprivacy.api.entities.v1.BatchJobState
	0,  // 8: strmprivacy.api.batch_jobs.v1.BatchJobsService.GetBatchJob:input_type -> strmprivacy.api.batch_jobs.v1.GetBatchJobRequest
	2,  // 9: strmprivacy.api.batch_jobs.v1.BatchJobsService.ListBatchJobs:input_type -> strmprivacy.api.batch_jobs.v1.ListBatchJobsRequest
	4,  // 10: strmprivacy.api.batch_jobs.v1.BatchJobsService.CreateBatchJob:input_type -> strmprivacy.api.batch_jobs.v1.CreateBatchJobRequest
	6,  // 11: strmprivacy.api.batch_jobs.v1.BatchJobsService.DeleteBatchJob:input_type -> strmprivacy.api.batch_jobs.v1.DeleteBatchJobRequest
	8,  // 12: strmprivacy.api.batch_jobs.v1.BatchJobsService.UpdateBatchJobState:input_type -> strmprivacy.api.batch_jobs.v1.UpdateBatchJobStateRequest
	1,  // 13: strmprivacy.api.batch_jobs.v1.BatchJobsService.GetBatchJob:output_type -> strmprivacy.api.batch_jobs.v1.GetBatchJobResponse
	3,  // 14: strmprivacy.api.batch_jobs.v1.BatchJobsService.ListBatchJobs:output_type -> strmprivacy.api.batch_jobs.v1.ListBatchJobsResponse
	5,  // 15: strmprivacy.api.batch_jobs.v1.BatchJobsService.CreateBatchJob:output_type -> strmprivacy.api.batch_jobs.v1.CreateBatchJobResponse
	7,  // 16: strmprivacy.api.batch_jobs.v1.BatchJobsService.DeleteBatchJob:output_type -> strmprivacy.api.batch_jobs.v1.DeleteBatchJobResponse
	9,  // 17: strmprivacy.api.batch_jobs.v1.BatchJobsService.UpdateBatchJobState:output_type -> strmprivacy.api.batch_jobs.v1.UpdateBatchJobStateResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_init() }
func file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_init() {
	if File_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchJobRequest); i {
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
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchJobResponse); i {
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
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBatchJobsRequest); i {
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
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBatchJobsResponse); i {
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
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchJobRequest); i {
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
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBatchJobResponse); i {
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
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBatchJobRequest); i {
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
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBatchJobResponse); i {
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
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBatchJobStateRequest); i {
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
		file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBatchJobStateResponse); i {
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
			RawDescriptor: file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto = out.File
	file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_rawDesc = nil
	file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_goTypes = nil
	file_strmprivacy_api_batch_jobs_v1_batch_jobs_v1_proto_depIdxs = nil
}
