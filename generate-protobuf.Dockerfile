FROM golang:1.23

ARG PROTOC_VER=30.2
ARG PTOCOC_GEN_GO=1.36.6
ARG PROTOC_GEN_GO_GRPC=1.5.1

WORKDIR /project

RUN apt-get update && apt-get install -y unzip && apt-get clean

RUN \
    wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-linux-x86_64.zip && \
    unzip -o protoc-*.zip -d /usr/local bin/protoc && \
    unzip -o protoc-*.zip -d /usr/local 'include/*' && \
    rm -f protoc-*.zip

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PTOCOC_GEN_GO}
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC}

