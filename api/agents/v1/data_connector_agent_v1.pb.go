// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: strmprivacy/api/agents/v1/data_connector_agent_v1.proto

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

type GetDesiredDataConnectorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDesiredDataConnectorsRequest) Reset() {
	*x = GetDesiredDataConnectorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDesiredDataConnectorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDesiredDataConnectorsRequest) ProtoMessage() {}

func (x *GetDesiredDataConnectorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDesiredDataConnectorsRequest.ProtoReflect.Descriptor instead.
func (*GetDesiredDataConnectorsRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDescGZIP(), []int{0}
}

type GetDesiredDataConnectorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataConnectors []*v1.DataConnector `protobuf:"bytes,1,rep,name=data_connectors,json=dataConnectors,proto3" json:"data_connectors,omitempty"`
}

func (x *GetDesiredDataConnectorsResponse) Reset() {
	*x = GetDesiredDataConnectorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDesiredDataConnectorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDesiredDataConnectorsResponse) ProtoMessage() {}

func (x *GetDesiredDataConnectorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDesiredDataConnectorsResponse.ProtoReflect.Descriptor instead.
func (*GetDesiredDataConnectorsResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GetDesiredDataConnectorsResponse) GetDataConnectors() []*v1.DataConnector {
	if x != nil {
		return x.DataConnectors
	}
	return nil
}

var File_strmprivacy_api_agents_v1_data_connector_agent_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDesc = []byte{
	0x0a, 0x37, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x77, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x32,
	0xb2, 0x01, 0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x93,
	0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x3a, 0x2e, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x72,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x63, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67,
	0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDescData = file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDesc
)

func file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDescData
}

var file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_goTypes = []interface{}{
	(*GetDesiredDataConnectorsRequest)(nil),  // 0: strmprivacy.api.agents.v1.GetDesiredDataConnectorsRequest
	(*GetDesiredDataConnectorsResponse)(nil), // 1: strmprivacy.api.agents.v1.GetDesiredDataConnectorsResponse
	(*v1.DataConnector)(nil),                 // 2: strmprivacy.api.entities.v1.DataConnector
}
var file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_depIdxs = []int32{
	2, // 0: strmprivacy.api.agents.v1.GetDesiredDataConnectorsResponse.data_connectors:type_name -> strmprivacy.api.entities.v1.DataConnector
	0, // 1: strmprivacy.api.agents.v1.DataConnectorsAgentService.GetDesiredDataConnectors:input_type -> strmprivacy.api.agents.v1.GetDesiredDataConnectorsRequest
	1, // 2: strmprivacy.api.agents.v1.DataConnectorsAgentService.GetDesiredDataConnectors:output_type -> strmprivacy.api.agents.v1.GetDesiredDataConnectorsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_init() }
func file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_init() {
	if File_strmprivacy_api_agents_v1_data_connector_agent_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDesiredDataConnectorsRequest); i {
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
		file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDesiredDataConnectorsResponse); i {
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
			RawDescriptor: file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_agents_v1_data_connector_agent_v1_proto = out.File
	file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_rawDesc = nil
	file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_goTypes = nil
	file_strmprivacy_api_agents_v1_data_connector_agent_v1_proto_depIdxs = nil
}