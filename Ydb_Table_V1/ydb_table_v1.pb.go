// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: ydb_table_v1.proto

package Ydb_Table_V1

import (
	Ydb_Table "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Table"
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

var File_ydb_table_v1_proto protoreflect.FileDescriptor

var file_ydb_table_v1_proto_rawDesc = []byte{
	0x0a, 0x12, 0x79, 0x64, 0x62, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x56, 0x31, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x79, 0x64, 0x62, 0x5f, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4, 0x0e, 0x0a, 0x0c, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65,
	0x12, 0x1b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4b, 0x65, 0x65,
	0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x44, 0x72, 0x6f,
	0x70, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x44, 0x72, 0x6f, 0x70, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x0a, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x1c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09,
	0x43, 0x6f, 0x70, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x6f, 0x70, 0x79, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x12, 0x1c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43,
	0x6f, 0x70, 0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x70,
	0x79, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x0c, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12,
	0x1e, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x52, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x1f, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x22, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x59,
	0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5b, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x22, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b,
	0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x22, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x12, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b,
	0x0a, 0x10, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x13, 0x52,
	0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x52,
	0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x67, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x59, 0x64, 0x62, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0f, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x2e,
	0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x59, 0x64, 0x62,
	0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x0a, 0x42, 0x75,
	0x6c, 0x6b, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x12, 0x1c, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x22, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x59, 0x64, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x49, 0x0a, 0x11, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x79, 0x64, 0x62, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x59, 0x64, 0x62, 0x5f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_ydb_table_v1_proto_goTypes = []interface{}{
	(*Ydb_Table.CreateSessionRequest)(nil),            // 0: Ydb.Table.CreateSessionRequest
	(*Ydb_Table.DeleteSessionRequest)(nil),            // 1: Ydb.Table.DeleteSessionRequest
	(*Ydb_Table.KeepAliveRequest)(nil),                // 2: Ydb.Table.KeepAliveRequest
	(*Ydb_Table.CreateTableRequest)(nil),              // 3: Ydb.Table.CreateTableRequest
	(*Ydb_Table.DropTableRequest)(nil),                // 4: Ydb.Table.DropTableRequest
	(*Ydb_Table.AlterTableRequest)(nil),               // 5: Ydb.Table.AlterTableRequest
	(*Ydb_Table.CopyTableRequest)(nil),                // 6: Ydb.Table.CopyTableRequest
	(*Ydb_Table.CopyTablesRequest)(nil),               // 7: Ydb.Table.CopyTablesRequest
	(*Ydb_Table.RenameTablesRequest)(nil),             // 8: Ydb.Table.RenameTablesRequest
	(*Ydb_Table.DescribeTableRequest)(nil),            // 9: Ydb.Table.DescribeTableRequest
	(*Ydb_Table.ExplainDataQueryRequest)(nil),         // 10: Ydb.Table.ExplainDataQueryRequest
	(*Ydb_Table.PrepareDataQueryRequest)(nil),         // 11: Ydb.Table.PrepareDataQueryRequest
	(*Ydb_Table.ExecuteDataQueryRequest)(nil),         // 12: Ydb.Table.ExecuteDataQueryRequest
	(*Ydb_Table.ExecuteSchemeQueryRequest)(nil),       // 13: Ydb.Table.ExecuteSchemeQueryRequest
	(*Ydb_Table.BeginTransactionRequest)(nil),         // 14: Ydb.Table.BeginTransactionRequest
	(*Ydb_Table.CommitTransactionRequest)(nil),        // 15: Ydb.Table.CommitTransactionRequest
	(*Ydb_Table.RollbackTransactionRequest)(nil),      // 16: Ydb.Table.RollbackTransactionRequest
	(*Ydb_Table.DescribeTableOptionsRequest)(nil),     // 17: Ydb.Table.DescribeTableOptionsRequest
	(*Ydb_Table.ReadTableRequest)(nil),                // 18: Ydb.Table.ReadTableRequest
	(*Ydb_Table.BulkUpsertRequest)(nil),               // 19: Ydb.Table.BulkUpsertRequest
	(*Ydb_Table.ExecuteScanQueryRequest)(nil),         // 20: Ydb.Table.ExecuteScanQueryRequest
	(*Ydb_Table.CreateSessionResponse)(nil),           // 21: Ydb.Table.CreateSessionResponse
	(*Ydb_Table.DeleteSessionResponse)(nil),           // 22: Ydb.Table.DeleteSessionResponse
	(*Ydb_Table.KeepAliveResponse)(nil),               // 23: Ydb.Table.KeepAliveResponse
	(*Ydb_Table.CreateTableResponse)(nil),             // 24: Ydb.Table.CreateTableResponse
	(*Ydb_Table.DropTableResponse)(nil),               // 25: Ydb.Table.DropTableResponse
	(*Ydb_Table.AlterTableResponse)(nil),              // 26: Ydb.Table.AlterTableResponse
	(*Ydb_Table.CopyTableResponse)(nil),               // 27: Ydb.Table.CopyTableResponse
	(*Ydb_Table.CopyTablesResponse)(nil),              // 28: Ydb.Table.CopyTablesResponse
	(*Ydb_Table.RenameTablesResponse)(nil),            // 29: Ydb.Table.RenameTablesResponse
	(*Ydb_Table.DescribeTableResponse)(nil),           // 30: Ydb.Table.DescribeTableResponse
	(*Ydb_Table.ExplainDataQueryResponse)(nil),        // 31: Ydb.Table.ExplainDataQueryResponse
	(*Ydb_Table.PrepareDataQueryResponse)(nil),        // 32: Ydb.Table.PrepareDataQueryResponse
	(*Ydb_Table.ExecuteDataQueryResponse)(nil),        // 33: Ydb.Table.ExecuteDataQueryResponse
	(*Ydb_Table.ExecuteSchemeQueryResponse)(nil),      // 34: Ydb.Table.ExecuteSchemeQueryResponse
	(*Ydb_Table.BeginTransactionResponse)(nil),        // 35: Ydb.Table.BeginTransactionResponse
	(*Ydb_Table.CommitTransactionResponse)(nil),       // 36: Ydb.Table.CommitTransactionResponse
	(*Ydb_Table.RollbackTransactionResponse)(nil),     // 37: Ydb.Table.RollbackTransactionResponse
	(*Ydb_Table.DescribeTableOptionsResponse)(nil),    // 38: Ydb.Table.DescribeTableOptionsResponse
	(*Ydb_Table.ReadTableResponse)(nil),               // 39: Ydb.Table.ReadTableResponse
	(*Ydb_Table.BulkUpsertResponse)(nil),              // 40: Ydb.Table.BulkUpsertResponse
	(*Ydb_Table.ExecuteScanQueryPartialResponse)(nil), // 41: Ydb.Table.ExecuteScanQueryPartialResponse
}
var file_ydb_table_v1_proto_depIdxs = []int32{
	0,  // 0: Ydb.Table.V1.TableService.CreateSession:input_type -> Ydb.Table.CreateSessionRequest
	1,  // 1: Ydb.Table.V1.TableService.DeleteSession:input_type -> Ydb.Table.DeleteSessionRequest
	2,  // 2: Ydb.Table.V1.TableService.KeepAlive:input_type -> Ydb.Table.KeepAliveRequest
	3,  // 3: Ydb.Table.V1.TableService.CreateTable:input_type -> Ydb.Table.CreateTableRequest
	4,  // 4: Ydb.Table.V1.TableService.DropTable:input_type -> Ydb.Table.DropTableRequest
	5,  // 5: Ydb.Table.V1.TableService.AlterTable:input_type -> Ydb.Table.AlterTableRequest
	6,  // 6: Ydb.Table.V1.TableService.CopyTable:input_type -> Ydb.Table.CopyTableRequest
	7,  // 7: Ydb.Table.V1.TableService.CopyTables:input_type -> Ydb.Table.CopyTablesRequest
	8,  // 8: Ydb.Table.V1.TableService.RenameTables:input_type -> Ydb.Table.RenameTablesRequest
	9,  // 9: Ydb.Table.V1.TableService.DescribeTable:input_type -> Ydb.Table.DescribeTableRequest
	10, // 10: Ydb.Table.V1.TableService.ExplainDataQuery:input_type -> Ydb.Table.ExplainDataQueryRequest
	11, // 11: Ydb.Table.V1.TableService.PrepareDataQuery:input_type -> Ydb.Table.PrepareDataQueryRequest
	12, // 12: Ydb.Table.V1.TableService.ExecuteDataQuery:input_type -> Ydb.Table.ExecuteDataQueryRequest
	13, // 13: Ydb.Table.V1.TableService.ExecuteSchemeQuery:input_type -> Ydb.Table.ExecuteSchemeQueryRequest
	14, // 14: Ydb.Table.V1.TableService.BeginTransaction:input_type -> Ydb.Table.BeginTransactionRequest
	15, // 15: Ydb.Table.V1.TableService.CommitTransaction:input_type -> Ydb.Table.CommitTransactionRequest
	16, // 16: Ydb.Table.V1.TableService.RollbackTransaction:input_type -> Ydb.Table.RollbackTransactionRequest
	17, // 17: Ydb.Table.V1.TableService.DescribeTableOptions:input_type -> Ydb.Table.DescribeTableOptionsRequest
	18, // 18: Ydb.Table.V1.TableService.StreamReadTable:input_type -> Ydb.Table.ReadTableRequest
	19, // 19: Ydb.Table.V1.TableService.BulkUpsert:input_type -> Ydb.Table.BulkUpsertRequest
	20, // 20: Ydb.Table.V1.TableService.StreamExecuteScanQuery:input_type -> Ydb.Table.ExecuteScanQueryRequest
	21, // 21: Ydb.Table.V1.TableService.CreateSession:output_type -> Ydb.Table.CreateSessionResponse
	22, // 22: Ydb.Table.V1.TableService.DeleteSession:output_type -> Ydb.Table.DeleteSessionResponse
	23, // 23: Ydb.Table.V1.TableService.KeepAlive:output_type -> Ydb.Table.KeepAliveResponse
	24, // 24: Ydb.Table.V1.TableService.CreateTable:output_type -> Ydb.Table.CreateTableResponse
	25, // 25: Ydb.Table.V1.TableService.DropTable:output_type -> Ydb.Table.DropTableResponse
	26, // 26: Ydb.Table.V1.TableService.AlterTable:output_type -> Ydb.Table.AlterTableResponse
	27, // 27: Ydb.Table.V1.TableService.CopyTable:output_type -> Ydb.Table.CopyTableResponse
	28, // 28: Ydb.Table.V1.TableService.CopyTables:output_type -> Ydb.Table.CopyTablesResponse
	29, // 29: Ydb.Table.V1.TableService.RenameTables:output_type -> Ydb.Table.RenameTablesResponse
	30, // 30: Ydb.Table.V1.TableService.DescribeTable:output_type -> Ydb.Table.DescribeTableResponse
	31, // 31: Ydb.Table.V1.TableService.ExplainDataQuery:output_type -> Ydb.Table.ExplainDataQueryResponse
	32, // 32: Ydb.Table.V1.TableService.PrepareDataQuery:output_type -> Ydb.Table.PrepareDataQueryResponse
	33, // 33: Ydb.Table.V1.TableService.ExecuteDataQuery:output_type -> Ydb.Table.ExecuteDataQueryResponse
	34, // 34: Ydb.Table.V1.TableService.ExecuteSchemeQuery:output_type -> Ydb.Table.ExecuteSchemeQueryResponse
	35, // 35: Ydb.Table.V1.TableService.BeginTransaction:output_type -> Ydb.Table.BeginTransactionResponse
	36, // 36: Ydb.Table.V1.TableService.CommitTransaction:output_type -> Ydb.Table.CommitTransactionResponse
	37, // 37: Ydb.Table.V1.TableService.RollbackTransaction:output_type -> Ydb.Table.RollbackTransactionResponse
	38, // 38: Ydb.Table.V1.TableService.DescribeTableOptions:output_type -> Ydb.Table.DescribeTableOptionsResponse
	39, // 39: Ydb.Table.V1.TableService.StreamReadTable:output_type -> Ydb.Table.ReadTableResponse
	40, // 40: Ydb.Table.V1.TableService.BulkUpsert:output_type -> Ydb.Table.BulkUpsertResponse
	41, // 41: Ydb.Table.V1.TableService.StreamExecuteScanQuery:output_type -> Ydb.Table.ExecuteScanQueryPartialResponse
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ydb_table_v1_proto_init() }
func file_ydb_table_v1_proto_init() {
	if File_ydb_table_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_table_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ydb_table_v1_proto_goTypes,
		DependencyIndexes: file_ydb_table_v1_proto_depIdxs,
	}.Build()
	File_ydb_table_v1_proto = out.File
	file_ydb_table_v1_proto_rawDesc = nil
	file_ydb_table_v1_proto_goTypes = nil
	file_ydb_table_v1_proto_depIdxs = nil
}
