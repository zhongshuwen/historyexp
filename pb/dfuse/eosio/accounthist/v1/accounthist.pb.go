// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: dfuse/eosio/accounthist/v1/accounthist.proto

package pbaccounthist

import (
	v1 "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/codec/v1"
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

type GetActionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account uint64  `protobuf:"varint,1,opt,name=account,proto3" json:"account,omitempty"`
	Limit   uint32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Cursor  *Cursor `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (x *GetActionsRequest) Reset() {
	*x = GetActionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActionsRequest) ProtoMessage() {}

func (x *GetActionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActionsRequest.ProtoReflect.Descriptor instead.
func (*GetActionsRequest) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescGZIP(), []int{0}
}

func (x *GetActionsRequest) GetAccount() uint64 {
	if x != nil {
		return x.Account
	}
	return 0
}

func (x *GetActionsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetActionsRequest) GetCursor() *Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor      *Cursor         `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	ActionTrace *v1.ActionTrace `protobuf:"bytes,2,opt,name=action_trace,json=actionTrace,proto3" json:"action_trace,omitempty"`
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescGZIP(), []int{1}
}

func (x *ActionResponse) GetCursor() *Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *ActionResponse) GetActionTrace() *v1.ActionTrace {
	if x != nil {
		return x.ActionTrace
	}
	return nil
}

type ActionRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        uint32          `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ActionTrace    *v1.ActionTrace `protobuf:"bytes,2,opt,name=action_trace,json=actionTrace,proto3" json:"action_trace,omitempty"`
	LastDeletedSeq uint64          `protobuf:"varint,3,opt,name=last_deleted_seq,json=lastDeletedSeq,proto3" json:"last_deleted_seq,omitempty"`
}

func (x *ActionRow) Reset() {
	*x = ActionRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRow) ProtoMessage() {}

func (x *ActionRow) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRow.ProtoReflect.Descriptor instead.
func (*ActionRow) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescGZIP(), []int{2}
}

func (x *ActionRow) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ActionRow) GetActionTrace() *v1.ActionTrace {
	if x != nil {
		return x.ActionTrace
	}
	return nil
}

func (x *ActionRow) GetLastDeletedSeq() uint64 {
	if x != nil {
		return x.LastDeletedSeq
	}
	return 0
}

type ShardCheckpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitialStartBlock   uint64 `protobuf:"varint,1,opt,name=initial_start_block,json=initialStartBlock,proto3" json:"initial_start_block,omitempty"`
	TargetStopBlock     uint64 `protobuf:"varint,2,opt,name=target_stop_block,json=targetStopBlock,proto3" json:"target_stop_block,omitempty"` // exclusive
	LastWrittenBlockNum uint64 `protobuf:"varint,3,opt,name=last_written_block_num,json=lastWrittenBlockNum,proto3" json:"last_written_block_num,omitempty"`
	LastWrittenBlockId  string `protobuf:"bytes,4,opt,name=last_written_block_id,json=lastWrittenBlockId,proto3" json:"last_written_block_id,omitempty"`
}

func (x *ShardCheckpoint) Reset() {
	*x = ShardCheckpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardCheckpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardCheckpoint) ProtoMessage() {}

func (x *ShardCheckpoint) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardCheckpoint.ProtoReflect.Descriptor instead.
func (*ShardCheckpoint) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescGZIP(), []int{3}
}

func (x *ShardCheckpoint) GetInitialStartBlock() uint64 {
	if x != nil {
		return x.InitialStartBlock
	}
	return 0
}

func (x *ShardCheckpoint) GetTargetStopBlock() uint64 {
	if x != nil {
		return x.TargetStopBlock
	}
	return 0
}

func (x *ShardCheckpoint) GetLastWrittenBlockNum() uint64 {
	if x != nil {
		return x.LastWrittenBlockNum
	}
	return 0
}

func (x *ShardCheckpoint) GetLastWrittenBlockId() string {
	if x != nil {
		return x.LastWrittenBlockId
	}
	return ""
}

type ActionRowAppend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastDeletedSeq uint64 `protobuf:"varint,3,opt,name=last_deleted_seq,json=lastDeletedSeq,proto3" json:"last_deleted_seq,omitempty"` // Align with `ActionRow.last_deleted_seq`
}

func (x *ActionRowAppend) Reset() {
	*x = ActionRowAppend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRowAppend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRowAppend) ProtoMessage() {}

func (x *ActionRowAppend) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRowAppend.ProtoReflect.Descriptor instead.
func (*ActionRowAppend) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescGZIP(), []int{4}
}

func (x *ActionRowAppend) GetLastDeletedSeq() uint64 {
	if x != nil {
		return x.LastDeletedSeq
	}
	return 0
}

type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version        uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Magic          uint32 `protobuf:"varint,2,opt,name=magic,proto3" json:"magic,omitempty"` // fixed at 4374 = acth in l33t = account history
	Account        uint64 `protobuf:"varint,3,opt,name=account,proto3" json:"account,omitempty"`
	ShardNum       uint32 `protobuf:"varint,4,opt,name=shard_num,json=shardNum,proto3" json:"shard_num,omitempty"`
	SequenceNumber uint64 `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescGZIP(), []int{5}
}

func (x *Cursor) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Cursor) GetMagic() uint32 {
	if x != nil {
		return x.Magic
	}
	return 0
}

func (x *Cursor) GetAccount() uint64 {
	if x != nil {
		return x.Account
	}
	return 0
}

func (x *Cursor) GetShardNum() uint32 {
	if x != nil {
		return x.ShardNum
	}
	return 0
}

func (x *Cursor) GetSequenceNumber() uint64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

var File_dfuse_eosio_accounthist_v1_accounthist_proto protoreflect.FileDescriptor

var file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2f, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x68, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x68, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x68, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x64, 0x66, 0x75, 0x73,
	0x65, 0x2f, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x3a, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x68, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x92, 0x01,
	0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x68, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x44, 0x0a, 0x0c,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x65, 0x6f, 0x73, 0x69, 0x6f,
	0x2e, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x77,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0c, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x53, 0x65, 0x71, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x53,
	0x68, 0x61, 0x72, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2e,
	0x0a, 0x13, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2a,
	0x0a, 0x11, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x53, 0x74, 0x6f, 0x70, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x33, 0x0a, 0x16, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74,
	0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x12,
	0x31, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x6c, 0x61, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x64, 0x22, 0x3b, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x77, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x53, 0x65, 0x71, 0x22,
	0x98, 0x01, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x4e, 0x75,
	0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x7b, 0x0a, 0x0e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x69, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x2e, 0x64, 0x66, 0x75,
	0x73, 0x65, 0x2e, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x68, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x64, 0x66, 0x75, 0x73,
	0x65, 0x2e, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x68,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x6f, 0x6e, 0x67, 0x73, 0x68, 0x75, 0x77, 0x65,
	0x6e, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x65, 0x78, 0x70, 0x2f, 0x70, 0x62, 0x2f,
	0x64, 0x66, 0x75, 0x73, 0x65, 0x2f, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x68, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x68, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescOnce sync.Once
	file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescData = file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDesc
)

func file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescGZIP() []byte {
	file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescOnce.Do(func() {
		file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescData = protoimpl.X.CompressGZIP(file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescData)
	})
	return file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDescData
}

var file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dfuse_eosio_accounthist_v1_accounthist_proto_goTypes = []interface{}{
	(*GetActionsRequest)(nil), // 0: dfuse.eosio.accounthist.v1.GetActionsRequest
	(*ActionResponse)(nil),    // 1: dfuse.eosio.accounthist.v1.ActionResponse
	(*ActionRow)(nil),         // 2: dfuse.eosio.accounthist.v1.ActionRow
	(*ShardCheckpoint)(nil),   // 3: dfuse.eosio.accounthist.v1.ShardCheckpoint
	(*ActionRowAppend)(nil),   // 4: dfuse.eosio.accounthist.v1.ActionRowAppend
	(*Cursor)(nil),            // 5: dfuse.eosio.accounthist.v1.Cursor
	(*v1.ActionTrace)(nil),    // 6: dfuse.eosio.codec.v1.ActionTrace
}
var file_dfuse_eosio_accounthist_v1_accounthist_proto_depIdxs = []int32{
	5, // 0: dfuse.eosio.accounthist.v1.GetActionsRequest.cursor:type_name -> dfuse.eosio.accounthist.v1.Cursor
	5, // 1: dfuse.eosio.accounthist.v1.ActionResponse.cursor:type_name -> dfuse.eosio.accounthist.v1.Cursor
	6, // 2: dfuse.eosio.accounthist.v1.ActionResponse.action_trace:type_name -> dfuse.eosio.codec.v1.ActionTrace
	6, // 3: dfuse.eosio.accounthist.v1.ActionRow.action_trace:type_name -> dfuse.eosio.codec.v1.ActionTrace
	0, // 4: dfuse.eosio.accounthist.v1.AccountHistory.GetActions:input_type -> dfuse.eosio.accounthist.v1.GetActionsRequest
	1, // 5: dfuse.eosio.accounthist.v1.AccountHistory.GetActions:output_type -> dfuse.eosio.accounthist.v1.ActionResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_dfuse_eosio_accounthist_v1_accounthist_proto_init() }
func file_dfuse_eosio_accounthist_v1_accounthist_proto_init() {
	if File_dfuse_eosio_accounthist_v1_accounthist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActionsRequest); i {
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
		file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResponse); i {
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
		file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRow); i {
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
		file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardCheckpoint); i {
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
		file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRowAppend); i {
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
		file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
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
			RawDescriptor: file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dfuse_eosio_accounthist_v1_accounthist_proto_goTypes,
		DependencyIndexes: file_dfuse_eosio_accounthist_v1_accounthist_proto_depIdxs,
		MessageInfos:      file_dfuse_eosio_accounthist_v1_accounthist_proto_msgTypes,
	}.Build()
	File_dfuse_eosio_accounthist_v1_accounthist_proto = out.File
	file_dfuse_eosio_accounthist_v1_accounthist_proto_rawDesc = nil
	file_dfuse_eosio_accounthist_v1_accounthist_proto_goTypes = nil
	file_dfuse_eosio_accounthist_v1_accounthist_proto_depIdxs = nil
}
