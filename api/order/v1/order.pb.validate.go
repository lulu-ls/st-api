// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: order/v1/order.proto

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

// Validate checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderRequestMultiError, or nil if none found.
func (m *CreateOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAppId() <= 0 {
		err := CreateOrderRequestValidationError{
			field:  "AppId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetChannelId() <= 0 {
		err := CreateOrderRequestValidationError{
			field:  "ChannelId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUserId() <= 0 {
		err := CreateOrderRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetOrderAmount() <= 0 {
		err := CreateOrderRequestValidationError{
			field:  "OrderAmount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Ip

	// no validation rules for Description

	// no validation rules for OrderType

	// no validation rules for OutOrderId

	if len(errors) > 0 {
		return CreateOrderRequestMultiError(errors)
	}

	return nil
}

// CreateOrderRequestMultiError is an error wrapping multiple validation errors
// returned by CreateOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderRequestMultiError) AllErrors() []error { return m }

// CreateOrderRequestValidationError is the validation error returned by
// CreateOrderRequest.Validate if the designated constraints aren't met.
type CreateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRequestValidationError) ErrorName() string {
	return "CreateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRequestValidationError{}

// Validate checks the field values on CreateOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderReplyMultiError, or nil if none found.
func (m *CreateOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return CreateOrderReplyMultiError(errors)
	}

	return nil
}

// CreateOrderReplyMultiError is an error wrapping multiple validation errors
// returned by CreateOrderReply.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderReplyMultiError) AllErrors() []error { return m }

// CreateOrderReplyValidationError is the validation error returned by
// CreateOrderReply.Validate if the designated constraints aren't met.
type CreateOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderReplyValidationError) ErrorName() string { return "CreateOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderReplyValidationError{}

// Validate checks the field values on CreateOrderWareItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderWareItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderWareItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderWareItemMultiError, or nil if none found.
func (m *CreateOrderWareItem) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderWareItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WareId

	// no validation rules for WareQuantity

	if len(errors) > 0 {
		return CreateOrderWareItemMultiError(errors)
	}

	return nil
}

// CreateOrderWareItemMultiError is an error wrapping multiple validation
// errors returned by CreateOrderWareItem.ValidateAll() if the designated
// constraints aren't met.
type CreateOrderWareItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderWareItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderWareItemMultiError) AllErrors() []error { return m }

// CreateOrderWareItemValidationError is the validation error returned by
// CreateOrderWareItem.Validate if the designated constraints aren't met.
type CreateOrderWareItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderWareItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderWareItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderWareItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderWareItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderWareItemValidationError) ErrorName() string {
	return "CreateOrderWareItemValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderWareItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderWareItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderWareItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderWareItemValidationError{}

// Validate checks the field values on GetOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrderRequestMultiError, or nil if none found.
func (m *GetOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return GetOrderRequestMultiError(errors)
	}

	return nil
}

// GetOrderRequestMultiError is an error wrapping multiple validation errors
// returned by GetOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type GetOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderRequestMultiError) AllErrors() []error { return m }

// GetOrderRequestValidationError is the validation error returned by
// GetOrderRequest.Validate if the designated constraints aren't met.
type GetOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderRequestValidationError) ErrorName() string { return "GetOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderRequestValidationError{}

// Validate checks the field values on GetOrderReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetOrderReplyMultiError, or
// nil if none found.
func (m *GetOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	// no validation rules for AppId

	// no validation rules for AccessChannelId

	// no validation rules for OrderType

	// no validation rules for OutOrderId

	// no validation rules for Phone

	// no validation rules for OrderAmount

	// no validation rules for OrderTime

	// no validation rules for PaymentAmount

	// no validation rules for PaymentDiscount

	// no validation rules for PaymentTime

	// no validation rules for PaymentInfo

	// no validation rules for Status

	// no validation rules for PaymentStatus

	// no validation rules for ExpireTime

	// no validation rules for UserId

	// no validation rules for Description

	if len(errors) > 0 {
		return GetOrderReplyMultiError(errors)
	}

	return nil
}

// GetOrderReplyMultiError is an error wrapping multiple validation errors
// returned by GetOrderReply.ValidateAll() if the designated constraints
// aren't met.
type GetOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrderReplyMultiError) AllErrors() []error { return m }

// GetOrderReplyValidationError is the validation error returned by
// GetOrderReply.Validate if the designated constraints aren't met.
type GetOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrderReplyValidationError) ErrorName() string { return "GetOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrderReplyValidationError{}

// Validate checks the field values on ListOrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListOrderRequestMultiError, or nil if none found.
func (m *ListOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListOrderRequestMultiError(errors)
	}

	return nil
}

// ListOrderRequestMultiError is an error wrapping multiple validation errors
// returned by ListOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type ListOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrderRequestMultiError) AllErrors() []error { return m }

// ListOrderRequestValidationError is the validation error returned by
// ListOrderRequest.Validate if the designated constraints aren't met.
type ListOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderRequestValidationError) ErrorName() string { return "ListOrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderRequestValidationError{}

// Validate checks the field values on ListOrderReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListOrderReplyMultiError,
// or nil if none found.
func (m *ListOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListOrderReplyMultiError(errors)
	}

	return nil
}

// ListOrderReplyMultiError is an error wrapping multiple validation errors
// returned by ListOrderReply.ValidateAll() if the designated constraints
// aren't met.
type ListOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOrderReplyMultiError) AllErrors() []error { return m }

// ListOrderReplyValidationError is the validation error returned by
// ListOrderReply.Validate if the designated constraints aren't met.
type ListOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOrderReplyValidationError) ErrorName() string { return "ListOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOrderReplyValidationError{}

// Validate checks the field values on PayResultRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PayResultRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayResultRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PayResultRequestMultiError, or nil if none found.
func (m *PayResultRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PayResultRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderStatus

	// no validation rules for OrderId

	// no validation rules for ChannelTransactionId

	// no validation rules for ChannelTxnInfo

	// no validation rules for ChannelTxnTime

	// no validation rules for ChannelTxnAmount

	// no validation rules for ChannelDiscount

	// no validation rules for ChannelSettleDate

	// no validation rules for ChannelSettleAmount

	// no validation rules for OrderRefundStatus

	if len(errors) > 0 {
		return PayResultRequestMultiError(errors)
	}

	return nil
}

// PayResultRequestMultiError is an error wrapping multiple validation errors
// returned by PayResultRequest.ValidateAll() if the designated constraints
// aren't met.
type PayResultRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayResultRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayResultRequestMultiError) AllErrors() []error { return m }

// PayResultRequestValidationError is the validation error returned by
// PayResultRequest.Validate if the designated constraints aren't met.
type PayResultRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayResultRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayResultRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayResultRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayResultRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayResultRequestValidationError) ErrorName() string { return "PayResultRequestValidationError" }

// Error satisfies the builtin error interface
func (e PayResultRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayResultRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayResultRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayResultRequestValidationError{}

// Validate checks the field values on PayResultReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PayResultReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PayResultReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PayResultReplyMultiError,
// or nil if none found.
func (m *PayResultReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PayResultReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PayResultReplyMultiError(errors)
	}

	return nil
}

// PayResultReplyMultiError is an error wrapping multiple validation errors
// returned by PayResultReply.ValidateAll() if the designated constraints
// aren't met.
type PayResultReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayResultReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayResultReplyMultiError) AllErrors() []error { return m }

// PayResultReplyValidationError is the validation error returned by
// PayResultReply.Validate if the designated constraints aren't met.
type PayResultReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayResultReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayResultReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayResultReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayResultReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayResultReplyValidationError) ErrorName() string { return "PayResultReplyValidationError" }

// Error satisfies the builtin error interface
func (e PayResultReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayResultReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayResultReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayResultReplyValidationError{}

// Validate checks the field values on RefundRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RefundRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefundRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RefundRequestMultiError, or
// nil if none found.
func (m *RefundRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RefundRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return RefundRequestMultiError(errors)
	}

	return nil
}

// RefundRequestMultiError is an error wrapping multiple validation errors
// returned by RefundRequest.ValidateAll() if the designated constraints
// aren't met.
type RefundRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefundRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefundRequestMultiError) AllErrors() []error { return m }

// RefundRequestValidationError is the validation error returned by
// RefundRequest.Validate if the designated constraints aren't met.
type RefundRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefundRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefundRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefundRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefundRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefundRequestValidationError) ErrorName() string { return "RefundRequestValidationError" }

// Error satisfies the builtin error interface
func (e RefundRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefundRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefundRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefundRequestValidationError{}

// Validate checks the field values on RefundReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RefundReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefundReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RefundReplyMultiError, or
// nil if none found.
func (m *RefundReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RefundReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RefundReplyMultiError(errors)
	}

	return nil
}

// RefundReplyMultiError is an error wrapping multiple validation errors
// returned by RefundReply.ValidateAll() if the designated constraints aren't met.
type RefundReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefundReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefundReplyMultiError) AllErrors() []error { return m }

// RefundReplyValidationError is the validation error returned by
// RefundReply.Validate if the designated constraints aren't met.
type RefundReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefundReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefundReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefundReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefundReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefundReplyValidationError) ErrorName() string { return "RefundReplyValidationError" }

// Error satisfies the builtin error interface
func (e RefundReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefundReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefundReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefundReplyValidationError{}

// Validate checks the field values on ConfirmPayRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConfirmPayRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfirmPayRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfirmPayRequestMultiError, or nil if none found.
func (m *ConfirmPayRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfirmPayRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return ConfirmPayRequestMultiError(errors)
	}

	return nil
}

// ConfirmPayRequestMultiError is an error wrapping multiple validation errors
// returned by ConfirmPayRequest.ValidateAll() if the designated constraints
// aren't met.
type ConfirmPayRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfirmPayRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfirmPayRequestMultiError) AllErrors() []error { return m }

// ConfirmPayRequestValidationError is the validation error returned by
// ConfirmPayRequest.Validate if the designated constraints aren't met.
type ConfirmPayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfirmPayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfirmPayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfirmPayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfirmPayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfirmPayRequestValidationError) ErrorName() string {
	return "ConfirmPayRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ConfirmPayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfirmPayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfirmPayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfirmPayRequestValidationError{}

// Validate checks the field values on ConfirmPayReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConfirmPayReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfirmPayReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfirmPayReplyMultiError, or nil if none found.
func (m *ConfirmPayReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfirmPayReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ResInfo

	if all {
		switch v := interface{}(m.GetPayInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfirmPayReplyValidationError{
					field:  "PayInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfirmPayReplyValidationError{
					field:  "PayInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPayInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfirmPayReplyValidationError{
				field:  "PayInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConfirmPayReplyMultiError(errors)
	}

	return nil
}

// ConfirmPayReplyMultiError is an error wrapping multiple validation errors
// returned by ConfirmPayReply.ValidateAll() if the designated constraints
// aren't met.
type ConfirmPayReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfirmPayReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfirmPayReplyMultiError) AllErrors() []error { return m }

// ConfirmPayReplyValidationError is the validation error returned by
// ConfirmPayReply.Validate if the designated constraints aren't met.
type ConfirmPayReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfirmPayReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfirmPayReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfirmPayReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfirmPayReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfirmPayReplyValidationError) ErrorName() string { return "ConfirmPayReplyValidationError" }

// Error satisfies the builtin error interface
func (e ConfirmPayReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfirmPayReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfirmPayReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfirmPayReplyValidationError{}
