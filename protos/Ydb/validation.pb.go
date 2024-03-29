// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: protos/annotations/validation.proto

package Ydb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Limit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*Limit_Range_
	//	*Limit_Lt
	//	*Limit_Le
	//	*Limit_Eq
	//	*Limit_Ge
	//	*Limit_Gt
	Kind isLimit_Kind `protobuf_oneof:"kind"`
}

func (x *Limit) Reset() {
	*x = Limit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_annotations_validation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Limit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Limit) ProtoMessage() {}

func (x *Limit) ProtoReflect() protoreflect.Message {
	mi := &file_protos_annotations_validation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Limit.ProtoReflect.Descriptor instead.
func (*Limit) Descriptor() ([]byte, []int) {
	return file_protos_annotations_validation_proto_rawDescGZIP(), []int{0}
}

func (m *Limit) GetKind() isLimit_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Limit) GetRange() *Limit_Range {
	if x, ok := x.GetKind().(*Limit_Range_); ok {
		return x.Range
	}
	return nil
}

func (x *Limit) GetLt() uint32 {
	if x, ok := x.GetKind().(*Limit_Lt); ok {
		return x.Lt
	}
	return 0
}

func (x *Limit) GetLe() uint32 {
	if x, ok := x.GetKind().(*Limit_Le); ok {
		return x.Le
	}
	return 0
}

func (x *Limit) GetEq() uint32 {
	if x, ok := x.GetKind().(*Limit_Eq); ok {
		return x.Eq
	}
	return 0
}

func (x *Limit) GetGe() uint32 {
	if x, ok := x.GetKind().(*Limit_Ge); ok {
		return x.Ge
	}
	return 0
}

func (x *Limit) GetGt() uint32 {
	if x, ok := x.GetKind().(*Limit_Gt); ok {
		return x.Gt
	}
	return 0
}

type isLimit_Kind interface {
	isLimit_Kind()
}

type Limit_Range_ struct {
	Range *Limit_Range `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}

type Limit_Lt struct {
	Lt uint32 `protobuf:"varint,2,opt,name=lt,proto3,oneof"`
}

type Limit_Le struct {
	Le uint32 `protobuf:"varint,3,opt,name=le,proto3,oneof"`
}

type Limit_Eq struct {
	Eq uint32 `protobuf:"varint,4,opt,name=eq,proto3,oneof"`
}

type Limit_Ge struct {
	Ge uint32 `protobuf:"varint,5,opt,name=ge,proto3,oneof"`
}

type Limit_Gt struct {
	Gt uint32 `protobuf:"varint,6,opt,name=gt,proto3,oneof"`
}

func (*Limit_Range_) isLimit_Kind() {}

func (*Limit_Lt) isLimit_Kind() {}

func (*Limit_Le) isLimit_Kind() {}

func (*Limit_Eq) isLimit_Kind() {}

func (*Limit_Ge) isLimit_Kind() {}

func (*Limit_Gt) isLimit_Kind() {}

type MapKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length *Limit `protobuf:"bytes,1,opt,name=length,proto3" json:"length,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MapKey) Reset() {
	*x = MapKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_annotations_validation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapKey) ProtoMessage() {}

func (x *MapKey) ProtoReflect() protoreflect.Message {
	mi := &file_protos_annotations_validation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapKey.ProtoReflect.Descriptor instead.
func (*MapKey) Descriptor() ([]byte, []int) {
	return file_protos_annotations_validation_proto_rawDescGZIP(), []int{1}
}

func (x *MapKey) GetLength() *Limit {
	if x != nil {
		return x.Length
	}
	return nil
}

func (x *MapKey) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Limit_Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min uint32 `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max uint32 `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *Limit_Range) Reset() {
	*x = Limit_Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_annotations_validation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Limit_Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Limit_Range) ProtoMessage() {}

func (x *Limit_Range) ProtoReflect() protoreflect.Message {
	mi := &file_protos_annotations_validation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Limit_Range.ProtoReflect.Descriptor instead.
func (*Limit_Range) Descriptor() ([]byte, []int) {
	return file_protos_annotations_validation_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Limit_Range) GetMin() uint32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Limit_Range) GetMax() uint32 {
	if x != nil {
		return x.Max
	}
	return 0
}

var file_protos_annotations_validation_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         87650,
		Name:          "Ydb.required",
		Tag:           "varint,87650,opt,name=required",
		Filename:      "protos/annotations/validation.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Limit)(nil),
		Field:         87651,
		Name:          "Ydb.size",
		Tag:           "bytes,87651,opt,name=size",
		Filename:      "protos/annotations/validation.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Limit)(nil),
		Field:         87652,
		Name:          "Ydb.length",
		Tag:           "bytes,87652,opt,name=length",
		Filename:      "protos/annotations/validation.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*MapKey)(nil),
		Field:         87653,
		Name:          "Ydb.map_key",
		Tag:           "bytes,87653,opt,name=map_key",
		Filename:      "protos/annotations/validation.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         87654,
		Name:          "Ydb.value",
		Tag:           "bytes,87654,opt,name=value",
		Filename:      "protos/annotations/validation.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool required = 87650;
	E_Required = &file_protos_annotations_validation_proto_extTypes[0]
	// optional Ydb.Limit size = 87651;
	E_Size = &file_protos_annotations_validation_proto_extTypes[1]
	// optional Ydb.Limit length = 87652;
	E_Length = &file_protos_annotations_validation_proto_extTypes[2]
	// optional Ydb.MapKey map_key = 87653;
	E_MapKey = &file_protos_annotations_validation_proto_extTypes[3]
	// optional string value = 87654;
	E_Value = &file_protos_annotations_validation_proto_extTypes[4]
)

var File_protos_annotations_validation_proto protoreflect.FileDescriptor

var file_protos_annotations_validation_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x59, 0x64, 0x62, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x02, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02,
	0x6c, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00,
	0x52, 0x02, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x02, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x02, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x02, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x02, 0x67, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x67, 0x74, 0x1a, 0x2b, 0x0a, 0x05, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22,
	0x42, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe2,
	0xac, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x3a, 0x3f, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe3, 0xac, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x3a, 0x43, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe4, 0xac, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x45, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x5f, 0x6b, 0x65,
	0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xe5, 0xac, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4d,
	0x61, 0x70, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x6d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x3a, 0x35, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe6, 0xac, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x47, 0x0a, 0x0e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x59, 0x64, 0x62, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_annotations_validation_proto_rawDescOnce sync.Once
	file_protos_annotations_validation_proto_rawDescData = file_protos_annotations_validation_proto_rawDesc
)

func file_protos_annotations_validation_proto_rawDescGZIP() []byte {
	file_protos_annotations_validation_proto_rawDescOnce.Do(func() {
		file_protos_annotations_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_annotations_validation_proto_rawDescData)
	})
	return file_protos_annotations_validation_proto_rawDescData
}

var file_protos_annotations_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_annotations_validation_proto_goTypes = []interface{}{
	(*Limit)(nil),                     // 0: Ydb.Limit
	(*MapKey)(nil),                    // 1: Ydb.MapKey
	(*Limit_Range)(nil),               // 2: Ydb.Limit.Range
	(*descriptorpb.FieldOptions)(nil), // 3: google.protobuf.FieldOptions
}
var file_protos_annotations_validation_proto_depIdxs = []int32{
	2,  // 0: Ydb.Limit.range:type_name -> Ydb.Limit.Range
	0,  // 1: Ydb.MapKey.length:type_name -> Ydb.Limit
	3,  // 2: Ydb.required:extendee -> google.protobuf.FieldOptions
	3,  // 3: Ydb.size:extendee -> google.protobuf.FieldOptions
	3,  // 4: Ydb.length:extendee -> google.protobuf.FieldOptions
	3,  // 5: Ydb.map_key:extendee -> google.protobuf.FieldOptions
	3,  // 6: Ydb.value:extendee -> google.protobuf.FieldOptions
	0,  // 7: Ydb.size:type_name -> Ydb.Limit
	0,  // 8: Ydb.length:type_name -> Ydb.Limit
	1,  // 9: Ydb.map_key:type_name -> Ydb.MapKey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	7,  // [7:10] is the sub-list for extension type_name
	2,  // [2:7] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_protos_annotations_validation_proto_init() }
func file_protos_annotations_validation_proto_init() {
	if File_protos_annotations_validation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_annotations_validation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Limit); i {
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
		file_protos_annotations_validation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapKey); i {
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
		file_protos_annotations_validation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Limit_Range); i {
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
	file_protos_annotations_validation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Limit_Range_)(nil),
		(*Limit_Lt)(nil),
		(*Limit_Le)(nil),
		(*Limit_Eq)(nil),
		(*Limit_Ge)(nil),
		(*Limit_Gt)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_annotations_validation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_protos_annotations_validation_proto_goTypes,
		DependencyIndexes: file_protos_annotations_validation_proto_depIdxs,
		MessageInfos:      file_protos_annotations_validation_proto_msgTypes,
		ExtensionInfos:    file_protos_annotations_validation_proto_extTypes,
	}.Build()
	File_protos_annotations_validation_proto = out.File
	file_protos_annotations_validation_proto_rawDesc = nil
	file_protos_annotations_validation_proto_goTypes = nil
	file_protos_annotations_validation_proto_depIdxs = nil
}
