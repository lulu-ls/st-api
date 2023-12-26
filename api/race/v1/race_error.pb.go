// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: race/v1/race_error.proto

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
	ErrorReason_RACE_SIGNUP_TIME_LIMIT         ErrorReason = 11
	ErrorReason_RACE_IS_START                  ErrorReason = 12 // 比赛已开始
	ErrorReason_RACE_SIGNUP_FAIL_PEOPLE        ErrorReason = 13
	ErrorReason_RACE_SIGNUP_CANCEL             ErrorReason = 14 // 已经取消报名
	ErrorReason_RACE_IS_IN_GAME                ErrorReason = 15 // 已在比赛中
	ErrorReason_RACE_IS_SIGNUP                 ErrorReason = 16 // 已经报名其他比赛
	ErrorReason_RACE_AUTH_SERIES_USER          ErrorReason = 17 // 是否有权限参与比才 497
	ErrorReason_RACE_SERVER_ERROR              ErrorReason = 18 // 服务器维护中
	ErrorReason_RACE_SIGNUP_FAIL_TIME          ErrorReason = 19 // 比赛可用时间已过
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:  "USER_NOT_FOUND",
		1:  "CONTENT_MISSING",
		2:  "TOKEN_VERIFY",
		10: "AUTH_CHANNEL_APPID_IS_REQUIRED",
		11: "RACE_SIGNUP_TIME_LIMIT",
		12: "RACE_IS_START",
		13: "RACE_SIGNUP_FAIL_PEOPLE",
		14: "RACE_SIGNUP_CANCEL",
		15: "RACE_IS_IN_GAME",
		16: "RACE_IS_SIGNUP",
		17: "RACE_AUTH_SERIES_USER",
		18: "RACE_SERVER_ERROR",
		19: "RACE_SIGNUP_FAIL_TIME",
	}
	ErrorReason_value = map[string]int32{
		"USER_NOT_FOUND":                 0,
		"CONTENT_MISSING":                1,
		"TOKEN_VERIFY":                   2,
		"AUTH_CHANNEL_APPID_IS_REQUIRED": 10,
		"RACE_SIGNUP_TIME_LIMIT":         11,
		"RACE_IS_START":                  12,
		"RACE_SIGNUP_FAIL_PEOPLE":        13,
		"RACE_SIGNUP_CANCEL":             14,
		"RACE_IS_IN_GAME":                15,
		"RACE_IS_SIGNUP":                 16,
		"RACE_AUTH_SERIES_USER":          17,
		"RACE_SERVER_ERROR":              18,
		"RACE_SIGNUP_FAIL_TIME":          19,
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
	return file_race_v1_race_error_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_race_v1_race_error_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_race_v1_race_error_proto_rawDescGZIP(), []int{0}
}

var File_race_v1_race_error_proto protoreflect.FileDescriptor

var file_race_v1_race_error_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x9a, 0x03, 0x0a,
	0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x0e,
	0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00,
	0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e,
	0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x90,
	0x03, 0x12, 0x16, 0x0a, 0x0c, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x59, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x28, 0x0a, 0x1e, 0x41, 0x55, 0x54,
	0x48, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x41, 0x50, 0x50, 0x49, 0x44, 0x5f,
	0x49, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0a, 0x1a, 0x04, 0xa8,
	0x45, 0x90, 0x03, 0x12, 0x20, 0x0a, 0x16, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e,
	0x55, 0x50, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x0b, 0x1a,
	0x04, 0xa8, 0x45, 0x9a, 0x03, 0x12, 0x17, 0x0a, 0x0d, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x49, 0x53,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x0c, 0x1a, 0x04, 0xa8, 0x45, 0x9c, 0x03, 0x12, 0x21,
	0x0a, 0x17, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x5f, 0x50, 0x45, 0x4f, 0x50, 0x4c, 0x45, 0x10, 0x0d, 0x1a, 0x04, 0xa8, 0x45, 0x9d,
	0x03, 0x12, 0x1c, 0x0a, 0x12, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x0e, 0x1a, 0x04, 0xa8, 0x45, 0x9e, 0x03, 0x12,
	0x19, 0x0a, 0x0f, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x49, 0x4e, 0x5f, 0x47, 0x41,
	0x4d, 0x45, 0x10, 0x0f, 0x1a, 0x04, 0xa8, 0x45, 0xf2, 0x03, 0x12, 0x18, 0x0a, 0x0e, 0x52, 0x41,
	0x43, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50, 0x10, 0x10, 0x1a, 0x04,
	0xa8, 0x45, 0x9f, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x5f, 0x53, 0x45, 0x52, 0x49, 0x45, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x11, 0x1a,
	0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x12, 0x1a, 0x04, 0xa8, 0x45,
	0x94, 0x03, 0x12, 0x1f, 0x0a, 0x15, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55,
	0x50, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x13, 0x1a, 0x04, 0xa8,
	0x45, 0x9f, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x39, 0x0a, 0x0b, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x6c, 0x75, 0x2d, 0x6c, 0x73, 0x2f, 0x73,
	0x74, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_race_v1_race_error_proto_rawDescOnce sync.Once
	file_race_v1_race_error_proto_rawDescData = file_race_v1_race_error_proto_rawDesc
)

func file_race_v1_race_error_proto_rawDescGZIP() []byte {
	file_race_v1_race_error_proto_rawDescOnce.Do(func() {
		file_race_v1_race_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_race_v1_race_error_proto_rawDescData)
	})
	return file_race_v1_race_error_proto_rawDescData
}

var file_race_v1_race_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_race_v1_race_error_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: api.race.v1.ErrorReason
}
var file_race_v1_race_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_race_v1_race_error_proto_init() }
func file_race_v1_race_error_proto_init() {
	if File_race_v1_race_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_race_v1_race_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_race_v1_race_error_proto_goTypes,
		DependencyIndexes: file_race_v1_race_error_proto_depIdxs,
		EnumInfos:         file_race_v1_race_error_proto_enumTypes,
	}.Build()
	File_race_v1_race_error_proto = out.File
	file_race_v1_race_error_proto_rawDesc = nil
	file_race_v1_race_error_proto_goTypes = nil
	file_race_v1_race_error_proto_depIdxs = nil
}
