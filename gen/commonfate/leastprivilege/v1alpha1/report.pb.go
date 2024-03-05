// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: commonfate/leastprivilege/v1alpha1/report.proto

package leastprivilegev1alpha1

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

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsageSummaries  []*UsageSummary        `protobuf:"bytes,1,rep,name=usage_summaries,json=usageSummaries,proto3" json:"usage_summaries,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Timestamp       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	UsageCutoffTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=usage_cutoff_time,json=usageCutoffTime,proto3" json:"usage_cutoff_time,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_leastprivilege_v1alpha1_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_report_proto_rawDescGZIP(), []int{0}
}

func (x *Report) GetUsageSummaries() []*UsageSummary {
	if x != nil {
		return x.UsageSummaries
	}
	return nil
}

func (x *Report) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Report) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Report) GetUsageCutoffTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UsageCutoffTime
	}
	return nil
}

type UsageSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label              string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	InUseCount         int32  `protobuf:"varint,3,opt,name=in_use_count,json=inUseCount,proto3" json:"in_use_count,omitempty"`
	UnusedCount        int32  `protobuf:"varint,4,opt,name=unused_count,json=unusedCount,proto3" json:"unused_count,omitempty"`
	IndeterminateCount int32  `protobuf:"varint,5,opt,name=indeterminate_count,json=indeterminateCount,proto3" json:"indeterminate_count,omitempty"`
}

func (x *UsageSummary) Reset() {
	*x = UsageSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commonfate_leastprivilege_v1alpha1_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageSummary) ProtoMessage() {}

func (x *UsageSummary) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_leastprivilege_v1alpha1_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageSummary.ProtoReflect.Descriptor instead.
func (*UsageSummary) Descriptor() ([]byte, []int) {
	return file_commonfate_leastprivilege_v1alpha1_report_proto_rawDescGZIP(), []int{1}
}

func (x *UsageSummary) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *UsageSummary) GetInUseCount() int32 {
	if x != nil {
		return x.InUseCount
	}
	return 0
}

func (x *UsageSummary) GetUnusedCount() int32 {
	if x != nil {
		return x.UnusedCount
	}
	return 0
}

func (x *UsageSummary) GetIndeterminateCount() int32 {
	if x != nil {
		return x.IndeterminateCount
	}
	return 0
}

var File_commonfate_leastprivilege_v1alpha1_report_proto protoreflect.FileDescriptor

var file_commonfate_leastprivilege_v1alpha1_report_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65,
	0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x59, 0x0a, 0x0f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69,
	0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0e, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x46, 0x0a, 0x11, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x75, 0x74, 0x6f, 0x66,
	0x66, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x75, 0x74, 0x6f, 0x66, 0x66, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x20, 0x0a, 0x0c, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x75, 0x6e, 0x75, 0x73, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x12, 0x69, 0x6e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xb9, 0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x65, 0x61, 0x73,
	0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x3b, 0x6c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x4c, 0x58,
	0xaa, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x4c, 0x65,
	0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x22, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x5c, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x5c, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x70, 0x72, 0x69,
	0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x70,
	0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_leastprivilege_v1alpha1_report_proto_rawDescOnce sync.Once
	file_commonfate_leastprivilege_v1alpha1_report_proto_rawDescData = file_commonfate_leastprivilege_v1alpha1_report_proto_rawDesc
)

func file_commonfate_leastprivilege_v1alpha1_report_proto_rawDescGZIP() []byte {
	file_commonfate_leastprivilege_v1alpha1_report_proto_rawDescOnce.Do(func() {
		file_commonfate_leastprivilege_v1alpha1_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_leastprivilege_v1alpha1_report_proto_rawDescData)
	})
	return file_commonfate_leastprivilege_v1alpha1_report_proto_rawDescData
}

var file_commonfate_leastprivilege_v1alpha1_report_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_commonfate_leastprivilege_v1alpha1_report_proto_goTypes = []interface{}{
	(*Report)(nil),                // 0: commonfate.leastprivilege.v1alpha1.Report
	(*UsageSummary)(nil),          // 1: commonfate.leastprivilege.v1alpha1.UsageSummary
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_commonfate_leastprivilege_v1alpha1_report_proto_depIdxs = []int32{
	1, // 0: commonfate.leastprivilege.v1alpha1.Report.usage_summaries:type_name -> commonfate.leastprivilege.v1alpha1.UsageSummary
	2, // 1: commonfate.leastprivilege.v1alpha1.Report.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: commonfate.leastprivilege.v1alpha1.Report.timestamp:type_name -> google.protobuf.Timestamp
	2, // 3: commonfate.leastprivilege.v1alpha1.Report.usage_cutoff_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_commonfate_leastprivilege_v1alpha1_report_proto_init() }
func file_commonfate_leastprivilege_v1alpha1_report_proto_init() {
	if File_commonfate_leastprivilege_v1alpha1_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commonfate_leastprivilege_v1alpha1_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_commonfate_leastprivilege_v1alpha1_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsageSummary); i {
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
			RawDescriptor: file_commonfate_leastprivilege_v1alpha1_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_leastprivilege_v1alpha1_report_proto_goTypes,
		DependencyIndexes: file_commonfate_leastprivilege_v1alpha1_report_proto_depIdxs,
		MessageInfos:      file_commonfate_leastprivilege_v1alpha1_report_proto_msgTypes,
	}.Build()
	File_commonfate_leastprivilege_v1alpha1_report_proto = out.File
	file_commonfate_leastprivilege_v1alpha1_report_proto_rawDesc = nil
	file_commonfate_leastprivilege_v1alpha1_report_proto_goTypes = nil
	file_commonfate_leastprivilege_v1alpha1_report_proto_depIdxs = nil
}