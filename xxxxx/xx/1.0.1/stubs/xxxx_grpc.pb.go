// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package xx

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

// XxxxClient is the client API for Xxxx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XxxxClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type xxxxClient struct {
	cc grpc.ClientConnInterface
}

func NewXxxxClient(cc grpc.ClientConnInterface) XxxxClient {
	return &xxxxClient{cc}
}

func (c *xxxxClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/xxxxx.xx.xxxx/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XxxxServer is the server API for Xxxx service.
// All implementations must embed UnimplementedXxxxServer
// for forward compatibility
type XxxxServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedXxxxServer()
}

// UnimplementedXxxxServer must be embedded to have forward compatible implementations.
type UnimplementedXxxxServer struct {
}

func (UnimplementedXxxxServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedXxxxServer) mustEmbedUnimplementedXxxxServer() {}

// UnsafeXxxxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XxxxServer will
// result in compilation errors.
type UnsafeXxxxServer interface {
	mustEmbedUnimplementedXxxxServer()
}

func RegisterXxxxServer(s grpc.ServiceRegistrar, srv XxxxServer) {
	s.RegisterService(&Xxxx_ServiceDesc, srv)
}

func _Xxxx_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XxxxServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxxxx.xx.xxxx/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XxxxServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Xxxx_ServiceDesc is the grpc.ServiceDesc for Xxxx service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Xxxx_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xxxxx.xx.xxxx",
	HandlerType: (*XxxxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Xxxx_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "xxxx.proto",
}
