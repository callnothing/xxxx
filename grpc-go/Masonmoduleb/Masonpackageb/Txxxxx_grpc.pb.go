// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Masonpackageb

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

// TxxxxxClient is the client API for Txxxxx service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxxxxxClient interface {
	Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error)
	Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type txxxxxClient struct {
	cc grpc.ClientConnInterface
}

func NewTxxxxxClient(cc grpc.ClientConnInterface) TxxxxxClient {
	return &txxxxxClient{cc}
}

func (c *txxxxxClient) Heathcheck(ctx context.Context, in *HealthcheckRequest, opts ...grpc.CallOption) (*HealthcheckResponse, error) {
	out := new(HealthcheckResponse)
	err := c.cc.Invoke(ctx, "/Masonmoduleb.Masonpackageb.Txxxxx/Heathcheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txxxxxClient) Helloworld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/Masonmoduleb.Masonpackageb.Txxxxx/Helloworld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TxxxxxServer is the server API for Txxxxx service.
// All implementations must embed UnimplementedTxxxxxServer
// for forward compatibility
type TxxxxxServer interface {
	Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error)
	Helloworld(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedTxxxxxServer()
}

// UnimplementedTxxxxxServer must be embedded to have forward compatible implementations.
type UnimplementedTxxxxxServer struct {
}

func (UnimplementedTxxxxxServer) Heathcheck(context.Context, *HealthcheckRequest) (*HealthcheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heathcheck not implemented")
}
func (UnimplementedTxxxxxServer) Helloworld(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Helloworld not implemented")
}
func (UnimplementedTxxxxxServer) mustEmbedUnimplementedTxxxxxServer() {}

// UnsafeTxxxxxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TxxxxxServer will
// result in compilation errors.
type UnsafeTxxxxxServer interface {
	mustEmbedUnimplementedTxxxxxServer()
}

func RegisterTxxxxxServer(s grpc.ServiceRegistrar, srv TxxxxxServer) {
	s.RegisterService(&Txxxxx_ServiceDesc, srv)
}

func _Txxxxx_Heathcheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthcheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxxxxxServer).Heathcheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Masonmoduleb.Masonpackageb.Txxxxx/Heathcheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxxxxxServer).Heathcheck(ctx, req.(*HealthcheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Txxxxx_Helloworld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxxxxxServer).Helloworld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Masonmoduleb.Masonpackageb.Txxxxx/Helloworld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxxxxxServer).Helloworld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Txxxxx_ServiceDesc is the grpc.ServiceDesc for Txxxxx service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Txxxxx_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Masonmoduleb.Masonpackageb.Txxxxx",
	HandlerType: (*TxxxxxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heathcheck",
			Handler:    _Txxxxx_Heathcheck_Handler,
		},
		{
			MethodName: "Helloworld",
			Handler:    _Txxxxx_Helloworld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Txxxxx.proto",
}
