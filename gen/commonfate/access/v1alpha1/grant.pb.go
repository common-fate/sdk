// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: commonfate/access/v1alpha1/grant.proto

package accessv1alpha1

import (
	_ "github.com/common-fate/sdk/gen/commonfate/control/attest/v1alpha1"
	_ "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type GrantStatus int32

const (
	GrantStatus_GRANT_STATUS_UNSPECIFIED GrantStatus = 0
	// Awaiting a manual review.
	GrantStatus_GRANT_STATUS_PENDING GrantStatus = 1
	// The grant to the entitlement is currently active.
	GrantStatus_GRANT_STATUS_ACTIVE GrantStatus = 3
	// No longer active. Grants may be closed due to several reasons, such as:
	// - the grant was cancelled by the user
	// - the grant was revoked by an administrator
	// - the grant has expired
	GrantStatus_GRANT_STATUS_CLOSED GrantStatus = 4
)

// Enum value maps for GrantStatus.
var (
	GrantStatus_name = map[int32]string{
		0: "GRANT_STATUS_UNSPECIFIED",
		1: "GRANT_STATUS_PENDING",
		3: "GRANT_STATUS_ACTIVE",
		4: "GRANT_STATUS_CLOSED",
	}
	GrantStatus_value = map[string]int32{
		"GRANT_STATUS_UNSPECIFIED": 0,
		"GRANT_STATUS_PENDING":     1,
		"GRANT_STATUS_ACTIVE":      3,
		"GRANT_STATUS_CLOSED":      4,
	}
)

func (x GrantStatus) Enum() *GrantStatus {
	p := new(GrantStatus)
	*p = x
	return p
}

func (x GrantStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrantStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_access_v1alpha1_grant_proto_enumTypes[0].Descriptor()
}

func (GrantStatus) Type() protoreflect.EnumType {
	return &file_commonfate_access_v1alpha1_grant_proto_enumTypes[0]
}

func (x GrantStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantStatus.Descriptor instead.
func (GrantStatus) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_grant_proto_rawDescGZIP(), []int{0}
}

type ProvisioningStatus int32

const (
	ProvisioningStatus_PROVISIONING_STATUS_UNSPECIFIED ProvisioningStatus = 0
	// Provisioning has not been attempted. This is the value before the Grant is activated.
	ProvisioningStatus_PROVISIONING_STATUS_NOT_ATTEMPTED ProvisioningStatus = 1
	// The Grant is currently being provisioned.
	ProvisioningStatus_PROVISIONING_STATUS_ATTEMPTING ProvisioningStatus = 2
	// Provisioning has completed successfully.
	ProvisioningStatus_PROVISIONING_STATUS_SUCCESSFUL ProvisioningStatus = 3
)

// Enum value maps for ProvisioningStatus.
var (
	ProvisioningStatus_name = map[int32]string{
		0: "PROVISIONING_STATUS_UNSPECIFIED",
		1: "PROVISIONING_STATUS_NOT_ATTEMPTED",
		2: "PROVISIONING_STATUS_ATTEMPTING",
		3: "PROVISIONING_STATUS_SUCCESSFUL",
	}
	ProvisioningStatus_value = map[string]int32{
		"PROVISIONING_STATUS_UNSPECIFIED":   0,
		"PROVISIONING_STATUS_NOT_ATTEMPTED": 1,
		"PROVISIONING_STATUS_ATTEMPTING":    2,
		"PROVISIONING_STATUS_SUCCESSFUL":    3,
	}
)

func (x ProvisioningStatus) Enum() *ProvisioningStatus {
	p := new(ProvisioningStatus)
	*p = x
	return p
}

func (x ProvisioningStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProvisioningStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_access_v1alpha1_grant_proto_enumTypes[1].Descriptor()
}

func (ProvisioningStatus) Type() protoreflect.EnumType {
	return &file_commonfate_access_v1alpha1_grant_proto_enumTypes[1]
}

func (x ProvisioningStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProvisioningStatus.Descriptor instead.
func (ProvisioningStatus) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_grant_proto_rawDescGZIP(), []int{1}
}

type Grant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Grant.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A user-friendly name describing the access, such as "AdministratorAccess to tax-api-prod"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The resource being requested.
	Target *NamedEID `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	// The role being requested, such as "View" or "Admin".
	Role *NamedEID `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	// The user or service account that requested the access.
	//
	// If a service account requested access, the 'name' and 'email' may be empty.
	Principal *User `protobuf:"bytes,5,opt,name=principal,proto3" json:"principal,omitempty"`
	// The status of the Grant.
	Status GrantStatus `protobuf:"varint,6,opt,name=status,proto3,enum=commonfate.access.v1alpha1.GrantStatus" json:"status,omitempty"`
	// For active grants, the time that the access is due to expire.
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// For active grants, the time that the user may attempt to extend the grant.
	// If empty, extension is not permitted.
	TryExtendAfter *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=try_extend_after,json=tryExtendAfter,proto3" json:"try_extend_after,omitempty"`
	// The ID of the Access Request associated with the Grant.
	AccessRequestId string `protobuf:"bytes,10,opt,name=access_request_id,json=accessRequestId,proto3" json:"access_request_id,omitempty"`
	// Approved is true if there are any approving reviews on the Grant.
	Approved bool `protobuf:"varint,11,opt,name=approved,proto3" json:"approved,omitempty"`
	// The time that the grant was closed.
	ClosedAt *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=closed_at,json=closedAt,proto3" json:"closed_at,omitempty"`
	// For grants which were activated, the time that the access was activated.
	ActivatedAt *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=activated_at,json=activatedAt,proto3" json:"activated_at,omitempty"`
	// For grants which were activated, the time that the access was deprovisioned.
	DeprovisionedAt *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=deprovisioned_at,json=deprovisionedAt,proto3" json:"deprovisioned_at,omitempty"`
	// The provisioning status of the Grant. This status tracks whether the entitlements have been successfully provisioned in the integration that Common Fate is provisioning access to, such as AWS or GCP.
	ProvisioningStatus ProvisioningStatus   `protobuf:"varint,15,opt,name=provisioning_status,json=provisioningStatus,proto3,enum=commonfate.access.v1alpha1.ProvisioningStatus" json:"provisioning_status,omitempty"`
	Duration           *durationpb.Duration `protobuf:"bytes,16,opt,name=duration,proto3" json:"duration,omitempty"`
	// The default duration shown for the grant.
	DefaultDuration *durationpb.Duration `protobuf:"bytes,17,opt,name=default_duration,json=defaultDuration,proto3" json:"default_duration,omitempty"`
	// A URL which can be used to access the requested entitlement.
	ExternalUrl string `protobuf:"bytes,18,opt,name=external_url,json=externalUrl,proto3" json:"external_url,omitempty"`
	// Conditions for extending access.
	// If empty, extensions are not permitted.
	Extension *Extension `protobuf:"bytes,19,opt,name=extension,proto3" json:"extension,omitempty"`
	// Instructions in Markdown format for accessing the entitlement
	// using a CLI.
	CliAccessInstructions string `protobuf:"bytes,20,opt,name=cli_access_instructions,json=cliAccessInstructions,proto3" json:"cli_access_instructions,omitempty"`
}

func (x *Grant) Reset() {
	*x = Grant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_grant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grant) ProtoMessage() {}

func (x *Grant) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_grant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grant.ProtoReflect.Descriptor instead.
func (*Grant) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_grant_proto_rawDescGZIP(), []int{0}
}

func (x *Grant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Grant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Grant) GetTarget() *NamedEID {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Grant) GetRole() *NamedEID {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *Grant) GetPrincipal() *User {
	if x != nil {
		return x.Principal
	}
	return nil
}

func (x *Grant) GetStatus() GrantStatus {
	if x != nil {
		return x.Status
	}
	return GrantStatus_GRANT_STATUS_UNSPECIFIED
}

func (x *Grant) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *Grant) GetTryExtendAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.TryExtendAfter
	}
	return nil
}

func (x *Grant) GetAccessRequestId() string {
	if x != nil {
		return x.AccessRequestId
	}
	return ""
}

func (x *Grant) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

func (x *Grant) GetClosedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ClosedAt
	}
	return nil
}

func (x *Grant) GetActivatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ActivatedAt
	}
	return nil
}

func (x *Grant) GetDeprovisionedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeprovisionedAt
	}
	return nil
}

func (x *Grant) GetProvisioningStatus() ProvisioningStatus {
	if x != nil {
		return x.ProvisioningStatus
	}
	return ProvisioningStatus_PROVISIONING_STATUS_UNSPECIFIED
}

func (x *Grant) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Grant) GetDefaultDuration() *durationpb.Duration {
	if x != nil {
		return x.DefaultDuration
	}
	return nil
}

func (x *Grant) GetExternalUrl() string {
	if x != nil {
		return x.ExternalUrl
	}
	return ""
}

func (x *Grant) GetExtension() *Extension {
	if x != nil {
		return x.Extension
	}
	return nil
}

func (x *Grant) GetCliAccessInstructions() string {
	if x != nil {
		return x.CliAccessInstructions
	}
	return ""
}

type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the number of allowed extensions remaining
	Remaining int32 `protobuf:"varint,1,opt,name=remaining,proto3" json:"remaining,omitempty"`
	// the duration to extend by on each extension
	ExtensionDurationSeconds *durationpb.Duration `protobuf:"bytes,2,opt,name=extension_duration_seconds,json=extensionDurationSeconds,proto3" json:"extension_duration_seconds,omitempty"`
}

func (x *Extension) Reset() {
	*x = Extension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_access_v1alpha1_grant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_access_v1alpha1_grant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_commonfate_access_v1alpha1_grant_proto_rawDescGZIP(), []int{1}
}

func (x *Extension) GetRemaining() int32 {
	if x != nil {
		return x.Remaining
	}
	return 0
}

func (x *Extension) GetExtensionDurationSeconds() *durationpb.Duration {
	if x != nil {
		return x.ExtensionDurationSeconds
	}
	return nil
}

var File_commonfate_access_v1alpha1_grant_proto protoreflect.FileDescriptor

var file_commonfate_access_v1alpha1_grant_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x34, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x65, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x08, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x64, 0x45, 0x49, 0x44, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x38, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x64, 0x45, 0x49, 0x44, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x74, 0x72, 0x79, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x74, 0x72,
	0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a,
	0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x45, 0x0a, 0x10,
	0x64, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0f, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x5f, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x12, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x10, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x55, 0x72, 0x6c, 0x12, 0x43, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6c, 0x69,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x63, 0x6c, 0x69, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x82, 0x01, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x57, 0x0a,
	0x1a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x18, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2a, 0x77, 0x0a, 0x0b, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x04, 0x2a,
	0xa8, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53,
	0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x54, 0x54, 0x45, 0x4d, 0x50, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x54, 0x54, 0x45, 0x4d, 0x50,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53,
	0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55, 0x4c, 0x10, 0x03, 0x42, 0x80, 0x02, 0x0a, 0x1e, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0a, 0x47,
	0x72, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x1a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x26, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x5c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x1c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_access_v1alpha1_grant_proto_rawDescOnce sync.Once
	file_commonfate_access_v1alpha1_grant_proto_rawDescData = file_commonfate_access_v1alpha1_grant_proto_rawDesc
)

func file_commonfate_access_v1alpha1_grant_proto_rawDescGZIP() []byte {
	file_commonfate_access_v1alpha1_grant_proto_rawDescOnce.Do(func() {
		file_commonfate_access_v1alpha1_grant_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_access_v1alpha1_grant_proto_rawDescData)
	})
	return file_commonfate_access_v1alpha1_grant_proto_rawDescData
}

var file_commonfate_access_v1alpha1_grant_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_commonfate_access_v1alpha1_grant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commonfate_access_v1alpha1_grant_proto_goTypes = []any{
	(GrantStatus)(0),              // 0: commonfate.access.v1alpha1.GrantStatus
	(ProvisioningStatus)(0),       // 1: commonfate.access.v1alpha1.ProvisioningStatus
	(*Grant)(nil),                 // 2: commonfate.access.v1alpha1.Grant
	(*Extension)(nil),             // 3: commonfate.access.v1alpha1.Extension
	(*NamedEID)(nil),              // 4: commonfate.access.v1alpha1.NamedEID
	(*User)(nil),                  // 5: commonfate.access.v1alpha1.User
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
}
var file_commonfate_access_v1alpha1_grant_proto_depIdxs = []int32{
	4,  // 0: commonfate.access.v1alpha1.Grant.target:type_name -> commonfate.access.v1alpha1.NamedEID
	4,  // 1: commonfate.access.v1alpha1.Grant.role:type_name -> commonfate.access.v1alpha1.NamedEID
	5,  // 2: commonfate.access.v1alpha1.Grant.principal:type_name -> commonfate.access.v1alpha1.User
	0,  // 3: commonfate.access.v1alpha1.Grant.status:type_name -> commonfate.access.v1alpha1.GrantStatus
	6,  // 4: commonfate.access.v1alpha1.Grant.expires_at:type_name -> google.protobuf.Timestamp
	6,  // 5: commonfate.access.v1alpha1.Grant.try_extend_after:type_name -> google.protobuf.Timestamp
	6,  // 6: commonfate.access.v1alpha1.Grant.closed_at:type_name -> google.protobuf.Timestamp
	6,  // 7: commonfate.access.v1alpha1.Grant.activated_at:type_name -> google.protobuf.Timestamp
	6,  // 8: commonfate.access.v1alpha1.Grant.deprovisioned_at:type_name -> google.protobuf.Timestamp
	1,  // 9: commonfate.access.v1alpha1.Grant.provisioning_status:type_name -> commonfate.access.v1alpha1.ProvisioningStatus
	7,  // 10: commonfate.access.v1alpha1.Grant.duration:type_name -> google.protobuf.Duration
	7,  // 11: commonfate.access.v1alpha1.Grant.default_duration:type_name -> google.protobuf.Duration
	3,  // 12: commonfate.access.v1alpha1.Grant.extension:type_name -> commonfate.access.v1alpha1.Extension
	7,  // 13: commonfate.access.v1alpha1.Extension.extension_duration_seconds:type_name -> google.protobuf.Duration
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_commonfate_access_v1alpha1_grant_proto_init() }
func file_commonfate_access_v1alpha1_grant_proto_init() {
	if File_commonfate_access_v1alpha1_grant_proto != nil {
		return
	}
	file_commonfate_access_v1alpha1_user_proto_init()
	file_commonfate_access_v1alpha1_named_eid_proto_init()
	file_commonfate_access_v1alpha1_audit_logs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commonfate_access_v1alpha1_grant_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Grant); i {
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
		file_commonfate_access_v1alpha1_grant_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Extension); i {
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
			RawDescriptor: file_commonfate_access_v1alpha1_grant_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_access_v1alpha1_grant_proto_goTypes,
		DependencyIndexes: file_commonfate_access_v1alpha1_grant_proto_depIdxs,
		EnumInfos:         file_commonfate_access_v1alpha1_grant_proto_enumTypes,
		MessageInfos:      file_commonfate_access_v1alpha1_grant_proto_msgTypes,
	}.Build()
	File_commonfate_access_v1alpha1_grant_proto = out.File
	file_commonfate_access_v1alpha1_grant_proto_rawDesc = nil
	file_commonfate_access_v1alpha1_grant_proto_goTypes = nil
	file_commonfate_access_v1alpha1_grant_proto_depIdxs = nil
}
