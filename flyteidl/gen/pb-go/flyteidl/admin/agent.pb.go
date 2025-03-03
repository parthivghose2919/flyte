// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/agent.proto

package admin

import (
	fmt "fmt"
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The state of the execution is used to control its visibility in the UI/CLI.
type State int32

const (
	State_RETRYABLE_FAILURE State = 0
	State_PERMANENT_FAILURE State = 1
	State_PENDING           State = 2
	State_RUNNING           State = 3
	State_SUCCEEDED         State = 4
)

var State_name = map[int32]string{
	0: "RETRYABLE_FAILURE",
	1: "PERMANENT_FAILURE",
	2: "PENDING",
	3: "RUNNING",
	4: "SUCCEEDED",
}

var State_value = map[string]int32{
	"RETRYABLE_FAILURE": 0,
	"PERMANENT_FAILURE": 1,
	"PENDING":           2,
	"RUNNING":           3,
	"SUCCEEDED":         4,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{0}
}

// Represents a subset of runtime task execution metadata that are relevant to external plugins.
type TaskExecutionMetadata struct {
	// ID of the task execution
	TaskExecutionId *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=task_execution_id,json=taskExecutionId,proto3" json:"task_execution_id,omitempty"`
	// k8s namespace where the task is executed in
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Labels attached to the task execution
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations attached to the task execution
	Annotations map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s service account associated with the task execution
	K8SServiceAccount string `protobuf:"bytes,5,opt,name=k8s_service_account,json=k8sServiceAccount,proto3" json:"k8s_service_account,omitempty"`
	// Environment variables attached to the task execution
	EnvironmentVariables map[string]string `protobuf:"bytes,6,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskExecutionMetadata) Reset()         { *m = TaskExecutionMetadata{} }
func (m *TaskExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionMetadata) ProtoMessage()    {}
func (*TaskExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{0}
}

func (m *TaskExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionMetadata.Unmarshal(m, b)
}
func (m *TaskExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *TaskExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionMetadata.Merge(m, src)
}
func (m *TaskExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionMetadata.Size(m)
}
func (m *TaskExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionMetadata proto.InternalMessageInfo

func (m *TaskExecutionMetadata) GetTaskExecutionId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.TaskExecutionId
	}
	return nil
}

func (m *TaskExecutionMetadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TaskExecutionMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TaskExecutionMetadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *TaskExecutionMetadata) GetK8SServiceAccount() string {
	if m != nil {
		return m.K8SServiceAccount
	}
	return ""
}

func (m *TaskExecutionMetadata) GetEnvironmentVariables() map[string]string {
	if m != nil {
		return m.EnvironmentVariables
	}
	return nil
}

// Represents a request structure to create task.
type CreateTaskRequest struct {
	// The inputs required to start the execution. All required inputs must be
	// included in this map. If not required and not provided, defaults apply.
	// +optional
	Inputs *core.LiteralMap `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Template of the task that encapsulates all the metadata of the task.
	Template *core.TaskTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Prefix for where task output data will be written. (e.g. s3://my-bucket/randomstring)
	OutputPrefix string `protobuf:"bytes,3,opt,name=output_prefix,json=outputPrefix,proto3" json:"output_prefix,omitempty"`
	// subset of runtime task execution metadata.
	TaskExecutionMetadata *TaskExecutionMetadata `protobuf:"bytes,4,opt,name=task_execution_metadata,json=taskExecutionMetadata,proto3" json:"task_execution_metadata,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{1}
}

func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(m, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetInputs() *core.LiteralMap {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CreateTaskRequest) GetTemplate() *core.TaskTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *CreateTaskRequest) GetOutputPrefix() string {
	if m != nil {
		return m.OutputPrefix
	}
	return ""
}

func (m *CreateTaskRequest) GetTaskExecutionMetadata() *TaskExecutionMetadata {
	if m != nil {
		return m.TaskExecutionMetadata
	}
	return nil
}

// Represents a create response structure.
type CreateTaskResponse struct {
	// Metadata is created by the agent. It could be a string (jobId) or a dict (more complex metadata).
	ResourceMeta         []byte   `protobuf:"bytes,1,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{2}
}

func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(m, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// A message used to fetch a job resource from flyte agent server.
type GetTaskRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata about the resource to be pass to the agent.
	ResourceMeta         []byte   `protobuf:"bytes,2,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskRequest) Reset()         { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()    {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{3}
}

func (m *GetTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskRequest.Unmarshal(m, b)
}
func (m *GetTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskRequest.Marshal(b, m, deterministic)
}
func (m *GetTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskRequest.Merge(m, src)
}
func (m *GetTaskRequest) XXX_Size() int {
	return xxx_messageInfo_GetTaskRequest.Size(m)
}
func (m *GetTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskRequest proto.InternalMessageInfo

func (m *GetTaskRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *GetTaskRequest) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// Response to get an individual task resource.
type GetTaskResponse struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetTaskResponse) Reset()         { *m = GetTaskResponse{} }
func (m *GetTaskResponse) String() string { return proto.CompactTextString(m) }
func (*GetTaskResponse) ProtoMessage()    {}
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{4}
}

func (m *GetTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskResponse.Unmarshal(m, b)
}
func (m *GetTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskResponse.Marshal(b, m, deterministic)
}
func (m *GetTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskResponse.Merge(m, src)
}
func (m *GetTaskResponse) XXX_Size() int {
	return xxx_messageInfo_GetTaskResponse.Size(m)
}
func (m *GetTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskResponse proto.InternalMessageInfo

func (m *GetTaskResponse) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type Resource struct {
	// The state of the execution is used to control its visibility in the UI/CLI.
	State State `protobuf:"varint,1,opt,name=state,proto3,enum=flyteidl.admin.State" json:"state,omitempty"`
	// The outputs of the execution. It's typically used by sql task. Agent service will create a
	// Structured dataset pointing to the query result table.
	// +optional
	Outputs *core.LiteralMap `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"`
	// A descriptive message for the current state. e.g. waiting for cluster.
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{5}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetState() State {
	if m != nil {
		return m.State
	}
	return State_RETRYABLE_FAILURE
}

func (m *Resource) GetOutputs() *core.LiteralMap {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Resource) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A message used to delete a task.
type DeleteTaskRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata about the resource to be pass to the agent.
	ResourceMeta         []byte   `protobuf:"bytes,2,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskRequest) Reset()         { *m = DeleteTaskRequest{} }
func (m *DeleteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskRequest) ProtoMessage()    {}
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{6}
}

func (m *DeleteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskRequest.Unmarshal(m, b)
}
func (m *DeleteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskRequest.Merge(m, src)
}
func (m *DeleteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskRequest.Size(m)
}
func (m *DeleteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskRequest proto.InternalMessageInfo

func (m *DeleteTaskRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *DeleteTaskRequest) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// Response to delete a task.
type DeleteTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskResponse) Reset()         { *m = DeleteTaskResponse{} }
func (m *DeleteTaskResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskResponse) ProtoMessage()    {}
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{7}
}

func (m *DeleteTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskResponse.Unmarshal(m, b)
}
func (m *DeleteTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskResponse.Merge(m, src)
}
func (m *DeleteTaskResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskResponse.Size(m)
}
func (m *DeleteTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("flyteidl.admin.State", State_name, State_value)
	proto.RegisterType((*TaskExecutionMetadata)(nil), "flyteidl.admin.TaskExecutionMetadata")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.EnvironmentVariablesEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.LabelsEntry")
	proto.RegisterType((*CreateTaskRequest)(nil), "flyteidl.admin.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "flyteidl.admin.CreateTaskResponse")
	proto.RegisterType((*GetTaskRequest)(nil), "flyteidl.admin.GetTaskRequest")
	proto.RegisterType((*GetTaskResponse)(nil), "flyteidl.admin.GetTaskResponse")
	proto.RegisterType((*Resource)(nil), "flyteidl.admin.Resource")
	proto.RegisterType((*DeleteTaskRequest)(nil), "flyteidl.admin.DeleteTaskRequest")
	proto.RegisterType((*DeleteTaskResponse)(nil), "flyteidl.admin.DeleteTaskResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/agent.proto", fileDescriptor_c434e52bb0028071) }

var fileDescriptor_c434e52bb0028071 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x7f, 0x6b, 0x2a, 0x47,
	0x14, 0xad, 0x1a, 0x8d, 0x5e, 0xf3, 0x43, 0xa7, 0x91, 0x6e, 0x4c, 0x5a, 0x82, 0xa5, 0x25, 0xb4,
	0x74, 0x25, 0xa6, 0xb4, 0x49, 0x4b, 0x5b, 0x4c, 0xdc, 0x8a, 0x60, 0x24, 0x4c, 0xb4, 0xb4, 0x85,
	0x56, 0xc6, 0xf5, 0xea, 0x5b, 0x5c, 0x67, 0xf7, 0xed, 0xcc, 0x4a, 0xfc, 0xfb, 0xc1, 0xfb, 0x0c,
	0xef, 0xe3, 0x3e, 0x76, 0xf6, 0x47, 0x54, 0x7c, 0x8f, 0x84, 0xf7, 0xdf, 0xce, 0x3d, 0xf7, 0x9e,
	0x39, 0x73, 0xce, 0x2c, 0x03, 0xd5, 0x89, 0xbd, 0x94, 0x68, 0x8d, 0xed, 0x3a, 0x1b, 0xcf, 0x2d,
	0x5e, 0x67, 0x53, 0xe4, 0x52, 0x77, 0x3d, 0x47, 0x3a, 0xe4, 0x20, 0xc6, 0x74, 0x85, 0x55, 0x4f,
	0x93, 0x5e, 0xd3, 0xf1, 0xb0, 0x6e, 0x5b, 0x12, 0x3d, 0x66, 0x8b, 0xb0, 0xbb, 0x7a, 0xbc, 0x8e,
	0x4a, 0x26, 0x66, 0x31, 0xf4, 0xe5, 0x3a, 0x64, 0x71, 0x89, 0xde, 0x84, 0x99, 0x18, 0xc1, 0x5f,
	0x6d, 0xc0, 0x63, 0xe4, 0xd2, 0x9a, 0x58, 0xe8, 0x85, 0x78, 0xed, 0x5d, 0x16, 0x2a, 0x7d, 0x26,
	0x66, 0xc6, 0x23, 0x9a, 0xbe, 0xb4, 0x1c, 0x7e, 0x87, 0x92, 0x8d, 0x99, 0x64, 0x84, 0x42, 0x39,
	0xd8, 0x67, 0x88, 0x31, 0x32, 0xb4, 0xc6, 0x5a, 0xea, 0x2c, 0x75, 0x5e, 0x6c, 0x7c, 0xab, 0x27,
	0xea, 0x03, 0x56, 0x7d, 0x8d, 0xa0, 0x93, 0x6c, 0x41, 0x0f, 0xe5, 0x3a, 0x40, 0x4e, 0xa1, 0xc0,
	0xd9, 0x1c, 0x85, 0xcb, 0x4c, 0xd4, 0xd2, 0x67, 0xa9, 0xf3, 0x02, 0x7d, 0x2a, 0x90, 0x0e, 0xe4,
	0x6c, 0x36, 0x42, 0x5b, 0x68, 0x99, 0xb3, 0xcc, 0x79, 0xb1, 0x71, 0xa1, 0xaf, 0x9b, 0xa4, 0x6f,
	0x15, 0xaa, 0x77, 0xd5, 0x8c, 0xc1, 0xa5, 0xb7, 0xa4, 0x11, 0x01, 0xf9, 0x1b, 0x8a, 0x8c, 0x73,
	0x47, 0xb2, 0xa0, 0x53, 0x68, 0x3b, 0x8a, 0xef, 0xa7, 0xe7, 0xf1, 0x35, 0x9f, 0x06, 0x43, 0xd2,
	0x55, 0x2a, 0xa2, 0xc3, 0xe7, 0xb3, 0x2b, 0x31, 0x14, 0xe8, 0x2d, 0x2c, 0x13, 0x87, 0xcc, 0x34,
	0x1d, 0x9f, 0x4b, 0x2d, 0xab, 0x0e, 0x53, 0x9e, 0x5d, 0x89, 0x87, 0x10, 0x69, 0x86, 0x00, 0x91,
	0x50, 0x41, 0xbe, 0xb0, 0x3c, 0x87, 0xcf, 0x91, 0xcb, 0xe1, 0x82, 0x79, 0x16, 0x1b, 0xd9, 0x28,
	0xb4, 0x9c, 0xd2, 0xf4, 0xc7, 0xf3, 0x34, 0x19, 0x4f, 0x14, 0x7f, 0xc5, 0x0c, 0xa1, 0xb8, 0x23,
	0xdc, 0x02, 0x55, 0xaf, 0xa1, 0xb8, 0x62, 0x0b, 0x29, 0x41, 0x66, 0x86, 0x4b, 0x95, 0x5e, 0x81,
	0x06, 0x9f, 0xe4, 0x08, 0xb2, 0x0b, 0x66, 0xfb, 0x71, 0x0a, 0xe1, 0xe2, 0x97, 0xf4, 0x55, 0xaa,
	0xfa, 0x3b, 0x94, 0x36, 0x1d, 0x78, 0xd1, 0x7c, 0x1b, 0x8e, 0x3f, 0xa8, 0xf6, 0x25, 0x44, 0xb5,
	0x37, 0x69, 0x28, 0xdf, 0x7a, 0xc8, 0x24, 0x06, 0x9e, 0x50, 0x7c, 0xed, 0xa3, 0x90, 0xe4, 0x02,
	0x72, 0x16, 0x77, 0x7d, 0x29, 0xa2, 0xbb, 0x78, 0xbc, 0x71, 0x17, 0xbb, 0xe1, 0x9f, 0x73, 0xc7,
	0x5c, 0x1a, 0x35, 0x92, 0x9f, 0x21, 0x2f, 0x71, 0xee, 0xda, 0x4c, 0x86, 0xbb, 0x14, 0x1b, 0x27,
	0x5b, 0x2e, 0x70, 0x3f, 0x6a, 0xa1, 0x49, 0x33, 0xf9, 0x1a, 0xf6, 0x1d, 0x5f, 0xba, 0xbe, 0x1c,
	0xba, 0x1e, 0x4e, 0xac, 0x47, 0x2d, 0xa3, 0x34, 0xee, 0x85, 0xc5, 0x7b, 0x55, 0x23, 0xff, 0xc1,
	0x17, 0x1b, 0xff, 0xc9, 0x3c, 0x4a, 0x4d, 0xdb, 0x51, 0x9b, 0x7d, 0xf3, 0xac, 0x88, 0x69, 0x45,
	0x6e, 0x2b, 0xd7, 0xae, 0x81, 0xac, 0x9a, 0x20, 0x5c, 0x87, 0x0b, 0xa5, 0xcc, 0x43, 0xe1, 0xf8,
	0x9e, 0x89, 0x6a, 0x3b, 0x65, 0xc6, 0x1e, 0xdd, 0x8b, 0x8b, 0xc1, 0x78, 0x8d, 0xc2, 0x41, 0x1b,
	0xe5, 0xaa, 0x79, 0x27, 0x50, 0x50, 0x5a, 0xe5, 0xd2, 0xc5, 0x28, 0x84, 0x7c, 0x50, 0xe8, 0x2f,
	0xdd, 0x2d, 0x9c, 0xe9, 0x2d, 0x9c, 0x6d, 0x38, 0x4c, 0x38, 0x23, 0x2d, 0x3f, 0x42, 0x3e, 0x6e,
	0x89, 0x32, 0xd1, 0x36, 0x4f, 0x4c, 0x23, 0x9c, 0x26, 0x9d, 0xb5, 0xb7, 0x29, 0xc8, 0xc7, 0x65,
	0xf2, 0x3d, 0x64, 0x85, 0x0c, 0xe2, 0x09, 0xe6, 0x0f, 0x1a, 0x95, 0xcd, 0xf9, 0x87, 0x00, 0xa4,
	0x61, 0x0f, 0xb9, 0x84, 0xdd, 0x30, 0x00, 0x11, 0xa5, 0xf9, 0x91, 0x2b, 0x10, 0x77, 0x12, 0x0d,
	0x76, 0xe7, 0x28, 0x04, 0x9b, 0x62, 0x14, 0x62, 0xbc, 0xac, 0x0d, 0xa0, 0xdc, 0x42, 0x1b, 0xd7,
	0x6f, 0xd9, 0xa7, 0x1b, 0x75, 0x04, 0x64, 0x95, 0x36, 0xf4, 0xea, 0xbb, 0xff, 0x21, 0xab, 0xce,
	0x42, 0x2a, 0x50, 0xa6, 0x46, 0x9f, 0xfe, 0xd3, 0xbc, 0xe9, 0x1a, 0xc3, 0x3f, 0x9b, 0x9d, 0xee,
	0x80, 0x1a, 0xa5, 0xcf, 0x82, 0xf2, 0xbd, 0x41, 0xef, 0x9a, 0x3d, 0xa3, 0xd7, 0x4f, 0xca, 0x29,
	0x52, 0x84, 0xdd, 0x7b, 0xa3, 0xd7, 0xea, 0xf4, 0xda, 0xa5, 0x74, 0xb0, 0xa0, 0x83, 0x5e, 0x2f,
	0x58, 0x64, 0xc8, 0x3e, 0x14, 0x1e, 0x06, 0xb7, 0xb7, 0x86, 0xd1, 0x32, 0x5a, 0xa5, 0x9d, 0x9b,
	0xdf, 0xfe, 0xfd, 0x75, 0x6a, 0xc9, 0x57, 0xfe, 0x48, 0x37, 0x9d, 0x79, 0x5d, 0xd9, 0xe2, 0x78,
	0xd3, 0xf0, 0xa3, 0x9e, 0x3c, 0x05, 0x53, 0xe4, 0x75, 0x77, 0xf4, 0xc3, 0xd4, 0xa9, 0xaf, 0xbf,
	0x50, 0xa3, 0x9c, 0x7a, 0x14, 0x2e, 0xdf, 0x07, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x0a, 0x0e, 0xc8,
	0xba, 0x06, 0x00, 0x00,
}
