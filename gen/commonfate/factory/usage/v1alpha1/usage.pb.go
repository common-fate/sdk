// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: commonfate/factory/usage/v1alpha1/usage.proto

package usagev1alpha1

import (
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

type ReportUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Integrations   []*IntegrationUsage `protobuf:"bytes,3,rep,name=integrations,proto3" json:"integrations,omitempty"`
	Users          *UserUsage          `protobuf:"bytes,4,opt,name=users,proto3" json:"users,omitempty"`
	AccessRequests *AccessRequestUsage `protobuf:"bytes,5,opt,name=access_requests,json=accessRequests,proto3" json:"access_requests,omitempty"`
	Configuration  *ConfigurationUsage `protobuf:"bytes,6,opt,name=configuration,proto3" json:"configuration,omitempty"`
}

func (x *ReportUsageRequest) Reset() {
	*x = ReportUsageRequest{}
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportUsageRequest) ProtoMessage() {}

func (x *ReportUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportUsageRequest.ProtoReflect.Descriptor instead.
func (*ReportUsageRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescGZIP(), []int{0}
}

func (x *ReportUsageRequest) GetIntegrations() []*IntegrationUsage {
	if x != nil {
		return x.Integrations
	}
	return nil
}

func (x *ReportUsageRequest) GetUsers() *UserUsage {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ReportUsageRequest) GetAccessRequests() *AccessRequestUsage {
	if x != nil {
		return x.AccessRequests
	}
	return nil
}

func (x *ReportUsageRequest) GetConfiguration() *ConfigurationUsage {
	if x != nil {
		return x.Configuration
	}
	return nil
}

type UserUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalUserCount       int32 `protobuf:"varint,1,opt,name=total_user_count,json=totalUserCount,proto3" json:"total_user_count,omitempty"`
	ActiveUserCount      int32 `protobuf:"varint,2,opt,name=active_user_count,json=activeUserCount,proto3" json:"active_user_count,omitempty"`
	Requestors_30DCount  int32 `protobuf:"varint,3,opt,name=requestors_30d_count,json=requestors30dCount,proto3" json:"requestors_30d_count,omitempty"`
	TotalRequestorsCount int32 `protobuf:"varint,4,opt,name=total_requestors_count,json=totalRequestorsCount,proto3" json:"total_requestors_count,omitempty"`
}

func (x *UserUsage) Reset() {
	*x = UserUsage{}
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUsage) ProtoMessage() {}

func (x *UserUsage) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUsage.ProtoReflect.Descriptor instead.
func (*UserUsage) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescGZIP(), []int{1}
}

func (x *UserUsage) GetTotalUserCount() int32 {
	if x != nil {
		return x.TotalUserCount
	}
	return 0
}

func (x *UserUsage) GetActiveUserCount() int32 {
	if x != nil {
		return x.ActiveUserCount
	}
	return 0
}

func (x *UserUsage) GetRequestors_30DCount() int32 {
	if x != nil {
		return x.Requestors_30DCount
	}
	return 0
}

func (x *UserUsage) GetTotalRequestorsCount() int32 {
	if x != nil {
		return x.TotalRequestorsCount
	}
	return 0
}

type AccessRequestUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalAccessRequestCount          int32         `protobuf:"varint,1,opt,name=total_access_request_count,json=totalAccessRequestCount,proto3" json:"total_access_request_count,omitempty"`
	AccessRequest_30DCount           int32         `protobuf:"varint,2,opt,name=access_request_30d_count,json=accessRequest30dCount,proto3" json:"access_request_30d_count,omitempty"`
	AccessHoursReductionPercent_30D  float32       `protobuf:"fixed32,3,opt,name=access_hours_reduction_percent_30d,json=accessHoursReductionPercent30d,proto3" json:"access_hours_reduction_percent_30d,omitempty"`
	AccessHoursReduction_30D         int32         `protobuf:"varint,4,opt,name=access_hours_reduction_30d,json=accessHoursReduction30d,proto3" json:"access_hours_reduction_30d,omitempty"`
	ManualApproval_30DCount          int32         `protobuf:"varint,5,opt,name=manual_approval_30d_count,json=manualApproval30dCount,proto3" json:"manual_approval_30d_count,omitempty"`
	RequestDuration_30DDistribution  *Distribution `protobuf:"bytes,6,opt,name=request_duration_30d_distribution,json=requestDuration30dDistribution,proto3" json:"request_duration_30d_distribution,omitempty"`
	ApprovalWaitTime_30DDistribution *Distribution `protobuf:"bytes,7,opt,name=approval_wait_time_30d_distribution,json=approvalWaitTime30dDistribution,proto3" json:"approval_wait_time_30d_distribution,omitempty"`
}

func (x *AccessRequestUsage) Reset() {
	*x = AccessRequestUsage{}
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessRequestUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequestUsage) ProtoMessage() {}

func (x *AccessRequestUsage) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequestUsage.ProtoReflect.Descriptor instead.
func (*AccessRequestUsage) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescGZIP(), []int{2}
}

func (x *AccessRequestUsage) GetTotalAccessRequestCount() int32 {
	if x != nil {
		return x.TotalAccessRequestCount
	}
	return 0
}

func (x *AccessRequestUsage) GetAccessRequest_30DCount() int32 {
	if x != nil {
		return x.AccessRequest_30DCount
	}
	return 0
}

func (x *AccessRequestUsage) GetAccessHoursReductionPercent_30D() float32 {
	if x != nil {
		return x.AccessHoursReductionPercent_30D
	}
	return 0
}

func (x *AccessRequestUsage) GetAccessHoursReduction_30D() int32 {
	if x != nil {
		return x.AccessHoursReduction_30D
	}
	return 0
}

func (x *AccessRequestUsage) GetManualApproval_30DCount() int32 {
	if x != nil {
		return x.ManualApproval_30DCount
	}
	return 0
}

func (x *AccessRequestUsage) GetRequestDuration_30DDistribution() *Distribution {
	if x != nil {
		return x.RequestDuration_30DDistribution
	}
	return nil
}

func (x *AccessRequestUsage) GetApprovalWaitTime_30DDistribution() *Distribution {
	if x != nil {
		return x.ApprovalWaitTime_30DDistribution
	}
	return nil
}

type ConfigurationUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowCount         int32 `protobuf:"varint,1,opt,name=workflow_count,json=workflowCount,proto3" json:"workflow_count,omitempty"`
	AvailabilitySpecCount int32 `protobuf:"varint,2,opt,name=availability_spec_count,json=availabilitySpecCount,proto3" json:"availability_spec_count,omitempty"`
	PolicyCount           int32 `protobuf:"varint,3,opt,name=policy_count,json=policyCount,proto3" json:"policy_count,omitempty"`
	SlackAlertCount       int32 `protobuf:"varint,4,opt,name=slack_alert_count,json=slackAlertCount,proto3" json:"slack_alert_count,omitempty"`
	EntitlementCount      int32 `protobuf:"varint,5,opt,name=entitlement_count,json=entitlementCount,proto3" json:"entitlement_count,omitempty"`
}

func (x *ConfigurationUsage) Reset() {
	*x = ConfigurationUsage{}
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigurationUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurationUsage) ProtoMessage() {}

func (x *ConfigurationUsage) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurationUsage.ProtoReflect.Descriptor instead.
func (*ConfigurationUsage) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigurationUsage) GetWorkflowCount() int32 {
	if x != nil {
		return x.WorkflowCount
	}
	return 0
}

func (x *ConfigurationUsage) GetAvailabilitySpecCount() int32 {
	if x != nil {
		return x.AvailabilitySpecCount
	}
	return 0
}

func (x *ConfigurationUsage) GetPolicyCount() int32 {
	if x != nil {
		return x.PolicyCount
	}
	return 0
}

func (x *ConfigurationUsage) GetSlackAlertCount() int32 {
	if x != nil {
		return x.SlackAlertCount
	}
	return 0
}

func (x *ConfigurationUsage) GetEntitlementCount() int32 {
	if x != nil {
		return x.EntitlementCount
	}
	return 0
}

type IntegrationUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId    string `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	IntegrationType  string `protobuf:"bytes,2,opt,name=integration_type,json=integrationType,proto3" json:"integration_type,omitempty"`
	UserAccountCount int32  `protobuf:"varint,3,opt,name=user_account_count,json=userAccountCount,proto3" json:"user_account_count,omitempty"`
	GroupCount       int32  `protobuf:"varint,4,opt,name=group_count,json=groupCount,proto3" json:"group_count,omitempty"`
	TargetCount      int32  `protobuf:"varint,5,opt,name=target_count,json=targetCount,proto3" json:"target_count,omitempty"`
}

func (x *IntegrationUsage) Reset() {
	*x = IntegrationUsage{}
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntegrationUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationUsage) ProtoMessage() {}

func (x *IntegrationUsage) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationUsage.ProtoReflect.Descriptor instead.
func (*IntegrationUsage) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescGZIP(), []int{4}
}

func (x *IntegrationUsage) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *IntegrationUsage) GetIntegrationType() string {
	if x != nil {
		return x.IntegrationType
	}
	return ""
}

func (x *IntegrationUsage) GetUserAccountCount() int32 {
	if x != nil {
		return x.UserAccountCount
	}
	return 0
}

func (x *IntegrationUsage) GetGroupCount() int32 {
	if x != nil {
		return x.GroupCount
	}
	return 0
}

func (x *IntegrationUsage) GetTargetCount() int32 {
	if x != nil {
		return x.TargetCount
	}
	return 0
}

type Distribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min int32 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	P25 int32 `protobuf:"varint,2,opt,name=p25,proto3" json:"p25,omitempty"`
	P50 int32 `protobuf:"varint,3,opt,name=p50,proto3" json:"p50,omitempty"`
	P75 int32 `protobuf:"varint,4,opt,name=p75,proto3" json:"p75,omitempty"`
	P95 int32 `protobuf:"varint,5,opt,name=p95,proto3" json:"p95,omitempty"`
	Max int32 `protobuf:"varint,6,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *Distribution) Reset() {
	*x = Distribution{}
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Distribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution) ProtoMessage() {}

func (x *Distribution) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution.ProtoReflect.Descriptor instead.
func (*Distribution) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescGZIP(), []int{5}
}

func (x *Distribution) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Distribution) GetP25() int32 {
	if x != nil {
		return x.P25
	}
	return 0
}

func (x *Distribution) GetP50() int32 {
	if x != nil {
		return x.P50
	}
	return 0
}

func (x *Distribution) GetP75() int32 {
	if x != nil {
		return x.P75
	}
	return 0
}

func (x *Distribution) GetP95() int32 {
	if x != nil {
		return x.P95
	}
	return 0
}

func (x *Distribution) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

type ReportUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportUsageResponse) Reset() {
	*x = ReportUsageResponse{}
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportUsageResponse) ProtoMessage() {}

func (x *ReportUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportUsageResponse.ProtoReflect.Descriptor instead.
func (*ReportUsageResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescGZIP(), []int{6}
}

var File_commonfate_factory_usage_v1alpha1_usage_proto protoreflect.FileDescriptor

var file_commonfate_factory_usage_v1alpha1_usage_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xee, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x0c, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x42, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x5e, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x5b, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xc9, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x33, 0x30, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x73, 0x33, 0x30, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xc9, 0x04, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x33, 0x30, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x33, 0x30, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x22,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x33,
	0x30, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x1e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x33, 0x30, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x33, 0x30, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x33, 0x30, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x33, 0x30, 0x64, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x33, 0x30, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x7a, 0x0a, 0x21, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x33, 0x30, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x33, 0x30, 0x64,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7d, 0x0a, 0x23,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x33, 0x30, 0x64, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1f, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x57, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x33, 0x30, 0x64, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xef, 0x01, 0x0a, 0x12,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x11, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd6, 0x01,
	0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x32, 0x35, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x32, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x35,
	0x30, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x35, 0x30, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x37, 0x35, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x37, 0x35, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x39, 0x35, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x39, 0x35,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d,
	0x61, 0x78, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8e, 0x01, 0x0a, 0x0c, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x0b, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xaa, 0x02, 0x0a, 0x25, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x55, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x75, 0x73, 0x61, 0x67, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x46, 0x55, 0xaa, 0x02, 0x21, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x21, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x5c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x2d, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x46, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x55, 0x73, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x24, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x3a, 0x55, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescOnce sync.Once
	file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescData = file_commonfate_factory_usage_v1alpha1_usage_proto_rawDesc
)

func file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescGZIP() []byte {
	file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescOnce.Do(func() {
		file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescData)
	})
	return file_commonfate_factory_usage_v1alpha1_usage_proto_rawDescData
}

var file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_commonfate_factory_usage_v1alpha1_usage_proto_goTypes = []any{
	(*ReportUsageRequest)(nil),  // 0: commonfate.factory.usage.v1alpha1.ReportUsageRequest
	(*UserUsage)(nil),           // 1: commonfate.factory.usage.v1alpha1.UserUsage
	(*AccessRequestUsage)(nil),  // 2: commonfate.factory.usage.v1alpha1.AccessRequestUsage
	(*ConfigurationUsage)(nil),  // 3: commonfate.factory.usage.v1alpha1.ConfigurationUsage
	(*IntegrationUsage)(nil),    // 4: commonfate.factory.usage.v1alpha1.IntegrationUsage
	(*Distribution)(nil),        // 5: commonfate.factory.usage.v1alpha1.Distribution
	(*ReportUsageResponse)(nil), // 6: commonfate.factory.usage.v1alpha1.ReportUsageResponse
}
var file_commonfate_factory_usage_v1alpha1_usage_proto_depIdxs = []int32{
	4, // 0: commonfate.factory.usage.v1alpha1.ReportUsageRequest.integrations:type_name -> commonfate.factory.usage.v1alpha1.IntegrationUsage
	1, // 1: commonfate.factory.usage.v1alpha1.ReportUsageRequest.users:type_name -> commonfate.factory.usage.v1alpha1.UserUsage
	2, // 2: commonfate.factory.usage.v1alpha1.ReportUsageRequest.access_requests:type_name -> commonfate.factory.usage.v1alpha1.AccessRequestUsage
	3, // 3: commonfate.factory.usage.v1alpha1.ReportUsageRequest.configuration:type_name -> commonfate.factory.usage.v1alpha1.ConfigurationUsage
	5, // 4: commonfate.factory.usage.v1alpha1.AccessRequestUsage.request_duration_30d_distribution:type_name -> commonfate.factory.usage.v1alpha1.Distribution
	5, // 5: commonfate.factory.usage.v1alpha1.AccessRequestUsage.approval_wait_time_30d_distribution:type_name -> commonfate.factory.usage.v1alpha1.Distribution
	0, // 6: commonfate.factory.usage.v1alpha1.UsageService.ReportUsage:input_type -> commonfate.factory.usage.v1alpha1.ReportUsageRequest
	6, // 7: commonfate.factory.usage.v1alpha1.UsageService.ReportUsage:output_type -> commonfate.factory.usage.v1alpha1.ReportUsageResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_commonfate_factory_usage_v1alpha1_usage_proto_init() }
func file_commonfate_factory_usage_v1alpha1_usage_proto_init() {
	if File_commonfate_factory_usage_v1alpha1_usage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_factory_usage_v1alpha1_usage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_factory_usage_v1alpha1_usage_proto_goTypes,
		DependencyIndexes: file_commonfate_factory_usage_v1alpha1_usage_proto_depIdxs,
		MessageInfos:      file_commonfate_factory_usage_v1alpha1_usage_proto_msgTypes,
	}.Build()
	File_commonfate_factory_usage_v1alpha1_usage_proto = out.File
	file_commonfate_factory_usage_v1alpha1_usage_proto_rawDesc = nil
	file_commonfate_factory_usage_v1alpha1_usage_proto_goTypes = nil
	file_commonfate_factory_usage_v1alpha1_usage_proto_depIdxs = nil
}
