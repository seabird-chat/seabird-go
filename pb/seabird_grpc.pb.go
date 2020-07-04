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

// SeabirdClient is the client API for Seabird service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeabirdClient interface {
	StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (Seabird_StreamEventsClient, error)
	// Chat actions
	PerformAction(ctx context.Context, in *PerformActionRequest, opts ...grpc.CallOption) (*PerformActionResponse, error)
	PerformPrivateAction(ctx context.Context, in *PerformPrivateActionRequest, opts ...grpc.CallOption) (*PerformPrivateActionResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	SendPrivateMessage(ctx context.Context, in *SendPrivateMessageRequest, opts ...grpc.CallOption) (*SendPrivateMessageResponse, error)
	JoinChannel(ctx context.Context, in *JoinChannelRequest, opts ...grpc.CallOption) (*JoinChannelResponse, error)
	LeaveChannel(ctx context.Context, in *LeaveChannelRequest, opts ...grpc.CallOption) (*LeaveChannelResponse, error)
	UpdateChannelInfo(ctx context.Context, in *UpdateChannelInfoRequest, opts ...grpc.CallOption) (*UpdateChannelInfoResponse, error)
	// Chat backend introspection
	ListBackends(ctx context.Context, in *ListBackendsRequest, opts ...grpc.CallOption) (*ListBackendsResponse, error)
	GetBackendInfo(ctx context.Context, in *BackendInfoRequest, opts ...grpc.CallOption) (*BackendInfoResponse, error)
	// Chat connection introspection
	ListChannels(ctx context.Context, in *ListChannelsRequest, opts ...grpc.CallOption) (*ListChannelsResponse, error)
	GetChannelInfo(ctx context.Context, in *ChannelInfoRequest, opts ...grpc.CallOption) (*ChannelInfoResponse, error)
	// Seabird introspection
	GetCoreInfo(ctx context.Context, in *CoreInfoRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error)
}

type seabirdClient struct {
	cc grpc.ClientConnInterface
}

func NewSeabirdClient(cc grpc.ClientConnInterface) SeabirdClient {
	return &seabirdClient{cc}
}

func (c *seabirdClient) StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (Seabird_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Seabird_serviceDesc.Streams[0], "/seabird.Seabird/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &seabirdStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Seabird_StreamEventsClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type seabirdStreamEventsClient struct {
	grpc.ClientStream
}

func (x *seabirdStreamEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seabirdClient) PerformAction(ctx context.Context, in *PerformActionRequest, opts ...grpc.CallOption) (*PerformActionResponse, error) {
	out := new(PerformActionResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/PerformAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) PerformPrivateAction(ctx context.Context, in *PerformPrivateActionRequest, opts ...grpc.CallOption) (*PerformPrivateActionResponse, error) {
	out := new(PerformPrivateActionResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/PerformPrivateAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) SendPrivateMessage(ctx context.Context, in *SendPrivateMessageRequest, opts ...grpc.CallOption) (*SendPrivateMessageResponse, error) {
	out := new(SendPrivateMessageResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/SendPrivateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) JoinChannel(ctx context.Context, in *JoinChannelRequest, opts ...grpc.CallOption) (*JoinChannelResponse, error) {
	out := new(JoinChannelResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/JoinChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) LeaveChannel(ctx context.Context, in *LeaveChannelRequest, opts ...grpc.CallOption) (*LeaveChannelResponse, error) {
	out := new(LeaveChannelResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/LeaveChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) UpdateChannelInfo(ctx context.Context, in *UpdateChannelInfoRequest, opts ...grpc.CallOption) (*UpdateChannelInfoResponse, error) {
	out := new(UpdateChannelInfoResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/UpdateChannelInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) ListBackends(ctx context.Context, in *ListBackendsRequest, opts ...grpc.CallOption) (*ListBackendsResponse, error) {
	out := new(ListBackendsResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/ListBackends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) GetBackendInfo(ctx context.Context, in *BackendInfoRequest, opts ...grpc.CallOption) (*BackendInfoResponse, error) {
	out := new(BackendInfoResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/GetBackendInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) ListChannels(ctx context.Context, in *ListChannelsRequest, opts ...grpc.CallOption) (*ListChannelsResponse, error) {
	out := new(ListChannelsResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/ListChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) GetChannelInfo(ctx context.Context, in *ChannelInfoRequest, opts ...grpc.CallOption) (*ChannelInfoResponse, error) {
	out := new(ChannelInfoResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/GetChannelInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seabirdClient) GetCoreInfo(ctx context.Context, in *CoreInfoRequest, opts ...grpc.CallOption) (*CoreInfoResponse, error) {
	out := new(CoreInfoResponse)
	err := c.cc.Invoke(ctx, "/seabird.Seabird/GetCoreInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeabirdServer is the server API for Seabird service.
// All implementations should embed UnimplementedSeabirdServer
// for forward compatibility
type SeabirdServer interface {
	StreamEvents(*StreamEventsRequest, Seabird_StreamEventsServer) error
	// Chat actions
	PerformAction(context.Context, *PerformActionRequest) (*PerformActionResponse, error)
	PerformPrivateAction(context.Context, *PerformPrivateActionRequest) (*PerformPrivateActionResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	SendPrivateMessage(context.Context, *SendPrivateMessageRequest) (*SendPrivateMessageResponse, error)
	JoinChannel(context.Context, *JoinChannelRequest) (*JoinChannelResponse, error)
	LeaveChannel(context.Context, *LeaveChannelRequest) (*LeaveChannelResponse, error)
	UpdateChannelInfo(context.Context, *UpdateChannelInfoRequest) (*UpdateChannelInfoResponse, error)
	// Chat backend introspection
	ListBackends(context.Context, *ListBackendsRequest) (*ListBackendsResponse, error)
	GetBackendInfo(context.Context, *BackendInfoRequest) (*BackendInfoResponse, error)
	// Chat connection introspection
	ListChannels(context.Context, *ListChannelsRequest) (*ListChannelsResponse, error)
	GetChannelInfo(context.Context, *ChannelInfoRequest) (*ChannelInfoResponse, error)
	// Seabird introspection
	GetCoreInfo(context.Context, *CoreInfoRequest) (*CoreInfoResponse, error)
}

// UnimplementedSeabirdServer should be embedded to have forward compatible implementations.
type UnimplementedSeabirdServer struct {
}

func (*UnimplementedSeabirdServer) StreamEvents(*StreamEventsRequest, Seabird_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (*UnimplementedSeabirdServer) PerformAction(context.Context, *PerformActionRequest) (*PerformActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformAction not implemented")
}
func (*UnimplementedSeabirdServer) PerformPrivateAction(context.Context, *PerformPrivateActionRequest) (*PerformPrivateActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformPrivateAction not implemented")
}
func (*UnimplementedSeabirdServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedSeabirdServer) SendPrivateMessage(context.Context, *SendPrivateMessageRequest) (*SendPrivateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPrivateMessage not implemented")
}
func (*UnimplementedSeabirdServer) JoinChannel(context.Context, *JoinChannelRequest) (*JoinChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinChannel not implemented")
}
func (*UnimplementedSeabirdServer) LeaveChannel(context.Context, *LeaveChannelRequest) (*LeaveChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveChannel not implemented")
}
func (*UnimplementedSeabirdServer) UpdateChannelInfo(context.Context, *UpdateChannelInfoRequest) (*UpdateChannelInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannelInfo not implemented")
}
func (*UnimplementedSeabirdServer) ListBackends(context.Context, *ListBackendsRequest) (*ListBackendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBackends not implemented")
}
func (*UnimplementedSeabirdServer) GetBackendInfo(context.Context, *BackendInfoRequest) (*BackendInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBackendInfo not implemented")
}
func (*UnimplementedSeabirdServer) ListChannels(context.Context, *ListChannelsRequest) (*ListChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannels not implemented")
}
func (*UnimplementedSeabirdServer) GetChannelInfo(context.Context, *ChannelInfoRequest) (*ChannelInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannelInfo not implemented")
}
func (*UnimplementedSeabirdServer) GetCoreInfo(context.Context, *CoreInfoRequest) (*CoreInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoreInfo not implemented")
}

func RegisterSeabirdServer(s *grpc.Server, srv SeabirdServer) {
	s.RegisterService(&_Seabird_serviceDesc, srv)
}

func _Seabird_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeabirdServer).StreamEvents(m, &seabirdStreamEventsServer{stream})
}

type Seabird_StreamEventsServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type seabirdStreamEventsServer struct {
	grpc.ServerStream
}

func (x *seabirdStreamEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Seabird_PerformAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).PerformAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/PerformAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).PerformAction(ctx, req.(*PerformActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_PerformPrivateAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformPrivateActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).PerformPrivateAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/PerformPrivateAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).PerformPrivateAction(ctx, req.(*PerformPrivateActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_SendPrivateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPrivateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).SendPrivateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/SendPrivateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).SendPrivateMessage(ctx, req.(*SendPrivateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_JoinChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).JoinChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/JoinChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).JoinChannel(ctx, req.(*JoinChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_LeaveChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).LeaveChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/LeaveChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).LeaveChannel(ctx, req.(*LeaveChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_UpdateChannelInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).UpdateChannelInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/UpdateChannelInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).UpdateChannelInfo(ctx, req.(*UpdateChannelInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_ListBackends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBackendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).ListBackends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/ListBackends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).ListBackends(ctx, req.(*ListBackendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_GetBackendInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackendInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).GetBackendInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/GetBackendInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).GetBackendInfo(ctx, req.(*BackendInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_ListChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).ListChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/ListChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).ListChannels(ctx, req.(*ListChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_GetChannelInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).GetChannelInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/GetChannelInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).GetChannelInfo(ctx, req.(*ChannelInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seabird_GetCoreInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeabirdServer).GetCoreInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seabird.Seabird/GetCoreInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeabirdServer).GetCoreInfo(ctx, req.(*CoreInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Seabird_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seabird.Seabird",
	HandlerType: (*SeabirdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PerformAction",
			Handler:    _Seabird_PerformAction_Handler,
		},
		{
			MethodName: "PerformPrivateAction",
			Handler:    _Seabird_PerformPrivateAction_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _Seabird_SendMessage_Handler,
		},
		{
			MethodName: "SendPrivateMessage",
			Handler:    _Seabird_SendPrivateMessage_Handler,
		},
		{
			MethodName: "JoinChannel",
			Handler:    _Seabird_JoinChannel_Handler,
		},
		{
			MethodName: "LeaveChannel",
			Handler:    _Seabird_LeaveChannel_Handler,
		},
		{
			MethodName: "UpdateChannelInfo",
			Handler:    _Seabird_UpdateChannelInfo_Handler,
		},
		{
			MethodName: "ListBackends",
			Handler:    _Seabird_ListBackends_Handler,
		},
		{
			MethodName: "GetBackendInfo",
			Handler:    _Seabird_GetBackendInfo_Handler,
		},
		{
			MethodName: "ListChannels",
			Handler:    _Seabird_ListChannels_Handler,
		},
		{
			MethodName: "GetChannelInfo",
			Handler:    _Seabird_GetChannelInfo_Handler,
		},
		{
			MethodName: "GetCoreInfo",
			Handler:    _Seabird_GetCoreInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _Seabird_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "seabird.proto",
}