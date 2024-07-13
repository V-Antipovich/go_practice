// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: psql/accounts/accounts.proto

package accounts

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_psql_accounts_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_psql_accounts_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_psql_accounts_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ChangeAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Newname string `protobuf:"bytes,2,opt,name=newname,proto3" json:"newname,omitempty"`
}

func (x *ChangeAccount) Reset() {
	*x = ChangeAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_psql_accounts_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAccount) ProtoMessage() {}

func (x *ChangeAccount) ProtoReflect() protoreflect.Message {
	mi := &file_psql_accounts_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAccount.ProtoReflect.Descriptor instead.
func (*ChangeAccount) Descriptor() ([]byte, []int) {
	return file_psql_accounts_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChangeAccount) GetNewname() string {
	if x != nil {
		return x.Newname
	}
	return ""
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_psql_accounts_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_psql_accounts_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_psql_accounts_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_psql_accounts_accounts_proto protoreflect.FileDescriptor

var file_psql_accounts_accounts_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x73, 0x71, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3d, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x77, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x32, 0xb9, 0x01, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x22, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x08, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x12, 0x1f, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x05, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x1a, 0x08, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x0e, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a,
	0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0c, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x08, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x1a, 0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x05, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_psql_accounts_accounts_proto_rawDescOnce sync.Once
	file_psql_accounts_accounts_proto_rawDescData = file_psql_accounts_accounts_proto_rawDesc
)

func file_psql_accounts_accounts_proto_rawDescGZIP() []byte {
	file_psql_accounts_accounts_proto_rawDescOnce.Do(func() {
		file_psql_accounts_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_psql_accounts_accounts_proto_rawDescData)
	})
	return file_psql_accounts_accounts_proto_rawDescData
}

var file_psql_accounts_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_psql_accounts_accounts_proto_goTypes = []any{
	(*Account)(nil),       // 0: Account
	(*ChangeAccount)(nil), // 1: ChangeAccount
	(*Name)(nil),          // 2: Name
}
var file_psql_accounts_accounts_proto_depIdxs = []int32{
	0, // 0: Bank.CreateAccount:input_type -> Account
	2, // 1: Bank.GetAccount:input_type -> Name
	1, // 2: Bank.UpdateAccount:input_type -> ChangeAccount
	0, // 3: Bank.PatchAccount:input_type -> Account
	2, // 4: Bank.DeleteAccount:input_type -> Name
	2, // 5: Bank.CreateAccount:output_type -> Name
	0, // 6: Bank.GetAccount:output_type -> Account
	2, // 7: Bank.UpdateAccount:output_type -> Name
	2, // 8: Bank.PatchAccount:output_type -> Name
	2, // 9: Bank.DeleteAccount:output_type -> Name
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_psql_accounts_accounts_proto_init() }
func file_psql_accounts_accounts_proto_init() {
	if File_psql_accounts_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_psql_accounts_accounts_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Account); i {
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
		file_psql_accounts_accounts_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ChangeAccount); i {
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
		file_psql_accounts_accounts_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Name); i {
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
			RawDescriptor: file_psql_accounts_accounts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_psql_accounts_accounts_proto_goTypes,
		DependencyIndexes: file_psql_accounts_accounts_proto_depIdxs,
		MessageInfos:      file_psql_accounts_accounts_proto_msgTypes,
	}.Build()
	File_psql_accounts_accounts_proto = out.File
	file_psql_accounts_accounts_proto_rawDesc = nil
	file_psql_accounts_accounts_proto_goTypes = nil
	file_psql_accounts_accounts_proto_depIdxs = nil
}
