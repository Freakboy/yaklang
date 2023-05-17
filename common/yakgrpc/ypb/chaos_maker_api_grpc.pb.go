// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: chaos_maker_api.proto

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
	ChaosMakerApi_ImportChaosMakerRules_FullMethodName    = "/ypb.ChaosMakerApi/ImportChaosMakerRules"
	ChaosMakerApi_QueryChaosMakerRule_FullMethodName      = "/ypb.ChaosMakerApi/QueryChaosMakerRule"
	ChaosMakerApi_DeleteChaosMakerRuleByID_FullMethodName = "/ypb.ChaosMakerApi/DeleteChaosMakerRuleByID"
	ChaosMakerApi_ExecuteChaosMakerRule_FullMethodName    = "/ypb.ChaosMakerApi/ExecuteChaosMakerRule"
	ChaosMakerApi_IsRemoteAddrAvailable_FullMethodName    = "/ypb.ChaosMakerApi/IsRemoteAddrAvailable"
)

// ChaosMakerApiClient is the client API for ChaosMakerApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChaosMakerApiClient interface {
	ImportChaosMakerRules(ctx context.Context, in *ImportChaosMakerRulesRequest, opts ...grpc.CallOption) (*Empty, error)
	QueryChaosMakerRule(ctx context.Context, in *QueryChaosMakerRuleRequest, opts ...grpc.CallOption) (*QueryChaosMakerRuleResponse, error)
	DeleteChaosMakerRuleByID(ctx context.Context, in *DeleteChaosMakerRuleByIDRequest, opts ...grpc.CallOption) (*Empty, error)
	ExecuteChaosMakerRule(ctx context.Context, in *ExecuteChaosMakerRuleRequest, opts ...grpc.CallOption) (ChaosMakerApi_ExecuteChaosMakerRuleClient, error)
	IsRemoteAddrAvailable(ctx context.Context, in *IsRemoteAddrAvailableRequest, opts ...grpc.CallOption) (*IsRemoteAddrAvailableResponse, error)
}

type chaosMakerApiClient struct {
	cc grpc.ClientConnInterface
}

func NewChaosMakerApiClient(cc grpc.ClientConnInterface) ChaosMakerApiClient {
	return &chaosMakerApiClient{cc}
}

func (c *chaosMakerApiClient) ImportChaosMakerRules(ctx context.Context, in *ImportChaosMakerRulesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChaosMakerApi_ImportChaosMakerRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosMakerApiClient) QueryChaosMakerRule(ctx context.Context, in *QueryChaosMakerRuleRequest, opts ...grpc.CallOption) (*QueryChaosMakerRuleResponse, error) {
	out := new(QueryChaosMakerRuleResponse)
	err := c.cc.Invoke(ctx, ChaosMakerApi_QueryChaosMakerRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosMakerApiClient) DeleteChaosMakerRuleByID(ctx context.Context, in *DeleteChaosMakerRuleByIDRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChaosMakerApi_DeleteChaosMakerRuleByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaosMakerApiClient) ExecuteChaosMakerRule(ctx context.Context, in *ExecuteChaosMakerRuleRequest, opts ...grpc.CallOption) (ChaosMakerApi_ExecuteChaosMakerRuleClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChaosMakerApi_ServiceDesc.Streams[0], ChaosMakerApi_ExecuteChaosMakerRule_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chaosMakerApiExecuteChaosMakerRuleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChaosMakerApi_ExecuteChaosMakerRuleClient interface {
	Recv() (*ExecResult, error)
	grpc.ClientStream
}

type chaosMakerApiExecuteChaosMakerRuleClient struct {
	grpc.ClientStream
}

func (x *chaosMakerApiExecuteChaosMakerRuleClient) Recv() (*ExecResult, error) {
	m := new(ExecResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chaosMakerApiClient) IsRemoteAddrAvailable(ctx context.Context, in *IsRemoteAddrAvailableRequest, opts ...grpc.CallOption) (*IsRemoteAddrAvailableResponse, error) {
	out := new(IsRemoteAddrAvailableResponse)
	err := c.cc.Invoke(ctx, ChaosMakerApi_IsRemoteAddrAvailable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChaosMakerApiServer is the server API for ChaosMakerApi service.
// All implementations must embed UnimplementedChaosMakerApiServer
// for forward compatibility
type ChaosMakerApiServer interface {
	ImportChaosMakerRules(context.Context, *ImportChaosMakerRulesRequest) (*Empty, error)
	QueryChaosMakerRule(context.Context, *QueryChaosMakerRuleRequest) (*QueryChaosMakerRuleResponse, error)
	DeleteChaosMakerRuleByID(context.Context, *DeleteChaosMakerRuleByIDRequest) (*Empty, error)
	ExecuteChaosMakerRule(*ExecuteChaosMakerRuleRequest, ChaosMakerApi_ExecuteChaosMakerRuleServer) error
	IsRemoteAddrAvailable(context.Context, *IsRemoteAddrAvailableRequest) (*IsRemoteAddrAvailableResponse, error)
	mustEmbedUnimplementedChaosMakerApiServer()
}

// UnimplementedChaosMakerApiServer must be embedded to have forward compatible implementations.
type UnimplementedChaosMakerApiServer struct {
}

func (UnimplementedChaosMakerApiServer) ImportChaosMakerRules(context.Context, *ImportChaosMakerRulesRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportChaosMakerRules not implemented")
}
func (UnimplementedChaosMakerApiServer) QueryChaosMakerRule(context.Context, *QueryChaosMakerRuleRequest) (*QueryChaosMakerRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryChaosMakerRule not implemented")
}
func (UnimplementedChaosMakerApiServer) DeleteChaosMakerRuleByID(context.Context, *DeleteChaosMakerRuleByIDRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChaosMakerRuleByID not implemented")
}
func (UnimplementedChaosMakerApiServer) ExecuteChaosMakerRule(*ExecuteChaosMakerRuleRequest, ChaosMakerApi_ExecuteChaosMakerRuleServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteChaosMakerRule not implemented")
}
func (UnimplementedChaosMakerApiServer) IsRemoteAddrAvailable(context.Context, *IsRemoteAddrAvailableRequest) (*IsRemoteAddrAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsRemoteAddrAvailable not implemented")
}
func (UnimplementedChaosMakerApiServer) mustEmbedUnimplementedChaosMakerApiServer() {}

// UnsafeChaosMakerApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChaosMakerApiServer will
// result in compilation errors.
type UnsafeChaosMakerApiServer interface {
	mustEmbedUnimplementedChaosMakerApiServer()
}

func RegisterChaosMakerApiServer(s grpc.ServiceRegistrar, srv ChaosMakerApiServer) {
	s.RegisterService(&ChaosMakerApi_ServiceDesc, srv)
}

func _ChaosMakerApi_ImportChaosMakerRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportChaosMakerRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosMakerApiServer).ImportChaosMakerRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaosMakerApi_ImportChaosMakerRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosMakerApiServer).ImportChaosMakerRules(ctx, req.(*ImportChaosMakerRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosMakerApi_QueryChaosMakerRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChaosMakerRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosMakerApiServer).QueryChaosMakerRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaosMakerApi_QueryChaosMakerRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosMakerApiServer).QueryChaosMakerRule(ctx, req.(*QueryChaosMakerRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosMakerApi_DeleteChaosMakerRuleByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChaosMakerRuleByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosMakerApiServer).DeleteChaosMakerRuleByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaosMakerApi_DeleteChaosMakerRuleByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosMakerApiServer).DeleteChaosMakerRuleByID(ctx, req.(*DeleteChaosMakerRuleByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaosMakerApi_ExecuteChaosMakerRule_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExecuteChaosMakerRuleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChaosMakerApiServer).ExecuteChaosMakerRule(m, &chaosMakerApiExecuteChaosMakerRuleServer{stream})
}

type ChaosMakerApi_ExecuteChaosMakerRuleServer interface {
	Send(*ExecResult) error
	grpc.ServerStream
}

type chaosMakerApiExecuteChaosMakerRuleServer struct {
	grpc.ServerStream
}

func (x *chaosMakerApiExecuteChaosMakerRuleServer) Send(m *ExecResult) error {
	return x.ServerStream.SendMsg(m)
}

func _ChaosMakerApi_IsRemoteAddrAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsRemoteAddrAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaosMakerApiServer).IsRemoteAddrAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChaosMakerApi_IsRemoteAddrAvailable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaosMakerApiServer).IsRemoteAddrAvailable(ctx, req.(*IsRemoteAddrAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChaosMakerApi_ServiceDesc is the grpc.ServiceDesc for ChaosMakerApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChaosMakerApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.ChaosMakerApi",
	HandlerType: (*ChaosMakerApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportChaosMakerRules",
			Handler:    _ChaosMakerApi_ImportChaosMakerRules_Handler,
		},
		{
			MethodName: "QueryChaosMakerRule",
			Handler:    _ChaosMakerApi_QueryChaosMakerRule_Handler,
		},
		{
			MethodName: "DeleteChaosMakerRuleByID",
			Handler:    _ChaosMakerApi_DeleteChaosMakerRuleByID_Handler,
		},
		{
			MethodName: "IsRemoteAddrAvailable",
			Handler:    _ChaosMakerApi_IsRemoteAddrAvailable_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteChaosMakerRule",
			Handler:       _ChaosMakerApi_ExecuteChaosMakerRule_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chaos_maker_api.proto",
}
