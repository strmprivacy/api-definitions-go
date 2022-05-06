// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: strmprivacy/api/handles/v1/handles_v1.proto

package handles

import (
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

type ListInstallationHandlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationId string `protobuf:"bytes,1,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
}

func (x *ListInstallationHandlesRequest) Reset() {
	*x = ListInstallationHandlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_handles_v1_handles_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstallationHandlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstallationHandlesRequest) ProtoMessage() {}

func (x *ListInstallationHandlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_handles_v1_handles_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstallationHandlesRequest.ProtoReflect.Descriptor instead.
func (*ListInstallationHandlesRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_handles_v1_handles_v1_proto_rawDescGZIP(), []int{0}
}

func (x *ListInstallationHandlesRequest) GetInstallationId() string {
	if x != nil {
		return x.InstallationId
	}
	return ""
}

type ListInstallationHandlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallationHandles []string `protobuf:"bytes,1,rep,name=installation_handles,json=installationHandles,proto3" json:"installation_handles,omitempty"`
}

func (x *ListInstallationHandlesResponse) Reset() {
	*x = ListInstallationHandlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_handles_v1_handles_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInstallationHandlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstallationHandlesResponse) ProtoMessage() {}

func (x *ListInstallationHandlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_handles_v1_handles_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstallationHandlesResponse.ProtoReflect.Descriptor instead.
func (*ListInstallationHandlesResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_handles_v1_handles_v1_proto_rawDescGZIP(), []int{1}
}

func (x *ListInstallationHandlesResponse) GetInstallationHandles() []string {
	if x != nil {
		return x.InstallationHandles
	}
	return nil
}

var File_strmprivacy_api_handles_v1_handles_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_handles_v1_handles_v1_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x49, 0x0a, 0x1e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x0e, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x92, 0x01,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x3a, 0x2e, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x66, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_strmprivacy_api_handles_v1_handles_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_handles_v1_handles_v1_proto_rawDescData = file_strmprivacy_api_handles_v1_handles_v1_proto_rawDesc
)

func file_strmprivacy_api_handles_v1_handles_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_handles_v1_handles_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_handles_v1_handles_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_handles_v1_handles_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_handles_v1_handles_v1_proto_rawDescData
}

var file_strmprivacy_api_handles_v1_handles_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_strmprivacy_api_handles_v1_handles_v1_proto_goTypes = []interface{}{
	(*ListInstallationHandlesRequest)(nil),  // 0: strmprivacy.api.handles.v1.ListInstallationHandlesRequest
	(*ListInstallationHandlesResponse)(nil), // 1: strmprivacy.api.handles.v1.ListInstallationHandlesResponse
}
var file_strmprivacy_api_handles_v1_handles_v1_proto_depIdxs = []int32{
	0, // 0: strmprivacy.api.handles.v1.HandlesService.ListInstallationHandles:input_type -> strmprivacy.api.handles.v1.ListInstallationHandlesRequest
	1, // 1: strmprivacy.api.handles.v1.HandlesService.ListInstallationHandles:output_type -> strmprivacy.api.handles.v1.ListInstallationHandlesResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_handles_v1_handles_v1_proto_init() }
func file_strmprivacy_api_handles_v1_handles_v1_proto_init() {
	if File_strmprivacy_api_handles_v1_handles_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_handles_v1_handles_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstallationHandlesRequest); i {
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
		file_strmprivacy_api_handles_v1_handles_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInstallationHandlesResponse); i {
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
			RawDescriptor: file_strmprivacy_api_handles_v1_handles_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_handles_v1_handles_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_handles_v1_handles_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_handles_v1_handles_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_handles_v1_handles_v1_proto = out.File
	file_strmprivacy_api_handles_v1_handles_v1_proto_rawDesc = nil
	file_strmprivacy_api_handles_v1_handles_v1_proto_goTypes = nil
	file_strmprivacy_api_handles_v1_handles_v1_proto_depIdxs = nil
}