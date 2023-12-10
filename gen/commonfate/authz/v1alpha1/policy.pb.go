// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commonfate/authz/v1alpha1/policy.proto

package authzv1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
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

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the policy.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Cedar code for the policy set.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Policy) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type PolicySet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the policy set.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The policies contained in the policy set.
	Policies []*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *PolicySet) Reset() {
	*x = PolicySet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicySet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicySet) ProtoMessage() {}

func (x *PolicySet) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicySet.ProtoReflect.Descriptor instead.
func (*PolicySet) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{1}
}

func (x *PolicySet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PolicySet) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

type PolicySetInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the policy.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Cedar code for the policy set.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PolicySetInput) Reset() {
	*x = PolicySetInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicySetInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicySetInput) ProtoMessage() {}

func (x *PolicySetInput) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicySetInput.ProtoReflect.Descriptor instead.
func (*PolicySetInput) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{2}
}

func (x *PolicySetInput) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PolicySetInput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ListPolicySetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The token for the next page.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListPolicySetsRequest) Reset() {
	*x = ListPolicySetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPolicySetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPolicySetsRequest) ProtoMessage() {}

func (x *ListPolicySetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPolicySetsRequest.ProtoReflect.Descriptor instead.
func (*ListPolicySetsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{3}
}

func (x *ListPolicySetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListPolicySetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicySets    []*PolicySet `protobuf:"bytes,1,rep,name=policy_sets,json=policySets,proto3" json:"policy_sets,omitempty"`
	NextPageToken string       `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPolicySetsResponse) Reset() {
	*x = ListPolicySetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPolicySetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPolicySetsResponse) ProtoMessage() {}

func (x *ListPolicySetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPolicySetsResponse.ProtoReflect.Descriptor instead.
func (*ListPolicySetsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{4}
}

func (x *ListPolicySetsResponse) GetPolicySets() []*PolicySet {
	if x != nil {
		return x.PolicySets
	}
	return nil
}

func (x *ListPolicySetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type BatchPutPolicySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicySets []*PolicySetInput `protobuf:"bytes,3,rep,name=policy_sets,json=policySets,proto3" json:"policy_sets,omitempty"`
}

func (x *BatchPutPolicySetRequest) Reset() {
	*x = BatchPutPolicySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchPutPolicySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchPutPolicySetRequest) ProtoMessage() {}

func (x *BatchPutPolicySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchPutPolicySetRequest.ProtoReflect.Descriptor instead.
func (*BatchPutPolicySetRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{5}
}

func (x *BatchPutPolicySetRequest) GetPolicySets() []*PolicySetInput {
	if x != nil {
		return x.PolicySets
	}
	return nil
}

type BatchPutPolicySetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BatchPutPolicySetResponse) Reset() {
	*x = BatchPutPolicySetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchPutPolicySetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchPutPolicySetResponse) ProtoMessage() {}

func (x *BatchPutPolicySetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchPutPolicySetResponse.ProtoReflect.Descriptor instead.
func (*BatchPutPolicySetResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{6}
}

type DeletePolicySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicySetIds []string `protobuf:"bytes,3,rep,name=policy_set_ids,json=policySetIds,proto3" json:"policy_set_ids,omitempty"`
}

func (x *DeletePolicySetRequest) Reset() {
	*x = DeletePolicySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicySetRequest) ProtoMessage() {}

func (x *DeletePolicySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicySetRequest.ProtoReflect.Descriptor instead.
func (*DeletePolicySetRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePolicySetRequest) GetPolicySetIds() []string {
	if x != nil {
		return x.PolicySetIds
	}
	return nil
}

type DeletePolicySetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicySetIds []string `protobuf:"bytes,1,rep,name=policy_set_ids,json=policySetIds,proto3" json:"policy_set_ids,omitempty"`
}

func (x *DeletePolicySetResponse) Reset() {
	*x = DeletePolicySetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicySetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicySetResponse) ProtoMessage() {}

func (x *DeletePolicySetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicySetResponse.ProtoReflect.Descriptor instead.
func (*DeletePolicySetResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePolicySetResponse) GetPolicySetIds() []string {
	if x != nil {
		return x.PolicySetIds
	}
	return nil
}

type GetPolicySetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The token for the next page.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPolicySetRequest) Reset() {
	*x = GetPolicySetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicySetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicySetRequest) ProtoMessage() {}

func (x *GetPolicySetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicySetRequest.ProtoReflect.Descriptor instead.
func (*GetPolicySetRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{9}
}

func (x *GetPolicySetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPolicySetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicySet *PolicySet `protobuf:"bytes,1,opt,name=policy_set,json=policySet,proto3" json:"policy_set,omitempty"`
}

func (x *GetPolicySetResponse) Reset() {
	*x = GetPolicySetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicySetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicySetResponse) ProtoMessage() {}

func (x *GetPolicySetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_policy_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicySetResponse.ProtoReflect.Descriptor instead.
func (*GetPolicySetResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP(), []int{10}
}

func (x *GetPolicySetResponse) GetPolicySet() *PolicySet {
	if x != nil {
		return x.PolicySet
	}
	return nil
}

var File_commonfate_authz_v1alpha1_policy_proto protoreflect.FileDescriptor

var file_commonfate_authz_v1alpha1_policy_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x5a, 0x0a, 0x09, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x0e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x53, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x36, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x87, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x65, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x66,
	0x0a, 0x18, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0b, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x65, 0x74, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74,
	0x49, 0x64, 0x73, 0x22, 0x3f, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24,
	0x0a, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65,
	0x74, 0x49, 0x64, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52, 0x09, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x32, 0xfa, 0x03, 0x0a, 0x0d, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x11, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74,
	0x12, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x73, 0x12,
	0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7a, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x31, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xfa, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75,
	0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_authz_v1alpha1_policy_proto_rawDescOnce sync.Once
	file_commonfate_authz_v1alpha1_policy_proto_rawDescData = file_commonfate_authz_v1alpha1_policy_proto_rawDesc
)

func file_commonfate_authz_v1alpha1_policy_proto_rawDescGZIP() []byte {
	file_commonfate_authz_v1alpha1_policy_proto_rawDescOnce.Do(func() {
		file_commonfate_authz_v1alpha1_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_authz_v1alpha1_policy_proto_rawDescData)
	})
	return file_commonfate_authz_v1alpha1_policy_proto_rawDescData
}

var file_commonfate_authz_v1alpha1_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_commonfate_authz_v1alpha1_policy_proto_goTypes = []interface{}{
	(*Policy)(nil),                    // 0: commonfate.authz.v1alpha1.Policy
	(*PolicySet)(nil),                 // 1: commonfate.authz.v1alpha1.PolicySet
	(*PolicySetInput)(nil),            // 2: commonfate.authz.v1alpha1.PolicySetInput
	(*ListPolicySetsRequest)(nil),     // 3: commonfate.authz.v1alpha1.ListPolicySetsRequest
	(*ListPolicySetsResponse)(nil),    // 4: commonfate.authz.v1alpha1.ListPolicySetsResponse
	(*BatchPutPolicySetRequest)(nil),  // 5: commonfate.authz.v1alpha1.BatchPutPolicySetRequest
	(*BatchPutPolicySetResponse)(nil), // 6: commonfate.authz.v1alpha1.BatchPutPolicySetResponse
	(*DeletePolicySetRequest)(nil),    // 7: commonfate.authz.v1alpha1.DeletePolicySetRequest
	(*DeletePolicySetResponse)(nil),   // 8: commonfate.authz.v1alpha1.DeletePolicySetResponse
	(*GetPolicySetRequest)(nil),       // 9: commonfate.authz.v1alpha1.GetPolicySetRequest
	(*GetPolicySetResponse)(nil),      // 10: commonfate.authz.v1alpha1.GetPolicySetResponse
}
var file_commonfate_authz_v1alpha1_policy_proto_depIdxs = []int32{
	0,  // 0: commonfate.authz.v1alpha1.PolicySet.policies:type_name -> commonfate.authz.v1alpha1.Policy
	1,  // 1: commonfate.authz.v1alpha1.ListPolicySetsResponse.policy_sets:type_name -> commonfate.authz.v1alpha1.PolicySet
	2,  // 2: commonfate.authz.v1alpha1.BatchPutPolicySetRequest.policy_sets:type_name -> commonfate.authz.v1alpha1.PolicySetInput
	1,  // 3: commonfate.authz.v1alpha1.GetPolicySetResponse.policy_set:type_name -> commonfate.authz.v1alpha1.PolicySet
	5,  // 4: commonfate.authz.v1alpha1.PolicyService.BatchPutPolicySet:input_type -> commonfate.authz.v1alpha1.BatchPutPolicySetRequest
	3,  // 5: commonfate.authz.v1alpha1.PolicyService.ListPolicySets:input_type -> commonfate.authz.v1alpha1.ListPolicySetsRequest
	9,  // 6: commonfate.authz.v1alpha1.PolicyService.GetPolicySet:input_type -> commonfate.authz.v1alpha1.GetPolicySetRequest
	7,  // 7: commonfate.authz.v1alpha1.PolicyService.DeletePolicySet:input_type -> commonfate.authz.v1alpha1.DeletePolicySetRequest
	6,  // 8: commonfate.authz.v1alpha1.PolicyService.BatchPutPolicySet:output_type -> commonfate.authz.v1alpha1.BatchPutPolicySetResponse
	4,  // 9: commonfate.authz.v1alpha1.PolicyService.ListPolicySets:output_type -> commonfate.authz.v1alpha1.ListPolicySetsResponse
	10, // 10: commonfate.authz.v1alpha1.PolicyService.GetPolicySet:output_type -> commonfate.authz.v1alpha1.GetPolicySetResponse
	8,  // 11: commonfate.authz.v1alpha1.PolicyService.DeletePolicySet:output_type -> commonfate.authz.v1alpha1.DeletePolicySetResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_commonfate_authz_v1alpha1_policy_proto_init() }
func file_commonfate_authz_v1alpha1_policy_proto_init() {
	if File_commonfate_authz_v1alpha1_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicySet); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicySetInput); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPolicySetsRequest); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPolicySetsResponse); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchPutPolicySetRequest); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchPutPolicySetResponse); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicySetRequest); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicySetResponse); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicySetRequest); i {
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
		file_commonfate_authz_v1alpha1_policy_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicySetResponse); i {
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
			RawDescriptor: file_commonfate_authz_v1alpha1_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_authz_v1alpha1_policy_proto_goTypes,
		DependencyIndexes: file_commonfate_authz_v1alpha1_policy_proto_depIdxs,
		MessageInfos:      file_commonfate_authz_v1alpha1_policy_proto_msgTypes,
	}.Build()
	File_commonfate_authz_v1alpha1_policy_proto = out.File
	file_commonfate_authz_v1alpha1_policy_proto_rawDesc = nil
	file_commonfate_authz_v1alpha1_policy_proto_goTypes = nil
	file_commonfate_authz_v1alpha1_policy_proto_depIdxs = nil
}
