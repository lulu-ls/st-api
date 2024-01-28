// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/activity/v1/red_packet.proto

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

// Validate checks the field values on ReceiveRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReceiveRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReceiveRequestMultiError,
// or nil if none found.
func (m *ReceiveRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for UserId

	// no validation rules for RedPacketId

	// no validation rules for ShareConfigId

	if len(errors) > 0 {
		return ReceiveRequestMultiError(errors)
	}

	return nil
}

// ReceiveRequestMultiError is an error wrapping multiple validation errors
// returned by ReceiveRequest.ValidateAll() if the designated constraints
// aren't met.
type ReceiveRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveRequestMultiError) AllErrors() []error { return m }

// ReceiveRequestValidationError is the validation error returned by
// ReceiveRequest.Validate if the designated constraints aren't met.
type ReceiveRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveRequestValidationError) ErrorName() string { return "ReceiveRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReceiveRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveRequestValidationError{}

// Validate checks the field values on ReceiveReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReceiveReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReceiveReplyMultiError, or
// nil if none found.
func (m *ReceiveReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReceiveReplyValidationError{
					field:  "Detail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReceiveReplyValidationError{
					field:  "Detail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReceiveReplyValidationError{
				field:  "Detail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Number

	for idx, item := range m.GetHistoryList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReceiveReplyValidationError{
						field:  fmt.Sprintf("HistoryList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReceiveReplyValidationError{
						field:  fmt.Sprintf("HistoryList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReceiveReplyValidationError{
					field:  fmt.Sprintf("HistoryList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetRewardList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReceiveReplyValidationError{
						field:  fmt.Sprintf("RewardList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReceiveReplyValidationError{
						field:  fmt.Sprintf("RewardList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReceiveReplyValidationError{
					field:  fmt.Sprintf("RewardList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Tips

	if len(errors) > 0 {
		return ReceiveReplyMultiError(errors)
	}

	return nil
}

// ReceiveReplyMultiError is an error wrapping multiple validation errors
// returned by ReceiveReply.ValidateAll() if the designated constraints aren't met.
type ReceiveReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveReplyMultiError) AllErrors() []error { return m }

// ReceiveReplyValidationError is the validation error returned by
// ReceiveReply.Validate if the designated constraints aren't met.
type ReceiveReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveReplyValidationError) ErrorName() string { return "ReceiveReplyValidationError" }

// Error satisfies the builtin error interface
func (e ReceiveReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveReplyValidationError{}

// Validate checks the field values on HistoryList with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HistoryList) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HistoryList with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HistoryListMultiError, or
// nil if none found.
func (m *HistoryList) ValidateAll() error {
	return m.validate(true)
}

func (m *HistoryList) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Nickname

	// no validation rules for Text

	if len(errors) > 0 {
		return HistoryListMultiError(errors)
	}

	return nil
}

// HistoryListMultiError is an error wrapping multiple validation errors
// returned by HistoryList.ValidateAll() if the designated constraints aren't met.
type HistoryListMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HistoryListMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HistoryListMultiError) AllErrors() []error { return m }

// HistoryListValidationError is the validation error returned by
// HistoryList.Validate if the designated constraints aren't met.
type HistoryListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HistoryListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HistoryListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HistoryListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HistoryListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HistoryListValidationError) ErrorName() string { return "HistoryListValidationError" }

// Error satisfies the builtin error interface
func (e HistoryListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHistoryList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HistoryListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HistoryListValidationError{}

// Validate checks the field values on RewardItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RewardItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RewardItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RewardItemMultiError, or
// nil if none found.
func (m *RewardItem) ValidateAll() error {
	return m.validate(true)
}

func (m *RewardItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ItemId

	// no validation rules for Quantity

	if len(errors) > 0 {
		return RewardItemMultiError(errors)
	}

	return nil
}

// RewardItemMultiError is an error wrapping multiple validation errors
// returned by RewardItem.ValidateAll() if the designated constraints aren't met.
type RewardItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RewardItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RewardItemMultiError) AllErrors() []error { return m }

// RewardItemValidationError is the validation error returned by
// RewardItem.Validate if the designated constraints aren't met.
type RewardItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RewardItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RewardItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RewardItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RewardItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RewardItemValidationError) ErrorName() string { return "RewardItemValidationError" }

// Error satisfies the builtin error interface
func (e RewardItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRewardItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RewardItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RewardItemValidationError{}

// Validate checks the field values on GetShareConfRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetShareConfRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetShareConfRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetShareConfRequestMultiError, or nil if none found.
func (m *GetShareConfRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetShareConfRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for UserId

	// no validation rules for Type

	if len(errors) > 0 {
		return GetShareConfRequestMultiError(errors)
	}

	return nil
}

// GetShareConfRequestMultiError is an error wrapping multiple validation
// errors returned by GetShareConfRequest.ValidateAll() if the designated
// constraints aren't met.
type GetShareConfRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetShareConfRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetShareConfRequestMultiError) AllErrors() []error { return m }

// GetShareConfRequestValidationError is the validation error returned by
// GetShareConfRequest.Validate if the designated constraints aren't met.
type GetShareConfRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetShareConfRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetShareConfRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetShareConfRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetShareConfRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetShareConfRequestValidationError) ErrorName() string {
	return "GetShareConfRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetShareConfRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetShareConfRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetShareConfRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetShareConfRequestValidationError{}

// Validate checks the field values on GetShareConfReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetShareConfReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetShareConfReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetShareConfReplyMultiError, or nil if none found.
func (m *GetShareConfReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetShareConfReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShareConfigId

	// no validation rules for Title

	// no validation rules for Subtitle

	// no validation rules for Desc

	// no validation rules for Information

	// no validation rules for Image

	// no validation rules for Type

	if len(errors) > 0 {
		return GetShareConfReplyMultiError(errors)
	}

	return nil
}

// GetShareConfReplyMultiError is an error wrapping multiple validation errors
// returned by GetShareConfReply.ValidateAll() if the designated constraints
// aren't met.
type GetShareConfReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetShareConfReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetShareConfReplyMultiError) AllErrors() []error { return m }

// GetShareConfReplyValidationError is the validation error returned by
// GetShareConfReply.Validate if the designated constraints aren't met.
type GetShareConfReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetShareConfReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetShareConfReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetShareConfReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetShareConfReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetShareConfReplyValidationError) ErrorName() string {
	return "GetShareConfReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetShareConfReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetShareConfReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetShareConfReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetShareConfReplyValidationError{}

// Validate checks the field values on RedPacketRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RedPacketRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RedPacketRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RedPacketRequestMultiError, or nil if none found.
func (m *RedPacketRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RedPacketRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for UserId

	// no validation rules for ShareConfigId

	if len(errors) > 0 {
		return RedPacketRequestMultiError(errors)
	}

	return nil
}

// RedPacketRequestMultiError is an error wrapping multiple validation errors
// returned by RedPacketRequest.ValidateAll() if the designated constraints
// aren't met.
type RedPacketRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RedPacketRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RedPacketRequestMultiError) AllErrors() []error { return m }

// RedPacketRequestValidationError is the validation error returned by
// RedPacketRequest.Validate if the designated constraints aren't met.
type RedPacketRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedPacketRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedPacketRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedPacketRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedPacketRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedPacketRequestValidationError) ErrorName() string { return "RedPacketRequestValidationError" }

// Error satisfies the builtin error interface
func (e RedPacketRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedPacketRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedPacketRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedPacketRequestValidationError{}

// Validate checks the field values on RedPacketReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RedPacketReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RedPacketReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RedPacketReplyMultiError,
// or nil if none found.
func (m *RedPacketReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RedPacketReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShareConfigId

	// no validation rules for RedPacketId

	if len(errors) > 0 {
		return RedPacketReplyMultiError(errors)
	}

	return nil
}

// RedPacketReplyMultiError is an error wrapping multiple validation errors
// returned by RedPacketReply.ValidateAll() if the designated constraints
// aren't met.
type RedPacketReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RedPacketReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RedPacketReplyMultiError) AllErrors() []error { return m }

// RedPacketReplyValidationError is the validation error returned by
// RedPacketReply.Validate if the designated constraints aren't met.
type RedPacketReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedPacketReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedPacketReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedPacketReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedPacketReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedPacketReplyValidationError) ErrorName() string { return "RedPacketReplyValidationError" }

// Error satisfies the builtin error interface
func (e RedPacketReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedPacketReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedPacketReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedPacketReplyValidationError{}
