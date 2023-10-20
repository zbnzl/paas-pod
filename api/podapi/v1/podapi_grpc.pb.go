// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: podapi/v1/podapi.proto

package v1

import (
	context "context"
	v1 "github.com/zbnzl/paas-pod/api/pod/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PodApi_FindPodById_FullMethodName   = "/api.podApi.v1.PodApi/FindPodById"
	PodApi_AddPod_FullMethodName        = "/api.podApi.v1.PodApi/AddPod"
	PodApi_DeletePodById_FullMethodName = "/api.podApi.v1.PodApi/DeletePodById"
	PodApi_UpdatePod_FullMethodName     = "/api.podApi.v1.PodApi/UpdatePod"
	PodApi_Call_FullMethodName          = "/api.podApi.v1.PodApi/Call"
)

// PodApiClient is the client API for PodApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PodApiClient interface {
	FindPodById(ctx context.Context, in *v1.PodId, opts ...grpc.CallOption) (*v1.PodInfo, error)
	AddPod(ctx context.Context, in *v1.PodInfo, opts ...grpc.CallOption) (*v1.Response, error)
	DeletePodById(ctx context.Context, in *v1.PodId, opts ...grpc.CallOption) (*v1.Response, error)
	UpdatePod(ctx context.Context, in *v1.PodInfo, opts ...grpc.CallOption) (*v1.Response, error)
	// 默认接口
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type podApiClient struct {
	cc grpc.ClientConnInterface
}

func NewPodApiClient(cc grpc.ClientConnInterface) PodApiClient {
	return &podApiClient{cc}
}

func (c *podApiClient) FindPodById(ctx context.Context, in *v1.PodId, opts ...grpc.CallOption) (*v1.PodInfo, error) {
	out := new(v1.PodInfo)
	err := c.cc.Invoke(ctx, PodApi_FindPodById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podApiClient) AddPod(ctx context.Context, in *v1.PodInfo, opts ...grpc.CallOption) (*v1.Response, error) {
	out := new(v1.Response)
	err := c.cc.Invoke(ctx, PodApi_AddPod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podApiClient) DeletePodById(ctx context.Context, in *v1.PodId, opts ...grpc.CallOption) (*v1.Response, error) {
	out := new(v1.Response)
	err := c.cc.Invoke(ctx, PodApi_DeletePodById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podApiClient) UpdatePod(ctx context.Context, in *v1.PodInfo, opts ...grpc.CallOption) (*v1.Response, error) {
	out := new(v1.Response)
	err := c.cc.Invoke(ctx, PodApi_UpdatePod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podApiClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, PodApi_Call_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodApiServer is the server API for PodApi service.
// All implementations must embed UnimplementedPodApiServer
// for forward compatibility
type PodApiServer interface {
	FindPodById(context.Context, *v1.PodId) (*v1.PodInfo, error)
	AddPod(context.Context, *v1.PodInfo) (*v1.Response, error)
	DeletePodById(context.Context, *v1.PodId) (*v1.Response, error)
	UpdatePod(context.Context, *v1.PodInfo) (*v1.Response, error)
	// 默认接口
	Call(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedPodApiServer()
}

// UnimplementedPodApiServer must be embedded to have forward compatible implementations.
type UnimplementedPodApiServer struct {
}

func (UnimplementedPodApiServer) FindPodById(context.Context, *v1.PodId) (*v1.PodInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPodById not implemented")
}
func (UnimplementedPodApiServer) AddPod(context.Context, *v1.PodInfo) (*v1.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPod not implemented")
}
func (UnimplementedPodApiServer) DeletePodById(context.Context, *v1.PodId) (*v1.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePodById not implemented")
}
func (UnimplementedPodApiServer) UpdatePod(context.Context, *v1.PodInfo) (*v1.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePod not implemented")
}
func (UnimplementedPodApiServer) Call(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedPodApiServer) mustEmbedUnimplementedPodApiServer() {}

// UnsafePodApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PodApiServer will
// result in compilation errors.
type UnsafePodApiServer interface {
	mustEmbedUnimplementedPodApiServer()
}

func RegisterPodApiServer(s grpc.ServiceRegistrar, srv PodApiServer) {
	s.RegisterService(&PodApi_ServiceDesc, srv)
}

func _PodApi_FindPodById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PodId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodApiServer).FindPodById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PodApi_FindPodById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodApiServer).FindPodById(ctx, req.(*v1.PodId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodApi_AddPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PodInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodApiServer).AddPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PodApi_AddPod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodApiServer).AddPod(ctx, req.(*v1.PodInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodApi_DeletePodById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PodId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodApiServer).DeletePodById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PodApi_DeletePodById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodApiServer).DeletePodById(ctx, req.(*v1.PodId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodApi_UpdatePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PodInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodApiServer).UpdatePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PodApi_UpdatePod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodApiServer).UpdatePod(ctx, req.(*v1.PodInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PodApi_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodApiServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PodApi_Call_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodApiServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PodApi_ServiceDesc is the grpc.ServiceDesc for PodApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PodApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.podApi.v1.PodApi",
	HandlerType: (*PodApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindPodById",
			Handler:    _PodApi_FindPodById_Handler,
		},
		{
			MethodName: "AddPod",
			Handler:    _PodApi_AddPod_Handler,
		},
		{
			MethodName: "DeletePodById",
			Handler:    _PodApi_DeletePodById_Handler,
		},
		{
			MethodName: "UpdatePod",
			Handler:    _PodApi_UpdatePod_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _PodApi_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "podapi/v1/podapi.proto",
}
