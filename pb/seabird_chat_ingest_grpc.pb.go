// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatIngestClient is the client API for ChatIngest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatIngestClient interface {
	IngestEvents(ctx context.Context, opts ...grpc.CallOption) (ChatIngest_IngestEventsClient, error)
}

type chatIngestClient struct {
	cc grpc.ClientConnInterface
}

func NewChatIngestClient(cc grpc.ClientConnInterface) ChatIngestClient {
	return &chatIngestClient{cc}
}

func (c *chatIngestClient) IngestEvents(ctx context.Context, opts ...grpc.CallOption) (ChatIngest_IngestEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatIngest_serviceDesc.Streams[0], "/seabird.ChatIngest/IngestEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatIngestIngestEventsClient{stream}
	return x, nil
}

type ChatIngest_IngestEventsClient interface {
	Send(*ChatEvent) error
	Recv() (*ChatRequest, error)
	grpc.ClientStream
}

type chatIngestIngestEventsClient struct {
	grpc.ClientStream
}

func (x *chatIngestIngestEventsClient) Send(m *ChatEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatIngestIngestEventsClient) Recv() (*ChatRequest, error) {
	m := new(ChatRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatIngestServer is the server API for ChatIngest service.
// All implementations should embed UnimplementedChatIngestServer
// for forward compatibility
type ChatIngestServer interface {
	IngestEvents(ChatIngest_IngestEventsServer) error
}

// UnimplementedChatIngestServer should be embedded to have forward compatible implementations.
type UnimplementedChatIngestServer struct {
}

func (*UnimplementedChatIngestServer) IngestEvents(ChatIngest_IngestEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method IngestEvents not implemented")
}

func RegisterChatIngestServer(s *grpc.Server, srv ChatIngestServer) {
	s.RegisterService(&_ChatIngest_serviceDesc, srv)
}

func _ChatIngest_IngestEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatIngestServer).IngestEvents(&chatIngestIngestEventsServer{stream})
}

type ChatIngest_IngestEventsServer interface {
	Send(*ChatRequest) error
	Recv() (*ChatEvent, error)
	grpc.ServerStream
}

type chatIngestIngestEventsServer struct {
	grpc.ServerStream
}

func (x *chatIngestIngestEventsServer) Send(m *ChatRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatIngestIngestEventsServer) Recv() (*ChatEvent, error) {
	m := new(ChatEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatIngest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seabird.ChatIngest",
	HandlerType: (*ChatIngestServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "IngestEvents",
			Handler:       _ChatIngest_IngestEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "seabird_chat_ingest.proto",
}
