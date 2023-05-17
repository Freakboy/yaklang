// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: online_api.proto

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
	OnlineApi_GetOnlineProfile_FullMethodName                  = "/ypb.OnlineApi/GetOnlineProfile"
	OnlineApi_SetOnlineProfile_FullMethodName                  = "/ypb.OnlineApi/SetOnlineProfile"
	OnlineApi_DownloadOnlinePluginById_FullMethodName          = "/ypb.OnlineApi/DownloadOnlinePluginById"
	OnlineApi_DownloadOnlinePluginByIds_FullMethodName         = "/ypb.OnlineApi/DownloadOnlinePluginByIds"
	OnlineApi_DownloadOnlinePluginAll_FullMethodName           = "/ypb.OnlineApi/DownloadOnlinePluginAll"
	OnlineApi_DeletePluginByUserID_FullMethodName              = "/ypb.OnlineApi/DeletePluginByUserID"
	OnlineApi_DeleteAllLocalPlugins_FullMethodName             = "/ypb.OnlineApi/DeleteAllLocalPlugins"
	OnlineApi_GetYakScriptTagsAndType_FullMethodName           = "/ypb.OnlineApi/GetYakScriptTagsAndType"
	OnlineApi_DeleteLocalPluginsByWhere_FullMethodName         = "/ypb.OnlineApi/DeleteLocalPluginsByWhere"
	OnlineApi_DownloadOnlinePluginByScriptNames_FullMethodName = "/ypb.OnlineApi/DownloadOnlinePluginByScriptNames"
)

// OnlineApiClient is the client API for OnlineApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OnlineApiClient interface {
	// Online
	GetOnlineProfile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OnlineProfile, error)
	SetOnlineProfile(ctx context.Context, in *OnlineProfile, opts ...grpc.CallOption) (*Empty, error)
	DownloadOnlinePluginById(ctx context.Context, in *DownloadOnlinePluginByIdRequest, opts ...grpc.CallOption) (*Empty, error)
	DownloadOnlinePluginByIds(ctx context.Context, in *DownloadOnlinePluginByIdsRequest, opts ...grpc.CallOption) (*Empty, error)
	DownloadOnlinePluginAll(ctx context.Context, in *DownloadOnlinePluginByTokenRequest, opts ...grpc.CallOption) (OnlineApi_DownloadOnlinePluginAllClient, error)
	DeletePluginByUserID(ctx context.Context, in *DeletePluginByUserIDRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteAllLocalPlugins(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	GetYakScriptTagsAndType(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetYakScriptTagsAndTypeResponse, error)
	DeleteLocalPluginsByWhere(ctx context.Context, in *DeleteLocalPluginsByWhereRequest, opts ...grpc.CallOption) (*Empty, error)
	DownloadOnlinePluginByScriptNames(ctx context.Context, in *DownloadOnlinePluginByScriptNamesRequest, opts ...grpc.CallOption) (*DownloadOnlinePluginByScriptNamesResponse, error)
}

type onlineApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOnlineApiClient(cc grpc.ClientConnInterface) OnlineApiClient {
	return &onlineApiClient{cc}
}

func (c *onlineApiClient) GetOnlineProfile(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OnlineProfile, error) {
	out := new(OnlineProfile)
	err := c.cc.Invoke(ctx, OnlineApi_GetOnlineProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineApiClient) SetOnlineProfile(ctx context.Context, in *OnlineProfile, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, OnlineApi_SetOnlineProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineApiClient) DownloadOnlinePluginById(ctx context.Context, in *DownloadOnlinePluginByIdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, OnlineApi_DownloadOnlinePluginById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineApiClient) DownloadOnlinePluginByIds(ctx context.Context, in *DownloadOnlinePluginByIdsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, OnlineApi_DownloadOnlinePluginByIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineApiClient) DownloadOnlinePluginAll(ctx context.Context, in *DownloadOnlinePluginByTokenRequest, opts ...grpc.CallOption) (OnlineApi_DownloadOnlinePluginAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &OnlineApi_ServiceDesc.Streams[0], OnlineApi_DownloadOnlinePluginAll_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &onlineApiDownloadOnlinePluginAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OnlineApi_DownloadOnlinePluginAllClient interface {
	Recv() (*DownloadOnlinePluginProgress, error)
	grpc.ClientStream
}

type onlineApiDownloadOnlinePluginAllClient struct {
	grpc.ClientStream
}

func (x *onlineApiDownloadOnlinePluginAllClient) Recv() (*DownloadOnlinePluginProgress, error) {
	m := new(DownloadOnlinePluginProgress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *onlineApiClient) DeletePluginByUserID(ctx context.Context, in *DeletePluginByUserIDRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, OnlineApi_DeletePluginByUserID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineApiClient) DeleteAllLocalPlugins(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, OnlineApi_DeleteAllLocalPlugins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineApiClient) GetYakScriptTagsAndType(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetYakScriptTagsAndTypeResponse, error) {
	out := new(GetYakScriptTagsAndTypeResponse)
	err := c.cc.Invoke(ctx, OnlineApi_GetYakScriptTagsAndType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineApiClient) DeleteLocalPluginsByWhere(ctx context.Context, in *DeleteLocalPluginsByWhereRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, OnlineApi_DeleteLocalPluginsByWhere_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineApiClient) DownloadOnlinePluginByScriptNames(ctx context.Context, in *DownloadOnlinePluginByScriptNamesRequest, opts ...grpc.CallOption) (*DownloadOnlinePluginByScriptNamesResponse, error) {
	out := new(DownloadOnlinePluginByScriptNamesResponse)
	err := c.cc.Invoke(ctx, OnlineApi_DownloadOnlinePluginByScriptNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnlineApiServer is the server API for OnlineApi service.
// All implementations must embed UnimplementedOnlineApiServer
// for forward compatibility
type OnlineApiServer interface {
	// Online
	GetOnlineProfile(context.Context, *Empty) (*OnlineProfile, error)
	SetOnlineProfile(context.Context, *OnlineProfile) (*Empty, error)
	DownloadOnlinePluginById(context.Context, *DownloadOnlinePluginByIdRequest) (*Empty, error)
	DownloadOnlinePluginByIds(context.Context, *DownloadOnlinePluginByIdsRequest) (*Empty, error)
	DownloadOnlinePluginAll(*DownloadOnlinePluginByTokenRequest, OnlineApi_DownloadOnlinePluginAllServer) error
	DeletePluginByUserID(context.Context, *DeletePluginByUserIDRequest) (*Empty, error)
	DeleteAllLocalPlugins(context.Context, *Empty) (*Empty, error)
	GetYakScriptTagsAndType(context.Context, *Empty) (*GetYakScriptTagsAndTypeResponse, error)
	DeleteLocalPluginsByWhere(context.Context, *DeleteLocalPluginsByWhereRequest) (*Empty, error)
	DownloadOnlinePluginByScriptNames(context.Context, *DownloadOnlinePluginByScriptNamesRequest) (*DownloadOnlinePluginByScriptNamesResponse, error)
	mustEmbedUnimplementedOnlineApiServer()
}

// UnimplementedOnlineApiServer must be embedded to have forward compatible implementations.
type UnimplementedOnlineApiServer struct {
}

func (UnimplementedOnlineApiServer) GetOnlineProfile(context.Context, *Empty) (*OnlineProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineProfile not implemented")
}
func (UnimplementedOnlineApiServer) SetOnlineProfile(context.Context, *OnlineProfile) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOnlineProfile not implemented")
}
func (UnimplementedOnlineApiServer) DownloadOnlinePluginById(context.Context, *DownloadOnlinePluginByIdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadOnlinePluginById not implemented")
}
func (UnimplementedOnlineApiServer) DownloadOnlinePluginByIds(context.Context, *DownloadOnlinePluginByIdsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadOnlinePluginByIds not implemented")
}
func (UnimplementedOnlineApiServer) DownloadOnlinePluginAll(*DownloadOnlinePluginByTokenRequest, OnlineApi_DownloadOnlinePluginAllServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadOnlinePluginAll not implemented")
}
func (UnimplementedOnlineApiServer) DeletePluginByUserID(context.Context, *DeletePluginByUserIDRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePluginByUserID not implemented")
}
func (UnimplementedOnlineApiServer) DeleteAllLocalPlugins(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllLocalPlugins not implemented")
}
func (UnimplementedOnlineApiServer) GetYakScriptTagsAndType(context.Context, *Empty) (*GetYakScriptTagsAndTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetYakScriptTagsAndType not implemented")
}
func (UnimplementedOnlineApiServer) DeleteLocalPluginsByWhere(context.Context, *DeleteLocalPluginsByWhereRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLocalPluginsByWhere not implemented")
}
func (UnimplementedOnlineApiServer) DownloadOnlinePluginByScriptNames(context.Context, *DownloadOnlinePluginByScriptNamesRequest) (*DownloadOnlinePluginByScriptNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadOnlinePluginByScriptNames not implemented")
}
func (UnimplementedOnlineApiServer) mustEmbedUnimplementedOnlineApiServer() {}

// UnsafeOnlineApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnlineApiServer will
// result in compilation errors.
type UnsafeOnlineApiServer interface {
	mustEmbedUnimplementedOnlineApiServer()
}

func RegisterOnlineApiServer(s grpc.ServiceRegistrar, srv OnlineApiServer) {
	s.RegisterService(&OnlineApi_ServiceDesc, srv)
}

func _OnlineApi_GetOnlineProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineApiServer).GetOnlineProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlineApi_GetOnlineProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineApiServer).GetOnlineProfile(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineApi_SetOnlineProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineApiServer).SetOnlineProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlineApi_SetOnlineProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineApiServer).SetOnlineProfile(ctx, req.(*OnlineProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineApi_DownloadOnlinePluginById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadOnlinePluginByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineApiServer).DownloadOnlinePluginById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlineApi_DownloadOnlinePluginById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineApiServer).DownloadOnlinePluginById(ctx, req.(*DownloadOnlinePluginByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineApi_DownloadOnlinePluginByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadOnlinePluginByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineApiServer).DownloadOnlinePluginByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlineApi_DownloadOnlinePluginByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineApiServer).DownloadOnlinePluginByIds(ctx, req.(*DownloadOnlinePluginByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineApi_DownloadOnlinePluginAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadOnlinePluginByTokenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OnlineApiServer).DownloadOnlinePluginAll(m, &onlineApiDownloadOnlinePluginAllServer{stream})
}

type OnlineApi_DownloadOnlinePluginAllServer interface {
	Send(*DownloadOnlinePluginProgress) error
	grpc.ServerStream
}

type onlineApiDownloadOnlinePluginAllServer struct {
	grpc.ServerStream
}

func (x *onlineApiDownloadOnlinePluginAllServer) Send(m *DownloadOnlinePluginProgress) error {
	return x.ServerStream.SendMsg(m)
}

func _OnlineApi_DeletePluginByUserID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePluginByUserIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineApiServer).DeletePluginByUserID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlineApi_DeletePluginByUserID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineApiServer).DeletePluginByUserID(ctx, req.(*DeletePluginByUserIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineApi_DeleteAllLocalPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineApiServer).DeleteAllLocalPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlineApi_DeleteAllLocalPlugins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineApiServer).DeleteAllLocalPlugins(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineApi_GetYakScriptTagsAndType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineApiServer).GetYakScriptTagsAndType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlineApi_GetYakScriptTagsAndType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineApiServer).GetYakScriptTagsAndType(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineApi_DeleteLocalPluginsByWhere_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLocalPluginsByWhereRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineApiServer).DeleteLocalPluginsByWhere(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlineApi_DeleteLocalPluginsByWhere_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineApiServer).DeleteLocalPluginsByWhere(ctx, req.(*DeleteLocalPluginsByWhereRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnlineApi_DownloadOnlinePluginByScriptNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadOnlinePluginByScriptNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineApiServer).DownloadOnlinePluginByScriptNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnlineApi_DownloadOnlinePluginByScriptNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineApiServer).DownloadOnlinePluginByScriptNames(ctx, req.(*DownloadOnlinePluginByScriptNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OnlineApi_ServiceDesc is the grpc.ServiceDesc for OnlineApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OnlineApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.OnlineApi",
	HandlerType: (*OnlineApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOnlineProfile",
			Handler:    _OnlineApi_GetOnlineProfile_Handler,
		},
		{
			MethodName: "SetOnlineProfile",
			Handler:    _OnlineApi_SetOnlineProfile_Handler,
		},
		{
			MethodName: "DownloadOnlinePluginById",
			Handler:    _OnlineApi_DownloadOnlinePluginById_Handler,
		},
		{
			MethodName: "DownloadOnlinePluginByIds",
			Handler:    _OnlineApi_DownloadOnlinePluginByIds_Handler,
		},
		{
			MethodName: "DeletePluginByUserID",
			Handler:    _OnlineApi_DeletePluginByUserID_Handler,
		},
		{
			MethodName: "DeleteAllLocalPlugins",
			Handler:    _OnlineApi_DeleteAllLocalPlugins_Handler,
		},
		{
			MethodName: "GetYakScriptTagsAndType",
			Handler:    _OnlineApi_GetYakScriptTagsAndType_Handler,
		},
		{
			MethodName: "DeleteLocalPluginsByWhere",
			Handler:    _OnlineApi_DeleteLocalPluginsByWhere_Handler,
		},
		{
			MethodName: "DownloadOnlinePluginByScriptNames",
			Handler:    _OnlineApi_DownloadOnlinePluginByScriptNames_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadOnlinePluginAll",
			Handler:       _OnlineApi_DownloadOnlinePluginAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "online_api.proto",
}
