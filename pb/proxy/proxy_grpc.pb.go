// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proxy

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

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyClient interface {
	GetActiveProxy(ctx context.Context, in *ByList, opts ...grpc.CallOption) (*ActiveProxy, error)
	GetActiveProxies(ctx context.Context, in *ByList, opts ...grpc.CallOption) (Proxy_GetActiveProxiesClient, error)
}

type proxyClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyClient(cc grpc.ClientConnInterface) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) GetActiveProxy(ctx context.Context, in *ByList, opts ...grpc.CallOption) (*ActiveProxy, error) {
	out := new(ActiveProxy)
	err := c.cc.Invoke(ctx, "/proxy.Proxy/GetActiveProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) GetActiveProxies(ctx context.Context, in *ByList, opts ...grpc.CallOption) (Proxy_GetActiveProxiesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Proxy_ServiceDesc.Streams[0], "/proxy.Proxy/GetActiveProxies", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyGetActiveProxiesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Proxy_GetActiveProxiesClient interface {
	Recv() (*ActiveProxy, error)
	grpc.ClientStream
}

type proxyGetActiveProxiesClient struct {
	grpc.ClientStream
}

func (x *proxyGetActiveProxiesClient) Recv() (*ActiveProxy, error) {
	m := new(ActiveProxy)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServer is the server API for Proxy service.
// All implementations must embed UnimplementedProxyServer
// for forward compatibility
type ProxyServer interface {
	GetActiveProxy(context.Context, *ByList) (*ActiveProxy, error)
	GetActiveProxies(*ByList, Proxy_GetActiveProxiesServer) error
	mustEmbedUnimplementedProxyServer()
}

// UnimplementedProxyServer must be embedded to have forward compatible implementations.
type UnimplementedProxyServer struct {
}

func (UnimplementedProxyServer) GetActiveProxy(context.Context, *ByList) (*ActiveProxy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveProxy not implemented")
}
func (UnimplementedProxyServer) GetActiveProxies(*ByList, Proxy_GetActiveProxiesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetActiveProxies not implemented")
}
func (UnimplementedProxyServer) mustEmbedUnimplementedProxyServer() {}

// UnsafeProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServer will
// result in compilation errors.
type UnsafeProxyServer interface {
	mustEmbedUnimplementedProxyServer()
}

func RegisterProxyServer(s grpc.ServiceRegistrar, srv ProxyServer) {
	s.RegisterService(&Proxy_ServiceDesc, srv)
}

func _Proxy_GetActiveProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).GetActiveProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proxy.Proxy/GetActiveProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).GetActiveProxy(ctx, req.(*ByList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_GetActiveProxies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ByList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServer).GetActiveProxies(m, &proxyGetActiveProxiesServer{stream})
}

type Proxy_GetActiveProxiesServer interface {
	Send(*ActiveProxy) error
	grpc.ServerStream
}

type proxyGetActiveProxiesServer struct {
	grpc.ServerStream
}

func (x *proxyGetActiveProxiesServer) Send(m *ActiveProxy) error {
	return x.ServerStream.SendMsg(m)
}

// Proxy_ServiceDesc is the grpc.ServiceDesc for Proxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Proxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActiveProxy",
			Handler:    _Proxy_GetActiveProxy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetActiveProxies",
			Handler:       _Proxy_GetActiveProxies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proxy.proto",
}

