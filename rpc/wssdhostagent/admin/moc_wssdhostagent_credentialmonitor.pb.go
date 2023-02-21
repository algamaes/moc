// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_wssdhostagent_credentialmonitor.proto

package admin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/microsoft/moc/rpc/common"
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

type CertificateStatus int32

const (
	CertificateStatus_Single  CertificateStatus = 0
	CertificateStatus_Overlap CertificateStatus = 1
)

var CertificateStatus_name = map[int32]string{
	0: "Single",
	1: "Overlap",
}

var CertificateStatus_value = map[string]int32{
	"Single":  0,
	"Overlap": 1,
}

func (x CertificateStatus) String() string {
	return proto.EnumName(CertificateStatus_name, int32(x))
}

func (CertificateStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13347e86c60e2a65, []int{0}
}

type CredentialMonitorRequest struct {
	CredentialMonitor    *CredentialMonitor `protobuf:"bytes,1,opt,name=CredentialMonitor,proto3" json:"CredentialMonitor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CredentialMonitorRequest) Reset()         { *m = CredentialMonitorRequest{} }
func (m *CredentialMonitorRequest) String() string { return proto.CompactTextString(m) }
func (*CredentialMonitorRequest) ProtoMessage()    {}
func (*CredentialMonitorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_13347e86c60e2a65, []int{0}
}

func (m *CredentialMonitorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialMonitorRequest.Unmarshal(m, b)
}
func (m *CredentialMonitorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialMonitorRequest.Marshal(b, m, deterministic)
}
func (m *CredentialMonitorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialMonitorRequest.Merge(m, src)
}
func (m *CredentialMonitorRequest) XXX_Size() int {
	return xxx_messageInfo_CredentialMonitorRequest.Size(m)
}
func (m *CredentialMonitorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialMonitorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialMonitorRequest proto.InternalMessageInfo

func (m *CredentialMonitorRequest) GetCredentialMonitor() *CredentialMonitor {
	if m != nil {
		return m.CredentialMonitor
	}
	return nil
}

type CredentialMonitorResponse struct {
	CredentialMonitor    *CredentialMonitor `protobuf:"bytes,1,opt,name=CredentialMonitor,proto3" json:"CredentialMonitor,omitempty"`
	Error                string             `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CredentialMonitorResponse) Reset()         { *m = CredentialMonitorResponse{} }
func (m *CredentialMonitorResponse) String() string { return proto.CompactTextString(m) }
func (*CredentialMonitorResponse) ProtoMessage()    {}
func (*CredentialMonitorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13347e86c60e2a65, []int{1}
}

func (m *CredentialMonitorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialMonitorResponse.Unmarshal(m, b)
}
func (m *CredentialMonitorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialMonitorResponse.Marshal(b, m, deterministic)
}
func (m *CredentialMonitorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialMonitorResponse.Merge(m, src)
}
func (m *CredentialMonitorResponse) XXX_Size() int {
	return xxx_messageInfo_CredentialMonitorResponse.Size(m)
}
func (m *CredentialMonitorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialMonitorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialMonitorResponse proto.InternalMessageInfo

func (m *CredentialMonitorResponse) GetCredentialMonitor() *CredentialMonitor {
	if m != nil {
		return m.CredentialMonitor
	}
	return nil
}

func (m *CredentialMonitorResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CredentialMonitor struct {
	Certificate          string            `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               CertificateStatus `protobuf:"varint,2,opt,name=status,proto3,enum=moc.wssdhostagent.admin.CertificateStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CredentialMonitor) Reset()         { *m = CredentialMonitor{} }
func (m *CredentialMonitor) String() string { return proto.CompactTextString(m) }
func (*CredentialMonitor) ProtoMessage()    {}
func (*CredentialMonitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_13347e86c60e2a65, []int{2}
}

func (m *CredentialMonitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialMonitor.Unmarshal(m, b)
}
func (m *CredentialMonitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialMonitor.Marshal(b, m, deterministic)
}
func (m *CredentialMonitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialMonitor.Merge(m, src)
}
func (m *CredentialMonitor) XXX_Size() int {
	return xxx_messageInfo_CredentialMonitor.Size(m)
}
func (m *CredentialMonitor) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialMonitor.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialMonitor proto.InternalMessageInfo

func (m *CredentialMonitor) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *CredentialMonitor) GetStatus() CertificateStatus {
	if m != nil {
		return m.Status
	}
	return CertificateStatus_Single
}

func init() {
	proto.RegisterEnum("moc.wssdhostagent.admin.CertificateStatus", CertificateStatus_name, CertificateStatus_value)
	proto.RegisterType((*CredentialMonitorRequest)(nil), "moc.wssdhostagent.admin.CredentialMonitorRequest")
	proto.RegisterType((*CredentialMonitorResponse)(nil), "moc.wssdhostagent.admin.CredentialMonitorResponse")
	proto.RegisterType((*CredentialMonitor)(nil), "moc.wssdhostagent.admin.CredentialMonitor")
}

func init() {
	proto.RegisterFile("moc_wssdhostagent_credentialmonitor.proto", fileDescriptor_13347e86c60e2a65)
}

var fileDescriptor_13347e86c60e2a65 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0x9b, 0xef, 0xd3, 0x4a, 0xa7, 0x20, 0x6d, 0x10, 0x5d, 0x7b, 0x2a, 0x7b, 0x90, 0x5a,
	0x64, 0x57, 0xd7, 0x5f, 0x60, 0x45, 0x3c, 0x89, 0xb0, 0xbd, 0x88, 0x97, 0x92, 0xa6, 0x69, 0x1b,
	0x68, 0x32, 0x6b, 0x32, 0xd5, 0xa3, 0xd7, 0x82, 0xbf, 0xc7, 0xff, 0x27, 0xdd, 0x2d, 0x6a, 0xd9,
	0x2a, 0xf4, 0xe0, 0x69, 0x59, 0xde, 0xbc, 0xcf, 0x33, 0x64, 0x02, 0xa7, 0x06, 0xe5, 0xe0, 0xc5,
	0xfb, 0xd1, 0x14, 0x3d, 0x89, 0x89, 0xb2, 0x34, 0x90, 0x4e, 0x8d, 0x94, 0x25, 0x2d, 0x66, 0x06,
	0xad, 0x26, 0x74, 0x51, 0xe6, 0x90, 0x90, 0x1f, 0x19, 0x94, 0xd1, 0xda, 0xd1, 0x48, 0x8c, 0x8c,
	0xb6, 0xad, 0x65, 0x30, 0x90, 0x68, 0x0c, 0xda, 0xd5, 0xa7, 0x68, 0x84, 0x04, 0xc1, 0xf5, 0x27,
	0xec, 0xae, 0x80, 0xa5, 0xea, 0x69, 0xae, 0x3c, 0xf1, 0x07, 0x68, 0x96, 0xb2, 0x80, 0xb5, 0x59,
	0xa7, 0x9e, 0x74, 0xa3, 0x1f, 0x4c, 0x51, 0x99, 0x56, 0x86, 0x84, 0x6f, 0x0c, 0x8e, 0x37, 0x68,
	0x7d, 0x86, 0xd6, 0xab, 0xbf, 0xf3, 0xf2, 0x03, 0xd8, 0xbd, 0x71, 0x0e, 0x5d, 0xf0, 0xaf, 0xcd,
	0x3a, 0xb5, 0xb4, 0xf8, 0x09, 0x5f, 0x37, 0xf8, 0xf8, 0x09, 0xd4, 0xa5, 0x72, 0xa4, 0xc7, 0x5a,
	0x0a, 0x52, 0xb9, 0xbe, 0xd6, 0xdb, 0x59, 0xbc, 0x07, 0x2c, 0xfd, 0x1e, 0xf0, 0x1e, 0x54, 0x3d,
	0x09, 0x9a, 0xfb, 0x9c, 0xb9, 0xff, 0xdb, 0x84, 0x5f, 0xad, 0x7e, 0xde, 0x48, 0x57, 0xcd, 0xee,
	0x19, 0x34, 0x4b, 0x21, 0x07, 0xa8, 0xf6, 0xb5, 0x9d, 0xcc, 0x54, 0xa3, 0xc2, 0xeb, 0xb0, 0x77,
	0xff, 0xac, 0xdc, 0x4c, 0x64, 0x0d, 0x96, 0x2c, 0x18, 0x1c, 0x96, 0xe6, 0xbd, 0x5a, 0x8a, 0xb8,
	0x85, 0xff, 0xb7, 0x8a, 0xf8, 0xc5, 0x16, 0xb7, 0x54, 0xec, 0xba, 0x95, 0x6c, 0x53, 0x29, 0xf6,
	0x14, 0x56, 0x7a, 0xc9, 0xe3, 0xf9, 0x44, 0xd3, 0x74, 0x3e, 0x8c, 0x24, 0x9a, 0xd8, 0x68, 0xe9,
	0xd0, 0xe3, 0x98, 0x62, 0x83, 0x32, 0x76, 0x99, 0x8c, 0xd7, 0x78, 0x71, 0xce, 0x1b, 0x56, 0xf3,
	0x87, 0x77, 0xf9, 0x11, 0x00, 0x00, 0xff, 0xff, 0x59, 0xd0, 0x38, 0x39, 0xd7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CredentialMonitorAgentClient is the client API for CredentialMonitorAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CredentialMonitorAgentClient interface {
	Get(ctx context.Context, in *CredentialMonitorRequest, opts ...grpc.CallOption) (*CredentialMonitorResponse, error)
}

type credentialMonitorAgentClient struct {
	cc *grpc.ClientConn
}

func NewCredentialMonitorAgentClient(cc *grpc.ClientConn) CredentialMonitorAgentClient {
	return &credentialMonitorAgentClient{cc}
}

func (c *credentialMonitorAgentClient) Get(ctx context.Context, in *CredentialMonitorRequest, opts ...grpc.CallOption) (*CredentialMonitorResponse, error) {
	out := new(CredentialMonitorResponse)
	err := c.cc.Invoke(ctx, "/moc.wssdhostagent.admin.CredentialMonitorAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialMonitorAgentServer is the server API for CredentialMonitorAgent service.
type CredentialMonitorAgentServer interface {
	Get(context.Context, *CredentialMonitorRequest) (*CredentialMonitorResponse, error)
}

// UnimplementedCredentialMonitorAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCredentialMonitorAgentServer struct {
}

func (*UnimplementedCredentialMonitorAgentServer) Get(ctx context.Context, req *CredentialMonitorRequest) (*CredentialMonitorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterCredentialMonitorAgentServer(s *grpc.Server, srv CredentialMonitorAgentServer) {
	s.RegisterService(&_CredentialMonitorAgent_serviceDesc, srv)
}

func _CredentialMonitorAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialMonitorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialMonitorAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.wssdhostagent.admin.CredentialMonitorAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialMonitorAgentServer).Get(ctx, req.(*CredentialMonitorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CredentialMonitorAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.wssdhostagent.admin.CredentialMonitorAgent",
	HandlerType: (*CredentialMonitorAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CredentialMonitorAgent_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_wssdhostagent_credentialmonitor.proto",
}
