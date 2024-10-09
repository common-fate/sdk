// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: commonfate/access/v1alpha1/regex_validation.proto

package accessv1alpha1

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

type RegexValidation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The regex pattern to that the reason should match on. It should be in a regex format similar to  "/CF-\d+/".
	RegexPattern string `protobuf:"bytes,1,opt,name=regex_pattern,json=regexPattern,proto3" json:"regex_pattern,omitempty"`
	// The custom error message to show if the reason doesn't match the regex pattern
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *RegexValidation) Reset() {
	*x = RegexValidation{}
	mi := &file_commonfate_access_v1alpha1_regex_validation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegexValidation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegexValidation) ProtoMessage() {}

func (x *RegexValidation) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_regex_validation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegexValidation.ProtoReflect.Descriptor instead.
func (*RegexValidation) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_regex_validation_proto_rawDescGZIP(), []int{0}
}

func (x *RegexValidation) GetRegexPattern() string {
	if x != nil {
		return x.RegexPattern
	}
	return ""
}

func (x *RegexValidation) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_commonfate_access_v1alpha1_regex_validation_proto protoreflect.FileDescriptor

var file_commonfate_access_v1alpha1_regex_validation_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x67,
	0x65, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22,
	0x5b, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x65, 0x78, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x65, 0x78, 0x5f, 0x70, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x65, 0x78,
	0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x8a, 0x02, 0x0a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x14, 0x52, 0x65, 0x67, 0x65, 0x78, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xe2, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_commonfate_access_v1alpha1_regex_validation_proto_rawDescOnce sync.Once
	file_commonfate_access_v1alpha1_regex_validation_proto_rawDescData = file_commonfate_access_v1alpha1_regex_validation_proto_rawDesc
)

func file_commonfate_access_v1alpha1_regex_validation_proto_rawDescGZIP() []byte {
	file_commonfate_access_v1alpha1_regex_validation_proto_rawDescOnce.Do(func() {
		file_commonfate_access_v1alpha1_regex_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_v1alpha1_regex_validation_proto_rawDescData)
	})
	return file_commonfate_access_v1alpha1_regex_validation_proto_rawDescData
}

var file_commonfate_access_v1alpha1_regex_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_commonfate_access_v1alpha1_regex_validation_proto_goTypes = []any{
	(*RegexValidation)(nil), // 0: commonfate.access.v1alpha1.RegexValidation
}
var file_commonfate_access_v1alpha1_regex_validation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commonfate_access_v1alpha1_regex_validation_proto_init() }
func file_commonfate_access_v1alpha1_regex_validation_proto_init() {
	if File_commonfate_access_v1alpha1_regex_validation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_access_v1alpha1_regex_validation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_access_v1alpha1_regex_validation_proto_goTypes,
		DependencyIndexes: file_commonfate_access_v1alpha1_regex_validation_proto_depIdxs,
		MessageInfos:      file_commonfate_access_v1alpha1_regex_validation_proto_msgTypes,
	}.Build()
	File_commonfate_access_v1alpha1_regex_validation_proto = out.File
	file_commonfate_access_v1alpha1_regex_validation_proto_rawDesc = nil
	file_commonfate_access_v1alpha1_regex_validation_proto_goTypes = nil
	file_commonfate_access_v1alpha1_regex_validation_proto_depIdxs = nil
}
