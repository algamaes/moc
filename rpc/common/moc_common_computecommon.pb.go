// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_common_computecommon.proto

package common

import (
	fmt "fmt"
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

type UserType int32

const (
	UserType_ROOT UserType = 0
	UserType_USER UserType = 1
)

var UserType_name = map[int32]string{
	0: "ROOT",
	1: "USER",
}

var UserType_value = map[string]int32{
	"ROOT": 0,
	"USER": 1,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}

func (UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{0}
}

type ProcessorType int32

const (
	ProcessorType_None    ProcessorType = 0
	ProcessorType_Intel   ProcessorType = 1
	ProcessorType_Intel64 ProcessorType = 2
	ProcessorType_AMD     ProcessorType = 3
	ProcessorType_AMD64   ProcessorType = 4
	ProcessorType_ARM     ProcessorType = 5
	ProcessorType_ARM64   ProcessorType = 6
)

var ProcessorType_name = map[int32]string{
	0: "None",
	1: "Intel",
	2: "Intel64",
	3: "AMD",
	4: "AMD64",
	5: "ARM",
	6: "ARM64",
}

var ProcessorType_value = map[string]int32{
	"None":    0,
	"Intel":   1,
	"Intel64": 2,
	"AMD":     3,
	"AMD64":   4,
	"ARM":     5,
	"ARM64":   6,
}

func (x ProcessorType) String() string {
	return proto.EnumName(ProcessorType_name, int32(x))
}

func (ProcessorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{1}
}

type OperatingSystemBootstrapEngine int32

const (
	OperatingSystemBootstrapEngine_CLOUD_INIT           OperatingSystemBootstrapEngine = 0
	OperatingSystemBootstrapEngine_WINDOWS_ANSWER_FILES OperatingSystemBootstrapEngine = 1
)

var OperatingSystemBootstrapEngine_name = map[int32]string{
	0: "CLOUD_INIT",
	1: "WINDOWS_ANSWER_FILES",
}

var OperatingSystemBootstrapEngine_value = map[string]int32{
	"CLOUD_INIT":           0,
	"WINDOWS_ANSWER_FILES": 1,
}

func (x OperatingSystemBootstrapEngine) String() string {
	return proto.EnumName(OperatingSystemBootstrapEngine_name, int32(x))
}

func (OperatingSystemBootstrapEngine) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{2}
}

type OperatingSystemType int32

const (
	OperatingSystemType_WINDOWS OperatingSystemType = 0
	OperatingSystemType_LINUX   OperatingSystemType = 1
)

var OperatingSystemType_name = map[int32]string{
	0: "WINDOWS",
	1: "LINUX",
}

var OperatingSystemType_value = map[string]int32{
	"WINDOWS": 0,
	"LINUX":   1,
}

func (x OperatingSystemType) String() string {
	return proto.EnumName(OperatingSystemType_name, int32(x))
}

func (OperatingSystemType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{3}
}

type VirtualMachineSizeType int32

const (
	VirtualMachineSizeType_Default             VirtualMachineSizeType = 0
	VirtualMachineSizeType_Standard_A2_v2      VirtualMachineSizeType = 2
	VirtualMachineSizeType_Standard_A4_v2      VirtualMachineSizeType = 3
	VirtualMachineSizeType_Standard_D2s_v3     VirtualMachineSizeType = 4
	VirtualMachineSizeType_Standard_D4s_v3     VirtualMachineSizeType = 5
	VirtualMachineSizeType_Standard_D8s_v3     VirtualMachineSizeType = 6
	VirtualMachineSizeType_Standard_D16s_v3    VirtualMachineSizeType = 7
	VirtualMachineSizeType_Standard_D32s_v3    VirtualMachineSizeType = 8
	VirtualMachineSizeType_Standard_DS2_v2     VirtualMachineSizeType = 9
	VirtualMachineSizeType_Standard_DS3_v2     VirtualMachineSizeType = 10
	VirtualMachineSizeType_Standard_DS4_v2     VirtualMachineSizeType = 11
	VirtualMachineSizeType_Standard_DS5_v2     VirtualMachineSizeType = 12
	VirtualMachineSizeType_Standard_DS13_v2    VirtualMachineSizeType = 13
	VirtualMachineSizeType_Standard_K8S_v1     VirtualMachineSizeType = 14
	VirtualMachineSizeType_Standard_K8S2_v1    VirtualMachineSizeType = 15
	VirtualMachineSizeType_Standard_K8S3_v1    VirtualMachineSizeType = 16
	VirtualMachineSizeType_Standard_K8S4_v1    VirtualMachineSizeType = 17
	VirtualMachineSizeType_Standard_NK6        VirtualMachineSizeType = 18
	VirtualMachineSizeType_Standard_NK12       VirtualMachineSizeType = 19
	VirtualMachineSizeType_Standard_NV6        VirtualMachineSizeType = 20
	VirtualMachineSizeType_Standard_NV12       VirtualMachineSizeType = 21
	VirtualMachineSizeType_Standard_K8S5_v1    VirtualMachineSizeType = 22
	VirtualMachineSizeType_Standard_DS2_v2_HPN VirtualMachineSizeType = 23
	VirtualMachineSizeType_Standard_DS3_v2_HPN VirtualMachineSizeType = 24
	VirtualMachineSizeType_Standard_DS4_v2_HPN VirtualMachineSizeType = 25
	VirtualMachineSizeType_Standard_F2s_HPN    VirtualMachineSizeType = 26
	VirtualMachineSizeType_Standard_F4s_HPN    VirtualMachineSizeType = 27
	VirtualMachineSizeType_Standard_F8s_HPN    VirtualMachineSizeType = 28
	VirtualMachineSizeType_Standard_F16s_HPN   VirtualMachineSizeType = 29
	VirtualMachineSizeType_Custom_NK           VirtualMachineSizeType = 96
	VirtualMachineSizeType_Custom_Gpupv        VirtualMachineSizeType = 97
	VirtualMachineSizeType_Custom              VirtualMachineSizeType = 98
	VirtualMachineSizeType_Unsupported         VirtualMachineSizeType = 99
)

var VirtualMachineSizeType_name = map[int32]string{
	0:  "Default",
	2:  "Standard_A2_v2",
	3:  "Standard_A4_v2",
	4:  "Standard_D2s_v3",
	5:  "Standard_D4s_v3",
	6:  "Standard_D8s_v3",
	7:  "Standard_D16s_v3",
	8:  "Standard_D32s_v3",
	9:  "Standard_DS2_v2",
	10: "Standard_DS3_v2",
	11: "Standard_DS4_v2",
	12: "Standard_DS5_v2",
	13: "Standard_DS13_v2",
	14: "Standard_K8S_v1",
	15: "Standard_K8S2_v1",
	16: "Standard_K8S3_v1",
	17: "Standard_K8S4_v1",
	18: "Standard_NK6",
	19: "Standard_NK12",
	20: "Standard_NV6",
	21: "Standard_NV12",
	22: "Standard_K8S5_v1",
	23: "Standard_DS2_v2_HPN",
	24: "Standard_DS3_v2_HPN",
	25: "Standard_DS4_v2_HPN",
	26: "Standard_F2s_HPN",
	27: "Standard_F4s_HPN",
	28: "Standard_F8s_HPN",
	29: "Standard_F16s_HPN",
	96: "Custom_NK",
	97: "Custom_Gpupv",
	98: "Custom",
	99: "Unsupported",
}

var VirtualMachineSizeType_value = map[string]int32{
	"Default":             0,
	"Standard_A2_v2":      2,
	"Standard_A4_v2":      3,
	"Standard_D2s_v3":     4,
	"Standard_D4s_v3":     5,
	"Standard_D8s_v3":     6,
	"Standard_D16s_v3":    7,
	"Standard_D32s_v3":    8,
	"Standard_DS2_v2":     9,
	"Standard_DS3_v2":     10,
	"Standard_DS4_v2":     11,
	"Standard_DS5_v2":     12,
	"Standard_DS13_v2":    13,
	"Standard_K8S_v1":     14,
	"Standard_K8S2_v1":    15,
	"Standard_K8S3_v1":    16,
	"Standard_K8S4_v1":    17,
	"Standard_NK6":        18,
	"Standard_NK12":       19,
	"Standard_NV6":        20,
	"Standard_NV12":       21,
	"Standard_K8S5_v1":    22,
	"Standard_DS2_v2_HPN": 23,
	"Standard_DS3_v2_HPN": 24,
	"Standard_DS4_v2_HPN": 25,
	"Standard_F2s_HPN":    26,
	"Standard_F4s_HPN":    27,
	"Standard_F8s_HPN":    28,
	"Standard_F16s_HPN":   29,
	"Custom_NK":           96,
	"Custom_Gpupv":        97,
	"Custom":              98,
	"Unsupported":         99,
}

func (x VirtualMachineSizeType) String() string {
	return proto.EnumName(VirtualMachineSizeType_name, int32(x))
}

func (VirtualMachineSizeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{4}
}

type WinRMProtocolType int32

const (
	WinRMProtocolType_HTTP  WinRMProtocolType = 0
	WinRMProtocolType_HTTPS WinRMProtocolType = 1
)

var WinRMProtocolType_name = map[int32]string{
	0: "HTTP",
	1: "HTTPS",
}

var WinRMProtocolType_value = map[string]int32{
	"HTTP":  0,
	"HTTPS": 1,
}

func (x WinRMProtocolType) String() string {
	return proto.EnumName(WinRMProtocolType_name, int32(x))
}

func (WinRMProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{5}
}

type PowerState int32

const (
	PowerState_Unknown  PowerState = 0
	PowerState_Running  PowerState = 1
	PowerState_Off      PowerState = 2
	PowerState_Paused   PowerState = 3
	PowerState_Critical PowerState = 4
)

var PowerState_name = map[int32]string{
	0: "Unknown",
	1: "Running",
	2: "Off",
	3: "Paused",
	4: "Critical",
}

var PowerState_value = map[string]int32{
	"Unknown":  0,
	"Running":  1,
	"Off":      2,
	"Paused":   3,
	"Critical": 4,
}

func (x PowerState) String() string {
	return proto.EnumName(PowerState_name, int32(x))
}

func (PowerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{6}
}

type VirtualMachineOperation int32

const (
	VirtualMachineOperation_START VirtualMachineOperation = 0
	VirtualMachineOperation_STOP  VirtualMachineOperation = 1
	VirtualMachineOperation_RESET VirtualMachineOperation = 2
)

var VirtualMachineOperation_name = map[int32]string{
	0: "START",
	1: "STOP",
	2: "RESET",
}

var VirtualMachineOperation_value = map[string]int32{
	"START": 0,
	"STOP":  1,
	"RESET": 2,
}

func (x VirtualMachineOperation) String() string {
	return proto.EnumName(VirtualMachineOperation_name, int32(x))
}

func (VirtualMachineOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{7}
}

type VirtualMachineRunCommandExecutionState int32

const (
	VirtualMachineRunCommandExecutionState_ExecutionState_UNKNOWN   VirtualMachineRunCommandExecutionState = 0
	VirtualMachineRunCommandExecutionState_ExecutionState_FAILED    VirtualMachineRunCommandExecutionState = 1
	VirtualMachineRunCommandExecutionState_ExecutionState_SUCCEEDED VirtualMachineRunCommandExecutionState = 2
)

var VirtualMachineRunCommandExecutionState_name = map[int32]string{
	0: "ExecutionState_UNKNOWN",
	1: "ExecutionState_FAILED",
	2: "ExecutionState_SUCCEEDED",
}

var VirtualMachineRunCommandExecutionState_value = map[string]int32{
	"ExecutionState_UNKNOWN":   0,
	"ExecutionState_FAILED":    1,
	"ExecutionState_SUCCEEDED": 2,
}

func (x VirtualMachineRunCommandExecutionState) String() string {
	return proto.EnumName(VirtualMachineRunCommandExecutionState_name, int32(x))
}

func (VirtualMachineRunCommandExecutionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{8}
}

type Architecture int32

const (
	Architecture_x86                Architecture = 0
	Architecture_MIPS               Architecture = 1
	Architecture_Alpha              Architecture = 2
	Architecture_PowerPC            Architecture = 3
	Architecture_ARM_Architecture   Architecture = 5
	Architecture_ia64               Architecture = 6
	Architecture_x64                Architecture = 9
	Architecture_ARM64_Architecture Architecture = 12
)

var Architecture_name = map[int32]string{
	0:  "x86",
	1:  "MIPS",
	2:  "Alpha",
	3:  "PowerPC",
	5:  "ARM_Architecture",
	6:  "ia64",
	9:  "x64",
	12: "ARM64_Architecture",
}

var Architecture_value = map[string]int32{
	"x86":                0,
	"MIPS":               1,
	"Alpha":              2,
	"PowerPC":            3,
	"ARM_Architecture":   5,
	"ia64":               6,
	"x64":                9,
	"ARM64_Architecture": 12,
}

func (x Architecture) String() string {
	return proto.EnumName(Architecture_name, int32(x))
}

func (Architecture) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{9}
}

type WinRMListener struct {
	Protocol             WinRMProtocolType `protobuf:"varint,1,opt,name=Protocol,proto3,enum=moc.WinRMProtocolType" json:"Protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WinRMListener) Reset()         { *m = WinRMListener{} }
func (m *WinRMListener) String() string { return proto.CompactTextString(m) }
func (*WinRMListener) ProtoMessage()    {}
func (*WinRMListener) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{0}
}

func (m *WinRMListener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WinRMListener.Unmarshal(m, b)
}
func (m *WinRMListener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WinRMListener.Marshal(b, m, deterministic)
}
func (m *WinRMListener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WinRMListener.Merge(m, src)
}
func (m *WinRMListener) XXX_Size() int {
	return xxx_messageInfo_WinRMListener.Size(m)
}
func (m *WinRMListener) XXX_DiscardUnknown() {
	xxx_messageInfo_WinRMListener.DiscardUnknown(m)
}

var xxx_messageInfo_WinRMListener proto.InternalMessageInfo

func (m *WinRMListener) GetProtocol() WinRMProtocolType {
	if m != nil {
		return m.Protocol
	}
	return WinRMProtocolType_HTTP
}

type WinRMConfiguration struct {
	Listeners            []*WinRMListener `protobuf:"bytes,1,rep,name=Listeners,proto3" json:"Listeners,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *WinRMConfiguration) Reset()         { *m = WinRMConfiguration{} }
func (m *WinRMConfiguration) String() string { return proto.CompactTextString(m) }
func (*WinRMConfiguration) ProtoMessage()    {}
func (*WinRMConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{1}
}

func (m *WinRMConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WinRMConfiguration.Unmarshal(m, b)
}
func (m *WinRMConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WinRMConfiguration.Marshal(b, m, deterministic)
}
func (m *WinRMConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WinRMConfiguration.Merge(m, src)
}
func (m *WinRMConfiguration) XXX_Size() int {
	return xxx_messageInfo_WinRMConfiguration.Size(m)
}
func (m *WinRMConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_WinRMConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_WinRMConfiguration proto.InternalMessageInfo

func (m *WinRMConfiguration) GetListeners() []*WinRMListener {
	if m != nil {
		return m.Listeners
	}
	return nil
}

type VirtualMachineCustomSize struct {
	CpuCount             int32    `protobuf:"varint,1,opt,name=cpuCount,proto3" json:"cpuCount,omitempty"`
	MemoryMB             int32    `protobuf:"varint,2,opt,name=memoryMB,proto3" json:"memoryMB,omitempty"`
	GpuCount             int32    `protobuf:"varint,3,opt,name=gpuCount,proto3" json:"gpuCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMachineCustomSize) Reset()         { *m = VirtualMachineCustomSize{} }
func (m *VirtualMachineCustomSize) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineCustomSize) ProtoMessage()    {}
func (*VirtualMachineCustomSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{2}
}

func (m *VirtualMachineCustomSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineCustomSize.Unmarshal(m, b)
}
func (m *VirtualMachineCustomSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineCustomSize.Marshal(b, m, deterministic)
}
func (m *VirtualMachineCustomSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineCustomSize.Merge(m, src)
}
func (m *VirtualMachineCustomSize) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineCustomSize.Size(m)
}
func (m *VirtualMachineCustomSize) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineCustomSize.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineCustomSize proto.InternalMessageInfo

func (m *VirtualMachineCustomSize) GetCpuCount() int32 {
	if m != nil {
		return m.CpuCount
	}
	return 0
}

func (m *VirtualMachineCustomSize) GetMemoryMB() int32 {
	if m != nil {
		return m.MemoryMB
	}
	return 0
}

func (m *VirtualMachineCustomSize) GetGpuCount() int32 {
	if m != nil {
		return m.GpuCount
	}
	return 0
}

type DynamicMemoryConfiguration struct {
	MaximumMemoryMB      uint64   `protobuf:"varint,1,opt,name=maximumMemoryMB,proto3" json:"maximumMemoryMB,omitempty"`
	MinimumMemoryMB      uint64   `protobuf:"varint,2,opt,name=minimumMemoryMB,proto3" json:"minimumMemoryMB,omitempty"`
	TargetMemoryBuffer   uint32   `protobuf:"varint,3,opt,name=targetMemoryBuffer,proto3" json:"targetMemoryBuffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DynamicMemoryConfiguration) Reset()         { *m = DynamicMemoryConfiguration{} }
func (m *DynamicMemoryConfiguration) String() string { return proto.CompactTextString(m) }
func (*DynamicMemoryConfiguration) ProtoMessage()    {}
func (*DynamicMemoryConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{3}
}

func (m *DynamicMemoryConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicMemoryConfiguration.Unmarshal(m, b)
}
func (m *DynamicMemoryConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicMemoryConfiguration.Marshal(b, m, deterministic)
}
func (m *DynamicMemoryConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicMemoryConfiguration.Merge(m, src)
}
func (m *DynamicMemoryConfiguration) XXX_Size() int {
	return xxx_messageInfo_DynamicMemoryConfiguration.Size(m)
}
func (m *DynamicMemoryConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicMemoryConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicMemoryConfiguration proto.InternalMessageInfo

func (m *DynamicMemoryConfiguration) GetMaximumMemoryMB() uint64 {
	if m != nil {
		return m.MaximumMemoryMB
	}
	return 0
}

func (m *DynamicMemoryConfiguration) GetMinimumMemoryMB() uint64 {
	if m != nil {
		return m.MinimumMemoryMB
	}
	return 0
}

func (m *DynamicMemoryConfiguration) GetTargetMemoryBuffer() uint32 {
	if m != nil {
		return m.TargetMemoryBuffer
	}
	return 0
}

type RegisterVirtualMachineInstanceView struct {
	Output               string   `protobuf:"bytes,1,opt,name=Output,proto3" json:"Output,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterVirtualMachineInstanceView) Reset()         { *m = RegisterVirtualMachineInstanceView{} }
func (m *RegisterVirtualMachineInstanceView) String() string { return proto.CompactTextString(m) }
func (*RegisterVirtualMachineInstanceView) ProtoMessage()    {}
func (*RegisterVirtualMachineInstanceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{4}
}

func (m *RegisterVirtualMachineInstanceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterVirtualMachineInstanceView.Unmarshal(m, b)
}
func (m *RegisterVirtualMachineInstanceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterVirtualMachineInstanceView.Marshal(b, m, deterministic)
}
func (m *RegisterVirtualMachineInstanceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterVirtualMachineInstanceView.Merge(m, src)
}
func (m *RegisterVirtualMachineInstanceView) XXX_Size() int {
	return xxx_messageInfo_RegisterVirtualMachineInstanceView.Size(m)
}
func (m *RegisterVirtualMachineInstanceView) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterVirtualMachineInstanceView.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterVirtualMachineInstanceView proto.InternalMessageInfo

func (m *RegisterVirtualMachineInstanceView) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *RegisterVirtualMachineInstanceView) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type DeregisterVirtualMachineInstanceView struct {
	Output               string   `protobuf:"bytes,1,opt,name=Output,proto3" json:"Output,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeregisterVirtualMachineInstanceView) Reset()         { *m = DeregisterVirtualMachineInstanceView{} }
func (m *DeregisterVirtualMachineInstanceView) String() string { return proto.CompactTextString(m) }
func (*DeregisterVirtualMachineInstanceView) ProtoMessage()    {}
func (*DeregisterVirtualMachineInstanceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{5}
}

func (m *DeregisterVirtualMachineInstanceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeregisterVirtualMachineInstanceView.Unmarshal(m, b)
}
func (m *DeregisterVirtualMachineInstanceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeregisterVirtualMachineInstanceView.Marshal(b, m, deterministic)
}
func (m *DeregisterVirtualMachineInstanceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterVirtualMachineInstanceView.Merge(m, src)
}
func (m *DeregisterVirtualMachineInstanceView) XXX_Size() int {
	return xxx_messageInfo_DeregisterVirtualMachineInstanceView.Size(m)
}
func (m *DeregisterVirtualMachineInstanceView) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterVirtualMachineInstanceView.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterVirtualMachineInstanceView proto.InternalMessageInfo

func (m *DeregisterVirtualMachineInstanceView) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *DeregisterVirtualMachineInstanceView) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VirtualMachineRunCommandInputParameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMachineRunCommandInputParameter) Reset() {
	*m = VirtualMachineRunCommandInputParameter{}
}
func (m *VirtualMachineRunCommandInputParameter) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineRunCommandInputParameter) ProtoMessage()    {}
func (*VirtualMachineRunCommandInputParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{6}
}

func (m *VirtualMachineRunCommandInputParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineRunCommandInputParameter.Unmarshal(m, b)
}
func (m *VirtualMachineRunCommandInputParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineRunCommandInputParameter.Marshal(b, m, deterministic)
}
func (m *VirtualMachineRunCommandInputParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineRunCommandInputParameter.Merge(m, src)
}
func (m *VirtualMachineRunCommandInputParameter) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineRunCommandInputParameter.Size(m)
}
func (m *VirtualMachineRunCommandInputParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineRunCommandInputParameter.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineRunCommandInputParameter proto.InternalMessageInfo

func (m *VirtualMachineRunCommandInputParameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualMachineRunCommandInputParameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type VirtualMachineRunCommandScriptSource struct {
	Script               string   `protobuf:"bytes,1,opt,name=Script,proto3" json:"Script,omitempty"`
	ScriptURI            string   `protobuf:"bytes,2,opt,name=ScriptURI,proto3" json:"ScriptURI,omitempty"`
	CommandID            string   `protobuf:"bytes,3,opt,name=CommandID,proto3" json:"CommandID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMachineRunCommandScriptSource) Reset()         { *m = VirtualMachineRunCommandScriptSource{} }
func (m *VirtualMachineRunCommandScriptSource) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineRunCommandScriptSource) ProtoMessage()    {}
func (*VirtualMachineRunCommandScriptSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{7}
}

func (m *VirtualMachineRunCommandScriptSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineRunCommandScriptSource.Unmarshal(m, b)
}
func (m *VirtualMachineRunCommandScriptSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineRunCommandScriptSource.Marshal(b, m, deterministic)
}
func (m *VirtualMachineRunCommandScriptSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineRunCommandScriptSource.Merge(m, src)
}
func (m *VirtualMachineRunCommandScriptSource) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineRunCommandScriptSource.Size(m)
}
func (m *VirtualMachineRunCommandScriptSource) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineRunCommandScriptSource.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineRunCommandScriptSource proto.InternalMessageInfo

func (m *VirtualMachineRunCommandScriptSource) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *VirtualMachineRunCommandScriptSource) GetScriptURI() string {
	if m != nil {
		return m.ScriptURI
	}
	return ""
}

func (m *VirtualMachineRunCommandScriptSource) GetCommandID() string {
	if m != nil {
		return m.CommandID
	}
	return ""
}

type VirtualMachineRunCommandInstanceView struct {
	ExecutionState       VirtualMachineRunCommandExecutionState `protobuf:"varint,1,opt,name=ExecutionState,proto3,enum=moc.VirtualMachineRunCommandExecutionState" json:"ExecutionState,omitempty"`
	ExitCode             int32                                  `protobuf:"varint,2,opt,name=ExitCode,proto3" json:"ExitCode,omitempty"`
	Output               string                                 `protobuf:"bytes,3,opt,name=Output,proto3" json:"Output,omitempty"`
	Error                string                                 `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *VirtualMachineRunCommandInstanceView) Reset()         { *m = VirtualMachineRunCommandInstanceView{} }
func (m *VirtualMachineRunCommandInstanceView) String() string { return proto.CompactTextString(m) }
func (*VirtualMachineRunCommandInstanceView) ProtoMessage()    {}
func (*VirtualMachineRunCommandInstanceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1a061f6c82445b, []int{8}
}

func (m *VirtualMachineRunCommandInstanceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMachineRunCommandInstanceView.Unmarshal(m, b)
}
func (m *VirtualMachineRunCommandInstanceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMachineRunCommandInstanceView.Marshal(b, m, deterministic)
}
func (m *VirtualMachineRunCommandInstanceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMachineRunCommandInstanceView.Merge(m, src)
}
func (m *VirtualMachineRunCommandInstanceView) XXX_Size() int {
	return xxx_messageInfo_VirtualMachineRunCommandInstanceView.Size(m)
}
func (m *VirtualMachineRunCommandInstanceView) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMachineRunCommandInstanceView.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMachineRunCommandInstanceView proto.InternalMessageInfo

func (m *VirtualMachineRunCommandInstanceView) GetExecutionState() VirtualMachineRunCommandExecutionState {
	if m != nil {
		return m.ExecutionState
	}
	return VirtualMachineRunCommandExecutionState_ExecutionState_UNKNOWN
}

func (m *VirtualMachineRunCommandInstanceView) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *VirtualMachineRunCommandInstanceView) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *VirtualMachineRunCommandInstanceView) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("moc.UserType", UserType_name, UserType_value)
	proto.RegisterEnum("moc.ProcessorType", ProcessorType_name, ProcessorType_value)
	proto.RegisterEnum("moc.OperatingSystemBootstrapEngine", OperatingSystemBootstrapEngine_name, OperatingSystemBootstrapEngine_value)
	proto.RegisterEnum("moc.OperatingSystemType", OperatingSystemType_name, OperatingSystemType_value)
	proto.RegisterEnum("moc.VirtualMachineSizeType", VirtualMachineSizeType_name, VirtualMachineSizeType_value)
	proto.RegisterEnum("moc.WinRMProtocolType", WinRMProtocolType_name, WinRMProtocolType_value)
	proto.RegisterEnum("moc.PowerState", PowerState_name, PowerState_value)
	proto.RegisterEnum("moc.VirtualMachineOperation", VirtualMachineOperation_name, VirtualMachineOperation_value)
	proto.RegisterEnum("moc.VirtualMachineRunCommandExecutionState", VirtualMachineRunCommandExecutionState_name, VirtualMachineRunCommandExecutionState_value)
	proto.RegisterEnum("moc.Architecture", Architecture_name, Architecture_value)
	proto.RegisterType((*WinRMListener)(nil), "moc.WinRMListener")
	proto.RegisterType((*WinRMConfiguration)(nil), "moc.WinRMConfiguration")
	proto.RegisterType((*VirtualMachineCustomSize)(nil), "moc.VirtualMachineCustomSize")
	proto.RegisterType((*DynamicMemoryConfiguration)(nil), "moc.DynamicMemoryConfiguration")
	proto.RegisterType((*RegisterVirtualMachineInstanceView)(nil), "moc.RegisterVirtualMachineInstanceView")
	proto.RegisterType((*DeregisterVirtualMachineInstanceView)(nil), "moc.DeregisterVirtualMachineInstanceView")
	proto.RegisterType((*VirtualMachineRunCommandInputParameter)(nil), "moc.VirtualMachineRunCommandInputParameter")
	proto.RegisterType((*VirtualMachineRunCommandScriptSource)(nil), "moc.VirtualMachineRunCommandScriptSource")
	proto.RegisterType((*VirtualMachineRunCommandInstanceView)(nil), "moc.VirtualMachineRunCommandInstanceView")
}

func init() { proto.RegisterFile("moc_common_computecommon.proto", fileDescriptor_7d1a061f6c82445b) }

var fileDescriptor_7d1a061f6c82445b = []byte{
	// 1138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x5b, 0x53, 0x1b, 0x37,
	0x14, 0xc7, 0xf1, 0x0d, 0xf0, 0xe1, 0x26, 0x04, 0x21, 0x4e, 0x9a, 0x66, 0x32, 0x6e, 0xda, 0x61,
	0xdc, 0x29, 0x14, 0xe3, 0x78, 0xe8, 0xa3, 0xb1, 0x97, 0xc6, 0x05, 0x5f, 0x46, 0x6b, 0x43, 0xdb,
	0x97, 0xed, 0x66, 0x91, 0xcd, 0x4e, 0xbd, 0xd2, 0x8e, 0x56, 0x4b, 0x20, 0x33, 0xfd, 0x2e, 0xfd,
	0x1e, 0x7d, 0xee, 0xf7, 0xea, 0x48, 0xeb, 0xc5, 0x68, 0x13, 0xde, 0xfa, 0x84, 0xce, 0xef, 0x7f,
	0xf4, 0xd7, 0xd9, 0xa3, 0x0b, 0x86, 0xd7, 0x01, 0xf7, 0x1c, 0x8f, 0x07, 0x01, 0x67, 0xea, 0x4f,
	0x18, 0x4b, 0x9a, 0x44, 0x07, 0xa1, 0xe0, 0x92, 0xe3, 0x42, 0xc0, 0xbd, 0x6a, 0x1b, 0x36, 0xae,
	0x7c, 0x46, 0x7a, 0x17, 0x7e, 0x24, 0x29, 0xa3, 0x02, 0xd7, 0x61, 0x75, 0xa8, 0x64, 0x8f, 0xcf,
	0x2a, 0xb9, 0x37, 0xb9, 0xfd, 0xcd, 0xfa, 0xde, 0x41, 0xc0, 0xbd, 0x03, 0x9d, 0x95, 0x2a, 0xa3,
	0xfb, 0x90, 0x92, 0x87, 0xbc, 0xea, 0x19, 0x60, 0x2d, 0xb7, 0x39, 0x9b, 0xf8, 0xd3, 0x58, 0xb8,
	0xd2, 0xe7, 0x0c, 0xff, 0x08, 0xe5, 0xd4, 0x35, 0xaa, 0xe4, 0xde, 0x14, 0xf6, 0xd7, 0xea, 0x78,
	0x61, 0x95, 0x4a, 0x64, 0x91, 0x54, 0x65, 0x50, 0xb9, 0xf4, 0x85, 0x8c, 0xdd, 0x59, 0xcf, 0xf5,
	0x6e, 0x7c, 0x46, 0xdb, 0x71, 0x24, 0x79, 0x60, 0xfb, 0x9f, 0x28, 0x7e, 0x09, 0xab, 0x5e, 0x18,
	0xb7, 0x79, 0xcc, 0xa4, 0xae, 0xab, 0x44, 0x1e, 0x62, 0xa5, 0x05, 0x34, 0xe0, 0xe2, 0xbe, 0x77,
	0x5a, 0xc9, 0x27, 0x5a, 0x1a, 0x2b, 0x6d, 0x9a, 0xce, 0x2b, 0x24, 0x5a, 0x1a, 0x57, 0xff, 0xce,
	0xc1, 0xcb, 0xce, 0x3d, 0x73, 0x03, 0xdf, 0xeb, 0xe9, 0x7c, 0xf3, 0x03, 0xf6, 0x61, 0x2b, 0x70,
	0xef, 0xfc, 0x20, 0x0e, 0x7a, 0xa9, 0xbb, 0x5a, 0xb9, 0x48, 0xb2, 0x58, 0x67, 0xfa, 0xcc, 0xc8,
	0xcc, 0xcf, 0x33, 0x4d, 0x8c, 0x0f, 0x00, 0x4b, 0x57, 0x4c, 0xa9, 0x4c, 0xc8, 0x69, 0x3c, 0x99,
	0x50, 0xa1, 0x0b, 0xdb, 0x20, 0x5f, 0x50, 0xaa, 0x04, 0xaa, 0x84, 0x4e, 0x55, 0x87, 0x84, 0xd9,
	0x9a, 0x2e, 0x8b, 0xa4, 0xcb, 0x3c, 0x7a, 0xe9, 0xd3, 0x8f, 0x78, 0x0f, 0x96, 0x07, 0xb1, 0x0c,
	0xe3, 0xa4, 0x35, 0x65, 0x32, 0x8f, 0xf0, 0x2e, 0x94, 0x2c, 0x21, 0xb8, 0xd0, 0xd5, 0x94, 0x49,
	0x12, 0x54, 0x47, 0xf0, 0xb6, 0x43, 0xc5, 0xff, 0xed, 0x4a, 0xe0, 0x3b, 0xd3, 0x8b, 0xc4, 0xac,
	0xcd, 0x83, 0xc0, 0x65, 0xd7, 0x5d, 0x16, 0xc6, 0x72, 0xe8, 0x0a, 0x37, 0xa0, 0x92, 0x0a, 0x8c,
	0xa1, 0xd8, 0x77, 0x03, 0x3a, 0x77, 0xd5, 0x63, 0xe5, 0x79, 0xe9, 0xce, 0x62, 0x9a, 0x7a, 0xea,
	0xa0, 0xfa, 0x09, 0xde, 0x3e, 0xe5, 0x69, 0x7b, 0xc2, 0x0f, 0xa5, 0xcd, 0x63, 0xe1, 0x51, 0x55,
	0x69, 0x12, 0xa7, 0x95, 0x26, 0x11, 0x7e, 0x05, 0xe5, 0x64, 0x34, 0x26, 0xdd, 0xb9, 0xf3, 0x02,
	0x28, 0x35, 0x2d, 0xb0, 0xa3, 0xb7, 0xa0, 0x4c, 0x16, 0xa0, 0xfa, 0x6f, 0xee, 0xe9, 0xc5, 0x8d,
	0x36, 0xd9, 0xb0, 0x69, 0xdd, 0x51, 0x2f, 0x56, 0x67, 0xc6, 0x96, 0xae, 0xa4, 0xf3, 0x7b, 0xf3,
	0xbd, 0x3e, 0xec, 0x4f, 0x59, 0x98, 0x53, 0x48, 0xc6, 0x42, 0x1d, 0x5b, 0xeb, 0xce, 0x97, 0x6d,
	0x7e, 0x4d, 0xd3, 0x23, 0x9d, 0xc6, 0x8f, 0xf6, 0xa5, 0xf0, 0xe5, 0x7d, 0x29, 0x3e, 0xda, 0x97,
	0xda, 0x6b, 0x58, 0x1d, 0x47, 0x54, 0xa8, 0x2b, 0x8b, 0x57, 0xa1, 0x48, 0x06, 0x83, 0x11, 0x5a,
	0x52, 0xa3, 0xb1, 0x6d, 0x11, 0x94, 0xab, 0xfd, 0x06, 0x1b, 0x43, 0xc1, 0x3d, 0x1a, 0x45, 0xfc,
	0x21, 0xa9, 0xcf, 0x19, 0x45, 0x4b, 0xb8, 0x0c, 0xa5, 0x2e, 0x93, 0x74, 0x86, 0x72, 0x78, 0x0d,
	0x56, 0xf4, 0xb0, 0xd9, 0x40, 0x79, 0xbc, 0x02, 0x85, 0x56, 0xaf, 0x83, 0x0a, 0x2a, 0xa1, 0xd5,
	0xeb, 0x34, 0x1b, 0xa8, 0xa8, 0x19, 0xe9, 0xa1, 0x92, 0x66, 0xa4, 0xd7, 0x6c, 0xa0, 0xe5, 0xda,
	0x2f, 0xf0, 0x7a, 0x10, 0x52, 0x75, 0x9b, 0xd8, 0xd4, 0xbe, 0x8f, 0x24, 0x0d, 0x4e, 0x39, 0x97,
	0x91, 0x14, 0x6e, 0x68, 0xb1, 0xa9, 0xcf, 0x28, 0xde, 0x04, 0x68, 0x5f, 0x0c, 0xc6, 0x1d, 0xa7,
	0xdb, 0xef, 0xaa, 0xb2, 0x2a, 0xb0, 0x7b, 0xd5, 0xed, 0x77, 0x06, 0x57, 0xb6, 0xd3, 0xea, 0xdb,
	0x57, 0x16, 0x71, 0xce, 0xba, 0x17, 0x96, 0x8d, 0x72, 0xb5, 0x1f, 0x60, 0x27, 0xe3, 0xa5, 0x8b,
	0x5d, 0x83, 0x95, 0xf9, 0x84, 0xa4, 0xde, 0x8b, 0x6e, 0x7f, 0xfc, 0x2b, 0xca, 0xd5, 0xfe, 0x29,
	0xc1, 0x9e, 0xd9, 0x7a, 0xf5, 0x8a, 0xa4, 0x53, 0x3a, 0x74, 0xe2, 0xc6, 0x33, 0x89, 0x96, 0x30,
	0x86, 0x4d, 0x5b, 0xba, 0xec, 0xda, 0x15, 0xd7, 0x4e, 0xab, 0xee, 0xdc, 0xd6, 0x51, 0xde, 0x64,
	0x0d, 0xc5, 0x0a, 0x78, 0x07, 0xb6, 0x1e, 0x58, 0xa7, 0x1e, 0x39, 0xb7, 0xc7, 0xa8, 0x68, 0xc2,
	0x86, 0x86, 0x25, 0x13, 0x9e, 0x68, 0xb8, 0x8c, 0x77, 0x01, 0x2d, 0xe0, 0x51, 0x53, 0xd3, 0x15,
	0x93, 0x1e, 0x27, 0xae, 0xab, 0xa6, 0x81, 0xad, 0x6b, 0x2a, 0x67, 0xe0, 0xb1, 0x82, 0x90, 0x81,
	0xba, 0xd2, 0xb5, 0x0c, 0x7c, 0xa7, 0xe0, 0xba, 0xb9, 0x92, 0x7d, 0xa4, 0xe7, 0x6f, 0x18, 0xa9,
	0xe7, 0x27, 0xb6, 0x73, 0x7b, 0x84, 0x36, 0x8d, 0xd4, 0xf3, 0x13, 0xb5, 0xfe, 0x11, 0xda, 0xca,
	0xd2, 0x63, 0x45, 0x51, 0x96, 0x36, 0x14, 0xdd, 0xc6, 0x08, 0xd6, 0x1f, 0x68, 0xff, 0xbc, 0x89,
	0x30, 0xde, 0x86, 0x8d, 0x47, 0xe4, 0xa8, 0x8e, 0x76, 0xcc, 0xa4, 0xcb, 0x26, 0xda, 0x35, 0x93,
	0x2e, 0x8f, 0xea, 0xe8, 0x59, 0xd6, 0xff, 0x9d, 0xf2, 0xdf, 0xc3, 0xcf, 0x61, 0x27, 0xd3, 0x20,
	0xe7, 0xfd, 0xb0, 0x8f, 0x9e, 0x67, 0x84, 0xe3, 0x54, 0xa8, 0x64, 0x84, 0x46, 0x2a, 0xbc, 0x30,
	0x16, 0x38, 0xab, 0x47, 0x9a, 0xbe, 0x34, 0x69, 0x23, 0xa1, 0x5f, 0x99, 0xf4, 0x24, 0xa1, 0xaf,
	0xf0, 0x33, 0xd8, 0x5e, 0x50, 0xb5, 0xb3, 0x0a, 0x7f, 0x8d, 0x37, 0xa0, 0x9c, 0xfc, 0xf3, 0x72,
	0xfa, 0xe7, 0xe8, 0x0f, 0xf5, 0xb5, 0xf3, 0xf0, 0xe7, 0x30, 0x0e, 0x6f, 0x91, 0x8b, 0x01, 0x96,
	0x13, 0x82, 0x3e, 0xe0, 0x2d, 0x58, 0x1b, 0xb3, 0x28, 0x0e, 0x43, 0x2e, 0x24, 0xbd, 0x46, 0x5e,
	0x6d, 0x1f, 0xb6, 0x3f, 0xfb, 0x7f, 0xab, 0xee, 0xe5, 0xfb, 0xd1, 0x68, 0x98, 0x9c, 0x73, 0x35,
	0x52, 0xd7, 0xa2, 0x0b, 0x30, 0xe4, 0x1f, 0xa9, 0x48, 0x5e, 0x8d, 0x35, 0x58, 0x19, 0xb3, 0x3f,
	0x19, 0xff, 0xc8, 0xd0, 0x92, 0x0a, 0x48, 0xcc, 0x98, 0xcf, 0xa6, 0x28, 0xa7, 0xae, 0xe7, 0x60,
	0x32, 0x41, 0x79, 0xb5, 0xee, 0xd0, 0x8d, 0x23, 0x7a, 0x8d, 0x0a, 0x78, 0x1d, 0x56, 0xdb, 0xc2,
	0x97, 0xbe, 0xe7, 0xce, 0x50, 0xb1, 0xf6, 0x13, 0x3c, 0x37, 0x6f, 0xcc, 0xfc, 0xbe, 0x71, 0xa6,
	0x16, 0xb4, 0x47, 0x2d, 0x32, 0x7f, 0x38, 0xec, 0xd1, 0x60, 0x88, 0x72, 0x0a, 0x12, 0xcb, 0xb6,
	0x46, 0x28, 0x5f, 0xfb, 0xeb, 0xe9, 0xb7, 0xff, 0xb3, 0x77, 0x6d, 0xcf, 0x24, 0xce, 0xb8, 0x7f,
	0xde, 0x1f, 0x5c, 0xf5, 0xd1, 0x12, 0x7e, 0x01, 0xcf, 0x32, 0xda, 0x59, 0xab, 0x7b, 0x61, 0x75,
	0x50, 0x0e, 0xbf, 0x82, 0x4a, 0x46, 0xb2, 0xc7, 0xed, 0xb6, 0x65, 0x75, 0xac, 0x0e, 0xca, 0xd7,
	0x3e, 0xc1, 0x7a, 0x4b, 0x78, 0x37, 0xbe, 0xa4, 0x9e, 0x8c, 0x05, 0x55, 0x1f, 0x7b, 0x77, 0xd2,
	0x4c, 0x8a, 0xed, 0x75, 0x55, 0x9f, 0xf4, 0xab, 0x34, 0x0b, 0x6f, 0x5c, 0x94, 0x57, 0x7d, 0xd1,
	0x2d, 0x1b, 0xb6, 0x51, 0x41, 0x6d, 0x6a, 0x8b, 0xf4, 0x9c, 0xc7, 0xd3, 0x51, 0x49, 0xcd, 0xf3,
	0x5d, 0xf5, 0x84, 0x69, 0xab, 0x66, 0x03, 0x95, 0xf1, 0x1e, 0x60, 0xfd, 0xac, 0x99, 0xa9, 0xeb,
	0xa7, 0xdf, 0xfe, 0xfe, 0xcd, 0xd4, 0x97, 0x37, 0xf1, 0x87, 0x03, 0x8f, 0x07, 0x87, 0x81, 0xef,
	0x09, 0x1e, 0xf1, 0x89, 0x3c, 0x0c, 0xb8, 0x77, 0x28, 0x42, 0xef, 0x30, 0xf9, 0xc9, 0xf5, 0x61,
	0x59, 0xff, 0xe6, 0x3a, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xab, 0x5c, 0xbc, 0x22, 0x95, 0x09,
	0x00, 0x00,
}
