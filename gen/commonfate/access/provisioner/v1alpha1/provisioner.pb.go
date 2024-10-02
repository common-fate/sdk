// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: commonfate/access/provisioner/v1alpha1/provisioner.proto

package provisionerv1alpha1

import (
	v1alpha11 "github.com/common-fate/sdk/gen/commonfate/control/integration/v1alpha1"
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

type ProvisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Grant.
	Grant     *v1alpha1.EID    `protobuf:"bytes,1,opt,name=grant,proto3" json:"grant,omitempty"`
	Principal *v1alpha1.Entity `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
	// DomainPrincipal is the principal to grant access to,
	// mapped to the identity domain of the system that we're provisioning in.
	//
	// For example, instead of sending a user's email address, we send the
	// ID of the user in AWS IAM Identity Center.
	//
	// This field may be empty if Common Fate fails to map the principal.
	// A well-designed provisioner should attempt to fall back on data from
	// the core principal entity instead to provision access.
	DomainPrincipal *v1alpha1.Entity `protobuf:"bytes,3,opt,name=domain_principal,json=domainPrincipal,proto3" json:"domain_principal,omitempty"`
	Target          *v1alpha1.Entity `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	Role            *v1alpha1.Entity `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	// The integration config for provisioning access. This will be removed in future,
	// and will be replaced with the 'integration' field.
	IntegrationConfig *v1alpha11.Config `protobuf:"bytes,6,opt,name=integration_config,json=integrationConfig,proto3" json:"integration_config,omitempty"`
	// The ID of the request.
	AccessRequestId *v1alpha1.EID `protobuf:"bytes,7,opt,name=access_request_id,json=accessRequestId,proto3" json:"access_request_id,omitempty"`
	// The integration to use for provisioning access.
	Integration *v1alpha11.Integration `protobuf:"bytes,8,opt,name=integration,proto3" json:"integration,omitempty"`
	// Additional integration configurations. This is used for AWS resource-based access
	// such as our EKS integration, which requires an AWS IAM Identity Center integration.
	AdditionalIntegrations []*v1alpha11.Integration `protobuf:"bytes,9,rep,name=additional_integrations,json=additionalIntegrations,proto3" json:"additional_integrations,omitempty"`
}

func (x *ProvisionRequest) Reset() {
	*x = ProvisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_provisioner_v1alpha1_provisioner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionRequest) ProtoMessage() {}

func (x *ProvisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_provisioner_v1alpha1_provisioner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionRequest.ProtoReflect.Descriptor instead.
func (*ProvisionRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDescGZIP(), []int{0}
}

func (x *ProvisionRequest) GetGrant() *v1alpha1.EID {
	if x != nil {
		return x.Grant
	}
	return nil
}

func (x *ProvisionRequest) GetPrincipal() *v1alpha1.Entity {
	if x != nil {
		return x.Principal
	}
	return nil
}

func (x *ProvisionRequest) GetDomainPrincipal() *v1alpha1.Entity {
	if x != nil {
		return x.DomainPrincipal
	}
	return nil
}

func (x *ProvisionRequest) GetTarget() *v1alpha1.Entity {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *ProvisionRequest) GetRole() *v1alpha1.Entity {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *ProvisionRequest) GetIntegrationConfig() *v1alpha11.Config {
	if x != nil {
		return x.IntegrationConfig
	}
	return nil
}

func (x *ProvisionRequest) GetAccessRequestId() *v1alpha1.EID {
	if x != nil {
		return x.AccessRequestId
	}
	return nil
}

func (x *ProvisionRequest) GetIntegration() *v1alpha11.Integration {
	if x != nil {
		return x.Integration
	}
	return nil
}

func (x *ProvisionRequest) GetAdditionalIntegrations() []*v1alpha11.Integration {
	if x != nil {
		return x.AdditionalIntegrations
	}
	return nil
}

type ProvisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output    *v1alpha1.Entity `protobuf:"bytes,3,opt,name=output,proto3,oneof" json:"output,omitempty"`
	HasAccess *bool            `protobuf:"varint,4,opt,name=has_access,json=hasAccess,proto3,oneof" json:"has_access,omitempty"`
}

func (x *ProvisionResponse) Reset() {
	*x = ProvisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_provisioner_v1alpha1_provisioner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProvisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProvisionResponse) ProtoMessage() {}

func (x *ProvisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_provisioner_v1alpha1_provisioner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProvisionResponse.ProtoReflect.Descriptor instead.
func (*ProvisionResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDescGZIP(), []int{1}
}

func (x *ProvisionResponse) GetOutput() *v1alpha1.Entity {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *ProvisionResponse) GetHasAccess() bool {
	if x != nil && x.HasAccess != nil {
		return *x.HasAccess
	}
	return false
}

var File_commonfate_access_provisioner_v1alpha1_provisioner_proto protoreflect.FileDescriptor

var file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDesc = []byte{
	0x0a, 0x38, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x39, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x05, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49,
	0x44, 0x52, 0x05, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e,
	0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x4d, 0x0a, 0x10, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x5e, 0x0a,
	0x12, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a,
	0x11, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x56, 0x0a, 0x0b, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x6d, 0x0a, 0x17, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x92, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x48, 0x00, 0x52, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x09,
	0x68, 0x61, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x68, 0x61, 0x73, 0x5f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0xd4, 0x02, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x10, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x50, 0xaa, 0x02, 0x26, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xca, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x32, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x29, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDescOnce sync.Once
	file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDescData = file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDesc
)

func file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDescGZIP() []byte {
	file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDescOnce.Do(func() {
		file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDescData)
	})
	return file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDescData
}

var file_commonfate_access_provisioner_v1alpha1_provisioner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commonfate_access_provisioner_v1alpha1_provisioner_proto_goTypes = []any{
	(*ProvisionRequest)(nil),      // 0: commonfate.access.provisioner.v1alpha1.ProvisionRequest
	(*ProvisionResponse)(nil),     // 1: commonfate.access.provisioner.v1alpha1.ProvisionResponse
	(*v1alpha1.EID)(nil),          // 2: commonfate.entity.v1alpha1.EID
	(*v1alpha1.Entity)(nil),       // 3: commonfate.entity.v1alpha1.Entity
	(*v1alpha11.Config)(nil),      // 4: commonfate.control.integration.v1alpha1.Config
	(*v1alpha11.Integration)(nil), // 5: commonfate.control.integration.v1alpha1.Integration
}
var file_commonfate_access_provisioner_v1alpha1_provisioner_proto_depIdxs = []int32{
	2,  // 0: commonfate.access.provisioner.v1alpha1.ProvisionRequest.grant:type_name -> commonfate.entity.v1alpha1.EID
	3,  // 1: commonfate.access.provisioner.v1alpha1.ProvisionRequest.principal:type_name -> commonfate.entity.v1alpha1.Entity
	3,  // 2: commonfate.access.provisioner.v1alpha1.ProvisionRequest.domain_principal:type_name -> commonfate.entity.v1alpha1.Entity
	3,  // 3: commonfate.access.provisioner.v1alpha1.ProvisionRequest.target:type_name -> commonfate.entity.v1alpha1.Entity
	3,  // 4: commonfate.access.provisioner.v1alpha1.ProvisionRequest.role:type_name -> commonfate.entity.v1alpha1.Entity
	4,  // 5: commonfate.access.provisioner.v1alpha1.ProvisionRequest.integration_config:type_name -> commonfate.control.integration.v1alpha1.Config
	2,  // 6: commonfate.access.provisioner.v1alpha1.ProvisionRequest.access_request_id:type_name -> commonfate.entity.v1alpha1.EID
	5,  // 7: commonfate.access.provisioner.v1alpha1.ProvisionRequest.integration:type_name -> commonfate.control.integration.v1alpha1.Integration
	5,  // 8: commonfate.access.provisioner.v1alpha1.ProvisionRequest.additional_integrations:type_name -> commonfate.control.integration.v1alpha1.Integration
	3,  // 9: commonfate.access.provisioner.v1alpha1.ProvisionResponse.output:type_name -> commonfate.entity.v1alpha1.Entity
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_commonfate_access_provisioner_v1alpha1_provisioner_proto_init() }
func file_commonfate_access_provisioner_v1alpha1_provisioner_proto_init() {
	if File_commonfate_access_provisioner_v1alpha1_provisioner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_access_provisioner_v1alpha1_provisioner_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProvisionRequest); i {
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
		file_commonfate_access_provisioner_v1alpha1_provisioner_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProvisionResponse); i {
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
	file_commonfate_access_provisioner_v1alpha1_provisioner_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_access_provisioner_v1alpha1_provisioner_proto_goTypes,
		DependencyIndexes: file_commonfate_access_provisioner_v1alpha1_provisioner_proto_depIdxs,
		MessageInfos:      file_commonfate_access_provisioner_v1alpha1_provisioner_proto_msgTypes,
	}.Build()
	File_commonfate_access_provisioner_v1alpha1_provisioner_proto = out.File
	file_commonfate_access_provisioner_v1alpha1_provisioner_proto_rawDesc = nil
	file_commonfate_access_provisioner_v1alpha1_provisioner_proto_goTypes = nil
	file_commonfate_access_provisioner_v1alpha1_provisioner_proto_depIdxs = nil
}
