#!/bin/bash
protoc --go_out=./ --go_opt=module=github.com/ydb-platform/ydb-go-genproto --go-grpc_out=./ --go-grpc_opt=module=github.com/ydb-platform/ydb-go-genproto --proto_path=./api/ ./api/*.proto
protoc --go_out=./ --go_opt=module=github.com/ydb-platform/ydb-go-genproto --go-grpc_out=./ --go-grpc_opt=module=github.com/ydb-platform/ydb-go-genproto --proto_path=./api/ ./api/protos/*.proto
protoc --go_out=./ --go_opt=module=github.com/ydb-platform/ydb-go-genproto --go-grpc_out=./ --go-grpc_opt=module=github.com/ydb-platform/ydb-go-genproto --proto_path=./api/ ./api/protos/annotations/*.proto
protoc --go_out=./ --go_opt=module=github.com/ydb-platform/ydb-go-genproto --go-grpc_out=./ --go-grpc_opt=module=github.com/ydb-platform/ydb-go-genproto --proto_path=./api/ ./api/draft/protos/*.proto
protoc --go_out=./ --go_opt=module=github.com/ydb-platform/ydb-go-genproto --go-grpc_out=./ --go-grpc_opt=module=github.com/ydb-platform/ydb-go-genproto --proto_path=./api/ ./api/draft/*.proto
