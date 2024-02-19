proto: clean protobuf

clean:
	rm -rf draft protos Ydb_*

protobuf:
	docker build -f generate-protobuf.Dockerfile . -t ydb-go-sdk-proto-generator-env
	docker run -v ${PWD}:/project -it ydb-go-sdk-proto-generator-env:latest bash ./generate_proto.sh
