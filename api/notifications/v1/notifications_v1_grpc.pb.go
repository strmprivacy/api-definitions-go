// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: strmprivacy/api/notifications/v1/notifications_v1.proto

package notifications

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

// NotificationsServiceClient is the client API for NotificationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationsServiceClient interface {
	// Tell the notification service that the given user should receive notifications about the given entity.
	SubscribeNotifications(ctx context.Context, in *SubscribeNotificationsRequest, opts ...grpc.CallOption) (*SubscribeNotificationsResponse, error)
	// Tell the notification service that the given user should *not* receive notifications about the given entity.
	UnSubscribeNotifications(ctx context.Context, in *UnSubscribeNotificationsRequest, opts ...grpc.CallOption) (*UnSubscribeNotificationsResponse, error)
	// List the subscriptions of a certain user.
	ListSubscriptions(ctx context.Context, in *ListSubscriptionsRequest, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error)
	// Called by systems (schema-registry for instance) that want to notify subscribed users about something.
	NotifySubscribers(ctx context.Context, in *NotifySubscribersRequest, opts ...grpc.CallOption) (*NotifySubscribersResponse, error)
	// Notify specific users directly, regardless of their subscriptions.
	NotifyUsers(ctx context.Context, in *NotifyUsersRequest, opts ...grpc.CallOption) (*NotifyUsersResponse, error)
	// Called by the user's browser to receive notifications.
	ReceiveNotifications(ctx context.Context, in *ReceiveNotificationsRequest, opts ...grpc.CallOption) (NotificationsService_ReceiveNotificationsClient, error)
	// List notifications for a user.
	ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error)
	// set the has_been_seen status of notification(s)
	AlterNotificationsHasBeenSeen(ctx context.Context, in *AlterNotificationsHasBeenSeenRequest, opts ...grpc.CallOption) (*AlterNotificationsHasBeenSeenResponse, error)
}

type notificationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsServiceClient(cc grpc.ClientConnInterface) NotificationsServiceClient {
	return &notificationsServiceClient{cc}
}

func (c *notificationsServiceClient) SubscribeNotifications(ctx context.Context, in *SubscribeNotificationsRequest, opts ...grpc.CallOption) (*SubscribeNotificationsResponse, error) {
	out := new(SubscribeNotificationsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.notifications.v1.NotificationsService/SubscribeNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) UnSubscribeNotifications(ctx context.Context, in *UnSubscribeNotificationsRequest, opts ...grpc.CallOption) (*UnSubscribeNotificationsResponse, error) {
	out := new(UnSubscribeNotificationsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.notifications.v1.NotificationsService/UnSubscribeNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) ListSubscriptions(ctx context.Context, in *ListSubscriptionsRequest, opts ...grpc.CallOption) (*ListSubscriptionsResponse, error) {
	out := new(ListSubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.notifications.v1.NotificationsService/ListSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) NotifySubscribers(ctx context.Context, in *NotifySubscribersRequest, opts ...grpc.CallOption) (*NotifySubscribersResponse, error) {
	out := new(NotifySubscribersResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.notifications.v1.NotificationsService/NotifySubscribers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) NotifyUsers(ctx context.Context, in *NotifyUsersRequest, opts ...grpc.CallOption) (*NotifyUsersResponse, error) {
	out := new(NotifyUsersResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.notifications.v1.NotificationsService/NotifyUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) ReceiveNotifications(ctx context.Context, in *ReceiveNotificationsRequest, opts ...grpc.CallOption) (NotificationsService_ReceiveNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotificationsService_ServiceDesc.Streams[0], "/strmprivacy.api.notifications.v1.NotificationsService/ReceiveNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationsServiceReceiveNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotificationsService_ReceiveNotificationsClient interface {
	Recv() (*ReceiveNotificationsResponse, error)
	grpc.ClientStream
}

type notificationsServiceReceiveNotificationsClient struct {
	grpc.ClientStream
}

func (x *notificationsServiceReceiveNotificationsClient) Recv() (*ReceiveNotificationsResponse, error) {
	m := new(ReceiveNotificationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notificationsServiceClient) ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error) {
	out := new(ListNotificationsResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.notifications.v1.NotificationsService/ListNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsServiceClient) AlterNotificationsHasBeenSeen(ctx context.Context, in *AlterNotificationsHasBeenSeenRequest, opts ...grpc.CallOption) (*AlterNotificationsHasBeenSeenResponse, error) {
	out := new(AlterNotificationsHasBeenSeenResponse)
	err := c.cc.Invoke(ctx, "/strmprivacy.api.notifications.v1.NotificationsService/AlterNotificationsHasBeenSeen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServiceServer is the server API for NotificationsService service.
// All implementations should embed UnimplementedNotificationsServiceServer
// for forward compatibility
type NotificationsServiceServer interface {
	// Tell the notification service that the given user should receive notifications about the given entity.
	SubscribeNotifications(context.Context, *SubscribeNotificationsRequest) (*SubscribeNotificationsResponse, error)
	// Tell the notification service that the given user should *not* receive notifications about the given entity.
	UnSubscribeNotifications(context.Context, *UnSubscribeNotificationsRequest) (*UnSubscribeNotificationsResponse, error)
	// List the subscriptions of a certain user.
	ListSubscriptions(context.Context, *ListSubscriptionsRequest) (*ListSubscriptionsResponse, error)
	// Called by systems (schema-registry for instance) that want to notify subscribed users about something.
	NotifySubscribers(context.Context, *NotifySubscribersRequest) (*NotifySubscribersResponse, error)
	// Notify specific users directly, regardless of their subscriptions.
	NotifyUsers(context.Context, *NotifyUsersRequest) (*NotifyUsersResponse, error)
	// Called by the user's browser to receive notifications.
	ReceiveNotifications(*ReceiveNotificationsRequest, NotificationsService_ReceiveNotificationsServer) error
	// List notifications for a user.
	ListNotifications(context.Context, *ListNotificationsRequest) (*ListNotificationsResponse, error)
	// set the has_been_seen status of notification(s)
	AlterNotificationsHasBeenSeen(context.Context, *AlterNotificationsHasBeenSeenRequest) (*AlterNotificationsHasBeenSeenResponse, error)
}

// UnimplementedNotificationsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNotificationsServiceServer struct {
}

func (UnimplementedNotificationsServiceServer) SubscribeNotifications(context.Context, *SubscribeNotificationsRequest) (*SubscribeNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeNotifications not implemented")
}
func (UnimplementedNotificationsServiceServer) UnSubscribeNotifications(context.Context, *UnSubscribeNotificationsRequest) (*UnSubscribeNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSubscribeNotifications not implemented")
}
func (UnimplementedNotificationsServiceServer) ListSubscriptions(context.Context, *ListSubscriptionsRequest) (*ListSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubscriptions not implemented")
}
func (UnimplementedNotificationsServiceServer) NotifySubscribers(context.Context, *NotifySubscribersRequest) (*NotifySubscribersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifySubscribers not implemented")
}
func (UnimplementedNotificationsServiceServer) NotifyUsers(context.Context, *NotifyUsersRequest) (*NotifyUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyUsers not implemented")
}
func (UnimplementedNotificationsServiceServer) ReceiveNotifications(*ReceiveNotificationsRequest, NotificationsService_ReceiveNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveNotifications not implemented")
}
func (UnimplementedNotificationsServiceServer) ListNotifications(context.Context, *ListNotificationsRequest) (*ListNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotifications not implemented")
}
func (UnimplementedNotificationsServiceServer) AlterNotificationsHasBeenSeen(context.Context, *AlterNotificationsHasBeenSeenRequest) (*AlterNotificationsHasBeenSeenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterNotificationsHasBeenSeen not implemented")
}

// UnsafeNotificationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServiceServer will
// result in compilation errors.
type UnsafeNotificationsServiceServer interface {
	mustEmbedUnimplementedNotificationsServiceServer()
}

func RegisterNotificationsServiceServer(s grpc.ServiceRegistrar, srv NotificationsServiceServer) {
	s.RegisterService(&NotificationsService_ServiceDesc, srv)
}

func _NotificationsService_SubscribeNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).SubscribeNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.notifications.v1.NotificationsService/SubscribeNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).SubscribeNotifications(ctx, req.(*SubscribeNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_UnSubscribeNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSubscribeNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).UnSubscribeNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.notifications.v1.NotificationsService/UnSubscribeNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).UnSubscribeNotifications(ctx, req.(*UnSubscribeNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_ListSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubscriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).ListSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.notifications.v1.NotificationsService/ListSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).ListSubscriptions(ctx, req.(*ListSubscriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_NotifySubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifySubscribersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).NotifySubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.notifications.v1.NotificationsService/NotifySubscribers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).NotifySubscribers(ctx, req.(*NotifySubscribersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_NotifyUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).NotifyUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.notifications.v1.NotificationsService/NotifyUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).NotifyUsers(ctx, req.(*NotifyUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_ReceiveNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveNotificationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationsServiceServer).ReceiveNotifications(m, &notificationsServiceReceiveNotificationsServer{stream})
}

type NotificationsService_ReceiveNotificationsServer interface {
	Send(*ReceiveNotificationsResponse) error
	grpc.ServerStream
}

type notificationsServiceReceiveNotificationsServer struct {
	grpc.ServerStream
}

func (x *notificationsServiceReceiveNotificationsServer) Send(m *ReceiveNotificationsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _NotificationsService_ListNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).ListNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.notifications.v1.NotificationsService/ListNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).ListNotifications(ctx, req.(*ListNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationsService_AlterNotificationsHasBeenSeen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlterNotificationsHasBeenSeenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).AlterNotificationsHasBeenSeen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strmprivacy.api.notifications.v1.NotificationsService/AlterNotificationsHasBeenSeen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).AlterNotificationsHasBeenSeen(ctx, req.(*AlterNotificationsHasBeenSeenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationsService_ServiceDesc is the grpc.ServiceDesc for NotificationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "strmprivacy.api.notifications.v1.NotificationsService",
	HandlerType: (*NotificationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubscribeNotifications",
			Handler:    _NotificationsService_SubscribeNotifications_Handler,
		},
		{
			MethodName: "UnSubscribeNotifications",
			Handler:    _NotificationsService_UnSubscribeNotifications_Handler,
		},
		{
			MethodName: "ListSubscriptions",
			Handler:    _NotificationsService_ListSubscriptions_Handler,
		},
		{
			MethodName: "NotifySubscribers",
			Handler:    _NotificationsService_NotifySubscribers_Handler,
		},
		{
			MethodName: "NotifyUsers",
			Handler:    _NotificationsService_NotifyUsers_Handler,
		},
		{
			MethodName: "ListNotifications",
			Handler:    _NotificationsService_ListNotifications_Handler,
		},
		{
			MethodName: "AlterNotificationsHasBeenSeen",
			Handler:    _NotificationsService_AlterNotificationsHasBeenSeen_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveNotifications",
			Handler:       _NotificationsService_ReceiveNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "strmprivacy/api/notifications/v1/notifications_v1.proto",
}
