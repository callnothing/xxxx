// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package AAA

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

// XXXXXClient is the client API for XXXXX service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XXXXXClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type xXXXXClient struct {
	cc grpc.ClientConnInterface
}

func NewXXXXXClient(cc grpc.ClientConnInterface) XXXXXClient {
	return &xXXXXClient{cc}
}

func (c *xXXXXClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/xxxx.AAA.XXXXX/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XXXXXServer is the server API for XXXXX service.
// All implementations must embed UnimplementedXXXXXServer
// for forward compatibility
type XXXXXServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedXXXXXServer()
}

// UnimplementedXXXXXServer must be embedded to have forward compatible implementations.
type UnimplementedXXXXXServer struct {
}

func (UnimplementedXXXXXServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedXXXXXServer) mustEmbedUnimplementedXXXXXServer() {}

// UnsafeXXXXXServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XXXXXServer will
// result in compilation errors.
type UnsafeXXXXXServer interface {
	mustEmbedUnimplementedXXXXXServer()
}

func RegisterXXXXXServer(s grpc.ServiceRegistrar, srv XXXXXServer) {
	s.RegisterService(&XXXXX_ServiceDesc, srv)
}

func _XXXXX_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XXXXXServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xxxx.AAA.XXXXX/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XXXXXServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// XXXXX_ServiceDesc is the grpc.ServiceDesc for XXXXX service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var XXXXX_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xxxx.AAA.XXXXX",
	HandlerType: (*XXXXXServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _XXXXX_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "XXXXX.proto",
}
