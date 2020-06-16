// Code generated by protoc-gen-go. DO NOT EDIT.
// source: timer.proto

package timer

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

type CreateTimerRequest struct {
	Mode                 *TimerMode      `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	Payload              string          `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Callback             *CallbackAction `protobuf:"bytes,3,opt,name=callback,proto3" json:"callback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateTimerRequest) Reset()         { *m = CreateTimerRequest{} }
func (m *CreateTimerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTimerRequest) ProtoMessage()    {}
func (*CreateTimerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{0}
}

func (m *CreateTimerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTimerRequest.Unmarshal(m, b)
}
func (m *CreateTimerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTimerRequest.Marshal(b, m, deterministic)
}
func (m *CreateTimerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTimerRequest.Merge(m, src)
}
func (m *CreateTimerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTimerRequest.Size(m)
}
func (m *CreateTimerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTimerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTimerRequest proto.InternalMessageInfo

func (m *CreateTimerRequest) GetMode() *TimerMode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *CreateTimerRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *CreateTimerRequest) GetCallback() *CallbackAction {
	if m != nil {
		return m.Callback
	}
	return nil
}

type CreateTimerReply struct {
	TimerID              string   `protobuf:"bytes,1,opt,name=timerID,proto3" json:"timerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTimerReply) Reset()         { *m = CreateTimerReply{} }
func (m *CreateTimerReply) String() string { return proto.CompactTextString(m) }
func (*CreateTimerReply) ProtoMessage()    {}
func (*CreateTimerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{1}
}

func (m *CreateTimerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTimerReply.Unmarshal(m, b)
}
func (m *CreateTimerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTimerReply.Marshal(b, m, deterministic)
}
func (m *CreateTimerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTimerReply.Merge(m, src)
}
func (m *CreateTimerReply) XXX_Size() int {
	return xxx_messageInfo_CreateTimerReply.Size(m)
}
func (m *CreateTimerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTimerReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTimerReply proto.InternalMessageInfo

func (m *CreateTimerReply) GetTimerID() string {
	if m != nil {
		return m.TimerID
	}
	return ""
}

type DeleteTimerRequest struct {
	TimerID              string   `protobuf:"bytes,1,opt,name=timerID,proto3" json:"timerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTimerRequest) Reset()         { *m = DeleteTimerRequest{} }
func (m *DeleteTimerRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTimerRequest) ProtoMessage()    {}
func (*DeleteTimerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{2}
}

func (m *DeleteTimerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTimerRequest.Unmarshal(m, b)
}
func (m *DeleteTimerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTimerRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTimerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTimerRequest.Merge(m, src)
}
func (m *DeleteTimerRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTimerRequest.Size(m)
}
func (m *DeleteTimerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTimerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTimerRequest proto.InternalMessageInfo

func (m *DeleteTimerRequest) GetTimerID() string {
	if m != nil {
		return m.TimerID
	}
	return ""
}

type DeleteTimerReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTimerReply) Reset()         { *m = DeleteTimerReply{} }
func (m *DeleteTimerReply) String() string { return proto.CompactTextString(m) }
func (*DeleteTimerReply) ProtoMessage()    {}
func (*DeleteTimerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{3}
}

func (m *DeleteTimerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTimerReply.Unmarshal(m, b)
}
func (m *DeleteTimerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTimerReply.Marshal(b, m, deterministic)
}
func (m *DeleteTimerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTimerReply.Merge(m, src)
}
func (m *DeleteTimerReply) XXX_Size() int {
	return xxx_messageInfo_DeleteTimerReply.Size(m)
}
func (m *DeleteTimerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTimerReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTimerReply proto.InternalMessageInfo

type CallbackAction struct {
	Type                 string            `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Method               string            `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Uri                  string            `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Headers              map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload              string            `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CallbackAction) Reset()         { *m = CallbackAction{} }
func (m *CallbackAction) String() string { return proto.CompactTextString(m) }
func (*CallbackAction) ProtoMessage()    {}
func (*CallbackAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{4}
}

func (m *CallbackAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackAction.Unmarshal(m, b)
}
func (m *CallbackAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackAction.Marshal(b, m, deterministic)
}
func (m *CallbackAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackAction.Merge(m, src)
}
func (m *CallbackAction) XXX_Size() int {
	return xxx_messageInfo_CallbackAction.Size(m)
}
func (m *CallbackAction) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackAction.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackAction proto.InternalMessageInfo

func (m *CallbackAction) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CallbackAction) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *CallbackAction) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *CallbackAction) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *CallbackAction) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type TimerMode struct {
	Mode                 string   `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	Interval             uint32   `protobuf:"varint,2,opt,name=interval,proto3" json:"interval,omitempty"`
	Timestamp            uint64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimerMode) Reset()         { *m = TimerMode{} }
func (m *TimerMode) String() string { return proto.CompactTextString(m) }
func (*TimerMode) ProtoMessage()    {}
func (*TimerMode) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{5}
}

func (m *TimerMode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimerMode.Unmarshal(m, b)
}
func (m *TimerMode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimerMode.Marshal(b, m, deterministic)
}
func (m *TimerMode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimerMode.Merge(m, src)
}
func (m *TimerMode) XXX_Size() int {
	return xxx_messageInfo_TimerMode.Size(m)
}
func (m *TimerMode) XXX_DiscardUnknown() {
	xxx_messageInfo_TimerMode.DiscardUnknown(m)
}

var xxx_messageInfo_TimerMode proto.InternalMessageInfo

func (m *TimerMode) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *TimerMode) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *TimerMode) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type TimerCreation struct {
	TimerID              string     `protobuf:"bytes,1,opt,name=timerID,proto3" json:"timerID,omitempty"`
	Timestamp            uint64     `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Info                 *TimerInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TimerCreation) Reset()         { *m = TimerCreation{} }
func (m *TimerCreation) String() string { return proto.CompactTextString(m) }
func (*TimerCreation) ProtoMessage()    {}
func (*TimerCreation) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{6}
}

func (m *TimerCreation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimerCreation.Unmarshal(m, b)
}
func (m *TimerCreation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimerCreation.Marshal(b, m, deterministic)
}
func (m *TimerCreation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimerCreation.Merge(m, src)
}
func (m *TimerCreation) XXX_Size() int {
	return xxx_messageInfo_TimerCreation.Size(m)
}
func (m *TimerCreation) XXX_DiscardUnknown() {
	xxx_messageInfo_TimerCreation.DiscardUnknown(m)
}

var xxx_messageInfo_TimerCreation proto.InternalMessageInfo

func (m *TimerCreation) GetTimerID() string {
	if m != nil {
		return m.TimerID
	}
	return ""
}

func (m *TimerCreation) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TimerCreation) GetInfo() *TimerInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type TimerInfo struct {
	Payload              string          `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Callback             *CallbackAction `protobuf:"bytes,2,opt,name=callback,proto3" json:"callback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TimerInfo) Reset()         { *m = TimerInfo{} }
func (m *TimerInfo) String() string { return proto.CompactTextString(m) }
func (*TimerInfo) ProtoMessage()    {}
func (*TimerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{7}
}

func (m *TimerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimerInfo.Unmarshal(m, b)
}
func (m *TimerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimerInfo.Marshal(b, m, deterministic)
}
func (m *TimerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimerInfo.Merge(m, src)
}
func (m *TimerInfo) XXX_Size() int {
	return xxx_messageInfo_TimerInfo.Size(m)
}
func (m *TimerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TimerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TimerInfo proto.InternalMessageInfo

func (m *TimerInfo) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *TimerInfo) GetCallback() *CallbackAction {
	if m != nil {
		return m.Callback
	}
	return nil
}

type TimerDeletion struct {
	TimerID              string   `protobuf:"bytes,1,opt,name=timerID,proto3" json:"timerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimerDeletion) Reset()         { *m = TimerDeletion{} }
func (m *TimerDeletion) String() string { return proto.CompactTextString(m) }
func (*TimerDeletion) ProtoMessage()    {}
func (*TimerDeletion) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{8}
}

func (m *TimerDeletion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimerDeletion.Unmarshal(m, b)
}
func (m *TimerDeletion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimerDeletion.Marshal(b, m, deterministic)
}
func (m *TimerDeletion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimerDeletion.Merge(m, src)
}
func (m *TimerDeletion) XXX_Size() int {
	return xxx_messageInfo_TimerDeletion.Size(m)
}
func (m *TimerDeletion) XXX_DiscardUnknown() {
	xxx_messageInfo_TimerDeletion.DiscardUnknown(m)
}

var xxx_messageInfo_TimerDeletion proto.InternalMessageInfo

func (m *TimerDeletion) GetTimerID() string {
	if m != nil {
		return m.TimerID
	}
	return ""
}

type TimerTriggerInfo struct {
	TimerID              string     `protobuf:"bytes,1,opt,name=timerID,proto3" json:"timerID,omitempty"`
	Info                 *TimerInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TimerTriggerInfo) Reset()         { *m = TimerTriggerInfo{} }
func (m *TimerTriggerInfo) String() string { return proto.CompactTextString(m) }
func (*TimerTriggerInfo) ProtoMessage()    {}
func (*TimerTriggerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad0307ee16b652d2, []int{9}
}

func (m *TimerTriggerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimerTriggerInfo.Unmarshal(m, b)
}
func (m *TimerTriggerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimerTriggerInfo.Marshal(b, m, deterministic)
}
func (m *TimerTriggerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimerTriggerInfo.Merge(m, src)
}
func (m *TimerTriggerInfo) XXX_Size() int {
	return xxx_messageInfo_TimerTriggerInfo.Size(m)
}
func (m *TimerTriggerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TimerTriggerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TimerTriggerInfo proto.InternalMessageInfo

func (m *TimerTriggerInfo) GetTimerID() string {
	if m != nil {
		return m.TimerID
	}
	return ""
}

func (m *TimerTriggerInfo) GetInfo() *TimerInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTimerRequest)(nil), "timer.CreateTimerRequest")
	proto.RegisterType((*CreateTimerReply)(nil), "timer.CreateTimerReply")
	proto.RegisterType((*DeleteTimerRequest)(nil), "timer.DeleteTimerRequest")
	proto.RegisterType((*DeleteTimerReply)(nil), "timer.DeleteTimerReply")
	proto.RegisterType((*CallbackAction)(nil), "timer.CallbackAction")
	proto.RegisterMapType((map[string]string)(nil), "timer.CallbackAction.HeadersEntry")
	proto.RegisterType((*TimerMode)(nil), "timer.TimerMode")
	proto.RegisterType((*TimerCreation)(nil), "timer.TimerCreation")
	proto.RegisterType((*TimerInfo)(nil), "timer.TimerInfo")
	proto.RegisterType((*TimerDeletion)(nil), "timer.TimerDeletion")
	proto.RegisterType((*TimerTriggerInfo)(nil), "timer.TimerTriggerInfo")
}

func init() { proto.RegisterFile("timer.proto", fileDescriptor_ad0307ee16b652d2) }

var fileDescriptor_ad0307ee16b652d2 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0xa4, 0xad, 0x27, 0x14, 0x59, 0x23, 0x3e, 0x4c, 0xc4, 0x21, 0x5a, 0x71, 0x08,
	0x12, 0x8a, 0x44, 0xb8, 0xa0, 0x8a, 0x0b, 0x4a, 0x91, 0xe8, 0x81, 0xcb, 0x2a, 0x07, 0x38, 0x6e,
	0x93, 0x69, 0x6b, 0xd5, 0xf6, 0x9a, 0xcd, 0xa6, 0x92, 0x7f, 0x01, 0x67, 0xfe, 0x22, 0xbf, 0x04,
	0xed, 0x64, 0xed, 0xda, 0x4d, 0x53, 0xf5, 0xb6, 0xb3, 0x3b, 0xf3, 0xfc, 0xde, 0xcc, 0x1b, 0xc3,
	0xd0, 0xa6, 0x39, 0x99, 0x69, 0x69, 0xb4, 0xd5, 0x38, 0xe0, 0x40, 0xfc, 0x09, 0x00, 0xe7, 0x86,
	0x94, 0xa5, 0x85, 0x8b, 0x25, 0xfd, 0xde, 0xd0, 0xda, 0xe2, 0x3b, 0xe8, 0xe7, 0x7a, 0x45, 0x49,
	0x30, 0x0e, 0x26, 0xc3, 0x59, 0x3c, 0xdd, 0x56, 0x72, 0xca, 0x0f, 0xbd, 0x22, 0xc9, 0xaf, 0x98,
	0xc0, 0x51, 0xa9, 0xaa, 0x4c, 0xab, 0x55, 0x12, 0x8e, 0x83, 0x49, 0x24, 0xeb, 0x10, 0x3f, 0xc2,
	0xf1, 0x52, 0x65, 0xd9, 0x85, 0x5a, 0xde, 0x24, 0x3d, 0xc6, 0x78, 0xe9, 0x31, 0xe6, 0xfe, 0xfa,
	0xeb, 0xd2, 0xa6, 0xba, 0x90, 0x4d, 0x9a, 0xf8, 0x00, 0x71, 0x87, 0x48, 0x99, 0x55, 0xee, 0x03,
	0x5c, 0x75, 0x7e, 0xc6, 0x4c, 0x22, 0x59, 0x87, 0x62, 0x0a, 0x78, 0x46, 0x19, 0xdd, 0xa3, 0xbd,
	0x3f, 0x1f, 0x21, 0xee, 0xe4, 0x97, 0x59, 0x25, 0xfe, 0x05, 0xf0, 0xbc, 0x4b, 0x07, 0x11, 0xfa,
	0x8b, 0xaa, 0x24, 0x5f, 0xcd, 0x67, 0x7c, 0x05, 0x87, 0x39, 0xd9, 0x6b, 0x5d, 0x8b, 0xf4, 0x11,
	0xc6, 0xd0, 0xdb, 0x98, 0x94, 0xe5, 0x45, 0xd2, 0x1d, 0xf1, 0x0b, 0x1c, 0x5d, 0x93, 0x5a, 0x91,
	0x59, 0x27, 0xfd, 0x71, 0x6f, 0x32, 0x9c, 0x89, 0x07, 0x45, 0x4f, 0xbf, 0x6f, 0x93, 0xbe, 0x15,
	0xd6, 0x54, 0xb2, 0x2e, 0x69, 0x77, 0x73, 0xd0, 0xe9, 0xe6, 0xe8, 0x14, 0x9e, 0xb5, 0x4b, 0xdc,
	0x97, 0x6f, 0xa8, 0xf2, 0x24, 0xdd, 0x11, 0x5f, 0xc0, 0xe0, 0x56, 0x65, 0x1b, 0xf2, 0x14, 0xb7,
	0xc1, 0x69, 0xf8, 0x39, 0x10, 0xbf, 0x20, 0x6a, 0xc6, 0xe6, 0xe4, 0x35, 0x63, 0x8d, 0xfc, 0x10,
	0x47, 0x70, 0x9c, 0x16, 0x96, 0xcc, 0xad, 0xca, 0xb8, 0xfa, 0x44, 0x36, 0x31, 0xbe, 0x85, 0xc8,
	0x09, 0x58, 0x5b, 0x95, 0x97, 0x2c, 0xb4, 0x2f, 0xef, 0x2e, 0x44, 0x0e, 0x27, 0x0c, 0xcd, 0x63,
	0x73, 0xdd, 0xdb, 0xdb, 0xfe, 0x2e, 0x50, 0x78, 0x0f, 0xc8, 0xb9, 0x2d, 0x2d, 0x2e, 0xb5, 0x77,
	0x4a, 0xc7, 0x6d, 0xe7, 0xc5, 0xa5, 0x96, 0xfc, 0x2a, 0x7e, 0x7a, 0x25, 0xee, 0xaa, 0xdd, 0xac,
	0x60, 0xbf, 0xf5, 0xc2, 0xa7, 0x59, 0xef, 0xbd, 0x17, 0xc2, 0x0e, 0x79, 0x54, 0x88, 0x90, 0x10,
	0x73, 0xea, 0xc2, 0xa4, 0x57, 0x57, 0x77, 0x5c, 0xf6, 0xc8, 0xae, 0x85, 0x85, 0x8f, 0x09, 0x9b,
	0xfd, 0x0d, 0x60, 0xc0, 0x77, 0x38, 0x87, 0x61, 0x6b, 0x07, 0xf0, 0x4d, 0x4d, 0x7c, 0x67, 0x41,
	0x47, 0xaf, 0x1f, 0x7a, 0x72, 0xa6, 0x3e, 0x70, 0x20, 0x2d, 0xab, 0x37, 0x20, 0xbb, 0xeb, 0xd2,
	0x80, 0xec, 0x6c, 0xc6, 0xc1, 0xc5, 0x21, 0xff, 0x25, 0x3e, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x37, 0xb7, 0x20, 0xed, 0x34, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimerClient is the client API for Timer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimerClient interface {
	CreateTimer(ctx context.Context, in *CreateTimerRequest, opts ...grpc.CallOption) (*CreateTimerReply, error)
	DeleteTimer(ctx context.Context, in *DeleteTimerRequest, opts ...grpc.CallOption) (*DeleteTimerReply, error)
}

type timerClient struct {
	cc *grpc.ClientConn
}

func NewTimerClient(cc *grpc.ClientConn) TimerClient {
	return &timerClient{cc}
}

func (c *timerClient) CreateTimer(ctx context.Context, in *CreateTimerRequest, opts ...grpc.CallOption) (*CreateTimerReply, error) {
	out := new(CreateTimerReply)
	err := c.cc.Invoke(ctx, "/timer.Timer/CreateTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timerClient) DeleteTimer(ctx context.Context, in *DeleteTimerRequest, opts ...grpc.CallOption) (*DeleteTimerReply, error) {
	out := new(DeleteTimerReply)
	err := c.cc.Invoke(ctx, "/timer.Timer/DeleteTimer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimerServer is the server API for Timer service.
type TimerServer interface {
	CreateTimer(context.Context, *CreateTimerRequest) (*CreateTimerReply, error)
	DeleteTimer(context.Context, *DeleteTimerRequest) (*DeleteTimerReply, error)
}

// UnimplementedTimerServer can be embedded to have forward compatible implementations.
type UnimplementedTimerServer struct {
}

func (*UnimplementedTimerServer) CreateTimer(ctx context.Context, req *CreateTimerRequest) (*CreateTimerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTimer not implemented")
}
func (*UnimplementedTimerServer) DeleteTimer(ctx context.Context, req *DeleteTimerRequest) (*DeleteTimerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTimer not implemented")
}

func RegisterTimerServer(s *grpc.Server, srv TimerServer) {
	s.RegisterService(&_Timer_serviceDesc, srv)
}

func _Timer_CreateTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServer).CreateTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timer.Timer/CreateTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServer).CreateTimer(ctx, req.(*CreateTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Timer_DeleteTimer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTimerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimerServer).DeleteTimer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timer.Timer/DeleteTimer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimerServer).DeleteTimer(ctx, req.(*DeleteTimerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Timer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "timer.Timer",
	HandlerType: (*TimerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTimer",
			Handler:    _Timer_CreateTimer_Handler,
		},
		{
			MethodName: "DeleteTimer",
			Handler:    _Timer_DeleteTimer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timer.proto",
}
