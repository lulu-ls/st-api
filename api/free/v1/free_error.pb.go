// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: free/v1/free_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	// 为某个枚举单独设置错误码
	ErrorReason_USER_NOT_FOUND                 ErrorReason = 0
	ErrorReason_CONTENT_MISSING                ErrorReason = 1
	ErrorReason_TOKEN_VERIFY                   ErrorReason = 2
	ErrorReason_AUTH_CHANNEL_APPID_IS_REQUIRED ErrorReason = 10
	ErrorReason_FREE_SIGNUP_ASSET_NOT_ENOUGH   ErrorReason = 11 // 财产余额不足
	ErrorReason_FREE_SIGNUP_ASSET_MIN_LIMIT    ErrorReason = 12 // 未达到最低限制
	ErrorReason_FREE_SIGNUP_ASSET_MAX_LIMIT    ErrorReason = 13 // 未达到最低限制
	ErrorReason_FREE_SIGNUP_ASSET_GET          ErrorReason = 14 // 获取账户信息失败
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:  "USER_NOT_FOUND",
		1:  "CONTENT_MISSING",
		2:  "TOKEN_VERIFY",
		10: "AUTH_CHANNEL_APPID_IS_REQUIRED",
		11: "FREE_SIGNUP_ASSET_NOT_ENOUGH",
		12: "FREE_SIGNUP_ASSET_MIN_LIMIT",
		13: "FREE_SIGNUP_ASSET_MAX_LIMIT",
		14: "FREE_SIGNUP_ASSET_GET",
	}
	ErrorReason_value = map[string]int32{
		"USER_NOT_FOUND":                 0,
		"CONTENT_MISSING":                1,
		"TOKEN_VERIFY":                   2,
		"AUTH_CHANNEL_APPID_IS_REQUIRED": 10,
		"FREE_SIGNUP_ASSET_NOT_ENOUGH":   11,
		"FREE_SIGNUP_ASSET_MIN_LIMIT":    12,
		"FREE_SIGNUP_ASSET_MAX_LIMIT":    13,
		"FREE_SIGNUP_ASSET_GET":          14,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_free_v1_free_error_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_free_v1_free_error_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_free_v1_free_error_proto_rawDescGZIP(), []int{0}
}

var File_free_v1_free_error_proto protoreflect.FileDescriptor

var file_free_v1_free_error_proto_rawDesc = []byte{
	0x0a, 0x18, 0x66, 0x72, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x72, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa1, 0x02, 0x0a,
	0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x0e,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00,
	0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e,
	0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90,
	0x03, 0x12, 0x16, 0x0a, 0x0c, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x59, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x28, 0x0a, 0x1e, 0x41, 0x55, 0x54,
	0x48, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x41, 0x50, 0x50, 0x49, 0x44, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0a, 0x1a, 0x04, 0xa8,
	0x45, 0x90, 0x03, 0x12, 0x26, 0x0a, 0x1c, 0x46, 0x52, 0x45, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e,
	0x55, 0x50, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f,
	0x55, 0x47, 0x48, 0x10, 0x0b, 0x1a, 0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x25, 0x0a, 0x1b, 0x46,
	0x52, 0x45, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54,
	0x5f, 0x4d, 0x49, 0x4e, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x0c, 0x1a, 0x04, 0xa8, 0x45,
	0x9b, 0x03, 0x12, 0x25, 0x0a, 0x1b, 0x46, 0x52, 0x45, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55,
	0x50, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x4c, 0x49, 0x4d, 0x49,
	0x54, 0x10, 0x0d, 0x1a, 0x04, 0xa8, 0x45, 0x9c, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x46, 0x52, 0x45,
	0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x47,
	0x45, 0x54, 0x10, 0x0e, 0x1a, 0x04, 0xa8, 0x45, 0x9d, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03,
	0x42, 0x39, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75,
	0x6c, 0x75, 0x2d, 0x6c, 0x73, 0x2f, 0x73, 0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_free_v1_free_error_proto_rawDescOnce sync.Once
	file_free_v1_free_error_proto_rawDescData = file_free_v1_free_error_proto_rawDesc
)

func file_free_v1_free_error_proto_rawDescGZIP() []byte {
	file_free_v1_free_error_proto_rawDescOnce.Do(func() {
		file_free_v1_free_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_free_v1_free_error_proto_rawDescData)
	})
	return file_free_v1_free_error_proto_rawDescData
}

var file_free_v1_free_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_free_v1_free_error_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: api.free.v1.ErrorReason
}
var file_free_v1_free_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_free_v1_free_error_proto_init() }
func file_free_v1_free_error_proto_init() {
	if File_free_v1_free_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_free_v1_free_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_free_v1_free_error_proto_goTypes,
		DependencyIndexes: file_free_v1_free_error_proto_depIdxs,
		EnumInfos:         file_free_v1_free_error_proto_enumTypes,
	}.Build()
	File_free_v1_free_error_proto = out.File
	file_free_v1_free_error_proto_rawDesc = nil
	file_free_v1_free_error_proto_goTypes = nil
	file_free_v1_free_error_proto_depIdxs = nil
}
