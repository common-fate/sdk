// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: commonfate/authz/v1alpha1/schema.proto

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

type GetSchemaJSONStringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSchemaJSONStringRequest) Reset() {
	*x = GetSchemaJSONStringRequest{}
	mi := &file_commonfate_authz_v1alpha1_schema_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSchemaJSONStringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaJSONStringRequest) ProtoMessage() {}

func (x *GetSchemaJSONStringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_schema_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaJSONStringRequest.ProtoReflect.Descriptor instead.
func (*GetSchemaJSONStringRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_schema_proto_rawDescGZIP(), []int{0}
}

type GetSchemaJSONStringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *GetSchemaJSONStringResponse) Reset() {
	*x = GetSchemaJSONStringResponse{}
	mi := &file_commonfate_authz_v1alpha1_schema_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSchemaJSONStringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaJSONStringResponse) ProtoMessage() {}

func (x *GetSchemaJSONStringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_schema_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaJSONStringResponse.ProtoReflect.Descriptor instead.
func (*GetSchemaJSONStringResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *GetSchemaJSONStringResponse) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

var File_commonfate_authz_v1alpha1_schema_proto protoreflect.FileDescriptor

var file_commonfate_authz_v1alpha1_schema_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x32, 0x9c, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x35, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5,
	0x18, 0x01, 0x42, 0xfa, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41,
	0x58, 0xaa, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a,
	0x41, 0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_authz_v1alpha1_schema_proto_rawDescOnce sync.Once
	file_commonfate_authz_v1alpha1_schema_proto_rawDescData = file_commonfate_authz_v1alpha1_schema_proto_rawDesc
)

func file_commonfate_authz_v1alpha1_schema_proto_rawDescGZIP() []byte {
	file_commonfate_authz_v1alpha1_schema_proto_rawDescOnce.Do(func() {
		file_commonfate_authz_v1alpha1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_authz_v1alpha1_schema_proto_rawDescData)
	})
	return file_commonfate_authz_v1alpha1_schema_proto_rawDescData
}

var file_commonfate_authz_v1alpha1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commonfate_authz_v1alpha1_schema_proto_goTypes = []any{
	(*GetSchemaJSONStringRequest)(nil),  // 0: commonfate.authz.v1alpha1.GetSchemaJSONStringRequest
	(*GetSchemaJSONStringResponse)(nil), // 1: commonfate.authz.v1alpha1.GetSchemaJSONStringResponse
}
var file_commonfate_authz_v1alpha1_schema_proto_depIdxs = []int32{
	0, // 0: commonfate.authz.v1alpha1.SchemaService.GetSchemaJSONString:input_type -> commonfate.authz.v1alpha1.GetSchemaJSONStringRequest
	1, // 1: commonfate.authz.v1alpha1.SchemaService.GetSchemaJSONString:output_type -> commonfate.authz.v1alpha1.GetSchemaJSONStringResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commonfate_authz_v1alpha1_schema_proto_init() }
func file_commonfate_authz_v1alpha1_schema_proto_init() {
	if File_commonfate_authz_v1alpha1_schema_proto != nil {
		return
	}
	file_commonfate_authz_v1alpha1_read_only_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_authz_v1alpha1_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_authz_v1alpha1_schema_proto_goTypes,
		DependencyIndexes: file_commonfate_authz_v1alpha1_schema_proto_depIdxs,
		MessageInfos:      file_commonfate_authz_v1alpha1_schema_proto_msgTypes,
	}.Build()
	File_commonfate_authz_v1alpha1_schema_proto = out.File
	file_commonfate_authz_v1alpha1_schema_proto_rawDesc = nil
	file_commonfate_authz_v1alpha1_schema_proto_goTypes = nil
	file_commonfate_authz_v1alpha1_schema_proto_depIdxs = nil
}
