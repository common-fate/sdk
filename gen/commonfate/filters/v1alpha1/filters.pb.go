// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: commonfate/filters/v1alpha1/filters.proto

package filtersv1alpha1

import (
	v1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
	v1alpha11 "github.com/common-fate/sdk/gen/commonfate/entity/v1alpha1"
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

type BoolComparison int32

const (
	BoolComparison_BOOL_COMPARISON_UNSPECIFIED BoolComparison = 0
	// Equal to the provided value.
	BoolComparison_BOOL_COMPARISON_EQUAL BoolComparison = 1
	// Not equal to the provided value.
	BoolComparison_BOOL_COMPARISON_NOT_EQUAL BoolComparison = 2
)

// Enum value maps for BoolComparison.
var (
	BoolComparison_name = map[int32]string{
		0: "BOOL_COMPARISON_UNSPECIFIED",
		1: "BOOL_COMPARISON_EQUAL",
		2: "BOOL_COMPARISON_NOT_EQUAL",
	}
	BoolComparison_value = map[string]int32{
		"BOOL_COMPARISON_UNSPECIFIED": 0,
		"BOOL_COMPARISON_EQUAL":       1,
		"BOOL_COMPARISON_NOT_EQUAL":   2,
	}
)

func (x BoolComparison) Enum() *BoolComparison {
	p := new(BoolComparison)
	*p = x
	return p
}

func (x BoolComparison) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BoolComparison) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_filters_v1alpha1_filters_proto_enumTypes[0].Descriptor()
}

func (BoolComparison) Type() protoreflect.EnumType {
	return &file_commonfate_filters_v1alpha1_filters_proto_enumTypes[0]
}

func (x BoolComparison) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BoolComparison.Descriptor instead.
func (BoolComparison) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{0}
}

type EqualityComparison int32

const (
	EqualityComparison_EQUALITY_COMPARISON_UNSPECIFIED EqualityComparison = 0
	// Equal to the provided value.
	EqualityComparison_EQUALITY_COMPARISON_EQUAL EqualityComparison = 1
	// Not equal to the provided value.
	EqualityComparison_EQUALITY_COMPARISON_NOT_EQUAL EqualityComparison = 2
	// Less than the provided value.
	EqualityComparison_EQUALITY_COMPARISON_LESS_THAN EqualityComparison = 3
	// Greater than the provided value.
	EqualityComparison_EQUALITY_COMPARISON_GREATER_THAN EqualityComparison = 4
	// Less than or equal to the provided value.
	EqualityComparison_EQUALITY_COMPARISON_LESS_THAN_EQUAL EqualityComparison = 5
	// Greater than or equal to the provided value.
	EqualityComparison_EQUALITY_COMPARISON_GREATER_THAN_EQUAL EqualityComparison = 6
)

// Enum value maps for EqualityComparison.
var (
	EqualityComparison_name = map[int32]string{
		0: "EQUALITY_COMPARISON_UNSPECIFIED",
		1: "EQUALITY_COMPARISON_EQUAL",
		2: "EQUALITY_COMPARISON_NOT_EQUAL",
		3: "EQUALITY_COMPARISON_LESS_THAN",
		4: "EQUALITY_COMPARISON_GREATER_THAN",
		5: "EQUALITY_COMPARISON_LESS_THAN_EQUAL",
		6: "EQUALITY_COMPARISON_GREATER_THAN_EQUAL",
	}
	EqualityComparison_value = map[string]int32{
		"EQUALITY_COMPARISON_UNSPECIFIED":        0,
		"EQUALITY_COMPARISON_EQUAL":              1,
		"EQUALITY_COMPARISON_NOT_EQUAL":          2,
		"EQUALITY_COMPARISON_LESS_THAN":          3,
		"EQUALITY_COMPARISON_GREATER_THAN":       4,
		"EQUALITY_COMPARISON_LESS_THAN_EQUAL":    5,
		"EQUALITY_COMPARISON_GREATER_THAN_EQUAL": 6,
	}
)

func (x EqualityComparison) Enum() *EqualityComparison {
	p := new(EqualityComparison)
	*p = x
	return p
}

func (x EqualityComparison) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EqualityComparison) Descriptor() protoreflect.EnumDescriptor {
	return file_commonfate_filters_v1alpha1_filters_proto_enumTypes[1].Descriptor()
}

func (EqualityComparison) Type() protoreflect.EnumType {
	return &file_commonfate_filters_v1alpha1_filters_proto_enumTypes[1]
}

func (x EqualityComparison) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EqualityComparison.Descriptor instead.
func (EqualityComparison) EnumDescriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{1}
}

// Filters items with a particular authorization decision
type DecisionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision v1alpha1.Decision `protobuf:"varint,1,opt,name=decision,proto3,enum=commonfate.authz.v1alpha1.Decision" json:"decision,omitempty"`
}

func (x *DecisionFilter) Reset() {
	*x = DecisionFilter{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecisionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionFilter) ProtoMessage() {}

func (x *DecisionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionFilter.ProtoReflect.Descriptor instead.
func (*DecisionFilter) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{0}
}

func (x *DecisionFilter) GetDecision() v1alpha1.Decision {
	if x != nil {
		return x.Decision
	}
	return v1alpha1.Decision(0)
}

// Filters by a particular key/value tag associated with items.
// Will filter for either items which match the tag, or do not match the tag,
// based on the 'comparison' value.
type TagFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key        string         `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value      string         `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Comparison BoolComparison `protobuf:"varint,3,opt,name=comparison,proto3,enum=commonfate.filters.v1alpha1.BoolComparison" json:"comparison,omitempty"`
}

func (x *TagFilter) Reset() {
	*x = TagFilter{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TagFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagFilter) ProtoMessage() {}

func (x *TagFilter) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagFilter.ProtoReflect.Descriptor instead.
func (*TagFilter) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{1}
}

func (x *TagFilter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TagFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TagFilter) GetComparison() BoolComparison {
	if x != nil {
		return x.Comparison
	}
	return BoolComparison_BOOL_COMPARISON_UNSPECIFIED
}

// Filters where the principal/action/resource matches particular entities.
type EntityFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IDs of entities to filter for. Will filter for a principal/action/resource matching any of these entities.
	Ids        []*v1alpha11.EID `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Comparison BoolComparison   `protobuf:"varint,2,opt,name=comparison,proto3,enum=commonfate.filters.v1alpha1.BoolComparison" json:"comparison,omitempty"`
}

func (x *EntityFilter) Reset() {
	*x = EntityFilter{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityFilter) ProtoMessage() {}

func (x *EntityFilter) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityFilter.ProtoReflect.Descriptor instead.
func (*EntityFilter) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{2}
}

func (x *EntityFilter) GetIds() []*v1alpha11.EID {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *EntityFilter) GetComparison() BoolComparison {
	if x != nil {
		return x.Comparison
	}
	return BoolComparison_BOOL_COMPARISON_UNSPECIFIED
}

// Filters where the principal/action/resource matches particular entity types.
type EntityTypeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types of entities to filter for. Will filter for a principal/action/resource matching any of these entities.
	Types      []string       `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
	Comparison BoolComparison `protobuf:"varint,2,opt,name=comparison,proto3,enum=commonfate.filters.v1alpha1.BoolComparison" json:"comparison,omitempty"`
}

func (x *EntityTypeFilter) Reset() {
	*x = EntityTypeFilter{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityTypeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityTypeFilter) ProtoMessage() {}

func (x *EntityTypeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityTypeFilter.ProtoReflect.Descriptor instead.
func (*EntityTypeFilter) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{3}
}

func (x *EntityTypeFilter) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *EntityTypeFilter) GetComparison() BoolComparison {
	if x != nil {
		return x.Comparison
	}
	return BoolComparison_BOOL_COMPARISON_UNSPECIFIED
}

// Filters based on an integer equality.
type IntegerFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      int64              `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Comparison EqualityComparison `protobuf:"varint,2,opt,name=comparison,proto3,enum=commonfate.filters.v1alpha1.EqualityComparison" json:"comparison,omitempty"`
}

func (x *IntegerFilter) Reset() {
	*x = IntegerFilter{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntegerFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegerFilter) ProtoMessage() {}

func (x *IntegerFilter) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegerFilter.ProtoReflect.Descriptor instead.
func (*IntegerFilter) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{4}
}

func (x *IntegerFilter) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *IntegerFilter) GetComparison() EqualityComparison {
	if x != nil {
		return x.Comparison
	}
	return EqualityComparison_EQUALITY_COMPARISON_UNSPECIFIED
}

// Filters based on a condition being true or false.
type BoolFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comparison BoolComparison `protobuf:"varint,2,opt,name=comparison,proto3,enum=commonfate.filters.v1alpha1.BoolComparison" json:"comparison,omitempty"`
}

func (x *BoolFilter) Reset() {
	*x = BoolFilter{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BoolFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolFilter) ProtoMessage() {}

func (x *BoolFilter) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolFilter.ProtoReflect.Descriptor instead.
func (*BoolFilter) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{5}
}

func (x *BoolFilter) GetComparison() BoolComparison {
	if x != nil {
		return x.Comparison
	}
	return BoolComparison_BOOL_COMPARISON_UNSPECIFIED
}

type RelativeTimeBound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration *durationpb.Duration `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *RelativeTimeBound) Reset() {
	*x = RelativeTimeBound{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelativeTimeBound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelativeTimeBound) ProtoMessage() {}

func (x *RelativeTimeBound) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelativeTimeBound.ProtoReflect.Descriptor instead.
func (*RelativeTimeBound) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{6}
}

func (x *RelativeTimeBound) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type AbsoluteTimeBound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *AbsoluteTimeBound) Reset() {
	*x = AbsoluteTimeBound{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AbsoluteTimeBound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbsoluteTimeBound) ProtoMessage() {}

func (x *AbsoluteTimeBound) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbsoluteTimeBound.ProtoReflect.Descriptor instead.
func (*AbsoluteTimeBound) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{7}
}

func (x *AbsoluteTimeBound) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type TimeBound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Bound:
	//
	//	*TimeBound_Relative
	//	*TimeBound_Absolute
	Bound isTimeBound_Bound `protobuf_oneof:"bound"`
}

func (x *TimeBound) Reset() {
	*x = TimeBound{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeBound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeBound) ProtoMessage() {}

func (x *TimeBound) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeBound.ProtoReflect.Descriptor instead.
func (*TimeBound) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{8}
}

func (m *TimeBound) GetBound() isTimeBound_Bound {
	if m != nil {
		return m.Bound
	}
	return nil
}

func (x *TimeBound) GetRelative() *RelativeTimeBound {
	if x, ok := x.GetBound().(*TimeBound_Relative); ok {
		return x.Relative
	}
	return nil
}

func (x *TimeBound) GetAbsolute() *AbsoluteTimeBound {
	if x, ok := x.GetBound().(*TimeBound_Absolute); ok {
		return x.Absolute
	}
	return nil
}

type isTimeBound_Bound interface {
	isTimeBound_Bound()
}

type TimeBound_Relative struct {
	Relative *RelativeTimeBound `protobuf:"bytes,1,opt,name=relative,proto3,oneof"`
}

type TimeBound_Absolute struct {
	Absolute *AbsoluteTimeBound `protobuf:"bytes,2,opt,name=absolute,proto3,oneof"`
}

func (*TimeBound_Relative) isTimeBound_Bound() {}

func (*TimeBound_Absolute) isTimeBound_Bound() {}

// Filters based on a time range, where start or end is not provided a relative time is used
// e.g everything before an end time or everything after a start time
type TimeRangeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optionaly provide a starting bound for the date range
	Start *TimeBound `protobuf:"bytes,1,opt,name=start,proto3,oneof" json:"start,omitempty"`
	// Optionaly provide an ending bound for the date range
	End *TimeBound `protobuf:"bytes,2,opt,name=end,proto3,oneof" json:"end,omitempty"`
}

func (x *TimeRangeFilter) Reset() {
	*x = TimeRangeFilter{}
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeRangeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRangeFilter) ProtoMessage() {}

func (x *TimeRangeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_commonfate_filters_v1alpha1_filters_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRangeFilter.ProtoReflect.Descriptor instead.
func (*TimeRangeFilter) Descriptor() ([]byte, []int) {
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP(), []int{9}
}

func (x *TimeRangeFilter) GetStart() *TimeBound {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *TimeRangeFilter) GetEnd() *TimeBound {
	if x != nil {
		return x.End
	}
	return nil
}

var File_commonfate_filters_v1alpha1_filters_proto protoreflect.FileDescriptor

var file_commonfate_filters_v1alpha1_filters_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x2a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0e, 0x44,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a,
	0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x80,
	0x01, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f,
	0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x31, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x49, 0x44,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x4b, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69,
	0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73,
	0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x52, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x0d, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x4f, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f,
	0x6e, 0x22, 0x59, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x4b, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e,
	0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x11,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x43, 0x0a, 0x11, 0x41, 0x62, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xb0, 0x01,
	0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x4c, 0x0a, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x00, 0x52,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x61, 0x62, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x08, 0x61,
	0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74,
	0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x01, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x65, 0x6e, 0x64, 0x2a, 0x6b, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6c,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x1b, 0x42, 0x4f,
	0x4f, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x42,
	0x4f, 0x4f, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x45,
	0x51, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x4f, 0x4f, 0x4c, 0x5f, 0x43,
	0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51,
	0x55, 0x41, 0x4c, 0x10, 0x02, 0x2a, 0x99, 0x02, 0x0a, 0x12, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x1f,
	0x45, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49,
	0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x01,
	0x12, 0x21, 0x0a, 0x1d, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41,
	0x4c, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f,
	0x54, 0x48, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x47, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x27, 0x0a, 0x23,
	0x45, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49,
	0x53, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45, 0x51,
	0x55, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x2a, 0x0a, 0x26, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10,
	0x06, 0x42, 0x89, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x46, 0x58, 0xaa, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x66, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61,
	0x74, 0x65, 0x5c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x27, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65,
	0x5c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x66, 0x61, 0x74, 0x65, 0x3a, 0x3a, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commonfate_filters_v1alpha1_filters_proto_rawDescOnce sync.Once
	file_commonfate_filters_v1alpha1_filters_proto_rawDescData = file_commonfate_filters_v1alpha1_filters_proto_rawDesc
)

func file_commonfate_filters_v1alpha1_filters_proto_rawDescGZIP() []byte {
	file_commonfate_filters_v1alpha1_filters_proto_rawDescOnce.Do(func() {
		file_commonfate_filters_v1alpha1_filters_proto_rawDescData = protoimpl.X.CompressGZIP(file_commonfate_filters_v1alpha1_filters_proto_rawDescData)
	})
	return file_commonfate_filters_v1alpha1_filters_proto_rawDescData
}

var file_commonfate_filters_v1alpha1_filters_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_commonfate_filters_v1alpha1_filters_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_commonfate_filters_v1alpha1_filters_proto_goTypes = []any{
	(BoolComparison)(0),           // 0: commonfate.filters.v1alpha1.BoolComparison
	(EqualityComparison)(0),       // 1: commonfate.filters.v1alpha1.EqualityComparison
	(*DecisionFilter)(nil),        // 2: commonfate.filters.v1alpha1.DecisionFilter
	(*TagFilter)(nil),             // 3: commonfate.filters.v1alpha1.TagFilter
	(*EntityFilter)(nil),          // 4: commonfate.filters.v1alpha1.EntityFilter
	(*EntityTypeFilter)(nil),      // 5: commonfate.filters.v1alpha1.EntityTypeFilter
	(*IntegerFilter)(nil),         // 6: commonfate.filters.v1alpha1.IntegerFilter
	(*BoolFilter)(nil),            // 7: commonfate.filters.v1alpha1.BoolFilter
	(*RelativeTimeBound)(nil),     // 8: commonfate.filters.v1alpha1.RelativeTimeBound
	(*AbsoluteTimeBound)(nil),     // 9: commonfate.filters.v1alpha1.AbsoluteTimeBound
	(*TimeBound)(nil),             // 10: commonfate.filters.v1alpha1.TimeBound
	(*TimeRangeFilter)(nil),       // 11: commonfate.filters.v1alpha1.TimeRangeFilter
	(v1alpha1.Decision)(0),        // 12: commonfate.authz.v1alpha1.Decision
	(*v1alpha11.EID)(nil),         // 13: commonfate.entity.v1alpha1.EID
	(*durationpb.Duration)(nil),   // 14: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 15: google.protobuf.Timestamp
}
var file_commonfate_filters_v1alpha1_filters_proto_depIdxs = []int32{
	12, // 0: commonfate.filters.v1alpha1.DecisionFilter.decision:type_name -> commonfate.authz.v1alpha1.Decision
	0,  // 1: commonfate.filters.v1alpha1.TagFilter.comparison:type_name -> commonfate.filters.v1alpha1.BoolComparison
	13, // 2: commonfate.filters.v1alpha1.EntityFilter.ids:type_name -> commonfate.entity.v1alpha1.EID
	0,  // 3: commonfate.filters.v1alpha1.EntityFilter.comparison:type_name -> commonfate.filters.v1alpha1.BoolComparison
	0,  // 4: commonfate.filters.v1alpha1.EntityTypeFilter.comparison:type_name -> commonfate.filters.v1alpha1.BoolComparison
	1,  // 5: commonfate.filters.v1alpha1.IntegerFilter.comparison:type_name -> commonfate.filters.v1alpha1.EqualityComparison
	0,  // 6: commonfate.filters.v1alpha1.BoolFilter.comparison:type_name -> commonfate.filters.v1alpha1.BoolComparison
	14, // 7: commonfate.filters.v1alpha1.RelativeTimeBound.duration:type_name -> google.protobuf.Duration
	15, // 8: commonfate.filters.v1alpha1.AbsoluteTimeBound.time:type_name -> google.protobuf.Timestamp
	8,  // 9: commonfate.filters.v1alpha1.TimeBound.relative:type_name -> commonfate.filters.v1alpha1.RelativeTimeBound
	9,  // 10: commonfate.filters.v1alpha1.TimeBound.absolute:type_name -> commonfate.filters.v1alpha1.AbsoluteTimeBound
	10, // 11: commonfate.filters.v1alpha1.TimeRangeFilter.start:type_name -> commonfate.filters.v1alpha1.TimeBound
	10, // 12: commonfate.filters.v1alpha1.TimeRangeFilter.end:type_name -> commonfate.filters.v1alpha1.TimeBound
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_commonfate_filters_v1alpha1_filters_proto_init() }
func file_commonfate_filters_v1alpha1_filters_proto_init() {
	if File_commonfate_filters_v1alpha1_filters_proto != nil {
		return
	}
	file_commonfate_filters_v1alpha1_filters_proto_msgTypes[8].OneofWrappers = []any{
		(*TimeBound_Relative)(nil),
		(*TimeBound_Absolute)(nil),
	}
	file_commonfate_filters_v1alpha1_filters_proto_msgTypes[9].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_commonfate_filters_v1alpha1_filters_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commonfate_filters_v1alpha1_filters_proto_goTypes,
		DependencyIndexes: file_commonfate_filters_v1alpha1_filters_proto_depIdxs,
		EnumInfos:         file_commonfate_filters_v1alpha1_filters_proto_enumTypes,
		MessageInfos:      file_commonfate_filters_v1alpha1_filters_proto_msgTypes,
	}.Build()
	File_commonfate_filters_v1alpha1_filters_proto = out.File
	file_commonfate_filters_v1alpha1_filters_proto_rawDesc = nil
	file_commonfate_filters_v1alpha1_filters_proto_goTypes = nil
	file_commonfate_filters_v1alpha1_filters_proto_depIdxs = nil
}
