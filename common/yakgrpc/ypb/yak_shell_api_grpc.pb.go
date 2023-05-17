// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: yak_shell_api.proto

package ypb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	YakShellApi_CreateYaklangShell_FullMethodName = "/ypb.YakShellApi/CreateYaklangShell"
)

// YakShellApiClient is the client API for YakShellApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YakShellApiClient interface {
	// Yaklang Shell
	// 创建一个交互式 Shell
	CreateYaklangShell(ctx context.Context, opts ...grpc.CallOption) (YakShellApi_CreateYaklangShellClient, error)
}

type yakShellApiClient struct {
	cc grpc.ClientConnInterface
}

func NewYakShellApiClient(cc grpc.ClientConnInterface) YakShellApiClient {
	return &yakShellApiClient{cc}
}

func (c *yakShellApiClient) CreateYaklangShell(ctx context.Context, opts ...grpc.CallOption) (YakShellApi_CreateYaklangShellClient, error) {
	stream, err := c.cc.NewStream(ctx, &YakShellApi_ServiceDesc.Streams[0], YakShellApi_CreateYaklangShell_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &yakShellApiCreateYaklangShellClient{stream}
	return x, nil
}

type YakShellApi_CreateYaklangShellClient interface {
	Send(*YaklangShellRequest) error
	Recv() (*YaklangShellResponse, error)
	grpc.ClientStream
}

type yakShellApiCreateYaklangShellClient struct {
	grpc.ClientStream
}

func (x *yakShellApiCreateYaklangShellClient) Send(m *YaklangShellRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *yakShellApiCreateYaklangShellClient) Recv() (*YaklangShellResponse, error) {
	m := new(YaklangShellResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// YakShellApiServer is the server API for YakShellApi service.
// All implementations must embed UnimplementedYakShellApiServer
// for forward compatibility
type YakShellApiServer interface {
	// Yaklang Shell
	// 创建一个交互式 Shell
	CreateYaklangShell(YakShellApi_CreateYaklangShellServer) error
	mustEmbedUnimplementedYakShellApiServer()
}

// UnimplementedYakShellApiServer must be embedded to have forward compatible implementations.
type UnimplementedYakShellApiServer struct {
}

func (UnimplementedYakShellApiServer) CreateYaklangShell(YakShellApi_CreateYaklangShellServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateYaklangShell not implemented")
}
func (UnimplementedYakShellApiServer) mustEmbedUnimplementedYakShellApiServer() {}

// UnsafeYakShellApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YakShellApiServer will
// result in compilation errors.
type UnsafeYakShellApiServer interface {
	mustEmbedUnimplementedYakShellApiServer()
}

func RegisterYakShellApiServer(s grpc.ServiceRegistrar, srv YakShellApiServer) {
	s.RegisterService(&YakShellApi_ServiceDesc, srv)
}

func _YakShellApi_CreateYaklangShell_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(YakShellApiServer).CreateYaklangShell(&yakShellApiCreateYaklangShellServer{stream})
}

type YakShellApi_CreateYaklangShellServer interface {
	Send(*YaklangShellResponse) error
	Recv() (*YaklangShellRequest, error)
	grpc.ServerStream
}

type yakShellApiCreateYaklangShellServer struct {
	grpc.ServerStream
}

func (x *yakShellApiCreateYaklangShellServer) Send(m *YaklangShellResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *yakShellApiCreateYaklangShellServer) Recv() (*YaklangShellRequest, error) {
	m := new(YaklangShellRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// YakShellApi_ServiceDesc is the grpc.ServiceDesc for YakShellApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YakShellApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.YakShellApi",
	HandlerType: (*YakShellApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateYaklangShell",
			Handler:       _YakShellApi_CreateYaklangShell_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "yak_shell_api.proto",
}
