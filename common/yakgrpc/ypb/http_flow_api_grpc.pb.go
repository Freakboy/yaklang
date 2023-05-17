// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: http_flow_api.proto

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
	HTTPFlowApi_GetHTTPFlowByHash_FullMethodName           = "/ypb.HTTPFlowApi/GetHTTPFlowByHash"
	HTTPFlowApi_GetHTTPFlowById_FullMethodName             = "/ypb.HTTPFlowApi/GetHTTPFlowById"
	HTTPFlowApi_QueryHTTPFlowByIds_FullMethodName          = "/ypb.HTTPFlowApi/QueryHTTPFlowByIds"
	HTTPFlowApi_GetHTTPFlowByIds_FullMethodName            = "/ypb.HTTPFlowApi/GetHTTPFlowByIds"
	HTTPFlowApi_QueryHTTPFlows_FullMethodName              = "/ypb.HTTPFlowApi/QueryHTTPFlows"
	HTTPFlowApi_DeleteHTTPFlows_FullMethodName             = "/ypb.HTTPFlowApi/DeleteHTTPFlows"
	HTTPFlowApi_SetTagForHTTPFlow_FullMethodName           = "/ypb.HTTPFlowApi/SetTagForHTTPFlow"
	HTTPFlowApi_QueryHTTPFlowsIds_FullMethodName           = "/ypb.HTTPFlowApi/QueryHTTPFlowsIds"
	HTTPFlowApi_HTTPFlowsFieldGroup_FullMethodName         = "/ypb.HTTPFlowApi/HTTPFlowsFieldGroup"
	HTTPFlowApi_GetRequestBodyByHTTPFlowID_FullMethodName  = "/ypb.HTTPFlowApi/GetRequestBodyByHTTPFlowID"
	HTTPFlowApi_GetResponseBodyByHTTPFlowID_FullMethodName = "/ypb.HTTPFlowApi/GetResponseBodyByHTTPFlowID"
	HTTPFlowApi_GetHTTPPacketBody_FullMethodName           = "/ypb.HTTPFlowApi/GetHTTPPacketBody"
)

// HTTPFlowApiClient is the client API for HTTPFlowApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HTTPFlowApiClient interface {
	GetHTTPFlowByHash(ctx context.Context, in *GetHTTPFlowByHashRequest, opts ...grpc.CallOption) (*HTTPFlow, error)
	GetHTTPFlowById(ctx context.Context, in *GetHTTPFlowByIdRequest, opts ...grpc.CallOption) (*HTTPFlow, error)
	QueryHTTPFlowByIds(ctx context.Context, in *GetHTTPFlowByIdsRequest, opts ...grpc.CallOption) (*HTTPFlows, error)
	GetHTTPFlowByIds(ctx context.Context, in *GetHTTPFlowByIdsRequest, opts ...grpc.CallOption) (*HTTPFlows, error)
	QueryHTTPFlows(ctx context.Context, in *QueryHTTPFlowRequest, opts ...grpc.CallOption) (*QueryHTTPFlowResponse, error)
	DeleteHTTPFlows(ctx context.Context, in *DeleteHTTPFlowRequest, opts ...grpc.CallOption) (*Empty, error)
	SetTagForHTTPFlow(ctx context.Context, in *SetTagForHTTPFlowRequest, opts ...grpc.CallOption) (*Empty, error)
	QueryHTTPFlowsIds(ctx context.Context, in *QueryHTTPFlowsIdsRequest, opts ...grpc.CallOption) (*QueryHTTPFlowsIdsResponse, error)
	HTTPFlowsFieldGroup(ctx context.Context, in *HTTPFlowsFieldGroupRequest, opts ...grpc.CallOption) (*HTTPFlowsFieldGroupResponse, error)
	// Response Body 的魔法操作
	GetRequestBodyByHTTPFlowID(ctx context.Context, in *DownloadBodyByHTTPFlowIDRequest, opts ...grpc.CallOption) (*Bytes, error)
	GetResponseBodyByHTTPFlowID(ctx context.Context, in *DownloadBodyByHTTPFlowIDRequest, opts ...grpc.CallOption) (*Bytes, error)
	GetHTTPPacketBody(ctx context.Context, in *GetHTTPPacketBodyRequest, opts ...grpc.CallOption) (*Bytes, error)
}

type hTTPFlowApiClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPFlowApiClient(cc grpc.ClientConnInterface) HTTPFlowApiClient {
	return &hTTPFlowApiClient{cc}
}

func (c *hTTPFlowApiClient) GetHTTPFlowByHash(ctx context.Context, in *GetHTTPFlowByHashRequest, opts ...grpc.CallOption) (*HTTPFlow, error) {
	out := new(HTTPFlow)
	err := c.cc.Invoke(ctx, HTTPFlowApi_GetHTTPFlowByHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) GetHTTPFlowById(ctx context.Context, in *GetHTTPFlowByIdRequest, opts ...grpc.CallOption) (*HTTPFlow, error) {
	out := new(HTTPFlow)
	err := c.cc.Invoke(ctx, HTTPFlowApi_GetHTTPFlowById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) QueryHTTPFlowByIds(ctx context.Context, in *GetHTTPFlowByIdsRequest, opts ...grpc.CallOption) (*HTTPFlows, error) {
	out := new(HTTPFlows)
	err := c.cc.Invoke(ctx, HTTPFlowApi_QueryHTTPFlowByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) GetHTTPFlowByIds(ctx context.Context, in *GetHTTPFlowByIdsRequest, opts ...grpc.CallOption) (*HTTPFlows, error) {
	out := new(HTTPFlows)
	err := c.cc.Invoke(ctx, HTTPFlowApi_GetHTTPFlowByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) QueryHTTPFlows(ctx context.Context, in *QueryHTTPFlowRequest, opts ...grpc.CallOption) (*QueryHTTPFlowResponse, error) {
	out := new(QueryHTTPFlowResponse)
	err := c.cc.Invoke(ctx, HTTPFlowApi_QueryHTTPFlows_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) DeleteHTTPFlows(ctx context.Context, in *DeleteHTTPFlowRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, HTTPFlowApi_DeleteHTTPFlows_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) SetTagForHTTPFlow(ctx context.Context, in *SetTagForHTTPFlowRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, HTTPFlowApi_SetTagForHTTPFlow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) QueryHTTPFlowsIds(ctx context.Context, in *QueryHTTPFlowsIdsRequest, opts ...grpc.CallOption) (*QueryHTTPFlowsIdsResponse, error) {
	out := new(QueryHTTPFlowsIdsResponse)
	err := c.cc.Invoke(ctx, HTTPFlowApi_QueryHTTPFlowsIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) HTTPFlowsFieldGroup(ctx context.Context, in *HTTPFlowsFieldGroupRequest, opts ...grpc.CallOption) (*HTTPFlowsFieldGroupResponse, error) {
	out := new(HTTPFlowsFieldGroupResponse)
	err := c.cc.Invoke(ctx, HTTPFlowApi_HTTPFlowsFieldGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) GetRequestBodyByHTTPFlowID(ctx context.Context, in *DownloadBodyByHTTPFlowIDRequest, opts ...grpc.CallOption) (*Bytes, error) {
	out := new(Bytes)
	err := c.cc.Invoke(ctx, HTTPFlowApi_GetRequestBodyByHTTPFlowID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) GetResponseBodyByHTTPFlowID(ctx context.Context, in *DownloadBodyByHTTPFlowIDRequest, opts ...grpc.CallOption) (*Bytes, error) {
	out := new(Bytes)
	err := c.cc.Invoke(ctx, HTTPFlowApi_GetResponseBodyByHTTPFlowID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hTTPFlowApiClient) GetHTTPPacketBody(ctx context.Context, in *GetHTTPPacketBodyRequest, opts ...grpc.CallOption) (*Bytes, error) {
	out := new(Bytes)
	err := c.cc.Invoke(ctx, HTTPFlowApi_GetHTTPPacketBody_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HTTPFlowApiServer is the server API for HTTPFlowApi service.
// All implementations must embed UnimplementedHTTPFlowApiServer
// for forward compatibility
type HTTPFlowApiServer interface {
	GetHTTPFlowByHash(context.Context, *GetHTTPFlowByHashRequest) (*HTTPFlow, error)
	GetHTTPFlowById(context.Context, *GetHTTPFlowByIdRequest) (*HTTPFlow, error)
	QueryHTTPFlowByIds(context.Context, *GetHTTPFlowByIdsRequest) (*HTTPFlows, error)
	GetHTTPFlowByIds(context.Context, *GetHTTPFlowByIdsRequest) (*HTTPFlows, error)
	QueryHTTPFlows(context.Context, *QueryHTTPFlowRequest) (*QueryHTTPFlowResponse, error)
	DeleteHTTPFlows(context.Context, *DeleteHTTPFlowRequest) (*Empty, error)
	SetTagForHTTPFlow(context.Context, *SetTagForHTTPFlowRequest) (*Empty, error)
	QueryHTTPFlowsIds(context.Context, *QueryHTTPFlowsIdsRequest) (*QueryHTTPFlowsIdsResponse, error)
	HTTPFlowsFieldGroup(context.Context, *HTTPFlowsFieldGroupRequest) (*HTTPFlowsFieldGroupResponse, error)
	// Response Body 的魔法操作
	GetRequestBodyByHTTPFlowID(context.Context, *DownloadBodyByHTTPFlowIDRequest) (*Bytes, error)
	GetResponseBodyByHTTPFlowID(context.Context, *DownloadBodyByHTTPFlowIDRequest) (*Bytes, error)
	GetHTTPPacketBody(context.Context, *GetHTTPPacketBodyRequest) (*Bytes, error)
	mustEmbedUnimplementedHTTPFlowApiServer()
}

// UnimplementedHTTPFlowApiServer must be embedded to have forward compatible implementations.
type UnimplementedHTTPFlowApiServer struct {
}

func (UnimplementedHTTPFlowApiServer) GetHTTPFlowByHash(context.Context, *GetHTTPFlowByHashRequest) (*HTTPFlow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHTTPFlowByHash not implemented")
}
func (UnimplementedHTTPFlowApiServer) GetHTTPFlowById(context.Context, *GetHTTPFlowByIdRequest) (*HTTPFlow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHTTPFlowById not implemented")
}
func (UnimplementedHTTPFlowApiServer) QueryHTTPFlowByIds(context.Context, *GetHTTPFlowByIdsRequest) (*HTTPFlows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryHTTPFlowByIds not implemented")
}
func (UnimplementedHTTPFlowApiServer) GetHTTPFlowByIds(context.Context, *GetHTTPFlowByIdsRequest) (*HTTPFlows, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHTTPFlowByIds not implemented")
}
func (UnimplementedHTTPFlowApiServer) QueryHTTPFlows(context.Context, *QueryHTTPFlowRequest) (*QueryHTTPFlowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryHTTPFlows not implemented")
}
func (UnimplementedHTTPFlowApiServer) DeleteHTTPFlows(context.Context, *DeleteHTTPFlowRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHTTPFlows not implemented")
}
func (UnimplementedHTTPFlowApiServer) SetTagForHTTPFlow(context.Context, *SetTagForHTTPFlowRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTagForHTTPFlow not implemented")
}
func (UnimplementedHTTPFlowApiServer) QueryHTTPFlowsIds(context.Context, *QueryHTTPFlowsIdsRequest) (*QueryHTTPFlowsIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryHTTPFlowsIds not implemented")
}
func (UnimplementedHTTPFlowApiServer) HTTPFlowsFieldGroup(context.Context, *HTTPFlowsFieldGroupRequest) (*HTTPFlowsFieldGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTTPFlowsFieldGroup not implemented")
}
func (UnimplementedHTTPFlowApiServer) GetRequestBodyByHTTPFlowID(context.Context, *DownloadBodyByHTTPFlowIDRequest) (*Bytes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestBodyByHTTPFlowID not implemented")
}
func (UnimplementedHTTPFlowApiServer) GetResponseBodyByHTTPFlowID(context.Context, *DownloadBodyByHTTPFlowIDRequest) (*Bytes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResponseBodyByHTTPFlowID not implemented")
}
func (UnimplementedHTTPFlowApiServer) GetHTTPPacketBody(context.Context, *GetHTTPPacketBodyRequest) (*Bytes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHTTPPacketBody not implemented")
}
func (UnimplementedHTTPFlowApiServer) mustEmbedUnimplementedHTTPFlowApiServer() {}

// UnsafeHTTPFlowApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HTTPFlowApiServer will
// result in compilation errors.
type UnsafeHTTPFlowApiServer interface {
	mustEmbedUnimplementedHTTPFlowApiServer()
}

func RegisterHTTPFlowApiServer(s grpc.ServiceRegistrar, srv HTTPFlowApiServer) {
	s.RegisterService(&HTTPFlowApi_ServiceDesc, srv)
}

func _HTTPFlowApi_GetHTTPFlowByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHTTPFlowByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).GetHTTPFlowByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_GetHTTPFlowByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).GetHTTPFlowByHash(ctx, req.(*GetHTTPFlowByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_GetHTTPFlowById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHTTPFlowByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).GetHTTPFlowById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_GetHTTPFlowById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).GetHTTPFlowById(ctx, req.(*GetHTTPFlowByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_QueryHTTPFlowByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHTTPFlowByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).QueryHTTPFlowByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_QueryHTTPFlowByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).QueryHTTPFlowByIds(ctx, req.(*GetHTTPFlowByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_GetHTTPFlowByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHTTPFlowByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).GetHTTPFlowByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_GetHTTPFlowByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).GetHTTPFlowByIds(ctx, req.(*GetHTTPFlowByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_QueryHTTPFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHTTPFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).QueryHTTPFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_QueryHTTPFlows_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).QueryHTTPFlows(ctx, req.(*QueryHTTPFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_DeleteHTTPFlows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHTTPFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).DeleteHTTPFlows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_DeleteHTTPFlows_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).DeleteHTTPFlows(ctx, req.(*DeleteHTTPFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_SetTagForHTTPFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTagForHTTPFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).SetTagForHTTPFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_SetTagForHTTPFlow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).SetTagForHTTPFlow(ctx, req.(*SetTagForHTTPFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_QueryHTTPFlowsIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHTTPFlowsIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).QueryHTTPFlowsIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_QueryHTTPFlowsIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).QueryHTTPFlowsIds(ctx, req.(*QueryHTTPFlowsIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_HTTPFlowsFieldGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPFlowsFieldGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).HTTPFlowsFieldGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_HTTPFlowsFieldGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).HTTPFlowsFieldGroup(ctx, req.(*HTTPFlowsFieldGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_GetRequestBodyByHTTPFlowID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadBodyByHTTPFlowIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).GetRequestBodyByHTTPFlowID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_GetRequestBodyByHTTPFlowID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).GetRequestBodyByHTTPFlowID(ctx, req.(*DownloadBodyByHTTPFlowIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_GetResponseBodyByHTTPFlowID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadBodyByHTTPFlowIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).GetResponseBodyByHTTPFlowID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_GetResponseBodyByHTTPFlowID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).GetResponseBodyByHTTPFlowID(ctx, req.(*DownloadBodyByHTTPFlowIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HTTPFlowApi_GetHTTPPacketBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHTTPPacketBodyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPFlowApiServer).GetHTTPPacketBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HTTPFlowApi_GetHTTPPacketBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPFlowApiServer).GetHTTPPacketBody(ctx, req.(*GetHTTPPacketBodyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HTTPFlowApi_ServiceDesc is the grpc.ServiceDesc for HTTPFlowApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HTTPFlowApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.HTTPFlowApi",
	HandlerType: (*HTTPFlowApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHTTPFlowByHash",
			Handler:    _HTTPFlowApi_GetHTTPFlowByHash_Handler,
		},
		{
			MethodName: "GetHTTPFlowById",
			Handler:    _HTTPFlowApi_GetHTTPFlowById_Handler,
		},
		{
			MethodName: "QueryHTTPFlowByIds",
			Handler:    _HTTPFlowApi_QueryHTTPFlowByIds_Handler,
		},
		{
			MethodName: "GetHTTPFlowByIds",
			Handler:    _HTTPFlowApi_GetHTTPFlowByIds_Handler,
		},
		{
			MethodName: "QueryHTTPFlows",
			Handler:    _HTTPFlowApi_QueryHTTPFlows_Handler,
		},
		{
			MethodName: "DeleteHTTPFlows",
			Handler:    _HTTPFlowApi_DeleteHTTPFlows_Handler,
		},
		{
			MethodName: "SetTagForHTTPFlow",
			Handler:    _HTTPFlowApi_SetTagForHTTPFlow_Handler,
		},
		{
			MethodName: "QueryHTTPFlowsIds",
			Handler:    _HTTPFlowApi_QueryHTTPFlowsIds_Handler,
		},
		{
			MethodName: "HTTPFlowsFieldGroup",
			Handler:    _HTTPFlowApi_HTTPFlowsFieldGroup_Handler,
		},
		{
			MethodName: "GetRequestBodyByHTTPFlowID",
			Handler:    _HTTPFlowApi_GetRequestBodyByHTTPFlowID_Handler,
		},
		{
			MethodName: "GetResponseBodyByHTTPFlowID",
			Handler:    _HTTPFlowApi_GetResponseBodyByHTTPFlowID_Handler,
		},
		{
			MethodName: "GetHTTPPacketBody",
			Handler:    _HTTPFlowApi_GetHTTPPacketBody_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "http_flow_api.proto",
}
