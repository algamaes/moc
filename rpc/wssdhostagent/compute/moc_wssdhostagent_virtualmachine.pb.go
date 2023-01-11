// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_wssdhostagent_virtualmachine.proto

package compute

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type VirtualMachine struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Entity               *common.Entity `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags   `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VirtualMachine) Reset()         { *m = VirtualMachine{} }
func (m *VirtualMachine) String() string { return proto.CompactTextString(m) }
func (*VirtualMachine) ProtoMessage()    {}
func (*VirtualMachine) Descriptor() ([]byte, []int) {
	return fileDescriptor_382e78bf5e2a26d1, []int{0}
}

func (m *VirtualMachine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachine.Unmarshal(m, b)
}
func (m *VirtualMachine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachine.Marshal(b, m, deterministic)
}
func (m *VirtualMachine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachine.Merge(m, src)
}
func (m *VirtualMachine) XXX_Size() int {
	return xxx_messageInfo_VirtualMachine.Size(m)
}
func (m *VirtualMachine) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachine.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachine proto.InternalMessageInfo

func (m *VirtualMachine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualMachine) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualMachine) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *VirtualMachine) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type RegisterVirtualMachineRequest struct {
	VirtualMachine       *VirtualMachine `protobuf:"bytes,1,opt,name=VirtualMachine,proto3" json:"VirtualMachine,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegisterVirtualMachineRequest) Reset()         { *m = RegisterVirtualMachineRequest{} }
func (m *RegisterVirtualMachineRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterVirtualMachineRequest) ProtoMessage()    {}
func (*RegisterVirtualMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_382e78bf5e2a26d1, []int{1}
}

func (m *RegisterVirtualMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterVirtualMachineRequest.Unmarshal(m, b)
}
func (m *RegisterVirtualMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterVirtualMachineRequest.Marshal(b, m, deterministic)
}
func (m *RegisterVirtualMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterVirtualMachineRequest.Merge(m, src)
}
func (m *RegisterVirtualMachineRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterVirtualMachineRequest.Size(m)
}
func (m *RegisterVirtualMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterVirtualMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterVirtualMachineRequest proto.InternalMessageInfo

func (m *RegisterVirtualMachineRequest) GetVirtualMachine() *VirtualMachine {
	if m != nil {
		return m.VirtualMachine
	}
	return nil
}

type RegisterVirtualMachineResponse struct {
	InstanceView         *common.RegisterVirtualMachineInstanceView `protobuf:"bytes,1,opt,name=InstanceView,proto3" json:"InstanceView,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *RegisterVirtualMachineResponse) Reset()         { *m = RegisterVirtualMachineResponse{} }
func (m *RegisterVirtualMachineResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterVirtualMachineResponse) ProtoMessage()    {}
func (*RegisterVirtualMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_382e78bf5e2a26d1, []int{2}
}

func (m *RegisterVirtualMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterVirtualMachineResponse.Unmarshal(m, b)
}
func (m *RegisterVirtualMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterVirtualMachineResponse.Marshal(b, m, deterministic)
}
func (m *RegisterVirtualMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterVirtualMachineResponse.Merge(m, src)
}
func (m *RegisterVirtualMachineResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterVirtualMachineResponse.Size(m)
}
func (m *RegisterVirtualMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterVirtualMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterVirtualMachineResponse proto.InternalMessageInfo

func (m *RegisterVirtualMachineResponse) GetInstanceView() *common.RegisterVirtualMachineInstanceView {
	if m != nil {
		return m.InstanceView
	}
	return nil
}

type DeregisterVirtualMachineRequest struct {
	VirtualMachine       *VirtualMachine `protobuf:"bytes,1,opt,name=VirtualMachine,proto3" json:"VirtualMachine,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DeregisterVirtualMachineRequest) Reset()         { *m = DeregisterVirtualMachineRequest{} }
func (m *DeregisterVirtualMachineRequest) String() string { return proto.CompactTextString(m) }
func (*DeregisterVirtualMachineRequest) ProtoMessage()    {}
func (*DeregisterVirtualMachineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_382e78bf5e2a26d1, []int{3}
}

func (m *DeregisterVirtualMachineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeregisterVirtualMachineRequest.Unmarshal(m, b)
}
func (m *DeregisterVirtualMachineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeregisterVirtualMachineRequest.Marshal(b, m, deterministic)
}
func (m *DeregisterVirtualMachineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterVirtualMachineRequest.Merge(m, src)
}
func (m *DeregisterVirtualMachineRequest) XXX_Size() int {
	return xxx_messageInfo_DeregisterVirtualMachineRequest.Size(m)
}
func (m *DeregisterVirtualMachineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterVirtualMachineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterVirtualMachineRequest proto.InternalMessageInfo

func (m *DeregisterVirtualMachineRequest) GetVirtualMachine() *VirtualMachine {
	if m != nil {
		return m.VirtualMachine
	}
	return nil
}

type DeregisterVirtualMachineResponse struct {
	InstanceView         *common.DeregisterVirtualMachineInstanceView `protobuf:"bytes,1,opt,name=InstanceView,proto3" json:"InstanceView,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *DeregisterVirtualMachineResponse) Reset()         { *m = DeregisterVirtualMachineResponse{} }
func (m *DeregisterVirtualMachineResponse) String() string { return proto.CompactTextString(m) }
func (*DeregisterVirtualMachineResponse) ProtoMessage()    {}
func (*DeregisterVirtualMachineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_382e78bf5e2a26d1, []int{4}
}

func (m *DeregisterVirtualMachineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeregisterVirtualMachineResponse.Unmarshal(m, b)
}
func (m *DeregisterVirtualMachineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeregisterVirtualMachineResponse.Marshal(b, m, deterministic)
}
func (m *DeregisterVirtualMachineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterVirtualMachineResponse.Merge(m, src)
}
func (m *DeregisterVirtualMachineResponse) XXX_Size() int {
	return xxx_messageInfo_DeregisterVirtualMachineResponse.Size(m)
}
func (m *DeregisterVirtualMachineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterVirtualMachineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterVirtualMachineResponse proto.InternalMessageInfo

func (m *DeregisterVirtualMachineResponse) GetInstanceView() *common.DeregisterVirtualMachineInstanceView {
	if m != nil {
		return m.InstanceView
	}
	return nil
}

type VirtualMachineRunCommandRequest struct {
	VirtualMachine            *VirtualMachine                                  `protobuf:"bytes,1,opt,name=VirtualMachine,proto3" json:"VirtualMachine,omitempty"`
	Source                    *common.VirtualMachineRunCommandScriptSource     `protobuf:"bytes,2,opt,name=Source,proto3" json:"Source,omitempty"`
	RunCommandInputParameters []*common.VirtualMachineRunCommandInputParameter `protobuf:"bytes,3,rep,name=RunCommandInputParameters,proto3" json:"RunCommandInputParameters,omitempty"`
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
	return fileDescriptor_382e78bf5e2a26d1, []int{5}
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

func (m *VirtualMachineRunCommandRequest) GetVirtualMachine() *VirtualMachine {
	if m != nil {
		return m.VirtualMachine
	}
	return nil
}

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

type VirtualMachineRunCommandResponse struct {
	VirtualMachine       *VirtualMachine                              `protobuf:"bytes,1,opt,name=VirtualMachine,proto3" json:"VirtualMachine,omitempty"`
	InstanceView         *common.VirtualMachineRunCommandInstanceView `protobuf:"bytes,2,opt,name=InstanceView,proto3" json:"InstanceView,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *VirtualMachineRunCommandResponse) Reset()         { *m = VirtualMachineRunCommandResponse{} }
func (m *VirtualMachineRunCommandResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineRunCommandResponse) ProtoMessage()    {}
func (*VirtualMachineRunCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_382e78bf5e2a26d1, []int{6}
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

func (m *VirtualMachineRunCommandResponse) GetVirtualMachine() *VirtualMachine {
	if m != nil {
		return m.VirtualMachine
	}
	return nil
}

func (m *VirtualMachineRunCommandResponse) GetInstanceView() *common.VirtualMachineRunCommandInstanceView {
	if m != nil {
		return m.InstanceView
	}
	return nil
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
	return fileDescriptor_382e78bf5e2a26d1, []int{7}
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

func init() {
	proto.RegisterType((*VirtualMachine)(nil), "moc.wssdhostagent.compute.VirtualMachine")
	proto.RegisterType((*RegisterVirtualMachineRequest)(nil), "moc.wssdhostagent.compute.RegisterVirtualMachineRequest")
	proto.RegisterType((*RegisterVirtualMachineResponse)(nil), "moc.wssdhostagent.compute.RegisterVirtualMachineResponse")
	proto.RegisterType((*DeregisterVirtualMachineRequest)(nil), "moc.wssdhostagent.compute.DeregisterVirtualMachineRequest")
	proto.RegisterType((*DeregisterVirtualMachineResponse)(nil), "moc.wssdhostagent.compute.DeregisterVirtualMachineResponse")
	proto.RegisterType((*VirtualMachineRunCommandRequest)(nil), "moc.wssdhostagent.compute.VirtualMachineRunCommandRequest")
	proto.RegisterType((*VirtualMachineRunCommandResponse)(nil), "moc.wssdhostagent.compute.VirtualMachineRunCommandResponse")
	proto.RegisterType((*VirtualMachineCommandResultRequest)(nil), "moc.wssdhostagent.compute.VirtualMachineCommandResultRequest")
}

func init() {
	proto.RegisterFile("moc_wssdhostagent_virtualmachine.proto", fileDescriptor_382e78bf5e2a26d1)
}

var fileDescriptor_382e78bf5e2a26d1 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0x13, 0x3f,
	0x10, 0xef, 0x26, 0xf9, 0x47, 0xca, 0xe4, 0x4f, 0x0f, 0x46, 0x82, 0x6d, 0xa4, 0xb6, 0xab, 0x45,
	0x82, 0x00, 0xd2, 0x46, 0x0a, 0x1c, 0xf8, 0x38, 0xa5, 0x14, 0xa4, 0x08, 0x55, 0x14, 0x17, 0x7a,
	0xe0, 0x52, 0xb9, 0x8e, 0x49, 0x2c, 0xd5, 0xf6, 0xd6, 0xf6, 0x12, 0x78, 0x02, 0xb8, 0x71, 0xe2,
	0x51, 0x78, 0x03, 0x2e, 0xbc, 0x0e, 0x4f, 0x80, 0xd6, 0xbb, 0x25, 0xd9, 0x55, 0x76, 0x83, 0x2a,
	0x95, 0x53, 0x92, 0xf1, 0xfc, 0x3e, 0x66, 0x3c, 0xe3, 0xc0, 0x6d, 0xa1, 0xe8, 0xc9, 0xdc, 0x98,
	0xc9, 0x4c, 0x19, 0x4b, 0xa6, 0x4c, 0xda, 0x93, 0x0f, 0x5c, 0xdb, 0x84, 0x9c, 0x09, 0x42, 0x67,
	0x5c, 0xb2, 0x28, 0xd6, 0xca, 0x2a, 0xb4, 0x25, 0x14, 0x8d, 0x0a, 0x79, 0x11, 0x55, 0x22, 0x4e,
	0x2c, 0xeb, 0xdd, 0x4c, 0x29, 0xa8, 0x12, 0x42, 0xc9, 0xfc, 0x23, 0xc3, 0xf4, 0x76, 0x8a, 0x07,
	0x69, 0xf2, 0xf2, 0x79, 0xf8, 0x11, 0x36, 0x8f, 0x33, 0xad, 0x83, 0x4c, 0x0b, 0x21, 0x68, 0x49,
	0x22, 0x98, 0xef, 0x05, 0x5e, 0xbf, 0x83, 0xdd, 0x77, 0xb4, 0x09, 0x0d, 0x3e, 0xf1, 0x1b, 0x2e,
	0xd2, 0xe0, 0x13, 0x74, 0x0b, 0xda, 0x4c, 0x5a, 0x6e, 0x3f, 0xf9, 0xcd, 0xc0, 0xeb, 0x77, 0x87,
	0xdd, 0x28, 0xb5, 0xf6, 0xdc, 0x85, 0x70, 0x7e, 0x84, 0xb6, 0xa1, 0x65, 0xc9, 0xd4, 0xf8, 0x2d,
	0x97, 0xd2, 0x71, 0x29, 0x6f, 0xc8, 0xd4, 0x60, 0x17, 0x0e, 0x35, 0x6c, 0x63, 0x36, 0xe5, 0xc6,
	0x32, 0x5d, 0x74, 0x80, 0xd9, 0x79, 0xc2, 0x8c, 0x45, 0xaf, 0xcb, 0xd6, 0x9c, 0xa5, 0xee, 0xf0,
	0x6e, 0x54, 0xd9, 0x87, 0xa8, 0xc4, 0x54, 0x22, 0x08, 0x05, 0xec, 0x54, 0x69, 0x9a, 0x58, 0x49,
	0xc3, 0xd0, 0x4b, 0xf8, 0x7f, 0x2c, 0x8d, 0x25, 0x92, 0xb2, 0x63, 0xce, 0xe6, 0xb9, 0xe4, 0x1d,
	0x27, 0xb9, 0x1a, 0xba, 0x9c, 0x8e, 0x0b, 0xe0, 0xd0, 0xc2, 0xee, 0x3e, 0xd3, 0xff, 0xba, 0xc8,
	0x73, 0x08, 0xaa, 0x55, 0xf3, 0x32, 0x0f, 0x56, 0x96, 0x99, 0x89, 0x56, 0x81, 0x6b, 0x0a, 0xfd,
	0xd5, 0x80, 0xdd, 0x92, 0x52, 0x22, 0x9f, 0x29, 0x21, 0x88, 0x9c, 0x5c, 0x5d, 0xa5, 0x68, 0x04,
	0xed, 0x23, 0x95, 0x68, 0xca, 0xdc, 0x68, 0x5e, 0x50, 0x55, 0x19, 0x39, 0xa2, 0x9a, 0xc7, 0x36,
	0x03, 0xe0, 0x1c, 0x88, 0x38, 0x6c, 0x2d, 0x32, 0xc6, 0x32, 0x4e, 0xec, 0x21, 0xd1, 0x44, 0x30,
	0xcb, 0xb4, 0xf1, 0x9b, 0x41, 0xb3, 0xdf, 0x1d, 0xde, 0xaf, 0x65, 0x2d, 0x62, 0x70, 0x35, 0x1b,
	0x0a, 0xa1, 0x83, 0x13, 0x39, 0x32, 0x6f, 0x0d, 0xd3, 0x6e, 0x29, 0x3a, 0x7b, 0xad, 0x2f, 0xdf,
	0x7d, 0x0f, 0x2f, 0xc2, 0xe8, 0x1e, 0x5c, 0x73, 0x3f, 0x0e, 0x89, 0x31, 0x73, 0xa5, 0x27, 0xfe,
	0x7f, 0x4b, 0x79, 0xc5, 0xa3, 0xf0, 0x87, 0x07, 0x41, 0x75, 0xd3, 0xf3, 0x8b, 0xbe, 0x82, 0xae,
	0x97, 0x67, 0xe7, 0x6f, 0x7a, 0x5f, 0x33, 0x3b, 0x2f, 0x20, 0x2c, 0xa2, 0x16, 0x25, 0x24, 0x67,
	0xf6, 0x62, 0x7a, 0x02, 0xe8, 0xbe, 0x8a, 0x99, 0x26, 0x96, 0x2b, 0x39, 0xde, 0xcf, 0x1f, 0xa7,
	0xe5, 0xd0, 0xf0, 0x67, 0x13, 0xae, 0x17, 0x89, 0x46, 0x69, 0x55, 0xe8, 0xab, 0x07, 0x37, 0x56,
	0x6f, 0x2e, 0x7a, 0x54, 0xd3, 0x84, 0xda, 0xb7, 0xa9, 0xf7, 0xf8, 0x12, 0xc8, 0xec, 0x46, 0xc2,
	0x0d, 0xf4, 0xcd, 0x03, 0xbf, 0x6a, 0xc9, 0xd0, 0x93, 0x1a, 0xe6, 0x35, 0x8f, 0x49, 0xef, 0xe9,
	0xa5, 0xb0, 0x7f, 0x7c, 0x7d, 0xf6, 0x00, 0x16, 0x57, 0x56, 0xeb, 0x64, 0xcd, 0xb2, 0xd7, 0x3a,
	0x59, 0x37, 0xb3, 0xe1, 0xc6, 0xde, 0xc3, 0x77, 0xc3, 0x29, 0xb7, 0xb3, 0xe4, 0x34, 0x05, 0x0d,
	0x04, 0xa7, 0x5a, 0x19, 0xf5, 0xde, 0x0e, 0x84, 0xa2, 0x03, 0x1d, 0xd3, 0x41, 0x81, 0x78, 0x90,
	0x13, 0x9f, 0xb6, 0xdd, 0x5f, 0xda, 0x83, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0x65, 0x1d,
	0xe3, 0x50, 0x07, 0x00, 0x00,
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
	RegisterVirtualMachine(ctx context.Context, in *RegisterVirtualMachineRequest, opts ...grpc.CallOption) (*RegisterVirtualMachineResponse, error)
	DeregisterVirtualMachine(ctx context.Context, in *DeregisterVirtualMachineRequest, opts ...grpc.CallOption) (*DeregisterVirtualMachineResponse, error)
	RunCommand(ctx context.Context, in *VirtualMachineRunCommandRequest, opts ...grpc.CallOption) (*VirtualMachineRunCommandResponse, error)
}

type virtualMachineAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualMachineAgentClient(cc *grpc.ClientConn) VirtualMachineAgentClient {
	return &virtualMachineAgentClient{cc}
}

func (c *virtualMachineAgentClient) RegisterVirtualMachine(ctx context.Context, in *RegisterVirtualMachineRequest, opts ...grpc.CallOption) (*RegisterVirtualMachineResponse, error) {
	out := new(RegisterVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/moc.wssdhostagent.compute.VirtualMachineAgent/RegisterVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineAgentClient) DeregisterVirtualMachine(ctx context.Context, in *DeregisterVirtualMachineRequest, opts ...grpc.CallOption) (*DeregisterVirtualMachineResponse, error) {
	out := new(DeregisterVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/moc.wssdhostagent.compute.VirtualMachineAgent/DeregisterVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachineAgentClient) RunCommand(ctx context.Context, in *VirtualMachineRunCommandRequest, opts ...grpc.CallOption) (*VirtualMachineRunCommandResponse, error) {
	out := new(VirtualMachineRunCommandResponse)
	err := c.cc.Invoke(ctx, "/moc.wssdhostagent.compute.VirtualMachineAgent/RunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualMachineAgentServer is the server API for VirtualMachineAgent service.
type VirtualMachineAgentServer interface {
	RegisterVirtualMachine(context.Context, *RegisterVirtualMachineRequest) (*RegisterVirtualMachineResponse, error)
	DeregisterVirtualMachine(context.Context, *DeregisterVirtualMachineRequest) (*DeregisterVirtualMachineResponse, error)
	RunCommand(context.Context, *VirtualMachineRunCommandRequest) (*VirtualMachineRunCommandResponse, error)
}

// UnimplementedVirtualMachineAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualMachineAgentServer struct {
}

func (*UnimplementedVirtualMachineAgentServer) RegisterVirtualMachine(ctx context.Context, req *RegisterVirtualMachineRequest) (*RegisterVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterVirtualMachine not implemented")
}
func (*UnimplementedVirtualMachineAgentServer) DeregisterVirtualMachine(ctx context.Context, req *DeregisterVirtualMachineRequest) (*DeregisterVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeregisterVirtualMachine not implemented")
}
func (*UnimplementedVirtualMachineAgentServer) RunCommand(ctx context.Context, req *VirtualMachineRunCommandRequest) (*VirtualMachineRunCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCommand not implemented")
}

func RegisterVirtualMachineAgentServer(s *grpc.Server, srv VirtualMachineAgentServer) {
	s.RegisterService(&_VirtualMachineAgent_serviceDesc, srv)
}

func _VirtualMachineAgent_RegisterVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServer).RegisterVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.wssdhostagent.compute.VirtualMachineAgent/RegisterVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServer).RegisterVirtualMachine(ctx, req.(*RegisterVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachineAgent_DeregisterVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachineAgentServer).DeregisterVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.wssdhostagent.compute.VirtualMachineAgent/DeregisterVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServer).DeregisterVirtualMachine(ctx, req.(*DeregisterVirtualMachineRequest))
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
		FullMethod: "/moc.wssdhostagent.compute.VirtualMachineAgent/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachineAgentServer).RunCommand(ctx, req.(*VirtualMachineRunCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualMachineAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.wssdhostagent.compute.VirtualMachineAgent",
	HandlerType: (*VirtualMachineAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterVirtualMachine",
			Handler:    _VirtualMachineAgent_RegisterVirtualMachine_Handler,
		},
		{
			MethodName: "DeregisterVirtualMachine",
			Handler:    _VirtualMachineAgent_DeregisterVirtualMachine_Handler,
		},
		{
			MethodName: "RunCommand",
			Handler:    _VirtualMachineAgent_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_wssdhostagent_virtualmachine.proto",
}