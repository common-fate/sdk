// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commonfate/access/v1alpha1/identity.proto

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

type GetCallerIdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCallerIdentityRequest) Reset() {
	*x = GetCallerIdentityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCallerIdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCallerIdentityRequest) ProtoMessage() {}

func (x *GetCallerIdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCallerIdentityRequest.ProtoReflect.Descriptor instead.
func (*GetCallerIdentityRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_identity_proto_rawDescGZIP(), []int{0}
}

type GetCallerIdentityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The chain of the identities for the user.
	// The final item in the chain is the current identity the user is acting as,
	// and is the 'principal' for authorization policy evaluations.
	//
	// For users authenticating with OIDC, the chain will usually look like:
	// 1. OIDC::Subject::"https://oidc-issuer.example.com/93fcac4b-ab67-405c-926b-184f8ba697a4"
	// 2. CF::User::"usr_2Z0WQkw9Ogpsn2Us6RuIBiDRYzJ"
	Chain []*IdentityLink `protobuf:"bytes,1,rep,name=chain,proto3" json:"chain,omitempty"`
	// The principal that the user is currently acting as.
	Principal *User `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
}

func (x *GetCallerIdentityResponse) Reset() {
	*x = GetCallerIdentityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCallerIdentityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCallerIdentityResponse) ProtoMessage() {}

func (x *GetCallerIdentityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCallerIdentityResponse.ProtoReflect.Descriptor instead.
func (*GetCallerIdentityResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_identity_proto_rawDescGZIP(), []int{1}
}

func (x *GetCallerIdentityResponse) GetChain() []*IdentityLink {
	if x != nil {
		return x.Chain
	}
	return nil
}

func (x *GetCallerIdentityResponse) GetPrincipal() *User {
	if x != nil {
		return x.Principal
	}
	return nil
}

type IdentityLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *v1alpha1.UID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string        `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *IdentityLink) Reset() {
	*x = IdentityLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityLink) ProtoMessage() {}

func (x *IdentityLink) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityLink.ProtoReflect.Descriptor instead.
func (*IdentityLink) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_identity_proto_rawDescGZIP(), []int{2}
}

func (x *IdentityLink) GetId() *v1alpha1.UID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *IdentityLink) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_commonfate_access_v1alpha1_identity_proto protoreflect.FileDescriptor

var file_commonfate_access_v1alpha1_identity_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x9b, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x3e,
	0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x22, 0x55,
	0x0a, 0x0c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2f,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x32, 0x96, 0x01, 0x0a, 0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x83,
	0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x0d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x41, 0x58, 0xaa, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x26, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_access_v1alpha1_identity_proto_rawDescOnce sync.Once
	file_commonfate_access_v1alpha1_identity_proto_rawDescData = file_commonfate_access_v1alpha1_identity_proto_rawDesc
)

func file_commonfate_access_v1alpha1_identity_proto_rawDescGZIP() []byte {
	file_commonfate_access_v1alpha1_identity_proto_rawDescOnce.Do(func() {
		file_commonfate_access_v1alpha1_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_v1alpha1_identity_proto_rawDescData)
	})
	return file_commonfate_access_v1alpha1_identity_proto_rawDescData
}

var file_commonfate_access_v1alpha1_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commonfate_access_v1alpha1_identity_proto_goTypes = []interface{}{
	(*GetCallerIdentityRequest)(nil),  // 0: commonfate.access.v1alpha1.GetCallerIdentityRequest
	(*GetCallerIdentityResponse)(nil), // 1: commonfate.access.v1alpha1.GetCallerIdentityResponse
	(*IdentityLink)(nil),              // 2: commonfate.access.v1alpha1.IdentityLink
	(*User)(nil),                      // 3: commonfate.access.v1alpha1.User
	(*v1alpha1.UID)(nil),              // 4: commonfate.entity.v1alpha1.UID
}
var file_commonfate_access_v1alpha1_identity_proto_depIdxs = []int32{
	2, // 0: commonfate.access.v1alpha1.GetCallerIdentityResponse.chain:type_name -> commonfate.access.v1alpha1.IdentityLink
	3, // 1: commonfate.access.v1alpha1.GetCallerIdentityResponse.principal:type_name -> commonfate.access.v1alpha1.User
	4, // 2: commonfate.access.v1alpha1.IdentityLink.id:type_name -> commonfate.entity.v1alpha1.UID
	0, // 3: commonfate.access.v1alpha1.IdentityService.GetCallerIdentity:input_type -> commonfate.access.v1alpha1.GetCallerIdentityRequest
	1, // 4: commonfate.access.v1alpha1.IdentityService.GetCallerIdentity:output_type -> commonfate.access.v1alpha1.GetCallerIdentityResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_commonfate_access_v1alpha1_identity_proto_init() }
func file_commonfate_access_v1alpha1_identity_proto_init() {
	if File_commonfate_access_v1alpha1_identity_proto != nil {
		return
	}
	file_commonfate_access_v1alpha1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_access_v1alpha1_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCallerIdentityRequest); i {
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
		file_commonfate_access_v1alpha1_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCallerIdentityResponse); i {
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
		file_commonfate_access_v1alpha1_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityLink); i {
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
			RawDescriptor: file_commonfate_access_v1alpha1_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_access_v1alpha1_identity_proto_goTypes,
		DependencyIndexes: file_commonfate_access_v1alpha1_identity_proto_depIdxs,
		MessageInfos:      file_commonfate_access_v1alpha1_identity_proto_msgTypes,
	}.Build()
	File_commonfate_access_v1alpha1_identity_proto = out.File
	file_commonfate_access_v1alpha1_identity_proto_rawDesc = nil
	file_commonfate_access_v1alpha1_identity_proto_goTypes = nil
	file_commonfate_access_v1alpha1_identity_proto_depIdxs = nil
}
