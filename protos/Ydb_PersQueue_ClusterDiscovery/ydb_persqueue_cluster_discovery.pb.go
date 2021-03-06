// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: protos/ydb_persqueue_cluster_discovery.proto

package Ydb_PersQueue_ClusterDiscovery

import (
	Ydb_Operations "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Operations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WriteSessionClusters_SelectionReason int32

const (
	WriteSessionClusters_SELECTION_REASON_UNSPECIFIED WriteSessionClusters_SelectionReason = 0
	WriteSessionClusters_CLIENT_PREFERENCE            WriteSessionClusters_SelectionReason = 1
	WriteSessionClusters_CLIENT_LOCATION              WriteSessionClusters_SelectionReason = 2
	WriteSessionClusters_CONSISTENT_DISTRIBUTION      WriteSessionClusters_SelectionReason = 3
)

// Enum value maps for WriteSessionClusters_SelectionReason.
var (
	WriteSessionClusters_SelectionReason_name = map[int32]string{
		0: "SELECTION_REASON_UNSPECIFIED",
		1: "CLIENT_PREFERENCE",
		2: "CLIENT_LOCATION",
		3: "CONSISTENT_DISTRIBUTION",
	}
	WriteSessionClusters_SelectionReason_value = map[string]int32{
		"SELECTION_REASON_UNSPECIFIED": 0,
		"CLIENT_PREFERENCE":            1,
		"CLIENT_LOCATION":              2,
		"CONSISTENT_DISTRIBUTION":      3,
	}
)

func (x WriteSessionClusters_SelectionReason) Enum() *WriteSessionClusters_SelectionReason {
	p := new(WriteSessionClusters_SelectionReason)
	*p = x
	return p
}

func (x WriteSessionClusters_SelectionReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WriteSessionClusters_SelectionReason) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_ydb_persqueue_cluster_discovery_proto_enumTypes[0].Descriptor()
}

func (WriteSessionClusters_SelectionReason) Type() protoreflect.EnumType {
	return &file_protos_ydb_persqueue_cluster_discovery_proto_enumTypes[0]
}

func (x WriteSessionClusters_SelectionReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WriteSessionClusters_SelectionReason.Descriptor instead.
func (WriteSessionClusters_SelectionReason) EnumDescriptor() ([]byte, []int) {
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP(), []int{3, 0}
}

type WriteSessionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the topic to write to.
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// Message group identifier.
	SourceId []byte `protobuf:"bytes,2,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	// Partition group to write to. 0 by default.
	PartitionGroup uint32 `protobuf:"varint,3,opt,name=partition_group,json=partitionGroup,proto3" json:"partition_group,omitempty"`
	// Force the specified cluster via its name. Leave it empty by default.
	PreferredClusterName string `protobuf:"bytes,4,opt,name=preferred_cluster_name,json=preferredClusterName,proto3" json:"preferred_cluster_name,omitempty"`
}

func (x *WriteSessionParams) Reset() {
	*x = WriteSessionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteSessionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteSessionParams) ProtoMessage() {}

func (x *WriteSessionParams) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteSessionParams.ProtoReflect.Descriptor instead.
func (*WriteSessionParams) Descriptor() ([]byte, []int) {
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *WriteSessionParams) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *WriteSessionParams) GetSourceId() []byte {
	if x != nil {
		return x.SourceId
	}
	return nil
}

func (x *WriteSessionParams) GetPartitionGroup() uint32 {
	if x != nil {
		return x.PartitionGroup
	}
	return 0
}

func (x *WriteSessionParams) GetPreferredClusterName() string {
	if x != nil {
		return x.PreferredClusterName
	}
	return ""
}

type ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A host discovery endpoint to use at the next step.
	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// An official cluster name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Is the cluster available right now?
	Available bool `protobuf:"varint,3,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *ClusterInfo) Reset() {
	*x = ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInfo) ProtoMessage() {}

func (x *ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterInfo.ProtoReflect.Descriptor instead.
func (*ClusterInfo) Descriptor() ([]byte, []int) {
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterInfo) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ClusterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterInfo) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

type ReadSessionParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the topic to read from.
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// Read mode is set according to the read rule.
	//
	// Types that are assignable to ReadRule:
	//	*ReadSessionParams_MirrorToCluster
	//	*ReadSessionParams_AllOriginal
	ReadRule isReadSessionParams_ReadRule `protobuf_oneof:"read_rule"`
}

func (x *ReadSessionParams) Reset() {
	*x = ReadSessionParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSessionParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSessionParams) ProtoMessage() {}

func (x *ReadSessionParams) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSessionParams.ProtoReflect.Descriptor instead.
func (*ReadSessionParams) Descriptor() ([]byte, []int) {
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *ReadSessionParams) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (m *ReadSessionParams) GetReadRule() isReadSessionParams_ReadRule {
	if m != nil {
		return m.ReadRule
	}
	return nil
}

func (x *ReadSessionParams) GetMirrorToCluster() string {
	if x, ok := x.GetReadRule().(*ReadSessionParams_MirrorToCluster); ok {
		return x.MirrorToCluster
	}
	return ""
}

func (x *ReadSessionParams) GetAllOriginal() *emptypb.Empty {
	if x, ok := x.GetReadRule().(*ReadSessionParams_AllOriginal); ok {
		return x.AllOriginal
	}
	return nil
}

type isReadSessionParams_ReadRule interface {
	isReadSessionParams_ReadRule()
}

type ReadSessionParams_MirrorToCluster struct {
	MirrorToCluster string `protobuf:"bytes,2,opt,name=mirror_to_cluster,json=mirrorToCluster,proto3,oneof"`
}

type ReadSessionParams_AllOriginal struct {
	AllOriginal *emptypb.Empty `protobuf:"bytes,3,opt,name=all_original,json=allOriginal,proto3,oneof"`
}

func (*ReadSessionParams_MirrorToCluster) isReadSessionParams_ReadRule() {}

func (*ReadSessionParams_AllOriginal) isReadSessionParams_ReadRule() {}

type WriteSessionClusters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ordered clusters with statuses.
	Clusters []*ClusterInfo `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// The reason why a particular cluster was prioritized.
	PrimaryClusterSelectionReason WriteSessionClusters_SelectionReason `protobuf:"varint,2,opt,name=primary_cluster_selection_reason,json=primaryClusterSelectionReason,proto3,enum=Ydb.PersQueue.ClusterDiscovery.WriteSessionClusters_SelectionReason" json:"primary_cluster_selection_reason,omitempty"`
}

func (x *WriteSessionClusters) Reset() {
	*x = WriteSessionClusters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteSessionClusters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteSessionClusters) ProtoMessage() {}

func (x *WriteSessionClusters) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteSessionClusters.ProtoReflect.Descriptor instead.
func (*WriteSessionClusters) Descriptor() ([]byte, []int) {
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *WriteSessionClusters) GetClusters() []*ClusterInfo {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *WriteSessionClusters) GetPrimaryClusterSelectionReason() WriteSessionClusters_SelectionReason {
	if x != nil {
		return x.PrimaryClusterSelectionReason
	}
	return WriteSessionClusters_SELECTION_REASON_UNSPECIFIED
}

type ReadSessionClusters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ordered clusters with statuses.
	Clusters []*ClusterInfo `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *ReadSessionClusters) Reset() {
	*x = ReadSessionClusters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSessionClusters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSessionClusters) ProtoMessage() {}

func (x *ReadSessionClusters) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSessionClusters.ProtoReflect.Descriptor instead.
func (*ReadSessionClusters) Descriptor() ([]byte, []int) {
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP(), []int{4}
}

func (x *ReadSessionClusters) GetClusters() []*ClusterInfo {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type DiscoverClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationParams *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	// Clusters will be discovered separately for each element of the list.
	WriteSessions []*WriteSessionParams `protobuf:"bytes,2,rep,name=write_sessions,json=writeSessions,proto3" json:"write_sessions,omitempty"`
	ReadSessions  []*ReadSessionParams  `protobuf:"bytes,3,rep,name=read_sessions,json=readSessions,proto3" json:"read_sessions,omitempty"`
	// Latest clusters status version known to the client application. Use 0 by default.
	MinimalVersion int64 `protobuf:"varint,4,opt,name=minimal_version,json=minimalVersion,proto3" json:"minimal_version,omitempty"`
}

func (x *DiscoverClustersRequest) Reset() {
	*x = DiscoverClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverClustersRequest) ProtoMessage() {}

func (x *DiscoverClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverClustersRequest.ProtoReflect.Descriptor instead.
func (*DiscoverClustersRequest) Descriptor() ([]byte, []int) {
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP(), []int{5}
}

func (x *DiscoverClustersRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if x != nil {
		return x.OperationParams
	}
	return nil
}

func (x *DiscoverClustersRequest) GetWriteSessions() []*WriteSessionParams {
	if x != nil {
		return x.WriteSessions
	}
	return nil
}

func (x *DiscoverClustersRequest) GetReadSessions() []*ReadSessionParams {
	if x != nil {
		return x.ReadSessions
	}
	return nil
}

func (x *DiscoverClustersRequest) GetMinimalVersion() int64 {
	if x != nil {
		return x.MinimalVersion
	}
	return 0
}

type DiscoverClustersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation contains the result of the request. Check the ydb_operation.proto.
	Operation *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *DiscoverClustersResponse) Reset() {
	*x = DiscoverClustersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverClustersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverClustersResponse) ProtoMessage() {}

func (x *DiscoverClustersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverClustersResponse.ProtoReflect.Descriptor instead.
func (*DiscoverClustersResponse) Descriptor() ([]byte, []int) {
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP(), []int{6}
}

func (x *DiscoverClustersResponse) GetOperation() *Ydb_Operations.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type DiscoverClustersResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Discovered per-session clusters.
	WriteSessionsClusters []*WriteSessionClusters `protobuf:"bytes,1,rep,name=write_sessions_clusters,json=writeSessionsClusters,proto3" json:"write_sessions_clusters,omitempty"`
	ReadSessionsClusters  []*ReadSessionClusters  `protobuf:"bytes,2,rep,name=read_sessions_clusters,json=readSessionsClusters,proto3" json:"read_sessions_clusters,omitempty"`
	// Latest clusters status version known to the cluster discovery service.
	Version int64 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *DiscoverClustersResult) Reset() {
	*x = DiscoverClustersResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverClustersResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverClustersResult) ProtoMessage() {}

func (x *DiscoverClustersResult) ProtoReflect() protoreflect.Message {
	mi := &file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverClustersResult.ProtoReflect.Descriptor instead.
func (*DiscoverClustersResult) Descriptor() ([]byte, []int) {
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP(), []int{7}
}

func (x *DiscoverClustersResult) GetWriteSessionsClusters() []*WriteSessionClusters {
	if x != nil {
		return x.WriteSessionsClusters
	}
	return nil
}

func (x *DiscoverClustersResult) GetReadSessionsClusters() []*ReadSessionClusters {
	if x != nil {
		return x.ReadSessionsClusters
	}
	return nil
}

func (x *DiscoverClustersResult) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_protos_ydb_persqueue_cluster_discovery_proto protoreflect.FileDescriptor

var file_protos_ydb_persqueue_cluster_discovery_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x72,
	0x73, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x12, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x5b, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xa1, 0x01,
	0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x6f,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x5f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x22, 0xed, 0x02, 0x0a, 0x14, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x08, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x20, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x44,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x1d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x0f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4c, 0x49, 0x45,
	0x4e, 0x54, 0x5f, 0x50, 0x52, 0x45, 0x46, 0x45, 0x52, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45,
	0x4e, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x03, 0x22, 0x5e, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x47, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x22, 0xc1, 0x02, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x59, 0x0a, 0x0e, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x56, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0c,
	0x72, 0x65, 0x61, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f,
	0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x18, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8b, 0x02, 0x0a, 0x16, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x6c, 0x0a, 0x17, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x15, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x69, 0x0a, 0x16, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x14, 0x72, 0x65, 0x61, 0x64, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x78, 0x0a, 0x24, 0x74, 0x65, 0x63, 0x68,
	0x2e, 0x79, 0x64, 0x62, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62,
	0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x59, 0x64, 0x62, 0x5f, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_ydb_persqueue_cluster_discovery_proto_rawDescOnce sync.Once
	file_protos_ydb_persqueue_cluster_discovery_proto_rawDescData = file_protos_ydb_persqueue_cluster_discovery_proto_rawDesc
)

func file_protos_ydb_persqueue_cluster_discovery_proto_rawDescGZIP() []byte {
	file_protos_ydb_persqueue_cluster_discovery_proto_rawDescOnce.Do(func() {
		file_protos_ydb_persqueue_cluster_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_ydb_persqueue_cluster_discovery_proto_rawDescData)
	})
	return file_protos_ydb_persqueue_cluster_discovery_proto_rawDescData
}

var file_protos_ydb_persqueue_cluster_discovery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_ydb_persqueue_cluster_discovery_proto_goTypes = []interface{}{
	(WriteSessionClusters_SelectionReason)(0), // 0: Ydb.PersQueue.ClusterDiscovery.WriteSessionClusters.SelectionReason
	(*WriteSessionParams)(nil),                // 1: Ydb.PersQueue.ClusterDiscovery.WriteSessionParams
	(*ClusterInfo)(nil),                       // 2: Ydb.PersQueue.ClusterDiscovery.ClusterInfo
	(*ReadSessionParams)(nil),                 // 3: Ydb.PersQueue.ClusterDiscovery.ReadSessionParams
	(*WriteSessionClusters)(nil),              // 4: Ydb.PersQueue.ClusterDiscovery.WriteSessionClusters
	(*ReadSessionClusters)(nil),               // 5: Ydb.PersQueue.ClusterDiscovery.ReadSessionClusters
	(*DiscoverClustersRequest)(nil),           // 6: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersRequest
	(*DiscoverClustersResponse)(nil),          // 7: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersResponse
	(*DiscoverClustersResult)(nil),            // 8: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersResult
	(*emptypb.Empty)(nil),                     // 9: google.protobuf.Empty
	(*Ydb_Operations.OperationParams)(nil),    // 10: Ydb.Operations.OperationParams
	(*Ydb_Operations.Operation)(nil),          // 11: Ydb.Operations.Operation
}
var file_protos_ydb_persqueue_cluster_discovery_proto_depIdxs = []int32{
	9,  // 0: Ydb.PersQueue.ClusterDiscovery.ReadSessionParams.all_original:type_name -> google.protobuf.Empty
	2,  // 1: Ydb.PersQueue.ClusterDiscovery.WriteSessionClusters.clusters:type_name -> Ydb.PersQueue.ClusterDiscovery.ClusterInfo
	0,  // 2: Ydb.PersQueue.ClusterDiscovery.WriteSessionClusters.primary_cluster_selection_reason:type_name -> Ydb.PersQueue.ClusterDiscovery.WriteSessionClusters.SelectionReason
	2,  // 3: Ydb.PersQueue.ClusterDiscovery.ReadSessionClusters.clusters:type_name -> Ydb.PersQueue.ClusterDiscovery.ClusterInfo
	10, // 4: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersRequest.operation_params:type_name -> Ydb.Operations.OperationParams
	1,  // 5: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersRequest.write_sessions:type_name -> Ydb.PersQueue.ClusterDiscovery.WriteSessionParams
	3,  // 6: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersRequest.read_sessions:type_name -> Ydb.PersQueue.ClusterDiscovery.ReadSessionParams
	11, // 7: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersResponse.operation:type_name -> Ydb.Operations.Operation
	4,  // 8: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersResult.write_sessions_clusters:type_name -> Ydb.PersQueue.ClusterDiscovery.WriteSessionClusters
	5,  // 9: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersResult.read_sessions_clusters:type_name -> Ydb.PersQueue.ClusterDiscovery.ReadSessionClusters
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_protos_ydb_persqueue_cluster_discovery_proto_init() }
func file_protos_ydb_persqueue_cluster_discovery_proto_init() {
	if File_protos_ydb_persqueue_cluster_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteSessionParams); i {
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
		file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterInfo); i {
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
		file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSessionParams); i {
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
		file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteSessionClusters); i {
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
		file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSessionClusters); i {
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
		file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverClustersRequest); i {
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
		file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverClustersResponse); i {
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
		file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverClustersResult); i {
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
	file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ReadSessionParams_MirrorToCluster)(nil),
		(*ReadSessionParams_AllOriginal)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_ydb_persqueue_cluster_discovery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_ydb_persqueue_cluster_discovery_proto_goTypes,
		DependencyIndexes: file_protos_ydb_persqueue_cluster_discovery_proto_depIdxs,
		EnumInfos:         file_protos_ydb_persqueue_cluster_discovery_proto_enumTypes,
		MessageInfos:      file_protos_ydb_persqueue_cluster_discovery_proto_msgTypes,
	}.Build()
	File_protos_ydb_persqueue_cluster_discovery_proto = out.File
	file_protos_ydb_persqueue_cluster_discovery_proto_rawDesc = nil
	file_protos_ydb_persqueue_cluster_discovery_proto_goTypes = nil
	file_protos_ydb_persqueue_cluster_discovery_proto_depIdxs = nil
}
