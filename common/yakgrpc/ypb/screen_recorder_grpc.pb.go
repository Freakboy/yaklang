// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: screen_recorder.proto

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
	ScreenRecorderService_IsScrecorderReady_FullMethodName    = "/ypb.ScreenRecorderService/IsScrecorderReady"
	ScreenRecorderService_InstallScrecorder_FullMethodName    = "/ypb.ScreenRecorderService/InstallScrecorder"
	ScreenRecorderService_StartScrecorder_FullMethodName      = "/ypb.ScreenRecorderService/StartScrecorder"
	ScreenRecorderService_QueryScreenRecorders_FullMethodName = "/ypb.ScreenRecorderService/QueryScreenRecorders"
)

// ScreenRecorderServiceClient is the client API for ScreenRecorderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScreenRecorderServiceClient interface {
	IsScrecorderReady(ctx context.Context, in *IsScrecorderReadyRequest, opts ...grpc.CallOption) (*IsScrecorderReadyResponse, error)
	InstallScrecorder(ctx context.Context, in *InstallScrecorderRequest, opts ...grpc.CallOption) (ScreenRecorderService_InstallScrecorderClient, error)
	StartScrecorder(ctx context.Context, in *StartScrecorderRequest, opts ...grpc.CallOption) (ScreenRecorderService_StartScrecorderClient, error)
	QueryScreenRecorders(ctx context.Context, in *QueryScreenRecorderRequest, opts ...grpc.CallOption) (*QueryScreenRecorderResponse, error)
}

type screenRecorderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScreenRecorderServiceClient(cc grpc.ClientConnInterface) ScreenRecorderServiceClient {
	return &screenRecorderServiceClient{cc}
}

func (c *screenRecorderServiceClient) IsScrecorderReady(ctx context.Context, in *IsScrecorderReadyRequest, opts ...grpc.CallOption) (*IsScrecorderReadyResponse, error) {
	out := new(IsScrecorderReadyResponse)
	err := c.cc.Invoke(ctx, ScreenRecorderService_IsScrecorderReady_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *screenRecorderServiceClient) InstallScrecorder(ctx context.Context, in *InstallScrecorderRequest, opts ...grpc.CallOption) (ScreenRecorderService_InstallScrecorderClient, error) {
	stream, err := c.cc.NewStream(ctx, &ScreenRecorderService_ServiceDesc.Streams[0], ScreenRecorderService_InstallScrecorder_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &screenRecorderServiceInstallScrecorderClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ScreenRecorderService_InstallScrecorderClient interface {
	Recv() (*ExecResult, error)
	grpc.ClientStream
}

type screenRecorderServiceInstallScrecorderClient struct {
	grpc.ClientStream
}

func (x *screenRecorderServiceInstallScrecorderClient) Recv() (*ExecResult, error) {
	m := new(ExecResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *screenRecorderServiceClient) StartScrecorder(ctx context.Context, in *StartScrecorderRequest, opts ...grpc.CallOption) (ScreenRecorderService_StartScrecorderClient, error) {
	stream, err := c.cc.NewStream(ctx, &ScreenRecorderService_ServiceDesc.Streams[1], ScreenRecorderService_StartScrecorder_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &screenRecorderServiceStartScrecorderClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ScreenRecorderService_StartScrecorderClient interface {
	Recv() (*ExecResult, error)
	grpc.ClientStream
}

type screenRecorderServiceStartScrecorderClient struct {
	grpc.ClientStream
}

func (x *screenRecorderServiceStartScrecorderClient) Recv() (*ExecResult, error) {
	m := new(ExecResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *screenRecorderServiceClient) QueryScreenRecorders(ctx context.Context, in *QueryScreenRecorderRequest, opts ...grpc.CallOption) (*QueryScreenRecorderResponse, error) {
	out := new(QueryScreenRecorderResponse)
	err := c.cc.Invoke(ctx, ScreenRecorderService_QueryScreenRecorders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScreenRecorderServiceServer is the server API for ScreenRecorderService service.
// All implementations must embed UnimplementedScreenRecorderServiceServer
// for forward compatibility
type ScreenRecorderServiceServer interface {
	IsScrecorderReady(context.Context, *IsScrecorderReadyRequest) (*IsScrecorderReadyResponse, error)
	InstallScrecorder(*InstallScrecorderRequest, ScreenRecorderService_InstallScrecorderServer) error
	StartScrecorder(*StartScrecorderRequest, ScreenRecorderService_StartScrecorderServer) error
	QueryScreenRecorders(context.Context, *QueryScreenRecorderRequest) (*QueryScreenRecorderResponse, error)
	mustEmbedUnimplementedScreenRecorderServiceServer()
}

// UnimplementedScreenRecorderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScreenRecorderServiceServer struct {
}

func (UnimplementedScreenRecorderServiceServer) IsScrecorderReady(context.Context, *IsScrecorderReadyRequest) (*IsScrecorderReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsScrecorderReady not implemented")
}
func (UnimplementedScreenRecorderServiceServer) InstallScrecorder(*InstallScrecorderRequest, ScreenRecorderService_InstallScrecorderServer) error {
	return status.Errorf(codes.Unimplemented, "method InstallScrecorder not implemented")
}
func (UnimplementedScreenRecorderServiceServer) StartScrecorder(*StartScrecorderRequest, ScreenRecorderService_StartScrecorderServer) error {
	return status.Errorf(codes.Unimplemented, "method StartScrecorder not implemented")
}
func (UnimplementedScreenRecorderServiceServer) QueryScreenRecorders(context.Context, *QueryScreenRecorderRequest) (*QueryScreenRecorderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryScreenRecorders not implemented")
}
func (UnimplementedScreenRecorderServiceServer) mustEmbedUnimplementedScreenRecorderServiceServer() {}

// UnsafeScreenRecorderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScreenRecorderServiceServer will
// result in compilation errors.
type UnsafeScreenRecorderServiceServer interface {
	mustEmbedUnimplementedScreenRecorderServiceServer()
}

func RegisterScreenRecorderServiceServer(s grpc.ServiceRegistrar, srv ScreenRecorderServiceServer) {
	s.RegisterService(&ScreenRecorderService_ServiceDesc, srv)
}

func _ScreenRecorderService_IsScrecorderReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsScrecorderReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenRecorderServiceServer).IsScrecorderReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScreenRecorderService_IsScrecorderReady_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenRecorderServiceServer).IsScrecorderReady(ctx, req.(*IsScrecorderReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScreenRecorderService_InstallScrecorder_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InstallScrecorderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ScreenRecorderServiceServer).InstallScrecorder(m, &screenRecorderServiceInstallScrecorderServer{stream})
}

type ScreenRecorderService_InstallScrecorderServer interface {
	Send(*ExecResult) error
	grpc.ServerStream
}

type screenRecorderServiceInstallScrecorderServer struct {
	grpc.ServerStream
}

func (x *screenRecorderServiceInstallScrecorderServer) Send(m *ExecResult) error {
	return x.ServerStream.SendMsg(m)
}

func _ScreenRecorderService_StartScrecorder_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StartScrecorderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ScreenRecorderServiceServer).StartScrecorder(m, &screenRecorderServiceStartScrecorderServer{stream})
}

type ScreenRecorderService_StartScrecorderServer interface {
	Send(*ExecResult) error
	grpc.ServerStream
}

type screenRecorderServiceStartScrecorderServer struct {
	grpc.ServerStream
}

func (x *screenRecorderServiceStartScrecorderServer) Send(m *ExecResult) error {
	return x.ServerStream.SendMsg(m)
}

func _ScreenRecorderService_QueryScreenRecorders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryScreenRecorderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScreenRecorderServiceServer).QueryScreenRecorders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScreenRecorderService_QueryScreenRecorders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScreenRecorderServiceServer).QueryScreenRecorders(ctx, req.(*QueryScreenRecorderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScreenRecorderService_ServiceDesc is the grpc.ServiceDesc for ScreenRecorderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScreenRecorderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.ScreenRecorderService",
	HandlerType: (*ScreenRecorderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsScrecorderReady",
			Handler:    _ScreenRecorderService_IsScrecorderReady_Handler,
		},
		{
			MethodName: "QueryScreenRecorders",
			Handler:    _ScreenRecorderService_QueryScreenRecorders_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InstallScrecorder",
			Handler:       _ScreenRecorderService_InstallScrecorder_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StartScrecorder",
			Handler:       _ScreenRecorderService_StartScrecorder_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "screen_recorder.proto",
}
