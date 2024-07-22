// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: commonfate/control/config/v1alpha1/aws_resource_scanner.proto

package configv1alpha1

import (
	_ "github.com/common-fate/sdk/gen/buf/validate"
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

type CreateAWSResourceScannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId string `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	// AWS regions to scan for resources in.
	Regions []string `protobuf:"bytes,2,rep,name=regions,proto3" json:"regions,omitempty"`
	// Resource types to scan for. If empty, Common Fate
	// will attempt to scan for all supported AWS resource types.
	//
	// Resource types should adhere to the Cedar format,
	// for example 'AWS::S3::Bucket'.
	ResourceTypes []string `protobuf:"bytes,3,rep,name=resource_types,json=resourceTypes,proto3" json:"resource_types,omitempty"`
	// If provided, only resources in the specified account IDs will be scanned.
	FilterForAccountIds []string `protobuf:"bytes,4,rep,name=filter_for_account_ids,json=filterForAccountIds,proto3" json:"filter_for_account_ids,omitempty"`
	// The AWS role name to use for scanning resources. Defaults to 'common-fate-audit' if not provided.
	RoleName string `protobuf:"bytes,5,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *CreateAWSResourceScannerRequest) Reset() {
	*x = CreateAWSResourceScannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAWSResourceScannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAWSResourceScannerRequest) ProtoMessage() {}

func (x *CreateAWSResourceScannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAWSResourceScannerRequest.ProtoReflect.Descriptor instead.
func (*CreateAWSResourceScannerRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAWSResourceScannerRequest) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *CreateAWSResourceScannerRequest) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *CreateAWSResourceScannerRequest) GetResourceTypes() []string {
	if x != nil {
		return x.ResourceTypes
	}
	return nil
}

func (x *CreateAWSResourceScannerRequest) GetFilterForAccountIds() []string {
	if x != nil {
		return x.FilterForAccountIds
	}
	return nil
}

func (x *CreateAWSResourceScannerRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type AWSResourceScanner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IntegrationId string `protobuf:"bytes,2,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	// AWS regions to scan for resources in.
	Regions []string `protobuf:"bytes,3,rep,name=regions,proto3" json:"regions,omitempty"`
	// Resource types to scan for. If empty, Common Fate
	// will attempt to scan for all supported AWS resource types.
	//
	// Resource types should adhere to the Cedar format,
	// for example 'AWS::S3::Bucket'.
	ResourceTypes []string `protobuf:"bytes,4,rep,name=resource_types,json=resourceTypes,proto3" json:"resource_types,omitempty"`
	// If provided, only resources in the specified account IDs will be scanned.
	FilterForAccountIds []string `protobuf:"bytes,5,rep,name=filter_for_account_ids,json=filterForAccountIds,proto3" json:"filter_for_account_ids,omitempty"`
	// The AWS role name to use for scanning resources. Defaults to 'common-fate-audit' if not provided.
	RoleName string `protobuf:"bytes,6,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *AWSResourceScanner) Reset() {
	*x = AWSResourceScanner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSResourceScanner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSResourceScanner) ProtoMessage() {}

func (x *AWSResourceScanner) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSResourceScanner.ProtoReflect.Descriptor instead.
func (*AWSResourceScanner) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP(), []int{1}
}

func (x *AWSResourceScanner) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AWSResourceScanner) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *AWSResourceScanner) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *AWSResourceScanner) GetResourceTypes() []string {
	if x != nil {
		return x.ResourceTypes
	}
	return nil
}

func (x *AWSResourceScanner) GetFilterForAccountIds() []string {
	if x != nil {
		return x.FilterForAccountIds
	}
	return nil
}

func (x *AWSResourceScanner) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type CreateAWSResourceScannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceScanner *AWSResourceScanner `protobuf:"bytes,1,opt,name=resource_scanner,json=resourceScanner,proto3" json:"resource_scanner,omitempty"`
}

func (x *CreateAWSResourceScannerResponse) Reset() {
	*x = CreateAWSResourceScannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAWSResourceScannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAWSResourceScannerResponse) ProtoMessage() {}

func (x *CreateAWSResourceScannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAWSResourceScannerResponse.ProtoReflect.Descriptor instead.
func (*CreateAWSResourceScannerResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAWSResourceScannerResponse) GetResourceScanner() *AWSResourceScanner {
	if x != nil {
		return x.ResourceScanner
	}
	return nil
}

type GetAWSResourceScannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAWSResourceScannerRequest) Reset() {
	*x = GetAWSResourceScannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAWSResourceScannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAWSResourceScannerRequest) ProtoMessage() {}

func (x *GetAWSResourceScannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAWSResourceScannerRequest.ProtoReflect.Descriptor instead.
func (*GetAWSResourceScannerRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP(), []int{3}
}

func (x *GetAWSResourceScannerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAWSResourceScannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceScanner *AWSResourceScanner `protobuf:"bytes,1,opt,name=resource_scanner,json=resourceScanner,proto3" json:"resource_scanner,omitempty"`
}

func (x *GetAWSResourceScannerResponse) Reset() {
	*x = GetAWSResourceScannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAWSResourceScannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAWSResourceScannerResponse) ProtoMessage() {}

func (x *GetAWSResourceScannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAWSResourceScannerResponse.ProtoReflect.Descriptor instead.
func (*GetAWSResourceScannerResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP(), []int{4}
}

func (x *GetAWSResourceScannerResponse) GetResourceScanner() *AWSResourceScanner {
	if x != nil {
		return x.ResourceScanner
	}
	return nil
}

type UpdateAWSResourceScannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceScanner *AWSResourceScanner `protobuf:"bytes,1,opt,name=resource_scanner,json=resourceScanner,proto3" json:"resource_scanner,omitempty"`
}

func (x *UpdateAWSResourceScannerRequest) Reset() {
	*x = UpdateAWSResourceScannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAWSResourceScannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAWSResourceScannerRequest) ProtoMessage() {}

func (x *UpdateAWSResourceScannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAWSResourceScannerRequest.ProtoReflect.Descriptor instead.
func (*UpdateAWSResourceScannerRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAWSResourceScannerRequest) GetResourceScanner() *AWSResourceScanner {
	if x != nil {
		return x.ResourceScanner
	}
	return nil
}

type UpdateAWSResourceScannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceScanner *AWSResourceScanner `protobuf:"bytes,1,opt,name=resource_scanner,json=resourceScanner,proto3" json:"resource_scanner,omitempty"`
}

func (x *UpdateAWSResourceScannerResponse) Reset() {
	*x = UpdateAWSResourceScannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAWSResourceScannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAWSResourceScannerResponse) ProtoMessage() {}

func (x *UpdateAWSResourceScannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAWSResourceScannerResponse.ProtoReflect.Descriptor instead.
func (*UpdateAWSResourceScannerResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAWSResourceScannerResponse) GetResourceScanner() *AWSResourceScanner {
	if x != nil {
		return x.ResourceScanner
	}
	return nil
}

type DeleteAWSResourceScannerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAWSResourceScannerRequest) Reset() {
	*x = DeleteAWSResourceScannerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAWSResourceScannerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAWSResourceScannerRequest) ProtoMessage() {}

func (x *DeleteAWSResourceScannerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAWSResourceScannerRequest.ProtoReflect.Descriptor instead.
func (*DeleteAWSResourceScannerRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAWSResourceScannerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAWSResourceScannerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAWSResourceScannerResponse) Reset() {
	*x = DeleteAWSResourceScannerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAWSResourceScannerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAWSResourceScannerResponse) ProtoMessage() {}

func (x *DeleteAWSResourceScannerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAWSResourceScannerResponse.ProtoReflect.Descriptor instead.
func (*DeleteAWSResourceScannerResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAWSResourceScannerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_commonfate_control_config_v1alpha1_aws_resource_scanner_proto protoreflect.FileDescriptor

var file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x1f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x46, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xde, 0x01, 0x0a,
	0x12, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x85, 0x01,
	0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x57, 0x53, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x57, 0x53,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x1f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x61,
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57,
	0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65,
	0x72, 0x22, 0x85, 0x01, 0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x1f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x20,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x32, 0xbe, 0x05, 0x0a, 0x19, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa7,
	0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x43, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa2, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x12, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x57, 0x53, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x57, 0x53,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0xa7, 0x01,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x43, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa7, 0x01, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x57, 0x53, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xbe, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x17, 0x41, 0x77,
	0x73, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x43, 0xaa,
	0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescOnce sync.Once
	file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescData = file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDesc
)

func file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescGZIP() []byte {
	file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescOnce.Do(func() {
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescData)
	})
	return file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDescData
}

var file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_goTypes = []any{
	(*CreateAWSResourceScannerRequest)(nil),  // 0: commonfate.control.config.v1alpha1.CreateAWSResourceScannerRequest
	(*AWSResourceScanner)(nil),               // 1: commonfate.control.config.v1alpha1.AWSResourceScanner
	(*CreateAWSResourceScannerResponse)(nil), // 2: commonfate.control.config.v1alpha1.CreateAWSResourceScannerResponse
	(*GetAWSResourceScannerRequest)(nil),     // 3: commonfate.control.config.v1alpha1.GetAWSResourceScannerRequest
	(*GetAWSResourceScannerResponse)(nil),    // 4: commonfate.control.config.v1alpha1.GetAWSResourceScannerResponse
	(*UpdateAWSResourceScannerRequest)(nil),  // 5: commonfate.control.config.v1alpha1.UpdateAWSResourceScannerRequest
	(*UpdateAWSResourceScannerResponse)(nil), // 6: commonfate.control.config.v1alpha1.UpdateAWSResourceScannerResponse
	(*DeleteAWSResourceScannerRequest)(nil),  // 7: commonfate.control.config.v1alpha1.DeleteAWSResourceScannerRequest
	(*DeleteAWSResourceScannerResponse)(nil), // 8: commonfate.control.config.v1alpha1.DeleteAWSResourceScannerResponse
}
var file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_depIdxs = []int32{
	1, // 0: commonfate.control.config.v1alpha1.CreateAWSResourceScannerResponse.resource_scanner:type_name -> commonfate.control.config.v1alpha1.AWSResourceScanner
	1, // 1: commonfate.control.config.v1alpha1.GetAWSResourceScannerResponse.resource_scanner:type_name -> commonfate.control.config.v1alpha1.AWSResourceScanner
	1, // 2: commonfate.control.config.v1alpha1.UpdateAWSResourceScannerRequest.resource_scanner:type_name -> commonfate.control.config.v1alpha1.AWSResourceScanner
	1, // 3: commonfate.control.config.v1alpha1.UpdateAWSResourceScannerResponse.resource_scanner:type_name -> commonfate.control.config.v1alpha1.AWSResourceScanner
	0, // 4: commonfate.control.config.v1alpha1.AWSResourceScannerService.CreateAWSResourceScanner:input_type -> commonfate.control.config.v1alpha1.CreateAWSResourceScannerRequest
	3, // 5: commonfate.control.config.v1alpha1.AWSResourceScannerService.GetAWSResourceScanner:input_type -> commonfate.control.config.v1alpha1.GetAWSResourceScannerRequest
	5, // 6: commonfate.control.config.v1alpha1.AWSResourceScannerService.UpdateAWSResourceScanner:input_type -> commonfate.control.config.v1alpha1.UpdateAWSResourceScannerRequest
	7, // 7: commonfate.control.config.v1alpha1.AWSResourceScannerService.DeleteAWSResourceScanner:input_type -> commonfate.control.config.v1alpha1.DeleteAWSResourceScannerRequest
	2, // 8: commonfate.control.config.v1alpha1.AWSResourceScannerService.CreateAWSResourceScanner:output_type -> commonfate.control.config.v1alpha1.CreateAWSResourceScannerResponse
	4, // 9: commonfate.control.config.v1alpha1.AWSResourceScannerService.GetAWSResourceScanner:output_type -> commonfate.control.config.v1alpha1.GetAWSResourceScannerResponse
	6, // 10: commonfate.control.config.v1alpha1.AWSResourceScannerService.UpdateAWSResourceScanner:output_type -> commonfate.control.config.v1alpha1.UpdateAWSResourceScannerResponse
	8, // 11: commonfate.control.config.v1alpha1.AWSResourceScannerService.DeleteAWSResourceScanner:output_type -> commonfate.control.config.v1alpha1.DeleteAWSResourceScannerResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_init() }
func file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_init() {
	if File_commonfate_control_config_v1alpha1_aws_resource_scanner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAWSResourceScannerRequest); i {
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
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AWSResourceScanner); i {
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
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAWSResourceScannerResponse); i {
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
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAWSResourceScannerRequest); i {
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
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAWSResourceScannerResponse); i {
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
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAWSResourceScannerRequest); i {
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
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAWSResourceScannerResponse); i {
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
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAWSResourceScannerRequest); i {
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
		file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteAWSResourceScannerResponse); i {
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
			RawDescriptor: file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_goTypes,
		DependencyIndexes: file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_depIdxs,
		MessageInfos:      file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_msgTypes,
	}.Build()
	File_commonfate_control_config_v1alpha1_aws_resource_scanner_proto = out.File
	file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_rawDesc = nil
	file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_goTypes = nil
	file_commonfate_control_config_v1alpha1_aws_resource_scanner_proto_depIdxs = nil
}
