// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_wssdhostagent_certificate.proto

package security

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

type CertificateType int32

const (
	CertificateType_Client CertificateType = 0
	CertificateType_Server CertificateType = 1
)

var CertificateType_name = map[int32]string{
	0: "Client",
	1: "Server",
}

var CertificateType_value = map[string]int32{
	"Client": 0,
	"Server": 1,
}

func (x CertificateType) String() string {
	return proto.EnumName(CertificateType_name, int32(x))
}

func (CertificateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b8e52b6eda84a90f, []int{0}
}

type CertificateRequest struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8e52b6eda84a90f, []int{0}
}

func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}
func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}
func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}
func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}
func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CertificateResponse struct {
	Certificates         []*Certificate      `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8e52b6eda84a90f, []int{1}
}

func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *CertificateResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CertificateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Certificate struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NotBefore            int64           `protobuf:"varint,3,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
	NotAfter             int64           `protobuf:"varint,4,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
	Certificate          string          `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               *common.Status  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Type                 CertificateType `protobuf:"varint,7,opt,name=type,proto3,enum=moc.wssdhostagent.security.CertificateType" json:"type,omitempty"`
	Entity               *common.Entity  `protobuf:"bytes,8,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags    `protobuf:"bytes,9,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8e52b6eda84a90f, []int{2}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Certificate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Certificate) GetNotBefore() int64 {
	if m != nil {
		return m.NotBefore
	}
	return 0
}

func (m *Certificate) GetNotAfter() int64 {
	if m != nil {
		return m.NotAfter
	}
	return 0
}

func (m *Certificate) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *Certificate) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Certificate) GetType() CertificateType {
	if m != nil {
		return m.Type
	}
	return CertificateType_Client
}

func (m *Certificate) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Certificate) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.wssdhostagent.security.CertificateType", CertificateType_name, CertificateType_value)
	proto.RegisterType((*CertificateRequest)(nil), "moc.wssdhostagent.security.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "moc.wssdhostagent.security.CertificateResponse")
	proto.RegisterType((*Certificate)(nil), "moc.wssdhostagent.security.Certificate")
}

func init() {
	proto.RegisterFile("moc_wssdhostagent_certificate.proto", fileDescriptor_b8e52b6eda84a90f)
}

var fileDescriptor_b8e52b6eda84a90f = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x93, 0xd4, 0x34, 0x13, 0x14, 0xa2, 0x05, 0x89, 0x95, 0x05, 0x28, 0x4a, 0x25, 0x08,
	0x20, 0xd9, 0x52, 0x2a, 0xce, 0xa8, 0x09, 0x15, 0x07, 0x0e, 0x48, 0xdb, 0xc2, 0x81, 0x4b, 0xe5,
	0x38, 0x63, 0xc7, 0xc2, 0xf6, 0x98, 0xdd, 0x31, 0x55, 0xfe, 0x80, 0x8f, 0xe1, 0xca, 0x91, 0x5f,
	0xe1, 0x5b, 0x90, 0xd7, 0x29, 0x75, 0x90, 0x90, 0x72, 0x40, 0x3d, 0x65, 0xf7, 0xbd, 0x37, 0xef,
	0xed, 0x4c, 0xc6, 0x70, 0x9c, 0x53, 0x74, 0x79, 0x65, 0xcc, 0x6a, 0x4d, 0x86, 0xc3, 0x04, 0x0b,
	0xbe, 0x8c, 0x50, 0x73, 0x1a, 0xa7, 0x51, 0xc8, 0xe8, 0x97, 0x9a, 0x98, 0x84, 0x97, 0x53, 0xe4,
	0xef, 0x88, 0x7c, 0x83, 0x51, 0xa5, 0x53, 0xde, 0x78, 0x4f, 0x12, 0xa2, 0x24, 0xc3, 0xc0, 0x2a,
	0x97, 0x55, 0x1c, 0x5c, 0xe9, 0xb0, 0x2c, 0x51, 0x9b, 0xa6, 0xd6, 0x7b, 0x58, 0x07, 0x44, 0x94,
	0xe7, 0x54, 0x6c, 0x7f, 0x1a, 0x62, 0x12, 0x82, 0x58, 0xdc, 0x24, 0x29, 0xfc, 0x52, 0xa1, 0x61,
	0xf1, 0x0e, 0xee, 0xb6, 0x50, 0x23, 0x9d, 0x71, 0x77, 0x3a, 0x98, 0x3d, 0xf3, 0xff, 0xfd, 0x02,
	0xbf, 0xed, 0xb2, 0x53, 0x3c, 0xf9, 0xee, 0xc0, 0xfd, 0x9d, 0x0c, 0x53, 0x52, 0x61, 0xf0, 0xbf,
	0x86, 0x88, 0x19, 0xb8, 0x0a, 0x4d, 0x95, 0xb1, 0xec, 0x8c, 0x9d, 0xe9, 0x60, 0xe6, 0xf9, 0xcd,
	0x44, 0xfc, 0xeb, 0x89, 0xf8, 0x73, 0xa2, 0xec, 0x63, 0x98, 0x55, 0xa8, 0xb6, 0x4a, 0xf1, 0x00,
	0x0e, 0xcf, 0xb4, 0x26, 0x2d, 0xbb, 0x63, 0x67, 0xda, 0x57, 0xcd, 0x65, 0xf2, 0xb3, 0x03, 0x83,
	0x96, 0xb5, 0x10, 0xd0, 0x2b, 0xc2, 0x1c, 0xa5, 0x63, 0x45, 0xf6, 0x2c, 0x86, 0xd0, 0x49, 0x57,
	0x36, 0xa9, 0xaf, 0x3a, 0xe9, 0x4a, 0x3c, 0x82, 0x7e, 0x41, 0x3c, 0xc7, 0x98, 0x34, 0x5a, 0xb7,
	0xae, 0xba, 0x01, 0x84, 0x07, 0x47, 0x05, 0xf1, 0x69, 0xcc, 0xa8, 0x65, 0xcf, 0x92, 0x7f, 0xee,
	0xe2, 0x29, 0x0c, 0x5a, 0xff, 0xb4, 0x3c, 0xac, 0x2d, 0xe7, 0xbd, 0x6f, 0x3f, 0xa4, 0xa3, 0xda,
	0x84, 0x38, 0x06, 0xd7, 0x70, 0xc8, 0x95, 0x91, 0xae, 0xed, 0x6f, 0x60, 0xc7, 0x74, 0x6e, 0x21,
	0xb5, 0xa5, 0xc4, 0x6b, 0xe8, 0xf1, 0xa6, 0x44, 0x79, 0x67, 0xec, 0x4c, 0x87, 0xb3, 0x97, 0x7b,
	0x4e, 0xf2, 0x62, 0x53, 0xa2, 0xb2, 0x85, 0x75, 0x0a, 0x16, 0x9c, 0xf2, 0x46, 0x1e, 0xb5, 0x52,
	0xce, 0x2c, 0xa4, 0xb6, 0x94, 0x78, 0x0c, 0x3d, 0x0e, 0x13, 0x23, 0xfb, 0x56, 0xd2, 0xb7, 0x92,
	0x8b, 0x30, 0x31, 0xca, 0xc2, 0x2f, 0x9e, 0xc3, 0xbd, 0xbf, 0xcc, 0x05, 0x80, 0xbb, 0xc8, 0x52,
	0x2c, 0x78, 0x74, 0x50, 0x9f, 0xcf, 0x51, 0x7f, 0x45, 0x3d, 0x72, 0x66, 0xbf, 0x3a, 0x30, 0x6a,
	0x69, 0x4f, 0xeb, 0x27, 0x0a, 0x03, 0xc3, 0x85, 0xc6, 0x90, 0xf1, 0xbd, 0xfe, 0x50, 0xae, 0xea,
	0xde, 0xfd, 0x7d, 0x57, 0xa2, 0xd9, 0x5e, 0x2f, 0xd8, 0x5b, 0xdf, 0x6c, 0xe2, 0xe4, 0x40, 0xac,
	0xa1, 0xfb, 0x16, 0xf9, 0x36, 0x92, 0x3e, 0x83, 0xfb, 0x06, 0x33, 0xbc, 0x95, 0xb6, 0xe6, 0xaf,
	0x3e, 0x9d, 0x24, 0x29, 0xaf, 0xab, 0xa5, 0x1f, 0x51, 0x1e, 0xe4, 0x69, 0xa4, 0xc9, 0x50, 0xcc,
	0x41, 0x4e, 0x51, 0xa0, 0xcb, 0x28, 0xd8, 0x31, 0x0b, 0xae, 0xcd, 0x96, 0xae, 0xfd, 0x68, 0x4e,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x88, 0x3f, 0x57, 0x97, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertificateAgentClient is the client API for CertificateAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateAgentClient interface {
	CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateAgentClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAgentClient(cc *grpc.ClientConn) CertificateAgentClient {
	return &certificateAgentClient{cc}
}

func (c *certificateAgentClient) CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.wssdhostagent.security.CertificateAgent/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.wssdhostagent.security.CertificateAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.wssdhostagent.security.CertificateAgent/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAgentServer is the server API for CertificateAgent service.
type CertificateAgentServer interface {
	CreateOrUpdate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Get(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Delete(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

// UnimplementedCertificateAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAgentServer struct {
}

func (*UnimplementedCertificateAgentServer) CreateOrUpdate(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdate not implemented")
}
func (*UnimplementedCertificateAgentServer) Get(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCertificateAgentServer) Delete(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCertificateAgentServer(s *grpc.Server, srv CertificateAgentServer) {
	s.RegisterService(&_CertificateAgent_serviceDesc, srv)
}

func _CertificateAgent_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.wssdhostagent.security.CertificateAgent/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.wssdhostagent.security.CertificateAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Get(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.wssdhostagent.security.CertificateAgent/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Delete(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.wssdhostagent.security.CertificateAgent",
	HandlerType: (*CertificateAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _CertificateAgent_CreateOrUpdate_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CertificateAgent_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateAgent_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_wssdhostagent_certificate.proto",
}
