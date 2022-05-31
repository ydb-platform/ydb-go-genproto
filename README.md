# ydb-go-genproto

`ydb-go-genproto` includes code generation from YDB protos

## Installation `protoc-gen-go` and `protoc-gen-go-grpc`

The protocol buffer compiler requires a plugin to generate Go code. Install it using Go 1.16 or higher by running:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
This will install a `protoc-gen-go` and `protoc-gen-go-grpc` binaries in `$GOBIN`. Set the `$GOBIN` environment variable to change the installation location. It must be in your `$PATH` for the protocol buffer compiler to find it.

## Re-generation `*.pb.go` code

1. First you must get latest YDB proto files which includes into this project as git submodule:
```
cd api
git checkout master
git pull origin master
cd ../
```
2. Run generation `*.pb.go` code from command:
```
make proto
```
This command re-generate `*.pb.go` files

