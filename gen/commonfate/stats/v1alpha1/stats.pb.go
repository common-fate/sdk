// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: commonfate/stats/v1alpha1/stats.proto

package statsv1alpha1

import (
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
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

type GetStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_stats_v1alpha1_stats_proto_rawDescGZIP(), []int{0}
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats *Stats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_stats_v1alpha1_stats_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatsResponse) GetStats() *Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtectedUsers int32 `protobuf:"varint,1,opt,name=protected_users,json=protectedUsers,proto3" json:"protected_users,omitempty"`
	// from 0 - 100
	AccessHoursReductionPercent float32                     `protobuf:"fixed32,2,opt,name=access_hours_reduction_percent,json=accessHoursReductionPercent,proto3" json:"access_hours_reduction_percent,omitempty"`
	AccessHoursReduced          int32                       `protobuf:"varint,3,opt,name=access_hours_reduced,json=accessHoursReduced,proto3" json:"access_hours_reduced,omitempty"`
	AccessRequestsLast_90Days   []*AccessRequestDailyMetric `protobuf:"bytes,4,rep,name=access_requests_last_90_days,json=accessRequestsLast90Days,proto3" json:"access_requests_last_90_days,omitempty"`
	EntitlementUsageLast_90Days []*EntitlementUsageMetric   `protobuf:"bytes,5,rep,name=entitlement_usage_last_90_days,json=entitlementUsageLast90Days,proto3" json:"entitlement_usage_last_90_days,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_commonfate_stats_v1alpha1_stats_proto_rawDescGZIP(), []int{2}
}

func (x *Stats) GetProtectedUsers() int32 {
	if x != nil {
		return x.ProtectedUsers
	}
	return 0
}

func (x *Stats) GetAccessHoursReductionPercent() float32 {
	if x != nil {
		return x.AccessHoursReductionPercent
	}
	return 0
}

func (x *Stats) GetAccessHoursReduced() int32 {
	if x != nil {
		return x.AccessHoursReduced
	}
	return 0
}

func (x *Stats) GetAccessRequestsLast_90Days() []*AccessRequestDailyMetric {
	if x != nil {
		return x.AccessRequestsLast_90Days
	}
	return nil
}

func (x *Stats) GetEntitlementUsageLast_90Days() []*EntitlementUsageMetric {
	if x != nil {
		return x.EntitlementUsageLast_90Days
	}
	return nil
}

type AccessRequestDailyMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date  string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AccessRequestDailyMetric) Reset() {
	*x = AccessRequestDailyMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequestDailyMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequestDailyMetric) ProtoMessage() {}

func (x *AccessRequestDailyMetric) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequestDailyMetric.ProtoReflect.Descriptor instead.
func (*AccessRequestDailyMetric) Descriptor() ([]byte, []int) {
	return file_commonfate_stats_v1alpha1_stats_proto_rawDescGZIP(), []int{3}
}

func (x *AccessRequestDailyMetric) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *AccessRequestDailyMetric) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type EntitlementUsageMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *v1alpha1.NamedEID `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Role   *v1alpha1.NamedEID `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Count  int32              `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *EntitlementUsageMetric) Reset() {
	*x = EntitlementUsageMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntitlementUsageMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitlementUsageMetric) ProtoMessage() {}

func (x *EntitlementUsageMetric) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_stats_v1alpha1_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitlementUsageMetric.ProtoReflect.Descriptor instead.
func (*EntitlementUsageMetric) Descriptor() ([]byte, []int) {
	return file_commonfate_stats_v1alpha1_stats_proto_rawDescGZIP(), []int{4}
}

func (x *EntitlementUsageMetric) GetTarget() *v1alpha1.NamedEID {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *EntitlementUsageMetric) GetRole() *v1alpha1.NamedEID {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *EntitlementUsageMetric) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_commonfate_stats_v1alpha1_stats_proto protoreflect.FileDescriptor

var file_commonfate_stats_v1alpha1_stats_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e,
	0x61, 0x6d, 0x65, 0x64, 0x5f, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x93, 0x03, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x70, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x1b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73,
	0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x12, 0x30, 0x0a, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x64, 0x12, 0x73, 0x0a, 0x1c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x39, 0x30, 0x5f, 0x64, 0x61,
	0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x18, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x4c, 0x61, 0x73,
	0x74, 0x39, 0x30, 0x44, 0x61, 0x79, 0x73, 0x12, 0x75, 0x0a, 0x1e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x39, 0x30, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x1a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x39, 0x30, 0x44, 0x61, 0x79, 0x73, 0x22, 0x44,
	0x0a, 0x18, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa6, 0x01, 0x0a, 0x16, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12,
	0x3c, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x64, 0x45, 0x49, 0x44, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x38, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x45, 0x49,
	0x44, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x79, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x42, 0xf9, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_stats_v1alpha1_stats_proto_rawDescOnce sync.Once
	file_commonfate_stats_v1alpha1_stats_proto_rawDescData = file_commonfate_stats_v1alpha1_stats_proto_rawDesc
)

func file_commonfate_stats_v1alpha1_stats_proto_rawDescGZIP() []byte {
	file_commonfate_stats_v1alpha1_stats_proto_rawDescOnce.Do(func() {
		file_commonfate_stats_v1alpha1_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_stats_v1alpha1_stats_proto_rawDescData)
	})
	return file_commonfate_stats_v1alpha1_stats_proto_rawDescData
}

var file_commonfate_stats_v1alpha1_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_commonfate_stats_v1alpha1_stats_proto_goTypes = []any{
	(*GetStatsRequest)(nil),          // 0: commonfate.stats.v1alpha1.GetStatsRequest
	(*GetStatsResponse)(nil),         // 1: commonfate.stats.v1alpha1.GetStatsResponse
	(*Stats)(nil),                    // 2: commonfate.stats.v1alpha1.Stats
	(*AccessRequestDailyMetric)(nil), // 3: commonfate.stats.v1alpha1.AccessRequestDailyMetric
	(*EntitlementUsageMetric)(nil),   // 4: commonfate.stats.v1alpha1.EntitlementUsageMetric
	(*v1alpha1.NamedEID)(nil),        // 5: commonfate.access.v1alpha1.NamedEID
}
var file_commonfate_stats_v1alpha1_stats_proto_depIdxs = []int32{
	2, // 0: commonfate.stats.v1alpha1.GetStatsResponse.stats:type_name -> commonfate.stats.v1alpha1.Stats
	3, // 1: commonfate.stats.v1alpha1.Stats.access_requests_last_90_days:type_name -> commonfate.stats.v1alpha1.AccessRequestDailyMetric
	4, // 2: commonfate.stats.v1alpha1.Stats.entitlement_usage_last_90_days:type_name -> commonfate.stats.v1alpha1.EntitlementUsageMetric
	5, // 3: commonfate.stats.v1alpha1.EntitlementUsageMetric.target:type_name -> commonfate.access.v1alpha1.NamedEID
	5, // 4: commonfate.stats.v1alpha1.EntitlementUsageMetric.role:type_name -> commonfate.access.v1alpha1.NamedEID
	0, // 5: commonfate.stats.v1alpha1.StatsService.GetStats:input_type -> commonfate.stats.v1alpha1.GetStatsRequest
	1, // 6: commonfate.stats.v1alpha1.StatsService.GetStats:output_type -> commonfate.stats.v1alpha1.GetStatsResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_commonfate_stats_v1alpha1_stats_proto_init() }
func file_commonfate_stats_v1alpha1_stats_proto_init() {
	if File_commonfate_stats_v1alpha1_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_stats_v1alpha1_stats_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetStatsRequest); i {
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
		file_commonfate_stats_v1alpha1_stats_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetStatsResponse); i {
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
		file_commonfate_stats_v1alpha1_stats_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Stats); i {
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
		file_commonfate_stats_v1alpha1_stats_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AccessRequestDailyMetric); i {
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
		file_commonfate_stats_v1alpha1_stats_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EntitlementUsageMetric); i {
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
			RawDescriptor: file_commonfate_stats_v1alpha1_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_stats_v1alpha1_stats_proto_goTypes,
		DependencyIndexes: file_commonfate_stats_v1alpha1_stats_proto_depIdxs,
		MessageInfos:      file_commonfate_stats_v1alpha1_stats_proto_msgTypes,
	}.Build()
	File_commonfate_stats_v1alpha1_stats_proto = out.File
	file_commonfate_stats_v1alpha1_stats_proto_rawDesc = nil
	file_commonfate_stats_v1alpha1_stats_proto_goTypes = nil
	file_commonfate_stats_v1alpha1_stats_proto_depIdxs = nil
}
