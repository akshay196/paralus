// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/system/oidc_provider.proto

package rpcv3

import (
	context "context"
	v3 "github.com/RafaySystems/rcloud-base/proto/types/systempb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OIDCProviderClient is the client API for OIDCProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OIDCProviderClient interface {
	CreateOIDCProvider(ctx context.Context, in *v3.OIDCProvider, opts ...grpc.CallOption) (*v3.OIDCProvider, error)
	GetOIDCProvider(ctx context.Context, in *v3.OIDCProvider, opts ...grpc.CallOption) (*v3.OIDCProvider, error)
	ListOIDCProvider(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v3.OIDCProviderList, error)
	UpdateOIDCProvider(ctx context.Context, in *v3.OIDCProvider, opts ...grpc.CallOption) (*v3.OIDCProvider, error)
	DeleteOIDCProvider(ctx context.Context, in *v3.OIDCProvider, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type oIDCProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewOIDCProviderClient(cc grpc.ClientConnInterface) OIDCProviderClient {
	return &oIDCProviderClient{cc}
}

func (c *oIDCProviderClient) CreateOIDCProvider(ctx context.Context, in *v3.OIDCProvider, opts ...grpc.CallOption) (*v3.OIDCProvider, error) {
	out := new(v3.OIDCProvider)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.OIDCProvider/CreateOIDCProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCProviderClient) GetOIDCProvider(ctx context.Context, in *v3.OIDCProvider, opts ...grpc.CallOption) (*v3.OIDCProvider, error) {
	out := new(v3.OIDCProvider)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.OIDCProvider/GetOIDCProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCProviderClient) ListOIDCProvider(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*v3.OIDCProviderList, error) {
	out := new(v3.OIDCProviderList)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.OIDCProvider/ListOIDCProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCProviderClient) UpdateOIDCProvider(ctx context.Context, in *v3.OIDCProvider, opts ...grpc.CallOption) (*v3.OIDCProvider, error) {
	out := new(v3.OIDCProvider)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.OIDCProvider/UpdateOIDCProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCProviderClient) DeleteOIDCProvider(ctx context.Context, in *v3.OIDCProvider, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.OIDCProvider/DeleteOIDCProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OIDCProviderServer is the server API for OIDCProvider service.
// All implementations should embed UnimplementedOIDCProviderServer
// for forward compatibility
type OIDCProviderServer interface {
	CreateOIDCProvider(context.Context, *v3.OIDCProvider) (*v3.OIDCProvider, error)
	GetOIDCProvider(context.Context, *v3.OIDCProvider) (*v3.OIDCProvider, error)
	ListOIDCProvider(context.Context, *emptypb.Empty) (*v3.OIDCProviderList, error)
	UpdateOIDCProvider(context.Context, *v3.OIDCProvider) (*v3.OIDCProvider, error)
	DeleteOIDCProvider(context.Context, *v3.OIDCProvider) (*emptypb.Empty, error)
}

// UnimplementedOIDCProviderServer should be embedded to have forward compatible implementations.
type UnimplementedOIDCProviderServer struct {
}

func (UnimplementedOIDCProviderServer) CreateOIDCProvider(context.Context, *v3.OIDCProvider) (*v3.OIDCProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOIDCProvider not implemented")
}
func (UnimplementedOIDCProviderServer) GetOIDCProvider(context.Context, *v3.OIDCProvider) (*v3.OIDCProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOIDCProvider not implemented")
}
func (UnimplementedOIDCProviderServer) ListOIDCProvider(context.Context, *emptypb.Empty) (*v3.OIDCProviderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOIDCProvider not implemented")
}
func (UnimplementedOIDCProviderServer) UpdateOIDCProvider(context.Context, *v3.OIDCProvider) (*v3.OIDCProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOIDCProvider not implemented")
}
func (UnimplementedOIDCProviderServer) DeleteOIDCProvider(context.Context, *v3.OIDCProvider) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOIDCProvider not implemented")
}

// UnsafeOIDCProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OIDCProviderServer will
// result in compilation errors.
type UnsafeOIDCProviderServer interface {
	mustEmbedUnimplementedOIDCProviderServer()
}

func RegisterOIDCProviderServer(s grpc.ServiceRegistrar, srv OIDCProviderServer) {
	s.RegisterService(&OIDCProvider_ServiceDesc, srv)
}

func _OIDCProvider_CreateOIDCProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.OIDCProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCProviderServer).CreateOIDCProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.OIDCProvider/CreateOIDCProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCProviderServer).CreateOIDCProvider(ctx, req.(*v3.OIDCProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCProvider_GetOIDCProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.OIDCProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCProviderServer).GetOIDCProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.OIDCProvider/GetOIDCProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCProviderServer).GetOIDCProvider(ctx, req.(*v3.OIDCProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCProvider_ListOIDCProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCProviderServer).ListOIDCProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.OIDCProvider/ListOIDCProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCProviderServer).ListOIDCProvider(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCProvider_UpdateOIDCProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.OIDCProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCProviderServer).UpdateOIDCProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.OIDCProvider/UpdateOIDCProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCProviderServer).UpdateOIDCProvider(ctx, req.(*v3.OIDCProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCProvider_DeleteOIDCProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.OIDCProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCProviderServer).DeleteOIDCProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.OIDCProvider/DeleteOIDCProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCProviderServer).DeleteOIDCProvider(ctx, req.(*v3.OIDCProvider))
	}
	return interceptor(ctx, in, info, handler)
}

// OIDCProvider_ServiceDesc is the grpc.ServiceDesc for OIDCProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OIDCProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rafay.dev.rpc.v3.OIDCProvider",
	HandlerType: (*OIDCProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOIDCProvider",
			Handler:    _OIDCProvider_CreateOIDCProvider_Handler,
		},
		{
			MethodName: "GetOIDCProvider",
			Handler:    _OIDCProvider_GetOIDCProvider_Handler,
		},
		{
			MethodName: "ListOIDCProvider",
			Handler:    _OIDCProvider_ListOIDCProvider_Handler,
		},
		{
			MethodName: "UpdateOIDCProvider",
			Handler:    _OIDCProvider_UpdateOIDCProvider_Handler,
		},
		{
			MethodName: "DeleteOIDCProvider",
			Handler:    _OIDCProvider_DeleteOIDCProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/system/oidc_provider.proto",
}