// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtualmachine/moc_mocguestagent_virtualmachine.proto

package compute

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type VirtualMachineRunCommandRequest struct {
	Source                    *common.VirtualMachineRunCommandScriptSource     `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	RunCommandInputParameters []*common.VirtualMachineRunCommandInputParameter `protobuf:"bytes,2,rep,name=RunCommandInputParameters,proto3" json:"RunCommandInputParameters,omitempty"`
	OperationID               string                                           `protobuf:"bytes,3,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	RunAsUser                 string                                           `protobuf:"bytes,4,opt,name=RunAsUser,proto3" json:"RunAsUser,omitempty"`
	RunAsPassword             string                                           `protobuf:"bytes,5,opt,name=RunAsPassword,proto3" json:"RunAsPassword,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                                         `json:"-"`
	XXX_unrecognized          []byte                                           `json:"-"`
	XXX_sizecache             int32                                            `json:"-"`
}

func (m *VirtualMachineRunCommandRequest) Reset()         { *m = VirtualMachineRunCommandRequest{} }
func (m *VirtualMachineRunCommandRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineRunCommandRequest) ProtoMessage()    {}
func (*VirtualMachineRunCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44dd44b4a8829e1, []int{0}
}

func (m *VirtualMachineRunCommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineRunCommandRequest.Unmarshal(m, b)
}
func (m *VirtualMachineRunCommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineRunCommandRequest.Marshal(b, m, deterministic)
}
func (m *VirtualMachineRunCommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineRunCommandRequest.Merge(m, src)
}
func (m *VirtualMachineRunCommandRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineRunCommandRequest.Size(m)
}
func (m *VirtualMachineRunCommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineRunCommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineRunCommandRequest proto.InternalMessageInfo

func (m *VirtualMachineRunCommandRequest) GetSource() *common.VirtualMachineRunCommandScriptSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *VirtualMachineRunCommandRequest) GetRunCommandInputParameters() []*common.VirtualMachineRunCommandInputParameter {
	if m != nil {
		return m.RunCommandInputParameters
	}
	return nil
}

func (m *VirtualMachineRunCommandRequest) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

func (m *VirtualMachineRunCommandRequest) GetRunAsUser() string {
	if m != nil {
		return m.RunAsUser
	}
	return ""
}

func (m *VirtualMachineRunCommandRequest) GetRunAsPassword() string {
	if m != nil {
		return m.RunAsPassword
	}
	return ""
}

type VirtualMachineCommandResultRequest struct {
	OperationID          string   `protobuf:"bytes,1,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMachineCommandResultRequest) Reset()         { *m = VirtualMachineCommandResultRequest{} }
func (m *VirtualMachineCommandResultRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineCommandResultRequest) ProtoMessage()    {}
func (*VirtualMachineCommandResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44dd44b4a8829e1, []int{1}
}

func (m *VirtualMachineCommandResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineCommandResultRequest.Unmarshal(m, b)
}
func (m *VirtualMachineCommandResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineCommandResultRequest.Marshal(b, m, deterministic)
}
func (m *VirtualMachineCommandResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineCommandResultRequest.Merge(m, src)
}
func (m *VirtualMachineCommandResultRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineCommandResultRequest.Size(m)
}
func (m *VirtualMachineCommandResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineCommandResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineCommandResultRequest proto.InternalMessageInfo

func (m *VirtualMachineCommandResultRequest) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type VirtualMachineRunCommandResponse struct {
	InstanceView         *common.VirtualMachineRunCommandInstanceView `protobuf:"bytes,1,opt,name=InstanceView,proto3" json:"InstanceView,omitempty"`
	OperationID          string                                       `protobuf:"bytes,2,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *VirtualMachineRunCommandResponse) Reset()         { *m = VirtualMachineRunCommandResponse{} }
func (m *VirtualMachineRunCommandResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineRunCommandResponse) ProtoMessage()    {}
func (*VirtualMachineRunCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44dd44b4a8829e1, []int{2}
}

func (m *VirtualMachineRunCommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineRunCommandResponse.Unmarshal(m, b)
}
func (m *VirtualMachineRunCommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineRunCommandResponse.Marshal(b, m, deterministic)
}
func (m *VirtualMachineRunCommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineRunCommandResponse.Merge(m, src)
}
func (m *VirtualMachineRunCommandResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineRunCommandResponse.Size(m)
}
func (m *VirtualMachineRunCommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineRunCommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineRunCommandResponse proto.InternalMessageInfo

func (m *VirtualMachineRunCommandResponse) GetInstanceView() *common.VirtualMachineRunCommandInstanceView {
	if m != nil {
		return m.InstanceView
	}
	return nil
}

func (m *VirtualMachineRunCommandResponse) GetOperationID() string {
	if m != nil {
		return m.OperationID
	}
	return ""
}

type UpdateAgentRequest struct {
	AgentBinary          []byte   `protobuf:"bytes,1,opt,name=AgentBinary,proto3" json:"AgentBinary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAgentRequest) Reset()         { *m = UpdateAgentRequest{} }
func (m *UpdateAgentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAgentRequest) ProtoMessage()    {}
func (*UpdateAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44dd44b4a8829e1, []int{3}
}

func (m *UpdateAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAgentRequest.Unmarshal(m, b)
}
func (m *UpdateAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAgentRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAgentRequest.Merge(m, src)
}
func (m *UpdateAgentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAgentRequest.Size(m)
}
func (m *UpdateAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAgentRequest proto.InternalMessageInfo

func (m *UpdateAgentRequest) GetAgentBinary() []byte {
	if m != nil {
		return m.AgentBinary
	}
	return nil
}

type UpdateAgentResponse struct {
	Result               *wrappers.BoolValue `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpdateAgentResponse) Reset()         { *m = UpdateAgentResponse{} }
func (m *UpdateAgentResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateAgentResponse) ProtoMessage()    {}
func (*UpdateAgentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44dd44b4a8829e1, []int{4}
}

func (m *UpdateAgentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAgentResponse.Unmarshal(m, b)
}
func (m *UpdateAgentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAgentResponse.Marshal(b, m, deterministic)
}
func (m *UpdateAgentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAgentResponse.Merge(m, src)
}
func (m *UpdateAgentResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateAgentResponse.Size(m)
}
func (m *UpdateAgentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAgentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAgentResponse proto.InternalMessageInfo

func (m *UpdateAgentResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UpdateAgentResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*VirtualMachineRunCommandRequest)(nil), "moc.mocguestagent.compute.VirtualMachineRunCommandRequest")
	proto.RegisterType((*VirtualMachineCommandResultRequest)(nil), "moc.mocguestagent.compute.VirtualMachineCommandResultRequest")
	proto.RegisterType((*VirtualMachineRunCommandResponse)(nil), "moc.mocguestagent.compute.VirtualMachineRunCommandResponse")
	proto.RegisterType((*UpdateAgentRequest)(nil), "moc.mocguestagent.compute.UpdateAgentRequest")
	proto.RegisterType((*UpdateAgentResponse)(nil), "moc.mocguestagent.compute.UpdateAgentResponse")
}

func init() {
	proto.RegisterFile("virtualmachine/moc_mocguestagent_virtualmachine.proto", fileDescriptor_c44dd44b4a8829e1)
}

var fileDescriptor_c44dd44b4a8829e1 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x65, 0xfa, 0x88, 0x54, 0xa7, 0x48, 0xe0, 0x22, 0x98, 0x0c, 0x02, 0xa2, 0x59, 0x05, 0x10,
	0x33, 0x52, 0x78, 0x2c, 0x40, 0x2c, 0x92, 0x52, 0x20, 0x8b, 0x42, 0x35, 0x55, 0xb3, 0x60, 0x13,
	0x39, 0x8e, 0x93, 0x58, 0xc4, 0x0f, 0xfc, 0x20, 0xea, 0x17, 0xc0, 0x12, 0x89, 0x6f, 0xe1, 0x03,
	0xf8, 0x1f, 0x3e, 0x02, 0xcd, 0xab, 0x19, 0x27, 0xa4, 0x44, 0xb0, 0x1a, 0xd9, 0xe7, 0x9c, 0x7b,
	0xcf, 0x7d, 0x78, 0xc0, 0xd3, 0xcf, 0x54, 0x19, 0x8b, 0x66, 0x0c, 0xe1, 0x29, 0xe5, 0x24, 0x66,
	0x02, 0x0f, 0x98, 0xc0, 0x13, 0x4b, 0xb4, 0x41, 0x13, 0xc2, 0xcd, 0xc0, 0x25, 0x44, 0x52, 0x09,
	0x23, 0x60, 0x83, 0x09, 0x1c, 0x39, 0xbc, 0x08, 0x0b, 0x26, 0xad, 0x21, 0xc1, 0xed, 0x89, 0x10,
	0x93, 0x19, 0x89, 0x33, 0xe2, 0xd0, 0x8e, 0x63, 0xc2, 0xa4, 0x39, 0xcf, 0x75, 0xc1, 0xdd, 0x65,
	0x70, 0xae, 0x90, 0x94, 0x44, 0xe9, 0x02, 0xbf, 0x95, 0xe6, 0xc7, 0x82, 0x31, 0xc1, 0x8b, 0x4f,
	0x29, 0x74, 0x81, 0x34, 0x93, 0x83, 0xdf, 0xa9, 0xe0, 0x5c, 0x18, 0x3a, 0xa6, 0x18, 0x19, 0x5a,
	0xc2, 0xe1, 0xcf, 0x2d, 0x70, 0xaf, 0x9f, 0x17, 0x72, 0x9c, 0x17, 0x92, 0x58, 0x7e, 0x28, 0x18,
	0x43, 0x7c, 0x94, 0x90, 0x4f, 0x69, 0x15, 0xb0, 0x03, 0x6a, 0xa7, 0xc2, 0x2a, 0x4c, 0x7c, 0xaf,
	0xe9, 0xb5, 0xea, 0xed, 0xfb, 0x69, 0x81, 0xd1, 0x3a, 0xd5, 0x29, 0x56, 0x54, 0x9a, 0x5c, 0x90,
	0x14, 0x42, 0x48, 0x41, 0x63, 0xc1, 0xe8, 0x71, 0x69, 0xcd, 0x09, 0x52, 0x88, 0x11, 0x43, 0x94,
	0xf6, 0xb7, 0x9a, 0xdb, 0xad, 0x7a, 0xfb, 0xe1, 0xa5, 0x51, 0x5d, 0x4d, 0xb2, 0x3e, 0x1a, 0x6c,
	0x82, 0xfa, 0x7b, 0x49, 0x54, 0x56, 0x64, 0xef, 0x95, 0xbf, 0xdd, 0xf4, 0x5a, 0x7b, 0x49, 0xf5,
	0x0a, 0x86, 0x60, 0x2f, 0xb1, 0xbc, 0xa3, 0xcf, 0x34, 0x51, 0xfe, 0x4e, 0x8a, 0x77, 0x77, 0xbe,
	0xfe, 0xf0, 0xbd, 0x64, 0x71, 0x0d, 0x1f, 0x80, 0xab, 0xd9, 0xe1, 0x04, 0x69, 0x3d, 0x17, 0x6a,
	0xe4, 0xef, 0x56, 0x78, 0x2e, 0x14, 0xbe, 0x06, 0xa1, 0x6b, 0xfb, 0xa2, 0x7f, 0xda, 0xce, 0x4c,
	0xd9, 0xc5, 0x25, 0x5f, 0xde, 0x8a, 0xaf, 0xf0, 0xbb, 0x07, 0x9a, 0xeb, 0x67, 0xa1, 0xa5, 0xe0,
	0x9a, 0xc0, 0x63, 0xb0, 0xdf, 0xe3, 0xda, 0x20, 0x8e, 0x49, 0x9f, 0x92, 0xf9, 0x46, 0x23, 0xa9,
	0x0a, 0x12, 0x47, 0xbe, 0xec, 0x6a, 0x6b, 0xd5, 0xd5, 0x33, 0x00, 0xcf, 0xe4, 0x08, 0x19, 0xd2,
	0x49, 0xb7, 0xb9, 0x52, 0x4d, 0x76, 0xee, 0x52, 0x8e, 0xd4, 0x79, 0xe6, 0x62, 0x3f, 0xa9, 0x5e,
	0x85, 0x03, 0x70, 0xe0, 0xe8, 0x0a, 0xff, 0x6d, 0x50, 0xcb, 0xfb, 0x52, 0x38, 0x0f, 0xa2, 0x7c,
	0xf3, 0xa3, 0x72, 0xf3, 0xa3, 0xae, 0x10, 0xb3, 0x3e, 0x9a, 0x59, 0x92, 0x14, 0x4c, 0x78, 0x03,
	0xec, 0x1e, 0x29, 0x25, 0x54, 0x61, 0x2f, 0x3f, 0xb4, 0x7f, 0x6d, 0x83, 0x03, 0xb7, 0xe2, 0x2c,
	0x13, 0x7c, 0x0b, 0xae, 0x1f, 0x4e, 0x09, 0xfe, 0xf8, 0xae, 0xb2, 0xed, 0xf0, 0xe6, 0x4a, 0x9a,
	0xa3, 0xf4, 0xf5, 0x05, 0x8d, 0xac, 0x71, 0x55, 0x6a, 0xe9, 0x34, 0xbc, 0x02, 0xbf, 0x78, 0x00,
	0x2c, 0xba, 0x08, 0x9f, 0x47, 0x6b, 0x1f, 0x77, 0xf4, 0x97, 0x37, 0x14, 0xbc, 0xf8, 0x27, 0xed,
	0x85, 0x93, 0x6f, 0x1e, 0xb8, 0xf6, 0x86, 0x18, 0x67, 0xb1, 0xe0, 0xcb, 0x8d, 0x63, 0xfe, 0x69,
	0x21, 0xff, 0xd7, 0x12, 0x07, 0xf5, 0xca, 0x7c, 0xe1, 0xa3, 0x4b, 0xa2, 0xad, 0xee, 0x4f, 0x10,
	0x6d, 0x4a, 0x2f, 0xf3, 0x75, 0x9f, 0x7c, 0x68, 0x4f, 0xa8, 0x99, 0xda, 0x61, 0xca, 0x8b, 0x19,
	0xc5, 0x4a, 0x68, 0x31, 0x36, 0xe9, 0x8f, 0x39, 0x56, 0x12, 0xc7, 0x4e, 0xac, 0xb8, 0x88, 0x35,
	0xac, 0x65, 0xf3, 0x7e, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x36, 0x3d, 0x6a, 0xb9, 0xcf, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualMachineAgentClient is the client API for VirtualMachineAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualMachineAgentClient interface {
	CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error)
	RunCommand(ctx context.Context, in *VirtualMachineRunCommandRequest, opts ...grpc.CallOption) (*VirtualMachineRunCommandResponse, error)
	GetCommandResult(ctx context.Context, in *VirtualMachineCommandResultRequest, opts ...grpc.CallOption) (*VirtualMachineRunCommandResponse, error)
	UpdateAgent(ctx context.Context, in *UpdateAgentRequest, opts ...grpc.CallOption) (*UpdateAgentResponse, error)
}

type virtualMachineAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualMachineAgentClient(cc *grpc.ClientConn) VirtualMachineAgentClient {
	return &virtualMachineAgentClient{cc}
}

func (c *virtualMachineAgentClient) CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error) {
	out := new(common.NotificationResponse)
	err := c.cc.Invoke(ctx, "/moc.mocguestagent.compute.VirtualMachineAgent/CheckNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineAgentClient) RunCommand(ctx context.Context, in *VirtualMachineRunCommandRequest, opts ...grpc.CallOption) (*VirtualMachineRunCommandResponse, error) {
	out := new(VirtualMachineRunCommandResponse)
	err := c.cc.Invoke(ctx, "/moc.mocguestagent.compute.VirtualMachineAgent/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineAgentClient) GetCommandResult(ctx context.Context, in *VirtualMachineCommandResultRequest, opts ...grpc.CallOption) (*VirtualMachineRunCommandResponse, error) {
	out := new(VirtualMachineRunCommandResponse)
	err := c.cc.Invoke(ctx, "/moc.mocguestagent.compute.VirtualMachineAgent/GetCommandResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineAgentClient) UpdateAgent(ctx context.Context, in *UpdateAgentRequest, opts ...grpc.CallOption) (*UpdateAgentResponse, error) {
	out := new(UpdateAgentResponse)
	err := c.cc.Invoke(ctx, "/moc.mocguestagent.compute.VirtualMachineAgent/UpdateAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineAgentServer is the server API for VirtualMachineAgent service.
type VirtualMachineAgentServer interface {
	CheckNotification(context.Context, *empty.Empty) (*common.NotificationResponse, error)
	RunCommand(context.Context, *VirtualMachineRunCommandRequest) (*VirtualMachineRunCommandResponse, error)
	GetCommandResult(context.Context, *VirtualMachineCommandResultRequest) (*VirtualMachineRunCommandResponse, error)
	UpdateAgent(context.Context, *UpdateAgentRequest) (*UpdateAgentResponse, error)
}

// UnimplementedVirtualMachineAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineAgentServer struct {
}

func (*UnimplementedVirtualMachineAgentServer) CheckNotification(ctx context.Context, req *empty.Empty) (*common.NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNotification not implemented")
}
func (*UnimplementedVirtualMachineAgentServer) RunCommand(ctx context.Context, req *VirtualMachineRunCommandRequest) (*VirtualMachineRunCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}
func (*UnimplementedVirtualMachineAgentServer) GetCommandResult(ctx context.Context, req *VirtualMachineCommandResultRequest) (*VirtualMachineRunCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommandResult not implemented")
}
func (*UnimplementedVirtualMachineAgentServer) UpdateAgent(ctx context.Context, req *UpdateAgentRequest) (*UpdateAgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgent not implemented")
}

func RegisterVirtualMachineAgentServer(s *grpc.Server, srv VirtualMachineAgentServer) {
	s.RegisterService(&_VirtualMachineAgent_serviceDesc, srv)
}

func _VirtualMachineAgent_CheckNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServer).CheckNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.mocguestagent.compute.VirtualMachineAgent/CheckNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServer).CheckNotification(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineAgent_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineRunCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.mocguestagent.compute.VirtualMachineAgent/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServer).RunCommand(ctx, req.(*VirtualMachineRunCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineAgent_GetCommandResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualMachineCommandResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServer).GetCommandResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.mocguestagent.compute.VirtualMachineAgent/GetCommandResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServer).GetCommandResult(ctx, req.(*VirtualMachineCommandResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineAgent_UpdateAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServer).UpdateAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.mocguestagent.compute.VirtualMachineAgent/UpdateAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServer).UpdateAgent(ctx, req.(*UpdateAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMachineAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.mocguestagent.compute.VirtualMachineAgent",
	HandlerType: (*VirtualMachineAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckNotification",
			Handler:    _VirtualMachineAgent_CheckNotification_Handler,
		},
		{
			MethodName: "RunCommand",
			Handler:    _VirtualMachineAgent_RunCommand_Handler,
		},
		{
			MethodName: "GetCommandResult",
			Handler:    _VirtualMachineAgent_GetCommandResult_Handler,
		},
		{
			MethodName: "UpdateAgent",
			Handler:    _VirtualMachineAgent_UpdateAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtualmachine/moc_mocguestagent_virtualmachine.proto",
}
