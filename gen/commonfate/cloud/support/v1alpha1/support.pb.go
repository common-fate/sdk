// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: commonfate/cloud/support/v1alpha1/support.proto

package supportv1alpha1

import (
	v1alpha11 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/control/support/v1alpha1"
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

type ContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *v1alpha1.Ticket `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	User   *v1alpha11.User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ContactRequest) Reset() {
	*x = ContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactRequest) ProtoMessage() {}

func (x *ContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactRequest.ProtoReflect.Descriptor instead.
func (*ContactRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_cloud_support_v1alpha1_support_proto_rawDescGZIP(), []int{0}
}

func (x *ContactRequest) GetTicket() *v1alpha1.Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

func (x *ContactRequest) GetUser() *v1alpha11.User {
	if x != nil {
		return x.User
	}
	return nil
}

type ContactResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An initial response indicating whether the support request has been received.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ContactResponse) Reset() {
	*x = ContactResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactResponse) ProtoMessage() {}

func (x *ContactResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactResponse.ProtoReflect.Descriptor instead.
func (*ContactResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_cloud_support_v1alpha1_support_proto_rawDescGZIP(), []int{1}
}

func (x *ContactResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAttachmentUploadURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileInput *v1alpha1.FileInput `protobuf:"bytes,1,opt,name=file_input,json=fileInput,proto3" json:"file_input,omitempty"`
	User      *v1alpha11.User     `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetAttachmentUploadURLRequest) Reset() {
	*x = GetAttachmentUploadURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachmentUploadURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachmentUploadURLRequest) ProtoMessage() {}

func (x *GetAttachmentUploadURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachmentUploadURLRequest.ProtoReflect.Descriptor instead.
func (*GetAttachmentUploadURLRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_cloud_support_v1alpha1_support_proto_rawDescGZIP(), []int{2}
}

func (x *GetAttachmentUploadURLRequest) GetFileInput() *v1alpha1.FileInput {
	if x != nil {
		return x.FileInput
	}
	return nil
}

func (x *GetAttachmentUploadURLRequest) GetUser() *v1alpha11.User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetAttachmentUploadURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAttachmentUploadURLResponse) Reset() {
	*x = GetAttachmentUploadURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachmentUploadURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachmentUploadURLResponse) ProtoMessage() {}

func (x *GetAttachmentUploadURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachmentUploadURLResponse.ProtoReflect.Descriptor instead.
func (*GetAttachmentUploadURLResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_cloud_support_v1alpha1_support_proto_rawDescGZIP(), []int{3}
}

var File_commonfate_cloud_support_v1alpha1_support_proto protoreflect.FileDescriptor

var file_commonfate_cloud_support_v1alpha1_support_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x31, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b,
	0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x43, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x06,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x0f,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0a, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x20, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xa6, 0x02, 0x0a, 0x0e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9f, 0x01, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x55, 0x52, 0x4c, 0x12, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xae, 0x02, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x53, 0xaa, 0x02, 0x21,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xca, 0x02, 0x21, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_cloud_support_v1alpha1_support_proto_rawDescOnce sync.Once
	file_commonfate_cloud_support_v1alpha1_support_proto_rawDescData = file_commonfate_cloud_support_v1alpha1_support_proto_rawDesc
)

func file_commonfate_cloud_support_v1alpha1_support_proto_rawDescGZIP() []byte {
	file_commonfate_cloud_support_v1alpha1_support_proto_rawDescOnce.Do(func() {
		file_commonfate_cloud_support_v1alpha1_support_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_cloud_support_v1alpha1_support_proto_rawDescData)
	})
	return file_commonfate_cloud_support_v1alpha1_support_proto_rawDescData
}

var file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_commonfate_cloud_support_v1alpha1_support_proto_goTypes = []interface{}{
	(*ContactRequest)(nil),                 // 0: commonfate.cloud.support.v1alpha1.ContactRequest
	(*ContactResponse)(nil),                // 1: commonfate.cloud.support.v1alpha1.ContactResponse
	(*GetAttachmentUploadURLRequest)(nil),  // 2: commonfate.cloud.support.v1alpha1.GetAttachmentUploadURLRequest
	(*GetAttachmentUploadURLResponse)(nil), // 3: commonfate.cloud.support.v1alpha1.GetAttachmentUploadURLResponse
	(*v1alpha1.Ticket)(nil),                // 4: commonfate.control.support.v1alpha1.Ticket
	(*v1alpha11.User)(nil),                 // 5: commonfate.access.v1alpha1.User
	(*v1alpha1.FileInput)(nil),             // 6: commonfate.control.support.v1alpha1.FileInput
}
var file_commonfate_cloud_support_v1alpha1_support_proto_depIdxs = []int32{
	4, // 0: commonfate.cloud.support.v1alpha1.ContactRequest.ticket:type_name -> commonfate.control.support.v1alpha1.Ticket
	5, // 1: commonfate.cloud.support.v1alpha1.ContactRequest.user:type_name -> commonfate.access.v1alpha1.User
	6, // 2: commonfate.cloud.support.v1alpha1.GetAttachmentUploadURLRequest.file_input:type_name -> commonfate.control.support.v1alpha1.FileInput
	5, // 3: commonfate.cloud.support.v1alpha1.GetAttachmentUploadURLRequest.user:type_name -> commonfate.access.v1alpha1.User
	0, // 4: commonfate.cloud.support.v1alpha1.SupportService.Contact:input_type -> commonfate.cloud.support.v1alpha1.ContactRequest
	2, // 5: commonfate.cloud.support.v1alpha1.SupportService.GetAttachmentUploadURL:input_type -> commonfate.cloud.support.v1alpha1.GetAttachmentUploadURLRequest
	1, // 6: commonfate.cloud.support.v1alpha1.SupportService.Contact:output_type -> commonfate.cloud.support.v1alpha1.ContactResponse
	3, // 7: commonfate.cloud.support.v1alpha1.SupportService.GetAttachmentUploadURL:output_type -> commonfate.cloud.support.v1alpha1.GetAttachmentUploadURLResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_commonfate_cloud_support_v1alpha1_support_proto_init() }
func file_commonfate_cloud_support_v1alpha1_support_proto_init() {
	if File_commonfate_cloud_support_v1alpha1_support_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactRequest); i {
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
		file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactResponse); i {
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
		file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachmentUploadURLRequest); i {
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
		file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachmentUploadURLResponse); i {
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
			RawDescriptor: file_commonfate_cloud_support_v1alpha1_support_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_cloud_support_v1alpha1_support_proto_goTypes,
		DependencyIndexes: file_commonfate_cloud_support_v1alpha1_support_proto_depIdxs,
		MessageInfos:      file_commonfate_cloud_support_v1alpha1_support_proto_msgTypes,
	}.Build()
	File_commonfate_cloud_support_v1alpha1_support_proto = out.File
	file_commonfate_cloud_support_v1alpha1_support_proto_rawDesc = nil
	file_commonfate_cloud_support_v1alpha1_support_proto_goTypes = nil
	file_commonfate_cloud_support_v1alpha1_support_proto_depIdxs = nil
}
