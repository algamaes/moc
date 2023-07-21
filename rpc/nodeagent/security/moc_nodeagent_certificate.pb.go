// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_nodeagent_certificate.proto

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
	return fileDescriptor_ae3bff7543638e48, []int{0}
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
	return fileDescriptor_ae3bff7543638e48, []int{0}
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
	return fileDescriptor_ae3bff7543638e48, []int{1}
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
	Type                 CertificateType `protobuf:"varint,7,opt,name=type,proto3,enum=moc.nodeagent.security.CertificateType" json:"type,omitempty"`
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
	return fileDescriptor_ae3bff7543638e48, []int{2}
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

type CSRRequest struct {
	CSRs                 []*CertificateSigningRequest `protobuf:"bytes,1,rep,name=CSRs,proto3" json:"CSRs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CSRRequest) Reset()         { *m = CSRRequest{} }
func (m *CSRRequest) String() string { return proto.CompactTextString(m) }
func (*CSRRequest) ProtoMessage()    {}
func (*CSRRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae3bff7543638e48, []int{3}
}

func (m *CSRRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSRRequest.Unmarshal(m, b)
}
func (m *CSRRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSRRequest.Marshal(b, m, deterministic)
}
func (m *CSRRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSRRequest.Merge(m, src)
}
func (m *CSRRequest) XXX_Size() int {
	return xxx_messageInfo_CSRRequest.Size(m)
}
func (m *CSRRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CSRRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CSRRequest proto.InternalMessageInfo

func (m *CSRRequest) GetCSRs() []*CertificateSigningRequest {
	if m != nil {
		return m.CSRs
	}
	return nil
}

type CertificateSigningRequest struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Csr                  string              `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	OldCertificate       string              `protobuf:"bytes,3,opt,name=oldCertificate,proto3" json:"oldCertificate,omitempty"`
	Status               *common.Status      `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CaName               string              `protobuf:"bytes,5,opt,name=caName,proto3" json:"caName,omitempty"`
	Validity             int64               `protobuf:"varint,6,opt,name=validity,proto3" json:"validity,omitempty"`
	ServerAuth           *wrappers.BoolValue `protobuf:"bytes,7,opt,name=serverAuth,proto3" json:"serverAuth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CertificateSigningRequest) Reset()         { *m = CertificateSigningRequest{} }
func (m *CertificateSigningRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateSigningRequest) ProtoMessage()    {}
func (*CertificateSigningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae3bff7543638e48, []int{4}
}

func (m *CertificateSigningRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateSigningRequest.Unmarshal(m, b)
}
func (m *CertificateSigningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateSigningRequest.Marshal(b, m, deterministic)
}
func (m *CertificateSigningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateSigningRequest.Merge(m, src)
}
func (m *CertificateSigningRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateSigningRequest.Size(m)
}
func (m *CertificateSigningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateSigningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateSigningRequest proto.InternalMessageInfo

func (m *CertificateSigningRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CertificateSigningRequest) GetCsr() string {
	if m != nil {
		return m.Csr
	}
	return ""
}

func (m *CertificateSigningRequest) GetOldCertificate() string {
	if m != nil {
		return m.OldCertificate
	}
	return ""
}

func (m *CertificateSigningRequest) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateSigningRequest) GetCaName() string {
	if m != nil {
		return m.CaName
	}
	return ""
}

func (m *CertificateSigningRequest) GetValidity() int64 {
	if m != nil {
		return m.Validity
	}
	return 0
}

func (m *CertificateSigningRequest) GetServerAuth() *wrappers.BoolValue {
	if m != nil {
		return m.ServerAuth
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.nodeagent.security.CertificateType", CertificateType_name, CertificateType_value)
	proto.RegisterType((*CertificateRequest)(nil), "moc.nodeagent.security.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "moc.nodeagent.security.CertificateResponse")
	proto.RegisterType((*Certificate)(nil), "moc.nodeagent.security.Certificate")
	proto.RegisterType((*CSRRequest)(nil), "moc.nodeagent.security.CSRRequest")
	proto.RegisterType((*CertificateSigningRequest)(nil), "moc.nodeagent.security.CertificateSigningRequest")
}

func init() { proto.RegisterFile("moc_nodeagent_certificate.proto", fileDescriptor_ae3bff7543638e48) }

var fileDescriptor_ae3bff7543638e48 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xad, 0x13, 0xd7, 0x5f, 0x33, 0xf9, 0x14, 0xaa, 0x05, 0x15, 0x13, 0xf1, 0x13, 0xb9, 0x12,
	0x84, 0x82, 0x6c, 0x35, 0xdc, 0xc1, 0x55, 0x13, 0xaa, 0xde, 0x81, 0xb4, 0x29, 0x48, 0x54, 0x42,
	0xc5, 0xdd, 0x4c, 0x5c, 0x0b, 0x7b, 0xd7, 0xec, 0xae, 0x5b, 0xe5, 0x0d, 0x10, 0x4f, 0xc2, 0x0b,
	0x70, 0xc1, 0xc3, 0x21, 0x21, 0xaf, 0x9d, 0xd4, 0xad, 0x5a, 0x14, 0x24, 0x72, 0x65, 0xcf, 0xcc,
	0x99, 0x33, 0xbb, 0xe7, 0xcc, 0xc2, 0xa3, 0x54, 0xb0, 0x63, 0x2e, 0x26, 0x18, 0x46, 0xc8, 0xf5,
	0x31, 0x43, 0xa9, 0xe3, 0x69, 0xcc, 0x42, 0x8d, 0x7e, 0x26, 0x85, 0x16, 0x64, 0x2b, 0x15, 0xcc,
	0x5f, 0x00, 0x7c, 0x85, 0x2c, 0x97, 0xb1, 0x9e, 0x75, 0x1f, 0x46, 0x42, 0x44, 0x09, 0x06, 0x06,
	0x75, 0x92, 0x4f, 0x83, 0x73, 0x19, 0x66, 0x19, 0x4a, 0x55, 0xf6, 0x75, 0xef, 0x16, 0xc4, 0x4c,
	0xa4, 0xa9, 0xe0, 0xd5, 0xa7, 0x2c, 0x78, 0x1f, 0x81, 0x8c, 0x2e, 0xa6, 0x50, 0xfc, 0x92, 0xa3,
	0xd2, 0xe4, 0x00, 0xfe, 0xaf, 0x65, 0x95, 0x6b, 0xf5, 0x9a, 0xfd, 0xf6, 0x60, 0xdb, 0xbf, 0x7e,
	0xba, 0x5f, 0x67, 0xb8, 0xd4, 0xe8, 0x7d, 0xb7, 0xe0, 0xf6, 0x25, 0x7e, 0x95, 0x09, 0xae, 0xf0,
	0x9f, 0x0d, 0x20, 0x03, 0x70, 0x28, 0xaa, 0x3c, 0xd1, 0x6e, 0xa3, 0x67, 0xf5, 0xdb, 0x83, 0xae,
	0x5f, 0x2a, 0xe1, 0xcf, 0x95, 0xf0, 0x87, 0x42, 0x24, 0xef, 0xc3, 0x24, 0x47, 0x5a, 0x21, 0xc9,
	0x1d, 0x58, 0xdf, 0x97, 0x52, 0x48, 0xb7, 0xd9, 0xb3, 0xfa, 0x2d, 0x5a, 0x06, 0xde, 0xcf, 0x06,
	0xb4, 0x6b, 0xd4, 0x84, 0x80, 0xcd, 0xc3, 0x14, 0x5d, 0xcb, 0x80, 0xcc, 0x3f, 0xe9, 0x40, 0x23,
	0x9e, 0x98, 0x49, 0x2d, 0xda, 0x88, 0x27, 0xe4, 0x3e, 0xb4, 0xb8, 0xd0, 0x43, 0x9c, 0x0a, 0x89,
	0x86, 0xad, 0x49, 0x2f, 0x12, 0xa4, 0x0b, 0x1b, 0x5c, 0xe8, 0xbd, 0xa9, 0x46, 0xe9, 0xda, 0xa6,
	0xb8, 0x88, 0xc9, 0x63, 0x68, 0xd7, 0xdc, 0x75, 0xd7, 0x0b, 0xca, 0xa1, 0xfd, 0xf5, 0x87, 0x6b,
	0xd1, 0x7a, 0x81, 0x6c, 0x83, 0xa3, 0x74, 0xa8, 0x73, 0xe5, 0x3a, 0xe6, 0x7e, 0x6d, 0x23, 0xd1,
	0xd8, 0xa4, 0x68, 0x55, 0x22, 0xaf, 0xc0, 0xd6, 0xb3, 0x0c, 0xdd, 0xff, 0x7a, 0x56, 0xbf, 0x33,
	0x78, 0xb2, 0x84, 0x8a, 0x87, 0xb3, 0x0c, 0xa9, 0x69, 0x2a, 0x26, 0x20, 0xd7, 0xb1, 0x9e, 0xb9,
	0x1b, 0xb5, 0x09, 0xfb, 0x26, 0x45, 0xab, 0x12, 0x79, 0x00, 0xb6, 0x0e, 0x23, 0xe5, 0xb6, 0x0c,
	0xa4, 0x65, 0x20, 0x87, 0x61, 0xa4, 0xa8, 0x49, 0x7b, 0x63, 0x80, 0xd1, 0x98, 0xce, 0xb7, 0x67,
	0x1f, 0xec, 0xd1, 0x98, 0xce, 0x4d, 0xdd, 0x5d, 0xe2, 0x38, 0xe3, 0x38, 0xe2, 0x31, 0x8f, 0x2a,
	0x02, 0x6a, 0xda, 0xbd, 0x6f, 0x0d, 0xb8, 0x77, 0x23, 0xe6, 0x5a, 0x7b, 0xb6, 0xa0, 0xc9, 0x94,
	0x2c, 0xfd, 0xa9, 0xc4, 0x2c, 0x12, 0xe4, 0x39, 0x74, 0x44, 0x32, 0xa9, 0x71, 0x95, 0xce, 0x57,
	0x90, 0x2b, 0xb5, 0x9a, 0xe4, 0xf6, 0xcd, 0x92, 0x6f, 0x81, 0xc3, 0xc2, 0x37, 0xc5, 0x01, 0x8c,
	0x75, 0xb4, 0x8a, 0x0a, 0xcf, 0xcf, 0xc2, 0x24, 0x9e, 0x14, 0x7a, 0x3a, 0xa5, 0xe7, 0xf3, 0x98,
	0xbc, 0x04, 0x50, 0x28, 0xcf, 0x50, 0xee, 0xe5, 0xfa, 0xd4, 0x98, 0xf5, 0xe7, 0x7d, 0xad, 0xa1,
	0x77, 0x9e, 0xc2, 0xad, 0x2b, 0xf6, 0x11, 0x00, 0x67, 0x94, 0xc4, 0xc8, 0xf5, 0xe6, 0x5a, 0xf1,
	0x3f, 0x36, 0xe0, 0x4d, 0x6b, 0xf0, 0xab, 0x09, 0x9b, 0x35, 0xec, 0x5e, 0xa1, 0x3a, 0xf9, 0x0c,
	0x9d, 0x91, 0xc4, 0x50, 0xe3, 0x5b, 0xf9, 0x2e, 0x9b, 0x14, 0xd7, 0xdc, 0x59, 0xe6, 0xb1, 0x95,
	0x62, 0x77, 0x9f, 0x2d, 0x85, 0x2d, 0xdf, 0xb6, 0xb7, 0x46, 0x3e, 0x41, 0xf3, 0x00, 0xf5, 0x2a,
	0x27, 0x30, 0x70, 0x5e, 0x63, 0x82, 0xab, 0xbd, 0xc6, 0x07, 0xb0, 0x8b, 0xa5, 0x23, 0xde, 0x8d,
	0x6d, 0x8b, 0x9d, 0xff, 0x5b, 0xea, 0x23, 0x58, 0xa7, 0xc8, 0xf1, 0x7c, 0x05, 0xdc, 0xc3, 0xdd,
	0xa3, 0x20, 0x8a, 0xf5, 0x69, 0x7e, 0xe2, 0x33, 0x91, 0x06, 0x69, 0xcc, 0xa4, 0x50, 0x62, 0xaa,
	0x83, 0x54, 0xb0, 0x40, 0x66, 0x2c, 0x58, 0x10, 0x05, 0x73, 0xa2, 0x13, 0xc7, 0x6c, 0xdf, 0x8b,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x92, 0x9f, 0xcb, 0x80, 0x06, 0x00, 0x00,
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
	Sign(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Renew(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateAgentClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAgentClient(cc *grpc.ClientConn) CertificateAgentClient {
	return &certificateAgentClient{cc}
}

func (c *certificateAgentClient) CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Sign(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Renew(ctx context.Context, in *CSRRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/Renew", in, out, opts...)
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
	Sign(context.Context, *CSRRequest) (*CertificateResponse, error)
	Renew(context.Context, *CSRRequest) (*CertificateResponse, error)
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
func (*UnimplementedCertificateAgentServer) Sign(ctx context.Context, req *CSRRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (*UnimplementedCertificateAgentServer) Renew(ctx context.Context, req *CSRRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Renew not implemented")
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
		FullMethod: "/moc.nodeagent.security.CertificateAgent/CreateOrUpdate",
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
		FullMethod: "/moc.nodeagent.security.CertificateAgent/Get",
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
		FullMethod: "/moc.nodeagent.security.CertificateAgent/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Delete(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.CertificateAgent/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Sign(ctx, req.(*CSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.CertificateAgent/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Renew(ctx, req.(*CSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.security.CertificateAgent",
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
		{
			MethodName: "Sign",
			Handler:    _CertificateAgent_Sign_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _CertificateAgent_Renew_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_nodeagent_certificate.proto",
}
