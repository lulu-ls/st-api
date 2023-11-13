// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/v1/auth.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetCode()) < 5 {
		err := LoginRequestValidationError{
			field:  "Code",
			reason: "value length must be at least 5 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for AppId

	// no validation rules for ChannelId

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReplyMultiError, or
// nil if none found.
func (m *LoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	// no validation rules for ExpireTime

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginReplyValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginReplyValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginReplyValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LoginReplyMultiError(errors)
	}

	return nil
}

// LoginReplyMultiError is an error wrapping multiple validation errors
// returned by LoginReply.ValidateAll() if the designated constraints aren't met.
type LoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReplyMultiError) AllErrors() []error { return m }

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on LoginTestRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginTestRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginTestRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginTestRequestMultiError, or nil if none found.
func (m *LoginTestRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginTestRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserName

	// no validation rules for Password

	// no validation rules for AppId

	// no validation rules for ChannelId

	if len(errors) > 0 {
		return LoginTestRequestMultiError(errors)
	}

	return nil
}

// LoginTestRequestMultiError is an error wrapping multiple validation errors
// returned by LoginTestRequest.ValidateAll() if the designated constraints
// aren't met.
type LoginTestRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginTestRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginTestRequestMultiError) AllErrors() []error { return m }

// LoginTestRequestValidationError is the validation error returned by
// LoginTestRequest.Validate if the designated constraints aren't met.
type LoginTestRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginTestRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginTestRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginTestRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginTestRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginTestRequestValidationError) ErrorName() string { return "LoginTestRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginTestRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginTestRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginTestRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginTestRequestValidationError{}

// Validate checks the field values on LoginTestReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginTestReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginTestReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginTestReplyMultiError,
// or nil if none found.
func (m *LoginTestReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginTestReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LoginTestReplyMultiError(errors)
	}

	return nil
}

// LoginTestReplyMultiError is an error wrapping multiple validation errors
// returned by LoginTestReply.ValidateAll() if the designated constraints
// aren't met.
type LoginTestReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginTestReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginTestReplyMultiError) AllErrors() []error { return m }

// LoginTestReplyValidationError is the validation error returned by
// LoginTestReply.Validate if the designated constraints aren't met.
type LoginTestReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginTestReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginTestReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginTestReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginTestReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginTestReplyValidationError) ErrorName() string { return "LoginTestReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginTestReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginTestReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginTestReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginTestReplyValidationError{}

// Validate checks the field values on DecryptRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DecryptRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DecryptRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DecryptRequestMultiError,
// or nil if none found.
func (m *DecryptRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DecryptRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetEncryptedData()) < 5 {
		err := DecryptRequestValidationError{
			field:  "EncryptedData",
			reason: "value length must be at least 5 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetIv()) < 5 {
		err := DecryptRequestValidationError{
			field:  "Iv",
			reason: "value length must be at least 5 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUserInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DecryptRequestValidationError{
					field:  "UserInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DecryptRequestValidationError{
					field:  "UserInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DecryptRequestValidationError{
				field:  "UserInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	// no validation rules for AppId

	// no validation rules for ChannelId

	if len(errors) > 0 {
		return DecryptRequestMultiError(errors)
	}

	return nil
}

// DecryptRequestMultiError is an error wrapping multiple validation errors
// returned by DecryptRequest.ValidateAll() if the designated constraints
// aren't met.
type DecryptRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DecryptRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DecryptRequestMultiError) AllErrors() []error { return m }

// DecryptRequestValidationError is the validation error returned by
// DecryptRequest.Validate if the designated constraints aren't met.
type DecryptRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecryptRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecryptRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecryptRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecryptRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecryptRequestValidationError) ErrorName() string { return "DecryptRequestValidationError" }

// Error satisfies the builtin error interface
func (e DecryptRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecryptRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecryptRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecryptRequestValidationError{}

// Validate checks the field values on DecryptUserInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DecryptUserInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DecryptUserInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DecryptUserInfoMultiError, or nil if none found.
func (m *DecryptUserInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *DecryptUserInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Language

	// no validation rules for City

	// no validation rules for Province

	// no validation rules for AvatarUrl

	// no validation rules for NickName

	// no validation rules for Gender

	// no validation rules for Country

	// no validation rules for UnionId

	if len(errors) > 0 {
		return DecryptUserInfoMultiError(errors)
	}

	return nil
}

// DecryptUserInfoMultiError is an error wrapping multiple validation errors
// returned by DecryptUserInfo.ValidateAll() if the designated constraints
// aren't met.
type DecryptUserInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DecryptUserInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DecryptUserInfoMultiError) AllErrors() []error { return m }

// DecryptUserInfoValidationError is the validation error returned by
// DecryptUserInfo.Validate if the designated constraints aren't met.
type DecryptUserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecryptUserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecryptUserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecryptUserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecryptUserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecryptUserInfoValidationError) ErrorName() string { return "DecryptUserInfoValidationError" }

// Error satisfies the builtin error interface
func (e DecryptUserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecryptUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecryptUserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecryptUserInfoValidationError{}

// Validate checks the field values on DecryptReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DecryptReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DecryptReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DecryptReplyMultiError, or
// nil if none found.
func (m *DecryptReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DecryptReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Language

	// no validation rules for City

	// no validation rules for Province

	// no validation rules for Avatar

	// no validation rules for NickName

	// no validation rules for Gender

	// no validation rules for Country

	// no validation rules for UnionId

	// no validation rules for UserNo

	// no validation rules for UserId

	// no validation rules for Status

	// no validation rules for Figure

	if len(errors) > 0 {
		return DecryptReplyMultiError(errors)
	}

	return nil
}

// DecryptReplyMultiError is an error wrapping multiple validation errors
// returned by DecryptReply.ValidateAll() if the designated constraints aren't met.
type DecryptReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DecryptReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DecryptReplyMultiError) AllErrors() []error { return m }

// DecryptReplyValidationError is the validation error returned by
// DecryptReply.Validate if the designated constraints aren't met.
type DecryptReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecryptReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecryptReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecryptReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecryptReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecryptReplyValidationError) ErrorName() string { return "DecryptReplyValidationError" }

// Error satisfies the builtin error interface
func (e DecryptReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecryptReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecryptReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecryptReplyValidationError{}

// Validate checks the field values on LoginUser with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginUser) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginUser with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginUserMultiError, or nil
// if none found.
func (m *LoginUser) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginUser) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserNo

	// no validation rules for NickName

	// no validation rules for Avatar

	// no validation rules for Phone

	// no validation rules for Email

	// no validation rules for Gender

	// no validation rules for BirthDate

	// no validation rules for Status

	// no validation rules for Gold

	// no validation rules for Diamond

	// no validation rules for Ticket

	// no validation rules for UserId

	// no validation rules for Figure

	if len(errors) > 0 {
		return LoginUserMultiError(errors)
	}

	return nil
}

// LoginUserMultiError is an error wrapping multiple validation errors returned
// by LoginUser.ValidateAll() if the designated constraints aren't met.
type LoginUserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginUserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginUserMultiError) AllErrors() []error { return m }

// LoginUserValidationError is the validation error returned by
// LoginUser.Validate if the designated constraints aren't met.
type LoginUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUserValidationError) ErrorName() string { return "LoginUserValidationError" }

// Error satisfies the builtin error interface
func (e LoginUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUserValidationError{}

// Validate checks the field values on Common with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Common) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Common with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CommonMultiError, or nil if none found.
func (m *Common) ValidateAll() error {
	return m.validate(true)
}

func (m *Common) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WxAppId

	// no validation rules for ChannelId

	if len(errors) > 0 {
		return CommonMultiError(errors)
	}

	return nil
}

// CommonMultiError is an error wrapping multiple validation errors returned by
// Common.ValidateAll() if the designated constraints aren't met.
type CommonMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonMultiError) AllErrors() []error { return m }

// CommonValidationError is the validation error returned by Common.Validate if
// the designated constraints aren't met.
type CommonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonValidationError) ErrorName() string { return "CommonValidationError" }

// Error satisfies the builtin error interface
func (e CommonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommon.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonValidationError{}

// Validate checks the field values on GetInfoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetInfoRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetInfoRequestMultiError,
// or nil if none found.
func (m *GetInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for AppId

	// no validation rules for ChannelId

	if len(errors) > 0 {
		return GetInfoRequestMultiError(errors)
	}

	return nil
}

// GetInfoRequestMultiError is an error wrapping multiple validation errors
// returned by GetInfoRequest.ValidateAll() if the designated constraints
// aren't met.
type GetInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetInfoRequestMultiError) AllErrors() []error { return m }

// GetInfoRequestValidationError is the validation error returned by
// GetInfoRequest.Validate if the designated constraints aren't met.
type GetInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoRequestValidationError) ErrorName() string { return "GetInfoRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoRequestValidationError{}

// Validate checks the field values on GetInfoReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetInfoReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetInfoReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetInfoReplyMultiError, or
// nil if none found.
func (m *GetInfoReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetInfoReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetInfoReplyMultiError(errors)
	}

	return nil
}

// GetInfoReplyMultiError is an error wrapping multiple validation errors
// returned by GetInfoReply.ValidateAll() if the designated constraints aren't met.
type GetInfoReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetInfoReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetInfoReplyMultiError) AllErrors() []error { return m }

// GetInfoReplyValidationError is the validation error returned by
// GetInfoReply.Validate if the designated constraints aren't met.
type GetInfoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoReplyValidationError) ErrorName() string { return "GetInfoReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetInfoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoReplyValidationError{}
