// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: commonfate/access/v1alpha1/named_eid.proto

package accessv1alpha1

import (
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
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

type NamedEID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eid  *v1alpha1.EID `protobuf:"bytes,1,opt,name=eid,proto3" json:"eid,omitempty"`
	Name string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NamedEID) Reset() {
	*x = NamedEID{}
	mi := &file_commonfate_access_v1alpha1_named_eid_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamedEID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamedEID) ProtoMessage() {}

func (x *NamedEID) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_named_eid_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamedEID.ProtoReflect.Descriptor instead.
func (*NamedEID) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_named_eid_proto_rawDescGZIP(), []int{0}
}

func (x *NamedEID) GetEid() *v1alpha1.EID {
	if x != nil {
		return x.Eid
	}
	return nil
}

func (x *NamedEID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_commonfate_access_v1alpha1_named_eid_proto protoreflect.FileDescriptor

var file_commonfate_access_v1alpha1_named_eid_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x61, 0x6d,
	0x65, 0x64, 0x5f, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51,
	0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x45, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x03, 0x65, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x03, 0x65, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x83, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x45, 0x69, 0x64, 0x50, 0x72,
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
	file_commonfate_access_v1alpha1_named_eid_proto_rawDescOnce sync.Once
	file_commonfate_access_v1alpha1_named_eid_proto_rawDescData = file_commonfate_access_v1alpha1_named_eid_proto_rawDesc
)

func file_commonfate_access_v1alpha1_named_eid_proto_rawDescGZIP() []byte {
	file_commonfate_access_v1alpha1_named_eid_proto_rawDescOnce.Do(func() {
		file_commonfate_access_v1alpha1_named_eid_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_v1alpha1_named_eid_proto_rawDescData)
	})
	return file_commonfate_access_v1alpha1_named_eid_proto_rawDescData
}

var file_commonfate_access_v1alpha1_named_eid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_commonfate_access_v1alpha1_named_eid_proto_goTypes = []any{
	(*NamedEID)(nil),     // 0: commonfate.access.v1alpha1.NamedEID
	(*v1alpha1.EID)(nil), // 1: commonfate.entity.v1alpha1.EID
}
var file_commonfate_access_v1alpha1_named_eid_proto_depIdxs = []int32{
	1, // 0: commonfate.access.v1alpha1.NamedEID.eid:type_name -> commonfate.entity.v1alpha1.EID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_commonfate_access_v1alpha1_named_eid_proto_init() }
func file_commonfate_access_v1alpha1_named_eid_proto_init() {
	if File_commonfate_access_v1alpha1_named_eid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_access_v1alpha1_named_eid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_access_v1alpha1_named_eid_proto_goTypes,
		DependencyIndexes: file_commonfate_access_v1alpha1_named_eid_proto_depIdxs,
		MessageInfos:      file_commonfate_access_v1alpha1_named_eid_proto_msgTypes,
	}.Build()
	File_commonfate_access_v1alpha1_named_eid_proto = out.File
	file_commonfate_access_v1alpha1_named_eid_proto_rawDesc = nil
	file_commonfate_access_v1alpha1_named_eid_proto_goTypes = nil
	file_commonfate_access_v1alpha1_named_eid_proto_depIdxs = nil
}
