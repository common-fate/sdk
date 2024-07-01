// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: commonfate/control/integration/v1alpha1/aws_rds.proto

package integrationv1alpha1

import (
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

type AWSRDSProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdcProvisionerRoleArn string `protobuf:"bytes,1,opt,name=idc_provisioner_role_arn,json=idcProvisionerRoleArn,proto3" json:"idc_provisioner_role_arn,omitempty"`
	IdcInstanceArn        string `protobuf:"bytes,2,opt,name=idc_instance_arn,json=idcInstanceArn,proto3" json:"idc_instance_arn,omitempty"`
	IdcRegion             string `protobuf:"bytes,3,opt,name=idc_region,json=idcRegion,proto3" json:"idc_region,omitempty"`
	// The id of the AWS account where the proxy is deployed
	AwsAccount string `protobuf:"bytes,4,opt,name=aws_account,json=awsAccount,proto3" json:"aws_account,omitempty"`
	// The AWS region where the proxy is deployed
	AwsRegion string `protobuf:"bytes,5,opt,name=aws_region,json=awsRegion,proto3" json:"aws_region,omitempty"`
	// When deployed to ECS, the name of the cluster where the proxy is deployed
	EcsClusterName string `protobuf:"bytes,6,opt,name=ecs_cluster_name,json=ecsClusterName,proto3" json:"ecs_cluster_name,omitempty"`
	// When deployed to ECS, the name of the proxy task definition
	EcsTaskDefinitionFamily string `protobuf:"bytes,7,opt,name=ecs_task_definition_family,json=ecsTaskDefinitionFamily,proto3" json:"ecs_task_definition_family,omitempty"`
	// When deployed to ECS, the name of the container for the proxy
	EcsContainerName string `protobuf:"bytes,8,opt,name=ecs_container_name,json=ecsContainerName,proto3" json:"ecs_container_name,omitempty"`
}

func (x *AWSRDSProxy) Reset() {
	*x = AWSRDSProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSRDSProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSRDSProxy) ProtoMessage() {}

func (x *AWSRDSProxy) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSRDSProxy.ProtoReflect.Descriptor instead.
func (*AWSRDSProxy) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescGZIP(), []int{0}
}

func (x *AWSRDSProxy) GetIdcProvisionerRoleArn() string {
	if x != nil {
		return x.IdcProvisionerRoleArn
	}
	return ""
}

func (x *AWSRDSProxy) GetIdcInstanceArn() string {
	if x != nil {
		return x.IdcInstanceArn
	}
	return ""
}

func (x *AWSRDSProxy) GetIdcRegion() string {
	if x != nil {
		return x.IdcRegion
	}
	return ""
}

func (x *AWSRDSProxy) GetAwsAccount() string {
	if x != nil {
		return x.AwsAccount
	}
	return ""
}

func (x *AWSRDSProxy) GetAwsRegion() string {
	if x != nil {
		return x.AwsRegion
	}
	return ""
}

func (x *AWSRDSProxy) GetEcsClusterName() string {
	if x != nil {
		return x.EcsClusterName
	}
	return ""
}

func (x *AWSRDSProxy) GetEcsTaskDefinitionFamily() string {
	if x != nil {
		return x.EcsTaskDefinitionFamily
	}
	return ""
}

func (x *AWSRDSProxy) GetEcsContainerName() string {
	if x != nil {
		return x.EcsContainerName
	}
	return ""
}

type AWSRDSDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Engine     string                `protobuf:"bytes,2,opt,name=engine,proto3" json:"engine,omitempty"`
	Region     string                `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Account    string                `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	InstanceId string                `protobuf:"bytes,5,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Database   string                `protobuf:"bytes,6,opt,name=database,proto3" json:"database,omitempty"`
	Users      []*AWSRDSDatabaseUser `protobuf:"bytes,7,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *AWSRDSDatabase) Reset() {
	*x = AWSRDSDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSRDSDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSRDSDatabase) ProtoMessage() {}

func (x *AWSRDSDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSRDSDatabase.ProtoReflect.Descriptor instead.
func (*AWSRDSDatabase) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescGZIP(), []int{1}
}

func (x *AWSRDSDatabase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AWSRDSDatabase) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *AWSRDSDatabase) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AWSRDSDatabase) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AWSRDSDatabase) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *AWSRDSDatabase) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *AWSRDSDatabase) GetUsers() []*AWSRDSDatabaseUser {
	if x != nil {
		return x.Users
	}
	return nil
}

type AWSRDSDatabaseUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *AWSRDSDatabaseUser) Reset() {
	*x = AWSRDSDatabaseUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSRDSDatabaseUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSRDSDatabaseUser) ProtoMessage() {}

func (x *AWSRDSDatabaseUser) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSRDSDatabaseUser.ProtoReflect.Descriptor instead.
func (*AWSRDSDatabaseUser) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescGZIP(), []int{2}
}

func (x *AWSRDSDatabaseUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AWSRDSDatabaseUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_commonfate_control_integration_v1alpha1_aws_rds_proto protoreflect.FileDescriptor

var file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x22, 0xe4, 0x02, 0x0a, 0x0b, 0x41, 0x57, 0x53, 0x52, 0x44, 0x53, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x12, 0x37, 0x0a, 0x18, 0x69, 0x64, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x69, 0x64, 0x63, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x64, 0x63,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x63, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x41, 0x72, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x63, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x77, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x77, 0x73, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x77, 0x73, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x63,
	0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x1a,
	0x65, 0x63, 0x73, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x17, 0x65, 0x63, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x63, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x63, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xfe, 0x01, 0x0a, 0x0e, 0x41, 0x57, 0x53, 0x52,
	0x44, 0x53, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x57, 0x53, 0x52, 0x44, 0x53, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x44, 0x0a, 0x12, 0x41, 0x57, 0x53, 0x52,
	0x44, 0x53, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xd5,
	0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0b,
	0x41, 0x77, 0x73, 0x52, 0x64, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x49, 0xaa,
	0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x33, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a,
	0x3a, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescOnce sync.Once
	file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescData = file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDesc
)

func file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescGZIP() []byte {
	file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescOnce.Do(func() {
		file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescData)
	})
	return file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDescData
}

var file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commonfate_control_integration_v1alpha1_aws_rds_proto_goTypes = []any{
	(*AWSRDSProxy)(nil),        // 0: commonfate.control.integration.v1alpha1.AWSRDSProxy
	(*AWSRDSDatabase)(nil),     // 1: commonfate.control.integration.v1alpha1.AWSRDSDatabase
	(*AWSRDSDatabaseUser)(nil), // 2: commonfate.control.integration.v1alpha1.AWSRDSDatabaseUser
}
var file_commonfate_control_integration_v1alpha1_aws_rds_proto_depIdxs = []int32{
	2, // 0: commonfate.control.integration.v1alpha1.AWSRDSDatabase.users:type_name -> commonfate.control.integration.v1alpha1.AWSRDSDatabaseUser
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_commonfate_control_integration_v1alpha1_aws_rds_proto_init() }
func file_commonfate_control_integration_v1alpha1_aws_rds_proto_init() {
	if File_commonfate_control_integration_v1alpha1_aws_rds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AWSRDSProxy); i {
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
		file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AWSRDSDatabase); i {
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
		file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AWSRDSDatabaseUser); i {
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
			RawDescriptor: file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_control_integration_v1alpha1_aws_rds_proto_goTypes,
		DependencyIndexes: file_commonfate_control_integration_v1alpha1_aws_rds_proto_depIdxs,
		MessageInfos:      file_commonfate_control_integration_v1alpha1_aws_rds_proto_msgTypes,
	}.Build()
	File_commonfate_control_integration_v1alpha1_aws_rds_proto = out.File
	file_commonfate_control_integration_v1alpha1_aws_rds_proto_rawDesc = nil
	file_commonfate_control_integration_v1alpha1_aws_rds_proto_goTypes = nil
	file_commonfate_control_integration_v1alpha1_aws_rds_proto_depIdxs = nil
}
