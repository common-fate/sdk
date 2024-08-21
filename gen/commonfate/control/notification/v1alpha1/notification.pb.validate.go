// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/notification/v1alpha1/notification.proto

package notificationv1alpha1

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

// Validate checks the field values on GetUserNotificationSettingsRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetUserNotificationSettingsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserNotificationSettingsRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetUserNotificationSettingsRequestMultiError, or nil if none found.
func (m *GetUserNotificationSettingsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserNotificationSettingsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetUserNotificationSettingsRequestMultiError(errors)
	}

	return nil
}

// GetUserNotificationSettingsRequestMultiError is an error wrapping multiple
// validation errors returned by
// GetUserNotificationSettingsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUserNotificationSettingsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserNotificationSettingsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserNotificationSettingsRequestMultiError) AllErrors() []error { return m }

// GetUserNotificationSettingsRequestValidationError is the validation error
// returned by GetUserNotificationSettingsRequest.Validate if the designated
// constraints aren't met.
type GetUserNotificationSettingsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserNotificationSettingsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserNotificationSettingsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserNotificationSettingsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserNotificationSettingsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserNotificationSettingsRequestValidationError) ErrorName() string {
	return "GetUserNotificationSettingsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserNotificationSettingsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserNotificationSettingsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserNotificationSettingsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserNotificationSettingsRequestValidationError{}

// Validate checks the field values on GetUserNotificationSettingsResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetUserNotificationSettingsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserNotificationSettingsResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetUserNotificationSettingsResponseMultiError, or nil if none found.
func (m *GetUserNotificationSettingsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserNotificationSettingsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUserNotificationSettings() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetUserNotificationSettingsResponseValidationError{
						field:  fmt.Sprintf("UserNotificationSettings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetUserNotificationSettingsResponseValidationError{
						field:  fmt.Sprintf("UserNotificationSettings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUserNotificationSettingsResponseValidationError{
					field:  fmt.Sprintf("UserNotificationSettings[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetUserNotificationSettingsResponseMultiError(errors)
	}

	return nil
}

// GetUserNotificationSettingsResponseMultiError is an error wrapping multiple
// validation errors returned by
// GetUserNotificationSettingsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetUserNotificationSettingsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserNotificationSettingsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserNotificationSettingsResponseMultiError) AllErrors() []error { return m }

// GetUserNotificationSettingsResponseValidationError is the validation error
// returned by GetUserNotificationSettingsResponse.Validate if the designated
// constraints aren't met.
type GetUserNotificationSettingsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserNotificationSettingsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserNotificationSettingsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserNotificationSettingsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserNotificationSettingsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserNotificationSettingsResponseValidationError) ErrorName() string {
	return "GetUserNotificationSettingsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserNotificationSettingsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserNotificationSettingsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserNotificationSettingsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserNotificationSettingsResponseValidationError{}

// Validate checks the field values on UserNotificationSettings with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserNotificationSettings) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserNotificationSettings with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserNotificationSettingsMultiError, or nil if none found.
func (m *UserNotificationSettings) ValidateAll() error {
	return m.validate(true)
}

func (m *UserNotificationSettings) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Enabled

	if len(errors) > 0 {
		return UserNotificationSettingsMultiError(errors)
	}

	return nil
}

// UserNotificationSettingsMultiError is an error wrapping multiple validation
// errors returned by UserNotificationSettings.ValidateAll() if the designated
// constraints aren't met.
type UserNotificationSettingsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserNotificationSettingsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserNotificationSettingsMultiError) AllErrors() []error { return m }

// UserNotificationSettingsValidationError is the validation error returned by
// UserNotificationSettings.Validate if the designated constraints aren't met.
type UserNotificationSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserNotificationSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserNotificationSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserNotificationSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserNotificationSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserNotificationSettingsValidationError) ErrorName() string {
	return "UserNotificationSettingsValidationError"
}

// Error satisfies the builtin error interface
func (e UserNotificationSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserNotificationSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserNotificationSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserNotificationSettingsValidationError{}

// Validate checks the field values on UpdateUserNotificationSettingsRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UpdateUserNotificationSettingsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserNotificationSettingsRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// UpdateUserNotificationSettingsRequestMultiError, or nil if none found.
func (m *UpdateUserNotificationSettingsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserNotificationSettingsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Enabled

	if len(errors) > 0 {
		return UpdateUserNotificationSettingsRequestMultiError(errors)
	}

	return nil
}

// UpdateUserNotificationSettingsRequestMultiError is an error wrapping
// multiple validation errors returned by
// UpdateUserNotificationSettingsRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateUserNotificationSettingsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserNotificationSettingsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserNotificationSettingsRequestMultiError) AllErrors() []error { return m }

// UpdateUserNotificationSettingsRequestValidationError is the validation error
// returned by UpdateUserNotificationSettingsRequest.Validate if the
// designated constraints aren't met.
type UpdateUserNotificationSettingsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserNotificationSettingsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserNotificationSettingsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserNotificationSettingsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserNotificationSettingsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserNotificationSettingsRequestValidationError) ErrorName() string {
	return "UpdateUserNotificationSettingsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserNotificationSettingsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserNotificationSettingsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserNotificationSettingsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserNotificationSettingsRequestValidationError{}

// Validate checks the field values on UpdateUserNotificationSettingsResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UpdateUserNotificationSettingsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// UpdateUserNotificationSettingsResponse with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// UpdateUserNotificationSettingsResponseMultiError, or nil if none found.
func (m *UpdateUserNotificationSettingsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserNotificationSettingsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUserNotificationSettings()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateUserNotificationSettingsResponseValidationError{
					field:  "UserNotificationSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateUserNotificationSettingsResponseValidationError{
					field:  "UserNotificationSettings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserNotificationSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateUserNotificationSettingsResponseValidationError{
				field:  "UserNotificationSettings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateUserNotificationSettingsResponseMultiError(errors)
	}

	return nil
}

// UpdateUserNotificationSettingsResponseMultiError is an error wrapping
// multiple validation errors returned by
// UpdateUserNotificationSettingsResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateUserNotificationSettingsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserNotificationSettingsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserNotificationSettingsResponseMultiError) AllErrors() []error { return m }

// UpdateUserNotificationSettingsResponseValidationError is the validation
// error returned by UpdateUserNotificationSettingsResponse.Validate if the
// designated constraints aren't met.
type UpdateUserNotificationSettingsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserNotificationSettingsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserNotificationSettingsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserNotificationSettingsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserNotificationSettingsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserNotificationSettingsResponseValidationError) ErrorName() string {
	return "UpdateUserNotificationSettingsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserNotificationSettingsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserNotificationSettingsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserNotificationSettingsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserNotificationSettingsResponseValidationError{}
