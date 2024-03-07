// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: commonfate/control/oauth/v1alpha1/oauth.proto

package oauthv1alpha1

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

type CreatePagerDutyIntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePagerDutyIntegrationRequest) Reset() {
	*x = CreatePagerDutyIntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePagerDutyIntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePagerDutyIntegrationRequest) ProtoMessage() {}

func (x *CreatePagerDutyIntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePagerDutyIntegrationRequest.ProtoReflect.Descriptor instead.
func (*CreatePagerDutyIntegrationRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{0}
}

type CreatePagerDutyIntegrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *CreatePagerDutyIntegrationResponse) Reset() {
	*x = CreatePagerDutyIntegrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePagerDutyIntegrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePagerDutyIntegrationResponse) ProtoMessage() {}

func (x *CreatePagerDutyIntegrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePagerDutyIntegrationResponse.ProtoReflect.Descriptor instead.
func (*CreatePagerDutyIntegrationResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePagerDutyIntegrationResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type GetPagerDutyIntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPagerDutyIntegrationRequest) Reset() {
	*x = GetPagerDutyIntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPagerDutyIntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPagerDutyIntegrationRequest) ProtoMessage() {}

func (x *GetPagerDutyIntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPagerDutyIntegrationRequest.ProtoReflect.Descriptor instead.
func (*GetPagerDutyIntegrationRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{2}
}

type GetPagerDutyIntegrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *GetPagerDutyIntegrationResponse) Reset() {
	*x = GetPagerDutyIntegrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPagerDutyIntegrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPagerDutyIntegrationResponse) ProtoMessage() {}

func (x *GetPagerDutyIntegrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPagerDutyIntegrationResponse.ProtoReflect.Descriptor instead.
func (*GetPagerDutyIntegrationResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{3}
}

func (x *GetPagerDutyIntegrationResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type RemovePagerDutyIntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemovePagerDutyIntegrationRequest) Reset() {
	*x = RemovePagerDutyIntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePagerDutyIntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePagerDutyIntegrationRequest) ProtoMessage() {}

func (x *RemovePagerDutyIntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePagerDutyIntegrationRequest.ProtoReflect.Descriptor instead.
func (*RemovePagerDutyIntegrationRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{4}
}

type RemovePagerDutyIntegrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemovePagerDutyIntegrationResponse) Reset() {
	*x = RemovePagerDutyIntegrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePagerDutyIntegrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePagerDutyIntegrationResponse) ProtoMessage() {}

func (x *RemovePagerDutyIntegrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePagerDutyIntegrationResponse.ProtoReflect.Descriptor instead.
func (*RemovePagerDutyIntegrationResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{5}
}

func (x *RemovePagerDutyIntegrationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CreateSlackIntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSlackIntegrationRequest) Reset() {
	*x = CreateSlackIntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlackIntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlackIntegrationRequest) ProtoMessage() {}

func (x *CreateSlackIntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlackIntegrationRequest.ProtoReflect.Descriptor instead.
func (*CreateSlackIntegrationRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{6}
}

type CreateSlackIntegrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *CreateSlackIntegrationResponse) Reset() {
	*x = CreateSlackIntegrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlackIntegrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlackIntegrationResponse) ProtoMessage() {}

func (x *CreateSlackIntegrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlackIntegrationResponse.ProtoReflect.Descriptor instead.
func (*CreateSlackIntegrationResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{7}
}

func (x *CreateSlackIntegrationResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type GetSlackIntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSlackIntegrationRequest) Reset() {
	*x = GetSlackIntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSlackIntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSlackIntegrationRequest) ProtoMessage() {}

func (x *GetSlackIntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSlackIntegrationRequest.ProtoReflect.Descriptor instead.
func (*GetSlackIntegrationRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{8}
}

type GetSlackIntegrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active bool `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *GetSlackIntegrationResponse) Reset() {
	*x = GetSlackIntegrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSlackIntegrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSlackIntegrationResponse) ProtoMessage() {}

func (x *GetSlackIntegrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSlackIntegrationResponse.ProtoReflect.Descriptor instead.
func (*GetSlackIntegrationResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{9}
}

func (x *GetSlackIntegrationResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type RemoveSlackIntegrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveSlackIntegrationRequest) Reset() {
	*x = RemoveSlackIntegrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSlackIntegrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSlackIntegrationRequest) ProtoMessage() {}

func (x *RemoveSlackIntegrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSlackIntegrationRequest.ProtoReflect.Descriptor instead.
func (*RemoveSlackIntegrationRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{10}
}

type RemoveSlackIntegrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveSlackIntegrationResponse) Reset() {
	*x = RemoveSlackIntegrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSlackIntegrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSlackIntegrationResponse) ProtoMessage() {}

func (x *RemoveSlackIntegrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSlackIntegrationResponse.ProtoReflect.Descriptor instead.
func (*RemoveSlackIntegrationResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP(), []int{11}
}

func (x *RemoveSlackIntegrationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_commonfate_control_oauth_v1alpha1_oauth_proto protoreflect.FileDescriptor

var file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x22, 0x23, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x44, 0x75, 0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22,
	0x20, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x39, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74,
	0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x23, 0x0a, 0x21,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3e, 0x0a, 0x22, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x67, 0x65, 0x72,
	0x44, 0x75, 0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x32, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x61, 0x63, 0x6b,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x1e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x9c, 0x05, 0x0a, 0x0c, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa2, 0x01, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xab,
	0x01, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75,
	0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74,
	0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x96, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x61, 0x63, 0x6b,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9f, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x53, 0x6c, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6c, 0x61,
	0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xaa, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x0a, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x43, 0x4f, 0xaa, 0x02, 0x21, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x4f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x21, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x4f,
	0x61, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2d,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5c, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescOnce sync.Once
	file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescData = file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDesc
)

func file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescGZIP() []byte {
	file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescOnce.Do(func() {
		file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescData)
	})
	return file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDescData
}

var file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_commonfate_control_oauth_v1alpha1_oauth_proto_goTypes = []interface{}{
	(*CreatePagerDutyIntegrationRequest)(nil),  // 0: commonfate.control.oauth.v1alpha1.CreatePagerDutyIntegrationRequest
	(*CreatePagerDutyIntegrationResponse)(nil), // 1: commonfate.control.oauth.v1alpha1.CreatePagerDutyIntegrationResponse
	(*GetPagerDutyIntegrationRequest)(nil),     // 2: commonfate.control.oauth.v1alpha1.GetPagerDutyIntegrationRequest
	(*GetPagerDutyIntegrationResponse)(nil),    // 3: commonfate.control.oauth.v1alpha1.GetPagerDutyIntegrationResponse
	(*RemovePagerDutyIntegrationRequest)(nil),  // 4: commonfate.control.oauth.v1alpha1.RemovePagerDutyIntegrationRequest
	(*RemovePagerDutyIntegrationResponse)(nil), // 5: commonfate.control.oauth.v1alpha1.RemovePagerDutyIntegrationResponse
	(*CreateSlackIntegrationRequest)(nil),      // 6: commonfate.control.oauth.v1alpha1.CreateSlackIntegrationRequest
	(*CreateSlackIntegrationResponse)(nil),     // 7: commonfate.control.oauth.v1alpha1.CreateSlackIntegrationResponse
	(*GetSlackIntegrationRequest)(nil),         // 8: commonfate.control.oauth.v1alpha1.GetSlackIntegrationRequest
	(*GetSlackIntegrationResponse)(nil),        // 9: commonfate.control.oauth.v1alpha1.GetSlackIntegrationResponse
	(*RemoveSlackIntegrationRequest)(nil),      // 10: commonfate.control.oauth.v1alpha1.RemoveSlackIntegrationRequest
	(*RemoveSlackIntegrationResponse)(nil),     // 11: commonfate.control.oauth.v1alpha1.RemoveSlackIntegrationResponse
}
var file_commonfate_control_oauth_v1alpha1_oauth_proto_depIdxs = []int32{
	2,  // 0: commonfate.control.oauth.v1alpha1.OAuthService.GetPagerDutyIntegration:input_type -> commonfate.control.oauth.v1alpha1.GetPagerDutyIntegrationRequest
	4,  // 1: commonfate.control.oauth.v1alpha1.OAuthService.RemovePagerDutyIntegration:input_type -> commonfate.control.oauth.v1alpha1.RemovePagerDutyIntegrationRequest
	8,  // 2: commonfate.control.oauth.v1alpha1.OAuthService.GetSlackIntegration:input_type -> commonfate.control.oauth.v1alpha1.GetSlackIntegrationRequest
	10, // 3: commonfate.control.oauth.v1alpha1.OAuthService.RemoveSlackIntegration:input_type -> commonfate.control.oauth.v1alpha1.RemoveSlackIntegrationRequest
	3,  // 4: commonfate.control.oauth.v1alpha1.OAuthService.GetPagerDutyIntegration:output_type -> commonfate.control.oauth.v1alpha1.GetPagerDutyIntegrationResponse
	5,  // 5: commonfate.control.oauth.v1alpha1.OAuthService.RemovePagerDutyIntegration:output_type -> commonfate.control.oauth.v1alpha1.RemovePagerDutyIntegrationResponse
	9,  // 6: commonfate.control.oauth.v1alpha1.OAuthService.GetSlackIntegration:output_type -> commonfate.control.oauth.v1alpha1.GetSlackIntegrationResponse
	11, // 7: commonfate.control.oauth.v1alpha1.OAuthService.RemoveSlackIntegration:output_type -> commonfate.control.oauth.v1alpha1.RemoveSlackIntegrationResponse
	4,  // [4:8] is the sub-list for method output_type
	0,  // [0:4] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_commonfate_control_oauth_v1alpha1_oauth_proto_init() }
func file_commonfate_control_oauth_v1alpha1_oauth_proto_init() {
	if File_commonfate_control_oauth_v1alpha1_oauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePagerDutyIntegrationRequest); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePagerDutyIntegrationResponse); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPagerDutyIntegrationRequest); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPagerDutyIntegrationResponse); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePagerDutyIntegrationRequest); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePagerDutyIntegrationResponse); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlackIntegrationRequest); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlackIntegrationResponse); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSlackIntegrationRequest); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSlackIntegrationResponse); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSlackIntegrationRequest); i {
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
		file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSlackIntegrationResponse); i {
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
			RawDescriptor: file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_oauth_v1alpha1_oauth_proto_goTypes,
		DependencyIndexes: file_commonfate_control_oauth_v1alpha1_oauth_proto_depIdxs,
		MessageInfos:      file_commonfate_control_oauth_v1alpha1_oauth_proto_msgTypes,
	}.Build()
	File_commonfate_control_oauth_v1alpha1_oauth_proto = out.File
	file_commonfate_control_oauth_v1alpha1_oauth_proto_rawDesc = nil
	file_commonfate_control_oauth_v1alpha1_oauth_proto_goTypes = nil
	file_commonfate_control_oauth_v1alpha1_oauth_proto_depIdxs = nil
}
