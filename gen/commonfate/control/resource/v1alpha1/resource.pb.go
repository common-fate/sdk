// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: commonfate/control/resource/v1alpha1/resource.proto

package resourcev1alpha1

import (
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	_ "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	v1alpha11 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
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

type ListTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTypesRequest) Reset() {
	*x = ListTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTypesRequest) ProtoMessage() {}

func (x *ListTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTypesRequest.ProtoReflect.Descriptor instead.
func (*ListTypesRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{0}
}

type ListTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []string `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *ListTypesResponse) Reset() {
	*x = ListTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTypesResponse) ProtoMessage() {}

func (x *ListTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTypesResponse.ProtoReflect.Descriptor instead.
func (*ListTypesResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{1}
}

func (x *ListTypesResponse) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

type QueryResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *QueryResourcesRequest) Reset() {
	*x = QueryResourcesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResourcesRequest) ProtoMessage() {}

func (x *QueryResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResourcesRequest.ProtoReflect.Descriptor instead.
func (*QueryResourcesRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{2}
}

func (x *QueryResourcesRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type QueryResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*v1alpha1.NamedEID `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *QueryResourcesResponse) Reset() {
	*x = QueryResourcesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResourcesResponse) ProtoMessage() {}

func (x *QueryResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResourcesResponse.ProtoReflect.Descriptor instead.
func (*QueryResourcesResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{3}
}

func (x *QueryResourcesResponse) GetResources() []*v1alpha1.NamedEID {
	if x != nil {
		return x.Resources
	}
	return nil
}

type GetResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid *v1alpha11.EID `protobuf:"bytes,1,opt,name=eid,proto3" json:"eid,omitempty"`
}

func (x *GetResourceRequest) Reset() {
	*x = GetResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceRequest) ProtoMessage() {}

func (x *GetResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceRequest.ProtoReflect.Descriptor instead.
func (*GetResourceRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{4}
}

func (x *GetResourceRequest) GetEid() *v1alpha11.EID {
	if x != nil {
		return x.Eid
	}
	return nil
}

type GetResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *v1alpha11.Entity `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Name     string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetResourceResponse) Reset() {
	*x = GetResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceResponse) ProtoMessage() {}

func (x *GetResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceResponse.ProtoReflect.Descriptor instead.
func (*GetResourceResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{5}
}

func (x *GetResourceResponse) GetResource() *v1alpha11.Entity {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *GetResourceResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListChildrenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid *v1alpha11.EID `protobuf:"bytes,1,opt,name=eid,proto3" json:"eid,omitempty"`
}

func (x *ListChildrenRequest) Reset() {
	*x = ListChildrenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChildrenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChildrenRequest) ProtoMessage() {}

func (x *ListChildrenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChildrenRequest.ProtoReflect.Descriptor instead.
func (*ListChildrenRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{6}
}

func (x *ListChildrenRequest) GetEid() *v1alpha11.EID {
	if x != nil {
		return x.Eid
	}
	return nil
}

type ListChildrenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Children []*v1alpha1.NamedEID `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ListChildrenResponse) Reset() {
	*x = ListChildrenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChildrenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChildrenResponse) ProtoMessage() {}

func (x *ListChildrenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChildrenResponse.ProtoReflect.Descriptor instead.
func (*ListChildrenResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{7}
}

func (x *ListChildrenResponse) GetChildren() []*v1alpha1.NamedEID {
	if x != nil {
		return x.Children
	}
	return nil
}

type ListParentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid *v1alpha11.EID `protobuf:"bytes,1,opt,name=eid,proto3" json:"eid,omitempty"`
}

func (x *ListParentsRequest) Reset() {
	*x = ListParentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListParentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListParentsRequest) ProtoMessage() {}

func (x *ListParentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListParentsRequest.ProtoReflect.Descriptor instead.
func (*ListParentsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{8}
}

func (x *ListParentsRequest) GetEid() *v1alpha11.EID {
	if x != nil {
		return x.Eid
	}
	return nil
}

type ListParentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parents []*v1alpha1.NamedEID `protobuf:"bytes,1,rep,name=parents,proto3" json:"parents,omitempty"`
}

func (x *ListParentsResponse) Reset() {
	*x = ListParentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListParentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListParentsResponse) ProtoMessage() {}

func (x *ListParentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListParentsResponse.ProtoReflect.Descriptor instead.
func (*ListParentsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP(), []int{9}
}

func (x *ListParentsResponse) GetParents() []*v1alpha1.NamedEID {
	if x != nil {
		return x.Parents
	}
	return nil
}

var File_commonfate_control_resource_v1alpha1_resource_proto protoreflect.FileDescriptor

var file_commonfate_control_resource_v1alpha1_resource_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x65, 0x69,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65,
	0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x2b, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5c, 0x0a,
	0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x45, 0x49, 0x44,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x03, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52,
	0x03, 0x65, 0x69, 0x64, 0x22, 0x69, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x48, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x03, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x49, 0x44, 0x52, 0x03, 0x65, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x45, 0x49, 0x44, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x72, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x03, 0x65, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x03, 0x65, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x45, 0x49, 0x44, 0x52, 0x07, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0xce, 0x05, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x91, 0x01, 0x0a,
	0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12,
	0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01,
	0x12, 0x88, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x8b, 0x01, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x39, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x88, 0x01, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04,
	0x88, 0xb5, 0x18, 0x01, 0x42, 0xc2, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x52, 0xaa, 0x02,
	0x24, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x24, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x30, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_commonfate_control_resource_v1alpha1_resource_proto_rawDescOnce sync.Once
	file_commonfate_control_resource_v1alpha1_resource_proto_rawDescData = file_commonfate_control_resource_v1alpha1_resource_proto_rawDesc
)

func file_commonfate_control_resource_v1alpha1_resource_proto_rawDescGZIP() []byte {
	file_commonfate_control_resource_v1alpha1_resource_proto_rawDescOnce.Do(func() {
		file_commonfate_control_resource_v1alpha1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_resource_v1alpha1_resource_proto_rawDescData)
	})
	return file_commonfate_control_resource_v1alpha1_resource_proto_rawDescData
}

var file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_commonfate_control_resource_v1alpha1_resource_proto_goTypes = []interface{}{
	(*ListTypesRequest)(nil),       // 0: commonfate.control.resource.v1alpha1.ListTypesRequest
	(*ListTypesResponse)(nil),      // 1: commonfate.control.resource.v1alpha1.ListTypesResponse
	(*QueryResourcesRequest)(nil),  // 2: commonfate.control.resource.v1alpha1.QueryResourcesRequest
	(*QueryResourcesResponse)(nil), // 3: commonfate.control.resource.v1alpha1.QueryResourcesResponse
	(*GetResourceRequest)(nil),     // 4: commonfate.control.resource.v1alpha1.GetResourceRequest
	(*GetResourceResponse)(nil),    // 5: commonfate.control.resource.v1alpha1.GetResourceResponse
	(*ListChildrenRequest)(nil),    // 6: commonfate.control.resource.v1alpha1.ListChildrenRequest
	(*ListChildrenResponse)(nil),   // 7: commonfate.control.resource.v1alpha1.ListChildrenResponse
	(*ListParentsRequest)(nil),     // 8: commonfate.control.resource.v1alpha1.ListParentsRequest
	(*ListParentsResponse)(nil),    // 9: commonfate.control.resource.v1alpha1.ListParentsResponse
	(*v1alpha1.NamedEID)(nil),      // 10: commonfate.access.v1alpha1.NamedEID
	(*v1alpha11.EID)(nil),          // 11: commonfate.entity.v1alpha1.EID
	(*v1alpha11.Entity)(nil),       // 12: commonfate.entity.v1alpha1.Entity
}
var file_commonfate_control_resource_v1alpha1_resource_proto_depIdxs = []int32{
	10, // 0: commonfate.control.resource.v1alpha1.QueryResourcesResponse.resources:type_name -> commonfate.access.v1alpha1.NamedEID
	11, // 1: commonfate.control.resource.v1alpha1.GetResourceRequest.eid:type_name -> commonfate.entity.v1alpha1.EID
	12, // 2: commonfate.control.resource.v1alpha1.GetResourceResponse.resource:type_name -> commonfate.entity.v1alpha1.Entity
	11, // 3: commonfate.control.resource.v1alpha1.ListChildrenRequest.eid:type_name -> commonfate.entity.v1alpha1.EID
	10, // 4: commonfate.control.resource.v1alpha1.ListChildrenResponse.children:type_name -> commonfate.access.v1alpha1.NamedEID
	11, // 5: commonfate.control.resource.v1alpha1.ListParentsRequest.eid:type_name -> commonfate.entity.v1alpha1.EID
	10, // 6: commonfate.control.resource.v1alpha1.ListParentsResponse.parents:type_name -> commonfate.access.v1alpha1.NamedEID
	0,  // 7: commonfate.control.resource.v1alpha1.ResourceService.ListTypes:input_type -> commonfate.control.resource.v1alpha1.ListTypesRequest
	2,  // 8: commonfate.control.resource.v1alpha1.ResourceService.QueryResources:input_type -> commonfate.control.resource.v1alpha1.QueryResourcesRequest
	4,  // 9: commonfate.control.resource.v1alpha1.ResourceService.GetResource:input_type -> commonfate.control.resource.v1alpha1.GetResourceRequest
	6,  // 10: commonfate.control.resource.v1alpha1.ResourceService.ListChildren:input_type -> commonfate.control.resource.v1alpha1.ListChildrenRequest
	8,  // 11: commonfate.control.resource.v1alpha1.ResourceService.ListParents:input_type -> commonfate.control.resource.v1alpha1.ListParentsRequest
	1,  // 12: commonfate.control.resource.v1alpha1.ResourceService.ListTypes:output_type -> commonfate.control.resource.v1alpha1.ListTypesResponse
	3,  // 13: commonfate.control.resource.v1alpha1.ResourceService.QueryResources:output_type -> commonfate.control.resource.v1alpha1.QueryResourcesResponse
	5,  // 14: commonfate.control.resource.v1alpha1.ResourceService.GetResource:output_type -> commonfate.control.resource.v1alpha1.GetResourceResponse
	7,  // 15: commonfate.control.resource.v1alpha1.ResourceService.ListChildren:output_type -> commonfate.control.resource.v1alpha1.ListChildrenResponse
	9,  // 16: commonfate.control.resource.v1alpha1.ResourceService.ListParents:output_type -> commonfate.control.resource.v1alpha1.ListParentsResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_commonfate_control_resource_v1alpha1_resource_proto_init() }
func file_commonfate_control_resource_v1alpha1_resource_proto_init() {
	if File_commonfate_control_resource_v1alpha1_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTypesRequest); i {
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
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTypesResponse); i {
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
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResourcesRequest); i {
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
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResourcesResponse); i {
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
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceRequest); i {
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
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceResponse); i {
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
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChildrenRequest); i {
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
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChildrenResponse); i {
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
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListParentsRequest); i {
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
		file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListParentsResponse); i {
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
			RawDescriptor: file_commonfate_control_resource_v1alpha1_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_resource_v1alpha1_resource_proto_goTypes,
		DependencyIndexes: file_commonfate_control_resource_v1alpha1_resource_proto_depIdxs,
		MessageInfos:      file_commonfate_control_resource_v1alpha1_resource_proto_msgTypes,
	}.Build()
	File_commonfate_control_resource_v1alpha1_resource_proto = out.File
	file_commonfate_control_resource_v1alpha1_resource_proto_rawDesc = nil
	file_commonfate_control_resource_v1alpha1_resource_proto_goTypes = nil
	file_commonfate_control_resource_v1alpha1_resource_proto_depIdxs = nil
}
