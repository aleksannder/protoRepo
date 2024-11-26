// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: workflow-service.proto

package workflow

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
	WorkflowService_GetWorkflowById_FullMethodName         = "/workflow.WorkflowService/GetWorkflowById"
	WorkflowService_InsertWorkflow_FullMethodName          = "/workflow.WorkflowService/InsertWorkflow"
	WorkflowService_GetWorkflowsByProjectId_FullMethodName = "/workflow.WorkflowService/GetWorkflowsByProjectId"
)

// WorkflowServiceClient is the client API for WorkflowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowServiceClient interface {
	GetWorkflowById(ctx context.Context, in *GetWorkflowByIdRequest, opts ...grpc.CallOption) (*GetWorkflowByIdResponse, error)
	InsertWorkflow(ctx context.Context, in *InsertWorkflowRequest, opts ...grpc.CallOption) (*InsertWorkflowResponse, error)
	GetWorkflowsByProjectId(ctx context.Context, in *GetWorkflowsByProjectIdRequest, opts ...grpc.CallOption) (*GetWorkflowsByProjectIdResponse, error)
}

type workflowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowServiceClient(cc grpc.ClientConnInterface) WorkflowServiceClient {
	return &workflowServiceClient{cc}
}

func (c *workflowServiceClient) GetWorkflowById(ctx context.Context, in *GetWorkflowByIdRequest, opts ...grpc.CallOption) (*GetWorkflowByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWorkflowByIdResponse)
	err := c.cc.Invoke(ctx, WorkflowService_GetWorkflowById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) InsertWorkflow(ctx context.Context, in *InsertWorkflowRequest, opts ...grpc.CallOption) (*InsertWorkflowResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InsertWorkflowResponse)
	err := c.cc.Invoke(ctx, WorkflowService_InsertWorkflow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) GetWorkflowsByProjectId(ctx context.Context, in *GetWorkflowsByProjectIdRequest, opts ...grpc.CallOption) (*GetWorkflowsByProjectIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWorkflowsByProjectIdResponse)
	err := c.cc.Invoke(ctx, WorkflowService_GetWorkflowsByProjectId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServiceServer is the server API for WorkflowService service.
// All implementations must embed UnimplementedWorkflowServiceServer
// for forward compatibility.
type WorkflowServiceServer interface {
	GetWorkflowById(context.Context, *GetWorkflowByIdRequest) (*GetWorkflowByIdResponse, error)
	InsertWorkflow(context.Context, *InsertWorkflowRequest) (*InsertWorkflowResponse, error)
	GetWorkflowsByProjectId(context.Context, *GetWorkflowsByProjectIdRequest) (*GetWorkflowsByProjectIdResponse, error)
	mustEmbedUnimplementedWorkflowServiceServer()
}

// UnimplementedWorkflowServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWorkflowServiceServer struct{}

func (UnimplementedWorkflowServiceServer) GetWorkflowById(context.Context, *GetWorkflowByIdRequest) (*GetWorkflowByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowById not implemented")
}
func (UnimplementedWorkflowServiceServer) InsertWorkflow(context.Context, *InsertWorkflowRequest) (*InsertWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertWorkflow not implemented")
}
func (UnimplementedWorkflowServiceServer) GetWorkflowsByProjectId(context.Context, *GetWorkflowsByProjectIdRequest) (*GetWorkflowsByProjectIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowsByProjectId not implemented")
}
func (UnimplementedWorkflowServiceServer) mustEmbedUnimplementedWorkflowServiceServer() {}
func (UnimplementedWorkflowServiceServer) testEmbeddedByValue()                         {}

// UnsafeWorkflowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowServiceServer will
// result in compilation errors.
type UnsafeWorkflowServiceServer interface {
	mustEmbedUnimplementedWorkflowServiceServer()
}

func RegisterWorkflowServiceServer(s grpc.ServiceRegistrar, srv WorkflowServiceServer) {
	// If the following call pancis, it indicates UnimplementedWorkflowServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WorkflowService_ServiceDesc, srv)
}

func _WorkflowService_GetWorkflowById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).GetWorkflowById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowService_GetWorkflowById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).GetWorkflowById(ctx, req.(*GetWorkflowByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_InsertWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).InsertWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowService_InsertWorkflow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).InsertWorkflow(ctx, req.(*InsertWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_GetWorkflowsByProjectId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowsByProjectIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).GetWorkflowsByProjectId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WorkflowService_GetWorkflowsByProjectId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).GetWorkflowsByProjectId(ctx, req.(*GetWorkflowsByProjectIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkflowService_ServiceDesc is the grpc.ServiceDesc for WorkflowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkflowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.WorkflowService",
	HandlerType: (*WorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkflowById",
			Handler:    _WorkflowService_GetWorkflowById_Handler,
		},
		{
			MethodName: "InsertWorkflow",
			Handler:    _WorkflowService_InsertWorkflow_Handler,
		},
		{
			MethodName: "GetWorkflowsByProjectId",
			Handler:    _WorkflowService_GetWorkflowsByProjectId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow-service.proto",
}