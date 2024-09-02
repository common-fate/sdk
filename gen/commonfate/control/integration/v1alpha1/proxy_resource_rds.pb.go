// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: commonfate/control/integration/v1alpha1/proxy_resource_rds.proto

package integrationv1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
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

type CreateProxyRdsResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProxyId     string          `protobuf:"bytes,1,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	RdsDatabase *AWSRDSDatabase `protobuf:"bytes,2,opt,name=rds_database,json=rdsDatabase,proto3" json:"rds_database,omitempty"`
}

func (x *CreateProxyRdsResourceRequest) Reset() {
	*x = CreateProxyRdsResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProxyRdsResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProxyRdsResourceRequest) ProtoMessage() {}

func (x *CreateProxyRdsResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProxyRdsResourceRequest.ProtoReflect.Descriptor instead.
func (*CreateProxyRdsResourceRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProxyRdsResourceRequest) GetProxyId() string {
	if x != nil {
		return x.ProxyId
	}
	return ""
}

func (x *CreateProxyRdsResourceRequest) GetRdsDatabase() *AWSRDSDatabase {
	if x != nil {
		return x.RdsDatabase
	}
	return nil
}

type CreateProxyRdsResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProxyId     string          `protobuf:"bytes,2,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	RdsDatabase *AWSRDSDatabase `protobuf:"bytes,3,opt,name=rds_database,json=rdsDatabase,proto3" json:"rds_database,omitempty"`
}

func (x *CreateProxyRdsResourceResponse) Reset() {
	*x = CreateProxyRdsResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProxyRdsResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProxyRdsResourceResponse) ProtoMessage() {}

func (x *CreateProxyRdsResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProxyRdsResourceResponse.ProtoReflect.Descriptor instead.
func (*CreateProxyRdsResourceResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProxyRdsResourceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateProxyRdsResourceResponse) GetProxyId() string {
	if x != nil {
		return x.ProxyId
	}
	return ""
}

func (x *CreateProxyRdsResourceResponse) GetRdsDatabase() *AWSRDSDatabase {
	if x != nil {
		return x.RdsDatabase
	}
	return nil
}

type GetProxyRdsResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProxyRdsResourceRequest) Reset() {
	*x = GetProxyRdsResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProxyRdsResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProxyRdsResourceRequest) ProtoMessage() {}

func (x *GetProxyRdsResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProxyRdsResourceRequest.ProtoReflect.Descriptor instead.
func (*GetProxyRdsResourceRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescGZIP(), []int{2}
}

func (x *GetProxyRdsResourceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetProxyRdsResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RdsDatabase *AWSRDSDatabase `protobuf:"bytes,2,opt,name=rds_database,json=rdsDatabase,proto3" json:"rds_database,omitempty"`
}

func (x *GetProxyRdsResourceResponse) Reset() {
	*x = GetProxyRdsResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProxyRdsResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProxyRdsResourceResponse) ProtoMessage() {}

func (x *GetProxyRdsResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProxyRdsResourceResponse.ProtoReflect.Descriptor instead.
func (*GetProxyRdsResourceResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescGZIP(), []int{3}
}

func (x *GetProxyRdsResourceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProxyRdsResourceResponse) GetRdsDatabase() *AWSRDSDatabase {
	if x != nil {
		return x.RdsDatabase
	}
	return nil
}

type UpdateProxyRdsResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProxyId     string          `protobuf:"bytes,2,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	RdsDatabase *AWSRDSDatabase `protobuf:"bytes,3,opt,name=rds_database,json=rdsDatabase,proto3" json:"rds_database,omitempty"`
}

func (x *UpdateProxyRdsResourceRequest) Reset() {
	*x = UpdateProxyRdsResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProxyRdsResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProxyRdsResourceRequest) ProtoMessage() {}

func (x *UpdateProxyRdsResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProxyRdsResourceRequest.ProtoReflect.Descriptor instead.
func (*UpdateProxyRdsResourceRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProxyRdsResourceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProxyRdsResourceRequest) GetProxyId() string {
	if x != nil {
		return x.ProxyId
	}
	return ""
}

func (x *UpdateProxyRdsResourceRequest) GetRdsDatabase() *AWSRDSDatabase {
	if x != nil {
		return x.RdsDatabase
	}
	return nil
}

type UpdateProxyRdsResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProxyId     string          `protobuf:"bytes,2,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	RdsDatabase *AWSRDSDatabase `protobuf:"bytes,3,opt,name=rds_database,json=rdsDatabase,proto3" json:"rds_database,omitempty"`
}

func (x *UpdateProxyRdsResourceResponse) Reset() {
	*x = UpdateProxyRdsResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProxyRdsResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProxyRdsResourceResponse) ProtoMessage() {}

func (x *UpdateProxyRdsResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProxyRdsResourceResponse.ProtoReflect.Descriptor instead.
func (*UpdateProxyRdsResourceResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateProxyRdsResourceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateProxyRdsResourceResponse) GetProxyId() string {
	if x != nil {
		return x.ProxyId
	}
	return ""
}

func (x *UpdateProxyRdsResourceResponse) GetRdsDatabase() *AWSRDSDatabase {
	if x != nil {
		return x.RdsDatabase
	}
	return nil
}

type DeleteProxyRdsResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProxyRdsResourceRequest) Reset() {
	*x = DeleteProxyRdsResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProxyRdsResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProxyRdsResourceRequest) ProtoMessage() {}

func (x *DeleteProxyRdsResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProxyRdsResourceRequest.ProtoReflect.Descriptor instead.
func (*DeleteProxyRdsResourceRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteProxyRdsResourceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteProxyRdsResourceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProxyRdsResourceResponse) Reset() {
	*x = DeleteProxyRdsResourceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProxyRdsResourceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProxyRdsResourceResponse) ProtoMessage() {}

func (x *DeleteProxyRdsResourceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProxyRdsResourceResponse.ProtoReflect.Descriptor instead.
func (*DeleteProxyRdsResourceResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProxyRdsResourceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto protoreflect.FileDescriptor

var file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x69, 0x64, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x37, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x61, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x33, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67,
	0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x6f, 0x6b, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f, 0x70, 0x73, 0x67, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x64, 0x75, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x33, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x33, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x1d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x64, 0x12, 0x5a, 0x0a, 0x0c, 0x72, 0x64, 0x73, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x52, 0x44, 0x53, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x0b, 0x72, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x52, 0x64, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x49, 0x64, 0x12, 0x5a, 0x0a, 0x0c, 0x72, 0x64, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x52, 0x44, 0x53, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x0b, 0x72, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0x2c,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x89, 0x01, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x64, 0x73, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x5a, 0x0a, 0x0c,
	0x72, 0x64, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53,
	0x52, 0x44, 0x53, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x0b, 0x72, 0x64, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x64, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x49, 0x64, 0x12, 0x5a, 0x0a, 0x0c, 0x72, 0x64, 0x73, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x52, 0x44, 0x53, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x0b, 0x72, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x22, 0xa7, 0x01, 0x0a, 0x1e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x52, 0x64, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x64, 0x12,
	0x5a, 0x0a, 0x0c, 0x72, 0x64, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x57, 0x53, 0x52, 0x44, 0x53, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x0b,
	0x72, 0x64, 0x73, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0x2f, 0x0a, 0x1d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x1e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0xdf,
	0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x15,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x64, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x49, 0xaa, 0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xca, 0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x33,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x2a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescOnce sync.Once
	file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescData = file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDesc
)

func file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescGZIP() []byte {
	file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescOnce.Do(func() {
		file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescData)
	})
	return file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDescData
}

var file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_goTypes = []any{
	(*CreateProxyRdsResourceRequest)(nil),  // 0: commonfate.control.integration.v1alpha1.CreateProxyRdsResourceRequest
	(*CreateProxyRdsResourceResponse)(nil), // 1: commonfate.control.integration.v1alpha1.CreateProxyRdsResourceResponse
	(*GetProxyRdsResourceRequest)(nil),     // 2: commonfate.control.integration.v1alpha1.GetProxyRdsResourceRequest
	(*GetProxyRdsResourceResponse)(nil),    // 3: commonfate.control.integration.v1alpha1.GetProxyRdsResourceResponse
	(*UpdateProxyRdsResourceRequest)(nil),  // 4: commonfate.control.integration.v1alpha1.UpdateProxyRdsResourceRequest
	(*UpdateProxyRdsResourceResponse)(nil), // 5: commonfate.control.integration.v1alpha1.UpdateProxyRdsResourceResponse
	(*DeleteProxyRdsResourceRequest)(nil),  // 6: commonfate.control.integration.v1alpha1.DeleteProxyRdsResourceRequest
	(*DeleteProxyRdsResourceResponse)(nil), // 7: commonfate.control.integration.v1alpha1.DeleteProxyRdsResourceResponse
	(*AWSRDSDatabase)(nil),                 // 8: commonfate.control.integration.v1alpha1.AWSRDSDatabase
}
var file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_depIdxs = []int32{
	8, // 0: commonfate.control.integration.v1alpha1.CreateProxyRdsResourceRequest.rds_database:type_name -> commonfate.control.integration.v1alpha1.AWSRDSDatabase
	8, // 1: commonfate.control.integration.v1alpha1.CreateProxyRdsResourceResponse.rds_database:type_name -> commonfate.control.integration.v1alpha1.AWSRDSDatabase
	8, // 2: commonfate.control.integration.v1alpha1.GetProxyRdsResourceResponse.rds_database:type_name -> commonfate.control.integration.v1alpha1.AWSRDSDatabase
	8, // 3: commonfate.control.integration.v1alpha1.UpdateProxyRdsResourceRequest.rds_database:type_name -> commonfate.control.integration.v1alpha1.AWSRDSDatabase
	8, // 4: commonfate.control.integration.v1alpha1.UpdateProxyRdsResourceResponse.rds_database:type_name -> commonfate.control.integration.v1alpha1.AWSRDSDatabase
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_init() }
func file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_init() {
	if File_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto != nil {
		return
	}
	file_commonfate_control_integration_v1alpha1_auth0_proto_init()
	file_commonfate_control_integration_v1alpha1_aws_idc_proto_init()
	file_commonfate_control_integration_v1alpha1_aws_proxy_proto_init()
	file_commonfate_control_integration_v1alpha1_datastax_proto_init()
	file_commonfate_control_integration_v1alpha1_entra_proto_init()
	file_commonfate_control_integration_v1alpha1_gcp_proto_init()
	file_commonfate_control_integration_v1alpha1_okta_proto_init()
	file_commonfate_control_integration_v1alpha1_opsgenie_proto_init()
	file_commonfate_control_integration_v1alpha1_pagerduty_proto_init()
	file_commonfate_control_integration_v1alpha1_s3_log_destination_proto_init()
	file_commonfate_control_integration_v1alpha1_slack_proto_init()
	file_commonfate_control_integration_v1alpha1_webhook_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProxyRdsResourceRequest); i {
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
		file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateProxyRdsResourceResponse); i {
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
		file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetProxyRdsResourceRequest); i {
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
		file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetProxyRdsResourceResponse); i {
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
		file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProxyRdsResourceRequest); i {
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
		file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProxyRdsResourceResponse); i {
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
		file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteProxyRdsResourceRequest); i {
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
		file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteProxyRdsResourceResponse); i {
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
			RawDescriptor: file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_goTypes,
		DependencyIndexes: file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_depIdxs,
		MessageInfos:      file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_msgTypes,
	}.Build()
	File_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto = out.File
	file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_rawDesc = nil
	file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_goTypes = nil
	file_commonfate_control_integration_v1alpha1_proxy_resource_rds_proto_depIdxs = nil
}