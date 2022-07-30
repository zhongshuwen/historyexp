// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: dfuse/eosio/search/v1/search.proto

package pbsearcheos

import (
	v1 "github.com/zhongshuwen/histexp/pb/dfuse/eosio/codec/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DocumentID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNum            uint64 `protobuf:"varint,1,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	ActionIndex         uint64 `protobuf:"varint,2,opt,name=actionIndex,proto3" json:"actionIndex,omitempty"`
	TransactionIndex    uint64 `protobuf:"varint,3,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	TransactionIDPrefix []byte `protobuf:"bytes,4,opt,name=transactionIDPrefix,proto3" json:"transactionIDPrefix,omitempty"`
}

func (x *DocumentID) Reset() {
	*x = DocumentID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_search_v1_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentID) ProtoMessage() {}

func (x *DocumentID) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_search_v1_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentID.ProtoReflect.Descriptor instead.
func (*DocumentID) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_search_v1_search_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentID) GetBlockNum() uint64 {
	if x != nil {
		return x.BlockNum
	}
	return 0
}

func (x *DocumentID) GetActionIndex() uint64 {
	if x != nil {
		return x.ActionIndex
	}
	return 0
}

func (x *DocumentID) GetTransactionIndex() uint64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *DocumentID) GetTransactionIDPrefix() []byte {
	if x != nil {
		return x.TransactionIDPrefix
	}
	return nil
}

type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionIndexes []uint32         `protobuf:"varint,1,rep,packed,name=actionIndexes,proto3" json:"actionIndexes,omitempty"`
	Block         *BlockTrxPayload `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_search_v1_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_search_v1_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_search_v1_search_proto_rawDescGZIP(), []int{1}
}

func (x *Match) GetActionIndexes() []uint32 {
	if x != nil {
		return x.ActionIndexes
	}
	return nil
}

func (x *Match) GetBlock() *BlockTrxPayload {
	if x != nil {
		return x.Block
	}
	return nil
}

type BlockTrxPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockID     string               `protobuf:"bytes,1,opt,name=blockID,proto3" json:"blockID,omitempty"`
	BlockHeader *v1.BlockHeader      `protobuf:"bytes,2,opt,name=blockHeader,proto3" json:"blockHeader,omitempty"`
	Trace       *v1.TransactionTrace `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
}

func (x *BlockTrxPayload) Reset() {
	*x = BlockTrxPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_search_v1_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockTrxPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockTrxPayload) ProtoMessage() {}

func (x *BlockTrxPayload) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_search_v1_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockTrxPayload.ProtoReflect.Descriptor instead.
func (*BlockTrxPayload) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_search_v1_search_proto_rawDescGZIP(), []int{2}
}

func (x *BlockTrxPayload) GetBlockID() string {
	if x != nil {
		return x.BlockID
	}
	return ""
}

func (x *BlockTrxPayload) GetBlockHeader() *v1.BlockHeader {
	if x != nil {
		return x.BlockHeader
	}
	return nil
}

func (x *BlockTrxPayload) GetTrace() *v1.TransactionTrace {
	if x != nil {
		return x.Trace
	}
	return nil
}

var File_dfuse_eosio_search_v1_search_proto protoreflect.FileDescriptor

var file_dfuse_eosio_search_v1_search_proto_rawDesc = []byte{
	0x0a, 0x22, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2f, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x65, 0x6f, 0x73, 0x69,
	0x6f, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x64, 0x66, 0x75,
	0x73, 0x65, 0x2f, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01,
	0x0a, 0x0a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x30, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x6b, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x65,
	0x6f, 0x73, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x72, 0x78, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xae, 0x01, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54,
	0x72, 0x78, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x44, 0x12, 0x43, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65,
	0x2e, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e,
	0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52,
	0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x6f, 0x6e, 0x67, 0x73, 0x68, 0x75, 0x77, 0x65, 0x6e,
	0x2f, 0x68, 0x69, 0x73, 0x74, 0x65, 0x78, 0x70, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x66, 0x75, 0x73,
	0x65, 0x2f, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x62, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dfuse_eosio_search_v1_search_proto_rawDescOnce sync.Once
	file_dfuse_eosio_search_v1_search_proto_rawDescData = file_dfuse_eosio_search_v1_search_proto_rawDesc
)

func file_dfuse_eosio_search_v1_search_proto_rawDescGZIP() []byte {
	file_dfuse_eosio_search_v1_search_proto_rawDescOnce.Do(func() {
		file_dfuse_eosio_search_v1_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_dfuse_eosio_search_v1_search_proto_rawDescData)
	})
	return file_dfuse_eosio_search_v1_search_proto_rawDescData
}

var file_dfuse_eosio_search_v1_search_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dfuse_eosio_search_v1_search_proto_goTypes = []interface{}{
	(*DocumentID)(nil),          // 0: dfuse.eosio.search.v1.DocumentID
	(*Match)(nil),               // 1: dfuse.eosio.search.v1.Match
	(*BlockTrxPayload)(nil),     // 2: dfuse.eosio.search.v1.BlockTrxPayload
	(*v1.BlockHeader)(nil),      // 3: dfuse.eosio.codec.v1.BlockHeader
	(*v1.TransactionTrace)(nil), // 4: dfuse.eosio.codec.v1.TransactionTrace
}
var file_dfuse_eosio_search_v1_search_proto_depIdxs = []int32{
	2, // 0: dfuse.eosio.search.v1.Match.block:type_name -> dfuse.eosio.search.v1.BlockTrxPayload
	3, // 1: dfuse.eosio.search.v1.BlockTrxPayload.blockHeader:type_name -> dfuse.eosio.codec.v1.BlockHeader
	4, // 2: dfuse.eosio.search.v1.BlockTrxPayload.trace:type_name -> dfuse.eosio.codec.v1.TransactionTrace
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dfuse_eosio_search_v1_search_proto_init() }
func file_dfuse_eosio_search_v1_search_proto_init() {
	if File_dfuse_eosio_search_v1_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dfuse_eosio_search_v1_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dfuse_eosio_search_v1_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dfuse_eosio_search_v1_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockTrxPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dfuse_eosio_search_v1_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dfuse_eosio_search_v1_search_proto_goTypes,
		DependencyIndexes: file_dfuse_eosio_search_v1_search_proto_depIdxs,
		MessageInfos:      file_dfuse_eosio_search_v1_search_proto_msgTypes,
	}.Build()
	File_dfuse_eosio_search_v1_search_proto = out.File
	file_dfuse_eosio_search_v1_search_proto_rawDesc = nil
	file_dfuse_eosio_search_v1_search_proto_goTypes = nil
	file_dfuse_eosio_search_v1_search_proto_depIdxs = nil
}
