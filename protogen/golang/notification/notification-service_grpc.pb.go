// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: notification-service.proto

package notification

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NotificationService_CreateNotification_FullMethodName        = "/notification.NotificationService/CreateNotification"
	NotificationService_GetAllNotifications_FullMethodName       = "/notification.NotificationService/GetAllNotifications"
	NotificationService_GetAllNotificationsByUser_FullMethodName = "/notification.NotificationService/GetAllNotificationsByUser"
	NotificationService_MarkAsRead_FullMethodName                = "/notification.NotificationService/MarkAsRead"
)

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error)
	GetAllNotifications(ctx context.Context, in *GetAllNotificationRequest, opts ...grpc.CallOption) (*GetAllNotificationsResponse, error)
	GetAllNotificationsByUser(ctx context.Context, in *GetAllNotificationsByUserRequest, opts ...grpc.CallOption) (*GetAllNotificationsByUserResponse, error)
	MarkAsRead(ctx context.Context, in *MarkAsReadNotificationRequest, opts ...grpc.CallOption) (*MarkAsReadNotificationResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) CreateNotification(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNotificationResponse)
	err := c.cc.Invoke(ctx, NotificationService_CreateNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetAllNotifications(ctx context.Context, in *GetAllNotificationRequest, opts ...grpc.CallOption) (*GetAllNotificationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllNotificationsResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetAllNotifications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) GetAllNotificationsByUser(ctx context.Context, in *GetAllNotificationsByUserRequest, opts ...grpc.CallOption) (*GetAllNotificationsByUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllNotificationsByUserResponse)
	err := c.cc.Invoke(ctx, NotificationService_GetAllNotificationsByUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) MarkAsRead(ctx context.Context, in *MarkAsReadNotificationRequest, opts ...grpc.CallOption) (*MarkAsReadNotificationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MarkAsReadNotificationResponse)
	err := c.cc.Invoke(ctx, NotificationService_MarkAsRead_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility.
type NotificationServiceServer interface {
	CreateNotification(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error)
	GetAllNotifications(context.Context, *GetAllNotificationRequest) (*GetAllNotificationsResponse, error)
	GetAllNotificationsByUser(context.Context, *GetAllNotificationsByUserRequest) (*GetAllNotificationsByUserResponse, error)
	MarkAsRead(context.Context, *MarkAsReadNotificationRequest) (*MarkAsReadNotificationResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotificationServiceServer struct{}

func (UnimplementedNotificationServiceServer) CreateNotification(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (UnimplementedNotificationServiceServer) GetAllNotifications(context.Context, *GetAllNotificationRequest) (*GetAllNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotifications not implemented")
}
func (UnimplementedNotificationServiceServer) GetAllNotificationsByUser(context.Context, *GetAllNotificationsByUserRequest) (*GetAllNotificationsByUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotificationsByUser not implemented")
}
func (UnimplementedNotificationServiceServer) MarkAsRead(context.Context, *MarkAsReadNotificationRequest) (*MarkAsReadNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsRead not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}
func (UnimplementedNotificationServiceServer) testEmbeddedByValue()                             {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	// If the following call pancis, it indicates UnimplementedNotificationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_CreateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).CreateNotification(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetAllNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetAllNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetAllNotifications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetAllNotifications(ctx, req.(*GetAllNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetAllNotificationsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllNotificationsByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetAllNotificationsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_GetAllNotificationsByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetAllNotificationsByUser(ctx, req.(*GetAllNotificationsByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_MarkAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsReadNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).MarkAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotificationService_MarkAsRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).MarkAsRead(ctx, req.(*MarkAsReadNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNotification",
			Handler:    _NotificationService_CreateNotification_Handler,
		},
		{
			MethodName: "GetAllNotifications",
			Handler:    _NotificationService_GetAllNotifications_Handler,
		},
		{
			MethodName: "GetAllNotificationsByUser",
			Handler:    _NotificationService_GetAllNotificationsByUser_Handler,
		},
		{
			MethodName: "MarkAsRead",
			Handler:    _NotificationService_MarkAsRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification-service.proto",
}
