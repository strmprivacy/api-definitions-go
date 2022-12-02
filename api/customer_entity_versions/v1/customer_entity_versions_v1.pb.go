// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: strmprivacy/api/customer_entity_versions/v1/customer_entity_versions_v1.proto

package customer_entity_versions

import (
	v1 "github.com/strmprivacy/api-definitions-go/v2/api/installations/v1"
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

type GetBatchExporterImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationType v1.InstallationType `protobuf:"varint,1,opt,name=installation_type,json=installationType,proto3,enum=strmprivacy.api.installations.v1.InstallationType" json:"installation_type,omitempty"`
}

func (x *GetBatchExporterImageRequest) Reset() {
	*x = GetBatchExporterImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchExporterImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchExporterImageRequest) ProtoMessage() {}

func (x *GetBatchExporterImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchExporterImageRequest.ProtoReflect.Descriptor instead.
func (*GetBatchExporterImageRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescGZIP(), []int{0}
}

func (x *GetBatchExporterImageRequest) GetInstallationType() v1.InstallationType {
	if x != nil {
		return x.InstallationType
	}
	return v1.InstallationType(0)
}

type GetBatchExporterImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Deprecated: Do not use.
	AwsMarketplaceImage string `protobuf:"bytes,2,opt,name=aws_marketplace_image,json=awsMarketplaceImage,proto3" json:"aws_marketplace_image,omitempty"`
	ImageTag            string `protobuf:"bytes,3,opt,name=image_tag,json=imageTag,proto3" json:"image_tag,omitempty"`
}

func (x *GetBatchExporterImageResponse) Reset() {
	*x = GetBatchExporterImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchExporterImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchExporterImageResponse) ProtoMessage() {}

func (x *GetBatchExporterImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchExporterImageResponse.ProtoReflect.Descriptor instead.
func (*GetBatchExporterImageResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetBatchExporterImageResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// Deprecated: Do not use.
func (x *GetBatchExporterImageResponse) GetAwsMarketplaceImage() string {
	if x != nil {
		return x.AwsMarketplaceImage
	}
	return ""
}

func (x *GetBatchExporterImageResponse) GetImageTag() string {
	if x != nil {
		return x.ImageTag
	}
	return ""
}

type GetKafkaExporterImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetKafkaExporterImageRequest) Reset() {
	*x = GetKafkaExporterImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKafkaExporterImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaExporterImageRequest) ProtoMessage() {}

func (x *GetKafkaExporterImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaExporterImageRequest.ProtoReflect.Descriptor instead.
func (*GetKafkaExporterImageRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescGZIP(), []int{2}
}

type GetKafkaExporterImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *GetKafkaExporterImageResponse) Reset() {
	*x = GetKafkaExporterImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKafkaExporterImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaExporterImageResponse) ProtoMessage() {}

func (x *GetKafkaExporterImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaExporterImageResponse.ProtoReflect.Descriptor instead.
func (*GetKafkaExporterImageResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescGZIP(), []int{3}
}

func (x *GetKafkaExporterImageResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type GetDecrypterImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationType v1.InstallationType `protobuf:"varint,1,opt,name=installation_type,json=installationType,proto3,enum=strmprivacy.api.installations.v1.InstallationType" json:"installation_type,omitempty"`
}

func (x *GetDecrypterImageRequest) Reset() {
	*x = GetDecrypterImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDecrypterImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDecrypterImageRequest) ProtoMessage() {}

func (x *GetDecrypterImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDecrypterImageRequest.ProtoReflect.Descriptor instead.
func (*GetDecrypterImageRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescGZIP(), []int{4}
}

func (x *GetDecrypterImageRequest) GetInstallationType() v1.InstallationType {
	if x != nil {
		return x.InstallationType
	}
	return v1.InstallationType(0)
}

type GetDecrypterImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Deprecated: Do not use.
	AwsMarketplaceImage string `protobuf:"bytes,2,opt,name=aws_marketplace_image,json=awsMarketplaceImage,proto3" json:"aws_marketplace_image,omitempty"`
	ImageTag            string `protobuf:"bytes,3,opt,name=image_tag,json=imageTag,proto3" json:"image_tag,omitempty"`
}

func (x *GetDecrypterImageResponse) Reset() {
	*x = GetDecrypterImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDecrypterImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDecrypterImageResponse) ProtoMessage() {}

func (x *GetDecrypterImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDecrypterImageResponse.ProtoReflect.Descriptor instead.
func (*GetDecrypterImageResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescGZIP(), []int{5}
}

func (x *GetDecrypterImageResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// Deprecated: Do not use.
func (x *GetDecrypterImageResponse) GetAwsMarketplaceImage() string {
	if x != nil {
		return x.AwsMarketplaceImage
	}
	return ""
}

func (x *GetDecrypterImageResponse) GetImageTag() string {
	if x != nil {
		return x.ImageTag
	}
	return ""
}

type GetBatchJobRunnerImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationType v1.InstallationType `protobuf:"varint,1,opt,name=installation_type,json=installationType,proto3,enum=strmprivacy.api.installations.v1.InstallationType" json:"installation_type,omitempty"`
}

func (x *GetBatchJobRunnerImageRequest) Reset() {
	*x = GetBatchJobRunnerImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchJobRunnerImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchJobRunnerImageRequest) ProtoMessage() {}

func (x *GetBatchJobRunnerImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchJobRunnerImageRequest.ProtoReflect.Descriptor instead.
func (*GetBatchJobRunnerImageRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescGZIP(), []int{6}
}

func (x *GetBatchJobRunnerImageRequest) GetInstallationType() v1.InstallationType {
	if x != nil {
		return x.InstallationType
	}
	return v1.InstallationType(0)
}

type GetBatchJobRunnerImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// Deprecated: Do not use.
	AwsMarketplaceImage string `protobuf:"bytes,2,opt,name=aws_marketplace_image,json=awsMarketplaceImage,proto3" json:"aws_marketplace_image,omitempty"`
	ImageTag            string `protobuf:"bytes,3,opt,name=image_tag,json=imageTag,proto3" json:"image_tag,omitempty"`
}

func (x *GetBatchJobRunnerImageResponse) Reset() {
	*x = GetBatchJobRunnerImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatchJobRunnerImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatchJobRunnerImageResponse) ProtoMessage() {}

func (x *GetBatchJobRunnerImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatchJobRunnerImageResponse.ProtoReflect.Descriptor instead.
func (*GetBatchJobRunnerImageResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescGZIP(), []int{7}
}

func (x *GetBatchJobRunnerImageResponse) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

// Deprecated: Do not use.
func (x *GetBatchJobRunnerImageResponse) GetAwsMarketplaceImage() string {
	if x != nil {
		return x.AwsMarketplaceImage
	}
	return ""
}

func (x *GetBatchJobRunnerImageResponse) GetImageTag() string {
	if x != nil {
		return x.ImageTag
	}
	return ""
}

var File_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDesc = []byte{
	0x0a, 0x4d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2b, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x84, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x64, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x15, 0x61, 0x77, 0x73, 0x5f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x18, 0x01, 0xe0, 0x41, 0x02, 0x52, 0x13, 0x61, 0x77, 0x73,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x20, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54,
	0x61, 0x67, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3a, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x80,
	0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x72, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x64, 0x0a, 0x11, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x93, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x15, 0x61, 0x77,
	0x73, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x18, 0x01, 0xe0, 0x41, 0x02,
	0x52, 0x13, 0x61, 0x77, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x54, 0x61, 0x67, 0x22, 0x85, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x64, 0x0a, 0x11, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x98, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a,
	0x15, 0x61, 0x77, 0x73, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x18, 0x01,
	0xe0, 0x41, 0x02, 0x52, 0x13, 0x61, 0x77, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x61, 0x67, 0x32, 0xda, 0x05, 0x0a, 0x1d, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xae, 0x01, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x49, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x4a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xae, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x49, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa2,
	0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x72, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x45, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x72, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0xb1, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x4a,
	0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4b, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x4a, 0x6f, 0x62, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x99, 0x01, 0x0a, 0x2e, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescData = file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDesc
)

func file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDescData
}

var file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_goTypes = []interface{}{
	(*GetBatchExporterImageRequest)(nil),   // 0: strmprivacy.api.customer_entity_versions.v1.GetBatchExporterImageRequest
	(*GetBatchExporterImageResponse)(nil),  // 1: strmprivacy.api.customer_entity_versions.v1.GetBatchExporterImageResponse
	(*GetKafkaExporterImageRequest)(nil),   // 2: strmprivacy.api.customer_entity_versions.v1.GetKafkaExporterImageRequest
	(*GetKafkaExporterImageResponse)(nil),  // 3: strmprivacy.api.customer_entity_versions.v1.GetKafkaExporterImageResponse
	(*GetDecrypterImageRequest)(nil),       // 4: strmprivacy.api.customer_entity_versions.v1.GetDecrypterImageRequest
	(*GetDecrypterImageResponse)(nil),      // 5: strmprivacy.api.customer_entity_versions.v1.GetDecrypterImageResponse
	(*GetBatchJobRunnerImageRequest)(nil),  // 6: strmprivacy.api.customer_entity_versions.v1.GetBatchJobRunnerImageRequest
	(*GetBatchJobRunnerImageResponse)(nil), // 7: strmprivacy.api.customer_entity_versions.v1.GetBatchJobRunnerImageResponse
	(v1.InstallationType)(0),               // 8: strmprivacy.api.installations.v1.InstallationType
}
var file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_depIdxs = []int32{
	8, // 0: strmprivacy.api.customer_entity_versions.v1.GetBatchExporterImageRequest.installation_type:type_name -> strmprivacy.api.installations.v1.InstallationType
	8, // 1: strmprivacy.api.customer_entity_versions.v1.GetDecrypterImageRequest.installation_type:type_name -> strmprivacy.api.installations.v1.InstallationType
	8, // 2: strmprivacy.api.customer_entity_versions.v1.GetBatchJobRunnerImageRequest.installation_type:type_name -> strmprivacy.api.installations.v1.InstallationType
	0, // 3: strmprivacy.api.customer_entity_versions.v1.CustomerEntityVersionsService.GetBatchExporterImage:input_type -> strmprivacy.api.customer_entity_versions.v1.GetBatchExporterImageRequest
	2, // 4: strmprivacy.api.customer_entity_versions.v1.CustomerEntityVersionsService.GetKafkaExporterImage:input_type -> strmprivacy.api.customer_entity_versions.v1.GetKafkaExporterImageRequest
	4, // 5: strmprivacy.api.customer_entity_versions.v1.CustomerEntityVersionsService.GetDecrypterImage:input_type -> strmprivacy.api.customer_entity_versions.v1.GetDecrypterImageRequest
	6, // 6: strmprivacy.api.customer_entity_versions.v1.CustomerEntityVersionsService.GetBatchJobRunnerImage:input_type -> strmprivacy.api.customer_entity_versions.v1.GetBatchJobRunnerImageRequest
	1, // 7: strmprivacy.api.customer_entity_versions.v1.CustomerEntityVersionsService.GetBatchExporterImage:output_type -> strmprivacy.api.customer_entity_versions.v1.GetBatchExporterImageResponse
	3, // 8: strmprivacy.api.customer_entity_versions.v1.CustomerEntityVersionsService.GetKafkaExporterImage:output_type -> strmprivacy.api.customer_entity_versions.v1.GetKafkaExporterImageResponse
	5, // 9: strmprivacy.api.customer_entity_versions.v1.CustomerEntityVersionsService.GetDecrypterImage:output_type -> strmprivacy.api.customer_entity_versions.v1.GetDecrypterImageResponse
	7, // 10: strmprivacy.api.customer_entity_versions.v1.CustomerEntityVersionsService.GetBatchJobRunnerImage:output_type -> strmprivacy.api.customer_entity_versions.v1.GetBatchJobRunnerImageResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() {
	file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_init()
}
func file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_init() {
	if File_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchExporterImageRequest); i {
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
		file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchExporterImageResponse); i {
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
		file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKafkaExporterImageRequest); i {
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
		file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKafkaExporterImageResponse); i {
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
		file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDecrypterImageRequest); i {
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
		file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDecrypterImageResponse); i {
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
		file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchJobRunnerImageRequest); i {
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
		file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatchJobRunnerImageResponse); i {
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
			RawDescriptor: file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto = out.File
	file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_rawDesc = nil
	file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_goTypes = nil
	file_strmprivacy_api_customer_entity_versions_v1_customer_entity_versions_v1_proto_depIdxs = nil
}
