// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commonfate/access/v1alpha1/diagnostic.proto

package accessv1alpha1

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

type DiagnosticLevel int32

const (
	DiagnosticLevel_DIAGNOSTIC_LEVEL_UNSPECIFIED DiagnosticLevel = 0
	DiagnosticLevel_DIAGNOSTIC_LEVEL_INFO        DiagnosticLevel = 1
	DiagnosticLevel_DIAGNOSTIC_LEVEL_WARNING     DiagnosticLevel = 2
	DiagnosticLevel_DIAGNOSTIC_LEVEL_ERROR       DiagnosticLevel = 3
)

// Enum value maps for DiagnosticLevel.
var (
	DiagnosticLevel_name = map[int32]string{
		0: "DIAGNOSTIC_LEVEL_UNSPECIFIED",
		1: "DIAGNOSTIC_LEVEL_INFO",
		2: "DIAGNOSTIC_LEVEL_WARNING",
		3: "DIAGNOSTIC_LEVEL_ERROR",
	}
	DiagnosticLevel_value = map[string]int32{
		"DIAGNOSTIC_LEVEL_UNSPECIFIED": 0,
		"DIAGNOSTIC_LEVEL_INFO":        1,
		"DIAGNOSTIC_LEVEL_WARNING":     2,
		"DIAGNOSTIC_LEVEL_ERROR":       3,
	}
)

func (x DiagnosticLevel) Enum() *DiagnosticLevel {
	p := new(DiagnosticLevel)
	*p = x
	return p
}

func (x DiagnosticLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiagnosticLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_access_v1alpha1_diagnostic_proto_enumTypes[0].Descriptor()
}

func (DiagnosticLevel) Type() protoreflect.EnumType {
	return &file_commonfate_access_v1alpha1_diagnostic_proto_enumTypes[0]
}

func (x DiagnosticLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiagnosticLevel.Descriptor instead.
func (DiagnosticLevel) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_diagnostic_proto_rawDescGZIP(), []int{0}
}

type Diagnostic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level   DiagnosticLevel `protobuf:"varint,1,opt,name=level,proto3,enum=commonfate.access.v1alpha1.DiagnosticLevel" json:"level,omitempty"`
	Message string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Optional timestamp indicating the time at which the diagnostic event occurred.
	// Not all diagnostics will have this attribute.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Diagnostic) Reset() {
	*x = Diagnostic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_diagnostic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Diagnostic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Diagnostic) ProtoMessage() {}

func (x *Diagnostic) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_diagnostic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Diagnostic.ProtoReflect.Descriptor instead.
func (*Diagnostic) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_diagnostic_proto_rawDescGZIP(), []int{0}
}

func (x *Diagnostic) GetLevel() DiagnosticLevel {
	if x != nil {
		return x.Level
	}
	return DiagnosticLevel_DIAGNOSTIC_LEVEL_UNSPECIFIED
}

func (x *Diagnostic) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Diagnostic) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_commonfate_access_v1alpha1_diagnostic_proto protoreflect.FileDescriptor

var file_commonfate_access_v1alpha1_diagnostic_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x61,
	0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x0a, 0x44,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x12, 0x41, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2a, 0x88, 0x01, 0x0a, 0x0f, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x49, 0x41, 0x47, 0x4e, 0x4f, 0x53, 0x54,
	0x49, 0x43, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x49, 0x41, 0x47, 0x4e, 0x4f,
	0x53, 0x54, 0x49, 0x43, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10,
	0x01, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x49, 0x41, 0x47, 0x4e, 0x4f, 0x53, 0x54, 0x49, 0x43, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x1a, 0x0a, 0x16, 0x44, 0x49, 0x41, 0x47, 0x4e, 0x4f, 0x53, 0x54, 0x49, 0x43, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x42, 0x85, 0x02, 0x0a, 0x1e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0f,
	0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41,
	0x58, 0xaa, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02,
	0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x26, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_access_v1alpha1_diagnostic_proto_rawDescOnce sync.Once
	file_commonfate_access_v1alpha1_diagnostic_proto_rawDescData = file_commonfate_access_v1alpha1_diagnostic_proto_rawDesc
)

func file_commonfate_access_v1alpha1_diagnostic_proto_rawDescGZIP() []byte {
	file_commonfate_access_v1alpha1_diagnostic_proto_rawDescOnce.Do(func() {
		file_commonfate_access_v1alpha1_diagnostic_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_v1alpha1_diagnostic_proto_rawDescData)
	})
	return file_commonfate_access_v1alpha1_diagnostic_proto_rawDescData
}

var file_commonfate_access_v1alpha1_diagnostic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_commonfate_access_v1alpha1_diagnostic_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_commonfate_access_v1alpha1_diagnostic_proto_goTypes = []interface{}{
	(DiagnosticLevel)(0),          // 0: commonfate.access.v1alpha1.DiagnosticLevel
	(*Diagnostic)(nil),            // 1: commonfate.access.v1alpha1.Diagnostic
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_commonfate_access_v1alpha1_diagnostic_proto_depIdxs = []int32{
	0, // 0: commonfate.access.v1alpha1.Diagnostic.level:type_name -> commonfate.access.v1alpha1.DiagnosticLevel
	2, // 1: commonfate.access.v1alpha1.Diagnostic.timestamp:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commonfate_access_v1alpha1_diagnostic_proto_init() }
func file_commonfate_access_v1alpha1_diagnostic_proto_init() {
	if File_commonfate_access_v1alpha1_diagnostic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_access_v1alpha1_diagnostic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Diagnostic); i {
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
			RawDescriptor: file_commonfate_access_v1alpha1_diagnostic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_access_v1alpha1_diagnostic_proto_goTypes,
		DependencyIndexes: file_commonfate_access_v1alpha1_diagnostic_proto_depIdxs,
		EnumInfos:         file_commonfate_access_v1alpha1_diagnostic_proto_enumTypes,
		MessageInfos:      file_commonfate_access_v1alpha1_diagnostic_proto_msgTypes,
	}.Build()
	File_commonfate_access_v1alpha1_diagnostic_proto = out.File
	file_commonfate_access_v1alpha1_diagnostic_proto_rawDesc = nil
	file_commonfate_access_v1alpha1_diagnostic_proto_goTypes = nil
	file_commonfate_access_v1alpha1_diagnostic_proto_depIdxs = nil
}