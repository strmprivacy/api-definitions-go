// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: strmprivacy/api/categories/v1alpha/categories_v1alpha.proto

package categories

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/strmprivacy/api-definitions-go/v3/api/entities/v1"
	v1alpha "github.com/strmprivacy/api-definitions-go/v3/api/entities/v1alpha"
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

// List all categories for a given type. The organization is derived from the calling user.
type ListCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the categories to list.
	Type v1alpha.Category_Type `protobuf:"varint,1,opt,name=type,proto3,enum=strmprivacy.api.entities.v1alpha.Category_Type" json:"type,omitempty"`
	// The maximum number of categories to return. The server may return fewer than
	// this number of categories, even if there are more categories available.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous `ListCategories` call. Provide this to
	// retrieve the subsequent page.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListCategoriesRequest) Reset() {
	*x = ListCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoriesRequest) ProtoMessage() {}

func (x *ListCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoriesRequest.ProtoReflect.Descriptor instead.
func (*ListCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescGZIP(), []int{0}
}

func (x *ListCategoriesRequest) GetType() v1alpha.Category_Type {
	if x != nil {
		return x.Type
	}
	return v1alpha.Category_Type(0)
}

func (x *ListCategoriesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCategoriesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Categories returned.
	Categories []*v1alpha.Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	// A token to retrieve the next page of results. Pass this value in the
	// `page_token` field in the subsequent call to `ListCategories` method to
	// retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCategoriesResponse) Reset() {
	*x = ListCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoriesResponse) ProtoMessage() {}

func (x *ListCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoriesResponse.ProtoReflect.Descriptor instead.
func (*ListCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescGZIP(), []int{1}
}

func (x *ListCategoriesResponse) GetCategories() []*v1alpha.Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *ListCategoriesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescGZIP(), []int{2}
}

func (x *GetCategoryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *v1alpha.Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetCategoryResponse) Reset() {
	*x = GetCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryResponse) ProtoMessage() {}

func (x *GetCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetCategoryResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescGZIP(), []int{3}
}

func (x *GetCategoryResponse) GetCategory() *v1alpha.Category {
	if x != nil {
		return x.Category
	}
	return nil
}

// To create a category, leave its id empty. To update a category, set its id.
type UpsertCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *v1alpha.Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *UpsertCategoryRequest) Reset() {
	*x = UpsertCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCategoryRequest) ProtoMessage() {}

func (x *UpsertCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpsertCategoryRequest) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertCategoryRequest) GetCategory() *v1alpha.Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type UpsertCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *v1alpha.Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *UpsertCategoryResponse) Reset() {
	*x = UpsertCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCategoryResponse) ProtoMessage() {}

func (x *UpsertCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCategoryResponse.ProtoReflect.Descriptor instead.
func (*UpsertCategoryResponse) Descriptor() ([]byte, []int) {
	return file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertCategoryResponse) GetCategory() *v1alpha.Category {
	if x != nil {
		return x.Category
	}
	return nil
}

var File_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto protoreflect.FileDescriptor

var file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x5f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x2d, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x5f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x73, 0x74, 0x72,
	0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8c, 0x01,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2e, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x5f, 0x0a, 0x15, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x60, 0x0a, 0x16,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0xa7,
	0x03, 0x0a, 0x11, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x39, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7e,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x36, 0x2e,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x87,
	0x01, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x39, 0x2e, 0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x79, 0x0a, 0x25, 0x69, 0x6f, 0x2e, 0x73,
	0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x74, 0x72, 0x6d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescOnce sync.Once
	file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescData = file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDesc
)

func file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescGZIP() []byte {
	file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescOnce.Do(func() {
		file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescData = protoimpl.X.CompressGZIP(file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescData)
	})
	return file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDescData
}

var file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_goTypes = []interface{}{
	(*ListCategoriesRequest)(nil),  // 0: strmprivacy.api.categories.v1alpha.ListCategoriesRequest
	(*ListCategoriesResponse)(nil), // 1: strmprivacy.api.categories.v1alpha.ListCategoriesResponse
	(*GetCategoryRequest)(nil),     // 2: strmprivacy.api.categories.v1alpha.GetCategoryRequest
	(*GetCategoryResponse)(nil),    // 3: strmprivacy.api.categories.v1alpha.GetCategoryResponse
	(*UpsertCategoryRequest)(nil),  // 4: strmprivacy.api.categories.v1alpha.UpsertCategoryRequest
	(*UpsertCategoryResponse)(nil), // 5: strmprivacy.api.categories.v1alpha.UpsertCategoryResponse
	(v1alpha.Category_Type)(0),     // 6: strmprivacy.api.entities.v1alpha.Category.Type
	(*v1alpha.Category)(nil),       // 7: strmprivacy.api.entities.v1alpha.Category
}
var file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_depIdxs = []int32{
	6, // 0: strmprivacy.api.categories.v1alpha.ListCategoriesRequest.type:type_name -> strmprivacy.api.entities.v1alpha.Category.Type
	7, // 1: strmprivacy.api.categories.v1alpha.ListCategoriesResponse.categories:type_name -> strmprivacy.api.entities.v1alpha.Category
	7, // 2: strmprivacy.api.categories.v1alpha.GetCategoryResponse.category:type_name -> strmprivacy.api.entities.v1alpha.Category
	7, // 3: strmprivacy.api.categories.v1alpha.UpsertCategoryRequest.category:type_name -> strmprivacy.api.entities.v1alpha.Category
	7, // 4: strmprivacy.api.categories.v1alpha.UpsertCategoryResponse.category:type_name -> strmprivacy.api.entities.v1alpha.Category
	0, // 5: strmprivacy.api.categories.v1alpha.CategoriesService.ListCategories:input_type -> strmprivacy.api.categories.v1alpha.ListCategoriesRequest
	2, // 6: strmprivacy.api.categories.v1alpha.CategoriesService.GetCategory:input_type -> strmprivacy.api.categories.v1alpha.GetCategoryRequest
	4, // 7: strmprivacy.api.categories.v1alpha.CategoriesService.UpsertCategory:input_type -> strmprivacy.api.categories.v1alpha.UpsertCategoryRequest
	1, // 8: strmprivacy.api.categories.v1alpha.CategoriesService.ListCategories:output_type -> strmprivacy.api.categories.v1alpha.ListCategoriesResponse
	3, // 9: strmprivacy.api.categories.v1alpha.CategoriesService.GetCategory:output_type -> strmprivacy.api.categories.v1alpha.GetCategoryResponse
	5, // 10: strmprivacy.api.categories.v1alpha.CategoriesService.UpsertCategory:output_type -> strmprivacy.api.categories.v1alpha.UpsertCategoryResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_init() }
func file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_init() {
	if File_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoriesRequest); i {
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
		file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoriesResponse); i {
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
		file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryRequest); i {
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
		file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryResponse); i {
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
		file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertCategoryRequest); i {
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
		file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertCategoryResponse); i {
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
			RawDescriptor: file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_goTypes,
		DependencyIndexes: file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_depIdxs,
		MessageInfos:      file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_msgTypes,
	}.Build()
	File_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto = out.File
	file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_rawDesc = nil
	file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_goTypes = nil
	file_strmprivacy_api_categories_v1alpha_categories_v1alpha_proto_depIdxs = nil
}