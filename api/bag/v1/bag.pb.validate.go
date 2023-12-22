// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bag/v1/bag.proto

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

// Validate checks the field values on ItemListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ItemListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ItemListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ItemListRequestMultiError, or nil if none found.
func (m *ItemListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ItemListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ItemListRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ItemListRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ItemListRequestValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AppId

	// no validation rules for UserId

	if len(errors) > 0 {
		return ItemListRequestMultiError(errors)
	}

	return nil
}

// ItemListRequestMultiError is an error wrapping multiple validation errors
// returned by ItemListRequest.ValidateAll() if the designated constraints
// aren't met.
type ItemListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemListRequestMultiError) AllErrors() []error { return m }

// ItemListRequestValidationError is the validation error returned by
// ItemListRequest.Validate if the designated constraints aren't met.
type ItemListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemListRequestValidationError) ErrorName() string { return "ItemListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ItemListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItemListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemListRequestValidationError{}

// Validate checks the field values on ItemListReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ItemListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ItemListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ItemListReplyMultiError, or
// nil if none found.
func (m *ItemListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ItemListReply) validate(all bool) error {
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
					errors = append(errors, ItemListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ItemListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ItemListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ItemListReplyMultiError(errors)
	}

	return nil
}

// ItemListReplyMultiError is an error wrapping multiple validation errors
// returned by ItemListReply.ValidateAll() if the designated constraints
// aren't met.
type ItemListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemListReplyMultiError) AllErrors() []error { return m }

// ItemListReplyValidationError is the validation error returned by
// ItemListReply.Validate if the designated constraints aren't met.
type ItemListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemListReplyValidationError) ErrorName() string { return "ItemListReplyValidationError" }

// Error satisfies the builtin error interface
func (e ItemListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItemListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemListReplyValidationError{}

// Validate checks the field values on ListItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListItemMultiError, or nil
// if none found.
func (m *ListItem) ValidateAll() error {
	return m.validate(true)
}

func (m *ListItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ItemId

	// no validation rules for ItemCategoryId

	// no validation rules for ItemFuncId

	// no validation rules for Name

	// no validation rules for Title

	// no validation rules for Image

	// no validation rules for MaxStack

	// no validation rules for Total

	// no validation rules for ItemExchangeWareId

	// no validation rules for ItemQuantity

	// no validation rules for WareId

	// no validation rules for WareCategoryId

	// no validation rules for WareName

	// no validation rules for WareImage

	// no validation rules for WareExternalUrl

	// no validation rules for Information

	// no validation rules for IsPhoneBill

	// no validation rules for WareType

	// no validation rules for VoucherId

	// no validation rules for VoucherCode

	// no validation rules for VoucherName

	// no validation rules for VoucherTitle

	// no validation rules for VoucherStart

	// no validation rules for VoucherEnd

	if len(errors) > 0 {
		return ListItemMultiError(errors)
	}

	return nil
}

// ListItemMultiError is an error wrapping multiple validation errors returned
// by ListItem.ValidateAll() if the designated constraints aren't met.
type ListItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListItemMultiError) AllErrors() []error { return m }

// ListItemValidationError is the validation error returned by
// ListItem.Validate if the designated constraints aren't met.
type ListItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListItemValidationError) ErrorName() string { return "ListItemValidationError" }

// Error satisfies the builtin error interface
func (e ListItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListItemValidationError{}

// Validate checks the field values on ListEmojiRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListEmojiRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEmojiRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListEmojiRequestMultiError, or nil if none found.
func (m *ListEmojiRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEmojiRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for UserId

	if len(errors) > 0 {
		return ListEmojiRequestMultiError(errors)
	}

	return nil
}

// ListEmojiRequestMultiError is an error wrapping multiple validation errors
// returned by ListEmojiRequest.ValidateAll() if the designated constraints
// aren't met.
type ListEmojiRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEmojiRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEmojiRequestMultiError) AllErrors() []error { return m }

// ListEmojiRequestValidationError is the validation error returned by
// ListEmojiRequest.Validate if the designated constraints aren't met.
type ListEmojiRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEmojiRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEmojiRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEmojiRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEmojiRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEmojiRequestValidationError) ErrorName() string { return "ListEmojiRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListEmojiRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEmojiRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEmojiRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEmojiRequestValidationError{}

// Validate checks the field values on ListEmojiReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListEmojiReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEmojiReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListEmojiReplyMultiError,
// or nil if none found.
func (m *ListEmojiReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEmojiReply) validate(all bool) error {
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
					errors = append(errors, ListEmojiReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListEmojiReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListEmojiReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListEmojiReplyMultiError(errors)
	}

	return nil
}

// ListEmojiReplyMultiError is an error wrapping multiple validation errors
// returned by ListEmojiReply.ValidateAll() if the designated constraints
// aren't met.
type ListEmojiReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEmojiReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEmojiReplyMultiError) AllErrors() []error { return m }

// ListEmojiReplyValidationError is the validation error returned by
// ListEmojiReply.Validate if the designated constraints aren't met.
type ListEmojiReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEmojiReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEmojiReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEmojiReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEmojiReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEmojiReplyValidationError) ErrorName() string { return "ListEmojiReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListEmojiReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEmojiReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEmojiReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEmojiReplyValidationError{}

// Validate checks the field values on ListEmojiItem with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListEmojiItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListEmojiItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListEmojiItemMultiError, or
// nil if none found.
func (m *ListEmojiItem) ValidateAll() error {
	return m.validate(true)
}

func (m *ListEmojiItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ItemId

	// no validation rules for ItemName

	// no validation rules for ItemImage

	// no validation rules for Total

	// no validation rules for ExchangeItemId

	// no validation rules for ExchangeItemName

	// no validation rules for ExchangeItemImage

	// no validation rules for ExchangeItemQuantity

	if len(errors) > 0 {
		return ListEmojiItemMultiError(errors)
	}

	return nil
}

// ListEmojiItemMultiError is an error wrapping multiple validation errors
// returned by ListEmojiItem.ValidateAll() if the designated constraints
// aren't met.
type ListEmojiItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListEmojiItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListEmojiItemMultiError) AllErrors() []error { return m }

// ListEmojiItemValidationError is the validation error returned by
// ListEmojiItem.Validate if the designated constraints aren't met.
type ListEmojiItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListEmojiItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListEmojiItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListEmojiItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListEmojiItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListEmojiItemValidationError) ErrorName() string { return "ListEmojiItemValidationError" }

// Error satisfies the builtin error interface
func (e ListEmojiItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListEmojiItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListEmojiItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListEmojiItemValidationError{}

// Validate checks the field values on UseItemRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UseItemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UseItemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UseItemRequestMultiError,
// or nil if none found.
func (m *UseItemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UseItemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AppId

	// no validation rules for ChannelId

	// no validation rules for UserId

	if m.GetItemId() <= 0 {
		err := UseItemRequestValidationError{
			field:  "ItemId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetQuantity() <= 0 {
		err := UseItemRequestValidationError{
			field:  "Quantity",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UseItemRequestMultiError(errors)
	}

	return nil
}

// UseItemRequestMultiError is an error wrapping multiple validation errors
// returned by UseItemRequest.ValidateAll() if the designated constraints
// aren't met.
type UseItemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UseItemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UseItemRequestMultiError) AllErrors() []error { return m }

// UseItemRequestValidationError is the validation error returned by
// UseItemRequest.Validate if the designated constraints aren't met.
type UseItemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UseItemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UseItemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UseItemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UseItemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UseItemRequestValidationError) ErrorName() string { return "UseItemRequestValidationError" }

// Error satisfies the builtin error interface
func (e UseItemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUseItemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UseItemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UseItemRequestValidationError{}

// Validate checks the field values on UseItemReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UseItemReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UseItemReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UseItemReplyMultiError, or
// nil if none found.
func (m *UseItemReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UseItemReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UseItemReplyMultiError(errors)
	}

	return nil
}

// UseItemReplyMultiError is an error wrapping multiple validation errors
// returned by UseItemReply.ValidateAll() if the designated constraints aren't met.
type UseItemReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UseItemReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UseItemReplyMultiError) AllErrors() []error { return m }

// UseItemReplyValidationError is the validation error returned by
// UseItemReply.Validate if the designated constraints aren't met.
type UseItemReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UseItemReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UseItemReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UseItemReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UseItemReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UseItemReplyValidationError) ErrorName() string { return "UseItemReplyValidationError" }

// Error satisfies the builtin error interface
func (e UseItemReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUseItemReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UseItemReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UseItemReplyValidationError{}
