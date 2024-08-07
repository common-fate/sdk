// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: commonfate/control/integration/v1alpha1/aws_proxy.proto

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

type AWSProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdcProvisionerRoleArn string `protobuf:"bytes,1,opt,name=idc_provisioner_role_arn,json=idcProvisionerRoleArn,proto3" json:"idc_provisioner_role_arn,omitempty"`
	IdcInstanceArn        string `protobuf:"bytes,2,opt,name=idc_instance_arn,json=idcInstanceArn,proto3" json:"idc_instance_arn,omitempty"`
	IdcRegion             string `protobuf:"bytes,3,opt,name=idc_region,json=idcRegion,proto3" json:"idc_region,omitempty"`
	// Types that are assignable to InstanceConfig:
	//
	//	*AWSProxy_AwsEcsProxyInstanceConfig
	InstanceConfig isAWSProxy_InstanceConfig `protobuf_oneof:"instance_config"`
}

func (x *AWSProxy) Reset() {
	*x = AWSProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSProxy) ProtoMessage() {}

func (x *AWSProxy) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSProxy.ProtoReflect.Descriptor instead.
func (*AWSProxy) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *AWSProxy) GetIdcProvisionerRoleArn() string {
	if x != nil {
		return x.IdcProvisionerRoleArn
	}
	return ""
}

func (x *AWSProxy) GetIdcInstanceArn() string {
	if x != nil {
		return x.IdcInstanceArn
	}
	return ""
}

func (x *AWSProxy) GetIdcRegion() string {
	if x != nil {
		return x.IdcRegion
	}
	return ""
}

func (m *AWSProxy) GetInstanceConfig() isAWSProxy_InstanceConfig {
	if m != nil {
		return m.InstanceConfig
	}
	return nil
}

func (x *AWSProxy) GetAwsEcsProxyInstanceConfig() *AWSECSProxyInstanceConfig {
	if x, ok := x.GetInstanceConfig().(*AWSProxy_AwsEcsProxyInstanceConfig); ok {
		return x.AwsEcsProxyInstanceConfig
	}
	return nil
}

type isAWSProxy_InstanceConfig interface {
	isAWSProxy_InstanceConfig()
}

type AWSProxy_AwsEcsProxyInstanceConfig struct {
	AwsEcsProxyInstanceConfig *AWSECSProxyInstanceConfig `protobuf:"bytes,4,opt,name=aws_ecs_proxy_instance_config,json=awsEcsProxyInstanceConfig,proto3,oneof"` // @TODO consider adding support for running the proxy on ec2 here
}

func (*AWSProxy_AwsEcsProxyInstanceConfig) isAWSProxy_InstanceConfig() {}

// This config describes the deployment of the proxy instance, and is used when provisioning access to it
type AWSECSProxyInstanceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the AWS account where the proxy is deployed
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// The AWS region where the proxy is deployed
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// The name of the cluster where the proxy is deployed
	EcsClusterName string `protobuf:"bytes,3,opt,name=ecs_cluster_name,json=ecsClusterName,proto3" json:"ecs_cluster_name,omitempty"`
	// The name of the proxy task definition
	EcsTaskDefinitionFamily string `protobuf:"bytes,4,opt,name=ecs_task_definition_family,json=ecsTaskDefinitionFamily,proto3" json:"ecs_task_definition_family,omitempty"`
	// The name of the container for the proxy
	EcsContainerName string `protobuf:"bytes,5,opt,name=ecs_container_name,json=ecsContainerName,proto3" json:"ecs_container_name,omitempty"`
	// The ARN of the role which can be used to read the task ID and runtime ID of the proxy when provisioning access
	EcsClusterReaderRoleArn string `protobuf:"bytes,6,opt,name=ecs_cluster_reader_role_arn,json=ecsClusterReaderRoleArn,proto3" json:"ecs_cluster_reader_role_arn,omitempty"`
}

func (x *AWSECSProxyInstanceConfig) Reset() {
	*x = AWSECSProxyInstanceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSECSProxyInstanceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSECSProxyInstanceConfig) ProtoMessage() {}

func (x *AWSECSProxyInstanceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSECSProxyInstanceConfig.ProtoReflect.Descriptor instead.
func (*AWSECSProxyInstanceConfig) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *AWSECSProxyInstanceConfig) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *AWSECSProxyInstanceConfig) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AWSECSProxyInstanceConfig) GetEcsClusterName() string {
	if x != nil {
		return x.EcsClusterName
	}
	return ""
}

func (x *AWSECSProxyInstanceConfig) GetEcsTaskDefinitionFamily() string {
	if x != nil {
		return x.EcsTaskDefinitionFamily
	}
	return ""
}

func (x *AWSECSProxyInstanceConfig) GetEcsContainerName() string {
	if x != nil {
		return x.EcsContainerName
	}
	return ""
}

func (x *AWSECSProxyInstanceConfig) GetEcsClusterReaderRoleArn() string {
	if x != nil {
		return x.EcsClusterReaderRoleArn
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
		mi := &file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSRDSDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSRDSDatabase) ProtoMessage() {}

func (x *AWSRDSDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[2]
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
	return file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescGZIP(), []int{2}
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
		mi := &file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSRDSDatabaseUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSRDSDatabaseUser) ProtoMessage() {}

func (x *AWSRDSDatabaseUser) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[3]
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
	return file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescGZIP(), []int{3}
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

var File_commonfate_control_integration_v1alpha1_aws_proxy_proto protoreflect.FileDescriptor

var file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDesc = []byte{
	0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x22, 0xa8, 0x02, 0x0a, 0x08, 0x41, 0x57, 0x53, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12,
	0x37, 0x0a, 0x18, 0x69, 0x64, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x69, 0x64, 0x63, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x64, 0x63, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x63, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x41,
	0x72, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x64, 0x63, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x63, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x86, 0x01, 0x0a, 0x1d, 0x61, 0x77, 0x73, 0x5f, 0x65, 0x63, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x45, 0x43, 0x53, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52,
	0x19, 0x61, 0x77, 0x73, 0x45, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x11, 0x0a, 0x0f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa0, 0x02,
	0x0a, 0x19, 0x41, 0x57, 0x53, 0x45, 0x43, 0x53, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x10, 0x65, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x63, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x65, 0x63, 0x73, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x65, 0x63, 0x73,
	0x54, 0x61, 0x73, 0x6b, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x65, 0x63, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x1b, 0x65, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x72,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x65, 0x63, 0x73, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x72, 0x6e,
	0x22, 0xfe, 0x01, 0x0a, 0x0e, 0x41, 0x57, 0x53, 0x52, 0x44, 0x53, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x52, 0x44, 0x53, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x22, 0x44, 0x0a, 0x12, 0x41, 0x57, 0x53, 0x52, 0x44, 0x53, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xd7, 0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0d, 0x41, 0x77, 0x73, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x49, 0xaa, 0x02, 0x27, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x33, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescOnce sync.Once
	file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescData = file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDesc
)

func file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescGZIP() []byte {
	file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescOnce.Do(func() {
		file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescData)
	})
	return file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDescData
}

var file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_commonfate_control_integration_v1alpha1_aws_proxy_proto_goTypes = []any{
	(*AWSProxy)(nil),                  // 0: commonfate.control.integration.v1alpha1.AWSProxy
	(*AWSECSProxyInstanceConfig)(nil), // 1: commonfate.control.integration.v1alpha1.AWSECSProxyInstanceConfig
	(*AWSRDSDatabase)(nil),            // 2: commonfate.control.integration.v1alpha1.AWSRDSDatabase
	(*AWSRDSDatabaseUser)(nil),        // 3: commonfate.control.integration.v1alpha1.AWSRDSDatabaseUser
}
var file_commonfate_control_integration_v1alpha1_aws_proxy_proto_depIdxs = []int32{
	1, // 0: commonfate.control.integration.v1alpha1.AWSProxy.aws_ecs_proxy_instance_config:type_name -> commonfate.control.integration.v1alpha1.AWSECSProxyInstanceConfig
	3, // 1: commonfate.control.integration.v1alpha1.AWSRDSDatabase.users:type_name -> commonfate.control.integration.v1alpha1.AWSRDSDatabaseUser
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_commonfate_control_integration_v1alpha1_aws_proxy_proto_init() }
func file_commonfate_control_integration_v1alpha1_aws_proxy_proto_init() {
	if File_commonfate_control_integration_v1alpha1_aws_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AWSProxy); i {
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
		file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AWSECSProxyInstanceConfig); i {
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
		file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
	file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes[0].OneofWrappers = []any{
		(*AWSProxy_AwsEcsProxyInstanceConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_control_integration_v1alpha1_aws_proxy_proto_goTypes,
		DependencyIndexes: file_commonfate_control_integration_v1alpha1_aws_proxy_proto_depIdxs,
		MessageInfos:      file_commonfate_control_integration_v1alpha1_aws_proxy_proto_msgTypes,
	}.Build()
	File_commonfate_control_integration_v1alpha1_aws_proxy_proto = out.File
	file_commonfate_control_integration_v1alpha1_aws_proxy_proto_rawDesc = nil
	file_commonfate_control_integration_v1alpha1_aws_proxy_proto_goTypes = nil
	file_commonfate_control_integration_v1alpha1_aws_proxy_proto_depIdxs = nil
}
