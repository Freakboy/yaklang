// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: chaos_maker.proto

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
	ChaosMakerService_ImportChaosMakerRules_FullMethodName    = "/ypb.ChaosMakerService/ImportChaosMakerRules"
	ChaosMakerService_QueryChaosMakerRule_FullMethodName      = "/ypb.ChaosMakerService/QueryChaosMakerRule"
	ChaosMakerService_DeleteChaosMakerRuleByID_FullMethodName = "/ypb.ChaosMakerService/DeleteChaosMakerRuleByID"
	ChaosMakerService_ExecuteChaosMakerRule_FullMethodName    = "/ypb.ChaosMakerService/ExecuteChaosMakerRule"
	ChaosMakerService_IsRemoteAddrAvailable_FullMethodName    = "/ypb.ChaosMakerService/IsRemoteAddrAvailable"
)

// ChaosMakerServiceClient is the client API for ChaosMakerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChaosMakerServiceClient interface {
	ImportChaosMakerRules(ctx context.Context, in *ImportChaosMakerRulesRequest, opts ...grpc.CallOption) (*Empty, error)
	QueryChaosMakerRule(ctx context.Context, in *QueryChaosMakerRuleRequest, opts ...grpc.CallOption) (*QueryChaosMakerRuleResponse, error)
	DeleteChaosMakerRuleByID(ctx context.Context, in *DeleteChaosMakerRuleByIDRequest, opts ...grpc.CallOption) (*Empty, error)
	ExecuteChaosMakerRule(ctx context.Context, in *ExecuteChaosMakerRuleRequest, opts ...grpc.CallOption) (ChaosMakerService_ExecuteChaosMakerRuleClient, error)
	IsRemoteAddrAvailable(ctx context.Context, in *IsRemoteAddrAvailableRequest, opts ...grpc.CallOption) (*IsRemoteAddrAvailableResponse, error)
}

type chaosMakerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChaosMakerServiceClient(cc grpc.ClientConnInterface) ChaosMakerServiceClient {
	return &chaosMakerServiceClient{cc}
}

func (c *chaosMakerServiceClient) ImportChaosMakerRules(ctx context.Context, in *ImportChaosMakerRulesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChaosMakerService_ImportChaosMakerRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosMakerServiceClient) QueryChaosMakerRule(ctx context.Context, in *QueryChaosMakerRuleRequest, opts ...grpc.CallOption) (*QueryChaosMakerRuleResponse, error) {
	out := new(QueryChaosMakerRuleResponse)
	err := c.cc.Invoke(ctx, ChaosMakerService_QueryChaosMakerRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosMakerServiceClient) DeleteChaosMakerRuleByID(ctx context.Context, in *DeleteChaosMakerRuleByIDRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChaosMakerService_DeleteChaosMakerRuleByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosMakerServiceClient) ExecuteChaosMakerRule(ctx context.Context, in *ExecuteChaosMakerRuleRequest, opts ...grpc.CallOption) (ChaosMakerService_ExecuteChaosMakerRuleClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChaosMakerService_ServiceDesc.Streams[0], ChaosMakerService_ExecuteChaosMakerRule_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chaosMakerServiceExecuteChaosMakerRuleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChaosMakerService_ExecuteChaosMakerRuleClient interface {
	Recv() (*ExecResult, error)
	grpc.ClientStream
}

type chaosMakerServiceExecuteChaosMakerRuleClient struct {
	grpc.ClientStream
}

func (x *chaosMakerServiceExecuteChaosMakerRuleClient) Recv() (*ExecResult, error) {
	m := new(ExecResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chaosMakerServiceClient) IsRemoteAddrAvailable(ctx context.Context, in *IsRemoteAddrAvailableRequest, opts ...grpc.CallOption) (*IsRemoteAddrAvailableResponse, error) {
	out := new(IsRemoteAddrAvailableResponse)
	err := c.cc.Invoke(ctx, ChaosMakerService_IsRemoteAddrAvailable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChaosMakerServiceServer is the server API for ChaosMakerService service.
// All implementations must embed UnimplementedChaosMakerServiceServer
// for forward compatibility
type ChaosMakerServiceServer interface {
	ImportChaosMakerRules(context.Context, *ImportChaosMakerRulesRequest) (*Empty, error)
	QueryChaosMakerRule(context.Context, *QueryChaosMakerRuleRequest) (*QueryChaosMakerRuleResponse, error)
	DeleteChaosMakerRuleByID(context.Context, *DeleteChaosMakerRuleByIDRequest) (*Empty, error)
	ExecuteChaosMakerRule(*ExecuteChaosMakerRuleRequest, ChaosMakerService_ExecuteChaosMakerRuleServer) error
	IsRemoteAddrAvailable(context.Context, *IsRemoteAddrAvailableRequest) (*IsRemoteAddrAvailableResponse, error)
	mustEmbedUnimplementedChaosMakerServiceServer()
}

// UnimplementedChaosMakerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChaosMakerServiceServer struct {
}

func (UnimplementedChaosMakerServiceServer) ImportChaosMakerRules(context.Context, *ImportChaosMakerRulesRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportChaosMakerRules not implemented")
}
func (UnimplementedChaosMakerServiceServer) QueryChaosMakerRule(context.Context, *QueryChaosMakerRuleRequest) (*QueryChaosMakerRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryChaosMakerRule not implemented")
}
func (UnimplementedChaosMakerServiceServer) DeleteChaosMakerRuleByID(context.Context, *DeleteChaosMakerRuleByIDRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChaosMakerRuleByID not implemented")
}
func (UnimplementedChaosMakerServiceServer) ExecuteChaosMakerRule(*ExecuteChaosMakerRuleRequest, ChaosMakerService_ExecuteChaosMakerRuleServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteChaosMakerRule not implemented")
}
func (UnimplementedChaosMakerServiceServer) IsRemoteAddrAvailable(context.Context, *IsRemoteAddrAvailableRequest) (*IsRemoteAddrAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsRemoteAddrAvailable not implemented")
}
func (UnimplementedChaosMakerServiceServer) mustEmbedUnimplementedChaosMakerServiceServer() {}

// UnsafeChaosMakerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChaosMakerServiceServer will
// result in compilation errors.
type UnsafeChaosMakerServiceServer interface {
	mustEmbedUnimplementedChaosMakerServiceServer()
}

func RegisterChaosMakerServiceServer(s grpc.ServiceRegistrar, srv ChaosMakerServiceServer) {
	s.RegisterService(&ChaosMakerService_ServiceDesc, srv)
}

func _ChaosMakerService_ImportChaosMakerRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportChaosMakerRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosMakerServiceServer).ImportChaosMakerRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaosMakerService_ImportChaosMakerRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosMakerServiceServer).ImportChaosMakerRules(ctx, req.(*ImportChaosMakerRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosMakerService_QueryChaosMakerRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChaosMakerRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosMakerServiceServer).QueryChaosMakerRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaosMakerService_QueryChaosMakerRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosMakerServiceServer).QueryChaosMakerRule(ctx, req.(*QueryChaosMakerRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosMakerService_DeleteChaosMakerRuleByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChaosMakerRuleByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosMakerServiceServer).DeleteChaosMakerRuleByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaosMakerService_DeleteChaosMakerRuleByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosMakerServiceServer).DeleteChaosMakerRuleByID(ctx, req.(*DeleteChaosMakerRuleByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosMakerService_ExecuteChaosMakerRule_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecuteChaosMakerRuleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChaosMakerServiceServer).ExecuteChaosMakerRule(m, &chaosMakerServiceExecuteChaosMakerRuleServer{stream})
}

type ChaosMakerService_ExecuteChaosMakerRuleServer interface {
	Send(*ExecResult) error
	grpc.ServerStream
}

type chaosMakerServiceExecuteChaosMakerRuleServer struct {
	grpc.ServerStream
}

func (x *chaosMakerServiceExecuteChaosMakerRuleServer) Send(m *ExecResult) error {
	return x.ServerStream.SendMsg(m)
}

func _ChaosMakerService_IsRemoteAddrAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsRemoteAddrAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosMakerServiceServer).IsRemoteAddrAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaosMakerService_IsRemoteAddrAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosMakerServiceServer).IsRemoteAddrAvailable(ctx, req.(*IsRemoteAddrAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChaosMakerService_ServiceDesc is the grpc.ServiceDesc for ChaosMakerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChaosMakerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.ChaosMakerService",
	HandlerType: (*ChaosMakerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportChaosMakerRules",
			Handler:    _ChaosMakerService_ImportChaosMakerRules_Handler,
		},
		{
			MethodName: "QueryChaosMakerRule",
			Handler:    _ChaosMakerService_QueryChaosMakerRule_Handler,
		},
		{
			MethodName: "DeleteChaosMakerRuleByID",
			Handler:    _ChaosMakerService_DeleteChaosMakerRuleByID_Handler,
		},
		{
			MethodName: "IsRemoteAddrAvailable",
			Handler:    _ChaosMakerService_IsRemoteAddrAvailable_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteChaosMakerRule",
			Handler:       _ChaosMakerService_ExecuteChaosMakerRule_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chaos_maker.proto",
}
