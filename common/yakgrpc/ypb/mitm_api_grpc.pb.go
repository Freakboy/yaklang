// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: mitm_api.proto

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
	MITMApi_MITM_FullMethodName             = "/ypb.MITMApi/MITM"
	MITMApi_DownloadMITMCert_FullMethodName = "/ypb.MITMApi/DownloadMITMCert"
	MITMApi_GetCurrentRules_FullMethodName  = "/ypb.MITMApi/GetCurrentRules"
	MITMApi_SetCurrentRules_FullMethodName  = "/ypb.MITMApi/SetCurrentRules"
)

// MITMApiClient is the client API for MITMApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MITMApiClient interface {
	// 中间人劫持
	MITM(ctx context.Context, opts ...grpc.CallOption) (MITMApi_MITMClient, error)
	DownloadMITMCert(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MITMCert, error)
	GetCurrentRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MITMContentReplacers, error)
	SetCurrentRules(ctx context.Context, in *MITMContentReplacers, opts ...grpc.CallOption) (*Empty, error)
}

type mITMApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMITMApiClient(cc grpc.ClientConnInterface) MITMApiClient {
	return &mITMApiClient{cc}
}

func (c *mITMApiClient) MITM(ctx context.Context, opts ...grpc.CallOption) (MITMApi_MITMClient, error) {
	stream, err := c.cc.NewStream(ctx, &MITMApi_ServiceDesc.Streams[0], MITMApi_MITM_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &mITMApiMITMClient{stream}
	return x, nil
}

type MITMApi_MITMClient interface {
	Send(*MITMRequest) error
	Recv() (*MITMResponse, error)
	grpc.ClientStream
}

type mITMApiMITMClient struct {
	grpc.ClientStream
}

func (x *mITMApiMITMClient) Send(m *MITMRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mITMApiMITMClient) Recv() (*MITMResponse, error) {
	m := new(MITMResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mITMApiClient) DownloadMITMCert(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MITMCert, error) {
	out := new(MITMCert)
	err := c.cc.Invoke(ctx, MITMApi_DownloadMITMCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mITMApiClient) GetCurrentRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MITMContentReplacers, error) {
	out := new(MITMContentReplacers)
	err := c.cc.Invoke(ctx, MITMApi_GetCurrentRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mITMApiClient) SetCurrentRules(ctx context.Context, in *MITMContentReplacers, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, MITMApi_SetCurrentRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MITMApiServer is the server API for MITMApi service.
// All implementations must embed UnimplementedMITMApiServer
// for forward compatibility
type MITMApiServer interface {
	// 中间人劫持
	MITM(MITMApi_MITMServer) error
	DownloadMITMCert(context.Context, *Empty) (*MITMCert, error)
	GetCurrentRules(context.Context, *Empty) (*MITMContentReplacers, error)
	SetCurrentRules(context.Context, *MITMContentReplacers) (*Empty, error)
	mustEmbedUnimplementedMITMApiServer()
}

// UnimplementedMITMApiServer must be embedded to have forward compatible implementations.
type UnimplementedMITMApiServer struct {
}

func (UnimplementedMITMApiServer) MITM(MITMApi_MITMServer) error {
	return status.Errorf(codes.Unimplemented, "method MITM not implemented")
}
func (UnimplementedMITMApiServer) DownloadMITMCert(context.Context, *Empty) (*MITMCert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadMITMCert not implemented")
}
func (UnimplementedMITMApiServer) GetCurrentRules(context.Context, *Empty) (*MITMContentReplacers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentRules not implemented")
}
func (UnimplementedMITMApiServer) SetCurrentRules(context.Context, *MITMContentReplacers) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCurrentRules not implemented")
}
func (UnimplementedMITMApiServer) mustEmbedUnimplementedMITMApiServer() {}

// UnsafeMITMApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MITMApiServer will
// result in compilation errors.
type UnsafeMITMApiServer interface {
	mustEmbedUnimplementedMITMApiServer()
}

func RegisterMITMApiServer(s grpc.ServiceRegistrar, srv MITMApiServer) {
	s.RegisterService(&MITMApi_ServiceDesc, srv)
}

func _MITMApi_MITM_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MITMApiServer).MITM(&mITMApiMITMServer{stream})
}

type MITMApi_MITMServer interface {
	Send(*MITMResponse) error
	Recv() (*MITMRequest, error)
	grpc.ServerStream
}

type mITMApiMITMServer struct {
	grpc.ServerStream
}

func (x *mITMApiMITMServer) Send(m *MITMResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mITMApiMITMServer) Recv() (*MITMRequest, error) {
	m := new(MITMRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MITMApi_DownloadMITMCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MITMApiServer).DownloadMITMCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MITMApi_DownloadMITMCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MITMApiServer).DownloadMITMCert(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MITMApi_GetCurrentRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MITMApiServer).GetCurrentRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MITMApi_GetCurrentRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MITMApiServer).GetCurrentRules(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MITMApi_SetCurrentRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MITMContentReplacers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MITMApiServer).SetCurrentRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MITMApi_SetCurrentRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MITMApiServer).SetCurrentRules(ctx, req.(*MITMContentReplacers))
	}
	return interceptor(ctx, in, info, handler)
}

// MITMApi_ServiceDesc is the grpc.ServiceDesc for MITMApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MITMApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.MITMApi",
	HandlerType: (*MITMApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownloadMITMCert",
			Handler:    _MITMApi_DownloadMITMCert_Handler,
		},
		{
			MethodName: "GetCurrentRules",
			Handler:    _MITMApi_GetCurrentRules_Handler,
		},
		{
			MethodName: "SetCurrentRules",
			Handler:    _MITMApi_SetCurrentRules_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MITM",
			Handler:       _MITMApi_MITM_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mitm_api.proto",
}

const (
	MITMReplacerApi_ExportMITMReplacerRules_FullMethodName = "/ypb.MITMReplacerApi/ExportMITMReplacerRules"
	MITMReplacerApi_ImportMITMReplacerRules_FullMethodName = "/ypb.MITMReplacerApi/ImportMITMReplacerRules"
)

// MITMReplacerApiClient is the client API for MITMReplacerApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MITMReplacerApiClient interface {
	// MITM 衍生功能：
	// Replacers 管理
	ExportMITMReplacerRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExportMITMReplacerRulesResponse, error)
	ImportMITMReplacerRules(ctx context.Context, in *ImportMITMReplacerRulesRequest, opts ...grpc.CallOption) (*Empty, error)
}

type mITMReplacerApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMITMReplacerApiClient(cc grpc.ClientConnInterface) MITMReplacerApiClient {
	return &mITMReplacerApiClient{cc}
}

func (c *mITMReplacerApiClient) ExportMITMReplacerRules(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ExportMITMReplacerRulesResponse, error) {
	out := new(ExportMITMReplacerRulesResponse)
	err := c.cc.Invoke(ctx, MITMReplacerApi_ExportMITMReplacerRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mITMReplacerApiClient) ImportMITMReplacerRules(ctx context.Context, in *ImportMITMReplacerRulesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, MITMReplacerApi_ImportMITMReplacerRules_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MITMReplacerApiServer is the server API for MITMReplacerApi service.
// All implementations must embed UnimplementedMITMReplacerApiServer
// for forward compatibility
type MITMReplacerApiServer interface {
	// MITM 衍生功能：
	// Replacers 管理
	ExportMITMReplacerRules(context.Context, *Empty) (*ExportMITMReplacerRulesResponse, error)
	ImportMITMReplacerRules(context.Context, *ImportMITMReplacerRulesRequest) (*Empty, error)
	mustEmbedUnimplementedMITMReplacerApiServer()
}

// UnimplementedMITMReplacerApiServer must be embedded to have forward compatible implementations.
type UnimplementedMITMReplacerApiServer struct {
}

func (UnimplementedMITMReplacerApiServer) ExportMITMReplacerRules(context.Context, *Empty) (*ExportMITMReplacerRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportMITMReplacerRules not implemented")
}
func (UnimplementedMITMReplacerApiServer) ImportMITMReplacerRules(context.Context, *ImportMITMReplacerRulesRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportMITMReplacerRules not implemented")
}
func (UnimplementedMITMReplacerApiServer) mustEmbedUnimplementedMITMReplacerApiServer() {}

// UnsafeMITMReplacerApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MITMReplacerApiServer will
// result in compilation errors.
type UnsafeMITMReplacerApiServer interface {
	mustEmbedUnimplementedMITMReplacerApiServer()
}

func RegisterMITMReplacerApiServer(s grpc.ServiceRegistrar, srv MITMReplacerApiServer) {
	s.RegisterService(&MITMReplacerApi_ServiceDesc, srv)
}

func _MITMReplacerApi_ExportMITMReplacerRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MITMReplacerApiServer).ExportMITMReplacerRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MITMReplacerApi_ExportMITMReplacerRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MITMReplacerApiServer).ExportMITMReplacerRules(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MITMReplacerApi_ImportMITMReplacerRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportMITMReplacerRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MITMReplacerApiServer).ImportMITMReplacerRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MITMReplacerApi_ImportMITMReplacerRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MITMReplacerApiServer).ImportMITMReplacerRules(ctx, req.(*ImportMITMReplacerRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MITMReplacerApi_ServiceDesc is the grpc.ServiceDesc for MITMReplacerApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MITMReplacerApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.MITMReplacerApi",
	HandlerType: (*MITMReplacerApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportMITMReplacerRules",
			Handler:    _MITMReplacerApi_ExportMITMReplacerRules_Handler,
		},
		{
			MethodName: "ImportMITMReplacerRules",
			Handler:    _MITMReplacerApi_ImportMITMReplacerRules_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitm_api.proto",
}

const (
	MITMExtractedDataApi_QueryMITMRuleExtractedData_FullMethodName = "/ypb.MITMExtractedDataApi/QueryMITMRuleExtractedData"
)

// MITMExtractedDataApiClient is the client API for MITMExtractedDataApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MITMExtractedDataApiClient interface {
	// 从规则中提取数据
	QueryMITMRuleExtractedData(ctx context.Context, in *QueryMITMRuleExtractedDataRequest, opts ...grpc.CallOption) (*QueryMITMRuleExtractedDataResponse, error)
}

type mITMExtractedDataApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMITMExtractedDataApiClient(cc grpc.ClientConnInterface) MITMExtractedDataApiClient {
	return &mITMExtractedDataApiClient{cc}
}

func (c *mITMExtractedDataApiClient) QueryMITMRuleExtractedData(ctx context.Context, in *QueryMITMRuleExtractedDataRequest, opts ...grpc.CallOption) (*QueryMITMRuleExtractedDataResponse, error) {
	out := new(QueryMITMRuleExtractedDataResponse)
	err := c.cc.Invoke(ctx, MITMExtractedDataApi_QueryMITMRuleExtractedData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MITMExtractedDataApiServer is the server API for MITMExtractedDataApi service.
// All implementations must embed UnimplementedMITMExtractedDataApiServer
// for forward compatibility
type MITMExtractedDataApiServer interface {
	// 从规则中提取数据
	QueryMITMRuleExtractedData(context.Context, *QueryMITMRuleExtractedDataRequest) (*QueryMITMRuleExtractedDataResponse, error)
	mustEmbedUnimplementedMITMExtractedDataApiServer()
}

// UnimplementedMITMExtractedDataApiServer must be embedded to have forward compatible implementations.
type UnimplementedMITMExtractedDataApiServer struct {
}

func (UnimplementedMITMExtractedDataApiServer) QueryMITMRuleExtractedData(context.Context, *QueryMITMRuleExtractedDataRequest) (*QueryMITMRuleExtractedDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMITMRuleExtractedData not implemented")
}
func (UnimplementedMITMExtractedDataApiServer) mustEmbedUnimplementedMITMExtractedDataApiServer() {}

// UnsafeMITMExtractedDataApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MITMExtractedDataApiServer will
// result in compilation errors.
type UnsafeMITMExtractedDataApiServer interface {
	mustEmbedUnimplementedMITMExtractedDataApiServer()
}

func RegisterMITMExtractedDataApiServer(s grpc.ServiceRegistrar, srv MITMExtractedDataApiServer) {
	s.RegisterService(&MITMExtractedDataApi_ServiceDesc, srv)
}

func _MITMExtractedDataApi_QueryMITMRuleExtractedData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMITMRuleExtractedDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MITMExtractedDataApiServer).QueryMITMRuleExtractedData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MITMExtractedDataApi_QueryMITMRuleExtractedData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MITMExtractedDataApiServer).QueryMITMRuleExtractedData(ctx, req.(*QueryMITMRuleExtractedDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MITMExtractedDataApi_ServiceDesc is the grpc.ServiceDesc for MITMExtractedDataApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MITMExtractedDataApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.MITMExtractedDataApi",
	HandlerType: (*MITMExtractedDataApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryMITMRuleExtractedData",
			Handler:    _MITMExtractedDataApi_QueryMITMRuleExtractedData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitm_api.proto",
}

const (
	MITMFilterApi_SetMITMFilter_FullMethodName = "/ypb.MITMFilterApi/SetMITMFilter"
	MITMFilterApi_GetMITMFilter_FullMethodName = "/ypb.MITMFilterApi/GetMITMFilter"
)

// MITMFilterApiClient is the client API for MITMFilterApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MITMFilterApiClient interface {
	// MITM 劫持的过滤器
	SetMITMFilter(ctx context.Context, in *SetMITMFilterRequest, opts ...grpc.CallOption) (*SetMITMFilterResponse, error)
	GetMITMFilter(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SetMITMFilterRequest, error)
}

type mITMFilterApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMITMFilterApiClient(cc grpc.ClientConnInterface) MITMFilterApiClient {
	return &mITMFilterApiClient{cc}
}

func (c *mITMFilterApiClient) SetMITMFilter(ctx context.Context, in *SetMITMFilterRequest, opts ...grpc.CallOption) (*SetMITMFilterResponse, error) {
	out := new(SetMITMFilterResponse)
	err := c.cc.Invoke(ctx, MITMFilterApi_SetMITMFilter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mITMFilterApiClient) GetMITMFilter(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SetMITMFilterRequest, error) {
	out := new(SetMITMFilterRequest)
	err := c.cc.Invoke(ctx, MITMFilterApi_GetMITMFilter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MITMFilterApiServer is the server API for MITMFilterApi service.
// All implementations must embed UnimplementedMITMFilterApiServer
// for forward compatibility
type MITMFilterApiServer interface {
	// MITM 劫持的过滤器
	SetMITMFilter(context.Context, *SetMITMFilterRequest) (*SetMITMFilterResponse, error)
	GetMITMFilter(context.Context, *Empty) (*SetMITMFilterRequest, error)
	mustEmbedUnimplementedMITMFilterApiServer()
}

// UnimplementedMITMFilterApiServer must be embedded to have forward compatible implementations.
type UnimplementedMITMFilterApiServer struct {
}

func (UnimplementedMITMFilterApiServer) SetMITMFilter(context.Context, *SetMITMFilterRequest) (*SetMITMFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMITMFilter not implemented")
}
func (UnimplementedMITMFilterApiServer) GetMITMFilter(context.Context, *Empty) (*SetMITMFilterRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMITMFilter not implemented")
}
func (UnimplementedMITMFilterApiServer) mustEmbedUnimplementedMITMFilterApiServer() {}

// UnsafeMITMFilterApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MITMFilterApiServer will
// result in compilation errors.
type UnsafeMITMFilterApiServer interface {
	mustEmbedUnimplementedMITMFilterApiServer()
}

func RegisterMITMFilterApiServer(s grpc.ServiceRegistrar, srv MITMFilterApiServer) {
	s.RegisterService(&MITMFilterApi_ServiceDesc, srv)
}

func _MITMFilterApi_SetMITMFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMITMFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MITMFilterApiServer).SetMITMFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MITMFilterApi_SetMITMFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MITMFilterApiServer).SetMITMFilter(ctx, req.(*SetMITMFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MITMFilterApi_GetMITMFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MITMFilterApiServer).GetMITMFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MITMFilterApi_GetMITMFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MITMFilterApiServer).GetMITMFilter(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MITMFilterApi_ServiceDesc is the grpc.ServiceDesc for MITMFilterApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MITMFilterApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ypb.MITMFilterApi",
	HandlerType: (*MITMFilterApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetMITMFilter",
			Handler:    _MITMFilterApi_SetMITMFilter_Handler,
		},
		{
			MethodName: "GetMITMFilter",
			Handler:    _MITMFilterApi_GetMITMFilter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitm_api.proto",
}
