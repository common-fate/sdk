// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: commonfate/control/integration/reset/v1alpha1/reset.proto

package resetv1alpha1

import (
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
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

type ResetEntraUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, will return a preview of changes rather than actually resetting them.
	DryRun bool `protobuf:"varint,1,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
}

func (x *ResetEntraUsersRequest) Reset() {
	*x = ResetEntraUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetEntraUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetEntraUsersRequest) ProtoMessage() {}

func (x *ResetEntraUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetEntraUsersRequest.ProtoReflect.Descriptor instead.
func (*ResetEntraUsersRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{0}
}

func (x *ResetEntraUsersRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

type ResetEntraUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletedEntities []*v1alpha1.EID `protobuf:"bytes,1,rep,name=deleted_entities,json=deletedEntities,proto3" json:"deleted_entities,omitempty"`
}

func (x *ResetEntraUsersResponse) Reset() {
	*x = ResetEntraUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetEntraUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetEntraUsersResponse) ProtoMessage() {}

func (x *ResetEntraUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetEntraUsersResponse.ProtoReflect.Descriptor instead.
func (*ResetEntraUsersResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{1}
}

func (x *ResetEntraUsersResponse) GetDeletedEntities() []*v1alpha1.EID {
	if x != nil {
		return x.DeletedEntities
	}
	return nil
}

type RelinkEntraUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RelinkEntraUsersRequest) Reset() {
	*x = RelinkEntraUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelinkEntraUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelinkEntraUsersRequest) ProtoMessage() {}

func (x *RelinkEntraUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelinkEntraUsersRequest.ProtoReflect.Descriptor instead.
func (*RelinkEntraUsersRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{2}
}

type RelinkEntraUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RelinkEntraUsersResponse) Reset() {
	*x = RelinkEntraUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelinkEntraUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelinkEntraUsersResponse) ProtoMessage() {}

func (x *RelinkEntraUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelinkEntraUsersResponse.ProtoReflect.Descriptor instead.
func (*RelinkEntraUsersResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{3}
}

type RemoveOAuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the token to remove.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveOAuthTokenRequest) Reset() {
	*x = RemoveOAuthTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveOAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOAuthTokenRequest) ProtoMessage() {}

func (x *RemoveOAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*RemoveOAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveOAuthTokenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveOAuthTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveOAuthTokenResponse) Reset() {
	*x = RemoveOAuthTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveOAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOAuthTokenResponse) ProtoMessage() {}

func (x *RemoveOAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*RemoveOAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{5}
}

type CancelBackgroundJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The job id to cancel
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelBackgroundJobRequest) Reset() {
	*x = CancelBackgroundJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelBackgroundJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBackgroundJobRequest) ProtoMessage() {}

func (x *CancelBackgroundJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBackgroundJobRequest.ProtoReflect.Descriptor instead.
func (*CancelBackgroundJobRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{6}
}

func (x *CancelBackgroundJobRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CancelBackgroundJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelBackgroundJobResponse) Reset() {
	*x = CancelBackgroundJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelBackgroundJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelBackgroundJobResponse) ProtoMessage() {}

func (x *CancelBackgroundJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelBackgroundJobResponse.ProtoReflect.Descriptor instead.
func (*CancelBackgroundJobResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{7}
}

type RetryBackgroundJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The job id to retry
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RetryBackgroundJobRequest) Reset() {
	*x = RetryBackgroundJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryBackgroundJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryBackgroundJobRequest) ProtoMessage() {}

func (x *RetryBackgroundJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryBackgroundJobRequest.ProtoReflect.Descriptor instead.
func (*RetryBackgroundJobRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{8}
}

func (x *RetryBackgroundJobRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RetryBackgroundJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RetryBackgroundJobResponse) Reset() {
	*x = RetryBackgroundJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryBackgroundJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryBackgroundJobResponse) ProtoMessage() {}

func (x *RetryBackgroundJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryBackgroundJobResponse.ProtoReflect.Descriptor instead.
func (*RetryBackgroundJobResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP(), []int{9}
}

var File_commonfate_control_integration_reset_v1alpha1_reset_proto protoreflect.FileDescriptor

var file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x31, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72,
	0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79,
	0x52, 0x75, 0x6e, 0x22, 0x65, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a,
	0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65,
	0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x45,
	0x6e, 0x74, 0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x29, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x1a, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a, 0x19, 0x52, 0x65, 0x74, 0x72, 0x79, 0x42, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x52, 0x65, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xe2, 0x06, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xa2, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x61,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x45, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x61,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa5, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x69, 0x6e,
	0x6b, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x46, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x69,
	0x6e, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa5,
	0x01, 0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x46, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xae, 0x01, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x12, 0x49,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xab, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x12, 0x48,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x42, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x49, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x42, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xf4, 0x02, 0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x65, 0x74, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x04, 0x43, 0x43, 0x49, 0x52, 0xaa, 0x02, 0x2d, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x2d, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x39, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x31, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescOnce sync.Once
	file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescData = file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDesc
)

func file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescGZIP() []byte {
	file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescOnce.Do(func() {
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescData)
	})
	return file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDescData
}

var file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_commonfate_control_integration_reset_v1alpha1_reset_proto_goTypes = []interface{}{
	(*ResetEntraUsersRequest)(nil),      // 0: commonfate.control.integration.reset.v1alpha1.ResetEntraUsersRequest
	(*ResetEntraUsersResponse)(nil),     // 1: commonfate.control.integration.reset.v1alpha1.ResetEntraUsersResponse
	(*RelinkEntraUsersRequest)(nil),     // 2: commonfate.control.integration.reset.v1alpha1.RelinkEntraUsersRequest
	(*RelinkEntraUsersResponse)(nil),    // 3: commonfate.control.integration.reset.v1alpha1.RelinkEntraUsersResponse
	(*RemoveOAuthTokenRequest)(nil),     // 4: commonfate.control.integration.reset.v1alpha1.RemoveOAuthTokenRequest
	(*RemoveOAuthTokenResponse)(nil),    // 5: commonfate.control.integration.reset.v1alpha1.RemoveOAuthTokenResponse
	(*CancelBackgroundJobRequest)(nil),  // 6: commonfate.control.integration.reset.v1alpha1.CancelBackgroundJobRequest
	(*CancelBackgroundJobResponse)(nil), // 7: commonfate.control.integration.reset.v1alpha1.CancelBackgroundJobResponse
	(*RetryBackgroundJobRequest)(nil),   // 8: commonfate.control.integration.reset.v1alpha1.RetryBackgroundJobRequest
	(*RetryBackgroundJobResponse)(nil),  // 9: commonfate.control.integration.reset.v1alpha1.RetryBackgroundJobResponse
	(*v1alpha1.EID)(nil),                // 10: commonfate.entity.v1alpha1.EID
}
var file_commonfate_control_integration_reset_v1alpha1_reset_proto_depIdxs = []int32{
	10, // 0: commonfate.control.integration.reset.v1alpha1.ResetEntraUsersResponse.deleted_entities:type_name -> commonfate.entity.v1alpha1.EID
	0,  // 1: commonfate.control.integration.reset.v1alpha1.ResetService.ResetEntraUsers:input_type -> commonfate.control.integration.reset.v1alpha1.ResetEntraUsersRequest
	2,  // 2: commonfate.control.integration.reset.v1alpha1.ResetService.RelinkEntraUsers:input_type -> commonfate.control.integration.reset.v1alpha1.RelinkEntraUsersRequest
	4,  // 3: commonfate.control.integration.reset.v1alpha1.ResetService.RemoveOAuthToken:input_type -> commonfate.control.integration.reset.v1alpha1.RemoveOAuthTokenRequest
	6,  // 4: commonfate.control.integration.reset.v1alpha1.ResetService.CancelBackgroundJob:input_type -> commonfate.control.integration.reset.v1alpha1.CancelBackgroundJobRequest
	8,  // 5: commonfate.control.integration.reset.v1alpha1.ResetService.RetryBackgroundJob:input_type -> commonfate.control.integration.reset.v1alpha1.RetryBackgroundJobRequest
	1,  // 6: commonfate.control.integration.reset.v1alpha1.ResetService.ResetEntraUsers:output_type -> commonfate.control.integration.reset.v1alpha1.ResetEntraUsersResponse
	3,  // 7: commonfate.control.integration.reset.v1alpha1.ResetService.RelinkEntraUsers:output_type -> commonfate.control.integration.reset.v1alpha1.RelinkEntraUsersResponse
	5,  // 8: commonfate.control.integration.reset.v1alpha1.ResetService.RemoveOAuthToken:output_type -> commonfate.control.integration.reset.v1alpha1.RemoveOAuthTokenResponse
	7,  // 9: commonfate.control.integration.reset.v1alpha1.ResetService.CancelBackgroundJob:output_type -> commonfate.control.integration.reset.v1alpha1.CancelBackgroundJobResponse
	9,  // 10: commonfate.control.integration.reset.v1alpha1.ResetService.RetryBackgroundJob:output_type -> commonfate.control.integration.reset.v1alpha1.RetryBackgroundJobResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_commonfate_control_integration_reset_v1alpha1_reset_proto_init() }
func file_commonfate_control_integration_reset_v1alpha1_reset_proto_init() {
	if File_commonfate_control_integration_reset_v1alpha1_reset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetEntraUsersRequest); i {
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
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetEntraUsersResponse); i {
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
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelinkEntraUsersRequest); i {
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
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelinkEntraUsersResponse); i {
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
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveOAuthTokenRequest); i {
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
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveOAuthTokenResponse); i {
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
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelBackgroundJobRequest); i {
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
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelBackgroundJobResponse); i {
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
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryBackgroundJobRequest); i {
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
		file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryBackgroundJobResponse); i {
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
			RawDescriptor: file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_integration_reset_v1alpha1_reset_proto_goTypes,
		DependencyIndexes: file_commonfate_control_integration_reset_v1alpha1_reset_proto_depIdxs,
		MessageInfos:      file_commonfate_control_integration_reset_v1alpha1_reset_proto_msgTypes,
	}.Build()
	File_commonfate_control_integration_reset_v1alpha1_reset_proto = out.File
	file_commonfate_control_integration_reset_v1alpha1_reset_proto_rawDesc = nil
	file_commonfate_control_integration_reset_v1alpha1_reset_proto_goTypes = nil
	file_commonfate_control_integration_reset_v1alpha1_reset_proto_depIdxs = nil
}
