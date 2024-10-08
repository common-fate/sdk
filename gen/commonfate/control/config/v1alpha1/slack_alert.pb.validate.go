// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/config/v1alpha1/slack_alert.proto

package configv1alpha1

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

// Validate checks the field values on CreateSlackAlertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSlackAlertRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSlackAlertRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSlackAlertRequestMultiError, or nil if none found.
func (m *CreateSlackAlertRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSlackAlertRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WorkflowId

	// no validation rules for SlackChannelId

	// no validation rules for SlackWorkspaceId

	// no validation rules for UseWebConsoleForApproveAction

	// no validation rules for SendDirectMessagesToApprovers

	// no validation rules for DisableInteractivityHandlers

	if all {
		switch v := interface{}(m.GetNotifyExpiryInSeconds()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateSlackAlertRequestValidationError{
					field:  "NotifyExpiryInSeconds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateSlackAlertRequestValidationError{
					field:  "NotifyExpiryInSeconds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNotifyExpiryInSeconds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSlackAlertRequestValidationError{
				field:  "NotifyExpiryInSeconds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DisableChannelMessageForAutoapprovedRequests

	if m.IntegrationId != nil {
		// no validation rules for IntegrationId
	}

	if len(errors) > 0 {
		return CreateSlackAlertRequestMultiError(errors)
	}

	return nil
}

// CreateSlackAlertRequestMultiError is an error wrapping multiple validation
// errors returned by CreateSlackAlertRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateSlackAlertRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSlackAlertRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSlackAlertRequestMultiError) AllErrors() []error { return m }

// CreateSlackAlertRequestValidationError is the validation error returned by
// CreateSlackAlertRequest.Validate if the designated constraints aren't met.
type CreateSlackAlertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSlackAlertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSlackAlertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSlackAlertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSlackAlertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSlackAlertRequestValidationError) ErrorName() string {
	return "CreateSlackAlertRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSlackAlertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSlackAlertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSlackAlertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSlackAlertRequestValidationError{}

// Validate checks the field values on SlackAlert with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SlackAlert) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SlackAlert with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SlackAlertMultiError, or
// nil if none found.
func (m *SlackAlert) ValidateAll() error {
	return m.validate(true)
}

func (m *SlackAlert) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for WorkflowId

	// no validation rules for SlackChannelId

	// no validation rules for SlackWorkspaceId

	// no validation rules for UseWebConsoleForApproveAction

	// no validation rules for SendDirectMessagesToApprovers

	// no validation rules for DisableInteractivityHandlers

	if all {
		switch v := interface{}(m.GetNotifyExpiryInSeconds()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SlackAlertValidationError{
					field:  "NotifyExpiryInSeconds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SlackAlertValidationError{
					field:  "NotifyExpiryInSeconds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNotifyExpiryInSeconds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SlackAlertValidationError{
				field:  "NotifyExpiryInSeconds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DisableChannelMessageForAutoapprovedRequests

	if m.IntegrationId != nil {
		// no validation rules for IntegrationId
	}

	if len(errors) > 0 {
		return SlackAlertMultiError(errors)
	}

	return nil
}

// SlackAlertMultiError is an error wrapping multiple validation errors
// returned by SlackAlert.ValidateAll() if the designated constraints aren't met.
type SlackAlertMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SlackAlertMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SlackAlertMultiError) AllErrors() []error { return m }

// SlackAlertValidationError is the validation error returned by
// SlackAlert.Validate if the designated constraints aren't met.
type SlackAlertValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SlackAlertValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SlackAlertValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SlackAlertValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SlackAlertValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SlackAlertValidationError) ErrorName() string { return "SlackAlertValidationError" }

// Error satisfies the builtin error interface
func (e SlackAlertValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSlackAlert.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SlackAlertValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SlackAlertValidationError{}

// Validate checks the field values on CreateSlackAlertResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSlackAlertResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSlackAlertResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSlackAlertResponseMultiError, or nil if none found.
func (m *CreateSlackAlertResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSlackAlertResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAlert()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateSlackAlertResponseValidationError{
					field:  "Alert",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateSlackAlertResponseValidationError{
					field:  "Alert",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlert()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSlackAlertResponseValidationError{
				field:  "Alert",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateSlackAlertResponseMultiError(errors)
	}

	return nil
}

// CreateSlackAlertResponseMultiError is an error wrapping multiple validation
// errors returned by CreateSlackAlertResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateSlackAlertResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSlackAlertResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSlackAlertResponseMultiError) AllErrors() []error { return m }

// CreateSlackAlertResponseValidationError is the validation error returned by
// CreateSlackAlertResponse.Validate if the designated constraints aren't met.
type CreateSlackAlertResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSlackAlertResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSlackAlertResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSlackAlertResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSlackAlertResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSlackAlertResponseValidationError) ErrorName() string {
	return "CreateSlackAlertResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSlackAlertResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSlackAlertResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSlackAlertResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSlackAlertResponseValidationError{}

// Validate checks the field values on GetSlackAlertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSlackAlertRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSlackAlertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSlackAlertRequestMultiError, or nil if none found.
func (m *GetSlackAlertRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSlackAlertRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetSlackAlertRequestMultiError(errors)
	}

	return nil
}

// GetSlackAlertRequestMultiError is an error wrapping multiple validation
// errors returned by GetSlackAlertRequest.ValidateAll() if the designated
// constraints aren't met.
type GetSlackAlertRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSlackAlertRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSlackAlertRequestMultiError) AllErrors() []error { return m }

// GetSlackAlertRequestValidationError is the validation error returned by
// GetSlackAlertRequest.Validate if the designated constraints aren't met.
type GetSlackAlertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSlackAlertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSlackAlertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSlackAlertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSlackAlertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSlackAlertRequestValidationError) ErrorName() string {
	return "GetSlackAlertRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSlackAlertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSlackAlertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSlackAlertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSlackAlertRequestValidationError{}

// Validate checks the field values on GetSlackAlertResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSlackAlertResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSlackAlertResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSlackAlertResponseMultiError, or nil if none found.
func (m *GetSlackAlertResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSlackAlertResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAlert()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetSlackAlertResponseValidationError{
					field:  "Alert",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetSlackAlertResponseValidationError{
					field:  "Alert",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlert()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetSlackAlertResponseValidationError{
				field:  "Alert",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetSlackAlertResponseMultiError(errors)
	}

	return nil
}

// GetSlackAlertResponseMultiError is an error wrapping multiple validation
// errors returned by GetSlackAlertResponse.ValidateAll() if the designated
// constraints aren't met.
type GetSlackAlertResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSlackAlertResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSlackAlertResponseMultiError) AllErrors() []error { return m }

// GetSlackAlertResponseValidationError is the validation error returned by
// GetSlackAlertResponse.Validate if the designated constraints aren't met.
type GetSlackAlertResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSlackAlertResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSlackAlertResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSlackAlertResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSlackAlertResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSlackAlertResponseValidationError) ErrorName() string {
	return "GetSlackAlertResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSlackAlertResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSlackAlertResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSlackAlertResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSlackAlertResponseValidationError{}

// Validate checks the field values on UpdateSlackAlertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSlackAlertRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSlackAlertRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSlackAlertRequestMultiError, or nil if none found.
func (m *UpdateSlackAlertRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSlackAlertRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAlert()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSlackAlertRequestValidationError{
					field:  "Alert",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSlackAlertRequestValidationError{
					field:  "Alert",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlert()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSlackAlertRequestValidationError{
				field:  "Alert",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateSlackAlertRequestMultiError(errors)
	}

	return nil
}

// UpdateSlackAlertRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateSlackAlertRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateSlackAlertRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSlackAlertRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSlackAlertRequestMultiError) AllErrors() []error { return m }

// UpdateSlackAlertRequestValidationError is the validation error returned by
// UpdateSlackAlertRequest.Validate if the designated constraints aren't met.
type UpdateSlackAlertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSlackAlertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSlackAlertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSlackAlertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSlackAlertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSlackAlertRequestValidationError) ErrorName() string {
	return "UpdateSlackAlertRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSlackAlertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSlackAlertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSlackAlertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSlackAlertRequestValidationError{}

// Validate checks the field values on UpdateSlackAlertResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSlackAlertResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSlackAlertResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSlackAlertResponseMultiError, or nil if none found.
func (m *UpdateSlackAlertResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSlackAlertResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAlert()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSlackAlertResponseValidationError{
					field:  "Alert",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSlackAlertResponseValidationError{
					field:  "Alert",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlert()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSlackAlertResponseValidationError{
				field:  "Alert",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateSlackAlertResponseMultiError(errors)
	}

	return nil
}

// UpdateSlackAlertResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateSlackAlertResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateSlackAlertResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSlackAlertResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSlackAlertResponseMultiError) AllErrors() []error { return m }

// UpdateSlackAlertResponseValidationError is the validation error returned by
// UpdateSlackAlertResponse.Validate if the designated constraints aren't met.
type UpdateSlackAlertResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSlackAlertResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSlackAlertResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSlackAlertResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSlackAlertResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSlackAlertResponseValidationError) ErrorName() string {
	return "UpdateSlackAlertResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSlackAlertResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSlackAlertResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSlackAlertResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSlackAlertResponseValidationError{}

// Validate checks the field values on DeleteSlackAlertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSlackAlertRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSlackAlertRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSlackAlertRequestMultiError, or nil if none found.
func (m *DeleteSlackAlertRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSlackAlertRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteSlackAlertRequestMultiError(errors)
	}

	return nil
}

// DeleteSlackAlertRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteSlackAlertRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteSlackAlertRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSlackAlertRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSlackAlertRequestMultiError) AllErrors() []error { return m }

// DeleteSlackAlertRequestValidationError is the validation error returned by
// DeleteSlackAlertRequest.Validate if the designated constraints aren't met.
type DeleteSlackAlertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSlackAlertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSlackAlertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSlackAlertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSlackAlertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSlackAlertRequestValidationError) ErrorName() string {
	return "DeleteSlackAlertRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSlackAlertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSlackAlertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSlackAlertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSlackAlertRequestValidationError{}

// Validate checks the field values on DeleteSlackAlertResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSlackAlertResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSlackAlertResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSlackAlertResponseMultiError, or nil if none found.
func (m *DeleteSlackAlertResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSlackAlertResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteSlackAlertResponseMultiError(errors)
	}

	return nil
}

// DeleteSlackAlertResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteSlackAlertResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteSlackAlertResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSlackAlertResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSlackAlertResponseMultiError) AllErrors() []error { return m }

// DeleteSlackAlertResponseValidationError is the validation error returned by
// DeleteSlackAlertResponse.Validate if the designated constraints aren't met.
type DeleteSlackAlertResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSlackAlertResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSlackAlertResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSlackAlertResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSlackAlertResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSlackAlertResponseValidationError) ErrorName() string {
	return "DeleteSlackAlertResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSlackAlertResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSlackAlertResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSlackAlertResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSlackAlertResponseValidationError{}
