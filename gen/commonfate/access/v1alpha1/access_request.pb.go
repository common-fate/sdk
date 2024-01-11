// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: commonfate/access/v1alpha1/access_request.proto

package accessv1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/control/attest/v1alpha1"
	_ "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type QueryAccessRequestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The token for the next page.
	PageToken string `protobuf:"bytes,10,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *QueryAccessRequestsRequest) Reset() {
	*x = QueryAccessRequestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAccessRequestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAccessRequestsRequest) ProtoMessage() {}

func (x *QueryAccessRequestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAccessRequestsRequest.ProtoReflect.Descriptor instead.
func (*QueryAccessRequestsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP(), []int{0}
}

func (x *QueryAccessRequestsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type QueryAccessRequestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessRequests []*AccessRequest `protobuf:"bytes,1,rep,name=access_requests,json=accessRequests,proto3" json:"access_requests,omitempty"`
	NextPageToken  string           `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *QueryAccessRequestsResponse) Reset() {
	*x = QueryAccessRequestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAccessRequestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAccessRequestsResponse) ProtoMessage() {}

func (x *QueryAccessRequestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAccessRequestsResponse.ProtoReflect.Descriptor instead.
func (*QueryAccessRequestsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP(), []int{1}
}

func (x *QueryAccessRequestsResponse) GetAccessRequests() []*AccessRequest {
	if x != nil {
		return x.AccessRequests
	}
	return nil
}

func (x *QueryAccessRequestsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetAccessRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Access Request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAccessRequestRequest) Reset() {
	*x = GetAccessRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessRequestRequest) ProtoMessage() {}

func (x *GetAccessRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessRequestRequest.ProtoReflect.Descriptor instead.
func (*GetAccessRequestRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccessRequestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAccessRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessRequest *AccessRequest `protobuf:"bytes,1,opt,name=access_request,json=accessRequest,proto3" json:"access_request,omitempty"`
}

func (x *GetAccessRequestResponse) Reset() {
	*x = GetAccessRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessRequestResponse) ProtoMessage() {}

func (x *GetAccessRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessRequestResponse.ProtoReflect.Descriptor instead.
func (*GetAccessRequestResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccessRequestResponse) GetAccessRequest() *AccessRequest {
	if x != nil {
		return x.AccessRequest
	}
	return nil
}

// A request to be given entitlements allowing <Action> to be performed on <Resource>.
//
// For example, if the action is 'Admin', and the resource is AWS::Account::123456789012,
// This request is for entitlements allowing Admin access to the account 123456789012.
type AccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Access Request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Batch ID of the Access Request. Access Requests created in the same
	// API call have the same Batch ID.
	BatchId string `protobuf:"bytes,2,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	// The grants under consideration for access.
	Grants []*Grant `protobuf:"bytes,3,rep,name=grants,proto3" json:"grants,omitempty"`
	// An access request is reviewable if one or more Grants in the request
	// requires a manual review.
	Reviewable bool `protobuf:"varint,4,opt,name=reviewable,proto3" json:"reviewable,omitempty"`
	// True if the Access Request already existed and wasn't created by the API call.
	Existing bool `protobuf:"varint,5,opt,name=existing,proto3" json:"existing,omitempty"`
	// The timestamp that the request was created at.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *AccessRequest) Reset() {
	*x = AccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequest) ProtoMessage() {}

func (x *AccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequest.ProtoReflect.Descriptor instead.
func (*AccessRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP(), []int{4}
}

func (x *AccessRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccessRequest) GetBatchId() string {
	if x != nil {
		return x.BatchId
	}
	return ""
}

func (x *AccessRequest) GetGrants() []*Grant {
	if x != nil {
		return x.Grants
	}
	return nil
}

func (x *AccessRequest) GetReviewable() bool {
	if x != nil {
		return x.Reviewable
	}
	return false
}

func (x *AccessRequest) GetExisting() bool {
	if x != nil {
		return x.Existing
	}
	return false
}

func (x *AccessRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type ApproveAccessRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Access Request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// If provided, will only approve the grants with the specified IDs
	ApproveGrants []string `protobuf:"bytes,2,rep,name=approve_grants,json=approveGrants,proto3" json:"approve_grants,omitempty"`
}

func (x *ApproveAccessRequestRequest) Reset() {
	*x = ApproveAccessRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveAccessRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveAccessRequestRequest) ProtoMessage() {}

func (x *ApproveAccessRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveAccessRequestRequest.ProtoReflect.Descriptor instead.
func (*ApproveAccessRequestRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP(), []int{5}
}

func (x *ApproveAccessRequestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApproveAccessRequestRequest) GetApproveGrants() []string {
	if x != nil {
		return x.ApproveGrants
	}
	return nil
}

type ApproveAccessRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diagnostics []*Diagnostic `protobuf:"bytes,1,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *ApproveAccessRequestResponse) Reset() {
	*x = ApproveAccessRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveAccessRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveAccessRequestResponse) ProtoMessage() {}

func (x *ApproveAccessRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveAccessRequestResponse.ProtoReflect.Descriptor instead.
func (*ApproveAccessRequestResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP(), []int{6}
}

func (x *ApproveAccessRequestResponse) GetDiagnostics() []*Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

type CloseAccessRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Access Request.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// If provided, will only close the grants with the specified IDs
	CloseGrants []string `protobuf:"bytes,2,rep,name=close_grants,json=closeGrants,proto3" json:"close_grants,omitempty"`
}

func (x *CloseAccessRequestRequest) Reset() {
	*x = CloseAccessRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseAccessRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseAccessRequestRequest) ProtoMessage() {}

func (x *CloseAccessRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseAccessRequestRequest.ProtoReflect.Descriptor instead.
func (*CloseAccessRequestRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP(), []int{7}
}

func (x *CloseAccessRequestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CloseAccessRequestRequest) GetCloseGrants() []string {
	if x != nil {
		return x.CloseGrants
	}
	return nil
}

type CloseAccessRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Diagnostics []*Diagnostic `protobuf:"bytes,1,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *CloseAccessRequestResponse) Reset() {
	*x = CloseAccessRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseAccessRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseAccessRequestResponse) ProtoMessage() {}

func (x *CloseAccessRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_access_request_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseAccessRequestResponse.ProtoReflect.Descriptor instead.
func (*CloseAccessRequestResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP(), []int{8}
}

func (x *CloseAccessRequestResponse) GetDiagnostics() []*Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

var File_commonfate_access_v1alpha1_access_request_proto protoreflect.FileDescriptor

var file_commonfate_access_v1alpha1_access_request_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x34, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x64,
	0x5f, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x99, 0x01, 0x0a, 0x1b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6c, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xec, 0x01, 0x0a, 0x0d, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x06, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x54, 0x0a, 0x1b, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x22,
	0x68, 0x0a, 0x1c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x48, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52, 0x0b, 0x64, 0x69,
	0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0x4e, 0x0a, 0x19, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x66, 0x0a, 0x1a, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x32, 0xb8, 0x04, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x88, 0x01, 0x0a, 0x13, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x88, 0x02, 0x0a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x12, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2,
	0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_access_v1alpha1_access_request_proto_rawDescOnce sync.Once
	file_commonfate_access_v1alpha1_access_request_proto_rawDescData = file_commonfate_access_v1alpha1_access_request_proto_rawDesc
)

func file_commonfate_access_v1alpha1_access_request_proto_rawDescGZIP() []byte {
	file_commonfate_access_v1alpha1_access_request_proto_rawDescOnce.Do(func() {
		file_commonfate_access_v1alpha1_access_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_v1alpha1_access_request_proto_rawDescData)
	})
	return file_commonfate_access_v1alpha1_access_request_proto_rawDescData
}

var file_commonfate_access_v1alpha1_access_request_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_commonfate_access_v1alpha1_access_request_proto_goTypes = []interface{}{
	(*QueryAccessRequestsRequest)(nil),   // 0: commonfate.access.v1alpha1.QueryAccessRequestsRequest
	(*QueryAccessRequestsResponse)(nil),  // 1: commonfate.access.v1alpha1.QueryAccessRequestsResponse
	(*GetAccessRequestRequest)(nil),      // 2: commonfate.access.v1alpha1.GetAccessRequestRequest
	(*GetAccessRequestResponse)(nil),     // 3: commonfate.access.v1alpha1.GetAccessRequestResponse
	(*AccessRequest)(nil),                // 4: commonfate.access.v1alpha1.AccessRequest
	(*ApproveAccessRequestRequest)(nil),  // 5: commonfate.access.v1alpha1.ApproveAccessRequestRequest
	(*ApproveAccessRequestResponse)(nil), // 6: commonfate.access.v1alpha1.ApproveAccessRequestResponse
	(*CloseAccessRequestRequest)(nil),    // 7: commonfate.access.v1alpha1.CloseAccessRequestRequest
	(*CloseAccessRequestResponse)(nil),   // 8: commonfate.access.v1alpha1.CloseAccessRequestResponse
	(*Grant)(nil),                        // 9: commonfate.access.v1alpha1.Grant
	(*timestamppb.Timestamp)(nil),        // 10: google.protobuf.Timestamp
	(*Diagnostic)(nil),                   // 11: commonfate.access.v1alpha1.Diagnostic
}
var file_commonfate_access_v1alpha1_access_request_proto_depIdxs = []int32{
	4,  // 0: commonfate.access.v1alpha1.QueryAccessRequestsResponse.access_requests:type_name -> commonfate.access.v1alpha1.AccessRequest
	4,  // 1: commonfate.access.v1alpha1.GetAccessRequestResponse.access_request:type_name -> commonfate.access.v1alpha1.AccessRequest
	9,  // 2: commonfate.access.v1alpha1.AccessRequest.grants:type_name -> commonfate.access.v1alpha1.Grant
	10, // 3: commonfate.access.v1alpha1.AccessRequest.created_at:type_name -> google.protobuf.Timestamp
	11, // 4: commonfate.access.v1alpha1.ApproveAccessRequestResponse.diagnostics:type_name -> commonfate.access.v1alpha1.Diagnostic
	11, // 5: commonfate.access.v1alpha1.CloseAccessRequestResponse.diagnostics:type_name -> commonfate.access.v1alpha1.Diagnostic
	0,  // 6: commonfate.access.v1alpha1.AccessRequestService.QueryAccessRequests:input_type -> commonfate.access.v1alpha1.QueryAccessRequestsRequest
	2,  // 7: commonfate.access.v1alpha1.AccessRequestService.GetAccessRequest:input_type -> commonfate.access.v1alpha1.GetAccessRequestRequest
	5,  // 8: commonfate.access.v1alpha1.AccessRequestService.ApproveAccessRequest:input_type -> commonfate.access.v1alpha1.ApproveAccessRequestRequest
	7,  // 9: commonfate.access.v1alpha1.AccessRequestService.CloseAccessRequest:input_type -> commonfate.access.v1alpha1.CloseAccessRequestRequest
	1,  // 10: commonfate.access.v1alpha1.AccessRequestService.QueryAccessRequests:output_type -> commonfate.access.v1alpha1.QueryAccessRequestsResponse
	3,  // 11: commonfate.access.v1alpha1.AccessRequestService.GetAccessRequest:output_type -> commonfate.access.v1alpha1.GetAccessRequestResponse
	6,  // 12: commonfate.access.v1alpha1.AccessRequestService.ApproveAccessRequest:output_type -> commonfate.access.v1alpha1.ApproveAccessRequestResponse
	8,  // 13: commonfate.access.v1alpha1.AccessRequestService.CloseAccessRequest:output_type -> commonfate.access.v1alpha1.CloseAccessRequestResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_commonfate_access_v1alpha1_access_request_proto_init() }
func file_commonfate_access_v1alpha1_access_request_proto_init() {
	if File_commonfate_access_v1alpha1_access_request_proto != nil {
		return
	}
	file_commonfate_access_v1alpha1_user_proto_init()
	file_commonfate_access_v1alpha1_diagnostic_proto_init()
	file_commonfate_access_v1alpha1_grant_proto_init()
	file_commonfate_access_v1alpha1_named_eid_proto_init()
	file_commonfate_access_v1alpha1_audit_logs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_access_v1alpha1_access_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAccessRequestsRequest); i {
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
		file_commonfate_access_v1alpha1_access_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryAccessRequestsResponse); i {
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
		file_commonfate_access_v1alpha1_access_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessRequestRequest); i {
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
		file_commonfate_access_v1alpha1_access_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessRequestResponse); i {
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
		file_commonfate_access_v1alpha1_access_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequest); i {
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
		file_commonfate_access_v1alpha1_access_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveAccessRequestRequest); i {
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
		file_commonfate_access_v1alpha1_access_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveAccessRequestResponse); i {
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
		file_commonfate_access_v1alpha1_access_request_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseAccessRequestRequest); i {
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
		file_commonfate_access_v1alpha1_access_request_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseAccessRequestResponse); i {
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
			RawDescriptor: file_commonfate_access_v1alpha1_access_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_access_v1alpha1_access_request_proto_goTypes,
		DependencyIndexes: file_commonfate_access_v1alpha1_access_request_proto_depIdxs,
		MessageInfos:      file_commonfate_access_v1alpha1_access_request_proto_msgTypes,
	}.Build()
	File_commonfate_access_v1alpha1_access_request_proto = out.File
	file_commonfate_access_v1alpha1_access_request_proto_rawDesc = nil
	file_commonfate_access_v1alpha1_access_request_proto_goTypes = nil
	file_commonfate_access_v1alpha1_access_request_proto_depIdxs = nil
}
