// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: commonfate/control/attest/v1alpha1/attestation.proto

package attestv1alpha1

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

type AttestationType int32

const (
	AttestationType_ATTESTATION_TYPE_UNSPECIFIED         AttestationType = 0
	AttestationType_ATTESTATION_TYPE_DEVICE_REGISTRATION AttestationType = 1
	AttestationType_ATTESTATION_TYPE_ACCESS_REQUEST      AttestationType = 2
)

// Enum value maps for AttestationType.
var (
	AttestationType_name = map[int32]string{
		0: "ATTESTATION_TYPE_UNSPECIFIED",
		1: "ATTESTATION_TYPE_DEVICE_REGISTRATION",
		2: "ATTESTATION_TYPE_ACCESS_REQUEST",
	}
	AttestationType_value = map[string]int32{
		"ATTESTATION_TYPE_UNSPECIFIED":         0,
		"ATTESTATION_TYPE_DEVICE_REGISTRATION": 1,
		"ATTESTATION_TYPE_ACCESS_REQUEST":      2,
	}
)

func (x AttestationType) Enum() *AttestationType {
	p := new(AttestationType)
	*p = x
	return p
}

func (x AttestationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AttestationType) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_control_attest_v1alpha1_attestation_proto_enumTypes[0].Descriptor()
}

func (AttestationType) Type() protoreflect.EnumType {
	return &file_commonfate_control_attest_v1alpha1_attestation_proto_enumTypes[0]
}

func (x AttestationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AttestationType.Descriptor instead.
func (AttestationType) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescGZIP(), []int{0}
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// should always be '1'
	Version int64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// the current timestamp in milliseconds since Unix epoch
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// the type of message being signed
	Type AttestationType `protobuf:"varint,3,opt,name=type,proto3,enum=commonfate.control.attest.v1alpha1.AttestationType" json:"type,omitempty"`
	// the SHA256 digest of the message contents
	ContentDigest []byte `protobuf:"bytes,4,opt,name=content_digest,json=contentDigest,proto3" json:"content_digest,omitempty"`
	// the ID of the key that signed the message
	Kid string `protobuf:"bytes,5,opt,name=kid,proto3" json:"kid,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	mi := &file_commonfate_control_attest_v1alpha1_attestation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_attest_v1alpha1_attestation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Header) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Header) GetType() AttestationType {
	if x != nil {
		return x.Type
	}
	return AttestationType_ATTESTATION_TYPE_UNSPECIFIED
}

func (x *Header) GetContentDigest() []byte {
	if x != nil {
		return x.ContentDigest
	}
	return nil
}

func (x *Header) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber  string `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	Uuid          string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Platform      string `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
	Architecture  string `protobuf:"bytes,5,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Hostname      string `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	KernelRelease string `protobuf:"bytes,7,opt,name=kernel_release,json=kernelRelease,proto3" json:"kernel_release,omitempty"`
	KernelVersion string `protobuf:"bytes,8,opt,name=kernel_version,json=kernelVersion,proto3" json:"kernel_version,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	mi := &file_commonfate_control_attest_v1alpha1_attestation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_attest_v1alpha1_attestation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Device) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Device) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Device) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *Device) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Device) GetKernelRelease() string {
	if x != nil {
		return x.KernelRelease
	}
	return ""
}

func (x *Device) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Signature []byte  `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	mi := &file_commonfate_control_attest_v1alpha1_attestation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_attest_v1alpha1_attestation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescGZIP(), []int{2}
}

func (x *Attestation) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Attestation) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_commonfate_control_attest_v1alpha1_attestation_proto protoreflect.FileDescriptor

var file_commonfate_control_attest_v1alpha1_attestation_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0xc2, 0x01, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x47, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x69, 0x64, 0x22,
	0xeb, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6f, 0x0a,
	0x0b, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2a, 0x82,
	0x01, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x23,
	0x0a, 0x1f, 0x41, 0x54, 0x54, 0x45, 0x53, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x02, 0x42, 0xb7, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x10,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x41, 0xaa, 0x02, 0x22, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescOnce sync.Once
	file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescData = file_commonfate_control_attest_v1alpha1_attestation_proto_rawDesc
)

func file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescGZIP() []byte {
	file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescOnce.Do(func() {
		file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescData)
	})
	return file_commonfate_control_attest_v1alpha1_attestation_proto_rawDescData
}

var file_commonfate_control_attest_v1alpha1_attestation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_commonfate_control_attest_v1alpha1_attestation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commonfate_control_attest_v1alpha1_attestation_proto_goTypes = []any{
	(AttestationType)(0), // 0: commonfate.control.attest.v1alpha1.AttestationType
	(*Header)(nil),       // 1: commonfate.control.attest.v1alpha1.Header
	(*Device)(nil),       // 2: commonfate.control.attest.v1alpha1.Device
	(*Attestation)(nil),  // 3: commonfate.control.attest.v1alpha1.Attestation
}
var file_commonfate_control_attest_v1alpha1_attestation_proto_depIdxs = []int32{
	0, // 0: commonfate.control.attest.v1alpha1.Header.type:type_name -> commonfate.control.attest.v1alpha1.AttestationType
	1, // 1: commonfate.control.attest.v1alpha1.Attestation.header:type_name -> commonfate.control.attest.v1alpha1.Header
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commonfate_control_attest_v1alpha1_attestation_proto_init() }
func file_commonfate_control_attest_v1alpha1_attestation_proto_init() {
	if File_commonfate_control_attest_v1alpha1_attestation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_control_attest_v1alpha1_attestation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_control_attest_v1alpha1_attestation_proto_goTypes,
		DependencyIndexes: file_commonfate_control_attest_v1alpha1_attestation_proto_depIdxs,
		EnumInfos:         file_commonfate_control_attest_v1alpha1_attestation_proto_enumTypes,
		MessageInfos:      file_commonfate_control_attest_v1alpha1_attestation_proto_msgTypes,
	}.Build()
	File_commonfate_control_attest_v1alpha1_attestation_proto = out.File
	file_commonfate_control_attest_v1alpha1_attestation_proto_rawDesc = nil
	file_commonfate_control_attest_v1alpha1_attestation_proto_goTypes = nil
	file_commonfate_control_attest_v1alpha1_attestation_proto_depIdxs = nil
}
