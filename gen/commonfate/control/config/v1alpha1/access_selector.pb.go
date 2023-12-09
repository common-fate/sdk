// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commonfate/control/config/v1alpha1/access_selector.proto

package configv1alpha1

import (
	v1alpha11 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
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

type CreateAccessSelectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ResourceType string        `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	BelongingTo  *v1alpha1.UID `protobuf:"bytes,3,opt,name=belonging_to,json=belongingTo,proto3" json:"belonging_to,omitempty"`
	When         string        `protobuf:"bytes,4,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *CreateAccessSelectorRequest) Reset() {
	*x = CreateAccessSelectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccessSelectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessSelectorRequest) ProtoMessage() {}

func (x *CreateAccessSelectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessSelectorRequest.ProtoReflect.Descriptor instead.
func (*CreateAccessSelectorRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccessSelectorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAccessSelectorRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *CreateAccessSelectorRequest) GetBelongingTo() *v1alpha1.UID {
	if x != nil {
		return x.BelongingTo
	}
	return nil
}

func (x *CreateAccessSelectorRequest) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

type AccessSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ResourceType string        `protobuf:"bytes,3,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	BelongingTo  *v1alpha1.UID `protobuf:"bytes,4,opt,name=belonging_to,json=belongingTo,proto3" json:"belonging_to,omitempty"`
	When         string        `protobuf:"bytes,5,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *AccessSelector) Reset() {
	*x = AccessSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessSelector) ProtoMessage() {}

func (x *AccessSelector) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessSelector.ProtoReflect.Descriptor instead.
func (*AccessSelector) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP(), []int{1}
}

func (x *AccessSelector) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccessSelector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccessSelector) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *AccessSelector) GetBelongingTo() *v1alpha1.UID {
	if x != nil {
		return x.BelongingTo
	}
	return nil
}

func (x *AccessSelector) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

type CreateAccessSelectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selector    *AccessSelector         `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	Diagnostics []*v1alpha11.Diagnostic `protobuf:"bytes,2,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *CreateAccessSelectorResponse) Reset() {
	*x = CreateAccessSelectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccessSelectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessSelectorResponse) ProtoMessage() {}

func (x *CreateAccessSelectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessSelectorResponse.ProtoReflect.Descriptor instead.
func (*CreateAccessSelectorResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccessSelectorResponse) GetSelector() *AccessSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *CreateAccessSelectorResponse) GetDiagnostics() []*v1alpha11.Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

type GetAccessSelectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAccessSelectorRequest) Reset() {
	*x = GetAccessSelectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessSelectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessSelectorRequest) ProtoMessage() {}

func (x *GetAccessSelectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessSelectorRequest.ProtoReflect.Descriptor instead.
func (*GetAccessSelectorRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccessSelectorRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAccessSelectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selector *AccessSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
}

func (x *GetAccessSelectorResponse) Reset() {
	*x = GetAccessSelectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessSelectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessSelectorResponse) ProtoMessage() {}

func (x *GetAccessSelectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessSelectorResponse.ProtoReflect.Descriptor instead.
func (*GetAccessSelectorResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccessSelectorResponse) GetSelector() *AccessSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

type UpdateAccessSelectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selector *AccessSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
}

func (x *UpdateAccessSelectorRequest) Reset() {
	*x = UpdateAccessSelectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccessSelectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccessSelectorRequest) ProtoMessage() {}

func (x *UpdateAccessSelectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccessSelectorRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccessSelectorRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAccessSelectorRequest) GetSelector() *AccessSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

type UpdateAccessSelectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selector    *AccessSelector         `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	Diagnostics []*v1alpha11.Diagnostic `protobuf:"bytes,2,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *UpdateAccessSelectorResponse) Reset() {
	*x = UpdateAccessSelectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccessSelectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccessSelectorResponse) ProtoMessage() {}

func (x *UpdateAccessSelectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccessSelectorResponse.ProtoReflect.Descriptor instead.
func (*UpdateAccessSelectorResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAccessSelectorResponse) GetSelector() *AccessSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *UpdateAccessSelectorResponse) GetDiagnostics() []*v1alpha11.Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

type DeleteAccessSelectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAccessSelectorRequest) Reset() {
	*x = DeleteAccessSelectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccessSelectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessSelectorRequest) ProtoMessage() {}

func (x *DeleteAccessSelectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessSelectorRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccessSelectorRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAccessSelectorRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAccessSelectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAccessSelectorResponse) Reset() {
	*x = DeleteAccessSelectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccessSelectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessSelectorResponse) ProtoMessage() {}

func (x *DeleteAccessSelectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessSelectorResponse.ProtoReflect.Descriptor instead.
func (*DeleteAccessSelectorResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAccessSelectorResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_commonfate_control_config_v1alpha1_access_selector_proto protoreflect.FileDescriptor

var file_commonfate_control_config_v1alpha1_access_selector_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x24,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x75, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xae, 0x01, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x62, 0x65,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x49,
	0x44, 0x52, 0x0b, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x68,
	0x65, 0x6e, 0x22, 0xb1, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42,
	0x0a, 0x0c, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x55, 0x49, 0x44, 0x52, 0x0b, 0x62, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x6e, 0x67,
	0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x22, 0xb8, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6b, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x08, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x6d, 0x0a, 0x1b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x08, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xb8, 0x01, 0x0a, 0x1c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x08, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x64, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x61,
	0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x22, 0x2d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0x86, 0x05, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9b, 0x01,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x92, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x9b, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9b,
	0x01, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xba, 0x02, 0x0a,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x43, 0x43, 0xaa, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x22, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescOnce sync.Once
	file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescData = file_commonfate_control_config_v1alpha1_access_selector_proto_rawDesc
)

func file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescGZIP() []byte {
	file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescOnce.Do(func() {
		file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescData)
	})
	return file_commonfate_control_config_v1alpha1_access_selector_proto_rawDescData
}

var file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_commonfate_control_config_v1alpha1_access_selector_proto_goTypes = []interface{}{
	(*CreateAccessSelectorRequest)(nil),  // 0: commonfate.control.config.v1alpha1.CreateAccessSelectorRequest
	(*AccessSelector)(nil),               // 1: commonfate.control.config.v1alpha1.AccessSelector
	(*CreateAccessSelectorResponse)(nil), // 2: commonfate.control.config.v1alpha1.CreateAccessSelectorResponse
	(*GetAccessSelectorRequest)(nil),     // 3: commonfate.control.config.v1alpha1.GetAccessSelectorRequest
	(*GetAccessSelectorResponse)(nil),    // 4: commonfate.control.config.v1alpha1.GetAccessSelectorResponse
	(*UpdateAccessSelectorRequest)(nil),  // 5: commonfate.control.config.v1alpha1.UpdateAccessSelectorRequest
	(*UpdateAccessSelectorResponse)(nil), // 6: commonfate.control.config.v1alpha1.UpdateAccessSelectorResponse
	(*DeleteAccessSelectorRequest)(nil),  // 7: commonfate.control.config.v1alpha1.DeleteAccessSelectorRequest
	(*DeleteAccessSelectorResponse)(nil), // 8: commonfate.control.config.v1alpha1.DeleteAccessSelectorResponse
	(*v1alpha1.UID)(nil),                 // 9: commonfate.entity.v1alpha1.UID
	(*v1alpha11.Diagnostic)(nil),         // 10: commonfate.access.v1alpha1.Diagnostic
}
var file_commonfate_control_config_v1alpha1_access_selector_proto_depIdxs = []int32{
	9,  // 0: commonfate.control.config.v1alpha1.CreateAccessSelectorRequest.belonging_to:type_name -> commonfate.entity.v1alpha1.UID
	9,  // 1: commonfate.control.config.v1alpha1.AccessSelector.belonging_to:type_name -> commonfate.entity.v1alpha1.UID
	1,  // 2: commonfate.control.config.v1alpha1.CreateAccessSelectorResponse.selector:type_name -> commonfate.control.config.v1alpha1.AccessSelector
	10, // 3: commonfate.control.config.v1alpha1.CreateAccessSelectorResponse.diagnostics:type_name -> commonfate.access.v1alpha1.Diagnostic
	1,  // 4: commonfate.control.config.v1alpha1.GetAccessSelectorResponse.selector:type_name -> commonfate.control.config.v1alpha1.AccessSelector
	1,  // 5: commonfate.control.config.v1alpha1.UpdateAccessSelectorRequest.selector:type_name -> commonfate.control.config.v1alpha1.AccessSelector
	1,  // 6: commonfate.control.config.v1alpha1.UpdateAccessSelectorResponse.selector:type_name -> commonfate.control.config.v1alpha1.AccessSelector
	10, // 7: commonfate.control.config.v1alpha1.UpdateAccessSelectorResponse.diagnostics:type_name -> commonfate.access.v1alpha1.Diagnostic
	0,  // 8: commonfate.control.config.v1alpha1.AccessSelectorService.CreateAccessSelector:input_type -> commonfate.control.config.v1alpha1.CreateAccessSelectorRequest
	3,  // 9: commonfate.control.config.v1alpha1.AccessSelectorService.GetAccessSelector:input_type -> commonfate.control.config.v1alpha1.GetAccessSelectorRequest
	5,  // 10: commonfate.control.config.v1alpha1.AccessSelectorService.UpdateAccessSelector:input_type -> commonfate.control.config.v1alpha1.UpdateAccessSelectorRequest
	7,  // 11: commonfate.control.config.v1alpha1.AccessSelectorService.DeleteAccessSelector:input_type -> commonfate.control.config.v1alpha1.DeleteAccessSelectorRequest
	2,  // 12: commonfate.control.config.v1alpha1.AccessSelectorService.CreateAccessSelector:output_type -> commonfate.control.config.v1alpha1.CreateAccessSelectorResponse
	4,  // 13: commonfate.control.config.v1alpha1.AccessSelectorService.GetAccessSelector:output_type -> commonfate.control.config.v1alpha1.GetAccessSelectorResponse
	6,  // 14: commonfate.control.config.v1alpha1.AccessSelectorService.UpdateAccessSelector:output_type -> commonfate.control.config.v1alpha1.UpdateAccessSelectorResponse
	8,  // 15: commonfate.control.config.v1alpha1.AccessSelectorService.DeleteAccessSelector:output_type -> commonfate.control.config.v1alpha1.DeleteAccessSelectorResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_commonfate_control_config_v1alpha1_access_selector_proto_init() }
func file_commonfate_control_config_v1alpha1_access_selector_proto_init() {
	if File_commonfate_control_config_v1alpha1_access_selector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccessSelectorRequest); i {
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
		file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessSelector); i {
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
		file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccessSelectorResponse); i {
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
		file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessSelectorRequest); i {
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
		file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessSelectorResponse); i {
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
		file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccessSelectorRequest); i {
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
		file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccessSelectorResponse); i {
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
		file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccessSelectorRequest); i {
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
		file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccessSelectorResponse); i {
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
			RawDescriptor: file_commonfate_control_config_v1alpha1_access_selector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_config_v1alpha1_access_selector_proto_goTypes,
		DependencyIndexes: file_commonfate_control_config_v1alpha1_access_selector_proto_depIdxs,
		MessageInfos:      file_commonfate_control_config_v1alpha1_access_selector_proto_msgTypes,
	}.Build()
	File_commonfate_control_config_v1alpha1_access_selector_proto = out.File
	file_commonfate_control_config_v1alpha1_access_selector_proto_rawDesc = nil
	file_commonfate_control_config_v1alpha1_access_selector_proto_goTypes = nil
	file_commonfate_control_config_v1alpha1_access_selector_proto_depIdxs = nil
}
