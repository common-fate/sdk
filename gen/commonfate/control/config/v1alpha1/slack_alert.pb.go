// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: commonfate/control/config/v1alpha1/slack_alert.proto

package configv1alpha1

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

type CreateSlackAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId       string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	SlackChannelId   string `protobuf:"bytes,3,opt,name=slack_channel_id,json=slackChannelId,proto3" json:"slack_channel_id,omitempty"`
	SlackWorkspaceId string `protobuf:"bytes,4,opt,name=slack_workspace_id,json=slackWorkspaceId,proto3" json:"slack_workspace_id,omitempty"`
	IntegrationId    string `protobuf:"bytes,5,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
}

func (x *CreateSlackAlertRequest) Reset() {
	*x = CreateSlackAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlackAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlackAlertRequest) ProtoMessage() {}

func (x *CreateSlackAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlackAlertRequest.ProtoReflect.Descriptor instead.
func (*CreateSlackAlertRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSlackAlertRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *CreateSlackAlertRequest) GetSlackChannelId() string {
	if x != nil {
		return x.SlackChannelId
	}
	return ""
}

func (x *CreateSlackAlertRequest) GetSlackWorkspaceId() string {
	if x != nil {
		return x.SlackWorkspaceId
	}
	return ""
}

func (x *CreateSlackAlertRequest) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

type SlackAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkflowId       string  `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	SlackChannelId   string  `protobuf:"bytes,3,opt,name=slack_channel_id,json=slackChannelId,proto3" json:"slack_channel_id,omitempty"`
	SlackWorkspaceId string  `protobuf:"bytes,4,opt,name=slack_workspace_id,json=slackWorkspaceId,proto3" json:"slack_workspace_id,omitempty"`
	IntegrationId    *string `protobuf:"bytes,5,opt,name=integration_id,json=integrationId,proto3,oneof" json:"integration_id,omitempty"`
}

func (x *SlackAlert) Reset() {
	*x = SlackAlert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlackAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlackAlert) ProtoMessage() {}

func (x *SlackAlert) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlackAlert.ProtoReflect.Descriptor instead.
func (*SlackAlert) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP(), []int{1}
}

func (x *SlackAlert) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SlackAlert) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *SlackAlert) GetSlackChannelId() string {
	if x != nil {
		return x.SlackChannelId
	}
	return ""
}

func (x *SlackAlert) GetSlackWorkspaceId() string {
	if x != nil {
		return x.SlackWorkspaceId
	}
	return ""
}

func (x *SlackAlert) GetIntegrationId() string {
	if x != nil && x.IntegrationId != nil {
		return *x.IntegrationId
	}
	return ""
}

type CreateSlackAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert *SlackAlert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *CreateSlackAlertResponse) Reset() {
	*x = CreateSlackAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSlackAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSlackAlertResponse) ProtoMessage() {}

func (x *CreateSlackAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSlackAlertResponse.ProtoReflect.Descriptor instead.
func (*CreateSlackAlertResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSlackAlertResponse) GetAlert() *SlackAlert {
	if x != nil {
		return x.Alert
	}
	return nil
}

type GetSlackAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSlackAlertRequest) Reset() {
	*x = GetSlackAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSlackAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSlackAlertRequest) ProtoMessage() {}

func (x *GetSlackAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSlackAlertRequest.ProtoReflect.Descriptor instead.
func (*GetSlackAlertRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP(), []int{3}
}

func (x *GetSlackAlertRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSlackAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert *SlackAlert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *GetSlackAlertResponse) Reset() {
	*x = GetSlackAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSlackAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSlackAlertResponse) ProtoMessage() {}

func (x *GetSlackAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSlackAlertResponse.ProtoReflect.Descriptor instead.
func (*GetSlackAlertResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP(), []int{4}
}

func (x *GetSlackAlertResponse) GetAlert() *SlackAlert {
	if x != nil {
		return x.Alert
	}
	return nil
}

type UpdateSlackAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert *SlackAlert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *UpdateSlackAlertRequest) Reset() {
	*x = UpdateSlackAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSlackAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSlackAlertRequest) ProtoMessage() {}

func (x *UpdateSlackAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSlackAlertRequest.ProtoReflect.Descriptor instead.
func (*UpdateSlackAlertRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSlackAlertRequest) GetAlert() *SlackAlert {
	if x != nil {
		return x.Alert
	}
	return nil
}

type UpdateSlackAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alert *SlackAlert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
}

func (x *UpdateSlackAlertResponse) Reset() {
	*x = UpdateSlackAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSlackAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSlackAlertResponse) ProtoMessage() {}

func (x *UpdateSlackAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSlackAlertResponse.ProtoReflect.Descriptor instead.
func (*UpdateSlackAlertResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSlackAlertResponse) GetAlert() *SlackAlert {
	if x != nil {
		return x.Alert
	}
	return nil
}

type DeleteSlackAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSlackAlertRequest) Reset() {
	*x = DeleteSlackAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSlackAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSlackAlertRequest) ProtoMessage() {}

func (x *DeleteSlackAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSlackAlertRequest.ProtoReflect.Descriptor instead.
func (*DeleteSlackAlertRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSlackAlertRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteSlackAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSlackAlertResponse) Reset() {
	*x = DeleteSlackAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSlackAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSlackAlertResponse) ProtoMessage() {}

func (x *DeleteSlackAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSlackAlertResponse.ProtoReflect.Descriptor instead.
func (*DeleteSlackAlertResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSlackAlertResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_commonfate_control_config_v1alpha1_slack_alert_proto protoreflect.FileDescriptor

var file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0xb9, 0x01, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6c, 0x61, 0x63, 0x6b,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xd4, 0x01, 0x0a, 0x0a, 0x53, 0x6c, 0x61, 0x63, 0x6b,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x6c,
	0x61, 0x63, 0x6b, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2a,
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x60, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x6c,
	0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22,
	0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x6c,
	0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52,
	0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x5f, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x44, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x60, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x05, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6c,
	0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x32, 0xd2, 0x04, 0x0a, 0x11, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x3b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x86, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x61,
	0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x8f, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6c, 0x61, 0x63,
	0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x8f, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6c,
	0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb6, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x0f, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x43, 0xaa, 0x02, 0x22, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescOnce sync.Once
	file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescData = file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDesc
)

func file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescGZIP() []byte {
	file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescOnce.Do(func() {
		file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescData)
	})
	return file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDescData
}

var file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_commonfate_control_config_v1alpha1_slack_alert_proto_goTypes = []interface{}{
	(*CreateSlackAlertRequest)(nil),  // 0: commonfate.control.config.v1alpha1.CreateSlackAlertRequest
	(*SlackAlert)(nil),               // 1: commonfate.control.config.v1alpha1.SlackAlert
	(*CreateSlackAlertResponse)(nil), // 2: commonfate.control.config.v1alpha1.CreateSlackAlertResponse
	(*GetSlackAlertRequest)(nil),     // 3: commonfate.control.config.v1alpha1.GetSlackAlertRequest
	(*GetSlackAlertResponse)(nil),    // 4: commonfate.control.config.v1alpha1.GetSlackAlertResponse
	(*UpdateSlackAlertRequest)(nil),  // 5: commonfate.control.config.v1alpha1.UpdateSlackAlertRequest
	(*UpdateSlackAlertResponse)(nil), // 6: commonfate.control.config.v1alpha1.UpdateSlackAlertResponse
	(*DeleteSlackAlertRequest)(nil),  // 7: commonfate.control.config.v1alpha1.DeleteSlackAlertRequest
	(*DeleteSlackAlertResponse)(nil), // 8: commonfate.control.config.v1alpha1.DeleteSlackAlertResponse
}
var file_commonfate_control_config_v1alpha1_slack_alert_proto_depIdxs = []int32{
	1, // 0: commonfate.control.config.v1alpha1.CreateSlackAlertResponse.alert:type_name -> commonfate.control.config.v1alpha1.SlackAlert
	1, // 1: commonfate.control.config.v1alpha1.GetSlackAlertResponse.alert:type_name -> commonfate.control.config.v1alpha1.SlackAlert
	1, // 2: commonfate.control.config.v1alpha1.UpdateSlackAlertRequest.alert:type_name -> commonfate.control.config.v1alpha1.SlackAlert
	1, // 3: commonfate.control.config.v1alpha1.UpdateSlackAlertResponse.alert:type_name -> commonfate.control.config.v1alpha1.SlackAlert
	0, // 4: commonfate.control.config.v1alpha1.SlackAlertService.CreateSlackAlert:input_type -> commonfate.control.config.v1alpha1.CreateSlackAlertRequest
	3, // 5: commonfate.control.config.v1alpha1.SlackAlertService.GetSlackAlert:input_type -> commonfate.control.config.v1alpha1.GetSlackAlertRequest
	5, // 6: commonfate.control.config.v1alpha1.SlackAlertService.UpdateSlackAlert:input_type -> commonfate.control.config.v1alpha1.UpdateSlackAlertRequest
	7, // 7: commonfate.control.config.v1alpha1.SlackAlertService.DeleteSlackAlert:input_type -> commonfate.control.config.v1alpha1.DeleteSlackAlertRequest
	2, // 8: commonfate.control.config.v1alpha1.SlackAlertService.CreateSlackAlert:output_type -> commonfate.control.config.v1alpha1.CreateSlackAlertResponse
	4, // 9: commonfate.control.config.v1alpha1.SlackAlertService.GetSlackAlert:output_type -> commonfate.control.config.v1alpha1.GetSlackAlertResponse
	6, // 10: commonfate.control.config.v1alpha1.SlackAlertService.UpdateSlackAlert:output_type -> commonfate.control.config.v1alpha1.UpdateSlackAlertResponse
	8, // 11: commonfate.control.config.v1alpha1.SlackAlertService.DeleteSlackAlert:output_type -> commonfate.control.config.v1alpha1.DeleteSlackAlertResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_commonfate_control_config_v1alpha1_slack_alert_proto_init() }
func file_commonfate_control_config_v1alpha1_slack_alert_proto_init() {
	if File_commonfate_control_config_v1alpha1_slack_alert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlackAlertRequest); i {
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
		file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlackAlert); i {
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
		file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSlackAlertResponse); i {
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
		file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSlackAlertRequest); i {
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
		file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSlackAlertResponse); i {
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
		file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSlackAlertRequest); i {
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
		file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSlackAlertResponse); i {
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
		file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSlackAlertRequest); i {
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
		file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSlackAlertResponse); i {
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
	file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_config_v1alpha1_slack_alert_proto_goTypes,
		DependencyIndexes: file_commonfate_control_config_v1alpha1_slack_alert_proto_depIdxs,
		MessageInfos:      file_commonfate_control_config_v1alpha1_slack_alert_proto_msgTypes,
	}.Build()
	File_commonfate_control_config_v1alpha1_slack_alert_proto = out.File
	file_commonfate_control_config_v1alpha1_slack_alert_proto_rawDesc = nil
	file_commonfate_control_config_v1alpha1_slack_alert_proto_goTypes = nil
	file_commonfate_control_config_v1alpha1_slack_alert_proto_depIdxs = nil
}
