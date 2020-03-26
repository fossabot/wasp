// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	packet "github.com/vx-labs/mqtt-protocol/packet"
	raftpb "go.etcd.io/etcd/raft/raftpb"
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

type Payload struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RaftContext struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RaftContext) Reset()         { *m = RaftContext{} }
func (m *RaftContext) String() string { return proto.CompactTextString(m) }
func (*RaftContext) ProtoMessage()    {}
func (*RaftContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *RaftContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RaftContext.Unmarshal(m, b)
}
func (m *RaftContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RaftContext.Marshal(b, m, deterministic)
}
func (m *RaftContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftContext.Merge(m, src)
}
func (m *RaftContext) XXX_Size() int {
	return xxx_messageInfo_RaftContext.Size(m)
}
func (m *RaftContext) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftContext.DiscardUnknown(m)
}

var xxx_messageInfo_RaftContext proto.InternalMessageInfo

func (m *RaftContext) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *RaftContext) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type JoinClusterResponse struct {
	Peers                []*RaftContext `protobuf:"bytes,1,rep,name=Peers,proto3" json:"Peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JoinClusterResponse) Reset()         { *m = JoinClusterResponse{} }
func (m *JoinClusterResponse) String() string { return proto.CompactTextString(m) }
func (*JoinClusterResponse) ProtoMessage()    {}
func (*JoinClusterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *JoinClusterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinClusterResponse.Unmarshal(m, b)
}
func (m *JoinClusterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinClusterResponse.Marshal(b, m, deterministic)
}
func (m *JoinClusterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinClusterResponse.Merge(m, src)
}
func (m *JoinClusterResponse) XXX_Size() int {
	return xxx_messageInfo_JoinClusterResponse.Size(m)
}
func (m *JoinClusterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinClusterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinClusterResponse proto.InternalMessageInfo

func (m *JoinClusterResponse) GetPeers() []*RaftContext {
	if m != nil {
		return m.Peers
	}
	return nil
}

type PeerResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerResponse) Reset()         { *m = PeerResponse{} }
func (m *PeerResponse) String() string { return proto.CompactTextString(m) }
func (*PeerResponse) ProtoMessage()    {}
func (*PeerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *PeerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerResponse.Unmarshal(m, b)
}
func (m *PeerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerResponse.Marshal(b, m, deterministic)
}
func (m *PeerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerResponse.Merge(m, src)
}
func (m *PeerResponse) XXX_Size() int {
	return xxx_messageInfo_PeerResponse.Size(m)
}
func (m *PeerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PeerResponse proto.InternalMessageInfo

func (m *PeerResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type CheckHealthRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckHealthRequest) Reset()         { *m = CheckHealthRequest{} }
func (m *CheckHealthRequest) String() string { return proto.CompactTextString(m) }
func (*CheckHealthRequest) ProtoMessage()    {}
func (*CheckHealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *CheckHealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckHealthRequest.Unmarshal(m, b)
}
func (m *CheckHealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckHealthRequest.Marshal(b, m, deterministic)
}
func (m *CheckHealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckHealthRequest.Merge(m, src)
}
func (m *CheckHealthRequest) XXX_Size() int {
	return xxx_messageInfo_CheckHealthRequest.Size(m)
}
func (m *CheckHealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckHealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckHealthRequest proto.InternalMessageInfo

type CheckHealthResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckHealthResponse) Reset()         { *m = CheckHealthResponse{} }
func (m *CheckHealthResponse) String() string { return proto.CompactTextString(m) }
func (*CheckHealthResponse) ProtoMessage()    {}
func (*CheckHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *CheckHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckHealthResponse.Unmarshal(m, b)
}
func (m *CheckHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckHealthResponse.Marshal(b, m, deterministic)
}
func (m *CheckHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckHealthResponse.Merge(m, src)
}
func (m *CheckHealthResponse) XXX_Size() int {
	return xxx_messageInfo_CheckHealthResponse.Size(m)
}
func (m *CheckHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckHealthResponse proto.InternalMessageInfo

type GetMembersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMembersRequest) Reset()         { *m = GetMembersRequest{} }
func (m *GetMembersRequest) String() string { return proto.CompactTextString(m) }
func (*GetMembersRequest) ProtoMessage()    {}
func (*GetMembersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *GetMembersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMembersRequest.Unmarshal(m, b)
}
func (m *GetMembersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMembersRequest.Marshal(b, m, deterministic)
}
func (m *GetMembersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMembersRequest.Merge(m, src)
}
func (m *GetMembersRequest) XXX_Size() int {
	return xxx_messageInfo_GetMembersRequest.Size(m)
}
func (m *GetMembersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMembersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMembersRequest proto.InternalMessageInfo

type Member struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	IsLeader             bool     `protobuf:"varint,3,opt,name=IsLeader,proto3" json:"IsLeader,omitempty"`
	IsAlive              bool     `protobuf:"varint,4,opt,name=IsAlive,proto3" json:"IsAlive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Member) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Member) GetIsLeader() bool {
	if m != nil {
		return m.IsLeader
	}
	return false
}

func (m *Member) GetIsAlive() bool {
	if m != nil {
		return m.IsAlive
	}
	return false
}

type GetMembersResponse struct {
	Members              []*Member `protobuf:"bytes,1,rep,name=Members,proto3" json:"Members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetMembersResponse) Reset()         { *m = GetMembersResponse{} }
func (m *GetMembersResponse) String() string { return proto.CompactTextString(m) }
func (*GetMembersResponse) ProtoMessage()    {}
func (*GetMembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *GetMembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMembersResponse.Unmarshal(m, b)
}
func (m *GetMembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMembersResponse.Marshal(b, m, deterministic)
}
func (m *GetMembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMembersResponse.Merge(m, src)
}
func (m *GetMembersResponse) XXX_Size() int {
	return xxx_messageInfo_GetMembersResponse.Size(m)
}
func (m *GetMembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMembersResponse proto.InternalMessageInfo

func (m *GetMembersResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type DistributeMessageRequest struct {
	Message              *packet.Publish `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DistributeMessageRequest) Reset()         { *m = DistributeMessageRequest{} }
func (m *DistributeMessageRequest) String() string { return proto.CompactTextString(m) }
func (*DistributeMessageRequest) ProtoMessage()    {}
func (*DistributeMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *DistributeMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributeMessageRequest.Unmarshal(m, b)
}
func (m *DistributeMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributeMessageRequest.Marshal(b, m, deterministic)
}
func (m *DistributeMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributeMessageRequest.Merge(m, src)
}
func (m *DistributeMessageRequest) XXX_Size() int {
	return xxx_messageInfo_DistributeMessageRequest.Size(m)
}
func (m *DistributeMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributeMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DistributeMessageRequest proto.InternalMessageInfo

func (m *DistributeMessageRequest) GetMessage() *packet.Publish {
	if m != nil {
		return m.Message
	}
	return nil
}

type DistributeMessageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributeMessageResponse) Reset()         { *m = DistributeMessageResponse{} }
func (m *DistributeMessageResponse) String() string { return proto.CompactTextString(m) }
func (*DistributeMessageResponse) ProtoMessage()    {}
func (*DistributeMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *DistributeMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributeMessageResponse.Unmarshal(m, b)
}
func (m *DistributeMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributeMessageResponse.Marshal(b, m, deterministic)
}
func (m *DistributeMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributeMessageResponse.Merge(m, src)
}
func (m *DistributeMessageResponse) XXX_Size() int {
	return xxx_messageInfo_DistributeMessageResponse.Size(m)
}
func (m *DistributeMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributeMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DistributeMessageResponse proto.InternalMessageInfo

type ShutdownRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownRequest) Reset()         { *m = ShutdownRequest{} }
func (m *ShutdownRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownRequest) ProtoMessage()    {}
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *ShutdownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownRequest.Unmarshal(m, b)
}
func (m *ShutdownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownRequest.Marshal(b, m, deterministic)
}
func (m *ShutdownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownRequest.Merge(m, src)
}
func (m *ShutdownRequest) XXX_Size() int {
	return xxx_messageInfo_ShutdownRequest.Size(m)
}
func (m *ShutdownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownRequest proto.InternalMessageInfo

type ShutdownResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutdownResponse) Reset()         { *m = ShutdownResponse{} }
func (m *ShutdownResponse) String() string { return proto.CompactTextString(m) }
func (*ShutdownResponse) ProtoMessage()    {}
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{12}
}

func (m *ShutdownResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutdownResponse.Unmarshal(m, b)
}
func (m *ShutdownResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutdownResponse.Marshal(b, m, deterministic)
}
func (m *ShutdownResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutdownResponse.Merge(m, src)
}
func (m *ShutdownResponse) XXX_Size() int {
	return xxx_messageInfo_ShutdownResponse.Size(m)
}
func (m *ShutdownResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutdownResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShutdownResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Payload)(nil), "api.Payload")
	proto.RegisterType((*RaftContext)(nil), "api.RaftContext")
	proto.RegisterType((*JoinClusterResponse)(nil), "api.JoinClusterResponse")
	proto.RegisterType((*PeerResponse)(nil), "api.PeerResponse")
	proto.RegisterType((*CheckHealthRequest)(nil), "api.CheckHealthRequest")
	proto.RegisterType((*CheckHealthResponse)(nil), "api.CheckHealthResponse")
	proto.RegisterType((*GetMembersRequest)(nil), "api.GetMembersRequest")
	proto.RegisterType((*Member)(nil), "api.Member")
	proto.RegisterType((*GetMembersResponse)(nil), "api.GetMembersResponse")
	proto.RegisterType((*DistributeMessageRequest)(nil), "api.DistributeMessageRequest")
	proto.RegisterType((*DistributeMessageResponse)(nil), "api.DistributeMessageResponse")
	proto.RegisterType((*ShutdownRequest)(nil), "api.ShutdownRequest")
	proto.RegisterType((*ShutdownResponse)(nil), "api.ShutdownResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x6d, 0x52, 0x93, 0xa6, 0xe3, 0xa8, 0x69, 0x36, 0xfd, 0x30, 0x46, 0x45, 0xd1, 0x4a, 0x44,
	0xe1, 0x50, 0x1b, 0xc2, 0x01, 0xa1, 0x0a, 0xa1, 0x90, 0x20, 0x08, 0x22, 0x28, 0x98, 0x1c, 0xb9,
	0xac, 0xed, 0x69, 0x6c, 0xd5, 0xc9, 0xba, 0xde, 0x75, 0x29, 0xff, 0x89, 0x1f, 0x89, 0xbc, 0xb6,
	0xd3, 0xb4, 0x49, 0x25, 0x2e, 0xde, 0x9d, 0x37, 0x5f, 0x3b, 0xf3, 0x9e, 0x0c, 0xfb, 0x2c, 0x0e,
	0xad, 0x38, 0xe1, 0x92, 0x93, 0x5d, 0x16, 0x87, 0x66, 0x77, 0xce, 0x2d, 0x94, 0x9e, 0x6f, 0x85,
	0xdc, 0xce, 0x4e, 0x3b, 0x61, 0x97, 0x52, 0x7d, 0x62, 0x57, 0x1d, 0x79, 0xb0, 0xf9, 0x6a, 0x1e,
	0xca, 0x20, 0x75, 0x2d, 0x8f, 0x2f, 0xec, 0x9b, 0xdb, 0xf3, 0x88, 0xb9, 0xc2, 0x5e, 0x5c, 0x4b,
	0x79, 0xae, 0xdc, 0x1e, 0x8f, 0xec, 0x98, 0x79, 0x57, 0x28, 0xed, 0xd8, 0xcd, 0x33, 0xe8, 0x19,
	0xec, 0x4d, 0xd9, 0x9f, 0x88, 0x33, 0x9f, 0x10, 0xd0, 0x46, 0x4c, 0x32, 0xa3, 0xd2, 0xa9, 0xf4,
	0x1a, 0x8e, 0xba, 0xd3, 0xb7, 0xa0, 0x3b, 0xec, 0x52, 0x0e, 0xf9, 0x52, 0xe2, 0xad, 0x24, 0x07,
	0x50, 0x1d, 0x8f, 0x54, 0x80, 0xe6, 0x54, 0xc7, 0x23, 0x62, 0xc0, 0xde, 0xc0, 0xf7, 0x13, 0x14,
	0xc2, 0xa8, 0x76, 0x2a, 0xbd, 0x7d, 0xa7, 0x34, 0xe9, 0x7b, 0x68, 0x7f, 0xe5, 0xe1, 0x72, 0x18,
	0xa5, 0x42, 0x62, 0xe2, 0xa0, 0x88, 0xf9, 0x52, 0x20, 0xe9, 0xc2, 0x93, 0x29, 0x62, 0x22, 0x8c,
	0x4a, 0x67, 0xb7, 0xa7, 0xf7, 0x0f, 0xad, 0x6c, 0xd0, 0xb5, 0x0e, 0x4e, 0xee, 0xa6, 0x5d, 0x68,
	0x64, 0x97, 0x55, 0xde, 0x09, 0xd4, 0x84, 0x64, 0x32, 0x15, 0xaa, 0x79, 0xdd, 0x29, 0x2c, 0x7a,
	0x04, 0x64, 0x18, 0xa0, 0x77, 0xf5, 0x05, 0x59, 0x24, 0x03, 0x07, 0xaf, 0x53, 0x14, 0x92, 0x1e,
	0x43, 0xfb, 0x1e, 0x9a, 0x17, 0xa1, 0x6d, 0x68, 0x7d, 0x46, 0x39, 0xc1, 0x85, 0x8b, 0x89, 0x28,
	0x63, 0x03, 0xa8, 0xe5, 0xc8, 0xff, 0x0f, 0x47, 0x4c, 0xa8, 0x8f, 0xc5, 0x37, 0x64, 0x3e, 0x26,
	0xc6, 0xae, 0x7a, 0xcf, 0xca, 0xce, 0xb2, 0xc6, 0x62, 0x10, 0x85, 0x37, 0x68, 0x68, 0xca, 0x55,
	0x9a, 0xf4, 0x02, 0xc8, 0x7a, 0xfb, 0x62, 0xb2, 0x17, 0xb0, 0x57, 0x40, 0xc5, 0x4e, 0x74, 0xb5,
	0x93, 0x1c, 0x73, 0x4a, 0x1f, 0xfd, 0x04, 0xc6, 0x28, 0x14, 0x32, 0x09, 0xdd, 0x54, 0xe2, 0x04,
	0x85, 0x60, 0x73, 0x2c, 0x46, 0x20, 0x2f, 0xb3, 0x12, 0x0a, 0x51, 0xaf, 0xd7, 0xfb, 0x4d, 0x2b,
	0xa7, 0xd9, 0x9a, 0xa6, 0x6e, 0x14, 0x8a, 0xc0, 0x29, 0xfd, 0xf4, 0x19, 0x3c, 0xdd, 0x52, 0xa6,
	0xd8, 0x4f, 0x0b, 0x9a, 0x3f, 0x83, 0x54, 0xfa, 0xfc, 0xf7, 0xb2, 0xdc, 0x0e, 0x81, 0xc3, 0x3b,
	0x28, 0x0f, 0xeb, 0xff, 0xad, 0x82, 0x96, 0x51, 0x46, 0x5e, 0xc3, 0xc1, 0x34, 0xe1, 0x1e, 0x0a,
	0x51, 0x54, 0x22, 0x4d, 0x2b, 0xd7, 0xa4, 0x55, 0x00, 0x66, 0x43, 0x0d, 0x53, 0x28, 0x8c, 0xee,
	0x90, 0x0b, 0xd0, 0xd7, 0x64, 0x41, 0x36, 0xf8, 0x37, 0x0d, 0x85, 0x6c, 0x91, 0x0e, 0xdd, 0x21,
	0x36, 0xd4, 0xc6, 0x22, 0x93, 0xc5, 0x96, 0xbc, 0x56, 0xde, 0x08, 0xef, 0x25, 0x7c, 0x04, 0x7d,
	0x4d, 0x07, 0xe4, 0x54, 0xc5, 0x6c, 0xea, 0xa5, 0x68, 0xba, 0x4d, 0x32, 0x3b, 0xe4, 0x03, 0xc0,
	0x1d, 0x6b, 0xe4, 0x44, 0x45, 0x6e, 0xa8, 0xc8, 0x3c, 0xdd, 0xc0, 0xcb, 0x02, 0xfd, 0x5f, 0xa0,
	0x4d, 0x7e, 0xcc, 0x66, 0x64, 0x06, 0xad, 0x8d, 0xd5, 0x93, 0x33, 0x95, 0xf7, 0x18, 0xb3, 0xe6,
	0xf3, 0xc7, 0xdc, 0xab, 0xea, 0x03, 0xd0, 0xbe, 0x73, 0x1f, 0xc9, 0x3b, 0xa8, 0x97, 0x44, 0x91,
	0x23, 0x95, 0xf5, 0x80, 0x4a, 0xf3, 0xf8, 0x01, 0x5a, 0x96, 0x70, 0x6b, 0xea, 0x4f, 0xf0, 0xe6,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0xb3, 0x1d, 0xbf, 0x75, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RaftClient is the client API for Raft service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RaftClient interface {
	ProcessMessage(ctx context.Context, in *raftpb.Message, opts ...grpc.CallOption) (*Payload, error)
	JoinCluster(ctx context.Context, in *RaftContext, opts ...grpc.CallOption) (*JoinClusterResponse, error)
	IsPeer(ctx context.Context, in *RaftContext, opts ...grpc.CallOption) (*PeerResponse, error)
	CheckHealth(ctx context.Context, in *CheckHealthRequest, opts ...grpc.CallOption) (*CheckHealthResponse, error)
	GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersResponse, error)
}

type raftClient struct {
	cc *grpc.ClientConn
}

func NewRaftClient(cc *grpc.ClientConn) RaftClient {
	return &raftClient{cc}
}

func (c *raftClient) ProcessMessage(ctx context.Context, in *raftpb.Message, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := c.cc.Invoke(ctx, "/api.Raft/ProcessMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) JoinCluster(ctx context.Context, in *RaftContext, opts ...grpc.CallOption) (*JoinClusterResponse, error) {
	out := new(JoinClusterResponse)
	err := c.cc.Invoke(ctx, "/api.Raft/JoinCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) IsPeer(ctx context.Context, in *RaftContext, opts ...grpc.CallOption) (*PeerResponse, error) {
	out := new(PeerResponse)
	err := c.cc.Invoke(ctx, "/api.Raft/IsPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) CheckHealth(ctx context.Context, in *CheckHealthRequest, opts ...grpc.CallOption) (*CheckHealthResponse, error) {
	out := new(CheckHealthResponse)
	err := c.cc.Invoke(ctx, "/api.Raft/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) GetMembers(ctx context.Context, in *GetMembersRequest, opts ...grpc.CallOption) (*GetMembersResponse, error) {
	out := new(GetMembersResponse)
	err := c.cc.Invoke(ctx, "/api.Raft/GetMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftServer is the server API for Raft service.
type RaftServer interface {
	ProcessMessage(context.Context, *raftpb.Message) (*Payload, error)
	JoinCluster(context.Context, *RaftContext) (*JoinClusterResponse, error)
	IsPeer(context.Context, *RaftContext) (*PeerResponse, error)
	CheckHealth(context.Context, *CheckHealthRequest) (*CheckHealthResponse, error)
	GetMembers(context.Context, *GetMembersRequest) (*GetMembersResponse, error)
}

// UnimplementedRaftServer can be embedded to have forward compatible implementations.
type UnimplementedRaftServer struct {
}

func (*UnimplementedRaftServer) ProcessMessage(ctx context.Context, req *raftpb.Message) (*Payload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessMessage not implemented")
}
func (*UnimplementedRaftServer) JoinCluster(ctx context.Context, req *RaftContext) (*JoinClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCluster not implemented")
}
func (*UnimplementedRaftServer) IsPeer(ctx context.Context, req *RaftContext) (*PeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsPeer not implemented")
}
func (*UnimplementedRaftServer) CheckHealth(ctx context.Context, req *CheckHealthRequest) (*CheckHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (*UnimplementedRaftServer) GetMembers(ctx context.Context, req *GetMembersRequest) (*GetMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}

func RegisterRaftServer(s *grpc.Server, srv RaftServer) {
	s.RegisterService(&_Raft_serviceDesc, srv)
}

func _Raft_ProcessMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(raftpb.Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).ProcessMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Raft/ProcessMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).ProcessMessage(ctx, req.(*raftpb.Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_JoinCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).JoinCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Raft/JoinCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).JoinCluster(ctx, req.(*RaftContext))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_IsPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftContext)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).IsPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Raft/IsPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).IsPeer(ctx, req.(*RaftContext))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Raft/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).CheckHealth(ctx, req.(*CheckHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Raft/GetMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).GetMembers(ctx, req.(*GetMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Raft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Raft",
	HandlerType: (*RaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessMessage",
			Handler:    _Raft_ProcessMessage_Handler,
		},
		{
			MethodName: "JoinCluster",
			Handler:    _Raft_JoinCluster_Handler,
		},
		{
			MethodName: "IsPeer",
			Handler:    _Raft_IsPeer_Handler,
		},
		{
			MethodName: "CheckHealth",
			Handler:    _Raft_CheckHealth_Handler,
		},
		{
			MethodName: "GetMembers",
			Handler:    _Raft_GetMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// MQTTClient is the client API for MQTT service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MQTTClient interface {
	DistributeMessage(ctx context.Context, in *DistributeMessageRequest, opts ...grpc.CallOption) (*DistributeMessageResponse, error)
}

type mQTTClient struct {
	cc *grpc.ClientConn
}

func NewMQTTClient(cc *grpc.ClientConn) MQTTClient {
	return &mQTTClient{cc}
}

func (c *mQTTClient) DistributeMessage(ctx context.Context, in *DistributeMessageRequest, opts ...grpc.CallOption) (*DistributeMessageResponse, error) {
	out := new(DistributeMessageResponse)
	err := c.cc.Invoke(ctx, "/api.MQTT/DistributeMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MQTTServer is the server API for MQTT service.
type MQTTServer interface {
	DistributeMessage(context.Context, *DistributeMessageRequest) (*DistributeMessageResponse, error)
}

// UnimplementedMQTTServer can be embedded to have forward compatible implementations.
type UnimplementedMQTTServer struct {
}

func (*UnimplementedMQTTServer) DistributeMessage(ctx context.Context, req *DistributeMessageRequest) (*DistributeMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DistributeMessage not implemented")
}

func RegisterMQTTServer(s *grpc.Server, srv MQTTServer) {
	s.RegisterService(&_MQTT_serviceDesc, srv)
}

func _MQTT_DistributeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DistributeMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MQTTServer).DistributeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MQTT/DistributeMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MQTTServer).DistributeMessage(ctx, req.(*DistributeMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MQTT_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.MQTT",
	HandlerType: (*MQTTServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DistributeMessage",
			Handler:    _MQTT_DistributeMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*ShutdownResponse, error) {
	out := new(ShutdownResponse)
	err := c.cc.Invoke(ctx, "/api.Node/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	Shutdown(context.Context, *ShutdownRequest) (*ShutdownResponse, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) Shutdown(ctx context.Context, req *ShutdownRequest) (*ShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shutdown",
			Handler:    _Node_Shutdown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
