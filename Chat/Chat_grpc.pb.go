// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Chat

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

// ChittychatClient is the client API for Chittychat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittychatClient interface {
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
}

type chittychatClient struct {
	cc grpc.ClientConnInterface
}

func NewChittychatClient(cc grpc.ClientConnInterface) ChittychatClient {
	return &chittychatClient{cc}
}

func (c *chittychatClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, "/Chat.Chittychat/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/Chat.Chittychat/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatClient) Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/Chat.Chittychat/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittychatServer is the server API for Chittychat service.
// All implementations must embed UnimplementedChittychatServer
// for forward compatibility
type ChittychatServer interface {
	Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error)
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	Connect(context.Context, *ConnectionRequest) (*ConnectionResponse, error)
	mustEmbedUnimplementedChittychatServer()
}

// UnimplementedChittychatServer must be embedded to have forward compatible implementations.
type UnimplementedChittychatServer struct {
}

func (UnimplementedChittychatServer) Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedChittychatServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedChittychatServer) Connect(context.Context, *ConnectionRequest) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChittychatServer) mustEmbedUnimplementedChittychatServer() {}

// UnsafeChittychatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittychatServer will
// result in compilation errors.
type UnsafeChittychatServer interface {
	mustEmbedUnimplementedChittychatServer()
}

func RegisterChittychatServer(s grpc.ServiceRegistrar, srv ChittychatServer) {
	s.RegisterService(&Chittychat_ServiceDesc, srv)
}

func _Chittychat_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat.Chittychat/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServer).Broadcast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chittychat_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat.Chittychat/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chittychat_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chat.Chittychat/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServer).Connect(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chittychat_ServiceDesc is the grpc.ServiceDesc for Chittychat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chittychat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Chat.Chittychat",
	HandlerType: (*ChittychatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _Chittychat_Broadcast_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Chittychat_Publish_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Chittychat_Connect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Chat/Chat.proto",
}
