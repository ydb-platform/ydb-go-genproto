proto: clean protobuf

clean:
	rm -rf draft protos Ydb_*

protobuf:
	docker build -f generate-protobuf.Dockerfile . -t ydb-go-sdk-proto-generator-env
	docker run -v ${PWD}:/github.com/ydb-platform/ydb-go-genproto/volumes -it ydb-go-sdk-proto-generator-env:latest bash /github.com/ydb-platform/ydb-go-genproto/volumes/generate_proto.sh
