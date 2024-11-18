// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: commonfate/control/integration/v1alpha1/healthcheck.proto

package integrationv1alpha1

import (
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/access/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProxyId         string                 `protobuf:"bytes,1,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	EcsTaskMetadata *ECSTaskMetadata       `protobuf:"bytes,2,opt,name=ecs_task_metadata,json=ecsTaskMetadata,proto3" json:"ecs_task_metadata,omitempty"`
	Diagnostics     []*v1alpha1.Diagnostic `protobuf:"bytes,3,rep,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	mi := &file_commonfate_control_integration_v1alpha1_healthcheck_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_healthcheck_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetProxyId() string {
	if x != nil {
		return x.ProxyId
	}
	return ""
}

func (x *PingRequest) GetEcsTaskMetadata() *ECSTaskMetadata {
	if x != nil {
		return x.EcsTaskMetadata
	}
	return nil
}

func (x *PingRequest) GetDiagnostics() []*v1alpha1.Diagnostic {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

type ECSTaskMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskArn          string  `protobuf:"bytes,1,opt,name=task_arn,json=taskArn,proto3" json:"task_arn,omitempty"`
	ClusterArn       string  `protobuf:"bytes,2,opt,name=cluster_arn,json=clusterArn,proto3" json:"cluster_arn,omitempty"`
	AvailabilityZone string  `protobuf:"bytes,3,opt,name=availability_zone,json=availabilityZone,proto3" json:"availability_zone,omitempty"`
	CpuLimit         float32 `protobuf:"fixed32,4,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	MemoryLimit      uint32  `protobuf:"varint,5,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	LaunchType       string  `protobuf:"bytes,6,opt,name=launch_type,json=launchType,proto3" json:"launch_type,omitempty"`
}

func (x *ECSTaskMetadata) Reset() {
	*x = ECSTaskMetadata{}
	mi := &file_commonfate_control_integration_v1alpha1_healthcheck_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ECSTaskMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECSTaskMetadata) ProtoMessage() {}

func (x *ECSTaskMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_healthcheck_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECSTaskMetadata.ProtoReflect.Descriptor instead.
func (*ECSTaskMetadata) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescGZIP(), []int{1}
}

func (x *ECSTaskMetadata) GetTaskArn() string {
	if x != nil {
		return x.TaskArn
	}
	return ""
}

func (x *ECSTaskMetadata) GetClusterArn() string {
	if x != nil {
		return x.ClusterArn
	}
	return ""
}

func (x *ECSTaskMetadata) GetAvailabilityZone() string {
	if x != nil {
		return x.AvailabilityZone
	}
	return ""
}

func (x *ECSTaskMetadata) GetCpuLimit() float32 {
	if x != nil {
		return x.CpuLimit
	}
	return 0
}

func (x *ECSTaskMetadata) GetMemoryLimit() uint32 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *ECSTaskMetadata) GetLaunchType() string {
	if x != nil {
		return x.LaunchType
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An indication from the server as to when the
	// service should next send a ping.
	NextPing *durationpb.Duration `protobuf:"bytes,1,opt,name=next_ping,json=nextPing,proto3" json:"next_ping,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	mi := &file_commonfate_control_integration_v1alpha1_healthcheck_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_control_integration_v1alpha1_healthcheck_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescGZIP(), []int{2}
}

func (x *PingResponse) GetNextPing() *durationpb.Duration {
	if x != nil {
		return x.NextPing
	}
	return nil
}

var File_commonfate_control_integration_v1alpha1_healthcheck_proto protoreflect.FileDescriptor

var file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDesc = []byte{
	0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x49, 0x64, 0x12, 0x64, 0x0a, 0x11,
	0x65, 0x63, 0x73, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x45, 0x43, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x0f, 0x65, 0x63, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x48, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x52,
	0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x22, 0xdb, 0x01, 0x0a,
	0x0f, 0x45, 0x43, 0x53, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x41, 0x72, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x6e, 0x12, 0x2b, 0x0a, 0x11,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x70,
	0x75, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x0c, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x69,
	0x6e, 0x67, 0x32, 0x90, 0x01, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66,
	0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xda, 0x02, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74,
	0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x49, 0xaa, 0x02, 0x27, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2,
	0x02, 0x33, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x3a, 0x3a, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescOnce sync.Once
	file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescData = file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDesc
)

func file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescGZIP() []byte {
	file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescOnce.Do(func() {
		file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescData)
	})
	return file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDescData
}

var file_commonfate_control_integration_v1alpha1_healthcheck_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commonfate_control_integration_v1alpha1_healthcheck_proto_goTypes = []any{
	(*PingRequest)(nil),         // 0: commonfate.control.integration.v1alpha1.PingRequest
	(*ECSTaskMetadata)(nil),     // 1: commonfate.control.integration.v1alpha1.ECSTaskMetadata
	(*PingResponse)(nil),        // 2: commonfate.control.integration.v1alpha1.PingResponse
	(*v1alpha1.Diagnostic)(nil), // 3: commonfate.access.v1alpha1.Diagnostic
	(*durationpb.Duration)(nil), // 4: google.protobuf.Duration
}
var file_commonfate_control_integration_v1alpha1_healthcheck_proto_depIdxs = []int32{
	1, // 0: commonfate.control.integration.v1alpha1.PingRequest.ecs_task_metadata:type_name -> commonfate.control.integration.v1alpha1.ECSTaskMetadata
	3, // 1: commonfate.control.integration.v1alpha1.PingRequest.diagnostics:type_name -> commonfate.access.v1alpha1.Diagnostic
	4, // 2: commonfate.control.integration.v1alpha1.PingResponse.next_ping:type_name -> google.protobuf.Duration
	0, // 3: commonfate.control.integration.v1alpha1.ProxyHealthCheckService.Ping:input_type -> commonfate.control.integration.v1alpha1.PingRequest
	2, // 4: commonfate.control.integration.v1alpha1.ProxyHealthCheckService.Ping:output_type -> commonfate.control.integration.v1alpha1.PingResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_commonfate_control_integration_v1alpha1_healthcheck_proto_init() }
func file_commonfate_control_integration_v1alpha1_healthcheck_proto_init() {
	if File_commonfate_control_integration_v1alpha1_healthcheck_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commonfate_control_integration_v1alpha1_healthcheck_proto_goTypes,
		DependencyIndexes: file_commonfate_control_integration_v1alpha1_healthcheck_proto_depIdxs,
		MessageInfos:      file_commonfate_control_integration_v1alpha1_healthcheck_proto_msgTypes,
	}.Build()
	File_commonfate_control_integration_v1alpha1_healthcheck_proto = out.File
	file_commonfate_control_integration_v1alpha1_healthcheck_proto_rawDesc = nil
	file_commonfate_control_integration_v1alpha1_healthcheck_proto_goTypes = nil
	file_commonfate_control_integration_v1alpha1_healthcheck_proto_depIdxs = nil
}
