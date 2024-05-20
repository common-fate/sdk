// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: commonfate/factory/deployment/v1alpha1/deployment.proto

package deploymentv1alpha1

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

type DNSRecordType int32

const (
	DNSRecordType_DNS_RECORD_TYPE_UNSPECIFIED DNSRecordType = 0
	DNSRecordType_DNS_RECORD_TYPE_TXT         DNSRecordType = 1
	DNSRecordType_DNS_RECORD_TYPE_CNAME       DNSRecordType = 2
)

// Enum value maps for DNSRecordType.
var (
	DNSRecordType_name = map[int32]string{
		0: "DNS_RECORD_TYPE_UNSPECIFIED",
		1: "DNS_RECORD_TYPE_TXT",
		2: "DNS_RECORD_TYPE_CNAME",
	}
	DNSRecordType_value = map[string]int32{
		"DNS_RECORD_TYPE_UNSPECIFIED": 0,
		"DNS_RECORD_TYPE_TXT":         1,
		"DNS_RECORD_TYPE_CNAME":       2,
	}
)

func (x DNSRecordType) Enum() *DNSRecordType {
	p := new(DNSRecordType)
	*p = x
	return p
}

func (x DNSRecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DNSRecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_enumTypes[0].Descriptor()
}

func (DNSRecordType) Type() protoreflect.EnumType {
	return &file_commonfate_factory_deployment_v1alpha1_deployment_proto_enumTypes[0]
}

func (x DNSRecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DNSRecordType.Descriptor instead.
func (DNSRecordType) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP(), []int{0}
}

type DNSRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type   DNSRecordType `protobuf:"varint,2,opt,name=type,proto3,enum=commonfate.factory.deployment.v1alpha1.DNSRecordType" json:"type,omitempty"`
	Values []string      `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *DNSRecord) Reset() {
	*x = DNSRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSRecord) ProtoMessage() {}

func (x *DNSRecord) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSRecord.ProtoReflect.Descriptor instead.
func (*DNSRecord) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *DNSRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DNSRecord) GetType() DNSRecordType {
	if x != nil {
		return x.Type
	}
	return DNSRecordType_DNS_RECORD_TYPE_UNSPECIFIED
}

func (x *DNSRecord) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type CreateDNSRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *DNSRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *CreateDNSRecordRequest) Reset() {
	*x = CreateDNSRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDNSRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDNSRecordRequest) ProtoMessage() {}

func (x *CreateDNSRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDNSRecordRequest.ProtoReflect.Descriptor instead.
func (*CreateDNSRecordRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDNSRecordRequest) GetRecord() *DNSRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

type CreateDNSRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateDNSRecordResponse) Reset() {
	*x = CreateDNSRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDNSRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDNSRecordResponse) ProtoMessage() {}

func (x *CreateDNSRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDNSRecordResponse.ProtoReflect.Descriptor instead.
func (*CreateDNSRecordResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP(), []int{2}
}

type GetDNSRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDNSRecordRequest) Reset() {
	*x = GetDNSRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDNSRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDNSRecordRequest) ProtoMessage() {}

func (x *GetDNSRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDNSRecordRequest.ProtoReflect.Descriptor instead.
func (*GetDNSRecordRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP(), []int{3}
}

func (x *GetDNSRecordRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDNSRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *DNSRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *GetDNSRecordResponse) Reset() {
	*x = GetDNSRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDNSRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDNSRecordResponse) ProtoMessage() {}

func (x *GetDNSRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDNSRecordResponse.ProtoReflect.Descriptor instead.
func (*GetDNSRecordResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP(), []int{4}
}

func (x *GetDNSRecordResponse) GetRecord() *DNSRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

type GetDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDeploymentRequest) Reset() {
	*x = GetDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentRequest) ProtoMessage() {}

func (x *GetDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP(), []int{5}
}

type GetDeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment *Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
}

func (x *GetDeploymentResponse) Reset() {
	*x = GetDeploymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentResponse) ProtoMessage() {}

func (x *GetDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentResponse.ProtoReflect.Descriptor instead.
func (*GetDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP(), []int{6}
}

func (x *GetDeploymentResponse) GetDeployment() *Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DefaultDomain string   `protobuf:"bytes,2,opt,name=default_domain,json=defaultDomain,proto3" json:"default_domain,omitempty"`
	Nameservers   []string `protobuf:"bytes,3,rep,name=nameservers,proto3" json:"nameservers,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP(), []int{7}
}

func (x *Deployment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Deployment) GetDefaultDomain() string {
	if x != nil {
		return x.DefaultDomain
	}
	return ""
}

func (x *Deployment) GetNameservers() []string {
	if x != nil {
		return x.Nameservers
	}
	return nil
}

var File_commonfate_factory_deployment_v1alpha1_deployment_proto protoreflect.FileDescriptor

var file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x22, 0x82, 0x01, 0x0a, 0x09, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x4e, 0x53, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x63, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x49, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x61, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x65, 0x0a, 0x0a, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x2a, 0x64, 0x0a, 0x0d, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x4e, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x4e, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x58, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x44,
	0x4e, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x32, 0xc9, 0x03, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x94, 0x01,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x3e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0xd2, 0x02, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x46, 0x44, 0xaa, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x26,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x32, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x5c, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x29, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescOnce sync.Once
	file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescData = file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDesc
)

func file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescGZIP() []byte {
	file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescOnce.Do(func() {
		file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescData)
	})
	return file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDescData
}

var file_commonfate_factory_deployment_v1alpha1_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_commonfate_factory_deployment_v1alpha1_deployment_proto_goTypes = []interface{}{
	(DNSRecordType)(0),              // 0: commonfate.factory.deployment.v1alpha1.DNSRecordType
	(*DNSRecord)(nil),               // 1: commonfate.factory.deployment.v1alpha1.DNSRecord
	(*CreateDNSRecordRequest)(nil),  // 2: commonfate.factory.deployment.v1alpha1.CreateDNSRecordRequest
	(*CreateDNSRecordResponse)(nil), // 3: commonfate.factory.deployment.v1alpha1.CreateDNSRecordResponse
	(*GetDNSRecordRequest)(nil),     // 4: commonfate.factory.deployment.v1alpha1.GetDNSRecordRequest
	(*GetDNSRecordResponse)(nil),    // 5: commonfate.factory.deployment.v1alpha1.GetDNSRecordResponse
	(*GetDeploymentRequest)(nil),    // 6: commonfate.factory.deployment.v1alpha1.GetDeploymentRequest
	(*GetDeploymentResponse)(nil),   // 7: commonfate.factory.deployment.v1alpha1.GetDeploymentResponse
	(*Deployment)(nil),              // 8: commonfate.factory.deployment.v1alpha1.Deployment
}
var file_commonfate_factory_deployment_v1alpha1_deployment_proto_depIdxs = []int32{
	0, // 0: commonfate.factory.deployment.v1alpha1.DNSRecord.type:type_name -> commonfate.factory.deployment.v1alpha1.DNSRecordType
	1, // 1: commonfate.factory.deployment.v1alpha1.CreateDNSRecordRequest.record:type_name -> commonfate.factory.deployment.v1alpha1.DNSRecord
	1, // 2: commonfate.factory.deployment.v1alpha1.GetDNSRecordResponse.record:type_name -> commonfate.factory.deployment.v1alpha1.DNSRecord
	8, // 3: commonfate.factory.deployment.v1alpha1.GetDeploymentResponse.deployment:type_name -> commonfate.factory.deployment.v1alpha1.Deployment
	6, // 4: commonfate.factory.deployment.v1alpha1.DeploymentService.GetDeployment:input_type -> commonfate.factory.deployment.v1alpha1.GetDeploymentRequest
	2, // 5: commonfate.factory.deployment.v1alpha1.DeploymentService.CreateDNSRecord:input_type -> commonfate.factory.deployment.v1alpha1.CreateDNSRecordRequest
	4, // 6: commonfate.factory.deployment.v1alpha1.DeploymentService.GetDNSRecord:input_type -> commonfate.factory.deployment.v1alpha1.GetDNSRecordRequest
	7, // 7: commonfate.factory.deployment.v1alpha1.DeploymentService.GetDeployment:output_type -> commonfate.factory.deployment.v1alpha1.GetDeploymentResponse
	3, // 8: commonfate.factory.deployment.v1alpha1.DeploymentService.CreateDNSRecord:output_type -> commonfate.factory.deployment.v1alpha1.CreateDNSRecordResponse
	5, // 9: commonfate.factory.deployment.v1alpha1.DeploymentService.GetDNSRecord:output_type -> commonfate.factory.deployment.v1alpha1.GetDNSRecordResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_commonfate_factory_deployment_v1alpha1_deployment_proto_init() }
func file_commonfate_factory_deployment_v1alpha1_deployment_proto_init() {
	if File_commonfate_factory_deployment_v1alpha1_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSRecord); i {
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
		file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDNSRecordRequest); i {
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
		file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDNSRecordResponse); i {
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
		file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDNSRecordRequest); i {
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
		file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDNSRecordResponse); i {
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
		file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentRequest); i {
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
		file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentResponse); i {
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
		file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
			RawDescriptor: file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_factory_deployment_v1alpha1_deployment_proto_goTypes,
		DependencyIndexes: file_commonfate_factory_deployment_v1alpha1_deployment_proto_depIdxs,
		EnumInfos:         file_commonfate_factory_deployment_v1alpha1_deployment_proto_enumTypes,
		MessageInfos:      file_commonfate_factory_deployment_v1alpha1_deployment_proto_msgTypes,
	}.Build()
	File_commonfate_factory_deployment_v1alpha1_deployment_proto = out.File
	file_commonfate_factory_deployment_v1alpha1_deployment_proto_rawDesc = nil
	file_commonfate_factory_deployment_v1alpha1_deployment_proto_goTypes = nil
	file_commonfate_factory_deployment_v1alpha1_deployment_proto_depIdxs = nil
}
