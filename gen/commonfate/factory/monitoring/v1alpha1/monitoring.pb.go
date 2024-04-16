// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: commonfate/factory/monitoring/v1alpha1/monitoring.proto

package deployidv1alpha1

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

type AWSClaims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature     string                 `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SecurityToken string                 `protobuf:"bytes,3,opt,name=security_token,json=securityToken,proto3" json:"security_token,omitempty"`
}

func (x *AWSClaims) Reset() {
	*x = AWSClaims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSClaims) ProtoMessage() {}

func (x *AWSClaims) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSClaims.ProtoReflect.Descriptor instead.
func (*AWSClaims) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescGZIP(), []int{0}
}

func (x *AWSClaims) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *AWSClaims) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AWSClaims) GetSecurityToken() string {
	if x != nil {
		return x.SecurityToken
	}
	return ""
}

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the Common Fate licence key.
	LicenceKey string `protobuf:"bytes,1,opt,name=licence_key,json=licenceKey,proto3" json:"licence_key,omitempty"`
	// a name for the deployment, like 'prod'.
	// used to differentiate deployments residing in the same AWS account.
	DeploymentName string `protobuf:"bytes,2,opt,name=deployment_name,json=deploymentName,proto3" json:"deployment_name,omitempty"`
	// Types that are assignable to Claims:
	//
	//	*AuthenticateRequest_Aws
	Claims isAuthenticateRequest_Claims `protobuf_oneof:"claims"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateRequest) GetLicenceKey() string {
	if x != nil {
		return x.LicenceKey
	}
	return ""
}

func (x *AuthenticateRequest) GetDeploymentName() string {
	if x != nil {
		return x.DeploymentName
	}
	return ""
}

func (m *AuthenticateRequest) GetClaims() isAuthenticateRequest_Claims {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (x *AuthenticateRequest) GetAws() *AWSClaims {
	if x, ok := x.GetClaims().(*AuthenticateRequest_Aws); ok {
		return x.Aws
	}
	return nil
}

type isAuthenticateRequest_Claims interface {
	isAuthenticateRequest_Claims()
}

type AuthenticateRequest_Aws struct {
	Aws *AWSClaims `protobuf:"bytes,3,opt,name=aws,proto3,oneof"`
}

func (*AuthenticateRequest_Aws) isAuthenticateRequest_Claims() {}

type AuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A write token which gives access to write OpenTelemetry data to our collector.
	MonitoringToken string `protobuf:"bytes,1,opt,name=monitoring_token,json=monitoringToken,proto3" json:"monitoring_token,omitempty"`
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticateResponse) GetMonitoringToken() string {
	if x != nil {
		return x.MonitoringToken
	}
	return ""
}

var File_commonfate_factory_monitoring_v1alpha1_monitoring_proto protoreflect.FileDescriptor

var file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8a, 0x01, 0x0a, 0x09, 0x41, 0x57, 0x53, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xae, 0x01,
	0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x63, 0x65,
	0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x43, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x48, 0x00, 0x52,
	0x03, 0x61, 0x77, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x22, 0x41,
	0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x32, 0x9d, 0x01, 0x0a, 0x11, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x39, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x69,
	0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xc6, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0f,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x69,
	0x64, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x46, 0x44, 0xaa,
	0x02, 0x24, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x69, 0x64, 0x2e, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x24, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x5c, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5c, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x69, 0x64, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x30,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x46, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x5c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x69, 0x64, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x69, 0x64,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescOnce sync.Once
	file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescData = file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDesc
)

func file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescGZIP() []byte {
	file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescOnce.Do(func() {
		file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescData)
	})
	return file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDescData
}

var file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_goTypes = []interface{}{
	(*AWSClaims)(nil),             // 0: commonfate.factory.deployid.v1alpha1.AWSClaims
	(*AuthenticateRequest)(nil),   // 1: commonfate.factory.deployid.v1alpha1.AuthenticateRequest
	(*AuthenticateResponse)(nil),  // 2: commonfate.factory.deployid.v1alpha1.AuthenticateResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_depIdxs = []int32{
	3, // 0: commonfate.factory.deployid.v1alpha1.AWSClaims.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: commonfate.factory.deployid.v1alpha1.AuthenticateRequest.aws:type_name -> commonfate.factory.deployid.v1alpha1.AWSClaims
	1, // 2: commonfate.factory.deployid.v1alpha1.MonitoringService.Authenticate:input_type -> commonfate.factory.deployid.v1alpha1.AuthenticateRequest
	2, // 3: commonfate.factory.deployid.v1alpha1.MonitoringService.Authenticate:output_type -> commonfate.factory.deployid.v1alpha1.AuthenticateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_init() }
func file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_init() {
	if File_commonfate_factory_monitoring_v1alpha1_monitoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSClaims); i {
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
		file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
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
		file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateResponse); i {
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
	file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AuthenticateRequest_Aws)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_goTypes,
		DependencyIndexes: file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_depIdxs,
		MessageInfos:      file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_msgTypes,
	}.Build()
	File_commonfate_factory_monitoring_v1alpha1_monitoring_proto = out.File
	file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_rawDesc = nil
	file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_goTypes = nil
	file_commonfate_factory_monitoring_v1alpha1_monitoring_proto_depIdxs = nil
}
