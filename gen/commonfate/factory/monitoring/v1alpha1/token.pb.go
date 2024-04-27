// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: commonfate/factory/monitoring/v1alpha1/token.proto

package monitoringv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateWriteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateWriteTokenRequest) Reset() {
	*x = CreateWriteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWriteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWriteTokenRequest) ProtoMessage() {}

func (x *CreateWriteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWriteTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateWriteTokenRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescGZIP(), []int{0}
}

type CreateWriteTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A write token which gives access to write OpenTelemetry data to our collector.
	WriteToken string `protobuf:"bytes,2,opt,name=write_token,json=writeToken,proto3" json:"write_token,omitempty"`
}

func (x *CreateWriteTokenResponse) Reset() {
	*x = CreateWriteTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWriteTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWriteTokenResponse) ProtoMessage() {}

func (x *CreateWriteTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWriteTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateWriteTokenResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWriteTokenResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateWriteTokenResponse) GetWriteToken() string {
	if x != nil {
		return x.WriteToken
	}
	return ""
}

type GetWriteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetWriteTokenRequest) Reset() {
	*x = GetWriteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWriteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWriteTokenRequest) ProtoMessage() {}

func (x *GetWriteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWriteTokenRequest.ProtoReflect.Descriptor instead.
func (*GetWriteTokenRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescGZIP(), []int{2}
}

func (x *GetWriteTokenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetWriteTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsValid bool   `protobuf:"varint,2,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
}

func (x *GetWriteTokenResponse) Reset() {
	*x = GetWriteTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWriteTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWriteTokenResponse) ProtoMessage() {}

func (x *GetWriteTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWriteTokenResponse.ProtoReflect.Descriptor instead.
func (*GetWriteTokenResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescGZIP(), []int{3}
}

func (x *GetWriteTokenResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetWriteTokenResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

type DeleteWriteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWriteTokenRequest) Reset() {
	*x = DeleteWriteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWriteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWriteTokenRequest) ProtoMessage() {}

func (x *DeleteWriteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWriteTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteWriteTokenRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteWriteTokenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteWriteTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletedId string `protobuf:"bytes,1,opt,name=deleted_id,json=deletedId,proto3" json:"deleted_id,omitempty"`
}

func (x *DeleteWriteTokenResponse) Reset() {
	*x = DeleteWriteTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWriteTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWriteTokenResponse) ProtoMessage() {}

func (x *DeleteWriteTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWriteTokenResponse.ProtoReflect.Descriptor instead.
func (*DeleteWriteTokenResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteWriteTokenResponse) GetDeletedId() string {
	if x != nil {
		return x.DeletedId
	}
	return ""
}

var File_commonfate_factory_monitoring_v1alpha1_token_proto protoreflect.FileDescriptor

var file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x18,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x49, 0x64, 0x32, 0xd3, 0x03, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x97, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x8e, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x97, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xcd, 0x02,
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x46, 0x4d, 0xaa, 0x02, 0x26, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xca, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x5c, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x32, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x29, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a,
	0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescOnce sync.Once
	file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescData = file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDesc
)

func file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescGZIP() []byte {
	file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescOnce.Do(func() {
		file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescData)
	})
	return file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDescData
}

var file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_commonfate_factory_monitoring_v1alpha1_token_proto_goTypes = []interface{}{
	(*CreateWriteTokenRequest)(nil),  // 0: commonfate.factory.monitoring.v1alpha1.CreateWriteTokenRequest
	(*CreateWriteTokenResponse)(nil), // 1: commonfate.factory.monitoring.v1alpha1.CreateWriteTokenResponse
	(*GetWriteTokenRequest)(nil),     // 2: commonfate.factory.monitoring.v1alpha1.GetWriteTokenRequest
	(*GetWriteTokenResponse)(nil),    // 3: commonfate.factory.monitoring.v1alpha1.GetWriteTokenResponse
	(*DeleteWriteTokenRequest)(nil),  // 4: commonfate.factory.monitoring.v1alpha1.DeleteWriteTokenRequest
	(*DeleteWriteTokenResponse)(nil), // 5: commonfate.factory.monitoring.v1alpha1.DeleteWriteTokenResponse
}
var file_commonfate_factory_monitoring_v1alpha1_token_proto_depIdxs = []int32{
	0, // 0: commonfate.factory.monitoring.v1alpha1.TokenService.CreateWriteToken:input_type -> commonfate.factory.monitoring.v1alpha1.CreateWriteTokenRequest
	2, // 1: commonfate.factory.monitoring.v1alpha1.TokenService.GetWriteToken:input_type -> commonfate.factory.monitoring.v1alpha1.GetWriteTokenRequest
	4, // 2: commonfate.factory.monitoring.v1alpha1.TokenService.DeleteWriteToken:input_type -> commonfate.factory.monitoring.v1alpha1.DeleteWriteTokenRequest
	1, // 3: commonfate.factory.monitoring.v1alpha1.TokenService.CreateWriteToken:output_type -> commonfate.factory.monitoring.v1alpha1.CreateWriteTokenResponse
	3, // 4: commonfate.factory.monitoring.v1alpha1.TokenService.GetWriteToken:output_type -> commonfate.factory.monitoring.v1alpha1.GetWriteTokenResponse
	5, // 5: commonfate.factory.monitoring.v1alpha1.TokenService.DeleteWriteToken:output_type -> commonfate.factory.monitoring.v1alpha1.DeleteWriteTokenResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commonfate_factory_monitoring_v1alpha1_token_proto_init() }
func file_commonfate_factory_monitoring_v1alpha1_token_proto_init() {
	if File_commonfate_factory_monitoring_v1alpha1_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWriteTokenRequest); i {
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
		file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWriteTokenResponse); i {
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
		file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWriteTokenRequest); i {
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
		file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWriteTokenResponse); i {
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
		file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWriteTokenRequest); i {
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
		file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWriteTokenResponse); i {
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
			RawDescriptor: file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_factory_monitoring_v1alpha1_token_proto_goTypes,
		DependencyIndexes: file_commonfate_factory_monitoring_v1alpha1_token_proto_depIdxs,
		MessageInfos:      file_commonfate_factory_monitoring_v1alpha1_token_proto_msgTypes,
	}.Build()
	File_commonfate_factory_monitoring_v1alpha1_token_proto = out.File
	file_commonfate_factory_monitoring_v1alpha1_token_proto_rawDesc = nil
	file_commonfate_factory_monitoring_v1alpha1_token_proto_goTypes = nil
	file_commonfate_factory_monitoring_v1alpha1_token_proto_depIdxs = nil
}