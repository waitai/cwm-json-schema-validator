// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: cisco.jsonschema.common.adapter.proto

package common

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

type ErrorCode int32

const (
	ErrorCode_DEFAULT ErrorCode = 0
	ErrorCode_INIT    ErrorCode = 1
	ErrorCode_FATAL   ErrorCode = 2
	ErrorCode_TIMEOUT ErrorCode = 3
	ErrorCode_PANIC   ErrorCode = 4
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "DEFAULT",
		1: "INIT",
		2: "FATAL",
		3: "TIMEOUT",
		4: "PANIC",
	}
	ErrorCode_value = map[string]int32{
		"DEFAULT": 0,
		"INIT":    1,
		"FATAL":   2,
		"TIMEOUT": 3,
		"PANIC":   4,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_cisco_jsonschema_common_adapter_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_cisco_jsonschema_common_adapter_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_cisco_jsonschema_common_adapter_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cisco_jsonschema_common_adapter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_cisco_jsonschema_common_adapter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_cisco_jsonschema_common_adapter_proto_rawDescGZIP(), []int{0}
}

var File_cisco_jsonschema_common_adapter_proto protoreflect.FileDescriptor

var file_cisco_jsonschema_common_adapter_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x2e, 0x6a,
	0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x22, 0x08, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0x45, 0x0a, 0x09, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d,
	0x45, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10,
	0x04, 0x42, 0x30, 0x5a, 0x2e, 0x63, 0x69, 0x73, 0x63, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x77, 0x6d, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x69, 0x73, 0x63,
	0x6f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cisco_jsonschema_common_adapter_proto_rawDescOnce sync.Once
	file_cisco_jsonschema_common_adapter_proto_rawDescData = file_cisco_jsonschema_common_adapter_proto_rawDesc
)

func file_cisco_jsonschema_common_adapter_proto_rawDescGZIP() []byte {
	file_cisco_jsonschema_common_adapter_proto_rawDescOnce.Do(func() {
		file_cisco_jsonschema_common_adapter_proto_rawDescData = protoimpl.X.CompressGZIP(file_cisco_jsonschema_common_adapter_proto_rawDescData)
	})
	return file_cisco_jsonschema_common_adapter_proto_rawDescData
}

var file_cisco_jsonschema_common_adapter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cisco_jsonschema_common_adapter_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cisco_jsonschema_common_adapter_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: cisco.jsonschema.common.ErrorCode
	(*Config)(nil), // 1: cisco.jsonschema.common.Config
}
var file_cisco_jsonschema_common_adapter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cisco_jsonschema_common_adapter_proto_init() }
func file_cisco_jsonschema_common_adapter_proto_init() {
	if File_cisco_jsonschema_common_adapter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cisco_jsonschema_common_adapter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_cisco_jsonschema_common_adapter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cisco_jsonschema_common_adapter_proto_goTypes,
		DependencyIndexes: file_cisco_jsonschema_common_adapter_proto_depIdxs,
		EnumInfos:         file_cisco_jsonschema_common_adapter_proto_enumTypes,
		MessageInfos:      file_cisco_jsonschema_common_adapter_proto_msgTypes,
	}.Build()
	File_cisco_jsonschema_common_adapter_proto = out.File
	file_cisco_jsonschema_common_adapter_proto_rawDesc = nil
	file_cisco_jsonschema_common_adapter_proto_goTypes = nil
	file_cisco_jsonschema_common_adapter_proto_depIdxs = nil
}
