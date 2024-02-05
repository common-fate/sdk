// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: commonfate/leastprivilege/v1alpha1/leastprivilege.proto

package leastprivilegev1alpha1

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

type QueryLatestEntitlementUsagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *QueryLatestEntitlementUsagesRequest) Reset() {
	*x = QueryLatestEntitlementUsagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLatestEntitlementUsagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLatestEntitlementUsagesRequest) ProtoMessage() {}

func (x *QueryLatestEntitlementUsagesRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QueryLatestEntitlementUsagesRequest.ProtoReflect.Descriptor instead.
func (*QueryLatestEntitlementUsagesRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescGZIP(), []int{2}
}

func (x *QueryLatestEntitlementUsagesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type QueryLatestEntitlementUsagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntitlementUsages []*EntitlementUsage `protobuf:"bytes,1,rep,name=entitlement_usages,json=entitlementUsages,proto3" json:"entitlement_usages,omitempty"`
	NextPageToken     string              `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *QueryLatestEntitlementUsagesResponse) Reset() {
	*x = QueryLatestEntitlementUsagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryLatestEntitlementUsagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryLatestEntitlementUsagesResponse) ProtoMessage() {}

func (x *QueryLatestEntitlementUsagesResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QueryLatestEntitlementUsagesResponse.ProtoReflect.Descriptor instead.
func (*QueryLatestEntitlementUsagesResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_rawDescGZIP(), []int{3}
}

func (x *QueryLatestEntitlementUsagesResponse) GetEntitlementUsages() []*EntitlementUsage {
	if x != nil {
		return x.EntitlementUsages
	}
	return nil
}

func (x *QueryLatestEntitlementUsagesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
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
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65,
	0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x22, 0x44, 0x0a, 0x23, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb3, 0x01, 0x0a, 0x24, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x12, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x11, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x32, 0xd4, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0xb3, 0x01, 0x0a, 0x1c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x47, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72,
	0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xc1, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74,
	0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x13, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c,
	0x65, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x4c, 0x58, 0xaa, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x4c, 0x65, 0x61, 0x73,
	0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x5c, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x3a, 0x3a, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_goTypes = []interface{}{
	(*GetLatestReportRequest)(nil),               // 0: commonfate.leastprivilege.v1alpha1.GetLatestReportRequest
	(*GetLatestReportResponse)(nil),              // 1: commonfate.leastprivilege.v1alpha1.GetLatestReportResponse
	(*QueryLatestEntitlementUsagesRequest)(nil),  // 2: commonfate.leastprivilege.v1alpha1.QueryLatestEntitlementUsagesRequest
	(*QueryLatestEntitlementUsagesResponse)(nil), // 3: commonfate.leastprivilege.v1alpha1.QueryLatestEntitlementUsagesResponse
	(*Report)(nil),           // 4: commonfate.leastprivilege.v1alpha1.Report
	(*EntitlementUsage)(nil), // 5: commonfate.leastprivilege.v1alpha1.EntitlementUsage
}
var file_commonfate_leastprivilege_v1alpha1_leastprivilege_proto_depIdxs = []int32{
	4, // 0: commonfate.leastprivilege.v1alpha1.GetLatestReportResponse.report:type_name -> commonfate.leastprivilege.v1alpha1.Report
	5, // 1: commonfate.leastprivilege.v1alpha1.QueryLatestEntitlementUsagesResponse.entitlement_usages:type_name -> commonfate.leastprivilege.v1alpha1.EntitlementUsage
	0, // 2: commonfate.leastprivilege.v1alpha1.ReportService.GetLatestReport:input_type -> commonfate.leastprivilege.v1alpha1.GetLatestReportRequest
	2, // 3: commonfate.leastprivilege.v1alpha1.ReportService.QueryLatestEntitlementUsages:input_type -> commonfate.leastprivilege.v1alpha1.QueryLatestEntitlementUsagesRequest
	1, // 4: commonfate.leastprivilege.v1alpha1.ReportService.GetLatestReport:output_type -> commonfate.leastprivilege.v1alpha1.GetLatestReportResponse
	3, // 5: commonfate.leastprivilege.v1alpha1.ReportService.QueryLatestEntitlementUsages:output_type -> commonfate.leastprivilege.v1alpha1.QueryLatestEntitlementUsagesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
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
			switch v := v.(*QueryLatestEntitlementUsagesRequest); i {
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
			switch v := v.(*QueryLatestEntitlementUsagesResponse); i {
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
			NumMessages:   4,
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
