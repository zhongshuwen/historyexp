// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/eosio/funnel/v1/funnel.proto

package pbfunnel

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/codec/v1"
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

type StreamBlockRequest struct {
	FromBlockNum         int64    `protobuf:"varint,2,opt,name=fromBlockNum,proto3" json:"fromBlockNum,omitempty"`
	IrreversibleOnly     bool     `protobuf:"varint,3,opt,name=irreversibleOnly,proto3" json:"irreversibleOnly,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamBlockRequest) Reset()         { *m = StreamBlockRequest{} }
func (m *StreamBlockRequest) String() string { return proto.CompactTextString(m) }
func (*StreamBlockRequest) ProtoMessage()    {}
func (*StreamBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_477a415e0c40d59c, []int{0}
}

func (m *StreamBlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamBlockRequest.Unmarshal(m, b)
}
func (m *StreamBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamBlockRequest.Marshal(b, m, deterministic)
}
func (m *StreamBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamBlockRequest.Merge(m, src)
}
func (m *StreamBlockRequest) XXX_Size() int {
	return xxx_messageInfo_StreamBlockRequest.Size(m)
}
func (m *StreamBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamBlockRequest proto.InternalMessageInfo

func (m *StreamBlockRequest) GetFromBlockNum() int64 {
	if m != nil {
		return m.FromBlockNum
	}
	return 0
}

func (m *StreamBlockRequest) GetIrreversibleOnly() bool {
	if m != nil {
		return m.IrreversibleOnly
	}
	return false
}

type StreamBlockResponse struct {
	Undo                 bool      `protobuf:"varint,1,opt,name=undo,proto3" json:"undo,omitempty"`
	Block                *v1.Block `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StreamBlockResponse) Reset()         { *m = StreamBlockResponse{} }
func (m *StreamBlockResponse) String() string { return proto.CompactTextString(m) }
func (*StreamBlockResponse) ProtoMessage()    {}
func (*StreamBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_477a415e0c40d59c, []int{1}
}

func (m *StreamBlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamBlockResponse.Unmarshal(m, b)
}
func (m *StreamBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamBlockResponse.Marshal(b, m, deterministic)
}
func (m *StreamBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamBlockResponse.Merge(m, src)
}
func (m *StreamBlockResponse) XXX_Size() int {
	return xxx_messageInfo_StreamBlockResponse.Size(m)
}
func (m *StreamBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamBlockResponse proto.InternalMessageInfo

func (m *StreamBlockResponse) GetUndo() bool {
	if m != nil {
		return m.Undo
	}
	return false
}

func (m *StreamBlockResponse) GetBlock() *v1.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamBlockRequest)(nil), "dfuse.eosio.funnel.v1.StreamBlockRequest")
	proto.RegisterType((*StreamBlockResponse)(nil), "dfuse.eosio.funnel.v1.StreamBlockResponse")
}

func init() {
	proto.RegisterFile("dfuse/eosio/funnel/v1/funnel.proto", fileDescriptor_477a415e0c40d59c)
}

var fileDescriptor_477a415e0c40d59c = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x3d, 0x4f, 0xc3, 0x30,
	0x14, 0x54, 0x28, 0x54, 0xc8, 0x74, 0x40, 0x46, 0x48, 0x55, 0x59, 0xa2, 0x4c, 0xa5, 0x83, 0x4d,
	0xca, 0xc8, 0x56, 0x04, 0x23, 0x48, 0x61, 0x43, 0x2c, 0x38, 0x79, 0xf9, 0x10, 0x89, 0x9f, 0xeb,
	0x8f, 0x40, 0xf9, 0xf5, 0xa8, 0x76, 0x87, 0x54, 0xed, 0xc0, 0x76, 0x3a, 0xdf, 0xf9, 0xde, 0xbd,
	0x47, 0x92, 0xa2, 0x74, 0x06, 0x38, 0xa0, 0x69, 0x90, 0x97, 0x4e, 0x4a, 0x68, 0x79, 0x9f, 0xee,
	0x10, 0x53, 0x1a, 0x2d, 0xd2, 0x6b, 0xaf, 0x61, 0x5e, 0xc3, 0x76, 0x2f, 0x7d, 0x3a, 0x8b, 0x87,
	0xd6, 0x1c, 0x0b, 0xc8, 0xb7, 0x4e, 0x0f, 0x82, 0x31, 0x29, 0x08, 0x7d, 0xb3, 0x1a, 0x3e, 0xbb,
	0x55, 0x8b, 0xf9, 0x57, 0x06, 0x6b, 0x07, 0xc6, 0xd2, 0x84, 0x4c, 0x4a, 0x8d, 0x81, 0x7b, 0x71,
	0xdd, 0xf4, 0x24, 0x8e, 0xe6, 0xa3, 0x6c, 0x8f, 0xa3, 0x0b, 0x72, 0xd9, 0x68, 0x0d, 0x3d, 0x68,
	0xd3, 0x88, 0x16, 0x5e, 0x65, 0xbb, 0x99, 0x8e, 0xe2, 0x68, 0x7e, 0x9e, 0x1d, 0xf0, 0xc9, 0x07,
	0xb9, 0xda, 0x4b, 0x31, 0x0a, 0xa5, 0x01, 0x4a, 0xc9, 0xa9, 0x93, 0x05, 0x4e, 0x23, 0x6f, 0xf3,
	0x98, 0xa6, 0xe4, 0x4c, 0x6c, 0x45, 0x3e, 0xf3, 0x62, 0x79, 0xc3, 0x86, 0xcd, 0xc2, 0xe4, 0x7d,
	0xca, 0xc2, 0x3f, 0x41, 0xb9, 0x5c, 0x93, 0xf1, 0xb3, 0xaf, 0x4c, 0x2b, 0x32, 0x19, 0xe4, 0x18,
	0x7a, 0xcb, 0x8e, 0xee, 0x85, 0x1d, 0x56, 0x9e, 0x2d, 0xfe, 0x23, 0x0d, 0x73, 0xdf, 0x45, 0xab,
	0xa7, 0xf7, 0xc7, 0xaa, 0xb1, 0xb5, 0x13, 0x2c, 0xc7, 0x8e, 0xff, 0xd6, 0x28, 0x2b, 0x53, 0xbb,
	0x6f, 0x90, 0xbc, 0x6e, 0x8c, 0x45, 0xbd, 0x81, 0x1f, 0xc5, 0x95, 0xe0, 0x47, 0x4f, 0xf7, 0xa0,
	0x44, 0xc0, 0x62, 0xec, 0x8f, 0x70, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x15, 0xc8, 0xdc, 0xf9,
	0xe3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FunnelClient is the client API for Funnel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FunnelClient interface {
	StreamBlocks(ctx context.Context, in *StreamBlockRequest, opts ...grpc.CallOption) (Funnel_StreamBlocksClient, error)
}

type funnelClient struct {
	cc grpc.ClientConnInterface
}

func NewFunnelClient(cc grpc.ClientConnInterface) FunnelClient {
	return &funnelClient{cc}
}

func (c *funnelClient) StreamBlocks(ctx context.Context, in *StreamBlockRequest, opts ...grpc.CallOption) (Funnel_StreamBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Funnel_serviceDesc.Streams[0], "/dfuse.eosio.funnel.v1.Funnel/StreamBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &funnelStreamBlocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Funnel_StreamBlocksClient interface {
	Recv() (*StreamBlockResponse, error)
	grpc.ClientStream
}

type funnelStreamBlocksClient struct {
	grpc.ClientStream
}

func (x *funnelStreamBlocksClient) Recv() (*StreamBlockResponse, error) {
	m := new(StreamBlockResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FunnelServer is the server API for Funnel service.
type FunnelServer interface {
	StreamBlocks(*StreamBlockRequest, Funnel_StreamBlocksServer) error
}

// UnimplementedFunnelServer can be embedded to have forward compatible implementations.
type UnimplementedFunnelServer struct {
}

func (*UnimplementedFunnelServer) StreamBlocks(req *StreamBlockRequest, srv Funnel_StreamBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamBlocks not implemented")
}

func RegisterFunnelServer(s *grpc.Server, srv FunnelServer) {
	s.RegisterService(&_Funnel_serviceDesc, srv)
}

func _Funnel_StreamBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamBlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FunnelServer).StreamBlocks(m, &funnelStreamBlocksServer{stream})
}

type Funnel_StreamBlocksServer interface {
	Send(*StreamBlockResponse) error
	grpc.ServerStream
}

type funnelStreamBlocksServer struct {
	grpc.ServerStream
}

func (x *funnelStreamBlocksServer) Send(m *StreamBlockResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Funnel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.eosio.funnel.v1.Funnel",
	HandlerType: (*FunnelServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamBlocks",
			Handler:       _Funnel_StreamBlocks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/eosio/funnel/v1/funnel.proto",
}
