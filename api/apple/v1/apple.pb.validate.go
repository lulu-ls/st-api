// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: apple/v1/apple.proto

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

// Validate checks the field values on VerifyReceiptRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyReceiptRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyReceiptRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyReceiptRequestMultiError, or nil if none found.
func (m *VerifyReceiptRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyReceiptRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for PurchaseId

	// no validation rules for UserId

	// no validation rules for OrderId

	if len(errors) > 0 {
		return VerifyReceiptRequestMultiError(errors)
	}

	return nil
}

// VerifyReceiptRequestMultiError is an error wrapping multiple validation
// errors returned by VerifyReceiptRequest.ValidateAll() if the designated
// constraints aren't met.
type VerifyReceiptRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyReceiptRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyReceiptRequestMultiError) AllErrors() []error { return m }

// VerifyReceiptRequestValidationError is the validation error returned by
// VerifyReceiptRequest.Validate if the designated constraints aren't met.
type VerifyReceiptRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyReceiptRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyReceiptRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyReceiptRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyReceiptRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyReceiptRequestValidationError) ErrorName() string {
	return "VerifyReceiptRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyReceiptRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyReceiptRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyReceiptRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyReceiptRequestValidationError{}

// Validate checks the field values on VerifyReceiptReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyReceiptReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyReceiptReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyReceiptReplyMultiError, or nil if none found.
func (m *VerifyReceiptReply) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyReceiptReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Env

	// no validation rules for Platform

	// no validation rules for OutTradNo

	// no validation rules for OrderAmount

	// no validation rules for Quantity

	// no validation rules for Mode

	// no validation rules for OrderId

	// no validation rules for Title

	if len(errors) > 0 {
		return VerifyReceiptReplyMultiError(errors)
	}

	return nil
}

// VerifyReceiptReplyMultiError is an error wrapping multiple validation errors
// returned by VerifyReceiptReply.ValidateAll() if the designated constraints
// aren't met.
type VerifyReceiptReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyReceiptReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyReceiptReplyMultiError) AllErrors() []error { return m }

// VerifyReceiptReplyValidationError is the validation error returned by
// VerifyReceiptReply.Validate if the designated constraints aren't met.
type VerifyReceiptReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyReceiptReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyReceiptReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyReceiptReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyReceiptReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyReceiptReplyValidationError) ErrorName() string {
	return "VerifyReceiptReplyValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyReceiptReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyReceiptReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyReceiptReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyReceiptReplyValidationError{}

// Validate checks the field values on PayNoticeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PayNoticeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayNoticeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PayNoticeRequestMultiError, or nil if none found.
func (m *PayNoticeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PayNoticeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PayNoticeRequestMultiError(errors)
	}

	return nil
}

// PayNoticeRequestMultiError is an error wrapping multiple validation errors
// returned by PayNoticeRequest.ValidateAll() if the designated constraints
// aren't met.
type PayNoticeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayNoticeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayNoticeRequestMultiError) AllErrors() []error { return m }

// PayNoticeRequestValidationError is the validation error returned by
// PayNoticeRequest.Validate if the designated constraints aren't met.
type PayNoticeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayNoticeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayNoticeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayNoticeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayNoticeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayNoticeRequestValidationError) ErrorName() string { return "PayNoticeRequestValidationError" }

// Error satisfies the builtin error interface
func (e PayNoticeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayNoticeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayNoticeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayNoticeRequestValidationError{}

// Validate checks the field values on PayNoticeReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PayNoticeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayNoticeReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PayNoticeReplyMultiError,
// or nil if none found.
func (m *PayNoticeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PayNoticeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PayNoticeReplyMultiError(errors)
	}

	return nil
}

// PayNoticeReplyMultiError is an error wrapping multiple validation errors
// returned by PayNoticeReply.ValidateAll() if the designated constraints
// aren't met.
type PayNoticeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayNoticeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayNoticeReplyMultiError) AllErrors() []error { return m }

// PayNoticeReplyValidationError is the validation error returned by
// PayNoticeReply.Validate if the designated constraints aren't met.
type PayNoticeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayNoticeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayNoticeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayNoticeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayNoticeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayNoticeReplyValidationError) ErrorName() string { return "PayNoticeReplyValidationError" }

// Error satisfies the builtin error interface
func (e PayNoticeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayNoticeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayNoticeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayNoticeReplyValidationError{}

// Validate checks the field values on ApplePayBuyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ApplePayBuyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApplePayBuyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ApplePayBuyRequestMultiError, or nil if none found.
func (m *ApplePayBuyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ApplePayBuyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for PurchaseId

	// no validation rules for UserId

	// no validation rules for OrderId

	if len(errors) > 0 {
		return ApplePayBuyRequestMultiError(errors)
	}

	return nil
}

// ApplePayBuyRequestMultiError is an error wrapping multiple validation errors
// returned by ApplePayBuyRequest.ValidateAll() if the designated constraints
// aren't met.
type ApplePayBuyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplePayBuyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplePayBuyRequestMultiError) AllErrors() []error { return m }

// ApplePayBuyRequestValidationError is the validation error returned by
// ApplePayBuyRequest.Validate if the designated constraints aren't met.
type ApplePayBuyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplePayBuyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplePayBuyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplePayBuyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplePayBuyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplePayBuyRequestValidationError) ErrorName() string {
	return "ApplePayBuyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ApplePayBuyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplePayBuyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplePayBuyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplePayBuyRequestValidationError{}

// Validate checks the field values on ApplePayBuyReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ApplePayBuyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApplePayBuyReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ApplePayBuyReplyMultiError, or nil if none found.
func (m *ApplePayBuyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ApplePayBuyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for OrderAmount

	// no validation rules for Quantity

	// no validation rules for Title

	// no validation rules for Env

	if len(errors) > 0 {
		return ApplePayBuyReplyMultiError(errors)
	}

	return nil
}

// ApplePayBuyReplyMultiError is an error wrapping multiple validation errors
// returned by ApplePayBuyReply.ValidateAll() if the designated constraints
// aren't met.
type ApplePayBuyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplePayBuyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplePayBuyReplyMultiError) AllErrors() []error { return m }

// ApplePayBuyReplyValidationError is the validation error returned by
// ApplePayBuyReply.Validate if the designated constraints aren't met.
type ApplePayBuyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplePayBuyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplePayBuyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplePayBuyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplePayBuyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplePayBuyReplyValidationError) ErrorName() string { return "ApplePayBuyReplyValidationError" }

// Error satisfies the builtin error interface
func (e ApplePayBuyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplePayBuyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplePayBuyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplePayBuyReplyValidationError{}

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

	if utf8.RuneCountInString(m.GetCode()) < 1 {
		err := LoginRequestValidationError{
			field:  "Code",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

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

	// no validation rules for UniqueId

	// no validation rules for Email

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