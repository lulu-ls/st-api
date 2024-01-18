// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/v1/user.proto

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

// Validate checks the field values on RaceRecordListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RaceRecordListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RaceRecordListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RaceRecordListRequestMultiError, or nil if none found.
func (m *RaceRecordListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RaceRecordListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RaceRecordListRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RaceRecordListRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RaceRecordListRequestValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserId

	// no validation rules for AppId

	if len(errors) > 0 {
		return RaceRecordListRequestMultiError(errors)
	}

	return nil
}

// RaceRecordListRequestMultiError is an error wrapping multiple validation
// errors returned by RaceRecordListRequest.ValidateAll() if the designated
// constraints aren't met.
type RaceRecordListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RaceRecordListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RaceRecordListRequestMultiError) AllErrors() []error { return m }

// RaceRecordListRequestValidationError is the validation error returned by
// RaceRecordListRequest.Validate if the designated constraints aren't met.
type RaceRecordListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RaceRecordListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RaceRecordListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RaceRecordListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RaceRecordListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RaceRecordListRequestValidationError) ErrorName() string {
	return "RaceRecordListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RaceRecordListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRaceRecordListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RaceRecordListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RaceRecordListRequestValidationError{}

// Validate checks the field values on RaceRecordListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RaceRecordListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RaceRecordListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RaceRecordListReplyMultiError, or nil if none found.
func (m *RaceRecordListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RaceRecordListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RaceRecordListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RaceRecordListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RaceRecordListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return RaceRecordListReplyMultiError(errors)
	}

	return nil
}

// RaceRecordListReplyMultiError is an error wrapping multiple validation
// errors returned by RaceRecordListReply.ValidateAll() if the designated
// constraints aren't met.
type RaceRecordListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RaceRecordListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RaceRecordListReplyMultiError) AllErrors() []error { return m }

// RaceRecordListReplyValidationError is the validation error returned by
// RaceRecordListReply.Validate if the designated constraints aren't met.
type RaceRecordListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RaceRecordListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RaceRecordListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RaceRecordListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RaceRecordListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RaceRecordListReplyValidationError) ErrorName() string {
	return "RaceRecordListReplyValidationError"
}

// Error satisfies the builtin error interface
func (e RaceRecordListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRaceRecordListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RaceRecordListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RaceRecordListReplyValidationError{}

// Validate checks the field values on RaceRecordListItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RaceRecordListItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RaceRecordListItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RaceRecordListItemMultiError, or nil if none found.
func (m *RaceRecordListItem) ValidateAll() error {
	return m.validate(true)
}

func (m *RaceRecordListItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for GameId

	// no validation rules for GameCategoryId

	// no validation rules for GameCategoryName

	// no validation rules for GameConfigId

	// no validation rules for GameConfigName

	// no validation rules for GameConfigTitle

	// no validation rules for Rounds

	// no validation rules for Rank

	// no validation rules for Win

	// no validation rules for Lose

	// no validation rules for EnterTime

	// no validation rules for LeaveTime

	if len(errors) > 0 {
		return RaceRecordListItemMultiError(errors)
	}

	return nil
}

// RaceRecordListItemMultiError is an error wrapping multiple validation errors
// returned by RaceRecordListItem.ValidateAll() if the designated constraints
// aren't met.
type RaceRecordListItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RaceRecordListItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RaceRecordListItemMultiError) AllErrors() []error { return m }

// RaceRecordListItemValidationError is the validation error returned by
// RaceRecordListItem.Validate if the designated constraints aren't met.
type RaceRecordListItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RaceRecordListItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RaceRecordListItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RaceRecordListItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RaceRecordListItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RaceRecordListItemValidationError) ErrorName() string {
	return "RaceRecordListItemValidationError"
}

// Error satisfies the builtin error interface
func (e RaceRecordListItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRaceRecordListItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RaceRecordListItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RaceRecordListItemValidationError{}

// Validate checks the field values on InfoEditRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *InfoEditRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InfoEditRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InfoEditRequestMultiError, or nil if none found.
func (m *InfoEditRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InfoEditRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for Figure

	// no validation rules for Gender

	if len(errors) > 0 {
		return InfoEditRequestMultiError(errors)
	}

	return nil
}

// InfoEditRequestMultiError is an error wrapping multiple validation errors
// returned by InfoEditRequest.ValidateAll() if the designated constraints
// aren't met.
type InfoEditRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InfoEditRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InfoEditRequestMultiError) AllErrors() []error { return m }

// InfoEditRequestValidationError is the validation error returned by
// InfoEditRequest.Validate if the designated constraints aren't met.
type InfoEditRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoEditRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoEditRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoEditRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoEditRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoEditRequestValidationError) ErrorName() string { return "InfoEditRequestValidationError" }

// Error satisfies the builtin error interface
func (e InfoEditRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoEditRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoEditRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoEditRequestValidationError{}

// Validate checks the field values on InfoEditReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InfoEditReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InfoEditReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InfoEditReplyMultiError, or
// nil if none found.
func (m *InfoEditReply) ValidateAll() error {
	return m.validate(true)
}

func (m *InfoEditReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return InfoEditReplyMultiError(errors)
	}

	return nil
}

// InfoEditReplyMultiError is an error wrapping multiple validation errors
// returned by InfoEditReply.ValidateAll() if the designated constraints
// aren't met.
type InfoEditReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InfoEditReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InfoEditReplyMultiError) AllErrors() []error { return m }

// InfoEditReplyValidationError is the validation error returned by
// InfoEditReply.Validate if the designated constraints aren't met.
type InfoEditReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoEditReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoEditReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoEditReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoEditReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoEditReplyValidationError) ErrorName() string { return "InfoEditReplyValidationError" }

// Error satisfies the builtin error interface
func (e InfoEditReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfoEditReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoEditReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoEditReplyValidationError{}

// Validate checks the field values on AssetGetRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AssetGetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AssetGetRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AssetGetRequestMultiError, or nil if none found.
func (m *AssetGetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AssetGetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for AppId

	// no validation rules for ChannelId

	if len(errors) > 0 {
		return AssetGetRequestMultiError(errors)
	}

	return nil
}

// AssetGetRequestMultiError is an error wrapping multiple validation errors
// returned by AssetGetRequest.ValidateAll() if the designated constraints
// aren't met.
type AssetGetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AssetGetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AssetGetRequestMultiError) AllErrors() []error { return m }

// AssetGetRequestValidationError is the validation error returned by
// AssetGetRequest.Validate if the designated constraints aren't met.
type AssetGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AssetGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AssetGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AssetGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AssetGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AssetGetRequestValidationError) ErrorName() string { return "AssetGetRequestValidationError" }

// Error satisfies the builtin error interface
func (e AssetGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAssetGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AssetGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AssetGetRequestValidationError{}

// Validate checks the field values on AssetGetReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AssetGetReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AssetGetReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AssetGetReplyMultiError, or
// nil if none found.
func (m *AssetGetReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AssetGetReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Gold

	// no validation rules for Diamond

	// no validation rules for Ticket

	// no validation rules for CardCounter

	if len(errors) > 0 {
		return AssetGetReplyMultiError(errors)
	}

	return nil
}

// AssetGetReplyMultiError is an error wrapping multiple validation errors
// returned by AssetGetReply.ValidateAll() if the designated constraints
// aren't met.
type AssetGetReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AssetGetReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AssetGetReplyMultiError) AllErrors() []error { return m }

// AssetGetReplyValidationError is the validation error returned by
// AssetGetReply.Validate if the designated constraints aren't met.
type AssetGetReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AssetGetReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AssetGetReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AssetGetReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AssetGetReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AssetGetReplyValidationError) ErrorName() string { return "AssetGetReplyValidationError" }

// Error satisfies the builtin error interface
func (e AssetGetReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAssetGetReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AssetGetReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AssetGetReplyValidationError{}

// Validate checks the field values on NotifyStateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NotifyStateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotifyStateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotifyStateRequestMultiError, or nil if none found.
func (m *NotifyStateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *NotifyStateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for State

	// no validation rules for Times

	if len(errors) > 0 {
		return NotifyStateRequestMultiError(errors)
	}

	return nil
}

// NotifyStateRequestMultiError is an error wrapping multiple validation errors
// returned by NotifyStateRequest.ValidateAll() if the designated constraints
// aren't met.
type NotifyStateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotifyStateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotifyStateRequestMultiError) AllErrors() []error { return m }

// NotifyStateRequestValidationError is the validation error returned by
// NotifyStateRequest.Validate if the designated constraints aren't met.
type NotifyStateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotifyStateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotifyStateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotifyStateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotifyStateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotifyStateRequestValidationError) ErrorName() string {
	return "NotifyStateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NotifyStateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotifyStateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotifyStateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotifyStateRequestValidationError{}

// Validate checks the field values on NotifyStateReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NotifyStateReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotifyStateReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotifyStateReplyMultiError, or nil if none found.
func (m *NotifyStateReply) ValidateAll() error {
	return m.validate(true)
}

func (m *NotifyStateReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return NotifyStateReplyMultiError(errors)
	}

	return nil
}

// NotifyStateReplyMultiError is an error wrapping multiple validation errors
// returned by NotifyStateReply.ValidateAll() if the designated constraints
// aren't met.
type NotifyStateReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotifyStateReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotifyStateReplyMultiError) AllErrors() []error { return m }

// NotifyStateReplyValidationError is the validation error returned by
// NotifyStateReply.Validate if the designated constraints aren't met.
type NotifyStateReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotifyStateReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotifyStateReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotifyStateReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotifyStateReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotifyStateReplyValidationError) ErrorName() string { return "NotifyStateReplyValidationError" }

// Error satisfies the builtin error interface
func (e NotifyStateReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotifyStateReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotifyStateReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotifyStateReplyValidationError{}

// Validate checks the field values on NotifyCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NotifyCheckRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotifyCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotifyCheckRequestMultiError, or nil if none found.
func (m *NotifyCheckRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *NotifyCheckRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for Type

	if len(errors) > 0 {
		return NotifyCheckRequestMultiError(errors)
	}

	return nil
}

// NotifyCheckRequestMultiError is an error wrapping multiple validation errors
// returned by NotifyCheckRequest.ValidateAll() if the designated constraints
// aren't met.
type NotifyCheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotifyCheckRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotifyCheckRequestMultiError) AllErrors() []error { return m }

// NotifyCheckRequestValidationError is the validation error returned by
// NotifyCheckRequest.Validate if the designated constraints aren't met.
type NotifyCheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotifyCheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotifyCheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotifyCheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotifyCheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotifyCheckRequestValidationError) ErrorName() string {
	return "NotifyCheckRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NotifyCheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotifyCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotifyCheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotifyCheckRequestValidationError{}

// Validate checks the field values on NotifyCheckReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NotifyCheckReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotifyCheckReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotifyCheckReplyMultiError, or nil if none found.
func (m *NotifyCheckReply) ValidateAll() error {
	return m.validate(true)
}

func (m *NotifyCheckReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return NotifyCheckReplyMultiError(errors)
	}

	return nil
}

// NotifyCheckReplyMultiError is an error wrapping multiple validation errors
// returned by NotifyCheckReply.ValidateAll() if the designated constraints
// aren't met.
type NotifyCheckReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotifyCheckReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotifyCheckReplyMultiError) AllErrors() []error { return m }

// NotifyCheckReplyValidationError is the validation error returned by
// NotifyCheckReply.Validate if the designated constraints aren't met.
type NotifyCheckReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotifyCheckReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotifyCheckReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotifyCheckReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotifyCheckReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotifyCheckReplyValidationError) ErrorName() string { return "NotifyCheckReplyValidationError" }

// Error satisfies the builtin error interface
func (e NotifyCheckReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotifyCheckReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotifyCheckReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotifyCheckReplyValidationError{}
