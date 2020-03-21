// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/todo.proto

package todop

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Todo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	DeadLineTime         string   `protobuf:"bytes,4,opt,name=deadLineTime,proto3" json:"deadLineTime,omitempty"`
	CreatedTime          string   `protobuf:"bytes,5,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{0}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetDeadLineTime() string {
	if m != nil {
		return m.DeadLineTime
	}
	return ""
}

func (m *Todo) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

type CreateTodoReq struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoReq) Reset()         { *m = CreateTodoReq{} }
func (m *CreateTodoReq) String() string { return proto.CompactTextString(m) }
func (*CreateTodoReq) ProtoMessage()    {}
func (*CreateTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{1}
}

func (m *CreateTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoReq.Unmarshal(m, b)
}
func (m *CreateTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoReq.Marshal(b, m, deterministic)
}
func (m *CreateTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoReq.Merge(m, src)
}
func (m *CreateTodoReq) XXX_Size() int {
	return xxx_messageInfo_CreateTodoReq.Size(m)
}
func (m *CreateTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoReq proto.InternalMessageInfo

func (m *CreateTodoReq) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type CreateTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRes) Reset()         { *m = CreateTodoRes{} }
func (m *CreateTodoRes) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRes) ProtoMessage()    {}
func (*CreateTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{2}
}

func (m *CreateTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRes.Unmarshal(m, b)
}
func (m *CreateTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRes.Marshal(b, m, deterministic)
}
func (m *CreateTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRes.Merge(m, src)
}
func (m *CreateTodoRes) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRes.Size(m)
}
func (m *CreateTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRes proto.InternalMessageInfo

func (m *CreateTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type UpdateTodoReq struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoReq) Reset()         { *m = UpdateTodoReq{} }
func (m *UpdateTodoReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoReq) ProtoMessage()    {}
func (*UpdateTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{3}
}

func (m *UpdateTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoReq.Unmarshal(m, b)
}
func (m *UpdateTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoReq.Marshal(b, m, deterministic)
}
func (m *UpdateTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoReq.Merge(m, src)
}
func (m *UpdateTodoReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoReq.Size(m)
}
func (m *UpdateTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoReq proto.InternalMessageInfo

func (m *UpdateTodoReq) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type UpdateTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTodoRes) Reset()         { *m = UpdateTodoRes{} }
func (m *UpdateTodoRes) String() string { return proto.CompactTextString(m) }
func (*UpdateTodoRes) ProtoMessage()    {}
func (*UpdateTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{4}
}

func (m *UpdateTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTodoRes.Unmarshal(m, b)
}
func (m *UpdateTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTodoRes.Marshal(b, m, deterministic)
}
func (m *UpdateTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTodoRes.Merge(m, src)
}
func (m *UpdateTodoRes) XXX_Size() int {
	return xxx_messageInfo_UpdateTodoRes.Size(m)
}
func (m *UpdateTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTodoRes proto.InternalMessageInfo

func (m *UpdateTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type DeleteTodoReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoReq) Reset()         { *m = DeleteTodoReq{} }
func (m *DeleteTodoReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoReq) ProtoMessage()    {}
func (*DeleteTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{5}
}

func (m *DeleteTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoReq.Unmarshal(m, b)
}
func (m *DeleteTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoReq.Marshal(b, m, deterministic)
}
func (m *DeleteTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoReq.Merge(m, src)
}
func (m *DeleteTodoReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoReq.Size(m)
}
func (m *DeleteTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoReq proto.InternalMessageInfo

func (m *DeleteTodoReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteTodoRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRes) Reset()         { *m = DeleteTodoRes{} }
func (m *DeleteTodoRes) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRes) ProtoMessage()    {}
func (*DeleteTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{6}
}

func (m *DeleteTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRes.Unmarshal(m, b)
}
func (m *DeleteTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRes.Marshal(b, m, deterministic)
}
func (m *DeleteTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRes.Merge(m, src)
}
func (m *DeleteTodoRes) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRes.Size(m)
}
func (m *DeleteTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRes proto.InternalMessageInfo

func (m *DeleteTodoRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListTodoReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoReq) Reset()         { *m = ListTodoReq{} }
func (m *ListTodoReq) String() string { return proto.CompactTextString(m) }
func (*ListTodoReq) ProtoMessage()    {}
func (*ListTodoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{7}
}

func (m *ListTodoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoReq.Unmarshal(m, b)
}
func (m *ListTodoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoReq.Marshal(b, m, deterministic)
}
func (m *ListTodoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoReq.Merge(m, src)
}
func (m *ListTodoReq) XXX_Size() int {
	return xxx_messageInfo_ListTodoReq.Size(m)
}
func (m *ListTodoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoReq proto.InternalMessageInfo

type ListTodoRes struct {
	Todo                 *Todo    `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoRes) Reset()         { *m = ListTodoRes{} }
func (m *ListTodoRes) String() string { return proto.CompactTextString(m) }
func (*ListTodoRes) ProtoMessage()    {}
func (*ListTodoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b1ccc7160a78f, []int{8}
}

func (m *ListTodoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoRes.Unmarshal(m, b)
}
func (m *ListTodoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoRes.Marshal(b, m, deterministic)
}
func (m *ListTodoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoRes.Merge(m, src)
}
func (m *ListTodoRes) XXX_Size() int {
	return xxx_messageInfo_ListTodoRes.Size(m)
}
func (m *ListTodoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoRes proto.InternalMessageInfo

func (m *ListTodoRes) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.Todo")
	proto.RegisterType((*CreateTodoReq)(nil), "todo.CreateTodoReq")
	proto.RegisterType((*CreateTodoRes)(nil), "todo.CreateTodoRes")
	proto.RegisterType((*UpdateTodoReq)(nil), "todo.UpdateTodoReq")
	proto.RegisterType((*UpdateTodoRes)(nil), "todo.UpdateTodoRes")
	proto.RegisterType((*DeleteTodoReq)(nil), "todo.DeleteTodoReq")
	proto.RegisterType((*DeleteTodoRes)(nil), "todo.DeleteTodoRes")
	proto.RegisterType((*ListTodoReq)(nil), "todo.ListTodoReq")
	proto.RegisterType((*ListTodoRes)(nil), "todo.ListTodoRes")
}

func init() {
	proto.RegisterFile("proto/todo.proto", fileDescriptor_d28b1ccc7160a78f)
}

var fileDescriptor_d28b1ccc7160a78f = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x0d, 0xb5, 0x14, 0x3b, 0x88, 0xd1, 0xd1, 0x44, 0xc2, 0x41, 0x9b, 0x3d, 0xe9, 0xc1, 0xd6,
	0xb4, 0x89, 0x1f, 0xa0, 0x5e, 0x4c, 0x7a, 0xc2, 0x7a, 0xf1, 0x62, 0x2a, 0x3b, 0x87, 0x4d, 0xaa,
	0x8b, 0xec, 0xd6, 0x9f, 0xf0, 0x4f, 0xfd, 0x0a, 0xb3, 0x83, 0x04, 0xa8, 0x24, 0xd5, 0xdb, 0xbe,
	0x37, 0xef, 0x31, 0x6f, 0x5e, 0x80, 0x83, 0xbc, 0xd0, 0x56, 0x4f, 0xac, 0x96, 0x7a, 0xcc, 0x4f,
	0xec, 0xbb, 0xb7, 0xf8, 0xf4, 0xa0, 0xbf, 0xd0, 0x52, 0xe3, 0x3e, 0xf4, 0x94, 0x8c, 0xbd, 0x91,
	0x77, 0x3e, 0x4c, 0x7b, 0x4a, 0xe2, 0x09, 0x04, 0x6b, 0x43, 0xc5, 0xb3, 0x92, 0x71, 0x8f, 0xc9,
	0x81, 0x83, 0xf7, 0x12, 0x8f, 0xc1, 0xb7, 0xca, 0xae, 0x28, 0xde, 0x61, 0xba, 0x04, 0x28, 0x60,
	0x4f, 0xd2, 0x52, 0xce, 0xd5, 0x1b, 0x2d, 0xd4, 0x2b, 0xc5, 0x7d, 0x1e, 0xb6, 0x38, 0x1c, 0x41,
	0x98, 0x15, 0xb4, 0xb4, 0x24, 0x59, 0xe2, 0xb3, 0xa4, 0x49, 0x89, 0x09, 0x44, 0xb7, 0x0c, 0x5d,
	0xa4, 0x94, 0xde, 0xf1, 0x14, 0x38, 0x26, 0xe7, 0x0a, 0xa7, 0x30, 0xe6, 0xfc, 0x3c, 0x2c, 0xe3,
	0x6f, 0x18, 0xcc, 0x5f, 0x0c, 0x8f, 0xb9, 0xfc, 0xdf, 0x86, 0xa6, 0x61, 0xfb, 0x86, 0x33, 0x88,
	0xee, 0x68, 0x45, 0xf5, 0x86, 0x8d, 0x66, 0xc5, 0x45, 0x5b, 0x60, 0x30, 0x86, 0xc0, 0xac, 0xb3,
	0x8c, 0x8c, 0x61, 0xd5, 0x6e, 0x5a, 0x41, 0x11, 0x41, 0x38, 0x57, 0xc6, 0xfe, 0x7c, 0x49, 0x5c,
	0x36, 0xe1, 0xd6, 0x24, 0xd3, 0x2f, 0x0f, 0x42, 0x07, 0x1f, 0xa8, 0xf8, 0x50, 0x19, 0xe1, 0x35,
	0x40, 0x5d, 0x16, 0x1e, 0x95, 0xfa, 0x56, 0xdf, 0x49, 0x07, 0x69, 0x9c, 0xaf, 0xae, 0xa0, 0xf2,
	0xb5, 0x5a, 0x4c, 0x3a, 0x48, 0xf6, 0xd5, 0x87, 0x56, 0xbe, 0x56, 0x37, 0x49, 0x07, 0x69, 0x70,
	0x06, 0xc3, 0xea, 0x4c, 0x83, 0x87, 0xa5, 0xa2, 0x51, 0x43, 0xf2, 0x8b, 0x32, 0x57, 0xde, 0x4d,
	0xf0, 0xe4, 0x3b, 0x36, 0x7f, 0x19, 0xf0, 0xef, 0x3d, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd8,
	0x1d, 0xde, 0x11, 0xf2, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoRes, error)
	UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoRes, error)
	DeleteTodo(ctx context.Context, in *DeleteTodoReq, opts ...grpc.CallOption) (*DeleteTodoRes, error)
	ListTodos(ctx context.Context, in *ListTodoReq, opts ...grpc.CallOption) (TodoService_ListTodosClient, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoReq, opts ...grpc.CallOption) (*CreateTodoRes, error) {
	out := new(CreateTodoRes)
	err := c.cc.Invoke(ctx, "/todo.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoReq, opts ...grpc.CallOption) (*UpdateTodoRes, error) {
	out := new(UpdateTodoRes)
	err := c.cc.Invoke(ctx, "/todo.TodoService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoReq, opts ...grpc.CallOption) (*DeleteTodoRes, error) {
	out := new(DeleteTodoRes)
	err := c.cc.Invoke(ctx, "/todo.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTodos(ctx context.Context, in *ListTodoReq, opts ...grpc.CallOption) (TodoService_ListTodosClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TodoService_serviceDesc.Streams[0], "/todo.TodoService/ListTodos", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceListTodosClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoService_ListTodosClient interface {
	Recv() (*ListTodoRes, error)
	grpc.ClientStream
}

type todoServiceListTodosClient struct {
	grpc.ClientStream
}

func (x *todoServiceListTodosClient) Recv() (*ListTodoRes, error) {
	m := new(ListTodoRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoReq) (*CreateTodoRes, error)
	UpdateTodo(context.Context, *UpdateTodoReq) (*UpdateTodoRes, error)
	DeleteTodo(context.Context, *DeleteTodoReq) (*DeleteTodoRes, error)
	ListTodos(*ListTodoReq, TodoService_ListTodosServer) error
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) CreateTodo(ctx context.Context, req *CreateTodoReq) (*CreateTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) UpdateTodo(ctx context.Context, req *UpdateTodoReq) (*UpdateTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (*UnimplementedTodoServiceServer) DeleteTodo(ctx context.Context, req *DeleteTodoReq) (*DeleteTodoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (*UnimplementedTodoServiceServer) ListTodos(req *ListTodoReq, srv TodoService_ListTodosServer) error {
	return status.Errorf(codes.Unimplemented, "method ListTodos not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodo(ctx, req.(*UpdateTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTodos_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListTodoReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServiceServer).ListTodos(m, &todoServiceListTodosServer{stream})
}

type TodoService_ListTodosServer interface {
	Send(*ListTodoRes) error
	grpc.ServerStream
}

type todoServiceListTodosServer struct {
	grpc.ServerStream
}

func (x *todoServiceListTodosServer) Send(m *ListTodoRes) error {
	return x.ServerStream.SendMsg(m)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoService_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTodos",
			Handler:       _TodoService_ListTodos_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/todo.proto",
}