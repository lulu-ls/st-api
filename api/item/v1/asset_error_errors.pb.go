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
func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

// 为某个枚举单独设置错误码
func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}

func IsTokenVerify(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_TOKEN_VERIFY.String() && e.Code == 401
}

func ErrorTokenVerify(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_TOKEN_VERIFY.String(), fmt.Sprintf(format, args...))
}

func IsAssetTypeNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ASSET_TYPE_NOT_EXIST.String() && e.Code == 410
}

func ErrorAssetTypeNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(410, ErrorReason_ASSET_TYPE_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

func IsAssetCreateConfirm(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ASSET_CREATE_CONFIRM.String() && e.Code == 411
}

func ErrorAssetCreateConfirm(format string, args ...interface{}) *errors.Error {
	return errors.New(411, ErrorReason_ASSET_CREATE_CONFIRM.String(), fmt.Sprintf(format, args...))
}

func IsAssetNotEnough(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ASSET_NOT_ENOUGH.String() && e.Code == 412
}

func ErrorAssetNotEnough(format string, args ...interface{}) *errors.Error {
	return errors.New(412, ErrorReason_ASSET_NOT_ENOUGH.String(), fmt.Sprintf(format, args...))
}

func IsAssetQueryFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ASSET_QUERY_FAILED.String() && e.Code == 420
}

func ErrorAssetQueryFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(420, ErrorReason_ASSET_QUERY_FAILED.String(), fmt.Sprintf(format, args...))
}
