// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: chat.proto

package __

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error)
	ReceiveStream(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (ChatService_ReceiveStreamClient, error)
	StreamConnectionCount(ctx context.Context, in *StreamConnectionCountRequest, opts ...grpc.CallOption) (*StreamConnectionCountResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (*ReceiveResponse, error) {
	out := new(ReceiveResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ReceiveStream(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (ChatService_ReceiveStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/chat.ChatService/ReceiveStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceReceiveStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ReceiveStreamClient interface {
	Recv() (*ReceiveResponse, error)
	grpc.ClientStream
}

type chatServiceReceiveStreamClient struct {
	grpc.ClientStream
}

func (x *chatServiceReceiveStreamClient) Recv() (*ReceiveResponse, error) {
	m := new(ReceiveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) StreamConnectionCount(ctx context.Context, in *StreamConnectionCountRequest, opts ...grpc.CallOption) (*StreamConnectionCountResponse, error) {
	out := new(StreamConnectionCountResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/StreamConnectionCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Send(context.Context, *SendRequest) (*SendResponse, error)
	Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error)
	ReceiveStream(*ReceiveRequest, ChatService_ReceiveStreamServer) error
	StreamConnectionCount(context.Context, *StreamConnectionCountRequest) (*StreamConnectionCountResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedChatServiceServer) Receive(context.Context, *ReceiveRequest) (*ReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedChatServiceServer) ReceiveStream(*ReceiveRequest, ChatService_ReceiveStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveStream not implemented")
}
func (UnimplementedChatServiceServer) StreamConnectionCount(context.Context, *StreamConnectionCountRequest) (*StreamConnectionCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamConnectionCount not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Receive(ctx, req.(*ReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ReceiveStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ReceiveStream(m, &chatServiceReceiveStreamServer{stream})
}

type ChatService_ReceiveStreamServer interface {
	Send(*ReceiveResponse) error
	grpc.ServerStream
}

type chatServiceReceiveStreamServer struct {
	grpc.ServerStream
}

func (x *chatServiceReceiveStreamServer) Send(m *ReceiveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_StreamConnectionCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamConnectionCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).StreamConnectionCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/StreamConnectionCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).StreamConnectionCount(ctx, req.(*StreamConnectionCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _ChatService_Send_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _ChatService_Receive_Handler,
		},
		{
			MethodName: "StreamConnectionCount",
			Handler:    _ChatService_StreamConnectionCount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveStream",
			Handler:       _ChatService_ReceiveStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}