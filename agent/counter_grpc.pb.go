// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/agent/counter.proto

package agent

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

// CountingServiceClient is the client API for CountingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CountingServiceClient interface {
	CreateCounter(ctx context.Context, in *CreateCounterRequest, opts ...grpc.CallOption) (*Empty, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*Empty, error)
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*Empty, error)
	DumpTraffic(ctx context.Context, in *DumpTrafficRequest, opts ...grpc.CallOption) (*DumpTrafficResponse, error)
}

type countingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCountingServiceClient(cc grpc.ClientConnInterface) CountingServiceClient {
	return &countingServiceClient{cc}
}

func (c *countingServiceClient) CreateCounter(ctx context.Context, in *CreateCounterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.CountingService/CreateCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countingServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.CountingService/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countingServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.CountingService/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countingServiceClient) DumpTraffic(ctx context.Context, in *DumpTrafficRequest, opts ...grpc.CallOption) (*DumpTrafficResponse, error) {
	out := new(DumpTrafficResponse)
	err := c.cc.Invoke(ctx, "/proto.CountingService/DumpTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountingServiceServer is the server API for CountingService service.
// All implementations must embed UnimplementedCountingServiceServer
// for forward compatibility
type CountingServiceServer interface {
	CreateCounter(context.Context, *CreateCounterRequest) (*Empty, error)
	Subscribe(context.Context, *SubscribeRequest) (*Empty, error)
	Unsubscribe(context.Context, *UnsubscribeRequest) (*Empty, error)
	DumpTraffic(context.Context, *DumpTrafficRequest) (*DumpTrafficResponse, error)
	mustEmbedUnimplementedCountingServiceServer()
}

// UnimplementedCountingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCountingServiceServer struct {
}

func (UnimplementedCountingServiceServer) CreateCounter(context.Context, *CreateCounterRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCounter not implemented")
}
func (UnimplementedCountingServiceServer) Subscribe(context.Context, *SubscribeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedCountingServiceServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedCountingServiceServer) DumpTraffic(context.Context, *DumpTrafficRequest) (*DumpTrafficResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DumpTraffic not implemented")
}
func (UnimplementedCountingServiceServer) mustEmbedUnimplementedCountingServiceServer() {}

// UnsafeCountingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CountingServiceServer will
// result in compilation errors.
type UnsafeCountingServiceServer interface {
	mustEmbedUnimplementedCountingServiceServer()
}

func RegisterCountingServiceServer(s grpc.ServiceRegistrar, srv CountingServiceServer) {
	s.RegisterService(&CountingService_ServiceDesc, srv)
}

func _CountingService_CreateCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountingServiceServer).CreateCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CountingService/CreateCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountingServiceServer).CreateCounter(ctx, req.(*CreateCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountingService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountingServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CountingService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountingServiceServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountingService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountingServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CountingService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountingServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountingService_DumpTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DumpTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountingServiceServer).DumpTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CountingService/DumpTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountingServiceServer).DumpTraffic(ctx, req.(*DumpTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CountingService_ServiceDesc is the grpc.ServiceDesc for CountingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CountingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CountingService",
	HandlerType: (*CountingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCounter",
			Handler:    _CountingService_CreateCounter_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _CountingService_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _CountingService_Unsubscribe_Handler,
		},
		{
			MethodName: "DumpTraffic",
			Handler:    _CountingService_DumpTraffic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/agent/counter.proto",
}
