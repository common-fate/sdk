// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: commonfate/access/v1alpha1/audit_logs.proto

package accessv1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryAuditLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Query for audit logs relating to a particular target
	Target *v1alpha1.EID `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The token for the next page.
	PageToken string `protobuf:"bytes,10,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *QueryAuditLogsRequest) Reset() {
	*x = QueryAuditLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAuditLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAuditLogsRequest) ProtoMessage() {}

func (x *QueryAuditLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAuditLogsRequest.ProtoReflect.Descriptor instead.
func (*QueryAuditLogsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_audit_logs_proto_rawDescGZIP(), []int{0}
}

func (x *QueryAuditLogsRequest) GetTarget() *v1alpha1.EID {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *QueryAuditLogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type QueryAuditLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuditLogs     []*AuditLog `protobuf:"bytes,1,rep,name=audit_logs,json=auditLogs,proto3" json:"audit_logs,omitempty"`
	NextPageToken string      `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *QueryAuditLogsResponse) Reset() {
	*x = QueryAuditLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAuditLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAuditLogsResponse) ProtoMessage() {}

func (x *QueryAuditLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAuditLogsResponse.ProtoReflect.Descriptor instead.
func (*QueryAuditLogsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_audit_logs_proto_rawDescGZIP(), []int{1}
}

func (x *QueryAuditLogsResponse) GetAuditLogs() []*AuditLog {
	if x != nil {
		return x.AuditLogs
	}
	return nil
}

func (x *QueryAuditLogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type AuditLogPreview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*AuditLog `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
	// An approximate count of activity events
	LogCount uint32 `protobuf:"varint,12,opt,name=log_count,json=logCount,proto3" json:"log_count,omitempty"`
}

func (x *AuditLogPreview) Reset() {
	*x = AuditLogPreview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLogPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLogPreview) ProtoMessage() {}

func (x *AuditLogPreview) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLogPreview.ProtoReflect.Descriptor instead.
func (*AuditLogPreview) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_audit_logs_proto_rawDescGZIP(), []int{2}
}

func (x *AuditLogPreview) GetLogs() []*AuditLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *AuditLogPreview) GetLogCount() uint32 {
	if x != nil {
		return x.LogCount
	}
	return 0
}

type AuditLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// an ID of the audit log event
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the action which occurred
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	// the actor which performed the action.
	Actor *User `protobuf:"bytes,3,opt,name=actor,proto3" json:"actor,omitempty"`
	// the timestamp the action occurred at.
	OccurredAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// all related resources affected by the action.
	Targets []*v1alpha1.EID `protobuf:"bytes,5,rep,name=targets,proto3" json:"targets,omitempty"`
	// a human-friendly message describing the action.
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	// additional context (omitted in user-facing audit logs,
	// as it contains sensitive information like IP addresses)
	Context *structpb.Struct `protobuf:"bytes,7,opt,name=context,proto3" json:"context,omitempty"`
	// the caller identity details, including information
	// such as the particular authentication token ID used
	// to perform the action.
	CallerIdentityChain []*IdentityLink `protobuf:"bytes,8,rep,name=caller_identity_chain,json=callerIdentityChain,proto3" json:"caller_identity_chain,omitempty"`
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLog.ProtoReflect.Descriptor instead.
func (*AuditLog) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_audit_logs_proto_rawDescGZIP(), []int{3}
}

func (x *AuditLog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuditLog) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *AuditLog) GetActor() *User {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *AuditLog) GetOccurredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OccurredAt
	}
	return nil
}

func (x *AuditLog) GetTargets() []*v1alpha1.EID {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *AuditLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AuditLog) GetContext() *structpb.Struct {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *AuditLog) GetCallerIdentityChain() []*IdentityLink {
	if x != nil {
		return x.CallerIdentityChain
	}
	return nil
}

var File_commonfate_access_v1alpha1_audit_logs_proto protoreflect.FileDescriptor

var file_commonfate_access_v1alpha1_audit_logs_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x15, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x85, 0x01, 0x0a,
	0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f,
	0x67, 0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x68, 0x0a, 0x0f, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x38, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8d,
	0x03, 0x0a, 0x08, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x6f,
	0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x5c, 0x0a, 0x15, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x13, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x32, 0x90,
	0x01, 0x0a, 0x0f, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7d, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x4c, 0x6f, 0x67, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18,
	0x01, 0x42, 0x84, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xca, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_access_v1alpha1_audit_logs_proto_rawDescOnce sync.Once
	file_commonfate_access_v1alpha1_audit_logs_proto_rawDescData = file_commonfate_access_v1alpha1_audit_logs_proto_rawDesc
)

func file_commonfate_access_v1alpha1_audit_logs_proto_rawDescGZIP() []byte {
	file_commonfate_access_v1alpha1_audit_logs_proto_rawDescOnce.Do(func() {
		file_commonfate_access_v1alpha1_audit_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_v1alpha1_audit_logs_proto_rawDescData)
	})
	return file_commonfate_access_v1alpha1_audit_logs_proto_rawDescData
}

var file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_commonfate_access_v1alpha1_audit_logs_proto_goTypes = []interface{}{
	(*QueryAuditLogsRequest)(nil),  // 0: commonfate.access.v1alpha1.QueryAuditLogsRequest
	(*QueryAuditLogsResponse)(nil), // 1: commonfate.access.v1alpha1.QueryAuditLogsResponse
	(*AuditLogPreview)(nil),        // 2: commonfate.access.v1alpha1.AuditLogPreview
	(*AuditLog)(nil),               // 3: commonfate.access.v1alpha1.AuditLog
	(*v1alpha1.EID)(nil),           // 4: commonfate.entity.v1alpha1.EID
	(*User)(nil),                   // 5: commonfate.access.v1alpha1.User
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
	(*structpb.Struct)(nil),        // 7: google.protobuf.Struct
	(*IdentityLink)(nil),           // 8: commonfate.access.v1alpha1.IdentityLink
}
var file_commonfate_access_v1alpha1_audit_logs_proto_depIdxs = []int32{
	4, // 0: commonfate.access.v1alpha1.QueryAuditLogsRequest.target:type_name -> commonfate.entity.v1alpha1.EID
	3, // 1: commonfate.access.v1alpha1.QueryAuditLogsResponse.audit_logs:type_name -> commonfate.access.v1alpha1.AuditLog
	3, // 2: commonfate.access.v1alpha1.AuditLogPreview.logs:type_name -> commonfate.access.v1alpha1.AuditLog
	5, // 3: commonfate.access.v1alpha1.AuditLog.actor:type_name -> commonfate.access.v1alpha1.User
	6, // 4: commonfate.access.v1alpha1.AuditLog.occurred_at:type_name -> google.protobuf.Timestamp
	4, // 5: commonfate.access.v1alpha1.AuditLog.targets:type_name -> commonfate.entity.v1alpha1.EID
	7, // 6: commonfate.access.v1alpha1.AuditLog.context:type_name -> google.protobuf.Struct
	8, // 7: commonfate.access.v1alpha1.AuditLog.caller_identity_chain:type_name -> commonfate.access.v1alpha1.IdentityLink
	0, // 8: commonfate.access.v1alpha1.AuditLogService.QueryAuditLogs:input_type -> commonfate.access.v1alpha1.QueryAuditLogsRequest
	1, // 9: commonfate.access.v1alpha1.AuditLogService.QueryAuditLogs:output_type -> commonfate.access.v1alpha1.QueryAuditLogsResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_commonfate_access_v1alpha1_audit_logs_proto_init() }
func file_commonfate_access_v1alpha1_audit_logs_proto_init() {
	if File_commonfate_access_v1alpha1_audit_logs_proto != nil {
		return
	}
	file_commonfate_access_v1alpha1_identity_proto_init()
	file_commonfate_access_v1alpha1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAuditLogsRequest); i {
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
		file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAuditLogsResponse); i {
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
		file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLogPreview); i {
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
		file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLog); i {
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
			RawDescriptor: file_commonfate_access_v1alpha1_audit_logs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_access_v1alpha1_audit_logs_proto_goTypes,
		DependencyIndexes: file_commonfate_access_v1alpha1_audit_logs_proto_depIdxs,
		MessageInfos:      file_commonfate_access_v1alpha1_audit_logs_proto_msgTypes,
	}.Build()
	File_commonfate_access_v1alpha1_audit_logs_proto = out.File
	file_commonfate_access_v1alpha1_audit_logs_proto_rawDesc = nil
	file_commonfate_access_v1alpha1_audit_logs_proto_goTypes = nil
	file_commonfate_access_v1alpha1_audit_logs_proto_depIdxs = nil
}
