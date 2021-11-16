// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.8
// source: strmprivacy/api/metrics/v1/metrics.proto

package metrics

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

type SystemState int32

const (
	SystemState_SYSTEM_STATE_UNSPECIFIED SystemState = 0
	// Everything is working fine (green).
	SystemState_UP SystemState = 1
	// One or more major parts of the system are down/unusable (red).
	SystemState_MAJOR_DOWN SystemState = 2
	// Small parts of the system are down/unusable, but the core is not affected. (orange).
	SystemState_MINOR_DOWN SystemState = 3
)

// Enum value maps for SystemState.
var (
	SystemState_name = map[int32]string{
		0: "SYSTEM_STATE_UNSPECIFIED",
		1: "UP",
		2: "MAJOR_DOWN",
		3: "MINOR_DOWN",
	}
	SystemState_value = map[string]int32{
		"SYSTEM_STATE_UNSPECIFIED": 0,
		"UP":                       1,
		"MAJOR_DOWN":               2,
		"MINOR_DOWN":               3,
	}
)

func (x SystemState) Enum() *SystemState {
	p := new(SystemState)
	*p = x
	return p
}

func (x SystemState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SystemState) Descriptor() protoreflect.EnumDescriptor {
	return file_strmprivacy_api_metrics_v1_metrics_proto_enumTypes[0].Descriptor()
}

func (SystemState) Type() protoreflect.EnumType {
	return &file_strmprivacy_api_metrics_v1_metrics_proto_enumTypes[0]
}

func (x SystemState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SystemState.Descriptor instead.
func (SystemState) EnumDescriptor() ([]byte, []int) {
	return file_strmprivacy_api_metrics_v1_metrics_proto_rawDescGZIP(), []int{0}
}

type GetSystemStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSystemStateRequest) Reset() {
	*x = GetSystemStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_metrics_v1_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemStateRequest) ProtoMessage() {}

func (x *GetSystemStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_metrics_v1_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemStateRequest.ProtoReflect.Descriptor instead.
func (*GetSystemStateRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_metrics_v1_metrics_proto_rawDescGZIP(), []int{0}
}

type GetSystemStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OverallState SystemState `protobuf:"varint,1,opt,name=overall_state,json=overallState,proto3,enum=strmprivacy.api.metrics.v1.SystemState" json:"overall_state,omitempty"`
}

func (x *GetSystemStateResponse) Reset() {
	*x = GetSystemStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_metrics_v1_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemStateResponse) ProtoMessage() {}

func (x *GetSystemStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_metrics_v1_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemStateResponse.ProtoReflect.Descriptor instead.
func (*GetSystemStateResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_metrics_v1_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *GetSystemStateResponse) GetOverallState() SystemState {
	if x != nil {
		return x.OverallState
	}
	return SystemState_SYSTEM_STATE_UNSPECIFIED
}

var File_strmprivacy_api_metrics_v1_metrics_proto protoreflect.FileDescriptor

var file_strmprivacy_api_metrics_v1_metrics_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x74, 0x72, 0x6d,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x6b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0d, 0x6f, 0x76,
	0x65, 0x72, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0c, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x53, 0x0a,
	0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x18,
	0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x41, 0x4a, 0x4f, 0x52, 0x5f, 0x44, 0x4f, 0x57, 0x4e,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x49, 0x4e, 0x4f, 0x52, 0x5f, 0x44, 0x4f, 0x57, 0x4e,
	0x10, 0x03, 0x32, 0x89, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x63,
	0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_metrics_v1_metrics_proto_rawDescOnce sync.Once
	file_strmprivacy_api_metrics_v1_metrics_proto_rawDescData = file_strmprivacy_api_metrics_v1_metrics_proto_rawDesc
)

func file_strmprivacy_api_metrics_v1_metrics_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_metrics_v1_metrics_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_metrics_v1_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_metrics_v1_metrics_proto_rawDescData)
	})
	return file_strmprivacy_api_metrics_v1_metrics_proto_rawDescData
}

var file_strmprivacy_api_metrics_v1_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_strmprivacy_api_metrics_v1_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_strmprivacy_api_metrics_v1_metrics_proto_goTypes = []interface{}{
	(SystemState)(0),               // 0: strmprivacy.api.metrics.v1.SystemState
	(*GetSystemStateRequest)(nil),  // 1: strmprivacy.api.metrics.v1.GetSystemStateRequest
	(*GetSystemStateResponse)(nil), // 2: strmprivacy.api.metrics.v1.GetSystemStateResponse
}
var file_strmprivacy_api_metrics_v1_metrics_proto_depIdxs = []int32{
	0, // 0: strmprivacy.api.metrics.v1.GetSystemStateResponse.overall_state:type_name -> strmprivacy.api.metrics.v1.SystemState
	1, // 1: strmprivacy.api.metrics.v1.MetricsService.GetSystemState:input_type -> strmprivacy.api.metrics.v1.GetSystemStateRequest
	2, // 2: strmprivacy.api.metrics.v1.MetricsService.GetSystemState:output_type -> strmprivacy.api.metrics.v1.GetSystemStateResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_metrics_v1_metrics_proto_init() }
func file_strmprivacy_api_metrics_v1_metrics_proto_init() {
	if File_strmprivacy_api_metrics_v1_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_metrics_v1_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemStateRequest); i {
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
		file_strmprivacy_api_metrics_v1_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemStateResponse); i {
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
			RawDescriptor: file_strmprivacy_api_metrics_v1_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_metrics_v1_metrics_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_metrics_v1_metrics_proto_depIdxs,
		EnumInfos:         file_strmprivacy_api_metrics_v1_metrics_proto_enumTypes,
		MessageInfos:      file_strmprivacy_api_metrics_v1_metrics_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_metrics_v1_metrics_proto = out.File
	file_strmprivacy_api_metrics_v1_metrics_proto_rawDesc = nil
	file_strmprivacy_api_metrics_v1_metrics_proto_goTypes = nil
	file_strmprivacy_api_metrics_v1_metrics_proto_depIdxs = nil
}
