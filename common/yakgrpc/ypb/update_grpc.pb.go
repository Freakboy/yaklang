// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: update.proto

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
	UpdateService_UpdateFromYakitResource_FullMethodName = "/ypb.UpdateService/UpdateFromYakitResource"
	UpdateService_UpdateFromGithub_FullMethodName        = "/ypb.UpdateService/UpdateFromGithub"
)

// UpdateServiceClient is the client API for UpdateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateServiceClient interface {
	// Yakit Store
	UpdateFromYakitResource(ctx context.Context, in *UpdateFromYakitResourceRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateFromGithub(ctx context.Context, in *UpdateFromGithubRequest, opts ...grpc.CallOption) (*Empty, error)
}

type updateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateServiceClient(cc grpc.ClientConnInterface) UpdateServiceClient {
	return &updateServiceClient{cc}
}

func (c *updateServiceClient) UpdateFromYakitResource(ctx context.Context, in *UpdateFromYakitResourceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UpdateService_UpdateFromYakitResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *updateServiceClient) UpdateFromGithub(ctx context.Context, in *UpdateFromGithubRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, UpdateService_UpdateFromGithub_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateServiceServer is the server API for UpdateService service.
// All implementations must embed UnimplementedUpdateServiceServer
// for forward compatibility
type UpdateServiceServer interface {
	// Yakit Store
	UpdateFromYakitResource(context.Context, *UpdateFromYakitResourceRequest) (*Empty, error)
	UpdateFromGithub(context.Context, *UpdateFromGithubRequest) (*Empty, error)
	mustEmbedUnimplementedUpdateServiceServer()
}

// UnimplementedUpdateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUpdateServiceServer struct {
}

func (UnimplementedUpdateServiceServer) UpdateFromYakitResource(context.Context, *UpdateFromYakitResourceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFromYakitResource not implemented")
}
func (UnimplementedUpdateServiceServer) UpdateFromGithub(context.Context, *UpdateFromGithubRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFromGithub not implemented")
}
func (UnimplementedUpdateServiceServer) mustEmbedUnimplementedUpdateServiceServer() {}

// UnsafeUpdateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateServiceServer will
// result in compilation errors.
type UnsafeUpdateServiceServer interface {
	mustEmbedUnimplementedUpdateServiceServer()
}

func RegisterUpdateServiceServer(s grpc.ServiceRegistrar, srv UpdateServiceServer) {
	s.RegisterService(&UpdateService_ServiceDesc, srv)
}

func _UpdateService_UpdateFromYakitResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFromYakitResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServiceServer).UpdateFromYakitResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UpdateService_UpdateFromYakitResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServiceServer).UpdateFromYakitResource(ctx, req.(*UpdateFromYakitResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UpdateService_UpdateFromGithub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFromGithubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServiceServer).UpdateFromGithub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UpdateService_UpdateFromGithub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServiceServer).UpdateFromGithub(ctx, req.(*UpdateFromGithubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UpdateService_ServiceDesc is the grpc.ServiceDesc for UpdateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.UpdateService",
	HandlerType: (*UpdateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateFromYakitResource",
			Handler:    _UpdateService_UpdateFromYakitResource_Handler,
		},
		{
			MethodName: "UpdateFromGithub",
			Handler:    _UpdateService_UpdateFromGithub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "update.proto",
}
