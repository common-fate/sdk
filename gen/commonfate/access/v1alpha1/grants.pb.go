// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: commonfate/access/v1alpha1/grants.proto

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

type QueryGrantsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The token for the next page.
	PageToken string        `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	Principal *v1alpha1.EID `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
	Target    *v1alpha1.EID `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Role      *v1alpha1.EID `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	// The status of the Grant.
	Status GrantStatus `protobuf:"varint,5,opt,name=status,proto3,enum=commonfate.access.v1alpha1.GrantStatus" json:"status,omitempty"`
}

func (x *QueryGrantsRequest) Reset() {
	*x = QueryGrantsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_grants_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGrantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGrantsRequest) ProtoMessage() {}

func (x *QueryGrantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_grants_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGrantsRequest.ProtoReflect.Descriptor instead.
func (*QueryGrantsRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_grants_proto_rawDescGZIP(), []int{0}
}

func (x *QueryGrantsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *QueryGrantsRequest) GetPrincipal() *v1alpha1.EID {
	if x != nil {
		return x.Principal
	}
	return nil
}

func (x *QueryGrantsRequest) GetTarget() *v1alpha1.EID {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *QueryGrantsRequest) GetRole() *v1alpha1.EID {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *QueryGrantsRequest) GetStatus() GrantStatus {
	if x != nil {
		return x.Status
	}
	return GrantStatus_GRANT_STATUS_UNSPECIFIED
}

type QueryGrantsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grants        []*Grant `protobuf:"bytes,1,rep,name=grants,proto3" json:"grants,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *QueryGrantsResponse) Reset() {
	*x = QueryGrantsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_grants_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGrantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGrantsResponse) ProtoMessage() {}

func (x *QueryGrantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_grants_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGrantsResponse.ProtoReflect.Descriptor instead.
func (*QueryGrantsResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_grants_proto_rawDescGZIP(), []int{1}
}

func (x *QueryGrantsResponse) GetGrants() []*Grant {
	if x != nil {
		return x.Grants
	}
	return nil
}

func (x *QueryGrantsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type QueryGrantChildrenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The token for the next page.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *QueryGrantChildrenRequest) Reset() {
	*x = QueryGrantChildrenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_grants_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGrantChildrenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGrantChildrenRequest) ProtoMessage() {}

func (x *QueryGrantChildrenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_grants_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGrantChildrenRequest.ProtoReflect.Descriptor instead.
func (*QueryGrantChildrenRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_grants_proto_rawDescGZIP(), []int{2}
}

func (x *QueryGrantChildrenRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QueryGrantChildrenRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type QueryGrantChildrenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities      []*v1alpha1.Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	NextPageToken string             `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *QueryGrantChildrenResponse) Reset() {
	*x = QueryGrantChildrenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_grants_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryGrantChildrenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryGrantChildrenResponse) ProtoMessage() {}

func (x *QueryGrantChildrenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_grants_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryGrantChildrenResponse.ProtoReflect.Descriptor instead.
func (*QueryGrantChildrenResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_grants_proto_rawDescGZIP(), []int{3}
}

func (x *QueryGrantChildrenResponse) GetEntities() []*v1alpha1.Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *QueryGrantChildrenResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_commonfate_access_v1alpha1_grants_proto protoreflect.FileDescriptor

var file_commonfate_access_v1alpha1_grants_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3d, 0x0a, 0x09, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45,
	0x49, 0x44, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x78, 0x0a, 0x13, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x06, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x52, 0x06, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x4a, 0x0a, 0x19, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x84, 0x01, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x89, 0x02, 0x0a, 0x0d, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x12, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x12, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x81, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xca, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_access_v1alpha1_grants_proto_rawDescOnce sync.Once
	file_commonfate_access_v1alpha1_grants_proto_rawDescData = file_commonfate_access_v1alpha1_grants_proto_rawDesc
)

func file_commonfate_access_v1alpha1_grants_proto_rawDescGZIP() []byte {
	file_commonfate_access_v1alpha1_grants_proto_rawDescOnce.Do(func() {
		file_commonfate_access_v1alpha1_grants_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_v1alpha1_grants_proto_rawDescData)
	})
	return file_commonfate_access_v1alpha1_grants_proto_rawDescData
}

var file_commonfate_access_v1alpha1_grants_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_commonfate_access_v1alpha1_grants_proto_goTypes = []interface{}{
	(*QueryGrantsRequest)(nil),         // 0: commonfate.access.v1alpha1.QueryGrantsRequest
	(*QueryGrantsResponse)(nil),        // 1: commonfate.access.v1alpha1.QueryGrantsResponse
	(*QueryGrantChildrenRequest)(nil),  // 2: commonfate.access.v1alpha1.QueryGrantChildrenRequest
	(*QueryGrantChildrenResponse)(nil), // 3: commonfate.access.v1alpha1.QueryGrantChildrenResponse
	(*v1alpha1.EID)(nil),               // 4: commonfate.entity.v1alpha1.EID
	(GrantStatus)(0),                   // 5: commonfate.access.v1alpha1.GrantStatus
	(*Grant)(nil),                      // 6: commonfate.access.v1alpha1.Grant
	(*v1alpha1.Entity)(nil),            // 7: commonfate.entity.v1alpha1.Entity
}
var file_commonfate_access_v1alpha1_grants_proto_depIdxs = []int32{
	4, // 0: commonfate.access.v1alpha1.QueryGrantsRequest.principal:type_name -> commonfate.entity.v1alpha1.EID
	4, // 1: commonfate.access.v1alpha1.QueryGrantsRequest.target:type_name -> commonfate.entity.v1alpha1.EID
	4, // 2: commonfate.access.v1alpha1.QueryGrantsRequest.role:type_name -> commonfate.entity.v1alpha1.EID
	5, // 3: commonfate.access.v1alpha1.QueryGrantsRequest.status:type_name -> commonfate.access.v1alpha1.GrantStatus
	6, // 4: commonfate.access.v1alpha1.QueryGrantsResponse.grants:type_name -> commonfate.access.v1alpha1.Grant
	7, // 5: commonfate.access.v1alpha1.QueryGrantChildrenResponse.entities:type_name -> commonfate.entity.v1alpha1.Entity
	0, // 6: commonfate.access.v1alpha1.GrantsService.QueryGrants:input_type -> commonfate.access.v1alpha1.QueryGrantsRequest
	2, // 7: commonfate.access.v1alpha1.GrantsService.QueryGrantChildren:input_type -> commonfate.access.v1alpha1.QueryGrantChildrenRequest
	1, // 8: commonfate.access.v1alpha1.GrantsService.QueryGrants:output_type -> commonfate.access.v1alpha1.QueryGrantsResponse
	3, // 9: commonfate.access.v1alpha1.GrantsService.QueryGrantChildren:output_type -> commonfate.access.v1alpha1.QueryGrantChildrenResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_commonfate_access_v1alpha1_grants_proto_init() }
func file_commonfate_access_v1alpha1_grants_proto_init() {
	if File_commonfate_access_v1alpha1_grants_proto != nil {
		return
	}
	file_commonfate_access_v1alpha1_access_proto_init()
	file_commonfate_access_v1alpha1_grant_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_access_v1alpha1_grants_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGrantsRequest); i {
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
		file_commonfate_access_v1alpha1_grants_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGrantsResponse); i {
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
		file_commonfate_access_v1alpha1_grants_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGrantChildrenRequest); i {
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
		file_commonfate_access_v1alpha1_grants_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryGrantChildrenResponse); i {
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
			RawDescriptor: file_commonfate_access_v1alpha1_grants_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_access_v1alpha1_grants_proto_goTypes,
		DependencyIndexes: file_commonfate_access_v1alpha1_grants_proto_depIdxs,
		MessageInfos:      file_commonfate_access_v1alpha1_grants_proto_msgTypes,
	}.Build()
	File_commonfate_access_v1alpha1_grants_proto = out.File
	file_commonfate_access_v1alpha1_grants_proto_rawDesc = nil
	file_commonfate_access_v1alpha1_grants_proto_goTypes = nil
	file_commonfate_access_v1alpha1_grants_proto_depIdxs = nil
}