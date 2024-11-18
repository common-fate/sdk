// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: commonfate/control/config/v1alpha1/availability_spec.proto

package configv1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AvailabilitySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkflowId     string        `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	Role           *v1alpha1.EID `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Target         *v1alpha1.EID `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	IdentityDomain *v1alpha1.EID `protobuf:"bytes,6,opt,name=identity_domain,json=identityDomain,proto3" json:"identity_domain,omitempty"`
	RolePriority   *int64        `protobuf:"varint,7,opt,name=role_priority,json=rolePriority,proto3,oneof" json:"role_priority,omitempty"`
}

func (x *AvailabilitySpec) Reset() {
	*x = AvailabilitySpec{}
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AvailabilitySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailabilitySpec) ProtoMessage() {}

func (x *AvailabilitySpec) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailabilitySpec.ProtoReflect.Descriptor instead.
func (*AvailabilitySpec) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP(), []int{0}
}

func (x *AvailabilitySpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AvailabilitySpec) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *AvailabilitySpec) GetRole() *v1alpha1.EID {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *AvailabilitySpec) GetTarget() *v1alpha1.EID {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *AvailabilitySpec) GetIdentityDomain() *v1alpha1.EID {
	if x != nil {
		return x.IdentityDomain
	}
	return nil
}

func (x *AvailabilitySpec) GetRolePriority() int64 {
	if x != nil && x.RolePriority != nil {
		return *x.RolePriority
	}
	return 0
}

type CreateAvailabilitySpecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId     string        `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	Role           *v1alpha1.EID `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Target         *v1alpha1.EID `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	IdentityDomain *v1alpha1.EID `protobuf:"bytes,4,opt,name=identity_domain,json=identityDomain,proto3" json:"identity_domain,omitempty"`
	RolePriority   *int64        `protobuf:"varint,5,opt,name=role_priority,json=rolePriority,proto3,oneof" json:"role_priority,omitempty"`
}

func (x *CreateAvailabilitySpecRequest) Reset() {
	*x = CreateAvailabilitySpecRequest{}
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAvailabilitySpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAvailabilitySpecRequest) ProtoMessage() {}

func (x *CreateAvailabilitySpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAvailabilitySpecRequest.ProtoReflect.Descriptor instead.
func (*CreateAvailabilitySpecRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAvailabilitySpecRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *CreateAvailabilitySpecRequest) GetRole() *v1alpha1.EID {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *CreateAvailabilitySpecRequest) GetTarget() *v1alpha1.EID {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *CreateAvailabilitySpecRequest) GetIdentityDomain() *v1alpha1.EID {
	if x != nil {
		return x.IdentityDomain
	}
	return nil
}

func (x *CreateAvailabilitySpecRequest) GetRolePriority() int64 {
	if x != nil && x.RolePriority != nil {
		return *x.RolePriority
	}
	return 0
}

type CreateAvailabilitySpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailabilitySpec *AvailabilitySpec `protobuf:"bytes,1,opt,name=availability_spec,json=availabilitySpec,proto3" json:"availability_spec,omitempty"`
}

func (x *CreateAvailabilitySpecResponse) Reset() {
	*x = CreateAvailabilitySpecResponse{}
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAvailabilitySpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAvailabilitySpecResponse) ProtoMessage() {}

func (x *CreateAvailabilitySpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAvailabilitySpecResponse.ProtoReflect.Descriptor instead.
func (*CreateAvailabilitySpecResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAvailabilitySpecResponse) GetAvailabilitySpec() *AvailabilitySpec {
	if x != nil {
		return x.AvailabilitySpec
	}
	return nil
}

type GetAvailabilitySpecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAvailabilitySpecRequest) Reset() {
	*x = GetAvailabilitySpecRequest{}
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAvailabilitySpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailabilitySpecRequest) ProtoMessage() {}

func (x *GetAvailabilitySpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailabilitySpecRequest.ProtoReflect.Descriptor instead.
func (*GetAvailabilitySpecRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP(), []int{3}
}

func (x *GetAvailabilitySpecRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAvailabilitySpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailabilitySpec *AvailabilitySpec `protobuf:"bytes,1,opt,name=availability_spec,json=availabilitySpec,proto3" json:"availability_spec,omitempty"`
}

func (x *GetAvailabilitySpecResponse) Reset() {
	*x = GetAvailabilitySpecResponse{}
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAvailabilitySpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvailabilitySpecResponse) ProtoMessage() {}

func (x *GetAvailabilitySpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvailabilitySpecResponse.ProtoReflect.Descriptor instead.
func (*GetAvailabilitySpecResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP(), []int{4}
}

func (x *GetAvailabilitySpecResponse) GetAvailabilitySpec() *AvailabilitySpec {
	if x != nil {
		return x.AvailabilitySpec
	}
	return nil
}

type UpdateAvailabilitySpecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailabilitySpec *AvailabilitySpec `protobuf:"bytes,1,opt,name=availability_spec,json=availabilitySpec,proto3" json:"availability_spec,omitempty"`
}

func (x *UpdateAvailabilitySpecRequest) Reset() {
	*x = UpdateAvailabilitySpecRequest{}
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAvailabilitySpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvailabilitySpecRequest) ProtoMessage() {}

func (x *UpdateAvailabilitySpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvailabilitySpecRequest.ProtoReflect.Descriptor instead.
func (*UpdateAvailabilitySpecRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAvailabilitySpecRequest) GetAvailabilitySpec() *AvailabilitySpec {
	if x != nil {
		return x.AvailabilitySpec
	}
	return nil
}

type UpdateAvailabilitySpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailabilitySpec *AvailabilitySpec `protobuf:"bytes,1,opt,name=availability_spec,json=availabilitySpec,proto3" json:"availability_spec,omitempty"`
}

func (x *UpdateAvailabilitySpecResponse) Reset() {
	*x = UpdateAvailabilitySpecResponse{}
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAvailabilitySpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvailabilitySpecResponse) ProtoMessage() {}

func (x *UpdateAvailabilitySpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvailabilitySpecResponse.ProtoReflect.Descriptor instead.
func (*UpdateAvailabilitySpecResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAvailabilitySpecResponse) GetAvailabilitySpec() *AvailabilitySpec {
	if x != nil {
		return x.AvailabilitySpec
	}
	return nil
}

type DeleteAvailabilitySpecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAvailabilitySpecRequest) Reset() {
	*x = DeleteAvailabilitySpecRequest{}
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAvailabilitySpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAvailabilitySpecRequest) ProtoMessage() {}

func (x *DeleteAvailabilitySpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAvailabilitySpecRequest.ProtoReflect.Descriptor instead.
func (*DeleteAvailabilitySpecRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAvailabilitySpecRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAvailabilitySpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAvailabilitySpecResponse) Reset() {
	*x = DeleteAvailabilitySpecResponse{}
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAvailabilitySpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAvailabilitySpecResponse) ProtoMessage() {}

func (x *DeleteAvailabilitySpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAvailabilitySpecResponse.ProtoReflect.Descriptor instead.
func (*DeleteAvailabilitySpecResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAvailabilitySpecResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_commonfate_control_config_v1alpha1_availability_spec_proto protoreflect.FileDescriptor

var file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb7, 0x02, 0x0a, 0x10, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x48, 0x0a, 0x0f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52,
	0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x28, 0x0a, 0x0d, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0xb4, 0x02, 0x0a, 0x1d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x33,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x49, 0x44, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x48, 0x0a, 0x0f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x0c, 0x72, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x22, 0x83, 0x01, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x22, 0x2c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x22, 0x82, 0x01, 0x0a, 0x1d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x61, 0x0a, 0x11, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x10, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x22, 0x83,
	0x01, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x61, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x53, 0x70, 0x65, 0x63, 0x22, 0x2f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa4, 0x05, 0x0a, 0x17, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x41,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9c, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0xa1, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa1, 0x01, 0x0a, 0x16, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xbc,
	0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x15, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x43, 0xaa, 0x02, 0x22, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescOnce sync.Once
	file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescData = file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDesc
)

func file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescGZIP() []byte {
	file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescOnce.Do(func() {
		file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescData)
	})
	return file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDescData
}

var file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_commonfate_control_config_v1alpha1_availability_spec_proto_goTypes = []any{
	(*AvailabilitySpec)(nil),               // 0: commonfate.control.config.v1alpha1.AvailabilitySpec
	(*CreateAvailabilitySpecRequest)(nil),  // 1: commonfate.control.config.v1alpha1.CreateAvailabilitySpecRequest
	(*CreateAvailabilitySpecResponse)(nil), // 2: commonfate.control.config.v1alpha1.CreateAvailabilitySpecResponse
	(*GetAvailabilitySpecRequest)(nil),     // 3: commonfate.control.config.v1alpha1.GetAvailabilitySpecRequest
	(*GetAvailabilitySpecResponse)(nil),    // 4: commonfate.control.config.v1alpha1.GetAvailabilitySpecResponse
	(*UpdateAvailabilitySpecRequest)(nil),  // 5: commonfate.control.config.v1alpha1.UpdateAvailabilitySpecRequest
	(*UpdateAvailabilitySpecResponse)(nil), // 6: commonfate.control.config.v1alpha1.UpdateAvailabilitySpecResponse
	(*DeleteAvailabilitySpecRequest)(nil),  // 7: commonfate.control.config.v1alpha1.DeleteAvailabilitySpecRequest
	(*DeleteAvailabilitySpecResponse)(nil), // 8: commonfate.control.config.v1alpha1.DeleteAvailabilitySpecResponse
	(*v1alpha1.EID)(nil),                   // 9: commonfate.entity.v1alpha1.EID
}
var file_commonfate_control_config_v1alpha1_availability_spec_proto_depIdxs = []int32{
	9,  // 0: commonfate.control.config.v1alpha1.AvailabilitySpec.role:type_name -> commonfate.entity.v1alpha1.EID
	9,  // 1: commonfate.control.config.v1alpha1.AvailabilitySpec.target:type_name -> commonfate.entity.v1alpha1.EID
	9,  // 2: commonfate.control.config.v1alpha1.AvailabilitySpec.identity_domain:type_name -> commonfate.entity.v1alpha1.EID
	9,  // 3: commonfate.control.config.v1alpha1.CreateAvailabilitySpecRequest.role:type_name -> commonfate.entity.v1alpha1.EID
	9,  // 4: commonfate.control.config.v1alpha1.CreateAvailabilitySpecRequest.target:type_name -> commonfate.entity.v1alpha1.EID
	9,  // 5: commonfate.control.config.v1alpha1.CreateAvailabilitySpecRequest.identity_domain:type_name -> commonfate.entity.v1alpha1.EID
	0,  // 6: commonfate.control.config.v1alpha1.CreateAvailabilitySpecResponse.availability_spec:type_name -> commonfate.control.config.v1alpha1.AvailabilitySpec
	0,  // 7: commonfate.control.config.v1alpha1.GetAvailabilitySpecResponse.availability_spec:type_name -> commonfate.control.config.v1alpha1.AvailabilitySpec
	0,  // 8: commonfate.control.config.v1alpha1.UpdateAvailabilitySpecRequest.availability_spec:type_name -> commonfate.control.config.v1alpha1.AvailabilitySpec
	0,  // 9: commonfate.control.config.v1alpha1.UpdateAvailabilitySpecResponse.availability_spec:type_name -> commonfate.control.config.v1alpha1.AvailabilitySpec
	1,  // 10: commonfate.control.config.v1alpha1.AvailabilitySpecService.CreateAvailabilitySpec:input_type -> commonfate.control.config.v1alpha1.CreateAvailabilitySpecRequest
	3,  // 11: commonfate.control.config.v1alpha1.AvailabilitySpecService.GetAvailabilitySpec:input_type -> commonfate.control.config.v1alpha1.GetAvailabilitySpecRequest
	5,  // 12: commonfate.control.config.v1alpha1.AvailabilitySpecService.UpdateAvailabilitySpec:input_type -> commonfate.control.config.v1alpha1.UpdateAvailabilitySpecRequest
	7,  // 13: commonfate.control.config.v1alpha1.AvailabilitySpecService.DeleteAvailabilitySpec:input_type -> commonfate.control.config.v1alpha1.DeleteAvailabilitySpecRequest
	2,  // 14: commonfate.control.config.v1alpha1.AvailabilitySpecService.CreateAvailabilitySpec:output_type -> commonfate.control.config.v1alpha1.CreateAvailabilitySpecResponse
	4,  // 15: commonfate.control.config.v1alpha1.AvailabilitySpecService.GetAvailabilitySpec:output_type -> commonfate.control.config.v1alpha1.GetAvailabilitySpecResponse
	6,  // 16: commonfate.control.config.v1alpha1.AvailabilitySpecService.UpdateAvailabilitySpec:output_type -> commonfate.control.config.v1alpha1.UpdateAvailabilitySpecResponse
	8,  // 17: commonfate.control.config.v1alpha1.AvailabilitySpecService.DeleteAvailabilitySpec:output_type -> commonfate.control.config.v1alpha1.DeleteAvailabilitySpecResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_commonfate_control_config_v1alpha1_availability_spec_proto_init() }
func file_commonfate_control_config_v1alpha1_availability_spec_proto_init() {
	if File_commonfate_control_config_v1alpha1_availability_spec_proto != nil {
		return
	}
	file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[0].OneofWrappers = []any{}
	file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_config_v1alpha1_availability_spec_proto_goTypes,
		DependencyIndexes: file_commonfate_control_config_v1alpha1_availability_spec_proto_depIdxs,
		MessageInfos:      file_commonfate_control_config_v1alpha1_availability_spec_proto_msgTypes,
	}.Build()
	File_commonfate_control_config_v1alpha1_availability_spec_proto = out.File
	file_commonfate_control_config_v1alpha1_availability_spec_proto_rawDesc = nil
	file_commonfate_control_config_v1alpha1_availability_spec_proto_goTypes = nil
	file_commonfate_control_config_v1alpha1_availability_spec_proto_depIdxs = nil
}
