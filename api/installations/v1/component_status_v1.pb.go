// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: strmprivacy/api/component_status/v1/component_status_v1.proto

package installations

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

type CustomerComponentType int32

const (
	CustomerComponentType_CUSTOMER_COMPONENT_TYPE_UNSPECIFIED CustomerComponentType = 0
	CustomerComponentType_BATCH_JOBS_AGENT                    CustomerComponentType = 1
	CustomerComponentType_BATCH_EXPORTER_AGENT                CustomerComponentType = 2
	CustomerComponentType_KAFKA_EXPORTERS_AGENT               CustomerComponentType = 3
	CustomerComponentType_STREAMS_AGENT                       CustomerComponentType = 4
	CustomerComponentType_EVENT_GATEWAY                       CustomerComponentType = 5
)

// Enum value maps for CustomerComponentType.
var (
	CustomerComponentType_name = map[int32]string{
		0: "CUSTOMER_COMPONENT_TYPE_UNSPECIFIED",
		1: "BATCH_JOBS_AGENT",
		2: "BATCH_EXPORTER_AGENT",
		3: "KAFKA_EXPORTERS_AGENT",
		4: "STREAMS_AGENT",
		5: "EVENT_GATEWAY",
	}
	CustomerComponentType_value = map[string]int32{
		"CUSTOMER_COMPONENT_TYPE_UNSPECIFIED": 0,
		"BATCH_JOBS_AGENT":                    1,
		"BATCH_EXPORTER_AGENT":                2,
		"KAFKA_EXPORTERS_AGENT":               3,
		"STREAMS_AGENT":                       4,
		"EVENT_GATEWAY":                       5,
	}
)

func (x CustomerComponentType) Enum() *CustomerComponentType {
	p := new(CustomerComponentType)
	*p = x
	return p
}

func (x CustomerComponentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomerComponentType) Descriptor() protoreflect.EnumDescriptor {
	return file_strmprivacy_api_component_status_v1_component_status_v1_proto_enumTypes[0].Descriptor()
}

func (CustomerComponentType) Type() protoreflect.EnumType {
	return &file_strmprivacy_api_component_status_v1_component_status_v1_proto_enumTypes[0]
}

func (x CustomerComponentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomerComponentType.Descriptor instead.
func (CustomerComponentType) EnumDescriptor() ([]byte, []int) {
	return file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescGZIP(), []int{0}
}

type ComponentStateType int32

const (
	ComponentStateType_COMPONENT_STATE_TYPE_UNSPECIFIED ComponentStateType = 0
	ComponentStateType_PENDING                          ComponentStateType = 1
	ComponentStateType_STARTED                          ComponentStateType = 2
	ComponentStateType_ERROR_STARTING                   ComponentStateType = 3
	ComponentStateType_RUNNING                          ComponentStateType = 4
	ComponentStateType_FINISHED                         ComponentStateType = 5
	ComponentStateType_ERROR                            ComponentStateType = 6
)

// Enum value maps for ComponentStateType.
var (
	ComponentStateType_name = map[int32]string{
		0: "COMPONENT_STATE_TYPE_UNSPECIFIED",
		1: "PENDING",
		2: "STARTED",
		3: "ERROR_STARTING",
		4: "RUNNING",
		5: "FINISHED",
		6: "ERROR",
	}
	ComponentStateType_value = map[string]int32{
		"COMPONENT_STATE_TYPE_UNSPECIFIED": 0,
		"PENDING":                          1,
		"STARTED":                          2,
		"ERROR_STARTING":                   3,
		"RUNNING":                          4,
		"FINISHED":                         5,
		"ERROR":                            6,
	}
)

func (x ComponentStateType) Enum() *ComponentStateType {
	p := new(ComponentStateType)
	*p = x
	return p
}

func (x ComponentStateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComponentStateType) Descriptor() protoreflect.EnumDescriptor {
	return file_strmprivacy_api_component_status_v1_component_status_v1_proto_enumTypes[1].Descriptor()
}

func (ComponentStateType) Type() protoreflect.EnumType {
	return &file_strmprivacy_api_component_status_v1_component_status_v1_proto_enumTypes[1]
}

func (x ComponentStateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComponentStateType.Descriptor instead.
func (ComponentStateType) EnumDescriptor() ([]byte, []int) {
	return file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescGZIP(), []int{1}
}

// (-- api-linter: core::0134::request-mask-required=disabled
//     aip.dev/not-precedent: This RPC only updates one field. --)
type UpdateComponentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillingId      string          `protobuf:"bytes,1,opt,name=billing_id,json=billingId,proto3" json:"billing_id,omitempty"`
	ComponentState *ComponentState `protobuf:"bytes,2,opt,name=component_state,json=componentState,proto3" json:"component_state,omitempty"`
}

func (x *UpdateComponentStatusRequest) Reset() {
	*x = UpdateComponentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateComponentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateComponentStatusRequest) ProtoMessage() {}

func (x *UpdateComponentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateComponentStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateComponentStatusRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateComponentStatusRequest) GetBillingId() string {
	if x != nil {
		return x.BillingId
	}
	return ""
}

func (x *UpdateComponentStatusRequest) GetComponentState() *ComponentState {
	if x != nil {
		return x.ComponentState
	}
	return nil
}

type UpdateComponentStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateComponentStatusResponse) Reset() {
	*x = UpdateComponentStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateComponentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateComponentStatusResponse) ProtoMessage() {}

func (x *UpdateComponentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateComponentStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateComponentStatusResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescGZIP(), []int{1}
}

type ComponentState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=state_time,json=stateTime,proto3" json:"state_time,omitempty"`
	Component *CustomerComponent     `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
	State     ComponentStateType     `protobuf:"varint,3,opt,name=state,proto3,enum=strmprivacy.api.component_status.v1.ComponentStateType" json:"state,omitempty"`
}

func (x *ComponentState) Reset() {
	*x = ComponentState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentState) ProtoMessage() {}

func (x *ComponentState) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentState.ProtoReflect.Descriptor instead.
func (*ComponentState) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescGZIP(), []int{2}
}

func (x *ComponentState) GetStateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StateTime
	}
	return nil
}

func (x *ComponentState) GetComponent() *CustomerComponent {
	if x != nil {
		return x.Component
	}
	return nil
}

func (x *ComponentState) GetState() ComponentStateType {
	if x != nil {
		return x.State
	}
	return ComponentStateType_COMPONENT_STATE_TYPE_UNSPECIFIED
}

type CustomerComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentType CustomerComponentType `protobuf:"varint,1,opt,name=component_type,json=componentType,proto3,enum=strmprivacy.api.component_status.v1.CustomerComponentType" json:"component_type,omitempty"`
	ComponentId   string                `protobuf:"bytes,2,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty"`
}

func (x *CustomerComponent) Reset() {
	*x = CustomerComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerComponent) ProtoMessage() {}

func (x *CustomerComponent) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerComponent.ProtoReflect.Descriptor instead.
func (*CustomerComponent) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescGZIP(), []int{3}
}

func (x *CustomerComponent) GetComponentType() CustomerComponentType {
	if x != nil {
		return x.ComponentType
	}
	return CustomerComponentType_CUSTOMER_COMPONENT_TYPE_UNSPECIFIED
}

func (x *CustomerComponent) GetComponentId() string {
	if x != nil {
		return x.ComponentId
	}
	return ""
}

var File_strmprivacy_api_component_status_v1_component_status_v1_proto protoreflect.FileDescriptor

var file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x23, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x09, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x61, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x1f,
	0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xff, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x59, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x52, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0xa3, 0x01, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x66, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x3a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x2a, 0xb1, 0x01, 0x0a, 0x15, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x4a, 0x4f, 0x42, 0x53, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x18, 0x0a, 0x14, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54,
	0x45, 0x52, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4b, 0x41,
	0x46, 0x4b, 0x41, 0x5f, 0x45, 0x58, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x52, 0x53, 0x5f, 0x41, 0x47,
	0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x53,
	0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x10, 0x05, 0x2a, 0x8e, 0x01, 0x0a, 0x12,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10,
	0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x32, 0xb9, 0x01, 0x0a,
	0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x41, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x78, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescOnce sync.Once
	file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescData = file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDesc
)

func file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescData)
	})
	return file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDescData
}

var file_strmprivacy_api_component_status_v1_component_status_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_strmprivacy_api_component_status_v1_component_status_v1_proto_goTypes = []interface{}{
	(CustomerComponentType)(0),            // 0: strmprivacy.api.component_status.v1.CustomerComponentType
	(ComponentStateType)(0),               // 1: strmprivacy.api.component_status.v1.ComponentStateType
	(*UpdateComponentStatusRequest)(nil),  // 2: strmprivacy.api.component_status.v1.UpdateComponentStatusRequest
	(*UpdateComponentStatusResponse)(nil), // 3: strmprivacy.api.component_status.v1.UpdateComponentStatusResponse
	(*ComponentState)(nil),                // 4: strmprivacy.api.component_status.v1.ComponentState
	(*CustomerComponent)(nil),             // 5: strmprivacy.api.component_status.v1.CustomerComponent
	(*timestamppb.Timestamp)(nil),         // 6: google.protobuf.Timestamp
}
var file_strmprivacy_api_component_status_v1_component_status_v1_proto_depIdxs = []int32{
	4, // 0: strmprivacy.api.component_status.v1.UpdateComponentStatusRequest.component_state:type_name -> strmprivacy.api.component_status.v1.ComponentState
	6, // 1: strmprivacy.api.component_status.v1.ComponentState.state_time:type_name -> google.protobuf.Timestamp
	5, // 2: strmprivacy.api.component_status.v1.ComponentState.component:type_name -> strmprivacy.api.component_status.v1.CustomerComponent
	1, // 3: strmprivacy.api.component_status.v1.ComponentState.state:type_name -> strmprivacy.api.component_status.v1.ComponentStateType
	0, // 4: strmprivacy.api.component_status.v1.CustomerComponent.component_type:type_name -> strmprivacy.api.component_status.v1.CustomerComponentType
	2, // 5: strmprivacy.api.component_status.v1.ComponentStatusService.UpdateComponentStatus:input_type -> strmprivacy.api.component_status.v1.UpdateComponentStatusRequest
	3, // 6: strmprivacy.api.component_status.v1.ComponentStatusService.UpdateComponentStatus:output_type -> strmprivacy.api.component_status.v1.UpdateComponentStatusResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_component_status_v1_component_status_v1_proto_init() }
func file_strmprivacy_api_component_status_v1_component_status_v1_proto_init() {
	if File_strmprivacy_api_component_status_v1_component_status_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateComponentStatusRequest); i {
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
		file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateComponentStatusResponse); i {
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
		file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentState); i {
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
		file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerComponent); i {
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
			RawDescriptor: file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_component_status_v1_component_status_v1_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_component_status_v1_component_status_v1_proto_depIdxs,
		EnumInfos:         file_strmprivacy_api_component_status_v1_component_status_v1_proto_enumTypes,
		MessageInfos:      file_strmprivacy_api_component_status_v1_component_status_v1_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_component_status_v1_component_status_v1_proto = out.File
	file_strmprivacy_api_component_status_v1_component_status_v1_proto_rawDesc = nil
	file_strmprivacy_api_component_status_v1_component_status_v1_proto_goTypes = nil
	file_strmprivacy_api_component_status_v1_component_status_v1_proto_depIdxs = nil
}