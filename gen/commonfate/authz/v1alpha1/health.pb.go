// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: commonfate/authz/v1alpha1/health.proto

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

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_health_proto_rawDescGZIP(), []int{0}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_health_proto_rawDescGZIP(), []int{1}
}

var File_commonfate_authz_v1alpha1_health_proto protoreflect.FileDescriptor

var file_commonfate_authz_v1alpha1_health_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14,
	0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x83, 0x01, 0x0a, 0x0d,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x2d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x88, 0xb5, 0x18,
	0x01, 0x42, 0xfa, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x42, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f,
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
	file_commonfate_authz_v1alpha1_health_proto_rawDescOnce sync.Once
	file_commonfate_authz_v1alpha1_health_proto_rawDescData = file_commonfate_authz_v1alpha1_health_proto_rawDesc
)

func file_commonfate_authz_v1alpha1_health_proto_rawDescGZIP() []byte {
	file_commonfate_authz_v1alpha1_health_proto_rawDescOnce.Do(func() {
		file_commonfate_authz_v1alpha1_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_authz_v1alpha1_health_proto_rawDescData)
	})
	return file_commonfate_authz_v1alpha1_health_proto_rawDescData
}

var file_commonfate_authz_v1alpha1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commonfate_authz_v1alpha1_health_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil),  // 0: commonfate.authz.v1alpha1.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 1: commonfate.authz.v1alpha1.HealthCheckResponse
}
var file_commonfate_authz_v1alpha1_health_proto_depIdxs = []int32{
	0, // 0: commonfate.authz.v1alpha1.HealthService.HealthCheck:input_type -> commonfate.authz.v1alpha1.HealthCheckRequest
	1, // 1: commonfate.authz.v1alpha1.HealthService.HealthCheck:output_type -> commonfate.authz.v1alpha1.HealthCheckResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commonfate_authz_v1alpha1_health_proto_init() }
func file_commonfate_authz_v1alpha1_health_proto_init() {
	if File_commonfate_authz_v1alpha1_health_proto != nil {
		return
	}
	file_commonfate_authz_v1alpha1_read_only_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_authz_v1alpha1_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_commonfate_authz_v1alpha1_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_commonfate_authz_v1alpha1_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_authz_v1alpha1_health_proto_goTypes,
		DependencyIndexes: file_commonfate_authz_v1alpha1_health_proto_depIdxs,
		MessageInfos:      file_commonfate_authz_v1alpha1_health_proto_msgTypes,
	}.Build()
	File_commonfate_authz_v1alpha1_health_proto = out.File
	file_commonfate_authz_v1alpha1_health_proto_rawDesc = nil
	file_commonfate_authz_v1alpha1_health_proto_goTypes = nil
	file_commonfate_authz_v1alpha1_health_proto_depIdxs = nil
}
