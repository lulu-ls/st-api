// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 为某个枚举单独设置错误码
func IsLootPoolNil(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_LOOT_POOL_NIL.String() && e.Code == 410
}

// 为某个枚举单独设置错误码
func ErrorLootPoolNil(format string, args ...interface{}) *errors.Error {
	return errors.New(410, ErrorReason_LOOT_POOL_NIL.String(), fmt.Sprintf(format, args...))
}

func IsBargainServerErr(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_BARGAIN_SERVER_ERR.String() && e.Code == 410
}

func ErrorBargainServerErr(format string, args ...interface{}) *errors.Error {
	return errors.New(410, ErrorReason_BARGAIN_SERVER_ERR.String(), fmt.Sprintf(format, args...))
}

func IsRedPacketServerErr(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_RedPacket_SERVER_ERR.String() && e.Code == 410
}

func ErrorRedPacketServerErr(format string, args ...interface{}) *errors.Error {
	return errors.New(410, ErrorReason_RedPacket_SERVER_ERR.String(), fmt.Sprintf(format, args...))
}
