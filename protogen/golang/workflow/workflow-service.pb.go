// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: workflow-service.proto

package workflow

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskState int32

const (
	TaskState_STATE_UNSPECIFIED  TaskState = 0 // Default value if none specified.
	TaskState_ON_HOLD            TaskState = 1 // Corresponds to "on hold" in Go.
	TaskState_UNDER_CONSTRUCTION TaskState = 2 // Corresponds to "under construction" in Go.
	TaskState_FINISHED           TaskState = 3 // Corresponds to "finished" in Go.
)

// Enum value maps for TaskState.
var (
	TaskState_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ON_HOLD",
		2: "UNDER_CONSTRUCTION",
		3: "FINISHED",
	}
	TaskState_value = map[string]int32{
		"STATE_UNSPECIFIED":  0,
		"ON_HOLD":            1,
		"UNDER_CONSTRUCTION": 2,
		"FINISHED":           3,
	}
)

func (x TaskState) Enum() *TaskState {
	p := new(TaskState)
	*p = x
	return p
}

func (x TaskState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskState) Descriptor() protoreflect.EnumDescriptor {
	return file_workflow_service_proto_enumTypes[0].Descriptor()
}

func (TaskState) Type() protoreflect.EnumType {
	return &file_workflow_service_proto_enumTypes[0]
}

func (x TaskState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskState.Descriptor instead.
func (TaskState) EnumDescriptor() ([]byte, []int) {
	return file_workflow_service_proto_rawDescGZIP(), []int{0}
}

type GetWorkflowByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId string `protobuf:"bytes,1,opt,name=workflowId,proto3" json:"workflowId,omitempty"`
}

func (x *GetWorkflowByIdRequest) Reset() {
	*x = GetWorkflowByIdRequest{}
	mi := &file_workflow_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkflowByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowByIdRequest) ProtoMessage() {}

func (x *GetWorkflowByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowByIdRequest.ProtoReflect.Descriptor instead.
func (*GetWorkflowByIdRequest) Descriptor() ([]byte, []int) {
	return file_workflow_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetWorkflowByIdRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

type GetWorkflowByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Tasks       []*Task `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
	ProjectId   string  `protobuf:"bytes,5,opt,name=projectId,proto3" json:"projectId,omitempty"`
}

func (x *GetWorkflowByIdResponse) Reset() {
	*x = GetWorkflowByIdResponse{}
	mi := &file_workflow_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkflowByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowByIdResponse) ProtoMessage() {}

func (x *GetWorkflowByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowByIdResponse.ProtoReflect.Descriptor instead.
func (*GetWorkflowByIdResponse) Descriptor() ([]byte, []int) {
	return file_workflow_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetWorkflowByIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetWorkflowByIdResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetWorkflowByIdResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetWorkflowByIdResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *GetWorkflowByIdResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetWorkflowsByProjectIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=projectId,proto3" json:"projectId,omitempty"`
}

func (x *GetWorkflowsByProjectIdRequest) Reset() {
	*x = GetWorkflowsByProjectIdRequest{}
	mi := &file_workflow_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkflowsByProjectIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowsByProjectIdRequest) ProtoMessage() {}

func (x *GetWorkflowsByProjectIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowsByProjectIdRequest.ProtoReflect.Descriptor instead.
func (*GetWorkflowsByProjectIdRequest) Descriptor() ([]byte, []int) {
	return file_workflow_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetWorkflowsByProjectIdRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetWorkflowsByProjectIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Workflows []*GetWorkflowByIdResponse `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows,omitempty"`
}

func (x *GetWorkflowsByProjectIdResponse) Reset() {
	*x = GetWorkflowsByProjectIdResponse{}
	mi := &file_workflow_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkflowsByProjectIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowsByProjectIdResponse) ProtoMessage() {}

func (x *GetWorkflowsByProjectIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowsByProjectIdResponse.ProtoReflect.Descriptor instead.
func (*GetWorkflowsByProjectIdResponse) Descriptor() ([]byte, []int) {
	return file_workflow_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetWorkflowsByProjectIdResponse) GetWorkflows() []*GetWorkflowByIdResponse {
	if x != nil {
		return x.Workflows
	}
	return nil
}

type InsertWorkflowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tasks       []*Task `protobuf:"bytes,3,rep,name=tasks,proto3" json:"tasks,omitempty"`
	ProjectId   string  `protobuf:"bytes,4,opt,name=projectId,proto3" json:"projectId,omitempty"`
}

func (x *InsertWorkflowRequest) Reset() {
	*x = InsertWorkflowRequest{}
	mi := &file_workflow_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertWorkflowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertWorkflowRequest) ProtoMessage() {}

func (x *InsertWorkflowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertWorkflowRequest.ProtoReflect.Descriptor instead.
func (*InsertWorkflowRequest) Descriptor() ([]byte, []int) {
	return file_workflow_service_proto_rawDescGZIP(), []int{4}
}

func (x *InsertWorkflowRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InsertWorkflowRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InsertWorkflowRequest) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *InsertWorkflowRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type InsertWorkflowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *GetWorkflowByIdResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *InsertWorkflowResponse) Reset() {
	*x = InsertWorkflowResponse{}
	mi := &file_workflow_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertWorkflowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertWorkflowResponse) ProtoMessage() {}

func (x *InsertWorkflowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertWorkflowResponse.ProtoReflect.Descriptor instead.
func (*InsertWorkflowResponse) Descriptor() ([]byte, []int) {
	return file_workflow_service_proto_rawDescGZIP(), []int{5}
}

func (x *InsertWorkflowResponse) GetResponse() *GetWorkflowByIdResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State      TaskState `protobuf:"varint,2,opt,name=state,proto3,enum=workflow.TaskState" json:"state,omitempty"`
	Name       string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	DependsOn  []*Task   `protobuf:"bytes,4,rep,name=dependsOn,proto3" json:"dependsOn,omitempty"`
	ProjectId  string    `protobuf:"bytes,5,opt,name=projectId,proto3" json:"projectId,omitempty"`
	WorkflowId string    `protobuf:"bytes,6,opt,name=workflowId,proto3" json:"workflowId,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_workflow_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_workflow_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_workflow_service_proto_rawDescGZIP(), []int{6}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetState() TaskState {
	if x != nil {
		return x.State
	}
	return TaskState_STATE_UNSPECIFIED
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetDependsOn() []*Task {
	if x != nil {
		return x.DependsOn
	}
	return nil
}

func (x *Task) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Task) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

var File_workflow_service_proto protoreflect.FileDescriptor

var file_workflow_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x38, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x3e, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x62, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x73, 0x22, 0x91, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x16, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xc1, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x64, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x73, 0x4f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x64, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x73, 0x4f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x2a, 0x55, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x4e, 0x5f,
	0x48, 0x4f, 0x4c, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f,
	0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03, 0x32, 0x96, 0x03, 0x0a,
	0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x7a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f,
	0x7b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x0e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x1f,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x9a, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x28, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x79, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x24, 0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x7d, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_workflow_service_proto_rawDescOnce sync.Once
	file_workflow_service_proto_rawDescData = file_workflow_service_proto_rawDesc
)

func file_workflow_service_proto_rawDescGZIP() []byte {
	file_workflow_service_proto_rawDescOnce.Do(func() {
		file_workflow_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_workflow_service_proto_rawDescData)
	})
	return file_workflow_service_proto_rawDescData
}

var file_workflow_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_workflow_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_workflow_service_proto_goTypes = []any{
	(TaskState)(0),                          // 0: workflow.TaskState
	(*GetWorkflowByIdRequest)(nil),          // 1: workflow.GetWorkflowByIdRequest
	(*GetWorkflowByIdResponse)(nil),         // 2: workflow.GetWorkflowByIdResponse
	(*GetWorkflowsByProjectIdRequest)(nil),  // 3: workflow.GetWorkflowsByProjectIdRequest
	(*GetWorkflowsByProjectIdResponse)(nil), // 4: workflow.GetWorkflowsByProjectIdResponse
	(*InsertWorkflowRequest)(nil),           // 5: workflow.InsertWorkflowRequest
	(*InsertWorkflowResponse)(nil),          // 6: workflow.InsertWorkflowResponse
	(*Task)(nil),                            // 7: workflow.Task
}
var file_workflow_service_proto_depIdxs = []int32{
	7, // 0: workflow.GetWorkflowByIdResponse.tasks:type_name -> workflow.Task
	2, // 1: workflow.GetWorkflowsByProjectIdResponse.workflows:type_name -> workflow.GetWorkflowByIdResponse
	7, // 2: workflow.InsertWorkflowRequest.tasks:type_name -> workflow.Task
	2, // 3: workflow.InsertWorkflowResponse.response:type_name -> workflow.GetWorkflowByIdResponse
	0, // 4: workflow.Task.state:type_name -> workflow.TaskState
	7, // 5: workflow.Task.dependsOn:type_name -> workflow.Task
	1, // 6: workflow.WorkflowService.GetWorkflowById:input_type -> workflow.GetWorkflowByIdRequest
	5, // 7: workflow.WorkflowService.InsertWorkflow:input_type -> workflow.InsertWorkflowRequest
	3, // 8: workflow.WorkflowService.GetWorkflowsByProjectId:input_type -> workflow.GetWorkflowsByProjectIdRequest
	2, // 9: workflow.WorkflowService.GetWorkflowById:output_type -> workflow.GetWorkflowByIdResponse
	6, // 10: workflow.WorkflowService.InsertWorkflow:output_type -> workflow.InsertWorkflowResponse
	4, // 11: workflow.WorkflowService.GetWorkflowsByProjectId:output_type -> workflow.GetWorkflowsByProjectIdResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_workflow_service_proto_init() }
func file_workflow_service_proto_init() {
	if File_workflow_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_workflow_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workflow_service_proto_goTypes,
		DependencyIndexes: file_workflow_service_proto_depIdxs,
		EnumInfos:         file_workflow_service_proto_enumTypes,
		MessageInfos:      file_workflow_service_proto_msgTypes,
	}.Build()
	File_workflow_service_proto = out.File
	file_workflow_service_proto_rawDesc = nil
	file_workflow_service_proto_goTypes = nil
	file_workflow_service_proto_depIdxs = nil
}
