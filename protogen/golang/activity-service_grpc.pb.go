// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: activity-service.proto

package activity

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
	ActivityService_CreateActivity_FullMethodName              = "/activity.ActivityService/CreateActivity"
	ActivityService_GetAllActivitiesByManagerId_FullMethodName = "/activity.ActivityService/GetAllActivitiesByManagerId"
)

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityServiceClient interface {
	CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityResponse, error)
	GetAllActivitiesByManagerId(ctx context.Context, in *GetAllActivitiesByManagerIdRequest, opts ...grpc.CallOption) (*GetAllActivitiesByManagerIdResponse, error)
}

type activityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityServiceClient(cc grpc.ClientConnInterface) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateActivityResponse)
	err := c.cc.Invoke(ctx, ActivityService_CreateActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) GetAllActivitiesByManagerId(ctx context.Context, in *GetAllActivitiesByManagerIdRequest, opts ...grpc.CallOption) (*GetAllActivitiesByManagerIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllActivitiesByManagerIdResponse)
	err := c.cc.Invoke(ctx, ActivityService_GetAllActivitiesByManagerId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServiceServer is the server API for ActivityService service.
// All implementations must embed UnimplementedActivityServiceServer
// for forward compatibility.
type ActivityServiceServer interface {
	CreateActivity(context.Context, *CreateActivityRequest) (*CreateActivityResponse, error)
	GetAllActivitiesByManagerId(context.Context, *GetAllActivitiesByManagerIdRequest) (*GetAllActivitiesByManagerIdResponse, error)
	mustEmbedUnimplementedActivityServiceServer()
}

// UnimplementedActivityServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActivityServiceServer struct{}

func (UnimplementedActivityServiceServer) CreateActivity(context.Context, *CreateActivityRequest) (*CreateActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (UnimplementedActivityServiceServer) GetAllActivitiesByManagerId(context.Context, *GetAllActivitiesByManagerIdRequest) (*GetAllActivitiesByManagerIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllActivitiesByManagerId not implemented")
}
func (UnimplementedActivityServiceServer) mustEmbedUnimplementedActivityServiceServer() {}
func (UnimplementedActivityServiceServer) testEmbeddedByValue()                         {}

// UnsafeActivityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServiceServer will
// result in compilation errors.
type UnsafeActivityServiceServer interface {
	mustEmbedUnimplementedActivityServiceServer()
}

func RegisterActivityServiceServer(s grpc.ServiceRegistrar, srv ActivityServiceServer) {
	// If the following call pancis, it indicates UnimplementedActivityServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ActivityService_ServiceDesc, srv)
}

func _ActivityService_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_CreateActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).CreateActivity(ctx, req.(*CreateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_GetAllActivitiesByManagerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllActivitiesByManagerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).GetAllActivitiesByManagerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_GetAllActivitiesByManagerId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).GetAllActivitiesByManagerId(ctx, req.(*GetAllActivitiesByManagerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityService_ServiceDesc is the grpc.ServiceDesc for ActivityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activity.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateActivity",
			Handler:    _ActivityService_CreateActivity_Handler,
		},
		{
			MethodName: "GetAllActivitiesByManagerId",
			Handler:    _ActivityService_GetAllActivitiesByManagerId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activity-service.proto",
}
