// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: commonfate/access/v1alpha1/attachments.proto

package accessv1alpha1

import (
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

type JiraIssue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Summary string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Url     string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *JiraIssue) Reset() {
	*x = JiraIssue{}
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JiraIssue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JiraIssue) ProtoMessage() {}

func (x *JiraIssue) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JiraIssue.ProtoReflect.Descriptor instead.
func (*JiraIssue) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_attachments_proto_rawDescGZIP(), []int{0}
}

func (x *JiraIssue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *JiraIssue) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *JiraIssue) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type QueryJiraIssuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SummaryFilter *string `protobuf:"bytes,1,opt,name=summary_filter,json=summaryFilter,proto3,oneof" json:"summary_filter,omitempty"`
}

func (x *QueryJiraIssuesRequest) Reset() {
	*x = QueryJiraIssuesRequest{}
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryJiraIssuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryJiraIssuesRequest) ProtoMessage() {}

func (x *QueryJiraIssuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryJiraIssuesRequest.ProtoReflect.Descriptor instead.
func (*QueryJiraIssuesRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_attachments_proto_rawDescGZIP(), []int{1}
}

func (x *QueryJiraIssuesRequest) GetSummaryFilter() string {
	if x != nil && x.SummaryFilter != nil {
		return *x.SummaryFilter
	}
	return ""
}

type QueryJiraIssuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issues []*JiraIssue `protobuf:"bytes,1,rep,name=issues,proto3" json:"issues,omitempty"`
}

func (x *QueryJiraIssuesResponse) Reset() {
	*x = QueryJiraIssuesResponse{}
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryJiraIssuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryJiraIssuesResponse) ProtoMessage() {}

func (x *QueryJiraIssuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryJiraIssuesResponse.ProtoReflect.Descriptor instead.
func (*QueryJiraIssuesResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_attachments_proto_rawDescGZIP(), []int{2}
}

func (x *QueryJiraIssuesResponse) GetIssues() []*JiraIssue {
	if x != nil {
		return x.Issues
	}
	return nil
}

type AttachRequestCommentToJiraIssuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessRequestReason string `protobuf:"bytes,1,opt,name=access_request_reason,json=accessRequestReason,proto3" json:"access_request_reason,omitempty"`
}

func (x *AttachRequestCommentToJiraIssuesRequest) Reset() {
	*x = AttachRequestCommentToJiraIssuesRequest{}
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttachRequestCommentToJiraIssuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachRequestCommentToJiraIssuesRequest) ProtoMessage() {}

func (x *AttachRequestCommentToJiraIssuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachRequestCommentToJiraIssuesRequest.ProtoReflect.Descriptor instead.
func (*AttachRequestCommentToJiraIssuesRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_attachments_proto_rawDescGZIP(), []int{3}
}

func (x *AttachRequestCommentToJiraIssuesRequest) GetAccessRequestReason() string {
	if x != nil {
		return x.AccessRequestReason
	}
	return ""
}

type AttachRequestCommentToJiraIssuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diagnostics []*Diagnostic `protobuf:"bytes,1,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *AttachRequestCommentToJiraIssuesResponse) Reset() {
	*x = AttachRequestCommentToJiraIssuesResponse{}
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttachRequestCommentToJiraIssuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachRequestCommentToJiraIssuesResponse) ProtoMessage() {}

func (x *AttachRequestCommentToJiraIssuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_attachments_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachRequestCommentToJiraIssuesResponse.ProtoReflect.Descriptor instead.
func (*AttachRequestCommentToJiraIssuesResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_attachments_proto_rawDescGZIP(), []int{4}
}

func (x *AttachRequestCommentToJiraIssuesResponse) GetDiagnostics() []*Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

var File_commonfate_access_v1alpha1_attachments_proto protoreflect.FileDescriptor

var file_commonfate_access_v1alpha1_attachments_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x49, 0x0a, 0x09, 0x4a, 0x69, 0x72, 0x61, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x57, 0x0a,
	0x16, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x69, 0x72, 0x61, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0e, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0d, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x58, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a,
	0x69, 0x72, 0x61, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4a,
	0x69, 0x72, 0x61, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x22, 0x5d, 0x0a, 0x27, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x4a, 0x69, 0x72, 0x61, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0x74, 0x0a, 0x28, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x4a, 0x69, 0x72, 0x61, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x64,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x32, 0xcd, 0x02, 0x0a, 0x12, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a,
	0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x69, 0x72, 0x61, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73,
	0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4a, 0x69, 0x72, 0x61, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x69, 0x72, 0x61, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12,
	0xb3, 0x01, 0x0a, 0x20, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x4a, 0x69, 0x72, 0x61, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x73, 0x12, 0x43, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x4a, 0x69, 0x72, 0x61, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x4a, 0x69, 0x72,
	0x61, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x04, 0x88, 0xb5, 0x18, 0x01, 0x42, 0x86, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x10, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x1a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_access_v1alpha1_attachments_proto_rawDescOnce sync.Once
	file_commonfate_access_v1alpha1_attachments_proto_rawDescData = file_commonfate_access_v1alpha1_attachments_proto_rawDesc
)

func file_commonfate_access_v1alpha1_attachments_proto_rawDescGZIP() []byte {
	file_commonfate_access_v1alpha1_attachments_proto_rawDescOnce.Do(func() {
		file_commonfate_access_v1alpha1_attachments_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_v1alpha1_attachments_proto_rawDescData)
	})
	return file_commonfate_access_v1alpha1_attachments_proto_rawDescData
}

var file_commonfate_access_v1alpha1_attachments_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_commonfate_access_v1alpha1_attachments_proto_goTypes = []any{
	(*JiraIssue)(nil),                                // 0: commonfate.access.v1alpha1.JiraIssue
	(*QueryJiraIssuesRequest)(nil),                   // 1: commonfate.access.v1alpha1.QueryJiraIssuesRequest
	(*QueryJiraIssuesResponse)(nil),                  // 2: commonfate.access.v1alpha1.QueryJiraIssuesResponse
	(*AttachRequestCommentToJiraIssuesRequest)(nil),  // 3: commonfate.access.v1alpha1.AttachRequestCommentToJiraIssuesRequest
	(*AttachRequestCommentToJiraIssuesResponse)(nil), // 4: commonfate.access.v1alpha1.AttachRequestCommentToJiraIssuesResponse
	(*Diagnostic)(nil),                               // 5: commonfate.access.v1alpha1.Diagnostic
}
var file_commonfate_access_v1alpha1_attachments_proto_depIdxs = []int32{
	0, // 0: commonfate.access.v1alpha1.QueryJiraIssuesResponse.issues:type_name -> commonfate.access.v1alpha1.JiraIssue
	5, // 1: commonfate.access.v1alpha1.AttachRequestCommentToJiraIssuesResponse.diagnostics:type_name -> commonfate.access.v1alpha1.Diagnostic
	1, // 2: commonfate.access.v1alpha1.AttachmentsService.QueryJiraIssues:input_type -> commonfate.access.v1alpha1.QueryJiraIssuesRequest
	3, // 3: commonfate.access.v1alpha1.AttachmentsService.AttachRequestCommentToJiraIssues:input_type -> commonfate.access.v1alpha1.AttachRequestCommentToJiraIssuesRequest
	2, // 4: commonfate.access.v1alpha1.AttachmentsService.QueryJiraIssues:output_type -> commonfate.access.v1alpha1.QueryJiraIssuesResponse
	4, // 5: commonfate.access.v1alpha1.AttachmentsService.AttachRequestCommentToJiraIssues:output_type -> commonfate.access.v1alpha1.AttachRequestCommentToJiraIssuesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commonfate_access_v1alpha1_attachments_proto_init() }
func file_commonfate_access_v1alpha1_attachments_proto_init() {
	if File_commonfate_access_v1alpha1_attachments_proto != nil {
		return
	}
	file_commonfate_access_v1alpha1_diagnostic_proto_init()
	file_commonfate_access_v1alpha1_attachments_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_access_v1alpha1_attachments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_access_v1alpha1_attachments_proto_goTypes,
		DependencyIndexes: file_commonfate_access_v1alpha1_attachments_proto_depIdxs,
		MessageInfos:      file_commonfate_access_v1alpha1_attachments_proto_msgTypes,
	}.Build()
	File_commonfate_access_v1alpha1_attachments_proto = out.File
	file_commonfate_access_v1alpha1_attachments_proto_rawDesc = nil
	file_commonfate_access_v1alpha1_attachments_proto_goTypes = nil
	file_commonfate_access_v1alpha1_attachments_proto_depIdxs = nil
}
