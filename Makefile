proto: protobuf-3

protobuf-3:
	docker build -f generate-protobuf.Dockerfile . -t ydb-go-sdk-proto-generator-env-3 --build-arg PROTOC_VER=25.3 --build-arg PTOCOC_GEN_GO=1.32.0 --build-arg PROTOC_GEN_GO_GRPC=1.3.0
	docker run -v ${PWD}:/github.com/ydb-platform/ydb-go-genproto/volumes -it ydb-go-sdk-proto-generator-env-3:latest bash generate_proto.sh
