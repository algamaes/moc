// Code generated by protoc-gen-go. DO NOT EDIT.
// source: exec/moc_wssdhostagent_exec.proto

package admin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type ExecRequest struct {
	Execs                []*Exec          `protobuf:"bytes,1,rep,name=Execs,proto3" json:"Execs,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ExecRequest) Reset()         { *m = ExecRequest{} }
func (m *ExecRequest) String() string { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()    {}
func (*ExecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d4daa61e3d5a27, []int{0}
}

func (m *ExecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecRequest.Unmarshal(m, b)
}
func (m *ExecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecRequest.Marshal(b, m, deterministic)
}
func (m *ExecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecRequest.Merge(m, src)
}
func (m *ExecRequest) XXX_Size() int {
	return xxx_messageInfo_ExecRequest.Size(m)
}
func (m *ExecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecRequest proto.InternalMessageInfo

func (m *ExecRequest) GetExecs() []*Exec {
	if m != nil {
		return m.Execs
	}
	return nil
}

func (m *ExecRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type ExecResponse struct {
	Execs                []*Exec             `protobuf:"bytes,1,rep,name=Execs,proto3" json:"Execs,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ExecResponse) Reset()         { *m = ExecResponse{} }
func (m *ExecResponse) String() string { return proto.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()    {}
func (*ExecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d4daa61e3d5a27, []int{1}
}

func (m *ExecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecResponse.Unmarshal(m, b)
}
func (m *ExecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecResponse.Marshal(b, m, deterministic)
}
func (m *ExecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecResponse.Merge(m, src)
}
func (m *ExecResponse) XXX_Size() int {
	return xxx_messageInfo_ExecResponse.Size(m)
}
func (m *ExecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecResponse proto.InternalMessageInfo

func (m *ExecResponse) GetExecs() []*Exec {
	if m != nil {
		return m.Execs
	}
	return nil
}

func (m *ExecResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ExecResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Exec struct {
	Command              string         `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Output               string         `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Status               *common.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Exec) Reset()         { *m = Exec{} }
func (m *Exec) String() string { return proto.CompactTextString(m) }
func (*Exec) ProtoMessage()    {}
func (*Exec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7d4daa61e3d5a27, []int{2}
}

func (m *Exec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Exec.Unmarshal(m, b)
}
func (m *Exec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Exec.Marshal(b, m, deterministic)
}
func (m *Exec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exec.Merge(m, src)
}
func (m *Exec) XXX_Size() int {
	return xxx_messageInfo_Exec.Size(m)
}
func (m *Exec) XXX_DiscardUnknown() {
	xxx_messageInfo_Exec.DiscardUnknown(m)
}

var xxx_messageInfo_Exec proto.InternalMessageInfo

func (m *Exec) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Exec) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Exec) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecRequest)(nil), "moc.wssdhostagent.admin.ExecRequest")
	proto.RegisterType((*ExecResponse)(nil), "moc.wssdhostagent.admin.ExecResponse")
	proto.RegisterType((*Exec)(nil), "moc.wssdhostagent.admin.Exec")
}

func init() { proto.RegisterFile("exec/moc_wssdhostagent_exec.proto", fileDescriptor_d7d4daa61e3d5a27) }

var fileDescriptor_d7d4daa61e3d5a27 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xdf, 0x4a, 0xeb, 0x30,
	0x1c, 0x3e, 0x3d, 0x3b, 0xeb, 0x61, 0xe9, 0x39, 0xbb, 0x08, 0xa2, 0xa5, 0xa0, 0xcc, 0xaa, 0xb0,
	0xab, 0x54, 0x3a, 0x5f, 0xc0, 0xc1, 0x2e, 0xbc, 0x12, 0xaa, 0x28, 0x08, 0x32, 0xba, 0x34, 0xeb,
	0x8a, 0x6d, 0x7f, 0x31, 0x7f, 0xdc, 0x7c, 0x08, 0xdf, 0x59, 0x92, 0x74, 0xc2, 0x2e, 0x54, 0xf0,
	0xaa, 0x7c, 0xf9, 0xfe, 0x35, 0x5f, 0x8b, 0x8e, 0xd9, 0x86, 0xd1, 0xa4, 0x01, 0x3a, 0x5f, 0x4b,
	0x59, 0xac, 0x40, 0xaa, 0xbc, 0x64, 0xad, 0x9a, 0x9b, 0x63, 0xc2, 0x05, 0x28, 0xc0, 0x07, 0x0d,
	0x50, 0xb2, 0xc3, 0x92, 0xbc, 0x68, 0xaa, 0x36, 0x3a, 0x2a, 0x01, 0xca, 0x9a, 0x25, 0x56, 0xb6,
	0xd0, 0xcb, 0x64, 0x2d, 0x72, 0xce, 0x99, 0x90, 0xce, 0x18, 0x19, 0xe3, 0x9c, 0x42, 0xd3, 0x40,
	0xdb, 0x3d, 0x1c, 0x11, 0x6f, 0x50, 0x30, 0xdb, 0x30, 0x9a, 0xb1, 0x67, 0xcd, 0xa4, 0xc2, 0x13,
	0xd4, 0x37, 0x50, 0x86, 0xde, 0xa8, 0x37, 0x0e, 0xd2, 0x43, 0xf2, 0x49, 0x21, 0xb1, 0x26, 0xa7,
	0xc5, 0x17, 0xe8, 0xff, 0x35, 0x67, 0x22, 0x57, 0x15, 0xb4, 0xb7, 0xaf, 0x9c, 0x85, 0xbf, 0x47,
	0xde, 0x78, 0x98, 0x0e, 0xad, 0xf9, 0x83, 0xc9, 0x76, 0x45, 0xf1, 0x9b, 0x87, 0xfe, 0xb9, 0x6a,
	0xc9, 0xa1, 0x95, 0xec, 0x67, 0xdd, 0x29, 0xf2, 0x33, 0x26, 0x75, 0xad, 0x6c, 0x69, 0x90, 0x46,
	0xc4, 0x2d, 0x41, 0xb6, 0x4b, 0x90, 0x29, 0x40, 0x7d, 0x97, 0xd7, 0x9a, 0x65, 0x9d, 0x12, 0xef,
	0xa1, 0xfe, 0x4c, 0x08, 0x10, 0x61, 0x6f, 0xe4, 0x8d, 0x07, 0x99, 0x03, 0xf1, 0x23, 0xfa, 0x63,
	0x22, 0x71, 0x88, 0xfe, 0x9a, 0x85, 0xf2, 0xb6, 0x08, 0x3d, 0xcb, 0x6f, 0x21, 0xde, 0x47, 0x3e,
	0x68, 0xc5, 0xb5, 0xeb, 0x1a, 0x64, 0x1d, 0xc2, 0x27, 0xc8, 0x97, 0x2a, 0x57, 0x5a, 0xda, 0xc0,
	0x20, 0x0d, 0xec, 0x9b, 0xdf, 0xd8, 0xa3, 0xac, 0xa3, 0xd2, 0x02, 0x0d, 0x4c, 0xfc, 0xa5, 0xb9,
	0x07, 0xbe, 0x47, 0xfe, 0x55, 0xfb, 0x02, 0x4f, 0x0c, 0x9f, 0x7e, 0x7d, 0x4b, 0xf7, 0x59, 0xa2,
	0xb3, 0x6f, 0x54, 0x6e, 0xc1, 0xf8, 0xd7, 0x34, 0x7d, 0x38, 0x2f, 0x2b, 0xb5, 0xd2, 0x0b, 0x42,
	0xa1, 0x49, 0x9a, 0x8a, 0x0a, 0x90, 0xb0, 0x54, 0xe6, 0xaf, 0x4a, 0x04, 0xa7, 0xc9, 0x4e, 0x44,
	0x62, 0x23, 0x16, 0xbe, 0x9d, 0x6a, 0xf2, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xd1, 0x14, 0x9e,
	0x80, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExecAgentClient is the client API for ExecAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExecAgentClient interface {
	Invoke(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
}

type execAgentClient struct {
	cc *grpc.ClientConn
}

func NewExecAgentClient(cc *grpc.ClientConn) ExecAgentClient {
	return &execAgentClient{cc}
}

func (c *execAgentClient) Invoke(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := c.cc.Invoke(ctx, "/moc.wssdhostagent.admin.ExecAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecAgentServer is the server API for ExecAgent service.
type ExecAgentServer interface {
	Invoke(context.Context, *ExecRequest) (*ExecResponse, error)
}

// UnimplementedExecAgentServer can be embedded to have forward compatible implementations.
type UnimplementedExecAgentServer struct {
}

func (*UnimplementedExecAgentServer) Invoke(ctx context.Context, req *ExecRequest) (*ExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterExecAgentServer(s *grpc.Server, srv ExecAgentServer) {
	s.RegisterService(&_ExecAgent_serviceDesc, srv)
}

func _ExecAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.wssdhostagent.admin.ExecAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecAgentServer).Invoke(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.wssdhostagent.admin.ExecAgent",
	HandlerType: (*ExecAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _ExecAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exec/moc_wssdhostagent_exec.proto",
}
