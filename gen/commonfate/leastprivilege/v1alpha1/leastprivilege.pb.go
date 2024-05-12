// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: commonfate/leastprivilege/v1alpha1/leastprivilege.proto

package leastprivilegev1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
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

type GetLatestReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLatestReportRequest) Reset() {
	*x = GetLatestReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestReportRequest) ProtoMessage() {}

func (x *GetLatestReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestReportRequest.ProtoReflect.Descriptor instead.
func (*GetLatestReportRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescGZIP(), []int{0}
}

type GetLatestReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *GetLatestReportResponse) Reset() {
	*x = GetLatestReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestReportResponse) ProtoMessage() {}

func (x *GetLatestReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestReportResponse.ProtoReflect.Descriptor instead.
func (*GetLatestReportResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescGZIP(), []int{1}
}

func (x *GetLatestReportResponse) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type GetHistoricalReportsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHistoricalReportsRequest) Reset() {
	*x = GetHistoricalReportsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoricalReportsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoricalReportsRequest) ProtoMessage() {}

func (x *GetHistoricalReportsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoricalReportsRequest.ProtoReflect.Descriptor instead.
func (*GetHistoricalReportsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescGZIP(), []int{2}
}

type GetHistoricalReportsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reports []*Report `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
}

func (x *GetHistoricalReportsResponse) Reset() {
	*x = GetHistoricalReportsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoricalReportsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoricalReportsResponse) ProtoMessage() {}

func (x *GetHistoricalReportsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoricalReportsResponse.ProtoReflect.Descriptor instead.
func (*GetHistoricalReportsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescGZIP(), []int{3}
}

func (x *GetHistoricalReportsResponse) GetReports() []*Report {
	if x != nil {
		return x.Reports
	}
	return nil
}

type DownloadEntitlementUsageReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DownloadEntitlementUsageReportRequest) Reset() {
	*x = DownloadEntitlementUsageReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadEntitlementUsageReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadEntitlementUsageReportRequest) ProtoMessage() {}

func (x *DownloadEntitlementUsageReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadEntitlementUsageReportRequest.ProtoReflect.Descriptor instead.
func (*DownloadEntitlementUsageReportRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescGZIP(), []int{4}
}

type DownloadEntitlementUsageReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DownloadUrl string `protobuf:"bytes,1,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
}

func (x *DownloadEntitlementUsageReportResponse) Reset() {
	*x = DownloadEntitlementUsageReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadEntitlementUsageReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadEntitlementUsageReportResponse) ProtoMessage() {}

func (x *DownloadEntitlementUsageReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadEntitlementUsageReportResponse.ProtoReflect.Descriptor instead.
func (*DownloadEntitlementUsageReportResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadEntitlementUsageReportResponse) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

var File_commonfate_leastprivilege_v1alpha1_leastprivilege_proto protoreflect.FileDescriptor

var file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c,
	0x65, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3a,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x74,
	0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x5d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72,
	0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x64,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x07, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x07, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x22, 0x27, 0x0a, 0x25, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4b, 0x0a,
	0x26, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x32, 0x84, 0x04, 0x0a, 0x0d, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x90, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65,
	0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12,
	0x9f, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61,
	0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18,
	0x01, 0x12, 0xbd, 0x01, 0x0a, 0x1e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x49, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x4a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18,
	0x01, 0x42, 0xc1, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c,
	0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x13, 0x4c, 0x65,
	0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x6c,
	0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x4c, 0x58, 0xaa, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x4c, 0x65, 0x61, 0x73, 0x74,
	0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x24, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x4c, 0x65, 0x61,
	0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescOnce sync.Once
	file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescData = file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDesc
)

func file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescGZIP() []byte {
	file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescOnce.Do(func() {
		file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescData)
	})
	return file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescData
}

var file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_goTypes = []interface{}{
	(*GetLatestReportRequest)(nil),                 // 0: commonfate.leastprivilege.v1alpha1.GetLatestReportRequest
	(*GetLatestReportResponse)(nil),                // 1: commonfate.leastprivilege.v1alpha1.GetLatestReportResponse
	(*GetHistoricalReportsRequest)(nil),            // 2: commonfate.leastprivilege.v1alpha1.GetHistoricalReportsRequest
	(*GetHistoricalReportsResponse)(nil),           // 3: commonfate.leastprivilege.v1alpha1.GetHistoricalReportsResponse
	(*DownloadEntitlementUsageReportRequest)(nil),  // 4: commonfate.leastprivilege.v1alpha1.DownloadEntitlementUsageReportRequest
	(*DownloadEntitlementUsageReportResponse)(nil), // 5: commonfate.leastprivilege.v1alpha1.DownloadEntitlementUsageReportResponse
	(*Report)(nil), // 6: commonfate.leastprivilege.v1alpha1.Report
}
var file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_depIdxs = []int32{
	6, // 0: commonfate.leastprivilege.v1alpha1.GetLatestReportResponse.report:type_name -> commonfate.leastprivilege.v1alpha1.Report
	6, // 1: commonfate.leastprivilege.v1alpha1.GetHistoricalReportsResponse.reports:type_name -> commonfate.leastprivilege.v1alpha1.Report
	0, // 2: commonfate.leastprivilege.v1alpha1.ReportService.GetLatestReport:input_type -> commonfate.leastprivilege.v1alpha1.GetLatestReportRequest
	2, // 3: commonfate.leastprivilege.v1alpha1.ReportService.GetHistoricalReports:input_type -> commonfate.leastprivilege.v1alpha1.GetHistoricalReportsRequest
	4, // 4: commonfate.leastprivilege.v1alpha1.ReportService.DownloadEntitlementUsageReport:input_type -> commonfate.leastprivilege.v1alpha1.DownloadEntitlementUsageReportRequest
	1, // 5: commonfate.leastprivilege.v1alpha1.ReportService.GetLatestReport:output_type -> commonfate.leastprivilege.v1alpha1.GetLatestReportResponse
	3, // 6: commonfate.leastprivilege.v1alpha1.ReportService.GetHistoricalReports:output_type -> commonfate.leastprivilege.v1alpha1.GetHistoricalReportsResponse
	5, // 7: commonfate.leastprivilege.v1alpha1.ReportService.DownloadEntitlementUsageReport:output_type -> commonfate.leastprivilege.v1alpha1.DownloadEntitlementUsageReportResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_init() }
func file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_init() {
	if File_commonfate_leastprivilege_v1alpha1_leastprivilege_proto != nil {
		return
	}
	file_commonfate_leastprivilege_v1alpha1_entitlement_usage_proto_init()
	file_commonfate_leastprivilege_v1alpha1_report_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestReportRequest); i {
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
		file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestReportResponse); i {
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
		file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoricalReportsRequest); i {
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
		file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoricalReportsResponse); i {
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
		file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadEntitlementUsageReportRequest); i {
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
		file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadEntitlementUsageReportResponse); i {
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
			RawDescriptor: file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_goTypes,
		DependencyIndexes: file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_depIdxs,
		MessageInfos:      file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes,
	}.Build()
	File_commonfate_leastprivilege_v1alpha1_leastprivilege_proto = out.File
	file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDesc = nil
	file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_goTypes = nil
	file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_depIdxs = nil
}
