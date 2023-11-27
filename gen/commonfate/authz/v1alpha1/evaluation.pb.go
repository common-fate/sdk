// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commonfate/authz/v1alpha1/evaluation.proto

package authzv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Decision int32

const (
	Decision_DECISION_UNSPECIFIED Decision = 0
	// Access is allowed.
	Decision_DECISION_ALLOW Decision = 2
	// Access is denied.
	Decision_DECISION_DENY Decision = 4
)

// Enum value maps for Decision.
var (
	Decision_name = map[int32]string{
		0: "DECISION_UNSPECIFIED",
		2: "DECISION_ALLOW",
		4: "DECISION_DENY",
	}
	Decision_value = map[string]int32{
		"DECISION_UNSPECIFIED": 0,
		"DECISION_ALLOW":       2,
		"DECISION_DENY":        4,
	}
)

func (x Decision) Enum() *Decision {
	p := new(Decision)
	*p = x
	return p
}

func (x Decision) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Decision) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_authz_v1alpha1_evaluation_proto_enumTypes[0].Descriptor()
}

func (Decision) Type() protoreflect.EnumType {
	return &file_commonfate_authz_v1alpha1_evaluation_proto_enumTypes[0]
}

func (x Decision) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Decision.Descriptor instead.
func (Decision) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_evaluation_proto_rawDescGZIP(), []int{0}
}

type DebugEvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DebugEvaluationRequest) Reset() {
	*x = DebugEvaluationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugEvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugEvaluationRequest) ProtoMessage() {}

func (x *DebugEvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugEvaluationRequest.ProtoReflect.Descriptor instead.
func (*DebugEvaluationRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_evaluation_proto_rawDescGZIP(), []int{0}
}

func (x *DebugEvaluationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DebugEvaluationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluation       *Evaluation       `protobuf:"bytes,1,opt,name=evaluation,proto3" json:"evaluation,omitempty"`
	DebugInformation *DebugInformation `protobuf:"bytes,2,opt,name=debug_information,json=debugInformation,proto3" json:"debug_information,omitempty"`
}

func (x *DebugEvaluationResponse) Reset() {
	*x = DebugEvaluationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugEvaluationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugEvaluationResponse) ProtoMessage() {}

func (x *DebugEvaluationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugEvaluationResponse.ProtoReflect.Descriptor instead.
func (*DebugEvaluationResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_evaluation_proto_rawDescGZIP(), []int{1}
}

func (x *DebugEvaluationResponse) GetEvaluation() *Evaluation {
	if x != nil {
		return x.Evaluation
	}
	return nil
}

func (x *DebugEvaluationResponse) GetDebugInformation() *DebugInformation {
	if x != nil {
		return x.DebugInformation
	}
	return nil
}

type QueryEvaluationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Universe      string `protobuf:"bytes,1,opt,name=universe,proto3" json:"universe,omitempty"`
	Environment   string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	NextPageToken string `protobuf:"bytes,10,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *QueryEvaluationsRequest) Reset() {
	*x = QueryEvaluationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryEvaluationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryEvaluationsRequest) ProtoMessage() {}

func (x *QueryEvaluationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryEvaluationsRequest.ProtoReflect.Descriptor instead.
func (*QueryEvaluationsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_evaluation_proto_rawDescGZIP(), []int{2}
}

func (x *QueryEvaluationsRequest) GetUniverse() string {
	if x != nil {
		return x.Universe
	}
	return ""
}

func (x *QueryEvaluationsRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *QueryEvaluationsRequest) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type QueryEvaluationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Evaluations   []*Evaluation `protobuf:"bytes,1,rep,name=evaluations,proto3" json:"evaluations,omitempty"`
	NextPageToken string        `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *QueryEvaluationsResponse) Reset() {
	*x = QueryEvaluationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryEvaluationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryEvaluationsResponse) ProtoMessage() {}

func (x *QueryEvaluationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryEvaluationsResponse.ProtoReflect.Descriptor instead.
func (*QueryEvaluationsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_evaluation_proto_rawDescGZIP(), []int{3}
}

func (x *QueryEvaluationsResponse) GetEvaluations() []*Evaluation {
	if x != nil {
		return x.Evaluations
	}
	return nil
}

func (x *QueryEvaluationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type Evaluation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Request     *Request     `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Decision    Decision     `protobuf:"varint,3,opt,name=decision,proto3,enum=commonfate.authz.v1alpha1.Decision" json:"decision,omitempty"`
	Diagnostics *Diagnostics `protobuf:"bytes,4,opt,name=diagnostics,proto3" json:"diagnostics,omitempty"`
	// The client-side identifier for the request specified in AuthorizationRequest.
	ClientKey   string                 `protobuf:"bytes,5,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	EvaluatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=evaluated_at,json=evaluatedAt,proto3" json:"evaluated_at,omitempty"`
}

func (x *Evaluation) Reset() {
	*x = Evaluation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Evaluation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Evaluation) ProtoMessage() {}

func (x *Evaluation) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Evaluation.ProtoReflect.Descriptor instead.
func (*Evaluation) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_evaluation_proto_rawDescGZIP(), []int{4}
}

func (x *Evaluation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Evaluation) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Evaluation) GetDecision() Decision {
	if x != nil {
		return x.Decision
	}
	return Decision_DECISION_UNSPECIFIED
}

func (x *Evaluation) GetDiagnostics() *Diagnostics {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

func (x *Evaluation) GetClientKey() string {
	if x != nil {
		return x.ClientKey
	}
	return ""
}

func (x *Evaluation) GetEvaluatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EvaluatedAt
	}
	return nil
}

type Diagnostics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason []string `protobuf:"bytes,1,rep,name=reason,proto3" json:"reason,omitempty"`
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Diagnostics) Reset() {
	*x = Diagnostics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Diagnostics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostics) ProtoMessage() {}

func (x *Diagnostics) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostics.ProtoReflect.Descriptor instead.
func (*Diagnostics) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_evaluation_proto_rawDescGZIP(), []int{5}
}

func (x *Diagnostics) GetReason() []string {
	if x != nil {
		return x.Reason
	}
	return nil
}

func (x *Diagnostics) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type DebugInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Policies which contributed to the authorization decision.
	Policies []*Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	// Entities which contributed to the authorization decision.
	Entities []*Entity `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
}

func (x *DebugInformation) Reset() {
	*x = DebugInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DebugInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugInformation) ProtoMessage() {}

func (x *DebugInformation) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugInformation.ProtoReflect.Descriptor instead.
func (*DebugInformation) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_evaluation_proto_rawDescGZIP(), []int{6}
}

func (x *DebugInformation) GetPolicies() []*Policy {
	if x != nil {
		return x.Policies
	}
	return nil
}

func (x *DebugInformation) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

var File_commonfate_authz_v1alpha1_evaluation_proto protoreflect.FileDescriptor

var file_commonfate_authz_v1alpha1_evaluation_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x62, 0x75, 0x67, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x17,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58,
	0x0a, 0x11, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x17, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8b, 0x01, 0x0a, 0x18, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc3, 0x02, 0x0a, 0x0a, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x3d,
	0x0a, 0x0c, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0b, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3d, 0x0a,
	0x0b, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x90, 0x01, 0x0a,
	0x10, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x12, 0x3d, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2a,
	0x4b, 0x0a, 0x08, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x14, 0x44,
	0x45, 0x43, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x45, 0x43, 0x49, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x43,
	0x49, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x04, 0x32, 0x8e, 0x02, 0x0a,
	0x11, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7d, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x7a, 0x0a, 0x0f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xfe, 0x01,
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x0f, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58,
	0xaa, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x7a, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x41,
	0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_authz_v1alpha1_evaluation_proto_rawDescOnce sync.Once
	file_commonfate_authz_v1alpha1_evaluation_proto_rawDescData = file_commonfate_authz_v1alpha1_evaluation_proto_rawDesc
)

func file_commonfate_authz_v1alpha1_evaluation_proto_rawDescGZIP() []byte {
	file_commonfate_authz_v1alpha1_evaluation_proto_rawDescOnce.Do(func() {
		file_commonfate_authz_v1alpha1_evaluation_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_authz_v1alpha1_evaluation_proto_rawDescData)
	})
	return file_commonfate_authz_v1alpha1_evaluation_proto_rawDescData
}

var file_commonfate_authz_v1alpha1_evaluation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_commonfate_authz_v1alpha1_evaluation_proto_goTypes = []interface{}{
	(Decision)(0),                    // 0: commonfate.authz.v1alpha1.Decision
	(*DebugEvaluationRequest)(nil),   // 1: commonfate.authz.v1alpha1.DebugEvaluationRequest
	(*DebugEvaluationResponse)(nil),  // 2: commonfate.authz.v1alpha1.DebugEvaluationResponse
	(*QueryEvaluationsRequest)(nil),  // 3: commonfate.authz.v1alpha1.QueryEvaluationsRequest
	(*QueryEvaluationsResponse)(nil), // 4: commonfate.authz.v1alpha1.QueryEvaluationsResponse
	(*Evaluation)(nil),               // 5: commonfate.authz.v1alpha1.Evaluation
	(*Diagnostics)(nil),              // 6: commonfate.authz.v1alpha1.Diagnostics
	(*DebugInformation)(nil),         // 7: commonfate.authz.v1alpha1.DebugInformation
	(*Request)(nil),                  // 8: commonfate.authz.v1alpha1.Request
	(*timestamppb.Timestamp)(nil),    // 9: google.protobuf.Timestamp
	(*Policy)(nil),                   // 10: commonfate.authz.v1alpha1.Policy
	(*Entity)(nil),                   // 11: commonfate.authz.v1alpha1.Entity
}
var file_commonfate_authz_v1alpha1_evaluation_proto_depIdxs = []int32{
	5,  // 0: commonfate.authz.v1alpha1.DebugEvaluationResponse.evaluation:type_name -> commonfate.authz.v1alpha1.Evaluation
	7,  // 1: commonfate.authz.v1alpha1.DebugEvaluationResponse.debug_information:type_name -> commonfate.authz.v1alpha1.DebugInformation
	5,  // 2: commonfate.authz.v1alpha1.QueryEvaluationsResponse.evaluations:type_name -> commonfate.authz.v1alpha1.Evaluation
	8,  // 3: commonfate.authz.v1alpha1.Evaluation.request:type_name -> commonfate.authz.v1alpha1.Request
	0,  // 4: commonfate.authz.v1alpha1.Evaluation.decision:type_name -> commonfate.authz.v1alpha1.Decision
	6,  // 5: commonfate.authz.v1alpha1.Evaluation.diagnostics:type_name -> commonfate.authz.v1alpha1.Diagnostics
	9,  // 6: commonfate.authz.v1alpha1.Evaluation.evaluated_at:type_name -> google.protobuf.Timestamp
	10, // 7: commonfate.authz.v1alpha1.DebugInformation.policies:type_name -> commonfate.authz.v1alpha1.Policy
	11, // 8: commonfate.authz.v1alpha1.DebugInformation.entities:type_name -> commonfate.authz.v1alpha1.Entity
	3,  // 9: commonfate.authz.v1alpha1.EvaluationService.QueryEvaluations:input_type -> commonfate.authz.v1alpha1.QueryEvaluationsRequest
	1,  // 10: commonfate.authz.v1alpha1.EvaluationService.DebugEvaluation:input_type -> commonfate.authz.v1alpha1.DebugEvaluationRequest
	4,  // 11: commonfate.authz.v1alpha1.EvaluationService.QueryEvaluations:output_type -> commonfate.authz.v1alpha1.QueryEvaluationsResponse
	2,  // 12: commonfate.authz.v1alpha1.EvaluationService.DebugEvaluation:output_type -> commonfate.authz.v1alpha1.DebugEvaluationResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_commonfate_authz_v1alpha1_evaluation_proto_init() }
func file_commonfate_authz_v1alpha1_evaluation_proto_init() {
	if File_commonfate_authz_v1alpha1_evaluation_proto != nil {
		return
	}
	file_commonfate_authz_v1alpha1_request_proto_init()
	file_commonfate_authz_v1alpha1_policy_proto_init()
	file_commonfate_authz_v1alpha1_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugEvaluationRequest); i {
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
		file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugEvaluationResponse); i {
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
		file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryEvaluationsRequest); i {
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
		file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryEvaluationsResponse); i {
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
		file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Evaluation); i {
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
		file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Diagnostics); i {
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
		file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DebugInformation); i {
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
			RawDescriptor: file_commonfate_authz_v1alpha1_evaluation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_authz_v1alpha1_evaluation_proto_goTypes,
		DependencyIndexes: file_commonfate_authz_v1alpha1_evaluation_proto_depIdxs,
		EnumInfos:         file_commonfate_authz_v1alpha1_evaluation_proto_enumTypes,
		MessageInfos:      file_commonfate_authz_v1alpha1_evaluation_proto_msgTypes,
	}.Build()
	File_commonfate_authz_v1alpha1_evaluation_proto = out.File
	file_commonfate_authz_v1alpha1_evaluation_proto_rawDesc = nil
	file_commonfate_authz_v1alpha1_evaluation_proto_goTypes = nil
	file_commonfate_authz_v1alpha1_evaluation_proto_depIdxs = nil
}
