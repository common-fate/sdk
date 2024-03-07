// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: commonfate/authz/v1alpha1/request.proto

package authzv1alpha1

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Principal *v1alpha1.EID `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	Action    *v1alpha1.EID `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Resource  *v1alpha1.EID `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// A client side identifier for the authorization request.
	// The client_key will be included in the corresponding evaluation for the
	// particular request made.
	//
	// Useful when calling BatchAuthorize to match the evaluations with the requests.
	ClientKey string `protobuf:"bytes,4,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	// Entities to 'overlay' temporarily on the entities stored in the authz service.
	OverlayEntities []*v1alpha1.Entity `protobuf:"bytes,5,rep,name=overlay_entities,json=overlayEntities,proto3" json:"overlay_entities,omitempty"`
	// Parent/child relationships to 'overlay' temporarily on the entities stored in the authz service.
	OverlayChildren []*v1alpha1.ChildRelation `protobuf:"bytes,6,rep,name=overlay_children,json=overlayChildren,proto3" json:"overlay_children,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_authz_v1alpha1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_authz_v1alpha1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_commonfate_authz_v1alpha1_request_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetPrincipal() *v1alpha1.EID {
	if x != nil {
		return x.Principal
	}
	return nil
}

func (x *Request) GetAction() *v1alpha1.EID {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *Request) GetResource() *v1alpha1.EID {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Request) GetClientKey() string {
	if x != nil {
		return x.ClientKey
	}
	return ""
}

func (x *Request) GetOverlayEntities() []*v1alpha1.Entity {
	if x != nil {
		return x.OverlayEntities
	}
	return nil
}

func (x *Request) GetOverlayChildren() []*v1alpha1.ChildRelation {
	if x != nil {
		return x.OverlayChildren
	}
	return nil
}

var File_commonfate_authz_v1alpha1_request_proto protoreflect.FileDescriptor

var file_commonfate_authz_v1alpha1_request_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3d, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x49, 0x44, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12,
	0x37, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x4d, 0x0a, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x5f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x5f, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x79, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x42, 0xfb, 0x01, 0x0a, 0x1d, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xe2, 0x02, 0x25, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41,
	0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_authz_v1alpha1_request_proto_rawDescOnce sync.Once
	file_commonfate_authz_v1alpha1_request_proto_rawDescData = file_commonfate_authz_v1alpha1_request_proto_rawDesc
)

func file_commonfate_authz_v1alpha1_request_proto_rawDescGZIP() []byte {
	file_commonfate_authz_v1alpha1_request_proto_rawDescOnce.Do(func() {
		file_commonfate_authz_v1alpha1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_authz_v1alpha1_request_proto_rawDescData)
	})
	return file_commonfate_authz_v1alpha1_request_proto_rawDescData
}

var file_commonfate_authz_v1alpha1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_commonfate_authz_v1alpha1_request_proto_goTypes = []interface{}{
	(*Request)(nil),                // 0: commonfate.authz.v1alpha1.Request
	(*v1alpha1.EID)(nil),           // 1: commonfate.entity.v1alpha1.EID
	(*v1alpha1.Entity)(nil),        // 2: commonfate.entity.v1alpha1.Entity
	(*v1alpha1.ChildRelation)(nil), // 3: commonfate.entity.v1alpha1.ChildRelation
}
var file_commonfate_authz_v1alpha1_request_proto_depIdxs = []int32{
	1, // 0: commonfate.authz.v1alpha1.Request.principal:type_name -> commonfate.entity.v1alpha1.EID
	1, // 1: commonfate.authz.v1alpha1.Request.action:type_name -> commonfate.entity.v1alpha1.EID
	1, // 2: commonfate.authz.v1alpha1.Request.resource:type_name -> commonfate.entity.v1alpha1.EID
	2, // 3: commonfate.authz.v1alpha1.Request.overlay_entities:type_name -> commonfate.entity.v1alpha1.Entity
	3, // 4: commonfate.authz.v1alpha1.Request.overlay_children:type_name -> commonfate.entity.v1alpha1.ChildRelation
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_commonfate_authz_v1alpha1_request_proto_init() }
func file_commonfate_authz_v1alpha1_request_proto_init() {
	if File_commonfate_authz_v1alpha1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_authz_v1alpha1_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_commonfate_authz_v1alpha1_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_authz_v1alpha1_request_proto_goTypes,
		DependencyIndexes: file_commonfate_authz_v1alpha1_request_proto_depIdxs,
		MessageInfos:      file_commonfate_authz_v1alpha1_request_proto_msgTypes,
	}.Build()
	File_commonfate_authz_v1alpha1_request_proto = out.File
	file_commonfate_authz_v1alpha1_request_proto_rawDesc = nil
	file_commonfate_authz_v1alpha1_request_proto_goTypes = nil
	file_commonfate_authz_v1alpha1_request_proto_depIdxs = nil
}
