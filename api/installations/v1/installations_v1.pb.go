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

type GetInstallationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationId string `protobuf:"bytes,1,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
}

func (x *GetInstallationRequest) Reset() {
	*x = GetInstallationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstallationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstallationRequest) ProtoMessage() {}

func (x *GetInstallationRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetInstallationRequest.ProtoReflect.Descriptor instead.
func (*GetInstallationRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{0}
}

func (x *GetInstallationRequest) GetInstallationId() string {
	if x != nil {
		return x.InstallationId
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
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInstallationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstallationResponse) ProtoMessage() {}

func (x *GetInstallationResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetInstallationResponse.ProtoReflect.Descriptor instead.
func (*GetInstallationResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{1}
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
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstallationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstallationsRequest) ProtoMessage() {}

func (x *ListInstallationsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListInstallationsRequest.ProtoReflect.Descriptor instead.
func (*ListInstallationsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{2}
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
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstallationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstallationsResponse) ProtoMessage() {}

func (x *ListInstallationsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListInstallationsResponse.ProtoReflect.Descriptor instead.
func (*ListInstallationsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{3}
}

func (x *ListInstallationsResponse) GetInstallations() []*Installation {
	if x != nil {
		return x.Installations
	}
	return nil
}

type GetProjectInstallationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetProjectInstallationRequest) Reset() {
	*x = GetProjectInstallationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectInstallationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectInstallationRequest) ProtoMessage() {}

func (x *GetProjectInstallationRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetProjectInstallationRequest.ProtoReflect.Descriptor instead.
func (*GetProjectInstallationRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{4}
}

func (x *GetProjectInstallationRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetProjectInstallationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationId string `protobuf:"bytes,1,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
}

func (x *GetProjectInstallationResponse) Reset() {
	*x = GetProjectInstallationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectInstallationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectInstallationResponse) ProtoMessage() {}

func (x *GetProjectInstallationResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetProjectInstallationResponse.ProtoReflect.Descriptor instead.
func (*GetProjectInstallationResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{5}
}

func (x *GetProjectInstallationResponse) GetInstallationId() string {
	if x != nil {
		return x.InstallationId
	}
	return ""
}

type ListInstallationProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationId string `protobuf:"bytes,1,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
}

func (x *ListInstallationProjectsRequest) Reset() {
	*x = ListInstallationProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstallationProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstallationProjectsRequest) ProtoMessage() {}

func (x *ListInstallationProjectsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListInstallationProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListInstallationProjectsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{6}
}

func (x *ListInstallationProjectsRequest) GetInstallationId() string {
	if x != nil {
		return x.InstallationId
	}
	return ""
}

type ListInstallationProjectsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (-- api-linter: core::0132::response-unknown-fields=disabled
	//     aip.dev/not-precedent: this is best descriptive. --)
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds,proto3" json:"project_ids,omitempty"`
}

func (x *ListInstallationProjectsResponse) Reset() {
	*x = ListInstallationProjectsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstallationProjectsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstallationProjectsResponse) ProtoMessage() {}

func (x *ListInstallationProjectsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListInstallationProjectsResponse.ProtoReflect.Descriptor instead.
func (*ListInstallationProjectsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_installations_v1_installations_v1_proto_rawDescGZIP(), []int{7}
}

func (x *ListInstallationProjectsResponse) GetProjectIds() []string {
	if x != nil {
		return x.ProjectIds
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
	0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1a, 0x0a, 0x18,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x76, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x43, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x32, 0xf0, 0x04, 0x0a, 0x14,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x39, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x8c, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9b, 0x01, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x18, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x41, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x78,
	0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d,
	0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*GetInstallationRequest)(nil),           // 0: strmprivacy.api.installations.v1.GetInstallationRequest
	(*GetInstallationResponse)(nil),          // 1: strmprivacy.api.installations.v1.GetInstallationResponse
	(*ListInstallationsRequest)(nil),         // 2: strmprivacy.api.installations.v1.ListInstallationsRequest
	(*ListInstallationsResponse)(nil),        // 3: strmprivacy.api.installations.v1.ListInstallationsResponse
	(*GetProjectInstallationRequest)(nil),    // 4: strmprivacy.api.installations.v1.GetProjectInstallationRequest
	(*GetProjectInstallationResponse)(nil),   // 5: strmprivacy.api.installations.v1.GetProjectInstallationResponse
	(*ListInstallationProjectsRequest)(nil),  // 6: strmprivacy.api.installations.v1.ListInstallationProjectsRequest
	(*ListInstallationProjectsResponse)(nil), // 7: strmprivacy.api.installations.v1.ListInstallationProjectsResponse
	(*Installation)(nil),                     // 8: strmprivacy.api.installations.v1.Installation
}
var file_strmprivacy_api_installations_v1_installations_v1_proto_depIdxs = []int32{
	8, // 0: strmprivacy.api.installations.v1.GetInstallationResponse.installation:type_name -> strmprivacy.api.installations.v1.Installation
	8, // 1: strmprivacy.api.installations.v1.ListInstallationsResponse.installations:type_name -> strmprivacy.api.installations.v1.Installation
	0, // 2: strmprivacy.api.installations.v1.InstallationsService.GetInstallation:input_type -> strmprivacy.api.installations.v1.GetInstallationRequest
	2, // 3: strmprivacy.api.installations.v1.InstallationsService.ListInstallations:input_type -> strmprivacy.api.installations.v1.ListInstallationsRequest
	4, // 4: strmprivacy.api.installations.v1.InstallationsService.GetProjectInstallation:input_type -> strmprivacy.api.installations.v1.GetProjectInstallationRequest
	6, // 5: strmprivacy.api.installations.v1.InstallationsService.ListInstallationProjects:input_type -> strmprivacy.api.installations.v1.ListInstallationProjectsRequest
	1, // 6: strmprivacy.api.installations.v1.InstallationsService.GetInstallation:output_type -> strmprivacy.api.installations.v1.GetInstallationResponse
	3, // 7: strmprivacy.api.installations.v1.InstallationsService.ListInstallations:output_type -> strmprivacy.api.installations.v1.ListInstallationsResponse
	5, // 8: strmprivacy.api.installations.v1.InstallationsService.GetProjectInstallation:output_type -> strmprivacy.api.installations.v1.GetProjectInstallationResponse
	7, // 9: strmprivacy.api.installations.v1.InstallationsService.ListInstallationProjects:output_type -> strmprivacy.api.installations.v1.ListInstallationProjectsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_installations_v1_installations_v1_proto_init() }
func file_strmprivacy_api_installations_v1_installations_v1_proto_init() {
	if File_strmprivacy_api_installations_v1_installations_v1_proto != nil {
		return
	}
	file_strmprivacy_api_installations_v1_entities_v1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_strmprivacy_api_installations_v1_installations_v1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProjectInstallationRequest); i {
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
			switch v := v.(*GetProjectInstallationResponse); i {
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
			switch v := v.(*ListInstallationProjectsRequest); i {
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
			switch v := v.(*ListInstallationProjectsResponse); i {
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
