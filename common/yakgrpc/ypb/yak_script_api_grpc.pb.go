// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: yak_script_api.proto

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
	YakScriptApi_QueryYakScript_FullMethodName                    = "/ypb.YakScriptApi/QueryYakScript"
	YakScriptApi_QueryYakScriptByYakScriptName_FullMethodName     = "/ypb.YakScriptApi/QueryYakScriptByYakScriptName"
	YakScriptApi_SaveYakScript_FullMethodName                     = "/ypb.YakScriptApi/SaveYakScript"
	YakScriptApi_DeleteYakScript_FullMethodName                   = "/ypb.YakScriptApi/DeleteYakScript"
	YakScriptApi_GetYakScriptById_FullMethodName                  = "/ypb.YakScriptApi/GetYakScriptById"
	YakScriptApi_GetYakScriptByName_FullMethodName                = "/ypb.YakScriptApi/GetYakScriptByName"
	YakScriptApi_GetYakScriptByOnlineID_FullMethodName            = "/ypb.YakScriptApi/GetYakScriptByOnlineID"
	YakScriptApi_IgnoreYakScript_FullMethodName                   = "/ypb.YakScriptApi/IgnoreYakScript"
	YakScriptApi_UnIgnoreYakScript_FullMethodName                 = "/ypb.YakScriptApi/UnIgnoreYakScript"
	YakScriptApi_ExportYakScript_FullMethodName                   = "/ypb.YakScriptApi/ExportYakScript"
	YakScriptApi_GetYakScriptTags_FullMethodName                  = "/ypb.YakScriptApi/GetYakScriptTags"
	YakScriptApi_QueryYakScriptLocalAndUser_FullMethodName        = "/ypb.YakScriptApi/QueryYakScriptLocalAndUser"
	YakScriptApi_QueryYakScriptByOnlineGroup_FullMethodName       = "/ypb.YakScriptApi/QueryYakScriptByOnlineGroup"
	YakScriptApi_QueryYakScriptLocalAll_FullMethodName            = "/ypb.YakScriptApi/QueryYakScriptLocalAll"
	YakScriptApi_QueryYakScriptExecResult_FullMethodName          = "/ypb.YakScriptApi/QueryYakScriptExecResult"
	YakScriptApi_QueryYakScriptNameInExecResult_FullMethodName    = "/ypb.YakScriptApi/QueryYakScriptNameInExecResult"
	YakScriptApi_DeleteYakScriptExecResult_FullMethodName         = "/ypb.YakScriptApi/DeleteYakScriptExecResult"
	YakScriptApi_DeleteYakScriptExec_FullMethodName               = "/ypb.YakScriptApi/DeleteYakScriptExec"
	YakScriptApi_GetAvailableYakScriptTags_FullMethodName         = "/ypb.YakScriptApi/GetAvailableYakScriptTags"
	YakScriptApi_ForceUpdateAvailableYakScriptTags_FullMethodName = "/ypb.YakScriptApi/ForceUpdateAvailableYakScriptTags"
)

// YakScriptApiClient is the client API for YakScriptApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YakScriptApiClient interface {
	// yakScript
	QueryYakScript(ctx context.Context, in *QueryYakScriptRequest, opts ...grpc.CallOption) (*QueryYakScriptResponse, error)
	QueryYakScriptByYakScriptName(ctx context.Context, in *QueryYakScriptRequest, opts ...grpc.CallOption) (YakScriptApi_QueryYakScriptByYakScriptNameClient, error)
	SaveYakScript(ctx context.Context, in *YakScript, opts ...grpc.CallOption) (*YakScript, error)
	DeleteYakScript(ctx context.Context, in *DeleteYakScriptRequest, opts ...grpc.CallOption) (*Empty, error)
	GetYakScriptById(ctx context.Context, in *GetYakScriptByIdRequest, opts ...grpc.CallOption) (*YakScript, error)
	GetYakScriptByName(ctx context.Context, in *GetYakScriptByNameRequest, opts ...grpc.CallOption) (*YakScript, error)
	GetYakScriptByOnlineID(ctx context.Context, in *GetYakScriptByOnlineIDRequest, opts ...grpc.CallOption) (*YakScript, error)
	IgnoreYakScript(ctx context.Context, in *DeleteYakScriptRequest, opts ...grpc.CallOption) (*Empty, error)
	UnIgnoreYakScript(ctx context.Context, in *DeleteYakScriptRequest, opts ...grpc.CallOption) (*Empty, error)
	ExportYakScript(ctx context.Context, in *ExportYakScriptRequest, opts ...grpc.CallOption) (*ExportYakScriptResponse, error)
	GetYakScriptTags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetYakScriptTagsResponse, error)
	QueryYakScriptLocalAndUser(ctx context.Context, in *QueryYakScriptLocalAndUserRequest, opts ...grpc.CallOption) (*QueryYakScriptLocalAndUserResponse, error)
	QueryYakScriptByOnlineGroup(ctx context.Context, in *QueryYakScriptByOnlineGroupRequest, opts ...grpc.CallOption) (*QueryYakScriptLocalAndUserResponse, error)
	QueryYakScriptLocalAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QueryYakScriptLocalAndUserResponse, error)
	// 对插件结果的操作
	QueryYakScriptExecResult(ctx context.Context, in *QueryYakScriptExecResultRequest, opts ...grpc.CallOption) (*QueryYakScriptExecResultResponse, error)
	QueryYakScriptNameInExecResult(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*YakScriptNames, error)
	DeleteYakScriptExecResult(ctx context.Context, in *DeleteYakScriptExecResultRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteYakScriptExec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// 获取 Tags
	GetAvailableYakScriptTags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Fields, error)
	ForceUpdateAvailableYakScriptTags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type yakScriptApiClient struct {
	cc grpc.ClientConnInterface
}

func NewYakScriptApiClient(cc grpc.ClientConnInterface) YakScriptApiClient {
	return &yakScriptApiClient{cc}
}

func (c *yakScriptApiClient) QueryYakScript(ctx context.Context, in *QueryYakScriptRequest, opts ...grpc.CallOption) (*QueryYakScriptResponse, error) {
	out := new(QueryYakScriptResponse)
	err := c.cc.Invoke(ctx, YakScriptApi_QueryYakScript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) QueryYakScriptByYakScriptName(ctx context.Context, in *QueryYakScriptRequest, opts ...grpc.CallOption) (YakScriptApi_QueryYakScriptByYakScriptNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &YakScriptApi_ServiceDesc.Streams[0], YakScriptApi_QueryYakScriptByYakScriptName_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &yakScriptApiQueryYakScriptByYakScriptNameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type YakScriptApi_QueryYakScriptByYakScriptNameClient interface {
	Recv() (*YakScript, error)
	grpc.ClientStream
}

type yakScriptApiQueryYakScriptByYakScriptNameClient struct {
	grpc.ClientStream
}

func (x *yakScriptApiQueryYakScriptByYakScriptNameClient) Recv() (*YakScript, error) {
	m := new(YakScript)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *yakScriptApiClient) SaveYakScript(ctx context.Context, in *YakScript, opts ...grpc.CallOption) (*YakScript, error) {
	out := new(YakScript)
	err := c.cc.Invoke(ctx, YakScriptApi_SaveYakScript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) DeleteYakScript(ctx context.Context, in *DeleteYakScriptRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, YakScriptApi_DeleteYakScript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) GetYakScriptById(ctx context.Context, in *GetYakScriptByIdRequest, opts ...grpc.CallOption) (*YakScript, error) {
	out := new(YakScript)
	err := c.cc.Invoke(ctx, YakScriptApi_GetYakScriptById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) GetYakScriptByName(ctx context.Context, in *GetYakScriptByNameRequest, opts ...grpc.CallOption) (*YakScript, error) {
	out := new(YakScript)
	err := c.cc.Invoke(ctx, YakScriptApi_GetYakScriptByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) GetYakScriptByOnlineID(ctx context.Context, in *GetYakScriptByOnlineIDRequest, opts ...grpc.CallOption) (*YakScript, error) {
	out := new(YakScript)
	err := c.cc.Invoke(ctx, YakScriptApi_GetYakScriptByOnlineID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) IgnoreYakScript(ctx context.Context, in *DeleteYakScriptRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, YakScriptApi_IgnoreYakScript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) UnIgnoreYakScript(ctx context.Context, in *DeleteYakScriptRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, YakScriptApi_UnIgnoreYakScript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) ExportYakScript(ctx context.Context, in *ExportYakScriptRequest, opts ...grpc.CallOption) (*ExportYakScriptResponse, error) {
	out := new(ExportYakScriptResponse)
	err := c.cc.Invoke(ctx, YakScriptApi_ExportYakScript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) GetYakScriptTags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetYakScriptTagsResponse, error) {
	out := new(GetYakScriptTagsResponse)
	err := c.cc.Invoke(ctx, YakScriptApi_GetYakScriptTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) QueryYakScriptLocalAndUser(ctx context.Context, in *QueryYakScriptLocalAndUserRequest, opts ...grpc.CallOption) (*QueryYakScriptLocalAndUserResponse, error) {
	out := new(QueryYakScriptLocalAndUserResponse)
	err := c.cc.Invoke(ctx, YakScriptApi_QueryYakScriptLocalAndUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) QueryYakScriptByOnlineGroup(ctx context.Context, in *QueryYakScriptByOnlineGroupRequest, opts ...grpc.CallOption) (*QueryYakScriptLocalAndUserResponse, error) {
	out := new(QueryYakScriptLocalAndUserResponse)
	err := c.cc.Invoke(ctx, YakScriptApi_QueryYakScriptByOnlineGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) QueryYakScriptLocalAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*QueryYakScriptLocalAndUserResponse, error) {
	out := new(QueryYakScriptLocalAndUserResponse)
	err := c.cc.Invoke(ctx, YakScriptApi_QueryYakScriptLocalAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) QueryYakScriptExecResult(ctx context.Context, in *QueryYakScriptExecResultRequest, opts ...grpc.CallOption) (*QueryYakScriptExecResultResponse, error) {
	out := new(QueryYakScriptExecResultResponse)
	err := c.cc.Invoke(ctx, YakScriptApi_QueryYakScriptExecResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) QueryYakScriptNameInExecResult(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*YakScriptNames, error) {
	out := new(YakScriptNames)
	err := c.cc.Invoke(ctx, YakScriptApi_QueryYakScriptNameInExecResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) DeleteYakScriptExecResult(ctx context.Context, in *DeleteYakScriptExecResultRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, YakScriptApi_DeleteYakScriptExecResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) DeleteYakScriptExec(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, YakScriptApi_DeleteYakScriptExec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) GetAvailableYakScriptTags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Fields, error) {
	out := new(Fields)
	err := c.cc.Invoke(ctx, YakScriptApi_GetAvailableYakScriptTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yakScriptApiClient) ForceUpdateAvailableYakScriptTags(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, YakScriptApi_ForceUpdateAvailableYakScriptTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YakScriptApiServer is the server API for YakScriptApi service.
// All implementations must embed UnimplementedYakScriptApiServer
// for forward compatibility
type YakScriptApiServer interface {
	// yakScript
	QueryYakScript(context.Context, *QueryYakScriptRequest) (*QueryYakScriptResponse, error)
	QueryYakScriptByYakScriptName(*QueryYakScriptRequest, YakScriptApi_QueryYakScriptByYakScriptNameServer) error
	SaveYakScript(context.Context, *YakScript) (*YakScript, error)
	DeleteYakScript(context.Context, *DeleteYakScriptRequest) (*Empty, error)
	GetYakScriptById(context.Context, *GetYakScriptByIdRequest) (*YakScript, error)
	GetYakScriptByName(context.Context, *GetYakScriptByNameRequest) (*YakScript, error)
	GetYakScriptByOnlineID(context.Context, *GetYakScriptByOnlineIDRequest) (*YakScript, error)
	IgnoreYakScript(context.Context, *DeleteYakScriptRequest) (*Empty, error)
	UnIgnoreYakScript(context.Context, *DeleteYakScriptRequest) (*Empty, error)
	ExportYakScript(context.Context, *ExportYakScriptRequest) (*ExportYakScriptResponse, error)
	GetYakScriptTags(context.Context, *Empty) (*GetYakScriptTagsResponse, error)
	QueryYakScriptLocalAndUser(context.Context, *QueryYakScriptLocalAndUserRequest) (*QueryYakScriptLocalAndUserResponse, error)
	QueryYakScriptByOnlineGroup(context.Context, *QueryYakScriptByOnlineGroupRequest) (*QueryYakScriptLocalAndUserResponse, error)
	QueryYakScriptLocalAll(context.Context, *Empty) (*QueryYakScriptLocalAndUserResponse, error)
	// 对插件结果的操作
	QueryYakScriptExecResult(context.Context, *QueryYakScriptExecResultRequest) (*QueryYakScriptExecResultResponse, error)
	QueryYakScriptNameInExecResult(context.Context, *Empty) (*YakScriptNames, error)
	DeleteYakScriptExecResult(context.Context, *DeleteYakScriptExecResultRequest) (*Empty, error)
	DeleteYakScriptExec(context.Context, *Empty) (*Empty, error)
	// 获取 Tags
	GetAvailableYakScriptTags(context.Context, *Empty) (*Fields, error)
	ForceUpdateAvailableYakScriptTags(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedYakScriptApiServer()
}

// UnimplementedYakScriptApiServer must be embedded to have forward compatible implementations.
type UnimplementedYakScriptApiServer struct {
}

func (UnimplementedYakScriptApiServer) QueryYakScript(context.Context, *QueryYakScriptRequest) (*QueryYakScriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryYakScript not implemented")
}
func (UnimplementedYakScriptApiServer) QueryYakScriptByYakScriptName(*QueryYakScriptRequest, YakScriptApi_QueryYakScriptByYakScriptNameServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryYakScriptByYakScriptName not implemented")
}
func (UnimplementedYakScriptApiServer) SaveYakScript(context.Context, *YakScript) (*YakScript, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveYakScript not implemented")
}
func (UnimplementedYakScriptApiServer) DeleteYakScript(context.Context, *DeleteYakScriptRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteYakScript not implemented")
}
func (UnimplementedYakScriptApiServer) GetYakScriptById(context.Context, *GetYakScriptByIdRequest) (*YakScript, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetYakScriptById not implemented")
}
func (UnimplementedYakScriptApiServer) GetYakScriptByName(context.Context, *GetYakScriptByNameRequest) (*YakScript, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetYakScriptByName not implemented")
}
func (UnimplementedYakScriptApiServer) GetYakScriptByOnlineID(context.Context, *GetYakScriptByOnlineIDRequest) (*YakScript, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetYakScriptByOnlineID not implemented")
}
func (UnimplementedYakScriptApiServer) IgnoreYakScript(context.Context, *DeleteYakScriptRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IgnoreYakScript not implemented")
}
func (UnimplementedYakScriptApiServer) UnIgnoreYakScript(context.Context, *DeleteYakScriptRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnIgnoreYakScript not implemented")
}
func (UnimplementedYakScriptApiServer) ExportYakScript(context.Context, *ExportYakScriptRequest) (*ExportYakScriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportYakScript not implemented")
}
func (UnimplementedYakScriptApiServer) GetYakScriptTags(context.Context, *Empty) (*GetYakScriptTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetYakScriptTags not implemented")
}
func (UnimplementedYakScriptApiServer) QueryYakScriptLocalAndUser(context.Context, *QueryYakScriptLocalAndUserRequest) (*QueryYakScriptLocalAndUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryYakScriptLocalAndUser not implemented")
}
func (UnimplementedYakScriptApiServer) QueryYakScriptByOnlineGroup(context.Context, *QueryYakScriptByOnlineGroupRequest) (*QueryYakScriptLocalAndUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryYakScriptByOnlineGroup not implemented")
}
func (UnimplementedYakScriptApiServer) QueryYakScriptLocalAll(context.Context, *Empty) (*QueryYakScriptLocalAndUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryYakScriptLocalAll not implemented")
}
func (UnimplementedYakScriptApiServer) QueryYakScriptExecResult(context.Context, *QueryYakScriptExecResultRequest) (*QueryYakScriptExecResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryYakScriptExecResult not implemented")
}
func (UnimplementedYakScriptApiServer) QueryYakScriptNameInExecResult(context.Context, *Empty) (*YakScriptNames, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryYakScriptNameInExecResult not implemented")
}
func (UnimplementedYakScriptApiServer) DeleteYakScriptExecResult(context.Context, *DeleteYakScriptExecResultRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteYakScriptExecResult not implemented")
}
func (UnimplementedYakScriptApiServer) DeleteYakScriptExec(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteYakScriptExec not implemented")
}
func (UnimplementedYakScriptApiServer) GetAvailableYakScriptTags(context.Context, *Empty) (*Fields, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableYakScriptTags not implemented")
}
func (UnimplementedYakScriptApiServer) ForceUpdateAvailableYakScriptTags(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForceUpdateAvailableYakScriptTags not implemented")
}
func (UnimplementedYakScriptApiServer) mustEmbedUnimplementedYakScriptApiServer() {}

// UnsafeYakScriptApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YakScriptApiServer will
// result in compilation errors.
type UnsafeYakScriptApiServer interface {
	mustEmbedUnimplementedYakScriptApiServer()
}

func RegisterYakScriptApiServer(s grpc.ServiceRegistrar, srv YakScriptApiServer) {
	s.RegisterService(&YakScriptApi_ServiceDesc, srv)
}

func _YakScriptApi_QueryYakScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryYakScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).QueryYakScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_QueryYakScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).QueryYakScript(ctx, req.(*QueryYakScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_QueryYakScriptByYakScriptName_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryYakScriptRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(YakScriptApiServer).QueryYakScriptByYakScriptName(m, &yakScriptApiQueryYakScriptByYakScriptNameServer{stream})
}

type YakScriptApi_QueryYakScriptByYakScriptNameServer interface {
	Send(*YakScript) error
	grpc.ServerStream
}

type yakScriptApiQueryYakScriptByYakScriptNameServer struct {
	grpc.ServerStream
}

func (x *yakScriptApiQueryYakScriptByYakScriptNameServer) Send(m *YakScript) error {
	return x.ServerStream.SendMsg(m)
}

func _YakScriptApi_SaveYakScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YakScript)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).SaveYakScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_SaveYakScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).SaveYakScript(ctx, req.(*YakScript))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_DeleteYakScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteYakScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).DeleteYakScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_DeleteYakScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).DeleteYakScript(ctx, req.(*DeleteYakScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_GetYakScriptById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetYakScriptByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).GetYakScriptById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_GetYakScriptById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).GetYakScriptById(ctx, req.(*GetYakScriptByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_GetYakScriptByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetYakScriptByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).GetYakScriptByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_GetYakScriptByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).GetYakScriptByName(ctx, req.(*GetYakScriptByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_GetYakScriptByOnlineID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetYakScriptByOnlineIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).GetYakScriptByOnlineID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_GetYakScriptByOnlineID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).GetYakScriptByOnlineID(ctx, req.(*GetYakScriptByOnlineIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_IgnoreYakScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteYakScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).IgnoreYakScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_IgnoreYakScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).IgnoreYakScript(ctx, req.(*DeleteYakScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_UnIgnoreYakScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteYakScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).UnIgnoreYakScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_UnIgnoreYakScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).UnIgnoreYakScript(ctx, req.(*DeleteYakScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_ExportYakScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportYakScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).ExportYakScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_ExportYakScript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).ExportYakScript(ctx, req.(*ExportYakScriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_GetYakScriptTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).GetYakScriptTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_GetYakScriptTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).GetYakScriptTags(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_QueryYakScriptLocalAndUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryYakScriptLocalAndUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).QueryYakScriptLocalAndUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_QueryYakScriptLocalAndUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).QueryYakScriptLocalAndUser(ctx, req.(*QueryYakScriptLocalAndUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_QueryYakScriptByOnlineGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryYakScriptByOnlineGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).QueryYakScriptByOnlineGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_QueryYakScriptByOnlineGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).QueryYakScriptByOnlineGroup(ctx, req.(*QueryYakScriptByOnlineGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_QueryYakScriptLocalAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).QueryYakScriptLocalAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_QueryYakScriptLocalAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).QueryYakScriptLocalAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_QueryYakScriptExecResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryYakScriptExecResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).QueryYakScriptExecResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_QueryYakScriptExecResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).QueryYakScriptExecResult(ctx, req.(*QueryYakScriptExecResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_QueryYakScriptNameInExecResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).QueryYakScriptNameInExecResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_QueryYakScriptNameInExecResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).QueryYakScriptNameInExecResult(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_DeleteYakScriptExecResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteYakScriptExecResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).DeleteYakScriptExecResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_DeleteYakScriptExecResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).DeleteYakScriptExecResult(ctx, req.(*DeleteYakScriptExecResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_DeleteYakScriptExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).DeleteYakScriptExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_DeleteYakScriptExec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).DeleteYakScriptExec(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_GetAvailableYakScriptTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).GetAvailableYakScriptTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_GetAvailableYakScriptTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).GetAvailableYakScriptTags(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _YakScriptApi_ForceUpdateAvailableYakScriptTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YakScriptApiServer).ForceUpdateAvailableYakScriptTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YakScriptApi_ForceUpdateAvailableYakScriptTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YakScriptApiServer).ForceUpdateAvailableYakScriptTags(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// YakScriptApi_ServiceDesc is the grpc.ServiceDesc for YakScriptApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YakScriptApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.YakScriptApi",
	HandlerType: (*YakScriptApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryYakScript",
			Handler:    _YakScriptApi_QueryYakScript_Handler,
		},
		{
			MethodName: "SaveYakScript",
			Handler:    _YakScriptApi_SaveYakScript_Handler,
		},
		{
			MethodName: "DeleteYakScript",
			Handler:    _YakScriptApi_DeleteYakScript_Handler,
		},
		{
			MethodName: "GetYakScriptById",
			Handler:    _YakScriptApi_GetYakScriptById_Handler,
		},
		{
			MethodName: "GetYakScriptByName",
			Handler:    _YakScriptApi_GetYakScriptByName_Handler,
		},
		{
			MethodName: "GetYakScriptByOnlineID",
			Handler:    _YakScriptApi_GetYakScriptByOnlineID_Handler,
		},
		{
			MethodName: "IgnoreYakScript",
			Handler:    _YakScriptApi_IgnoreYakScript_Handler,
		},
		{
			MethodName: "UnIgnoreYakScript",
			Handler:    _YakScriptApi_UnIgnoreYakScript_Handler,
		},
		{
			MethodName: "ExportYakScript",
			Handler:    _YakScriptApi_ExportYakScript_Handler,
		},
		{
			MethodName: "GetYakScriptTags",
			Handler:    _YakScriptApi_GetYakScriptTags_Handler,
		},
		{
			MethodName: "QueryYakScriptLocalAndUser",
			Handler:    _YakScriptApi_QueryYakScriptLocalAndUser_Handler,
		},
		{
			MethodName: "QueryYakScriptByOnlineGroup",
			Handler:    _YakScriptApi_QueryYakScriptByOnlineGroup_Handler,
		},
		{
			MethodName: "QueryYakScriptLocalAll",
			Handler:    _YakScriptApi_QueryYakScriptLocalAll_Handler,
		},
		{
			MethodName: "QueryYakScriptExecResult",
			Handler:    _YakScriptApi_QueryYakScriptExecResult_Handler,
		},
		{
			MethodName: "QueryYakScriptNameInExecResult",
			Handler:    _YakScriptApi_QueryYakScriptNameInExecResult_Handler,
		},
		{
			MethodName: "DeleteYakScriptExecResult",
			Handler:    _YakScriptApi_DeleteYakScriptExecResult_Handler,
		},
		{
			MethodName: "DeleteYakScriptExec",
			Handler:    _YakScriptApi_DeleteYakScriptExec_Handler,
		},
		{
			MethodName: "GetAvailableYakScriptTags",
			Handler:    _YakScriptApi_GetAvailableYakScriptTags_Handler,
		},
		{
			MethodName: "ForceUpdateAvailableYakScriptTags",
			Handler:    _YakScriptApi_ForceUpdateAvailableYakScriptTags_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueryYakScriptByYakScriptName",
			Handler:       _YakScriptApi_QueryYakScriptByYakScriptName_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "yak_script_api.proto",
}
