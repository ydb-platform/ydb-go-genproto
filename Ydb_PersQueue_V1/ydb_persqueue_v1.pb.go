// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: ydb_persqueue_v1.proto

package Ydb_PersQueue_V1

import (
	Ydb_PersQueue_ClusterDiscovery "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_PersQueue_ClusterDiscovery"
	Ydb_PersQueue_V1 "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_PersQueue_V1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ydb_persqueue_v1_proto protoreflect.FileDescriptor

var file_ydb_persqueue_v1_proto_rawDesc = []byte{
	0x0a, 0x16, 0x79, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f,
	0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x79, 0x64, 0x62, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x76,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9f, 0x07, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x73,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x0e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x2d,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2d, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01,
	0x12, 0x8a, 0x01, 0x0a, 0x16, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x61, 0x64, 0x12, 0x35, 0x2e, 0x59, 0x64,
	0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x35, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x5c, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x0d, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x26, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a,
	0x09, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x22, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x44, 0x72,
	0x6f, 0x70, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56,
	0x31, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x57, 0x0a, 0x0a, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x23, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31,
	0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x61, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e,
	0x56, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x61, 0x64, 0x52, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa1, 0x01, 0x0a, 0x17, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x37, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x51, 0x0a,
	0x15, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x59, 0x64, 0x62, 0x5f, 0x50, 0x65, 0x72, 0x73, 0x51, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_persqueue_v1_proto_goTypes = []interface{}{
	(*Ydb_PersQueue_V1.StreamingWriteClientMessage)(nil),            // 0: Ydb.PersQueue.V1.StreamingWriteClientMessage
	(*Ydb_PersQueue_V1.MigrationStreamingReadClientMessage)(nil),    // 1: Ydb.PersQueue.V1.MigrationStreamingReadClientMessage
	(*Ydb_PersQueue_V1.ReadInfoRequest)(nil),                        // 2: Ydb.PersQueue.V1.ReadInfoRequest
	(*Ydb_PersQueue_V1.DescribeTopicRequest)(nil),                   // 3: Ydb.PersQueue.V1.DescribeTopicRequest
	(*Ydb_PersQueue_V1.DropTopicRequest)(nil),                       // 4: Ydb.PersQueue.V1.DropTopicRequest
	(*Ydb_PersQueue_V1.CreateTopicRequest)(nil),                     // 5: Ydb.PersQueue.V1.CreateTopicRequest
	(*Ydb_PersQueue_V1.AlterTopicRequest)(nil),                      // 6: Ydb.PersQueue.V1.AlterTopicRequest
	(*Ydb_PersQueue_V1.AddReadRuleRequest)(nil),                     // 7: Ydb.PersQueue.V1.AddReadRuleRequest
	(*Ydb_PersQueue_V1.RemoveReadRuleRequest)(nil),                  // 8: Ydb.PersQueue.V1.RemoveReadRuleRequest
	(*Ydb_PersQueue_ClusterDiscovery.DiscoverClustersRequest)(nil),  // 9: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersRequest
	(*Ydb_PersQueue_V1.StreamingWriteServerMessage)(nil),            // 10: Ydb.PersQueue.V1.StreamingWriteServerMessage
	(*Ydb_PersQueue_V1.MigrationStreamingReadServerMessage)(nil),    // 11: Ydb.PersQueue.V1.MigrationStreamingReadServerMessage
	(*Ydb_PersQueue_V1.ReadInfoResponse)(nil),                       // 12: Ydb.PersQueue.V1.ReadInfoResponse
	(*Ydb_PersQueue_V1.DescribeTopicResponse)(nil),                  // 13: Ydb.PersQueue.V1.DescribeTopicResponse
	(*Ydb_PersQueue_V1.DropTopicResponse)(nil),                      // 14: Ydb.PersQueue.V1.DropTopicResponse
	(*Ydb_PersQueue_V1.CreateTopicResponse)(nil),                    // 15: Ydb.PersQueue.V1.CreateTopicResponse
	(*Ydb_PersQueue_V1.AlterTopicResponse)(nil),                     // 16: Ydb.PersQueue.V1.AlterTopicResponse
	(*Ydb_PersQueue_V1.AddReadRuleResponse)(nil),                    // 17: Ydb.PersQueue.V1.AddReadRuleResponse
	(*Ydb_PersQueue_V1.RemoveReadRuleResponse)(nil),                 // 18: Ydb.PersQueue.V1.RemoveReadRuleResponse
	(*Ydb_PersQueue_ClusterDiscovery.DiscoverClustersResponse)(nil), // 19: Ydb.PersQueue.ClusterDiscovery.DiscoverClustersResponse
}
var file_ydb_persqueue_v1_proto_depIdxs = []int32{
	0,  // 0: Ydb.PersQueue.V1.PersQueueService.StreamingWrite:input_type -> Ydb.PersQueue.V1.StreamingWriteClientMessage
	1,  // 1: Ydb.PersQueue.V1.PersQueueService.MigrationStreamingRead:input_type -> Ydb.PersQueue.V1.MigrationStreamingReadClientMessage
	2,  // 2: Ydb.PersQueue.V1.PersQueueService.GetReadSessionsInfo:input_type -> Ydb.PersQueue.V1.ReadInfoRequest
	3,  // 3: Ydb.PersQueue.V1.PersQueueService.DescribeTopic:input_type -> Ydb.PersQueue.V1.DescribeTopicRequest
	4,  // 4: Ydb.PersQueue.V1.PersQueueService.DropTopic:input_type -> Ydb.PersQueue.V1.DropTopicRequest
	5,  // 5: Ydb.PersQueue.V1.PersQueueService.CreateTopic:input_type -> Ydb.PersQueue.V1.CreateTopicRequest
	6,  // 6: Ydb.PersQueue.V1.PersQueueService.AlterTopic:input_type -> Ydb.PersQueue.V1.AlterTopicRequest
	7,  // 7: Ydb.PersQueue.V1.PersQueueService.AddReadRule:input_type -> Ydb.PersQueue.V1.AddReadRuleRequest
	8,  // 8: Ydb.PersQueue.V1.PersQueueService.RemoveReadRule:input_type -> Ydb.PersQueue.V1.RemoveReadRuleRequest
	9,  // 9: Ydb.PersQueue.V1.ClusterDiscoveryService.DiscoverClusters:input_type -> Ydb.PersQueue.ClusterDiscovery.DiscoverClustersRequest
	10, // 10: Ydb.PersQueue.V1.PersQueueService.StreamingWrite:output_type -> Ydb.PersQueue.V1.StreamingWriteServerMessage
	11, // 11: Ydb.PersQueue.V1.PersQueueService.MigrationStreamingRead:output_type -> Ydb.PersQueue.V1.MigrationStreamingReadServerMessage
	12, // 12: Ydb.PersQueue.V1.PersQueueService.GetReadSessionsInfo:output_type -> Ydb.PersQueue.V1.ReadInfoResponse
	13, // 13: Ydb.PersQueue.V1.PersQueueService.DescribeTopic:output_type -> Ydb.PersQueue.V1.DescribeTopicResponse
	14, // 14: Ydb.PersQueue.V1.PersQueueService.DropTopic:output_type -> Ydb.PersQueue.V1.DropTopicResponse
	15, // 15: Ydb.PersQueue.V1.PersQueueService.CreateTopic:output_type -> Ydb.PersQueue.V1.CreateTopicResponse
	16, // 16: Ydb.PersQueue.V1.PersQueueService.AlterTopic:output_type -> Ydb.PersQueue.V1.AlterTopicResponse
	17, // 17: Ydb.PersQueue.V1.PersQueueService.AddReadRule:output_type -> Ydb.PersQueue.V1.AddReadRuleResponse
	18, // 18: Ydb.PersQueue.V1.PersQueueService.RemoveReadRule:output_type -> Ydb.PersQueue.V1.RemoveReadRuleResponse
	19, // 19: Ydb.PersQueue.V1.ClusterDiscoveryService.DiscoverClusters:output_type -> Ydb.PersQueue.ClusterDiscovery.DiscoverClustersResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_persqueue_v1_proto_init() }
func file_ydb_persqueue_v1_proto_init() {
	if File_ydb_persqueue_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_persqueue_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_ydb_persqueue_v1_proto_goTypes,
		DependencyIndexes: file_ydb_persqueue_v1_proto_depIdxs,
	}.Build()
	File_ydb_persqueue_v1_proto = out.File
	file_ydb_persqueue_v1_proto_rawDesc = nil
	file_ydb_persqueue_v1_proto_goTypes = nil
	file_ydb_persqueue_v1_proto_depIdxs = nil
}
