// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: commonfate/control/notification/v1alpha1/notification.proto

package notificationv1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
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

type GetUserUserNotificationSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserUserNotificationSettingsRequest) Reset() {
	*x = GetUserUserNotificationSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserUserNotificationSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserUserNotificationSettingsRequest) ProtoMessage() {}

func (x *GetUserUserNotificationSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserUserNotificationSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetUserUserNotificationSettingsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_notification_v1alpha1_notification_proto_rawDescGZIP(), []int{0}
}

type GetUserUserNotificationSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notifications []*UserNotificationSettings `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *GetUserUserNotificationSettingsResponse) Reset() {
	*x = GetUserUserNotificationSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserUserNotificationSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserUserNotificationSettingsResponse) ProtoMessage() {}

func (x *GetUserUserNotificationSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserUserNotificationSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetUserUserNotificationSettingsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_notification_v1alpha1_notification_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserUserNotificationSettingsResponse) GetNotifications() []*UserNotificationSettings {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type UserNotificationSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configurable list of notifications enabled by the user
	FirstRequest bool `protobuf:"varint,1,opt,name=first_request,json=firstRequest,proto3" json:"first_request,omitempty"`
	Approve      bool `protobuf:"varint,2,opt,name=approve,proto3" json:"approve,omitempty"`
	Activate     bool `protobuf:"varint,3,opt,name=activate,proto3" json:"activate,omitempty"`
	Extend       bool `protobuf:"varint,4,opt,name=extend,proto3" json:"extend,omitempty"`
	Revoke       bool `protobuf:"varint,5,opt,name=revoke,proto3" json:"revoke,omitempty"`
	Deprovision  bool `protobuf:"varint,6,opt,name=deprovision,proto3" json:"deprovision,omitempty"`
}

func (x *UserNotificationSettings) Reset() {
	*x = UserNotificationSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNotificationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNotificationSettings) ProtoMessage() {}

func (x *UserNotificationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNotificationSettings.ProtoReflect.Descriptor instead.
func (*UserNotificationSettings) Descriptor() ([]byte, []int) {
	return file_commonfate_control_notification_v1alpha1_notification_proto_rawDescGZIP(), []int{2}
}

func (x *UserNotificationSettings) GetFirstRequest() bool {
	if x != nil {
		return x.FirstRequest
	}
	return false
}

func (x *UserNotificationSettings) GetApprove() bool {
	if x != nil {
		return x.Approve
	}
	return false
}

func (x *UserNotificationSettings) GetActivate() bool {
	if x != nil {
		return x.Activate
	}
	return false
}

func (x *UserNotificationSettings) GetExtend() bool {
	if x != nil {
		return x.Extend
	}
	return false
}

func (x *UserNotificationSettings) GetRevoke() bool {
	if x != nil {
		return x.Revoke
	}
	return false
}

func (x *UserNotificationSettings) GetDeprovision() bool {
	if x != nil {
		return x.Deprovision
	}
	return false
}

type UpdateUserNotificationSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *UpdateUserNotificationSettingsRequest) Reset() {
	*x = UpdateUserNotificationSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserNotificationSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserNotificationSettingsRequest) ProtoMessage() {}

func (x *UpdateUserNotificationSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserNotificationSettingsRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserNotificationSettingsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_notification_v1alpha1_notification_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserNotificationSettingsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserNotificationSettingsRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type UpdateUserNotificationSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification *UserNotificationSettings `protobuf:"bytes,1,opt,name=notification,proto3" json:"notification,omitempty"`
}

func (x *UpdateUserNotificationSettingsResponse) Reset() {
	*x = UpdateUserNotificationSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserNotificationSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserNotificationSettingsResponse) ProtoMessage() {}

func (x *UpdateUserNotificationSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserNotificationSettingsResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserNotificationSettingsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_notification_v1alpha1_notification_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserNotificationSettingsResponse) GetNotification() *UserNotificationSettings {
	if x != nil {
		return x.Notification
	}
	return nil
}

var File_commonfate_control_notification_v1alpha1_notification_proto protoreflect.FileDescriptor

var file_commonfate_control_notification_v1alpha1_notification_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x28, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x93, 0x01, 0x0a,
	0x27, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x0d, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x25,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x26, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66,
	0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb8, 0x03, 0x0a, 0x1f, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xcc, 0x01, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x50,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x51, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0xc5, 0x01, 0x0a, 0x1e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x50, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xe2, 0x02, 0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x4e, 0xaa, 0x02, 0x28, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x28, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xe2, 0x02, 0x34, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_notification_v1alpha1_notification_proto_rawDescOnce sync.Once
	file_commonfate_control_notification_v1alpha1_notification_proto_rawDescData = file_commonfate_control_notification_v1alpha1_notification_proto_rawDesc
)

func file_commonfate_control_notification_v1alpha1_notification_proto_rawDescGZIP() []byte {
	file_commonfate_control_notification_v1alpha1_notification_proto_rawDescOnce.Do(func() {
		file_commonfate_control_notification_v1alpha1_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_notification_v1alpha1_notification_proto_rawDescData)
	})
	return file_commonfate_control_notification_v1alpha1_notification_proto_rawDescData
}

var file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_commonfate_control_notification_v1alpha1_notification_proto_goTypes = []any{
	(*GetUserUserNotificationSettingsRequest)(nil),  // 0: commonfate.control.notification.v1alpha1.GetUserUserNotificationSettingsRequest
	(*GetUserUserNotificationSettingsResponse)(nil), // 1: commonfate.control.notification.v1alpha1.GetUserUserNotificationSettingsResponse
	(*UserNotificationSettings)(nil),                // 2: commonfate.control.notification.v1alpha1.UserNotificationSettings
	(*UpdateUserNotificationSettingsRequest)(nil),   // 3: commonfate.control.notification.v1alpha1.UpdateUserNotificationSettingsRequest
	(*UpdateUserNotificationSettingsResponse)(nil),  // 4: commonfate.control.notification.v1alpha1.UpdateUserNotificationSettingsResponse
}
var file_commonfate_control_notification_v1alpha1_notification_proto_depIdxs = []int32{
	2, // 0: commonfate.control.notification.v1alpha1.GetUserUserNotificationSettingsResponse.notifications:type_name -> commonfate.control.notification.v1alpha1.UserNotificationSettings
	2, // 1: commonfate.control.notification.v1alpha1.UpdateUserNotificationSettingsResponse.notification:type_name -> commonfate.control.notification.v1alpha1.UserNotificationSettings
	0, // 2: commonfate.control.notification.v1alpha1.UserNotificationSettingsService.GetUserUserNotificationSettings:input_type -> commonfate.control.notification.v1alpha1.GetUserUserNotificationSettingsRequest
	3, // 3: commonfate.control.notification.v1alpha1.UserNotificationSettingsService.UpdateUserNotificationSettings:input_type -> commonfate.control.notification.v1alpha1.UpdateUserNotificationSettingsRequest
	1, // 4: commonfate.control.notification.v1alpha1.UserNotificationSettingsService.GetUserUserNotificationSettings:output_type -> commonfate.control.notification.v1alpha1.GetUserUserNotificationSettingsResponse
	4, // 5: commonfate.control.notification.v1alpha1.UserNotificationSettingsService.UpdateUserNotificationSettings:output_type -> commonfate.control.notification.v1alpha1.UpdateUserNotificationSettingsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commonfate_control_notification_v1alpha1_notification_proto_init() }
func file_commonfate_control_notification_v1alpha1_notification_proto_init() {
	if File_commonfate_control_notification_v1alpha1_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserUserNotificationSettingsRequest); i {
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
		file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserUserNotificationSettingsResponse); i {
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
		file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UserNotificationSettings); i {
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
		file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserNotificationSettingsRequest); i {
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
		file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateUserNotificationSettingsResponse); i {
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
			RawDescriptor: file_commonfate_control_notification_v1alpha1_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_notification_v1alpha1_notification_proto_goTypes,
		DependencyIndexes: file_commonfate_control_notification_v1alpha1_notification_proto_depIdxs,
		MessageInfos:      file_commonfate_control_notification_v1alpha1_notification_proto_msgTypes,
	}.Build()
	File_commonfate_control_notification_v1alpha1_notification_proto = out.File
	file_commonfate_control_notification_v1alpha1_notification_proto_rawDesc = nil
	file_commonfate_control_notification_v1alpha1_notification_proto_goTypes = nil
	file_commonfate_control_notification_v1alpha1_notification_proto_depIdxs = nil
}
