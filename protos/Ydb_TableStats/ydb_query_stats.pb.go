// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: protos/ydb_query_stats.proto

package Ydb_TableStats

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

// Describes select, update (insert, upsert, replace) and delete operations
type OperationStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows  uint64 `protobuf:"varint,1,opt,name=rows,proto3" json:"rows,omitempty"`
	Bytes uint64 `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *OperationStats) Reset() {
	*x = OperationStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_query_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationStats) ProtoMessage() {}

func (x *OperationStats) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_query_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationStats.ProtoReflect.Descriptor instead.
func (*OperationStats) Descriptor() ([]byte, []int) {
	return file_protos_ydb_query_stats_proto_rawDescGZIP(), []int{0}
}

func (x *OperationStats) GetRows() uint64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *OperationStats) GetBytes() uint64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

// Describes all operations on a table
type TableAccessStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reads           *OperationStats `protobuf:"bytes,3,opt,name=reads,proto3" json:"reads,omitempty"`
	Updates         *OperationStats `protobuf:"bytes,4,opt,name=updates,proto3" json:"updates,omitempty"`
	Deletes         *OperationStats `protobuf:"bytes,5,opt,name=deletes,proto3" json:"deletes,omitempty"`
	PartitionsCount uint64          `protobuf:"varint,6,opt,name=partitions_count,json=partitionsCount,proto3" json:"partitions_count,omitempty"`
}

func (x *TableAccessStats) Reset() {
	*x = TableAccessStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_query_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableAccessStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableAccessStats) ProtoMessage() {}

func (x *TableAccessStats) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_query_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableAccessStats.ProtoReflect.Descriptor instead.
func (*TableAccessStats) Descriptor() ([]byte, []int) {
	return file_protos_ydb_query_stats_proto_rawDescGZIP(), []int{1}
}

func (x *TableAccessStats) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TableAccessStats) GetReads() *OperationStats {
	if x != nil {
		return x.Reads
	}
	return nil
}

func (x *TableAccessStats) GetUpdates() *OperationStats {
	if x != nil {
		return x.Updates
	}
	return nil
}

func (x *TableAccessStats) GetDeletes() *OperationStats {
	if x != nil {
		return x.Deletes
	}
	return nil
}

func (x *TableAccessStats) GetPartitionsCount() uint64 {
	if x != nil {
		return x.PartitionsCount
	}
	return 0
}

type QueryPhaseStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DurationUs     uint64              `protobuf:"varint,1,opt,name=duration_us,json=durationUs,proto3" json:"duration_us,omitempty"`
	TableAccess    []*TableAccessStats `protobuf:"bytes,2,rep,name=table_access,json=tableAccess,proto3" json:"table_access,omitempty"`
	CpuTimeUs      uint64              `protobuf:"varint,3,opt,name=cpu_time_us,json=cpuTimeUs,proto3" json:"cpu_time_us,omitempty"`
	AffectedShards uint64              `protobuf:"varint,4,opt,name=affected_shards,json=affectedShards,proto3" json:"affected_shards,omitempty"`
	LiteralPhase   bool                `protobuf:"varint,5,opt,name=literal_phase,json=literalPhase,proto3" json:"literal_phase,omitempty"`
}

func (x *QueryPhaseStats) Reset() {
	*x = QueryPhaseStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_query_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPhaseStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPhaseStats) ProtoMessage() {}

func (x *QueryPhaseStats) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_query_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPhaseStats.ProtoReflect.Descriptor instead.
func (*QueryPhaseStats) Descriptor() ([]byte, []int) {
	return file_protos_ydb_query_stats_proto_rawDescGZIP(), []int{2}
}

func (x *QueryPhaseStats) GetDurationUs() uint64 {
	if x != nil {
		return x.DurationUs
	}
	return 0
}

func (x *QueryPhaseStats) GetTableAccess() []*TableAccessStats {
	if x != nil {
		return x.TableAccess
	}
	return nil
}

func (x *QueryPhaseStats) GetCpuTimeUs() uint64 {
	if x != nil {
		return x.CpuTimeUs
	}
	return 0
}

func (x *QueryPhaseStats) GetAffectedShards() uint64 {
	if x != nil {
		return x.AffectedShards
	}
	return 0
}

func (x *QueryPhaseStats) GetLiteralPhase() bool {
	if x != nil {
		return x.LiteralPhase
	}
	return false
}

type CompilationStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromCache  bool   `protobuf:"varint,1,opt,name=from_cache,json=fromCache,proto3" json:"from_cache,omitempty"`
	DurationUs uint64 `protobuf:"varint,2,opt,name=duration_us,json=durationUs,proto3" json:"duration_us,omitempty"`
	CpuTimeUs  uint64 `protobuf:"varint,3,opt,name=cpu_time_us,json=cpuTimeUs,proto3" json:"cpu_time_us,omitempty"`
}

func (x *CompilationStats) Reset() {
	*x = CompilationStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_query_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationStats) ProtoMessage() {}

func (x *CompilationStats) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_query_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationStats.ProtoReflect.Descriptor instead.
func (*CompilationStats) Descriptor() ([]byte, []int) {
	return file_protos_ydb_query_stats_proto_rawDescGZIP(), []int{3}
}

func (x *CompilationStats) GetFromCache() bool {
	if x != nil {
		return x.FromCache
	}
	return false
}

func (x *CompilationStats) GetDurationUs() uint64 {
	if x != nil {
		return x.DurationUs
	}
	return 0
}

func (x *CompilationStats) GetCpuTimeUs() uint64 {
	if x != nil {
		return x.CpuTimeUs
	}
	return 0
}

type QueryStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A query might have one or more execution phases
	QueryPhases      []*QueryPhaseStats `protobuf:"bytes,1,rep,name=query_phases,json=queryPhases,proto3" json:"query_phases,omitempty"`
	Compilation      *CompilationStats  `protobuf:"bytes,2,opt,name=compilation,proto3" json:"compilation,omitempty"`
	ProcessCpuTimeUs uint64             `protobuf:"varint,3,opt,name=process_cpu_time_us,json=processCpuTimeUs,proto3" json:"process_cpu_time_us,omitempty"`
	QueryPlan        string             `protobuf:"bytes,4,opt,name=query_plan,json=queryPlan,proto3" json:"query_plan,omitempty"`
	QueryAst         string             `protobuf:"bytes,5,opt,name=query_ast,json=queryAst,proto3" json:"query_ast,omitempty"`
	TotalDurationUs  uint64             `protobuf:"varint,6,opt,name=total_duration_us,json=totalDurationUs,proto3" json:"total_duration_us,omitempty"`
	TotalCpuTimeUs   uint64             `protobuf:"varint,7,opt,name=total_cpu_time_us,json=totalCpuTimeUs,proto3" json:"total_cpu_time_us,omitempty"`
}

func (x *QueryStats) Reset() {
	*x = QueryStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_query_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStats) ProtoMessage() {}

func (x *QueryStats) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_query_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryStats.ProtoReflect.Descriptor instead.
func (*QueryStats) Descriptor() ([]byte, []int) {
	return file_protos_ydb_query_stats_proto_rawDescGZIP(), []int{4}
}

func (x *QueryStats) GetQueryPhases() []*QueryPhaseStats {
	if x != nil {
		return x.QueryPhases
	}
	return nil
}

func (x *QueryStats) GetCompilation() *CompilationStats {
	if x != nil {
		return x.Compilation
	}
	return nil
}

func (x *QueryStats) GetProcessCpuTimeUs() uint64 {
	if x != nil {
		return x.ProcessCpuTimeUs
	}
	return 0
}

func (x *QueryStats) GetQueryPlan() string {
	if x != nil {
		return x.QueryPlan
	}
	return ""
}

func (x *QueryStats) GetQueryAst() string {
	if x != nil {
		return x.QueryAst
	}
	return ""
}

func (x *QueryStats) GetTotalDurationUs() uint64 {
	if x != nil {
		return x.TotalDurationUs
	}
	return 0
}

func (x *QueryStats) GetTotalCpuTimeUs() uint64 {
	if x != nil {
		return x.TotalCpuTimeUs
	}
	return 0
}

var File_protos_ydb_query_stats_proto protoreflect.FileDescriptor

var file_protos_ydb_query_stats_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x3a,
	0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x81, 0x02, 0x0a, 0x10, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xe5,
	0x01, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x12, 0x43, 0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x63, 0x70, 0x75, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63,
	0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0e, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x70, 0x68, 0x61,
	0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61,
	0x6c, 0x50, 0x68, 0x61, 0x73, 0x65, 0x22, 0x72, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x66, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x63, 0x70,
	0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x63, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x22, 0xd6, 0x02, 0x0a, 0x0a, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x0c, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x68, 0x61, 0x73, 0x65, 0x73, 0x12, 0x42, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2d, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x70, 0x75,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x43, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x41, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x12, 0x29, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x70, 0x75, 0x54, 0x69, 0x6d,
	0x65, 0x55, 0x73, 0x42, 0x4c, 0x0a, 0x08, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x59, 0x64, 0x62, 0x5f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0xf8, 0x01,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_ydb_query_stats_proto_rawDescOnce sync.Once
	file_protos_ydb_query_stats_proto_rawDescData = file_protos_ydb_query_stats_proto_rawDesc
)

func file_protos_ydb_query_stats_proto_rawDescGZIP() []byte {
	file_protos_ydb_query_stats_proto_rawDescOnce.Do(func() {
		file_protos_ydb_query_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_ydb_query_stats_proto_rawDescData)
	})
	return file_protos_ydb_query_stats_proto_rawDescData
}

var file_protos_ydb_query_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_ydb_query_stats_proto_goTypes = []interface{}{
	(*OperationStats)(nil),   // 0: Ydb.TableStats.OperationStats
	(*TableAccessStats)(nil), // 1: Ydb.TableStats.TableAccessStats
	(*QueryPhaseStats)(nil),  // 2: Ydb.TableStats.QueryPhaseStats
	(*CompilationStats)(nil), // 3: Ydb.TableStats.CompilationStats
	(*QueryStats)(nil),       // 4: Ydb.TableStats.QueryStats
}
var file_protos_ydb_query_stats_proto_depIdxs = []int32{
	0, // 0: Ydb.TableStats.TableAccessStats.reads:type_name -> Ydb.TableStats.OperationStats
	0, // 1: Ydb.TableStats.TableAccessStats.updates:type_name -> Ydb.TableStats.OperationStats
	0, // 2: Ydb.TableStats.TableAccessStats.deletes:type_name -> Ydb.TableStats.OperationStats
	1, // 3: Ydb.TableStats.QueryPhaseStats.table_access:type_name -> Ydb.TableStats.TableAccessStats
	2, // 4: Ydb.TableStats.QueryStats.query_phases:type_name -> Ydb.TableStats.QueryPhaseStats
	3, // 5: Ydb.TableStats.QueryStats.compilation:type_name -> Ydb.TableStats.CompilationStats
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protos_ydb_query_stats_proto_init() }
func file_protos_ydb_query_stats_proto_init() {
	if File_protos_ydb_query_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_ydb_query_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationStats); i {
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
		file_protos_ydb_query_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableAccessStats); i {
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
		file_protos_ydb_query_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPhaseStats); i {
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
		file_protos_ydb_query_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationStats); i {
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
		file_protos_ydb_query_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryStats); i {
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
			RawDescriptor: file_protos_ydb_query_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_ydb_query_stats_proto_goTypes,
		DependencyIndexes: file_protos_ydb_query_stats_proto_depIdxs,
		MessageInfos:      file_protos_ydb_query_stats_proto_msgTypes,
	}.Build()
	File_protos_ydb_query_stats_proto = out.File
	file_protos_ydb_query_stats_proto_rawDesc = nil
	file_protos_ydb_query_stats_proto_goTypes = nil
	file_protos_ydb_query_stats_proto_depIdxs = nil
}
