// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/eosio/search/v1/search.proto

package pbsearcheos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/codec/v1"
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

type DocumentID struct {
	BlockNum             uint64   `protobuf:"varint,1,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	ActionIndex          uint64   `protobuf:"varint,2,opt,name=actionIndex,proto3" json:"actionIndex,omitempty"`
	TransactionIndex     uint64   `protobuf:"varint,3,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	TransactionIDPrefix  []byte   `protobuf:"bytes,4,opt,name=transactionIDPrefix,proto3" json:"transactionIDPrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocumentID) Reset()         { *m = DocumentID{} }
func (m *DocumentID) String() string { return proto.CompactTextString(m) }
func (*DocumentID) ProtoMessage()    {}
func (*DocumentID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6416b04c85aeead, []int{0}
}

func (m *DocumentID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentID.Unmarshal(m, b)
}
func (m *DocumentID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentID.Marshal(b, m, deterministic)
}
func (m *DocumentID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentID.Merge(m, src)
}
func (m *DocumentID) XXX_Size() int {
	return xxx_messageInfo_DocumentID.Size(m)
}
func (m *DocumentID) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentID.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentID proto.InternalMessageInfo

func (m *DocumentID) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *DocumentID) GetActionIndex() uint64 {
	if m != nil {
		return m.ActionIndex
	}
	return 0
}

func (m *DocumentID) GetTransactionIndex() uint64 {
	if m != nil {
		return m.TransactionIndex
	}
	return 0
}

func (m *DocumentID) GetTransactionIDPrefix() []byte {
	if m != nil {
		return m.TransactionIDPrefix
	}
	return nil
}

type Match struct {
	ActionIndexes        []uint32         `protobuf:"varint,1,rep,packed,name=actionIndexes,proto3" json:"actionIndexes,omitempty"`
	Block                *BlockTrxPayload `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6416b04c85aeead, []int{1}
}

func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (m *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(m, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetActionIndexes() []uint32 {
	if m != nil {
		return m.ActionIndexes
	}
	return nil
}

func (m *Match) GetBlock() *BlockTrxPayload {
	if m != nil {
		return m.Block
	}
	return nil
}

type BlockTrxPayload struct {
	BlockID              string               `protobuf:"bytes,1,opt,name=blockID,proto3" json:"blockID,omitempty"`
	BlockHeader          *v1.BlockHeader      `protobuf:"bytes,2,opt,name=blockHeader,proto3" json:"blockHeader,omitempty"`
	Trace                *v1.TransactionTrace `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockTrxPayload) Reset()         { *m = BlockTrxPayload{} }
func (m *BlockTrxPayload) String() string { return proto.CompactTextString(m) }
func (*BlockTrxPayload) ProtoMessage()    {}
func (*BlockTrxPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6416b04c85aeead, []int{2}
}

func (m *BlockTrxPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockTrxPayload.Unmarshal(m, b)
}
func (m *BlockTrxPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockTrxPayload.Marshal(b, m, deterministic)
}
func (m *BlockTrxPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTrxPayload.Merge(m, src)
}
func (m *BlockTrxPayload) XXX_Size() int {
	return xxx_messageInfo_BlockTrxPayload.Size(m)
}
func (m *BlockTrxPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTrxPayload.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTrxPayload proto.InternalMessageInfo

func (m *BlockTrxPayload) GetBlockID() string {
	if m != nil {
		return m.BlockID
	}
	return ""
}

func (m *BlockTrxPayload) GetBlockHeader() *v1.BlockHeader {
	if m != nil {
		return m.BlockHeader
	}
	return nil
}

func (m *BlockTrxPayload) GetTrace() *v1.TransactionTrace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func init() {
	proto.RegisterType((*DocumentID)(nil), "dfuse.eosio.search.v1.DocumentID")
	proto.RegisterType((*Match)(nil), "dfuse.eosio.search.v1.Match")
	proto.RegisterType((*BlockTrxPayload)(nil), "dfuse.eosio.search.v1.BlockTrxPayload")
}

func init() {
	proto.RegisterFile("dfuse/eosio/search/v1/search.proto", fileDescriptor_f6416b04c85aeead)
}

var fileDescriptor_f6416b04c85aeead = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4f, 0x4f, 0xfa, 0x30,
	0x18, 0xce, 0x7e, 0xc0, 0x4f, 0x2d, 0x12, 0x4d, 0x8d, 0xc9, 0xc2, 0x69, 0x2e, 0xc6, 0x10, 0x0f,
	0xab, 0xe8, 0x51, 0x4f, 0xb8, 0x18, 0x38, 0x68, 0x48, 0xc3, 0xc9, 0x5b, 0xd7, 0xbd, 0xb0, 0x05,
	0x58, 0x97, 0xb6, 0xc3, 0xe1, 0x27, 0xf2, 0xe4, 0x67, 0x34, 0xeb, 0x00, 0x07, 0xee, 0xf6, 0xf6,
	0xf9, 0xd3, 0xf7, 0xc9, 0x93, 0x17, 0xb9, 0xe1, 0x34, 0x53, 0x40, 0x40, 0xa8, 0x58, 0x10, 0x05,
	0x4c, 0xf2, 0x88, 0xac, 0xfa, 0x9b, 0xc9, 0x4b, 0xa5, 0xd0, 0x02, 0x5f, 0x1a, 0x8d, 0x67, 0x34,
	0xde, 0x86, 0x59, 0xf5, 0xbb, 0x4e, 0xd5, 0xca, 0x45, 0x08, 0xbc, 0x70, 0x9a, 0xa1, 0x34, 0xba,
	0x5f, 0x16, 0x42, 0xbe, 0xe0, 0xd9, 0x12, 0x12, 0x3d, 0xf2, 0x71, 0x17, 0x1d, 0x07, 0x0b, 0xc1,
	0xe7, 0x6f, 0xd9, 0xd2, 0xb6, 0x1c, 0xab, 0xd7, 0xa4, 0xbb, 0x37, 0x76, 0x50, 0x9b, 0x71, 0x1d,
	0x8b, 0x64, 0x94, 0x84, 0x90, 0xdb, 0xff, 0x0c, 0x5d, 0x85, 0xf0, 0x2d, 0x3a, 0xd7, 0x92, 0x25,
	0xaa, 0x2a, 0x6b, 0x18, 0xd9, 0x1f, 0x1c, 0xdf, 0xa1, 0x8b, 0x2a, 0xe6, 0x8f, 0x25, 0x4c, 0xe3,
	0xdc, 0x6e, 0x3a, 0x56, 0xef, 0x94, 0xd6, 0x51, 0xee, 0x1c, 0xb5, 0x5e, 0x99, 0xe6, 0x11, 0xbe,
	0x46, 0x9d, 0xca, 0x4f, 0xa0, 0x6c, 0xcb, 0x69, 0xf4, 0x3a, 0x74, 0x1f, 0xc4, 0x4f, 0xa8, 0x65,
	0xa2, 0x9b, 0xa0, 0xed, 0xfb, 0x1b, 0xaf, 0xb6, 0x22, 0x6f, 0x50, 0x68, 0x26, 0x32, 0x1f, 0xb3,
	0xf5, 0x42, 0xb0, 0x90, 0x96, 0x26, 0xf7, 0xdb, 0x42, 0x67, 0x07, 0x14, 0xb6, 0xd1, 0x91, 0x21,
	0x47, 0xbe, 0xe9, 0xe6, 0x84, 0x6e, 0x9f, 0xf8, 0x19, 0xb5, 0xcd, 0x38, 0x04, 0x16, 0x82, 0xdc,
	0x6c, 0xbc, 0xda, 0xdb, 0x58, 0x96, 0xbe, 0x5d, 0x58, 0x0a, 0x69, 0xd5, 0x55, 0x04, 0xd6, 0x92,
	0x71, 0x30, 0x95, 0x1d, 0x06, 0xde, 0xd9, 0x27, 0xbf, 0xcd, 0x4c, 0x0a, 0x35, 0x2d, 0x4d, 0x83,
	0xe1, 0xfb, 0xcb, 0x2c, 0xd6, 0x51, 0x16, 0x78, 0x5c, 0x2c, 0xc9, 0x67, 0x24, 0x92, 0x99, 0x8a,
	0xb2, 0x0f, 0x48, 0x48, 0x14, 0x2b, 0x2d, 0xe4, 0x1a, 0xf2, 0x94, 0xa4, 0x01, 0xa9, 0x3d, 0xa6,
	0xc7, 0x34, 0x28, 0x67, 0x10, 0x2a, 0xf8, 0x6f, 0x2e, 0xe3, 0xe1, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x7a, 0x22, 0x23, 0xd0, 0x78, 0x02, 0x00, 0x00,
}
