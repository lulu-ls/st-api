// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ticket/v1/ticket.proto

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

	// no validation rules for AccessToken

	// no validation rules for MerchantId

	// no validation rules for MerchantName

	// no validation rules for StoreId

	// no validation rules for StoreName

	// no validation rules for HasStores

	for idx, item := range m.GetStores() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LoginReplyValidationError{
						field:  fmt.Sprintf("Stores[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LoginReplyValidationError{
						field:  fmt.Sprintf("Stores[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LoginReplyValidationError{
					field:  fmt.Sprintf("Stores[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
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

// Validate checks the field values on Store with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Store) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Store with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StoreMultiError, or nil if none found.
func (m *Store) ValidateAll() error {
	return m.validate(true)
}

func (m *Store) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StoreId

	// no validation rules for StoreName

	if len(errors) > 0 {
		return StoreMultiError(errors)
	}

	return nil
}

// StoreMultiError is an error wrapping multiple validation errors returned by
// Store.ValidateAll() if the designated constraints aren't met.
type StoreMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StoreMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StoreMultiError) AllErrors() []error { return m }

// StoreValidationError is the validation error returned by Store.Validate if
// the designated constraints aren't met.
type StoreValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StoreValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StoreValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StoreValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StoreValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StoreValidationError) ErrorName() string { return "StoreValidationError" }

// Error satisfies the builtin error interface
func (e StoreValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStore.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StoreValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StoreValidationError{}

// Validate checks the field values on StatisticsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StatisticsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatisticsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StatisticsRequestMultiError, or nil if none found.
func (m *StatisticsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StatisticsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EmployeeId

	// no validation rules for StoreId

	if len(errors) > 0 {
		return StatisticsRequestMultiError(errors)
	}

	return nil
}

// StatisticsRequestMultiError is an error wrapping multiple validation errors
// returned by StatisticsRequest.ValidateAll() if the designated constraints
// aren't met.
type StatisticsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatisticsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatisticsRequestMultiError) AllErrors() []error { return m }

// StatisticsRequestValidationError is the validation error returned by
// StatisticsRequest.Validate if the designated constraints aren't met.
type StatisticsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatisticsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatisticsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatisticsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatisticsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatisticsRequestValidationError) ErrorName() string {
	return "StatisticsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StatisticsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatisticsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatisticsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatisticsRequestValidationError{}

// Validate checks the field values on StatisticsReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StatisticsReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatisticsReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StatisticsReplyMultiError, or nil if none found.
func (m *StatisticsReply) ValidateAll() error {
	return m.validate(true)
}

func (m *StatisticsReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Quantity

	// no validation rules for Amount

	if len(errors) > 0 {
		return StatisticsReplyMultiError(errors)
	}

	return nil
}

// StatisticsReplyMultiError is an error wrapping multiple validation errors
// returned by StatisticsReply.ValidateAll() if the designated constraints
// aren't met.
type StatisticsReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatisticsReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatisticsReplyMultiError) AllErrors() []error { return m }

// StatisticsReplyValidationError is the validation error returned by
// StatisticsReply.Validate if the designated constraints aren't met.
type StatisticsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatisticsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatisticsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatisticsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatisticsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatisticsReplyValidationError) ErrorName() string { return "StatisticsReplyValidationError" }

// Error satisfies the builtin error interface
func (e StatisticsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatisticsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatisticsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatisticsReplyValidationError{}

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

	// no validation rules for Phone

	// no validation rules for Password

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
