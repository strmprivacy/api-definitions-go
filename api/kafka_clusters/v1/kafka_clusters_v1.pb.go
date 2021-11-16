// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.8
// source: strmprivacy/api/kafka_clusters/v1/kafka_clusters_v1.proto

package kafka_clusters

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

type ListKafkaClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
}

func (x *ListKafkaClustersRequest) Reset() {
	*x = ListKafkaClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKafkaClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKafkaClustersRequest) ProtoMessage() {}

func (x *ListKafkaClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKafkaClustersRequest.ProtoReflect.Descriptor instead.
func (*ListKafkaClustersRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescGZIP(), []int{0}
}

func (x *ListKafkaClustersRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

type ListKafkaClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KafkaClusters []*v1.KafkaCluster `protobuf:"bytes,1,rep,name=kafka_clusters,json=kafkaClusters,proto3" json:"kafka_clusters,omitempty"`
}

func (x *ListKafkaClustersResponse) Reset() {
	*x = ListKafkaClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKafkaClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKafkaClustersResponse) ProtoMessage() {}

func (x *ListKafkaClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKafkaClustersResponse.ProtoReflect.Descriptor instead.
func (*ListKafkaClustersResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListKafkaClustersResponse) GetKafkaClusters() []*v1.KafkaCluster {
	if x != nil {
		return x.KafkaClusters
	}
	return nil
}

type DeleteKafkaClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string              `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	Ref       *v1.KafkaClusterRef `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *DeleteKafkaClusterRequest) Reset() {
	*x = DeleteKafkaClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKafkaClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKafkaClusterRequest) ProtoMessage() {}

func (x *DeleteKafkaClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKafkaClusterRequest.ProtoReflect.Descriptor instead.
func (*DeleteKafkaClusterRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteKafkaClusterRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *DeleteKafkaClusterRequest) GetRef() *v1.KafkaClusterRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type DeleteKafkaClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteKafkaClusterResponse) Reset() {
	*x = DeleteKafkaClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKafkaClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKafkaClusterResponse) ProtoMessage() {}

func (x *DeleteKafkaClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKafkaClusterResponse.ProtoReflect.Descriptor instead.
func (*DeleteKafkaClusterResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescGZIP(), []int{3}
}

type CreateKafkaClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KafkaCluster *v1.KafkaCluster `protobuf:"bytes,1,opt,name=kafka_cluster,json=kafkaCluster,proto3" json:"kafka_cluster,omitempty"`
}

func (x *CreateKafkaClusterRequest) Reset() {
	*x = CreateKafkaClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKafkaClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKafkaClusterRequest) ProtoMessage() {}

func (x *CreateKafkaClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKafkaClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateKafkaClusterRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescGZIP(), []int{4}
}

func (x *CreateKafkaClusterRequest) GetKafkaCluster() *v1.KafkaCluster {
	if x != nil {
		return x.KafkaCluster
	}
	return nil
}

type CreateKafkaClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KafkaCluster *v1.KafkaCluster `protobuf:"bytes,1,opt,name=kafka_cluster,json=kafkaCluster,proto3" json:"kafka_cluster,omitempty"`
}

func (x *CreateKafkaClusterResponse) Reset() {
	*x = CreateKafkaClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKafkaClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKafkaClusterResponse) ProtoMessage() {}

func (x *CreateKafkaClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKafkaClusterResponse.ProtoReflect.Descriptor instead.
func (*CreateKafkaClusterResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescGZIP(), []int{5}
}

func (x *CreateKafkaClusterResponse) GetKafkaCluster() *v1.KafkaCluster {
	if x != nil {
		return x.KafkaCluster
	}
	return nil
}

type GetKafkaClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId string              `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	Ref       *v1.KafkaClusterRef `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (x *GetKafkaClusterRequest) Reset() {
	*x = GetKafkaClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKafkaClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaClusterRequest) ProtoMessage() {}

func (x *GetKafkaClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaClusterRequest.ProtoReflect.Descriptor instead.
func (*GetKafkaClusterRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetKafkaClusterRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *GetKafkaClusterRequest) GetRef() *v1.KafkaClusterRef {
	if x != nil {
		return x.Ref
	}
	return nil
}

type GetKafkaClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KafkaCluster *v1.KafkaCluster `protobuf:"bytes,1,opt,name=kafka_cluster,json=kafkaCluster,proto3" json:"kafka_cluster,omitempty"`
}

func (x *GetKafkaClusterResponse) Reset() {
	*x = GetKafkaClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKafkaClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaClusterResponse) ProtoMessage() {}

func (x *GetKafkaClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaClusterResponse.ProtoReflect.Descriptor instead.
func (*GetKafkaClusterResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescGZIP(), []int{7}
}

func (x *GetKafkaClusterResponse) GetKafkaCluster() *v1.KafkaCluster {
	if x != nil {
		return x.KafkaCluster
	}
	return nil
}

var File_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDesc = []byte{
	0x0a, 0x39, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x6d,
	0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0e, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0d,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x84, 0x01,
	0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12,
	0x43, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x03, 0x72, 0x65, 0x66, 0x22, 0x1c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4e, 0x0a, 0x0d, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x0c, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x6c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x0d, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x0c, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x81, 0x01,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x03,
	0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x66, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x72, 0x65,
	0x66, 0x22, 0x69, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0c,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x32, 0xda, 0x04, 0x0a,
	0x14, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x39, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x91, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x91, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x61, 0x66,
	0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x7b, 0x0a, 0x24, 0x69, 0x6f, 0x2e,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescData = file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDesc
)

func file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDescData
}

var file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_goTypes = []interface{}{
	(*ListKafkaClustersRequest)(nil),   // 0: strmprivacy.api.kafka_clusters.v1.ListKafkaClustersRequest
	(*ListKafkaClustersResponse)(nil),  // 1: strmprivacy.api.kafka_clusters.v1.ListKafkaClustersResponse
	(*DeleteKafkaClusterRequest)(nil),  // 2: strmprivacy.api.kafka_clusters.v1.DeleteKafkaClusterRequest
	(*DeleteKafkaClusterResponse)(nil), // 3: strmprivacy.api.kafka_clusters.v1.DeleteKafkaClusterResponse
	(*CreateKafkaClusterRequest)(nil),  // 4: strmprivacy.api.kafka_clusters.v1.CreateKafkaClusterRequest
	(*CreateKafkaClusterResponse)(nil), // 5: strmprivacy.api.kafka_clusters.v1.CreateKafkaClusterResponse
	(*GetKafkaClusterRequest)(nil),     // 6: strmprivacy.api.kafka_clusters.v1.GetKafkaClusterRequest
	(*GetKafkaClusterResponse)(nil),    // 7: strmprivacy.api.kafka_clusters.v1.GetKafkaClusterResponse
	(*v1.KafkaCluster)(nil),            // 8: strmprivacy.api.entities.v1.KafkaCluster
	(*v1.KafkaClusterRef)(nil),         // 9: strmprivacy.api.entities.v1.KafkaClusterRef
}
var file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_depIdxs = []int32{
	8,  // 0: strmprivacy.api.kafka_clusters.v1.ListKafkaClustersResponse.kafka_clusters:type_name -> strmprivacy.api.entities.v1.KafkaCluster
	9,  // 1: strmprivacy.api.kafka_clusters.v1.DeleteKafkaClusterRequest.ref:type_name -> strmprivacy.api.entities.v1.KafkaClusterRef
	8,  // 2: strmprivacy.api.kafka_clusters.v1.CreateKafkaClusterRequest.kafka_cluster:type_name -> strmprivacy.api.entities.v1.KafkaCluster
	8,  // 3: strmprivacy.api.kafka_clusters.v1.CreateKafkaClusterResponse.kafka_cluster:type_name -> strmprivacy.api.entities.v1.KafkaCluster
	9,  // 4: strmprivacy.api.kafka_clusters.v1.GetKafkaClusterRequest.ref:type_name -> strmprivacy.api.entities.v1.KafkaClusterRef
	8,  // 5: strmprivacy.api.kafka_clusters.v1.GetKafkaClusterResponse.kafka_cluster:type_name -> strmprivacy.api.entities.v1.KafkaCluster
	0,  // 6: strmprivacy.api.kafka_clusters.v1.KafkaClustersService.ListKafkaClusters:input_type -> strmprivacy.api.kafka_clusters.v1.ListKafkaClustersRequest
	6,  // 7: strmprivacy.api.kafka_clusters.v1.KafkaClustersService.GetKafkaCluster:input_type -> strmprivacy.api.kafka_clusters.v1.GetKafkaClusterRequest
	2,  // 8: strmprivacy.api.kafka_clusters.v1.KafkaClustersService.DeleteKafkaCluster:input_type -> strmprivacy.api.kafka_clusters.v1.DeleteKafkaClusterRequest
	4,  // 9: strmprivacy.api.kafka_clusters.v1.KafkaClustersService.CreateKafkaCluster:input_type -> strmprivacy.api.kafka_clusters.v1.CreateKafkaClusterRequest
	1,  // 10: strmprivacy.api.kafka_clusters.v1.KafkaClustersService.ListKafkaClusters:output_type -> strmprivacy.api.kafka_clusters.v1.ListKafkaClustersResponse
	7,  // 11: strmprivacy.api.kafka_clusters.v1.KafkaClustersService.GetKafkaCluster:output_type -> strmprivacy.api.kafka_clusters.v1.GetKafkaClusterResponse
	3,  // 12: strmprivacy.api.kafka_clusters.v1.KafkaClustersService.DeleteKafkaCluster:output_type -> strmprivacy.api.kafka_clusters.v1.DeleteKafkaClusterResponse
	5,  // 13: strmprivacy.api.kafka_clusters.v1.KafkaClustersService.CreateKafkaCluster:output_type -> strmprivacy.api.kafka_clusters.v1.CreateKafkaClusterResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_init() }
func file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_init() {
	if File_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKafkaClustersRequest); i {
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
		file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKafkaClustersResponse); i {
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
		file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKafkaClusterRequest); i {
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
		file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKafkaClusterResponse); i {
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
		file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKafkaClusterRequest); i {
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
		file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKafkaClusterResponse); i {
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
		file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKafkaClusterRequest); i {
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
		file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKafkaClusterResponse); i {
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
			RawDescriptor: file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto = out.File
	file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_rawDesc = nil
	file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_goTypes = nil
	file_strmprivacy_api_kafka_clusters_v1_kafka_clusters_v1_proto_depIdxs = nil
}
