// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: commonfate/control/config/v1alpha1/access_workflow.proto

package configv1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	_ "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccessWorkflow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AccessDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=access_duration,json=accessDuration,proto3" json:"access_duration,omitempty"`
	TryExtendAfter *durationpb.Duration `protobuf:"bytes,4,opt,name=try_extend_after,json=tryExtendAfter,proto3" json:"try_extend_after,omitempty"`
	Priority       int32                `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *AccessWorkflow) Reset() {
	*x = AccessWorkflow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessWorkflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessWorkflow) ProtoMessage() {}

func (x *AccessWorkflow) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessWorkflow.ProtoReflect.Descriptor instead.
func (*AccessWorkflow) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *AccessWorkflow) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccessWorkflow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccessWorkflow) GetAccessDuration() *durationpb.Duration {
	if x != nil {
		return x.AccessDuration
	}
	return nil
}

func (x *AccessWorkflow) GetTryExtendAfter() *durationpb.Duration {
	if x != nil {
		return x.TryExtendAfter
	}
	return nil
}

func (x *AccessWorkflow) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type CreateAccessWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AccessDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=access_duration,json=accessDuration,proto3" json:"access_duration,omitempty"`
	TryExtendAfter *durationpb.Duration `protobuf:"bytes,4,opt,name=try_extend_after,json=tryExtendAfter,proto3" json:"try_extend_after,omitempty"`
	Priority       int32                `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *CreateAccessWorkflowRequest) Reset() {
	*x = CreateAccessWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccessWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessWorkflowRequest) ProtoMessage() {}

func (x *CreateAccessWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessWorkflowRequest.ProtoReflect.Descriptor instead.
func (*CreateAccessWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccessWorkflowRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAccessWorkflowRequest) GetAccessDuration() *durationpb.Duration {
	if x != nil {
		return x.AccessDuration
	}
	return nil
}

func (x *CreateAccessWorkflowRequest) GetTryExtendAfter() *durationpb.Duration {
	if x != nil {
		return x.TryExtendAfter
	}
	return nil
}

func (x *CreateAccessWorkflowRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type CreateAccessWorkflowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workflow *AccessWorkflow `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
}

func (x *CreateAccessWorkflowResponse) Reset() {
	*x = CreateAccessWorkflowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccessWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessWorkflowResponse) ProtoMessage() {}

func (x *CreateAccessWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessWorkflowResponse.ProtoReflect.Descriptor instead.
func (*CreateAccessWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccessWorkflowResponse) GetWorkflow() *AccessWorkflow {
	if x != nil {
		return x.Workflow
	}
	return nil
}

type GetAccessWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAccessWorkflowRequest) Reset() {
	*x = GetAccessWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessWorkflowRequest) ProtoMessage() {}

func (x *GetAccessWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessWorkflowRequest.ProtoReflect.Descriptor instead.
func (*GetAccessWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccessWorkflowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAccessWorkflowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workflow *AccessWorkflow `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
}

func (x *GetAccessWorkflowResponse) Reset() {
	*x = GetAccessWorkflowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessWorkflowResponse) ProtoMessage() {}

func (x *GetAccessWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessWorkflowResponse.ProtoReflect.Descriptor instead.
func (*GetAccessWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccessWorkflowResponse) GetWorkflow() *AccessWorkflow {
	if x != nil {
		return x.Workflow
	}
	return nil
}

type UpdateAccessWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workflow *AccessWorkflow `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
}

func (x *UpdateAccessWorkflowRequest) Reset() {
	*x = UpdateAccessWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccessWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccessWorkflowRequest) ProtoMessage() {}

func (x *UpdateAccessWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccessWorkflowRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccessWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAccessWorkflowRequest) GetWorkflow() *AccessWorkflow {
	if x != nil {
		return x.Workflow
	}
	return nil
}

type UpdateAccessWorkflowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workflow *AccessWorkflow `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
}

func (x *UpdateAccessWorkflowResponse) Reset() {
	*x = UpdateAccessWorkflowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccessWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccessWorkflowResponse) ProtoMessage() {}

func (x *UpdateAccessWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccessWorkflowResponse.ProtoReflect.Descriptor instead.
func (*UpdateAccessWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAccessWorkflowResponse) GetWorkflow() *AccessWorkflow {
	if x != nil {
		return x.Workflow
	}
	return nil
}

type DeleteAccessWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAccessWorkflowRequest) Reset() {
	*x = DeleteAccessWorkflowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccessWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessWorkflowRequest) ProtoMessage() {}

func (x *DeleteAccessWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessWorkflowRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccessWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAccessWorkflowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAccessWorkflowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAccessWorkflowResponse) Reset() {
	*x = DeleteAccessWorkflowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccessWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessWorkflowResponse) ProtoMessage() {}

func (x *DeleteAccessWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessWorkflowResponse.ProtoReflect.Descriptor instead.
func (*DeleteAccessWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAccessWorkflowResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_commonfate_control_config_v1alpha1_access_workflow_proto protoreflect.FileDescriptor

var file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd9, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x10, 0x74, 0x72,
	0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x74, 0x72, 0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xd6, 0x01, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x42, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x10, 0x74, 0x72, 0x79, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x74, 0x72, 0x79, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x22, 0x6e, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x6b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x6d, 0x0a,
	0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x6e, 0x0a, 0x1c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x08,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x2d, 0x0a, 0x1b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x1c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x8a, 0x05, 0x0a, 0x15,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9b, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x3f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x96, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x9b, 0x01, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9b, 0x01, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xba, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x43, 0x43, 0xaa, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescOnce sync.Once
	file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescData = file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDesc
)

func file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescGZIP() []byte {
	file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescOnce.Do(func() {
		file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescData)
	})
	return file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDescData
}

var file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_commonfate_control_config_v1alpha1_access_workflow_proto_goTypes = []interface{}{
	(*AccessWorkflow)(nil),               // 0: commonfate.control.config.v1alpha1.AccessWorkflow
	(*CreateAccessWorkflowRequest)(nil),  // 1: commonfate.control.config.v1alpha1.CreateAccessWorkflowRequest
	(*CreateAccessWorkflowResponse)(nil), // 2: commonfate.control.config.v1alpha1.CreateAccessWorkflowResponse
	(*GetAccessWorkflowRequest)(nil),     // 3: commonfate.control.config.v1alpha1.GetAccessWorkflowRequest
	(*GetAccessWorkflowResponse)(nil),    // 4: commonfate.control.config.v1alpha1.GetAccessWorkflowResponse
	(*UpdateAccessWorkflowRequest)(nil),  // 5: commonfate.control.config.v1alpha1.UpdateAccessWorkflowRequest
	(*UpdateAccessWorkflowResponse)(nil), // 6: commonfate.control.config.v1alpha1.UpdateAccessWorkflowResponse
	(*DeleteAccessWorkflowRequest)(nil),  // 7: commonfate.control.config.v1alpha1.DeleteAccessWorkflowRequest
	(*DeleteAccessWorkflowResponse)(nil), // 8: commonfate.control.config.v1alpha1.DeleteAccessWorkflowResponse
	(*durationpb.Duration)(nil),          // 9: google.protobuf.Duration
}
var file_commonfate_control_config_v1alpha1_access_workflow_proto_depIdxs = []int32{
	9,  // 0: commonfate.control.config.v1alpha1.AccessWorkflow.access_duration:type_name -> google.protobuf.Duration
	9,  // 1: commonfate.control.config.v1alpha1.AccessWorkflow.try_extend_after:type_name -> google.protobuf.Duration
	9,  // 2: commonfate.control.config.v1alpha1.CreateAccessWorkflowRequest.access_duration:type_name -> google.protobuf.Duration
	9,  // 3: commonfate.control.config.v1alpha1.CreateAccessWorkflowRequest.try_extend_after:type_name -> google.protobuf.Duration
	0,  // 4: commonfate.control.config.v1alpha1.CreateAccessWorkflowResponse.workflow:type_name -> commonfate.control.config.v1alpha1.AccessWorkflow
	0,  // 5: commonfate.control.config.v1alpha1.GetAccessWorkflowResponse.workflow:type_name -> commonfate.control.config.v1alpha1.AccessWorkflow
	0,  // 6: commonfate.control.config.v1alpha1.UpdateAccessWorkflowRequest.workflow:type_name -> commonfate.control.config.v1alpha1.AccessWorkflow
	0,  // 7: commonfate.control.config.v1alpha1.UpdateAccessWorkflowResponse.workflow:type_name -> commonfate.control.config.v1alpha1.AccessWorkflow
	1,  // 8: commonfate.control.config.v1alpha1.AccessWorkflowService.CreateAccessWorkflow:input_type -> commonfate.control.config.v1alpha1.CreateAccessWorkflowRequest
	3,  // 9: commonfate.control.config.v1alpha1.AccessWorkflowService.GetAccessWorkflow:input_type -> commonfate.control.config.v1alpha1.GetAccessWorkflowRequest
	5,  // 10: commonfate.control.config.v1alpha1.AccessWorkflowService.UpdateAccessWorkflow:input_type -> commonfate.control.config.v1alpha1.UpdateAccessWorkflowRequest
	7,  // 11: commonfate.control.config.v1alpha1.AccessWorkflowService.DeleteAccessWorkflow:input_type -> commonfate.control.config.v1alpha1.DeleteAccessWorkflowRequest
	2,  // 12: commonfate.control.config.v1alpha1.AccessWorkflowService.CreateAccessWorkflow:output_type -> commonfate.control.config.v1alpha1.CreateAccessWorkflowResponse
	4,  // 13: commonfate.control.config.v1alpha1.AccessWorkflowService.GetAccessWorkflow:output_type -> commonfate.control.config.v1alpha1.GetAccessWorkflowResponse
	6,  // 14: commonfate.control.config.v1alpha1.AccessWorkflowService.UpdateAccessWorkflow:output_type -> commonfate.control.config.v1alpha1.UpdateAccessWorkflowResponse
	8,  // 15: commonfate.control.config.v1alpha1.AccessWorkflowService.DeleteAccessWorkflow:output_type -> commonfate.control.config.v1alpha1.DeleteAccessWorkflowResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_commonfate_control_config_v1alpha1_access_workflow_proto_init() }
func file_commonfate_control_config_v1alpha1_access_workflow_proto_init() {
	if File_commonfate_control_config_v1alpha1_access_workflow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessWorkflow); i {
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
		file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccessWorkflowRequest); i {
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
		file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccessWorkflowResponse); i {
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
		file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessWorkflowRequest); i {
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
		file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessWorkflowResponse); i {
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
		file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccessWorkflowRequest); i {
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
		file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccessWorkflowResponse); i {
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
		file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccessWorkflowRequest); i {
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
		file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccessWorkflowResponse); i {
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
			RawDescriptor: file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_config_v1alpha1_access_workflow_proto_goTypes,
		DependencyIndexes: file_commonfate_control_config_v1alpha1_access_workflow_proto_depIdxs,
		MessageInfos:      file_commonfate_control_config_v1alpha1_access_workflow_proto_msgTypes,
	}.Build()
	File_commonfate_control_config_v1alpha1_access_workflow_proto = out.File
	file_commonfate_control_config_v1alpha1_access_workflow_proto_rawDesc = nil
	file_commonfate_control_config_v1alpha1_access_workflow_proto_goTypes = nil
	file_commonfate_control_config_v1alpha1_access_workflow_proto_depIdxs = nil
}
