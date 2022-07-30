// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: dfuse/eosio/statedb/v1/tablet.proto

package pbstatedb

import (
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

type AuthLinkValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permission uint64 `protobuf:"varint,1,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *AuthLinkValue) Reset() {
	*x = AuthLinkValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthLinkValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthLinkValue) ProtoMessage() {}

func (x *AuthLinkValue) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthLinkValue.ProtoReflect.Descriptor instead.
func (*AuthLinkValue) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_statedb_v1_tablet_proto_rawDescGZIP(), []int{0}
}

func (x *AuthLinkValue) GetPermission() uint64 {
	if x != nil {
		return x.Permission
	}
	return 0
}

type ContractStateValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payer uint64 `protobuf:"varint,1,opt,name=payer,proto3" json:"payer,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ContractStateValue) Reset() {
	*x = ContractStateValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractStateValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractStateValue) ProtoMessage() {}

func (x *ContractStateValue) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractStateValue.ProtoReflect.Descriptor instead.
func (*ContractStateValue) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_statedb_v1_tablet_proto_rawDescGZIP(), []int{1}
}

func (x *ContractStateValue) GetPayer() uint64 {
	if x != nil {
		return x.Payer
	}
	return 0
}

func (x *ContractStateValue) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ContractTableScopeValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payer uint64 `protobuf:"varint,1,opt,name=payer,proto3" json:"payer,omitempty"`
}

func (x *ContractTableScopeValue) Reset() {
	*x = ContractTableScopeValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractTableScopeValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractTableScopeValue) ProtoMessage() {}

func (x *ContractTableScopeValue) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractTableScopeValue.ProtoReflect.Descriptor instead.
func (*ContractTableScopeValue) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_statedb_v1_tablet_proto_rawDescGZIP(), []int{2}
}

func (x *ContractTableScopeValue) GetPayer() uint64 {
	if x != nil {
		return x.Payer
	}
	return 0
}

// KeyAccountValue is actual empty and the bool field is there mainly to permit a future
// extension where we could add more fields in there.
type KeyAccountValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Present bool `protobuf:"varint,1,opt,name=present,proto3" json:"present,omitempty"`
}

func (x *KeyAccountValue) Reset() {
	*x = KeyAccountValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyAccountValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyAccountValue) ProtoMessage() {}

func (x *KeyAccountValue) ProtoReflect() protoreflect.Message {
	mi := &file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyAccountValue.ProtoReflect.Descriptor instead.
func (*KeyAccountValue) Descriptor() ([]byte, []int) {
	return file_dfuse_eosio_statedb_v1_tablet_proto_rawDescGZIP(), []int{3}
}

func (x *KeyAccountValue) GetPresent() bool {
	if x != nil {
		return x.Present
	}
	return false
}

var File_dfuse_eosio_statedb_v1_tablet_proto protoreflect.FileDescriptor

var file_dfuse_eosio_statedb_v1_tablet_proto_rawDesc = []byte{
	0x0a, 0x23, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2f, 0x65, 0x6f, 0x73, 0x69, 0x6f, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2e, 0x65, 0x6f, 0x73,
	0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x22, 0x2f, 0x0a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x69, 0x6e, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3e,
	0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2f,
	0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x22,
	0x2b, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x42, 0x47, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x6f, 0x6e, 0x67,
	0x73, 0x68, 0x75, 0x77, 0x65, 0x6e, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x65, 0x78,
	0x70, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x66, 0x75, 0x73, 0x65, 0x2f, 0x65, 0x6f, 0x73, 0x69, 0x6f,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dfuse_eosio_statedb_v1_tablet_proto_rawDescOnce sync.Once
	file_dfuse_eosio_statedb_v1_tablet_proto_rawDescData = file_dfuse_eosio_statedb_v1_tablet_proto_rawDesc
)

func file_dfuse_eosio_statedb_v1_tablet_proto_rawDescGZIP() []byte {
	file_dfuse_eosio_statedb_v1_tablet_proto_rawDescOnce.Do(func() {
		file_dfuse_eosio_statedb_v1_tablet_proto_rawDescData = protoimpl.X.CompressGZIP(file_dfuse_eosio_statedb_v1_tablet_proto_rawDescData)
	})
	return file_dfuse_eosio_statedb_v1_tablet_proto_rawDescData
}

var file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dfuse_eosio_statedb_v1_tablet_proto_goTypes = []interface{}{
	(*AuthLinkValue)(nil),           // 0: dfuse.eosio.statedb.v1.AuthLinkValue
	(*ContractStateValue)(nil),      // 1: dfuse.eosio.statedb.v1.ContractStateValue
	(*ContractTableScopeValue)(nil), // 2: dfuse.eosio.statedb.v1.ContractTableScopeValue
	(*KeyAccountValue)(nil),         // 3: dfuse.eosio.statedb.v1.KeyAccountValue
}
var file_dfuse_eosio_statedb_v1_tablet_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dfuse_eosio_statedb_v1_tablet_proto_init() }
func file_dfuse_eosio_statedb_v1_tablet_proto_init() {
	if File_dfuse_eosio_statedb_v1_tablet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthLinkValue); i {
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
		file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractStateValue); i {
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
		file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractTableScopeValue); i {
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
		file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyAccountValue); i {
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
			RawDescriptor: file_dfuse_eosio_statedb_v1_tablet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dfuse_eosio_statedb_v1_tablet_proto_goTypes,
		DependencyIndexes: file_dfuse_eosio_statedb_v1_tablet_proto_depIdxs,
		MessageInfos:      file_dfuse_eosio_statedb_v1_tablet_proto_msgTypes,
	}.Build()
	File_dfuse_eosio_statedb_v1_tablet_proto = out.File
	file_dfuse_eosio_statedb_v1_tablet_proto_rawDesc = nil
	file_dfuse_eosio_statedb_v1_tablet_proto_goTypes = nil
	file_dfuse_eosio_statedb_v1_tablet_proto_depIdxs = nil
}
