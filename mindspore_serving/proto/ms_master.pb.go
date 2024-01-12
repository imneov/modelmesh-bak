// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mindspore_serving/proto/ms_master.proto

package mindspore_serving_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ServableRegSpec struct {
	Name                 string                        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	VersionNumber        uint64                        `protobuf:"varint,2,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
	BatchSize            uint64                        `protobuf:"varint,4,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	Methods              []*ServableRegSpec_MethodInfo `protobuf:"bytes,5,rep,name=methods,proto3" json:"methods,omitempty"`
	ModelInfos           *ModelInfos                   `protobuf:"bytes,6,opt,name=model_infos,json=modelInfos,proto3" json:"model_infos,omitempty"`
	OwnDevice            bool                          `protobuf:"varint,7,opt,name=own_device,json=ownDevice,proto3" json:"own_device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ServableRegSpec) Reset()         { *m = ServableRegSpec{} }
func (m *ServableRegSpec) String() string { return proto.CompactTextString(m) }
func (*ServableRegSpec) ProtoMessage()    {}
func (*ServableRegSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{0}
}
func (m *ServableRegSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServableRegSpec.Unmarshal(m, b)
}
func (m *ServableRegSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServableRegSpec.Marshal(b, m, deterministic)
}
func (m *ServableRegSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServableRegSpec.Merge(m, src)
}
func (m *ServableRegSpec) XXX_Size() int {
	return xxx_messageInfo_ServableRegSpec.Size(m)
}
func (m *ServableRegSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServableRegSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServableRegSpec proto.InternalMessageInfo

func (m *ServableRegSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServableRegSpec) GetVersionNumber() uint64 {
	if m != nil {
		return m.VersionNumber
	}
	return 0
}

func (m *ServableRegSpec) GetBatchSize() uint64 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *ServableRegSpec) GetMethods() []*ServableRegSpec_MethodInfo {
	if m != nil {
		return m.Methods
	}
	return nil
}

func (m *ServableRegSpec) GetModelInfos() *ModelInfos {
	if m != nil {
		return m.ModelInfos
	}
	return nil
}

func (m *ServableRegSpec) GetOwnDevice() bool {
	if m != nil {
		return m.OwnDevice
	}
	return false
}

type ServableRegSpec_MethodInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InputNames           []string `protobuf:"bytes,2,rep,name=input_names,json=inputNames,proto3" json:"input_names,omitempty"`
	OnlyModelStage       bool     `protobuf:"varint,3,opt,name=only_model_stage,json=onlyModelStage,proto3" json:"only_model_stage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServableRegSpec_MethodInfo) Reset()         { *m = ServableRegSpec_MethodInfo{} }
func (m *ServableRegSpec_MethodInfo) String() string { return proto.CompactTextString(m) }
func (*ServableRegSpec_MethodInfo) ProtoMessage()    {}
func (*ServableRegSpec_MethodInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{0, 0}
}
func (m *ServableRegSpec_MethodInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServableRegSpec_MethodInfo.Unmarshal(m, b)
}
func (m *ServableRegSpec_MethodInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServableRegSpec_MethodInfo.Marshal(b, m, deterministic)
}
func (m *ServableRegSpec_MethodInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServableRegSpec_MethodInfo.Merge(m, src)
}
func (m *ServableRegSpec_MethodInfo) XXX_Size() int {
	return xxx_messageInfo_ServableRegSpec_MethodInfo.Size(m)
}
func (m *ServableRegSpec_MethodInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ServableRegSpec_MethodInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ServableRegSpec_MethodInfo proto.InternalMessageInfo

func (m *ServableRegSpec_MethodInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServableRegSpec_MethodInfo) GetInputNames() []string {
	if m != nil {
		return m.InputNames
	}
	return nil
}

func (m *ServableRegSpec_MethodInfo) GetOnlyModelStage() bool {
	if m != nil {
		return m.OnlyModelStage
	}
	return false
}

type WorkerRegSpec struct {
	WorkerPid            uint64           `protobuf:"varint,1,opt,name=worker_pid,json=workerPid,proto3" json:"worker_pid,omitempty"`
	Address              string           `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	ServableSpec         *ServableRegSpec `protobuf:"bytes,4,opt,name=servable_spec,json=servableSpec,proto3" json:"servable_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *WorkerRegSpec) Reset()         { *m = WorkerRegSpec{} }
func (m *WorkerRegSpec) String() string { return proto.CompactTextString(m) }
func (*WorkerRegSpec) ProtoMessage()    {}
func (*WorkerRegSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{1}
}
func (m *WorkerRegSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerRegSpec.Unmarshal(m, b)
}
func (m *WorkerRegSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerRegSpec.Marshal(b, m, deterministic)
}
func (m *WorkerRegSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerRegSpec.Merge(m, src)
}
func (m *WorkerRegSpec) XXX_Size() int {
	return xxx_messageInfo_WorkerRegSpec.Size(m)
}
func (m *WorkerRegSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerRegSpec.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerRegSpec proto.InternalMessageInfo

func (m *WorkerRegSpec) GetWorkerPid() uint64 {
	if m != nil {
		return m.WorkerPid
	}
	return 0
}

func (m *WorkerRegSpec) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WorkerRegSpec) GetServableSpec() *ServableRegSpec {
	if m != nil {
		return m.ServableSpec
	}
	return nil
}

type RegisterRequest struct {
	WorkerSpec           *WorkerRegSpec `protobuf:"bytes,1,opt,name=worker_spec,json=workerSpec,proto3" json:"worker_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{2}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetWorkerSpec() *WorkerRegSpec {
	if m != nil {
		return m.WorkerSpec
	}
	return nil
}

type RegisterReply struct {
	ErrorMsg             *ErrorMsg `protobuf:"bytes,1,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegisterReply) Reset()         { *m = RegisterReply{} }
func (m *RegisterReply) String() string { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()    {}
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{3}
}
func (m *RegisterReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReply.Unmarshal(m, b)
}
func (m *RegisterReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReply.Marshal(b, m, deterministic)
}
func (m *RegisterReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReply.Merge(m, src)
}
func (m *RegisterReply) XXX_Size() int {
	return xxx_messageInfo_RegisterReply.Size(m)
}
func (m *RegisterReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReply proto.InternalMessageInfo

func (m *RegisterReply) GetErrorMsg() *ErrorMsg {
	if m != nil {
		return m.ErrorMsg
	}
	return nil
}

type ExitRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExitRequest) Reset()         { *m = ExitRequest{} }
func (m *ExitRequest) String() string { return proto.CompactTextString(m) }
func (*ExitRequest) ProtoMessage()    {}
func (*ExitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{4}
}
func (m *ExitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitRequest.Unmarshal(m, b)
}
func (m *ExitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitRequest.Marshal(b, m, deterministic)
}
func (m *ExitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitRequest.Merge(m, src)
}
func (m *ExitRequest) XXX_Size() int {
	return xxx_messageInfo_ExitRequest.Size(m)
}
func (m *ExitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExitRequest proto.InternalMessageInfo

func (m *ExitRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ExitReply struct {
	ErrorMsg             *ErrorMsg `protobuf:"bytes,1,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExitReply) Reset()         { *m = ExitReply{} }
func (m *ExitReply) String() string { return proto.CompactTextString(m) }
func (*ExitReply) ProtoMessage()    {}
func (*ExitReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{5}
}
func (m *ExitReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExitReply.Unmarshal(m, b)
}
func (m *ExitReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExitReply.Marshal(b, m, deterministic)
}
func (m *ExitReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExitReply.Merge(m, src)
}
func (m *ExitReply) XXX_Size() int {
	return xxx_messageInfo_ExitReply.Size(m)
}
func (m *ExitReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ExitReply.DiscardUnknown(m)
}

var xxx_messageInfo_ExitReply proto.InternalMessageInfo

func (m *ExitReply) GetErrorMsg() *ErrorMsg {
	if m != nil {
		return m.ErrorMsg
	}
	return nil
}

type NotifyFailedRequest struct {
	WorkerPid            uint64   `protobuf:"varint,1,opt,name=worker_pid,json=workerPid,proto3" json:"worker_pid,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,2,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyFailedRequest) Reset()         { *m = NotifyFailedRequest{} }
func (m *NotifyFailedRequest) String() string { return proto.CompactTextString(m) }
func (*NotifyFailedRequest) ProtoMessage()    {}
func (*NotifyFailedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{6}
}
func (m *NotifyFailedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyFailedRequest.Unmarshal(m, b)
}
func (m *NotifyFailedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyFailedRequest.Marshal(b, m, deterministic)
}
func (m *NotifyFailedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyFailedRequest.Merge(m, src)
}
func (m *NotifyFailedRequest) XXX_Size() int {
	return xxx_messageInfo_NotifyFailedRequest.Size(m)
}
func (m *NotifyFailedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyFailedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyFailedRequest proto.InternalMessageInfo

func (m *NotifyFailedRequest) GetWorkerPid() uint64 {
	if m != nil {
		return m.WorkerPid
	}
	return 0
}

func (m *NotifyFailedRequest) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

type NotifyFailedReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyFailedReply) Reset()         { *m = NotifyFailedReply{} }
func (m *NotifyFailedReply) String() string { return proto.CompactTextString(m) }
func (*NotifyFailedReply) ProtoMessage()    {}
func (*NotifyFailedReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{7}
}
func (m *NotifyFailedReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyFailedReply.Unmarshal(m, b)
}
func (m *NotifyFailedReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyFailedReply.Marshal(b, m, deterministic)
}
func (m *NotifyFailedReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyFailedReply.Merge(m, src)
}
func (m *NotifyFailedReply) XXX_Size() int {
	return xxx_messageInfo_NotifyFailedReply.Size(m)
}
func (m *NotifyFailedReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyFailedReply.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyFailedReply proto.InternalMessageInfo

type GetModelInfoRequest struct {
	ServableName         string   `protobuf:"bytes,1,opt,name=servable_name,json=servableName,proto3" json:"servable_name,omitempty"`
	VersionNumber        uint32   `protobuf:"varint,2,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetModelInfoRequest) Reset()         { *m = GetModelInfoRequest{} }
func (m *GetModelInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetModelInfoRequest) ProtoMessage()    {}
func (*GetModelInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{8}
}
func (m *GetModelInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelInfoRequest.Unmarshal(m, b)
}
func (m *GetModelInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetModelInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelInfoRequest.Merge(m, src)
}
func (m *GetModelInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetModelInfoRequest.Size(m)
}
func (m *GetModelInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelInfoRequest proto.InternalMessageInfo

func (m *GetModelInfoRequest) GetServableName() string {
	if m != nil {
		return m.ServableName
	}
	return ""
}

func (m *GetModelInfoRequest) GetVersionNumber() uint32 {
	if m != nil {
		return m.VersionNumber
	}
	return 0
}

type ModelSubGraphInfo struct {
	Inputs               []*TensorInfo `protobuf:"bytes,3,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []*TensorInfo `protobuf:"bytes,4,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ModelSubGraphInfo) Reset()         { *m = ModelSubGraphInfo{} }
func (m *ModelSubGraphInfo) String() string { return proto.CompactTextString(m) }
func (*ModelSubGraphInfo) ProtoMessage()    {}
func (*ModelSubGraphInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{9}
}
func (m *ModelSubGraphInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelSubGraphInfo.Unmarshal(m, b)
}
func (m *ModelSubGraphInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelSubGraphInfo.Marshal(b, m, deterministic)
}
func (m *ModelSubGraphInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelSubGraphInfo.Merge(m, src)
}
func (m *ModelSubGraphInfo) XXX_Size() int {
	return xxx_messageInfo_ModelSubGraphInfo.Size(m)
}
func (m *ModelSubGraphInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelSubGraphInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ModelSubGraphInfo proto.InternalMessageInfo

func (m *ModelSubGraphInfo) GetInputs() []*TensorInfo {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *ModelSubGraphInfo) GetOutputs() []*TensorInfo {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type ModelInfo struct {
	BatchSize            uint64               `protobuf:"varint,2,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	SubgraphInfos        []*ModelSubGraphInfo `protobuf:"bytes,1,rep,name=subgraph_infos,json=subgraphInfos,proto3" json:"subgraph_infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ModelInfo) Reset()         { *m = ModelInfo{} }
func (m *ModelInfo) String() string { return proto.CompactTextString(m) }
func (*ModelInfo) ProtoMessage()    {}
func (*ModelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{10}
}
func (m *ModelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelInfo.Unmarshal(m, b)
}
func (m *ModelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelInfo.Marshal(b, m, deterministic)
}
func (m *ModelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelInfo.Merge(m, src)
}
func (m *ModelInfo) XXX_Size() int {
	return xxx_messageInfo_ModelInfo.Size(m)
}
func (m *ModelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ModelInfo proto.InternalMessageInfo

func (m *ModelInfo) GetBatchSize() uint64 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *ModelInfo) GetSubgraphInfos() []*ModelSubGraphInfo {
	if m != nil {
		return m.SubgraphInfos
	}
	return nil
}

type ModelInfos struct {
	ModelInfos           map[string]*ModelInfo `protobuf:"bytes,1,rep,name=model_infos,json=modelInfos,proto3" json:"model_infos,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ModelInfos) Reset()         { *m = ModelInfos{} }
func (m *ModelInfos) String() string { return proto.CompactTextString(m) }
func (*ModelInfos) ProtoMessage()    {}
func (*ModelInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{11}
}
func (m *ModelInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelInfos.Unmarshal(m, b)
}
func (m *ModelInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelInfos.Marshal(b, m, deterministic)
}
func (m *ModelInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelInfos.Merge(m, src)
}
func (m *ModelInfos) XXX_Size() int {
	return xxx_messageInfo_ModelInfos.Size(m)
}
func (m *ModelInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelInfos.DiscardUnknown(m)
}

var xxx_messageInfo_ModelInfos proto.InternalMessageInfo

func (m *ModelInfos) GetModelInfos() map[string]*ModelInfo {
	if m != nil {
		return m.ModelInfos
	}
	return nil
}

type GetModelInfoReply struct {
	ServableName         string      `protobuf:"bytes,1,opt,name=servable_name,json=servableName,proto3" json:"servable_name,omitempty"`
	VersionNumber        uint32      `protobuf:"varint,2,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
	ModelInfos           *ModelInfos `protobuf:"bytes,3,opt,name=model_infos,json=modelInfos,proto3" json:"model_infos,omitempty"`
	ErrorMsg             *ErrorMsg   `protobuf:"bytes,4,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetModelInfoReply) Reset()         { *m = GetModelInfoReply{} }
func (m *GetModelInfoReply) String() string { return proto.CompactTextString(m) }
func (*GetModelInfoReply) ProtoMessage()    {}
func (*GetModelInfoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d87d022b788c635, []int{12}
}
func (m *GetModelInfoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelInfoReply.Unmarshal(m, b)
}
func (m *GetModelInfoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelInfoReply.Marshal(b, m, deterministic)
}
func (m *GetModelInfoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelInfoReply.Merge(m, src)
}
func (m *GetModelInfoReply) XXX_Size() int {
	return xxx_messageInfo_GetModelInfoReply.Size(m)
}
func (m *GetModelInfoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelInfoReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelInfoReply proto.InternalMessageInfo

func (m *GetModelInfoReply) GetServableName() string {
	if m != nil {
		return m.ServableName
	}
	return ""
}

func (m *GetModelInfoReply) GetVersionNumber() uint32 {
	if m != nil {
		return m.VersionNumber
	}
	return 0
}

func (m *GetModelInfoReply) GetModelInfos() *ModelInfos {
	if m != nil {
		return m.ModelInfos
	}
	return nil
}

func (m *GetModelInfoReply) GetErrorMsg() *ErrorMsg {
	if m != nil {
		return m.ErrorMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*ServableRegSpec)(nil), "mindspore.serving.proto.ServableRegSpec")
	proto.RegisterType((*ServableRegSpec_MethodInfo)(nil), "mindspore.serving.proto.ServableRegSpec.MethodInfo")
	proto.RegisterType((*WorkerRegSpec)(nil), "mindspore.serving.proto.WorkerRegSpec")
	proto.RegisterType((*RegisterRequest)(nil), "mindspore.serving.proto.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "mindspore.serving.proto.RegisterReply")
	proto.RegisterType((*ExitRequest)(nil), "mindspore.serving.proto.ExitRequest")
	proto.RegisterType((*ExitReply)(nil), "mindspore.serving.proto.ExitReply")
	proto.RegisterType((*NotifyFailedRequest)(nil), "mindspore.serving.proto.NotifyFailedRequest")
	proto.RegisterType((*NotifyFailedReply)(nil), "mindspore.serving.proto.NotifyFailedReply")
	proto.RegisterType((*GetModelInfoRequest)(nil), "mindspore.serving.proto.GetModelInfoRequest")
	proto.RegisterType((*ModelSubGraphInfo)(nil), "mindspore.serving.proto.ModelSubGraphInfo")
	proto.RegisterType((*ModelInfo)(nil), "mindspore.serving.proto.ModelInfo")
	proto.RegisterType((*ModelInfos)(nil), "mindspore.serving.proto.ModelInfos")
	proto.RegisterMapType((map[string]*ModelInfo)(nil), "mindspore.serving.proto.ModelInfos.ModelInfosEntry")
	proto.RegisterType((*GetModelInfoReply)(nil), "mindspore.serving.proto.GetModelInfoReply")
}

func init() {
	proto.RegisterFile("mindspore_serving/proto/ms_master.proto", fileDescriptor_5d87d022b788c635)
}

var fileDescriptor_5d87d022b788c635 = []byte{
	// 813 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x4f, 0xe3, 0x46,
	0x10, 0xc7, 0x71, 0x20, 0xf1, 0x98, 0xf0, 0x67, 0x79, 0xa8, 0x95, 0x0a, 0x35, 0x35, 0x05, 0x2c,
	0x54, 0x05, 0x29, 0xbc, 0xa0, 0x56, 0xed, 0x4b, 0xa1, 0xa8, 0xaa, 0x92, 0xc2, 0x06, 0xa9, 0x52,
	0xa5, 0xd6, 0x72, 0xe2, 0x25, 0xb8, 0xd8, 0x5e, 0x77, 0xd7, 0x09, 0x0d, 0xd2, 0x7d, 0x83, 0x93,
	0xee, 0xf1, 0x3e, 0xd1, 0x7d, 0x94, 0xfb, 0x08, 0xf7, 0x7e, 0xda, 0xb5, 0x9d, 0x38, 0x39, 0x9c,
	0xe4, 0x4e, 0xbc, 0xed, 0x8e, 0x67, 0x7e, 0x33, 0xb3, 0xf3, 0x9b, 0x9f, 0xe1, 0x38, 0xf0, 0x42,
	0x97, 0x47, 0x94, 0x11, 0x9b, 0x13, 0x36, 0xf2, 0xc2, 0xc1, 0x69, 0xc4, 0x68, 0x4c, 0x4f, 0x03,
	0x6e, 0x07, 0x0e, 0x8f, 0x09, 0x6b, 0xca, 0x3b, 0xfa, 0x6a, 0xe2, 0xd8, 0x4c, 0x1d, 0x93, 0x0f,
	0x75, 0x6b, 0x01, 0x82, 0x34, 0xf4, 0x49, 0xe2, 0x69, 0xbe, 0x56, 0x61, 0xbb, 0x4b, 0xd8, 0xc8,
	0xe9, 0xf9, 0x04, 0x93, 0x41, 0x37, 0x22, 0x7d, 0x84, 0xa0, 0x1c, 0x3a, 0x01, 0x31, 0x94, 0x86,
	0x62, 0x69, 0x58, 0x9e, 0xd1, 0x21, 0x6c, 0x8d, 0x08, 0xe3, 0x1e, 0x0d, 0xed, 0x70, 0x18, 0xf4,
	0x08, 0x33, 0x4a, 0x0d, 0xc5, 0x2a, 0xe3, 0x5a, 0x6a, 0xed, 0x48, 0x23, 0xda, 0x07, 0xe8, 0x39,
	0x71, 0xff, 0xde, 0xe6, 0xde, 0x13, 0x31, 0xca, 0xd2, 0x45, 0x93, 0x96, 0xae, 0xf7, 0x44, 0x50,
	0x1b, 0x2a, 0x01, 0x89, 0xef, 0xa9, 0xcb, 0x8d, 0xf5, 0x86, 0x6a, 0xe9, 0xad, 0xb3, 0x66, 0x41,
	0x0b, 0xcd, 0xb9, 0xa2, 0x9a, 0x6d, 0x19, 0xf7, 0x5b, 0x78, 0x47, 0x71, 0x86, 0x81, 0x2e, 0x40,
	0x0f, 0xa8, 0x4b, 0x7c, 0xdb, 0x0b, 0xef, 0x28, 0x37, 0x36, 0x1a, 0x8a, 0xa5, 0xb7, 0x0e, 0x0a,
	0x21, 0xdb, 0xc2, 0x57, 0x20, 0x70, 0x0c, 0xc1, 0xe4, 0x2c, 0x6a, 0xa6, 0x8f, 0xa1, 0xed, 0x12,
	0xf1, 0x2c, 0x46, 0xa5, 0xa1, 0x58, 0x55, 0xac, 0xd1, 0xc7, 0xf0, 0x42, 0x1a, 0xea, 0x0f, 0x00,
	0xd3, 0xdc, 0xcf, 0xbe, 0xcd, 0x37, 0xa0, 0x7b, 0x61, 0x34, 0x8c, 0x6d, 0x71, 0xe3, 0x46, 0xa9,
	0xa1, 0x5a, 0x1a, 0x06, 0x69, 0xea, 0x08, 0x0b, 0xb2, 0x60, 0x87, 0x86, 0xfe, 0xd8, 0x4e, 0x8a,
	0xe5, 0xb1, 0x33, 0x20, 0x86, 0x2a, 0xf3, 0x6c, 0x09, 0xbb, 0xac, 0xab, 0x2b, 0xac, 0xe6, 0x5b,
	0x05, 0x6a, 0x7f, 0x52, 0xf6, 0x40, 0x58, 0x36, 0x8c, 0x7d, 0x80, 0x47, 0x69, 0xb0, 0x23, 0xcf,
	0x95, 0x69, 0xcb, 0x58, 0x4b, 0x2c, 0xd7, 0x9e, 0x8b, 0x0c, 0xa8, 0x38, 0xae, 0xcb, 0x08, 0xe7,
	0x72, 0x20, 0x1a, 0xce, 0xae, 0xa8, 0x0d, 0x35, 0x9e, 0xbe, 0xa1, 0xcd, 0x23, 0xd2, 0x97, 0xd3,
	0xd0, 0x5b, 0xd6, 0xaa, 0x2f, 0x8e, 0x37, 0xb3, 0x70, 0x71, 0x33, 0xff, 0x82, 0x6d, 0x4c, 0x06,
	0x9e, 0x60, 0x1f, 0x26, 0xff, 0x0d, 0x09, 0x8f, 0xd1, 0x15, 0xe8, 0x69, 0x69, 0x12, 0x5f, 0x91,
	0xf8, 0x47, 0x85, 0xf8, 0x33, 0x7d, 0xe1, 0xb4, 0x2b, 0x89, 0xfd, 0x07, 0xd4, 0xa6, 0xd8, 0x91,
	0x3f, 0x46, 0x3f, 0x83, 0x46, 0x18, 0xa3, 0xcc, 0x0e, 0xf8, 0x20, 0xc5, 0xfd, 0xb6, 0x10, 0xf7,
	0x52, 0x78, 0xb6, 0xf9, 0x00, 0x57, 0x49, 0x7a, 0x32, 0x8f, 0x41, 0xbf, 0xfc, 0xdf, 0x8b, 0xb3,
	0x42, 0x73, 0x8f, 0xa4, 0xcc, 0x3c, 0x92, 0xf9, 0x3b, 0x68, 0x89, 0xe3, 0x4b, 0x64, 0xbd, 0x81,
	0xbd, 0x0e, 0x8d, 0xbd, 0xbb, 0xf1, 0xaf, 0x8e, 0xe7, 0x13, 0x37, 0xcb, 0xbe, 0x64, 0x82, 0x5f,
	0xe7, 0xb3, 0x26, 0x33, 0x9c, 0x42, 0xee, 0xc1, 0xee, 0x2c, 0x64, 0xe4, 0x8f, 0x4d, 0x07, 0xf6,
	0xae, 0x48, 0x3c, 0x61, 0x73, 0x96, 0xe7, 0x20, 0x37, 0xf0, 0x1c, 0x47, 0x27, 0x63, 0xec, 0x14,
	0xef, 0x71, 0x6d, 0x6e, 0x8f, 0xcd, 0x37, 0x0a, 0xec, 0x26, 0xb4, 0x1c, 0xf6, 0xae, 0x98, 0x13,
	0xdd, 0x4b, 0xf2, 0xff, 0x08, 0x1b, 0x92, 0xd5, 0xdc, 0x50, 0xe5, 0xf6, 0x16, 0xaf, 0xda, 0x2d,
	0x09, 0x39, 0x65, 0xb2, 0xba, 0x34, 0x04, 0xfd, 0x04, 0x15, 0x3a, 0x8c, 0x65, 0x74, 0x79, 0xf5,
	0xe8, 0x2c, 0xc6, 0x7c, 0x05, 0xda, 0xa4, 0xe3, 0x39, 0x99, 0x29, 0xcd, 0xcb, 0xcc, 0x0d, 0x6c,
	0xf1, 0x61, 0x6f, 0x20, 0xea, 0x4e, 0xa5, 0x41, 0x91, 0x19, 0x4f, 0x16, 0x4b, 0x43, 0xbe, 0x57,
	0x5c, 0xcb, 0x10, 0xa4, 0x48, 0x98, 0xef, 0x14, 0x80, 0xa9, 0x7e, 0xa0, 0xdb, 0x59, 0xe5, 0x51,
	0x96, 0x88, 0xd9, 0x34, 0x32, 0x77, 0xbc, 0x0c, 0x63, 0x36, 0xce, 0x2b, 0x51, 0xdd, 0x81, 0xed,
	0xb9, 0xcf, 0x68, 0x07, 0xd4, 0x07, 0x32, 0x4e, 0x47, 0x29, 0x8e, 0xe8, 0x1c, 0xd6, 0x47, 0x8e,
	0x3f, 0x4c, 0xda, 0xd6, 0x5b, 0xe6, 0xf2, 0xa4, 0x38, 0x09, 0xf8, 0xa1, 0x74, 0xae, 0x98, 0xef,
	0x15, 0xd8, 0x9d, 0x25, 0x8f, 0x60, 0xfe, 0x0b, 0x52, 0x67, 0x5e, 0x94, 0xd5, 0x2f, 0x13, 0xe5,
	0x99, 0x5d, 0x2c, 0x7f, 0xf6, 0x2e, 0xb6, 0x3e, 0xa8, 0x50, 0x6d, 0x77, 0xdb, 0xf2, 0x6f, 0x89,
	0xfe, 0x81, 0x6a, 0xa6, 0x2f, 0xa8, 0x58, 0xff, 0xe6, 0xe4, 0xad, 0x7e, 0xb4, 0x82, 0xa7, 0x58,
	0xc7, 0x35, 0x84, 0xa1, 0x2c, 0x54, 0x04, 0x7d, 0x57, 0x5c, 0xe1, 0x54, 0x8d, 0xea, 0xe6, 0x12,
	0xaf, 0x04, 0xf3, 0x5f, 0xd8, 0xcc, 0x6f, 0x3e, 0xfa, 0xbe, 0x30, 0xea, 0x19, 0xcd, 0xa9, 0x9f,
	0xac, 0xe8, 0x9d, 0xe4, 0xfa, 0x1b, 0xb4, 0x5f, 0x1c, 0xdf, 0x97, 0xa3, 0x40, 0xc7, 0x85, 0xa1,
	0xd7, 0x8c, 0xb8, 0x5e, 0x7f, 0xd2, 0xc7, 0xe1, 0x72, 0xc7, 0x49, 0x2b, 0x79, 0xca, 0x2d, 0x68,
	0xe5, 0x19, 0x59, 0x5b, 0xd0, 0xca, 0x27, 0x3c, 0x36, 0xd7, 0x7a, 0x1b, 0xf2, 0xd3, 0xd9, 0xc7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xdb, 0x60, 0x61, 0x44, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MSMasterClient is the client API for MSMaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MSMasterClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Exit(ctx context.Context, in *ExitRequest, opts ...grpc.CallOption) (*ExitReply, error)
	NotifyFailed(ctx context.Context, in *NotifyFailedRequest, opts ...grpc.CallOption) (*NotifyFailedReply, error)
	CallModel(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictReply, error)
	GetModelInfo(ctx context.Context, in *GetModelInfoRequest, opts ...grpc.CallOption) (*GetModelInfoReply, error)
}

type mSMasterClient struct {
	cc *grpc.ClientConn
}

func NewMSMasterClient(cc *grpc.ClientConn) MSMasterClient {
	return &mSMasterClient{cc}
}

func (c *mSMasterClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/mindspore.serving.proto.MSMaster/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mSMasterClient) Exit(ctx context.Context, in *ExitRequest, opts ...grpc.CallOption) (*ExitReply, error) {
	out := new(ExitReply)
	err := c.cc.Invoke(ctx, "/mindspore.serving.proto.MSMaster/Exit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mSMasterClient) NotifyFailed(ctx context.Context, in *NotifyFailedRequest, opts ...grpc.CallOption) (*NotifyFailedReply, error) {
	out := new(NotifyFailedReply)
	err := c.cc.Invoke(ctx, "/mindspore.serving.proto.MSMaster/NotifyFailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mSMasterClient) CallModel(ctx context.Context, in *PredictRequest, opts ...grpc.CallOption) (*PredictReply, error) {
	out := new(PredictReply)
	err := c.cc.Invoke(ctx, "/mindspore.serving.proto.MSMaster/CallModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mSMasterClient) GetModelInfo(ctx context.Context, in *GetModelInfoRequest, opts ...grpc.CallOption) (*GetModelInfoReply, error) {
	out := new(GetModelInfoReply)
	err := c.cc.Invoke(ctx, "/mindspore.serving.proto.MSMaster/GetModelInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MSMasterServer is the server API for MSMaster service.
type MSMasterServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Exit(context.Context, *ExitRequest) (*ExitReply, error)
	NotifyFailed(context.Context, *NotifyFailedRequest) (*NotifyFailedReply, error)
	CallModel(context.Context, *PredictRequest) (*PredictReply, error)
	GetModelInfo(context.Context, *GetModelInfoRequest) (*GetModelInfoReply, error)
}

func RegisterMSMasterServer(s *grpc.Server, srv MSMasterServer) {
	s.RegisterService(&_MSMaster_serviceDesc, srv)
}

func _MSMaster_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSMasterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindspore.serving.proto.MSMaster/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSMasterServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MSMaster_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSMasterServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindspore.serving.proto.MSMaster/Exit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSMasterServer).Exit(ctx, req.(*ExitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MSMaster_NotifyFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyFailedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSMasterServer).NotifyFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindspore.serving.proto.MSMaster/NotifyFailed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSMasterServer).NotifyFailed(ctx, req.(*NotifyFailedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MSMaster_CallModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSMasterServer).CallModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindspore.serving.proto.MSMaster/CallModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSMasterServer).CallModel(ctx, req.(*PredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MSMaster_GetModelInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MSMasterServer).GetModelInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mindspore.serving.proto.MSMaster/GetModelInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MSMasterServer).GetModelInfo(ctx, req.(*GetModelInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MSMaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mindspore.serving.proto.MSMaster",
	HandlerType: (*MSMasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _MSMaster_Register_Handler,
		},
		{
			MethodName: "Exit",
			Handler:    _MSMaster_Exit_Handler,
		},
		{
			MethodName: "NotifyFailed",
			Handler:    _MSMaster_NotifyFailed_Handler,
		},
		{
			MethodName: "CallModel",
			Handler:    _MSMaster_CallModel_Handler,
		},
		{
			MethodName: "GetModelInfo",
			Handler:    _MSMaster_GetModelInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mindspore_serving/proto/ms_master.proto",
}
