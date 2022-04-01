// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: strmprivacy/api/agents/v1/streams_agent_v1.proto

package agents

import (
	v1 "github.com/strmprivacy/api-definitions-go/v2/api/entities/v1"
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

type GetStreamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStreamsRequest) Reset() {
	*x = GetStreamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_agents_v1_streams_agent_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamsRequest) ProtoMessage() {}

func (x *GetStreamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_agents_v1_streams_agent_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamsRequest.ProtoReflect.Descriptor instead.
func (*GetStreamsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDescGZIP(), []int{0}
}

type GetStreamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtendedStreams []*v1.ExtendedStream `protobuf:"bytes,1,rep,name=extended_streams,json=extendedStreams,proto3" json:"extended_streams,omitempty"`
}

func (x *GetStreamsResponse) Reset() {
	*x = GetStreamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_agents_v1_streams_agent_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamsResponse) ProtoMessage() {}

func (x *GetStreamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_agents_v1_streams_agent_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamsResponse.ProtoReflect.Descriptor instead.
func (*GetStreamsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetStreamsResponse) GetExtendedStreams() []*v1.ExtendedStream {
	if x != nil {
		return x.ExtendedStreams
	}
	return nil
}

var File_strmprivacy_api_agents_v1_streams_agent_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2d, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x6c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x0f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x32,
	0x80, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x2c, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x63, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69,
	0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDescData = file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDesc
)

func file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDescData
}

var file_strmprivacy_api_agents_v1_streams_agent_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_strmprivacy_api_agents_v1_streams_agent_v1_proto_goTypes = []interface{}{
	(*GetStreamsRequest)(nil),  // 0: strmprivacy.api.agents.v1.GetStreamsRequest
	(*GetStreamsResponse)(nil), // 1: strmprivacy.api.agents.v1.GetStreamsResponse
	(*v1.ExtendedStream)(nil),  // 2: strmprivacy.api.entities.v1.ExtendedStream
}
var file_strmprivacy_api_agents_v1_streams_agent_v1_proto_depIdxs = []int32{
	2, // 0: strmprivacy.api.agents.v1.GetStreamsResponse.extended_streams:type_name -> strmprivacy.api.entities.v1.ExtendedStream
	0, // 1: strmprivacy.api.agents.v1.StreamsAgentService.GetStreams:input_type -> strmprivacy.api.agents.v1.GetStreamsRequest
	1, // 2: strmprivacy.api.agents.v1.StreamsAgentService.GetStreams:output_type -> strmprivacy.api.agents.v1.GetStreamsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_agents_v1_streams_agent_v1_proto_init() }
func file_strmprivacy_api_agents_v1_streams_agent_v1_proto_init() {
	if File_strmprivacy_api_agents_v1_streams_agent_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_agents_v1_streams_agent_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamsRequest); i {
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
		file_strmprivacy_api_agents_v1_streams_agent_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamsResponse); i {
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
			RawDescriptor: file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_agents_v1_streams_agent_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_agents_v1_streams_agent_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_agents_v1_streams_agent_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_agents_v1_streams_agent_v1_proto = out.File
	file_strmprivacy_api_agents_v1_streams_agent_v1_proto_rawDesc = nil
	file_strmprivacy_api_agents_v1_streams_agent_v1_proto_goTypes = nil
	file_strmprivacy_api_agents_v1_streams_agent_v1_proto_depIdxs = nil
}