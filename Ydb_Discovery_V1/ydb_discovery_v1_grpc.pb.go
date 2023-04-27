// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: ydb_discovery_v1.proto

package Ydb_Discovery_V1

import (
	context "context"
	Ydb_Discovery "github.com/ydb-platform/ydb-go-genproto/protos/Ydb_Discovery"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DiscoveryService_ListEndpoints_FullMethodName = "/Ydb.Discovery.V1.DiscoveryService/ListEndpoints"
	DiscoveryService_WhoAmI_FullMethodName        = "/Ydb.Discovery.V1.DiscoveryService/WhoAmI"
)

// DiscoveryServiceClient is the client API for DiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveryServiceClient interface {
	ListEndpoints(ctx context.Context, in *Ydb_Discovery.ListEndpointsRequest, opts ...grpc.CallOption) (*Ydb_Discovery.ListEndpointsResponse, error)
	WhoAmI(ctx context.Context, in *Ydb_Discovery.WhoAmIRequest, opts ...grpc.CallOption) (*Ydb_Discovery.WhoAmIResponse, error)
}

type discoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryServiceClient(cc grpc.ClientConnInterface) DiscoveryServiceClient {
	return &discoveryServiceClient{cc}
}

func (c *discoveryServiceClient) ListEndpoints(ctx context.Context, in *Ydb_Discovery.ListEndpointsRequest, opts ...grpc.CallOption) (*Ydb_Discovery.ListEndpointsResponse, error) {
	out := new(Ydb_Discovery.ListEndpointsResponse)
	err := c.cc.Invoke(ctx, DiscoveryService_ListEndpoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) WhoAmI(ctx context.Context, in *Ydb_Discovery.WhoAmIRequest, opts ...grpc.CallOption) (*Ydb_Discovery.WhoAmIResponse, error) {
	out := new(Ydb_Discovery.WhoAmIResponse)
	err := c.cc.Invoke(ctx, DiscoveryService_WhoAmI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServiceServer is the server API for DiscoveryService service.
// All implementations must embed UnimplementedDiscoveryServiceServer
// for forward compatibility
type DiscoveryServiceServer interface {
	ListEndpoints(context.Context, *Ydb_Discovery.ListEndpointsRequest) (*Ydb_Discovery.ListEndpointsResponse, error)
	WhoAmI(context.Context, *Ydb_Discovery.WhoAmIRequest) (*Ydb_Discovery.WhoAmIResponse, error)
	mustEmbedUnimplementedDiscoveryServiceServer()
}

// UnimplementedDiscoveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServiceServer struct {
}

func (UnimplementedDiscoveryServiceServer) ListEndpoints(context.Context, *Ydb_Discovery.ListEndpointsRequest) (*Ydb_Discovery.ListEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEndpoints not implemented")
}
func (UnimplementedDiscoveryServiceServer) WhoAmI(context.Context, *Ydb_Discovery.WhoAmIRequest) (*Ydb_Discovery.WhoAmIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhoAmI not implemented")
}
func (UnimplementedDiscoveryServiceServer) mustEmbedUnimplementedDiscoveryServiceServer() {}

// UnsafeDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveryServiceServer will
// result in compilation errors.
type UnsafeDiscoveryServiceServer interface {
	mustEmbedUnimplementedDiscoveryServiceServer()
}

func RegisterDiscoveryServiceServer(s grpc.ServiceRegistrar, srv DiscoveryServiceServer) {
	s.RegisterService(&DiscoveryService_ServiceDesc, srv)
}

func _DiscoveryService_ListEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Discovery.ListEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).ListEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryService_ListEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).ListEndpoints(ctx, req.(*Ydb_Discovery.ListEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ydb_Discovery.WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DiscoveryService_WhoAmI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).WhoAmI(ctx, req.(*Ydb_Discovery.WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DiscoveryService_ServiceDesc is the grpc.ServiceDesc for DiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ydb.Discovery.V1.DiscoveryService",
	HandlerType: (*DiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEndpoints",
			Handler:    _DiscoveryService_ListEndpoints_Handler,
		},
		{
			MethodName: "WhoAmI",
			Handler:    _DiscoveryService_WhoAmI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ydb_discovery_v1.proto",
}
