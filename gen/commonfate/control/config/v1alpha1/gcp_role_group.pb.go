// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: commonfate/control/config/v1alpha1/gcp_role_group.proto

package configv1alpha1

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

type CreateGCPRoleGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrganizationId string   `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	RoleIds        []string `protobuf:"bytes,3,rep,name=role_ids,json=roleIds,proto3" json:"role_ids,omitempty"`
}

func (x *CreateGCPRoleGroupRequest) Reset() {
	*x = CreateGCPRoleGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGCPRoleGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGCPRoleGroupRequest) ProtoMessage() {}

func (x *CreateGCPRoleGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGCPRoleGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGCPRoleGroupRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGCPRoleGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateGCPRoleGroupRequest) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *CreateGCPRoleGroupRequest) GetRoleIds() []string {
	if x != nil {
		return x.RoleIds
	}
	return nil
}

type GCPRoleGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OrganizationId string   `protobuf:"bytes,3,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	RoleIds        []string `protobuf:"bytes,4,rep,name=role_ids,json=roleIds,proto3" json:"role_ids,omitempty"`
}

func (x *GCPRoleGroup) Reset() {
	*x = GCPRoleGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCPRoleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCPRoleGroup) ProtoMessage() {}

func (x *GCPRoleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCPRoleGroup.ProtoReflect.Descriptor instead.
func (*GCPRoleGroup) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP(), []int{1}
}

func (x *GCPRoleGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GCPRoleGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GCPRoleGroup) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *GCPRoleGroup) GetRoleIds() []string {
	if x != nil {
		return x.RoleIds
	}
	return nil
}

type CreateGCPRoleGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleGroup *GCPRoleGroup `protobuf:"bytes,1,opt,name=role_group,json=roleGroup,proto3" json:"role_group,omitempty"`
}

func (x *CreateGCPRoleGroupResponse) Reset() {
	*x = CreateGCPRoleGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGCPRoleGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGCPRoleGroupResponse) ProtoMessage() {}

func (x *CreateGCPRoleGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGCPRoleGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateGCPRoleGroupResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP(), []int{2}
}

func (x *CreateGCPRoleGroupResponse) GetRoleGroup() *GCPRoleGroup {
	if x != nil {
		return x.RoleGroup
	}
	return nil
}

type GetGCPRoleGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGCPRoleGroupRequest) Reset() {
	*x = GetGCPRoleGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGCPRoleGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGCPRoleGroupRequest) ProtoMessage() {}

func (x *GetGCPRoleGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGCPRoleGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGCPRoleGroupRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP(), []int{3}
}

func (x *GetGCPRoleGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGCPRoleGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleGroup *GCPRoleGroup `protobuf:"bytes,1,opt,name=role_group,json=roleGroup,proto3" json:"role_group,omitempty"`
}

func (x *GetGCPRoleGroupResponse) Reset() {
	*x = GetGCPRoleGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGCPRoleGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGCPRoleGroupResponse) ProtoMessage() {}

func (x *GetGCPRoleGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGCPRoleGroupResponse.ProtoReflect.Descriptor instead.
func (*GetGCPRoleGroupResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP(), []int{4}
}

func (x *GetGCPRoleGroupResponse) GetRoleGroup() *GCPRoleGroup {
	if x != nil {
		return x.RoleGroup
	}
	return nil
}

type UpdateGCPRoleGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleGroup *GCPRoleGroup `protobuf:"bytes,1,opt,name=role_group,json=roleGroup,proto3" json:"role_group,omitempty"`
}

func (x *UpdateGCPRoleGroupRequest) Reset() {
	*x = UpdateGCPRoleGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGCPRoleGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGCPRoleGroupRequest) ProtoMessage() {}

func (x *UpdateGCPRoleGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGCPRoleGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGCPRoleGroupRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateGCPRoleGroupRequest) GetRoleGroup() *GCPRoleGroup {
	if x != nil {
		return x.RoleGroup
	}
	return nil
}

type UpdateGCPRoleGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleGroup *GCPRoleGroup `protobuf:"bytes,1,opt,name=role_group,json=roleGroup,proto3" json:"role_group,omitempty"`
}

func (x *UpdateGCPRoleGroupResponse) Reset() {
	*x = UpdateGCPRoleGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGCPRoleGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGCPRoleGroupResponse) ProtoMessage() {}

func (x *UpdateGCPRoleGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGCPRoleGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateGCPRoleGroupResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateGCPRoleGroupResponse) GetRoleGroup() *GCPRoleGroup {
	if x != nil {
		return x.RoleGroup
	}
	return nil
}

type DeleteGCPRoleGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGCPRoleGroupRequest) Reset() {
	*x = DeleteGCPRoleGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGCPRoleGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGCPRoleGroupRequest) ProtoMessage() {}

func (x *DeleteGCPRoleGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGCPRoleGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGCPRoleGroupRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteGCPRoleGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteGCPRoleGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGCPRoleGroupResponse) Reset() {
	*x = DeleteGCPRoleGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGCPRoleGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGCPRoleGroupResponse) ProtoMessage() {}

func (x *DeleteGCPRoleGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGCPRoleGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteGCPRoleGroupResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteGCPRoleGroupResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_commonfate_control_config_v1alpha1_gcp_role_group_proto protoreflect.FileDescriptor

var file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x63, 0x70, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x29, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x76, 0x0a,
	0x0c, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x6d, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x43, 0x50,
	0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x47, 0x43, 0x50, 0x52, 0x6f,
	0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6a,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x72, 0x6f, 0x6c,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x6c, 0x0a, 0x19, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x72,
	0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x6d, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x09, 0x72, 0x6f,
	0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x43,
	0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xf0, 0x04, 0x0a, 0x13, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x43, 0x50,
	0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x43, 0x50, 0x52,
	0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x90, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x43, 0x50, 0x52, 0x6f,
	0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x95, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x95, 0x01,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x47, 0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47,
	0x43, 0x50, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb8, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x11, 0x47, 0x63, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x43, 0xaa, 0x02, 0x22,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescOnce sync.Once
	file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescData = file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDesc
)

func file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescGZIP() []byte {
	file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescOnce.Do(func() {
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescData)
	})
	return file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDescData
}

var file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_commonfate_control_config_v1alpha1_gcp_role_group_proto_goTypes = []interface{}{
	(*CreateGCPRoleGroupRequest)(nil),  // 0: commonfate.control.config.v1alpha1.CreateGCPRoleGroupRequest
	(*GCPRoleGroup)(nil),               // 1: commonfate.control.config.v1alpha1.GCPRoleGroup
	(*CreateGCPRoleGroupResponse)(nil), // 2: commonfate.control.config.v1alpha1.CreateGCPRoleGroupResponse
	(*GetGCPRoleGroupRequest)(nil),     // 3: commonfate.control.config.v1alpha1.GetGCPRoleGroupRequest
	(*GetGCPRoleGroupResponse)(nil),    // 4: commonfate.control.config.v1alpha1.GetGCPRoleGroupResponse
	(*UpdateGCPRoleGroupRequest)(nil),  // 5: commonfate.control.config.v1alpha1.UpdateGCPRoleGroupRequest
	(*UpdateGCPRoleGroupResponse)(nil), // 6: commonfate.control.config.v1alpha1.UpdateGCPRoleGroupResponse
	(*DeleteGCPRoleGroupRequest)(nil),  // 7: commonfate.control.config.v1alpha1.DeleteGCPRoleGroupRequest
	(*DeleteGCPRoleGroupResponse)(nil), // 8: commonfate.control.config.v1alpha1.DeleteGCPRoleGroupResponse
}
var file_commonfate_control_config_v1alpha1_gcp_role_group_proto_depIdxs = []int32{
	1, // 0: commonfate.control.config.v1alpha1.CreateGCPRoleGroupResponse.role_group:type_name -> commonfate.control.config.v1alpha1.GCPRoleGroup
	1, // 1: commonfate.control.config.v1alpha1.GetGCPRoleGroupResponse.role_group:type_name -> commonfate.control.config.v1alpha1.GCPRoleGroup
	1, // 2: commonfate.control.config.v1alpha1.UpdateGCPRoleGroupRequest.role_group:type_name -> commonfate.control.config.v1alpha1.GCPRoleGroup
	1, // 3: commonfate.control.config.v1alpha1.UpdateGCPRoleGroupResponse.role_group:type_name -> commonfate.control.config.v1alpha1.GCPRoleGroup
	0, // 4: commonfate.control.config.v1alpha1.GCPRoleGroupService.CreateGCPRoleGroup:input_type -> commonfate.control.config.v1alpha1.CreateGCPRoleGroupRequest
	3, // 5: commonfate.control.config.v1alpha1.GCPRoleGroupService.GetGCPRoleGroup:input_type -> commonfate.control.config.v1alpha1.GetGCPRoleGroupRequest
	5, // 6: commonfate.control.config.v1alpha1.GCPRoleGroupService.UpdateGCPRoleGroup:input_type -> commonfate.control.config.v1alpha1.UpdateGCPRoleGroupRequest
	7, // 7: commonfate.control.config.v1alpha1.GCPRoleGroupService.DeleteGCPRoleGroup:input_type -> commonfate.control.config.v1alpha1.DeleteGCPRoleGroupRequest
	2, // 8: commonfate.control.config.v1alpha1.GCPRoleGroupService.CreateGCPRoleGroup:output_type -> commonfate.control.config.v1alpha1.CreateGCPRoleGroupResponse
	4, // 9: commonfate.control.config.v1alpha1.GCPRoleGroupService.GetGCPRoleGroup:output_type -> commonfate.control.config.v1alpha1.GetGCPRoleGroupResponse
	6, // 10: commonfate.control.config.v1alpha1.GCPRoleGroupService.UpdateGCPRoleGroup:output_type -> commonfate.control.config.v1alpha1.UpdateGCPRoleGroupResponse
	8, // 11: commonfate.control.config.v1alpha1.GCPRoleGroupService.DeleteGCPRoleGroup:output_type -> commonfate.control.config.v1alpha1.DeleteGCPRoleGroupResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_commonfate_control_config_v1alpha1_gcp_role_group_proto_init() }
func file_commonfate_control_config_v1alpha1_gcp_role_group_proto_init() {
	if File_commonfate_control_config_v1alpha1_gcp_role_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGCPRoleGroupRequest); i {
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
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCPRoleGroup); i {
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
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGCPRoleGroupResponse); i {
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
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGCPRoleGroupRequest); i {
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
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGCPRoleGroupResponse); i {
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
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGCPRoleGroupRequest); i {
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
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGCPRoleGroupResponse); i {
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
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGCPRoleGroupRequest); i {
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
		file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGCPRoleGroupResponse); i {
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
			RawDescriptor: file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_config_v1alpha1_gcp_role_group_proto_goTypes,
		DependencyIndexes: file_commonfate_control_config_v1alpha1_gcp_role_group_proto_depIdxs,
		MessageInfos:      file_commonfate_control_config_v1alpha1_gcp_role_group_proto_msgTypes,
	}.Build()
	File_commonfate_control_config_v1alpha1_gcp_role_group_proto = out.File
	file_commonfate_control_config_v1alpha1_gcp_role_group_proto_rawDesc = nil
	file_commonfate_control_config_v1alpha1_gcp_role_group_proto_goTypes = nil
	file_commonfate_control_config_v1alpha1_gcp_role_group_proto_depIdxs = nil
}
