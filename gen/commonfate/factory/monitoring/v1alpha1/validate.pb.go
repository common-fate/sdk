// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: commonfate/factory/monitoring/v1alpha1/validate.proto

package monitoringv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidateWriteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WriteToken string `protobuf:"bytes,1,opt,name=write_token,json=writeToken,proto3" json:"write_token,omitempty"`
}

func (x *ValidateWriteTokenRequest) Reset() {
	*x = ValidateWriteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_validate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateWriteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateWriteTokenRequest) ProtoMessage() {}

func (x *ValidateWriteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_validate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateWriteTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateWriteTokenRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateWriteTokenRequest) GetWriteToken() string {
	if x != nil {
		return x.WriteToken
	}
	return ""
}

type ValidateWriteTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The deployment ID associated with the write token.
	DeploymentId string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	// The Common Fate account ID associated with the write token.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *ValidateWriteTokenResponse) Reset() {
	*x = ValidateWriteTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_validate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateWriteTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateWriteTokenResponse) ProtoMessage() {}

func (x *ValidateWriteTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_validate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateWriteTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateWriteTokenResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateWriteTokenResponse) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *ValidateWriteTokenResponse) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

var File_commonfate_factory_monitoring_v1alpha1_validate_proto protoreflect.FileDescriptor

var file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3c, 0x0a, 0x19, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x60,
	0x0a, 0x1a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x32, 0xb3, 0x01, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x41, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd0, 0x02, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x46, 0x4d, 0xaa, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x32, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x29,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x3a, 0x3a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a,
	0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDescOnce sync.Once
	file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDescData = file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDesc
)

func file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDescGZIP() []byte {
	file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDescOnce.Do(func() {
		file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDescData)
	})
	return file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDescData
}

var file_commonfate_factory_monitoring_v1alpha1_validate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commonfate_factory_monitoring_v1alpha1_validate_proto_goTypes = []interface{}{
	(*ValidateWriteTokenRequest)(nil),  // 0: commonfate.factory.monitoring.v1alpha1.ValidateWriteTokenRequest
	(*ValidateWriteTokenResponse)(nil), // 1: commonfate.factory.monitoring.v1alpha1.ValidateWriteTokenResponse
}
var file_commonfate_factory_monitoring_v1alpha1_validate_proto_depIdxs = []int32{
	0, // 0: commonfate.factory.monitoring.v1alpha1.ValidationService.ValidateWriteToken:input_type -> commonfate.factory.monitoring.v1alpha1.ValidateWriteTokenRequest
	1, // 1: commonfate.factory.monitoring.v1alpha1.ValidationService.ValidateWriteToken:output_type -> commonfate.factory.monitoring.v1alpha1.ValidateWriteTokenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commonfate_factory_monitoring_v1alpha1_validate_proto_init() }
func file_commonfate_factory_monitoring_v1alpha1_validate_proto_init() {
	if File_commonfate_factory_monitoring_v1alpha1_validate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_factory_monitoring_v1alpha1_validate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateWriteTokenRequest); i {
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
		file_commonfate_factory_monitoring_v1alpha1_validate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateWriteTokenResponse); i {
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
			RawDescriptor: file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_factory_monitoring_v1alpha1_validate_proto_goTypes,
		DependencyIndexes: file_commonfate_factory_monitoring_v1alpha1_validate_proto_depIdxs,
		MessageInfos:      file_commonfate_factory_monitoring_v1alpha1_validate_proto_msgTypes,
	}.Build()
	File_commonfate_factory_monitoring_v1alpha1_validate_proto = out.File
	file_commonfate_factory_monitoring_v1alpha1_validate_proto_rawDesc = nil
	file_commonfate_factory_monitoring_v1alpha1_validate_proto_goTypes = nil
	file_commonfate_factory_monitoring_v1alpha1_validate_proto_depIdxs = nil
}
