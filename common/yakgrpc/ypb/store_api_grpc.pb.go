// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: store_api.proto

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
	StoreApi_GetKey_FullMethodName              = "/ypb.StoreApi/GetKey"
	StoreApi_SetKey_FullMethodName              = "/ypb.StoreApi/SetKey"
	StoreApi_DelKey_FullMethodName              = "/ypb.StoreApi/DelKey"
	StoreApi_GetAllProcessEnvKey_FullMethodName = "/ypb.StoreApi/GetAllProcessEnvKey"
	StoreApi_SetProcessEnvKey_FullMethodName    = "/ypb.StoreApi/SetProcessEnvKey"
)

// StoreApiClient is the client API for StoreApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreApiClient interface {
	// 设置通用存储
	GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyResult, error)
	SetKey(ctx context.Context, in *SetKeyRequest, opts ...grpc.CallOption) (*Empty, error)
	DelKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*Empty, error)
	GetAllProcessEnvKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetProcessEnvKeyResult, error)
	SetProcessEnvKey(ctx context.Context, in *SetKeyRequest, opts ...grpc.CallOption) (*Empty, error)
}

type storeApiClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreApiClient(cc grpc.ClientConnInterface) StoreApiClient {
	return &storeApiClient{cc}
}

func (c *storeApiClient) GetKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*GetKeyResult, error) {
	out := new(GetKeyResult)
	err := c.cc.Invoke(ctx, StoreApi_GetKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeApiClient) SetKey(ctx context.Context, in *SetKeyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, StoreApi_SetKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeApiClient) DelKey(ctx context.Context, in *GetKeyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, StoreApi_DelKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeApiClient) GetAllProcessEnvKey(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetProcessEnvKeyResult, error) {
	out := new(GetProcessEnvKeyResult)
	err := c.cc.Invoke(ctx, StoreApi_GetAllProcessEnvKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeApiClient) SetProcessEnvKey(ctx context.Context, in *SetKeyRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, StoreApi_SetProcessEnvKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreApiServer is the server API for StoreApi service.
// All implementations must embed UnimplementedStoreApiServer
// for forward compatibility
type StoreApiServer interface {
	// 设置通用存储
	GetKey(context.Context, *GetKeyRequest) (*GetKeyResult, error)
	SetKey(context.Context, *SetKeyRequest) (*Empty, error)
	DelKey(context.Context, *GetKeyRequest) (*Empty, error)
	GetAllProcessEnvKey(context.Context, *Empty) (*GetProcessEnvKeyResult, error)
	SetProcessEnvKey(context.Context, *SetKeyRequest) (*Empty, error)
	mustEmbedUnimplementedStoreApiServer()
}

// UnimplementedStoreApiServer must be embedded to have forward compatible implementations.
type UnimplementedStoreApiServer struct {
}

func (UnimplementedStoreApiServer) GetKey(context.Context, *GetKeyRequest) (*GetKeyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}
func (UnimplementedStoreApiServer) SetKey(context.Context, *SetKeyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetKey not implemented")
}
func (UnimplementedStoreApiServer) DelKey(context.Context, *GetKeyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelKey not implemented")
}
func (UnimplementedStoreApiServer) GetAllProcessEnvKey(context.Context, *Empty) (*GetProcessEnvKeyResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProcessEnvKey not implemented")
}
func (UnimplementedStoreApiServer) SetProcessEnvKey(context.Context, *SetKeyRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProcessEnvKey not implemented")
}
func (UnimplementedStoreApiServer) mustEmbedUnimplementedStoreApiServer() {}

// UnsafeStoreApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreApiServer will
// result in compilation errors.
type UnsafeStoreApiServer interface {
	mustEmbedUnimplementedStoreApiServer()
}

func RegisterStoreApiServer(s grpc.ServiceRegistrar, srv StoreApiServer) {
	s.RegisterService(&StoreApi_ServiceDesc, srv)
}

func _StoreApi_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreApiServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreApi_GetKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreApiServer).GetKey(ctx, req.(*GetKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreApi_SetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreApiServer).SetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreApi_SetKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreApiServer).SetKey(ctx, req.(*SetKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreApi_DelKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreApiServer).DelKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreApi_DelKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreApiServer).DelKey(ctx, req.(*GetKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreApi_GetAllProcessEnvKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreApiServer).GetAllProcessEnvKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreApi_GetAllProcessEnvKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreApiServer).GetAllProcessEnvKey(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreApi_SetProcessEnvKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreApiServer).SetProcessEnvKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreApi_SetProcessEnvKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreApiServer).SetProcessEnvKey(ctx, req.(*SetKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreApi_ServiceDesc is the grpc.ServiceDesc for StoreApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.StoreApi",
	HandlerType: (*StoreApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKey",
			Handler:    _StoreApi_GetKey_Handler,
		},
		{
			MethodName: "SetKey",
			Handler:    _StoreApi_SetKey_Handler,
		},
		{
			MethodName: "DelKey",
			Handler:    _StoreApi_DelKey_Handler,
		},
		{
			MethodName: "GetAllProcessEnvKey",
			Handler:    _StoreApi_GetAllProcessEnvKey_Handler,
		},
		{
			MethodName: "SetProcessEnvKey",
			Handler:    _StoreApi_SetProcessEnvKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store_api.proto",
}
