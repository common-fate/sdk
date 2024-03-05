// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: commonfate/control/integration/v1alpha1/reset.proto

package integrationv1alpha1

import (
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResetEntraUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, will return a preview of changes rather than actually resetting them.
	DryRun bool `protobuf:"varint,1,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
}

func (x *ResetEntraUsersRequest) Reset() {
	*x = ResetEntraUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_reset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetEntraUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetEntraUsersRequest) ProtoMessage() {}

func (x *ResetEntraUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_reset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetEntraUsersRequest.ProtoReflect.Descriptor instead.
func (*ResetEntraUsersRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_reset_proto_rawDescGZIP(), []int{0}
}

func (x *ResetEntraUsersRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

type ResetEntraUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeletedEntities []*v1alpha1.EID `protobuf:"bytes,1,rep,name=deleted_entities,json=deletedEntities,proto3" json:"deleted_entities,omitempty"`
}

func (x *ResetEntraUsersResponse) Reset() {
	*x = ResetEntraUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_reset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetEntraUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetEntraUsersResponse) ProtoMessage() {}

func (x *ResetEntraUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_reset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetEntraUsersResponse.ProtoReflect.Descriptor instead.
func (*ResetEntraUsersResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_reset_proto_rawDescGZIP(), []int{1}
}

func (x *ResetEntraUsersResponse) GetDeletedEntities() []*v1alpha1.EID {
	if x != nil {
		return x.DeletedEntities
	}
	return nil
}

var File_commonfate_control_integration_v1alpha1_reset_proto protoreflect.FileDescriptor

var file_commonfate_control_integration_v1alpha1_reset_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x64, 0x72, 0x79, 0x52, 0x75, 0x6e, 0x22, 0x65, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x0f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0xa7,
	0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x96, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x3f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x61, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd4, 0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x49, 0xaa, 0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x33, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x2a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a,
	0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_integration_v1alpha1_reset_proto_rawDescOnce sync.Once
	file_commonfate_control_integration_v1alpha1_reset_proto_rawDescData = file_commonfate_control_integration_v1alpha1_reset_proto_rawDesc
)

func file_commonfate_control_integration_v1alpha1_reset_proto_rawDescGZIP() []byte {
	file_commonfate_control_integration_v1alpha1_reset_proto_rawDescOnce.Do(func() {
		file_commonfate_control_integration_v1alpha1_reset_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_integration_v1alpha1_reset_proto_rawDescData)
	})
	return file_commonfate_control_integration_v1alpha1_reset_proto_rawDescData
}

var file_commonfate_control_integration_v1alpha1_reset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commonfate_control_integration_v1alpha1_reset_proto_goTypes = []interface{}{
	(*ResetEntraUsersRequest)(nil),  // 0: commonfate.control.integration.v1alpha1.ResetEntraUsersRequest
	(*ResetEntraUsersResponse)(nil), // 1: commonfate.control.integration.v1alpha1.ResetEntraUsersResponse
	(*v1alpha1.EID)(nil),            // 2: commonfate.entity.v1alpha1.EID
}
var file_commonfate_control_integration_v1alpha1_reset_proto_depIdxs = []int32{
	2, // 0: commonfate.control.integration.v1alpha1.ResetEntraUsersResponse.deleted_entities:type_name -> commonfate.entity.v1alpha1.EID
	0, // 1: commonfate.control.integration.v1alpha1.ResetService.ResetEntraUsers:input_type -> commonfate.control.integration.v1alpha1.ResetEntraUsersRequest
	1, // 2: commonfate.control.integration.v1alpha1.ResetService.ResetEntraUsers:output_type -> commonfate.control.integration.v1alpha1.ResetEntraUsersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_commonfate_control_integration_v1alpha1_reset_proto_init() }
func file_commonfate_control_integration_v1alpha1_reset_proto_init() {
	if File_commonfate_control_integration_v1alpha1_reset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_integration_v1alpha1_reset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetEntraUsersRequest); i {
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
		file_commonfate_control_integration_v1alpha1_reset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetEntraUsersResponse); i {
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
			RawDescriptor: file_commonfate_control_integration_v1alpha1_reset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_integration_v1alpha1_reset_proto_goTypes,
		DependencyIndexes: file_commonfate_control_integration_v1alpha1_reset_proto_depIdxs,
		MessageInfos:      file_commonfate_control_integration_v1alpha1_reset_proto_msgTypes,
	}.Build()
	File_commonfate_control_integration_v1alpha1_reset_proto = out.File
	file_commonfate_control_integration_v1alpha1_reset_proto_rawDesc = nil
	file_commonfate_control_integration_v1alpha1_reset_proto_goTypes = nil
	file_commonfate_control_integration_v1alpha1_reset_proto_depIdxs = nil
}
