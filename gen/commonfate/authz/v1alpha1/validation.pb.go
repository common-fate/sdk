// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: commonfate/authz/v1alpha1/validation.proto

package authzv1alpha1

import (
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

type ValidationErrorKind int32

const (
	ValidationErrorKind_VALIDATION_ERROR_KIND_UNSPECIFIED ValidationErrorKind = 0
	// A policy contains an entity type that is not declared in the schema.
	ValidationErrorKind_VALIDATION_ERROR_UNRECOGNIZED_ENTITY_TYPE ValidationErrorKind = 1
	// A policy contains an action that is not declared in the schema.
	ValidationErrorKind_VALIDATION_ERROR_UNRECOGNIZED_ACTION_ID ValidationErrorKind = 2
	// There is no action satisfying the action head constraint that can be
	// applied to a principal and resources that both satisfy their respective
	// head conditions.
	ValidationErrorKind_VALIDATION_ERROR_INVALID_ACTION_APPLICATION ValidationErrorKind = 3
	// The type checker found an error.
	ValidationErrorKind_VALIDATION_ERROR_TYPE_ERROR ValidationErrorKind = 4
	// An unspecified entity was used in a policy.
	ValidationErrorKind_VALIDATION_ERROR_UNSPECIFIED_ENTITY ValidationErrorKind = 5
)

// Enum value maps for ValidationErrorKind.
var (
	ValidationErrorKind_name = map[int32]string{
		0: "VALIDATION_ERROR_KIND_UNSPECIFIED",
		1: "VALIDATION_ERROR_UNRECOGNIZED_ENTITY_TYPE",
		2: "VALIDATION_ERROR_UNRECOGNIZED_ACTION_ID",
		3: "VALIDATION_ERROR_INVALID_ACTION_APPLICATION",
		4: "VALIDATION_ERROR_TYPE_ERROR",
		5: "VALIDATION_ERROR_UNSPECIFIED_ENTITY",
	}
	ValidationErrorKind_value = map[string]int32{
		"VALIDATION_ERROR_KIND_UNSPECIFIED":           0,
		"VALIDATION_ERROR_UNRECOGNIZED_ENTITY_TYPE":   1,
		"VALIDATION_ERROR_UNRECOGNIZED_ACTION_ID":     2,
		"VALIDATION_ERROR_INVALID_ACTION_APPLICATION": 3,
		"VALIDATION_ERROR_TYPE_ERROR":                 4,
		"VALIDATION_ERROR_UNSPECIFIED_ENTITY":         5,
	}
)

func (x ValidationErrorKind) Enum() *ValidationErrorKind {
	p := new(ValidationErrorKind)
	*p = x
	return p
}

func (x ValidationErrorKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidationErrorKind) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_authz_v1alpha1_validation_proto_enumTypes[0].Descriptor()
}

func (ValidationErrorKind) Type() protoreflect.EnumType {
	return &file_commonfate_authz_v1alpha1_validation_proto_enumTypes[0]
}

func (x ValidationErrorKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidationErrorKind.Descriptor instead.
func (ValidationErrorKind) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_validation_proto_rawDescGZIP(), []int{0}
}

type ValidationWarningKind int32

const (
	ValidationWarningKind_VALIDATION_WARNING_KIND_UNSPECIFIED ValidationWarningKind = 0
	// A string contains mixed scripts. Different scripts can contain visually similar characters which may be confused for each other.
	ValidationWarningKind_VALIDATION_WARNING_MIXED_SCRIPT_STRING ValidationWarningKind = 1
	// A string contains BIDI control characters. These can be used to create crafted pieces of code that obfuscate true control flow.
	ValidationWarningKind_VALIDATION_WARNING_BIDI_CHARS_IN_STRING ValidationWarningKind = 2
	// An id contains BIDI control characters. These can be used to create crafted pieces of code that obfuscate true control flow.
	ValidationWarningKind_VALIDATION_WARNING_BIDI_CHARS_IN_IDENTIFIER ValidationWarningKind = 3
	// An id contains mixed scripts. This can cause characters to be confused for each other.
	ValidationWarningKind_VALIDATION_WARNING_MIXED_SCRIPT_IDENTIFIER ValidationWarningKind = 4
	// An id contains characters that fall outside of the General Security Profile for Identifiers. We recommend adhering to this if possible. See Unicode® Technical Standard #39 for more info.
	ValidationWarningKind_VALIDATION_WARNING_CONFUSABLE_IDENTIFIER ValidationWarningKind = 5
	// An entity referenced in a policy does not exist.
	ValidationWarningKind_VALIDATION_WARNING_ENTITY_DOES_NOT_EXIST ValidationWarningKind = 6
)

// Enum value maps for ValidationWarningKind.
var (
	ValidationWarningKind_name = map[int32]string{
		0: "VALIDATION_WARNING_KIND_UNSPECIFIED",
		1: "VALIDATION_WARNING_MIXED_SCRIPT_STRING",
		2: "VALIDATION_WARNING_BIDI_CHARS_IN_STRING",
		3: "VALIDATION_WARNING_BIDI_CHARS_IN_IDENTIFIER",
		4: "VALIDATION_WARNING_MIXED_SCRIPT_IDENTIFIER",
		5: "VALIDATION_WARNING_CONFUSABLE_IDENTIFIER",
		6: "VALIDATION_WARNING_ENTITY_DOES_NOT_EXIST",
	}
	ValidationWarningKind_value = map[string]int32{
		"VALIDATION_WARNING_KIND_UNSPECIFIED":         0,
		"VALIDATION_WARNING_MIXED_SCRIPT_STRING":      1,
		"VALIDATION_WARNING_BIDI_CHARS_IN_STRING":     2,
		"VALIDATION_WARNING_BIDI_CHARS_IN_IDENTIFIER": 3,
		"VALIDATION_WARNING_MIXED_SCRIPT_IDENTIFIER":  4,
		"VALIDATION_WARNING_CONFUSABLE_IDENTIFIER":    5,
		"VALIDATION_WARNING_ENTITY_DOES_NOT_EXIST":    6,
	}
)

func (x ValidationWarningKind) Enum() *ValidationWarningKind {
	p := new(ValidationWarningKind)
	*p = x
	return p
}

func (x ValidationWarningKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidationWarningKind) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_authz_v1alpha1_validation_proto_enumTypes[1].Descriptor()
}

func (ValidationWarningKind) Type() protoreflect.EnumType {
	return &file_commonfate_authz_v1alpha1_validation_proto_enumTypes[1]
}

func (x ValidationWarningKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidationWarningKind.Descriptor instead.
func (ValidationWarningKind) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_validation_proto_rawDescGZIP(), []int{1}
}

type ValidatePolicySetTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicySetText string `protobuf:"bytes,1,opt,name=policy_set_text,json=policySetText,proto3" json:"policy_set_text,omitempty"`
}

func (x *ValidatePolicySetTextRequest) Reset() {
	*x = ValidatePolicySetTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePolicySetTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePolicySetTextRequest) ProtoMessage() {}

func (x *ValidatePolicySetTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePolicySetTextRequest.ProtoReflect.Descriptor instead.
func (*ValidatePolicySetTextRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_validation_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatePolicySetTextRequest) GetPolicySetText() string {
	if x != nil {
		return x.PolicySetText
	}
	return ""
}

type ValidatePolicySetTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *ValidationResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ValidatePolicySetTextResponse) Reset() {
	*x = ValidatePolicySetTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatePolicySetTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatePolicySetTextResponse) ProtoMessage() {}

func (x *ValidatePolicySetTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatePolicySetTextResponse.ProtoReflect.Descriptor instead.
func (*ValidatePolicySetTextResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_validation_proto_rawDescGZIP(), []int{1}
}

func (x *ValidatePolicySetTextResponse) GetResult() *ValidationResult {
	if x != nil {
		return x.Result
	}
	return nil
}

type ValidationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors   []*ValidationError   `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	Warnings []*ValidationWarning `protobuf:"bytes,2,rep,name=warnings,proto3" json:"warnings,omitempty"`
}

func (x *ValidationResult) Reset() {
	*x = ValidationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationResult) ProtoMessage() {}

func (x *ValidationResult) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationResult.ProtoReflect.Descriptor instead.
func (*ValidationResult) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_validation_proto_rawDescGZIP(), []int{2}
}

func (x *ValidationResult) GetErrors() []*ValidationError {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *ValidationResult) GetWarnings() []*ValidationWarning {
	if x != nil {
		return x.Warnings
	}
	return nil
}

type ValidationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ValidationError) Reset() {
	*x = ValidationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationError) ProtoMessage() {}

func (x *ValidationError) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationError.ProtoReflect.Descriptor instead.
func (*ValidationError) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_validation_proto_rawDescGZIP(), []int{3}
}

func (x *ValidationError) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *ValidationError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ValidationWarning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ValidationWarning) Reset() {
	*x = ValidationWarning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationWarning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationWarning) ProtoMessage() {}

func (x *ValidationWarning) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationWarning.ProtoReflect.Descriptor instead.
func (*ValidationWarning) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_validation_proto_rawDescGZIP(), []int{4}
}

func (x *ValidationWarning) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *ValidationWarning) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The range in the policy definition which contains the error
type SourceRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start uint32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   uint32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *SourceRange) Reset() {
	*x = SourceRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceRange) ProtoMessage() {}

func (x *SourceRange) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_validation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceRange.ProtoReflect.Descriptor instead.
func (*SourceRange) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_validation_proto_rawDescGZIP(), []int{5}
}

func (x *SourceRange) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *SourceRange) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_commonfate_authz_v1alpha1_validation_proto protoreflect.FileDescriptor

var file_commonfate_authz_v1alpha1_validation_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x46, 0x0a, 0x1c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x74,
	0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x22, 0x64, 0x0a, 0x1d, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0xa0, 0x01, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x42, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x48, 0x0a, 0x08, 0x77, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x22, 0x48, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4a, 0x0a,
	0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x72, 0x6e, 0x69,
	0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a, 0x0b, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x2a, 0x93, 0x02, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x21, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x2d, 0x0a, 0x29, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x43, 0x4f, 0x47, 0x4e, 0x49, 0x5a, 0x45, 0x44,
	0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x01, 0x12, 0x2b,
	0x0a, 0x27, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x43, 0x4f, 0x47, 0x4e, 0x49, 0x5a, 0x45, 0x44, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x12, 0x2f, 0x0a, 0x2b, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x27, 0x0a,
	0x23, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x45, 0x4e,
	0x54, 0x49, 0x54, 0x59, 0x10, 0x05, 0x2a, 0xd6, 0x02, 0x0a, 0x15, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x27, 0x0a, 0x23, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x57,
	0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2a, 0x0a, 0x26, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f,
	0x4d, 0x49, 0x58, 0x45, 0x44, 0x5f, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x5f, 0x53, 0x54, 0x52,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x2b, 0x0a, 0x27, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x42, 0x49, 0x44, 0x49,
	0x5f, 0x43, 0x48, 0x41, 0x52, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x2f, 0x0a, 0x2b, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x42, 0x49, 0x44, 0x49, 0x5f, 0x43, 0x48,
	0x41, 0x52, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45,
	0x52, 0x10, 0x03, 0x12, 0x2e, 0x0a, 0x2a, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x49, 0x58, 0x45, 0x44, 0x5f,
	0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45,
	0x52, 0x10, 0x04, 0x12, 0x2c, 0x0a, 0x28, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x55, 0x53,
	0x41, 0x42, 0x4c, 0x45, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10,
	0x05, 0x12, 0x2c, 0x0a, 0x28, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x4e, 0x54, 0x49, 0x54, 0x59, 0x5f, 0x44,
	0x4f, 0x45, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x06, 0x32,
	0xa6, 0x01, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x90, 0x01, 0x0a, 0x15, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x54, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x42, 0xfe, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0f, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x19, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_commonfate_authz_v1alpha1_validation_proto_rawDescOnce sync.Once
	file_commonfate_authz_v1alpha1_validation_proto_rawDescData = file_commonfate_authz_v1alpha1_validation_proto_rawDesc
)

func file_commonfate_authz_v1alpha1_validation_proto_rawDescGZIP() []byte {
	file_commonfate_authz_v1alpha1_validation_proto_rawDescOnce.Do(func() {
		file_commonfate_authz_v1alpha1_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_authz_v1alpha1_validation_proto_rawDescData)
	})
	return file_commonfate_authz_v1alpha1_validation_proto_rawDescData
}

var file_commonfate_authz_v1alpha1_validation_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_commonfate_authz_v1alpha1_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_commonfate_authz_v1alpha1_validation_proto_goTypes = []interface{}{
	(ValidationErrorKind)(0),              // 0: commonfate.authz.v1alpha1.ValidationErrorKind
	(ValidationWarningKind)(0),            // 1: commonfate.authz.v1alpha1.ValidationWarningKind
	(*ValidatePolicySetTextRequest)(nil),  // 2: commonfate.authz.v1alpha1.ValidatePolicySetTextRequest
	(*ValidatePolicySetTextResponse)(nil), // 3: commonfate.authz.v1alpha1.ValidatePolicySetTextResponse
	(*ValidationResult)(nil),              // 4: commonfate.authz.v1alpha1.ValidationResult
	(*ValidationError)(nil),               // 5: commonfate.authz.v1alpha1.ValidationError
	(*ValidationWarning)(nil),             // 6: commonfate.authz.v1alpha1.ValidationWarning
	(*SourceRange)(nil),                   // 7: commonfate.authz.v1alpha1.SourceRange
}
var file_commonfate_authz_v1alpha1_validation_proto_depIdxs = []int32{
	4, // 0: commonfate.authz.v1alpha1.ValidatePolicySetTextResponse.result:type_name -> commonfate.authz.v1alpha1.ValidationResult
	5, // 1: commonfate.authz.v1alpha1.ValidationResult.errors:type_name -> commonfate.authz.v1alpha1.ValidationError
	6, // 2: commonfate.authz.v1alpha1.ValidationResult.warnings:type_name -> commonfate.authz.v1alpha1.ValidationWarning
	2, // 3: commonfate.authz.v1alpha1.ValidationService.ValidatePolicySetText:input_type -> commonfate.authz.v1alpha1.ValidatePolicySetTextRequest
	3, // 4: commonfate.authz.v1alpha1.ValidationService.ValidatePolicySetText:output_type -> commonfate.authz.v1alpha1.ValidatePolicySetTextResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_commonfate_authz_v1alpha1_validation_proto_init() }
func file_commonfate_authz_v1alpha1_validation_proto_init() {
	if File_commonfate_authz_v1alpha1_validation_proto != nil {
		return
	}
	file_commonfate_authz_v1alpha1_read_only_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_authz_v1alpha1_validation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePolicySetTextRequest); i {
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
		file_commonfate_authz_v1alpha1_validation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatePolicySetTextResponse); i {
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
		file_commonfate_authz_v1alpha1_validation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationResult); i {
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
		file_commonfate_authz_v1alpha1_validation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationError); i {
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
		file_commonfate_authz_v1alpha1_validation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationWarning); i {
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
		file_commonfate_authz_v1alpha1_validation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceRange); i {
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
			RawDescriptor: file_commonfate_authz_v1alpha1_validation_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_authz_v1alpha1_validation_proto_goTypes,
		DependencyIndexes: file_commonfate_authz_v1alpha1_validation_proto_depIdxs,
		EnumInfos:         file_commonfate_authz_v1alpha1_validation_proto_enumTypes,
		MessageInfos:      file_commonfate_authz_v1alpha1_validation_proto_msgTypes,
	}.Build()
	File_commonfate_authz_v1alpha1_validation_proto = out.File
	file_commonfate_authz_v1alpha1_validation_proto_rawDesc = nil
	file_commonfate_authz_v1alpha1_validation_proto_goTypes = nil
	file_commonfate_authz_v1alpha1_validation_proto_depIdxs = nil
}