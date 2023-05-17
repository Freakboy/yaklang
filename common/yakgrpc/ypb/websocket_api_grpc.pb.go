// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: websocket_api.proto

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
	WebSocketApi_CreateWebsocketFuzzer_FullMethodName                      = "/ypb.WebSocketApi/CreateWebsocketFuzzer"
	WebSocketApi_QueryWebsocketFlowByHTTPFlowWebsocketHash_FullMethodName  = "/ypb.WebSocketApi/QueryWebsocketFlowByHTTPFlowWebsocketHash"
	WebSocketApi_DeleteWebsocketFlowByHTTPFlowWebsocketHash_FullMethodName = "/ypb.WebSocketApi/DeleteWebsocketFlowByHTTPFlowWebsocketHash"
	WebSocketApi_DeleteWebsocketFlowAll_FullMethodName                     = "/ypb.WebSocketApi/DeleteWebsocketFlowAll"
)

// WebSocketApiClient is the client API for WebSocketApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebSocketApiClient interface {
	// WebSocket
	CreateWebsocketFuzzer(ctx context.Context, opts ...grpc.CallOption) (WebSocketApi_CreateWebsocketFuzzerClient, error)
	QueryWebsocketFlowByHTTPFlowWebsocketHash(ctx context.Context, in *QueryWebsocketFlowByHTTPFlowWebsocketHashRequest, opts ...grpc.CallOption) (*WebsocketFlows, error)
	DeleteWebsocketFlowByHTTPFlowWebsocketHash(ctx context.Context, in *DeleteWebsocketFlowByHTTPFlowWebsocketHashRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteWebsocketFlowAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type webSocketApiClient struct {
	cc grpc.ClientConnInterface
}

func NewWebSocketApiClient(cc grpc.ClientConnInterface) WebSocketApiClient {
	return &webSocketApiClient{cc}
}

func (c *webSocketApiClient) CreateWebsocketFuzzer(ctx context.Context, opts ...grpc.CallOption) (WebSocketApi_CreateWebsocketFuzzerClient, error) {
	stream, err := c.cc.NewStream(ctx, &WebSocketApi_ServiceDesc.Streams[0], WebSocketApi_CreateWebsocketFuzzer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &webSocketApiCreateWebsocketFuzzerClient{stream}
	return x, nil
}

type WebSocketApi_CreateWebsocketFuzzerClient interface {
	Send(*ClientWebsocketRequest) error
	Recv() (*ClientWebsocketResponse, error)
	grpc.ClientStream
}

type webSocketApiCreateWebsocketFuzzerClient struct {
	grpc.ClientStream
}

func (x *webSocketApiCreateWebsocketFuzzerClient) Send(m *ClientWebsocketRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *webSocketApiCreateWebsocketFuzzerClient) Recv() (*ClientWebsocketResponse, error) {
	m := new(ClientWebsocketResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *webSocketApiClient) QueryWebsocketFlowByHTTPFlowWebsocketHash(ctx context.Context, in *QueryWebsocketFlowByHTTPFlowWebsocketHashRequest, opts ...grpc.CallOption) (*WebsocketFlows, error) {
	out := new(WebsocketFlows)
	err := c.cc.Invoke(ctx, WebSocketApi_QueryWebsocketFlowByHTTPFlowWebsocketHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webSocketApiClient) DeleteWebsocketFlowByHTTPFlowWebsocketHash(ctx context.Context, in *DeleteWebsocketFlowByHTTPFlowWebsocketHashRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, WebSocketApi_DeleteWebsocketFlowByHTTPFlowWebsocketHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webSocketApiClient) DeleteWebsocketFlowAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, WebSocketApi_DeleteWebsocketFlowAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebSocketApiServer is the server API for WebSocketApi service.
// All implementations must embed UnimplementedWebSocketApiServer
// for forward compatibility
type WebSocketApiServer interface {
	// WebSocket
	CreateWebsocketFuzzer(WebSocketApi_CreateWebsocketFuzzerServer) error
	QueryWebsocketFlowByHTTPFlowWebsocketHash(context.Context, *QueryWebsocketFlowByHTTPFlowWebsocketHashRequest) (*WebsocketFlows, error)
	DeleteWebsocketFlowByHTTPFlowWebsocketHash(context.Context, *DeleteWebsocketFlowByHTTPFlowWebsocketHashRequest) (*Empty, error)
	DeleteWebsocketFlowAll(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedWebSocketApiServer()
}

// UnimplementedWebSocketApiServer must be embedded to have forward compatible implementations.
type UnimplementedWebSocketApiServer struct {
}

func (UnimplementedWebSocketApiServer) CreateWebsocketFuzzer(WebSocketApi_CreateWebsocketFuzzerServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateWebsocketFuzzer not implemented")
}
func (UnimplementedWebSocketApiServer) QueryWebsocketFlowByHTTPFlowWebsocketHash(context.Context, *QueryWebsocketFlowByHTTPFlowWebsocketHashRequest) (*WebsocketFlows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWebsocketFlowByHTTPFlowWebsocketHash not implemented")
}
func (UnimplementedWebSocketApiServer) DeleteWebsocketFlowByHTTPFlowWebsocketHash(context.Context, *DeleteWebsocketFlowByHTTPFlowWebsocketHashRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWebsocketFlowByHTTPFlowWebsocketHash not implemented")
}
func (UnimplementedWebSocketApiServer) DeleteWebsocketFlowAll(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWebsocketFlowAll not implemented")
}
func (UnimplementedWebSocketApiServer) mustEmbedUnimplementedWebSocketApiServer() {}

// UnsafeWebSocketApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebSocketApiServer will
// result in compilation errors.
type UnsafeWebSocketApiServer interface {
	mustEmbedUnimplementedWebSocketApiServer()
}

func RegisterWebSocketApiServer(s grpc.ServiceRegistrar, srv WebSocketApiServer) {
	s.RegisterService(&WebSocketApi_ServiceDesc, srv)
}

func _WebSocketApi_CreateWebsocketFuzzer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WebSocketApiServer).CreateWebsocketFuzzer(&webSocketApiCreateWebsocketFuzzerServer{stream})
}

type WebSocketApi_CreateWebsocketFuzzerServer interface {
	Send(*ClientWebsocketResponse) error
	Recv() (*ClientWebsocketRequest, error)
	grpc.ServerStream
}

type webSocketApiCreateWebsocketFuzzerServer struct {
	grpc.ServerStream
}

func (x *webSocketApiCreateWebsocketFuzzerServer) Send(m *ClientWebsocketResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *webSocketApiCreateWebsocketFuzzerServer) Recv() (*ClientWebsocketRequest, error) {
	m := new(ClientWebsocketRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WebSocketApi_QueryWebsocketFlowByHTTPFlowWebsocketHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWebsocketFlowByHTTPFlowWebsocketHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebSocketApiServer).QueryWebsocketFlowByHTTPFlowWebsocketHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebSocketApi_QueryWebsocketFlowByHTTPFlowWebsocketHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebSocketApiServer).QueryWebsocketFlowByHTTPFlowWebsocketHash(ctx, req.(*QueryWebsocketFlowByHTTPFlowWebsocketHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebSocketApi_DeleteWebsocketFlowByHTTPFlowWebsocketHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWebsocketFlowByHTTPFlowWebsocketHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebSocketApiServer).DeleteWebsocketFlowByHTTPFlowWebsocketHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebSocketApi_DeleteWebsocketFlowByHTTPFlowWebsocketHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebSocketApiServer).DeleteWebsocketFlowByHTTPFlowWebsocketHash(ctx, req.(*DeleteWebsocketFlowByHTTPFlowWebsocketHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebSocketApi_DeleteWebsocketFlowAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebSocketApiServer).DeleteWebsocketFlowAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebSocketApi_DeleteWebsocketFlowAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebSocketApiServer).DeleteWebsocketFlowAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// WebSocketApi_ServiceDesc is the grpc.ServiceDesc for WebSocketApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebSocketApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.WebSocketApi",
	HandlerType: (*WebSocketApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryWebsocketFlowByHTTPFlowWebsocketHash",
			Handler:    _WebSocketApi_QueryWebsocketFlowByHTTPFlowWebsocketHash_Handler,
		},
		{
			MethodName: "DeleteWebsocketFlowByHTTPFlowWebsocketHash",
			Handler:    _WebSocketApi_DeleteWebsocketFlowByHTTPFlowWebsocketHash_Handler,
		},
		{
			MethodName: "DeleteWebsocketFlowAll",
			Handler:    _WebSocketApi_DeleteWebsocketFlowAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateWebsocketFuzzer",
			Handler:       _WebSocketApi_CreateWebsocketFuzzer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "websocket_api.proto",
}
