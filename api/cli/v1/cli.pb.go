// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.8
// source: strmprivacy/api/cli/v1/cli.proto

package cli

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The CLI version, get the latest version if unspecified.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetReleaseRequest) Reset() {
	*x = GetReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_cli_v1_cli_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReleaseRequest) ProtoMessage() {}

func (x *GetReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_cli_v1_cli_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReleaseRequest.ProtoReflect.Descriptor instead.
func (*GetReleaseRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_cli_v1_cli_proto_rawDescGZIP(), []int{0}
}

func (x *GetReleaseRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type GetReleaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Release *CliRelease `protobuf:"bytes,1,opt,name=release,proto3" json:"release,omitempty"`
}

func (x *GetReleaseResponse) Reset() {
	*x = GetReleaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_cli_v1_cli_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReleaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReleaseResponse) ProtoMessage() {}

func (x *GetReleaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_cli_v1_cli_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReleaseResponse.ProtoReflect.Descriptor instead.
func (*GetReleaseResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_cli_v1_cli_proto_rawDescGZIP(), []int{1}
}

func (x *GetReleaseResponse) GetRelease() *CliRelease {
	if x != nil {
		return x.Release
	}
	return nil
}

type CliRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version string for this CLI release.
	Version     string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	ReleaseTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=release_time,json=releaseTime,proto3" json:"release_time,omitempty"`
	// The absolute URL to the release page in GitHub.
	SourceUri string `protobuf:"bytes,3,opt,name=source_uri,json=sourceUri,proto3" json:"source_uri,omitempty"`
	// The release notes in Markdown format.
	ReleaseNotes string `protobuf:"bytes,4,opt,name=release_notes,json=releaseNotes,proto3" json:"release_notes,omitempty"`
}

func (x *CliRelease) Reset() {
	*x = CliRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_cli_v1_cli_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CliRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CliRelease) ProtoMessage() {}

func (x *CliRelease) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_cli_v1_cli_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CliRelease.ProtoReflect.Descriptor instead.
func (*CliRelease) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_cli_v1_cli_proto_rawDescGZIP(), []int{2}
}

func (x *CliRelease) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CliRelease) GetReleaseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseTime
	}
	return nil
}

func (x *CliRelease) GetSourceUri() string {
	if x != nil {
		return x.SourceUri
	}
	return ""
}

func (x *CliRelease) GetReleaseNotes() string {
	if x != nil {
		return x.ReleaseNotes
	}
	return ""
}

var File_strmprivacy_api_cli_v1_cli_proto protoreflect.FileDescriptor

var file_strmprivacy_api_cli_v1_cli_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x69, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x0a, 0x43, 0x6c,
	0x69, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x69, 0x12,
	0x28, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x32, 0x71, 0x0a, 0x0a, 0x43, 0x6c, 0x69,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x29, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x57, 0x0a, 0x19,
	0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x6c, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_cli_v1_cli_proto_rawDescOnce sync.Once
	file_strmprivacy_api_cli_v1_cli_proto_rawDescData = file_strmprivacy_api_cli_v1_cli_proto_rawDesc
)

func file_strmprivacy_api_cli_v1_cli_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_cli_v1_cli_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_cli_v1_cli_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_cli_v1_cli_proto_rawDescData)
	})
	return file_strmprivacy_api_cli_v1_cli_proto_rawDescData
}

var file_strmprivacy_api_cli_v1_cli_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_strmprivacy_api_cli_v1_cli_proto_goTypes = []interface{}{
	(*GetReleaseRequest)(nil),     // 0: strmprivacy.api.cli.v1.GetReleaseRequest
	(*GetReleaseResponse)(nil),    // 1: strmprivacy.api.cli.v1.GetReleaseResponse
	(*CliRelease)(nil),            // 2: strmprivacy.api.cli.v1.CliRelease
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_strmprivacy_api_cli_v1_cli_proto_depIdxs = []int32{
	2, // 0: strmprivacy.api.cli.v1.GetReleaseResponse.release:type_name -> strmprivacy.api.cli.v1.CliRelease
	3, // 1: strmprivacy.api.cli.v1.CliRelease.release_time:type_name -> google.protobuf.Timestamp
	0, // 2: strmprivacy.api.cli.v1.CliService.GetRelease:input_type -> strmprivacy.api.cli.v1.GetReleaseRequest
	1, // 3: strmprivacy.api.cli.v1.CliService.GetRelease:output_type -> strmprivacy.api.cli.v1.GetReleaseResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_cli_v1_cli_proto_init() }
func file_strmprivacy_api_cli_v1_cli_proto_init() {
	if File_strmprivacy_api_cli_v1_cli_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_cli_v1_cli_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReleaseRequest); i {
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
		file_strmprivacy_api_cli_v1_cli_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReleaseResponse); i {
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
		file_strmprivacy_api_cli_v1_cli_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CliRelease); i {
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
			RawDescriptor: file_strmprivacy_api_cli_v1_cli_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_cli_v1_cli_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_cli_v1_cli_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_cli_v1_cli_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_cli_v1_cli_proto = out.File
	file_strmprivacy_api_cli_v1_cli_proto_rawDesc = nil
	file_strmprivacy_api_cli_v1_cli_proto_goTypes = nil
	file_strmprivacy_api_cli_v1_cli_proto_depIdxs = nil
}
