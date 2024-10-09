// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: commonfate/control/feature/v1alpha1/feature.proto

package featurev1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
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

type GetFeaturesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFeaturesRequest) Reset() {
	*x = GetFeaturesRequest{}
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeaturesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeaturesRequest) ProtoMessage() {}

func (x *GetFeaturesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeaturesRequest.ProtoReflect.Descriptor instead.
func (*GetFeaturesRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_feature_v1alpha1_feature_proto_rawDescGZIP(), []int{0}
}

type GetFeaturesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features []*Feature `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *GetFeaturesResponse) Reset() {
	*x = GetFeaturesResponse{}
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeaturesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeaturesResponse) ProtoMessage() {}

func (x *GetFeaturesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeaturesResponse.ProtoReflect.Descriptor instead.
func (*GetFeaturesResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_feature_v1alpha1_feature_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeaturesResponse) GetFeatures() []*Feature {
	if x != nil {
		return x.Features
	}
	return nil
}

type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the feature. By convention this is 'snake_case',
	// for example: "least_privilege_analytics"
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_commonfate_control_feature_v1alpha1_feature_proto_rawDescGZIP(), []int{2}
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feature) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type UpdateFeatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *UpdateFeatureRequest) Reset() {
	*x = UpdateFeatureRequest{}
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFeatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeatureRequest) ProtoMessage() {}

func (x *UpdateFeatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeatureRequest.ProtoReflect.Descriptor instead.
func (*UpdateFeatureRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_feature_v1alpha1_feature_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFeatureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateFeatureRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type UpdateFeatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature *Feature `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *UpdateFeatureResponse) Reset() {
	*x = UpdateFeatureResponse{}
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateFeatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeatureResponse) ProtoMessage() {}

func (x *UpdateFeatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeatureResponse.ProtoReflect.Descriptor instead.
func (*UpdateFeatureResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_feature_v1alpha1_feature_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateFeatureResponse) GetFeature() *Feature {
	if x != nil {
		return x.Feature
	}
	return nil
}

var File_commonfate_control_feature_v1alpha1_feature_proto protoreflect.FileDescriptor

var file_commonfate_control_feature_v1alpha1_feature_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x07, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x22, 0x44, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x5f, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0xa4, 0x02, 0x0a, 0x0e, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x12, 0x88, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xba, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x46, 0xaa, 0x02, 0x23, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x23,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x5c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_feature_v1alpha1_feature_proto_rawDescOnce sync.Once
	file_commonfate_control_feature_v1alpha1_feature_proto_rawDescData = file_commonfate_control_feature_v1alpha1_feature_proto_rawDesc
)

func file_commonfate_control_feature_v1alpha1_feature_proto_rawDescGZIP() []byte {
	file_commonfate_control_feature_v1alpha1_feature_proto_rawDescOnce.Do(func() {
		file_commonfate_control_feature_v1alpha1_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_feature_v1alpha1_feature_proto_rawDescData)
	})
	return file_commonfate_control_feature_v1alpha1_feature_proto_rawDescData
}

var file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_commonfate_control_feature_v1alpha1_feature_proto_goTypes = []any{
	(*GetFeaturesRequest)(nil),    // 0: commonfate.control.feature.v1alpha1.GetFeaturesRequest
	(*GetFeaturesResponse)(nil),   // 1: commonfate.control.feature.v1alpha1.GetFeaturesResponse
	(*Feature)(nil),               // 2: commonfate.control.feature.v1alpha1.Feature
	(*UpdateFeatureRequest)(nil),  // 3: commonfate.control.feature.v1alpha1.UpdateFeatureRequest
	(*UpdateFeatureResponse)(nil), // 4: commonfate.control.feature.v1alpha1.UpdateFeatureResponse
}
var file_commonfate_control_feature_v1alpha1_feature_proto_depIdxs = []int32{
	2, // 0: commonfate.control.feature.v1alpha1.GetFeaturesResponse.features:type_name -> commonfate.control.feature.v1alpha1.Feature
	2, // 1: commonfate.control.feature.v1alpha1.UpdateFeatureResponse.feature:type_name -> commonfate.control.feature.v1alpha1.Feature
	0, // 2: commonfate.control.feature.v1alpha1.FeatureService.GetFeatures:input_type -> commonfate.control.feature.v1alpha1.GetFeaturesRequest
	3, // 3: commonfate.control.feature.v1alpha1.FeatureService.UpdateFeature:input_type -> commonfate.control.feature.v1alpha1.UpdateFeatureRequest
	1, // 4: commonfate.control.feature.v1alpha1.FeatureService.GetFeatures:output_type -> commonfate.control.feature.v1alpha1.GetFeaturesResponse
	4, // 5: commonfate.control.feature.v1alpha1.FeatureService.UpdateFeature:output_type -> commonfate.control.feature.v1alpha1.UpdateFeatureResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commonfate_control_feature_v1alpha1_feature_proto_init() }
func file_commonfate_control_feature_v1alpha1_feature_proto_init() {
	if File_commonfate_control_feature_v1alpha1_feature_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_control_feature_v1alpha1_feature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_feature_v1alpha1_feature_proto_goTypes,
		DependencyIndexes: file_commonfate_control_feature_v1alpha1_feature_proto_depIdxs,
		MessageInfos:      file_commonfate_control_feature_v1alpha1_feature_proto_msgTypes,
	}.Build()
	File_commonfate_control_feature_v1alpha1_feature_proto = out.File
	file_commonfate_control_feature_v1alpha1_feature_proto_rawDesc = nil
	file_commonfate_control_feature_v1alpha1_feature_proto_goTypes = nil
	file_commonfate_control_feature_v1alpha1_feature_proto_depIdxs = nil
}
