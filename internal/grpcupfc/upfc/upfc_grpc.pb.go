// Copyright 2015 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: upfc.proto

package upfc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UPFApp_SayHello_FullMethodName       = "/upfc.UPFApp/SayHello"
	UPFApp_SetBuffer_FullMethodName      = "/upfc.UPFApp/SetBuffer"
	UPFApp_AddUplinkPdr_FullMethodName   = "/upfc.UPFApp/AddUplinkPdr"
	UPFApp_AddUplinkFar_FullMethodName   = "/upfc.UPFApp/AddUplinkFar"
	UPFApp_AddDownlinkPdr_FullMethodName = "/upfc.UPFApp/AddDownlinkPdr"
	UPFApp_AddDownlinkFar_FullMethodName = "/upfc.UPFApp/AddDownlinkFar"
	UPFApp_AddRoute_FullMethodName       = "/upfc.UPFApp/AddRoute"
	UPFApp_AddQer_FullMethodName         = "/upfc.UPFApp/AddQer"
	UPFApp_AddSDFTemplate_FullMethodName = "/upfc.UPFApp/AddSDFTemplate"
)

// UPFAppClient is the client API for UPFApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The greeting service definition.
type UPFAppClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SetBuffer(ctx context.Context, in *BufferRequest, opts ...grpc.CallOption) (*RpcReply, error)
	AddUplinkPdr(ctx context.Context, in *UplinkPdrRequest, opts ...grpc.CallOption) (*RpcReply, error)
	AddUplinkFar(ctx context.Context, in *FarRequest, opts ...grpc.CallOption) (*RpcReply, error)
	AddDownlinkPdr(ctx context.Context, in *DownlinkPdrRequest, opts ...grpc.CallOption) (*RpcReply, error)
	AddDownlinkFar(ctx context.Context, in *FarRequest, opts ...grpc.CallOption) (*RpcReply, error)
	AddRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RpcReply, error)
	AddQer(ctx context.Context, in *QerRequest, opts ...grpc.CallOption) (*RpcReply, error)
	AddSDFTemplate(ctx context.Context, in *SDFTemplateRequest, opts ...grpc.CallOption) (*RpcReply, error)
}

type uPFAppClient struct {
	cc grpc.ClientConnInterface
}

func NewUPFAppClient(cc grpc.ClientConnInterface) UPFAppClient {
	return &uPFAppClient{cc}
}

func (c *uPFAppClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, UPFApp_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uPFAppClient) SetBuffer(ctx context.Context, in *BufferRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, UPFApp_SetBuffer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uPFAppClient) AddUplinkPdr(ctx context.Context, in *UplinkPdrRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, UPFApp_AddUplinkPdr_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uPFAppClient) AddUplinkFar(ctx context.Context, in *FarRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, UPFApp_AddUplinkFar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uPFAppClient) AddDownlinkPdr(ctx context.Context, in *DownlinkPdrRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, UPFApp_AddDownlinkPdr_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uPFAppClient) AddDownlinkFar(ctx context.Context, in *FarRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, UPFApp_AddDownlinkFar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uPFAppClient) AddRoute(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, UPFApp_AddRoute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uPFAppClient) AddQer(ctx context.Context, in *QerRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, UPFApp_AddQer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uPFAppClient) AddSDFTemplate(ctx context.Context, in *SDFTemplateRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, UPFApp_AddSDFTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UPFAppServer is the server API for UPFApp service.
// All implementations must embed UnimplementedUPFAppServer
// for forward compatibility.
//
// The greeting service definition.
type UPFAppServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SetBuffer(context.Context, *BufferRequest) (*RpcReply, error)
	AddUplinkPdr(context.Context, *UplinkPdrRequest) (*RpcReply, error)
	AddUplinkFar(context.Context, *FarRequest) (*RpcReply, error)
	AddDownlinkPdr(context.Context, *DownlinkPdrRequest) (*RpcReply, error)
	AddDownlinkFar(context.Context, *FarRequest) (*RpcReply, error)
	AddRoute(context.Context, *RouteRequest) (*RpcReply, error)
	AddQer(context.Context, *QerRequest) (*RpcReply, error)
	AddSDFTemplate(context.Context, *SDFTemplateRequest) (*RpcReply, error)
	mustEmbedUnimplementedUPFAppServer()
}

// UnimplementedUPFAppServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUPFAppServer struct{}

func (UnimplementedUPFAppServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedUPFAppServer) SetBuffer(context.Context, *BufferRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetBuffer not implemented")
}
func (UnimplementedUPFAppServer) AddUplinkPdr(context.Context, *UplinkPdrRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUplinkPdr not implemented")
}
func (UnimplementedUPFAppServer) AddUplinkFar(context.Context, *FarRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUplinkFar not implemented")
}
func (UnimplementedUPFAppServer) AddDownlinkPdr(context.Context, *DownlinkPdrRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDownlinkPdr not implemented")
}
func (UnimplementedUPFAppServer) AddDownlinkFar(context.Context, *FarRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDownlinkFar not implemented")
}
func (UnimplementedUPFAppServer) AddRoute(context.Context, *RouteRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRoute not implemented")
}
func (UnimplementedUPFAppServer) AddQer(context.Context, *QerRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddQer not implemented")
}
func (UnimplementedUPFAppServer) AddSDFTemplate(context.Context, *SDFTemplateRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSDFTemplate not implemented")
}
func (UnimplementedUPFAppServer) mustEmbedUnimplementedUPFAppServer() {}
func (UnimplementedUPFAppServer) testEmbeddedByValue()                {}

// UnsafeUPFAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UPFAppServer will
// result in compilation errors.
type UnsafeUPFAppServer interface {
	mustEmbedUnimplementedUPFAppServer()
}

func RegisterUPFAppServer(s grpc.ServiceRegistrar, srv UPFAppServer) {
	// If the following call pancis, it indicates UnimplementedUPFAppServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UPFApp_ServiceDesc, srv)
}

func _UPFApp_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UPFAppServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UPFApp_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UPFAppServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UPFApp_SetBuffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BufferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UPFAppServer).SetBuffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UPFApp_SetBuffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UPFAppServer).SetBuffer(ctx, req.(*BufferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UPFApp_AddUplinkPdr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UplinkPdrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UPFAppServer).AddUplinkPdr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UPFApp_AddUplinkPdr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UPFAppServer).AddUplinkPdr(ctx, req.(*UplinkPdrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UPFApp_AddUplinkFar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UPFAppServer).AddUplinkFar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UPFApp_AddUplinkFar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UPFAppServer).AddUplinkFar(ctx, req.(*FarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UPFApp_AddDownlinkPdr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkPdrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UPFAppServer).AddDownlinkPdr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UPFApp_AddDownlinkPdr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UPFAppServer).AddDownlinkPdr(ctx, req.(*DownlinkPdrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UPFApp_AddDownlinkFar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UPFAppServer).AddDownlinkFar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UPFApp_AddDownlinkFar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UPFAppServer).AddDownlinkFar(ctx, req.(*FarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UPFApp_AddRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UPFAppServer).AddRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UPFApp_AddRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UPFAppServer).AddRoute(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UPFApp_AddQer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UPFAppServer).AddQer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UPFApp_AddQer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UPFAppServer).AddQer(ctx, req.(*QerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UPFApp_AddSDFTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SDFTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UPFAppServer).AddSDFTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UPFApp_AddSDFTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UPFAppServer).AddSDFTemplate(ctx, req.(*SDFTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UPFApp_ServiceDesc is the grpc.ServiceDesc for UPFApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UPFApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "upfc.UPFApp",
	HandlerType: (*UPFAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _UPFApp_SayHello_Handler,
		},
		{
			MethodName: "SetBuffer",
			Handler:    _UPFApp_SetBuffer_Handler,
		},
		{
			MethodName: "AddUplinkPdr",
			Handler:    _UPFApp_AddUplinkPdr_Handler,
		},
		{
			MethodName: "AddUplinkFar",
			Handler:    _UPFApp_AddUplinkFar_Handler,
		},
		{
			MethodName: "AddDownlinkPdr",
			Handler:    _UPFApp_AddDownlinkPdr_Handler,
		},
		{
			MethodName: "AddDownlinkFar",
			Handler:    _UPFApp_AddDownlinkFar_Handler,
		},
		{
			MethodName: "AddRoute",
			Handler:    _UPFApp_AddRoute_Handler,
		},
		{
			MethodName: "AddQer",
			Handler:    _UPFApp_AddQer_Handler,
		},
		{
			MethodName: "AddSDFTemplate",
			Handler:    _UPFApp_AddSDFTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upfc.proto",
}
