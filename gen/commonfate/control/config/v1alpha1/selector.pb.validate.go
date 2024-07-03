// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/config/v1alpha1/selector.proto

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

// Validate checks the field values on CreateSelectorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSelectorRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSelectorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSelectorRequestMultiError, or nil if none found.
func (m *CreateSelectorRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSelectorRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSelector()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateSelectorRequestValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateSelectorRequestValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSelector()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSelectorRequestValidationError{
				field:  "Selector",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateSelectorRequestMultiError(errors)
	}

	return nil
}

// CreateSelectorRequestMultiError is an error wrapping multiple validation
// errors returned by CreateSelectorRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateSelectorRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSelectorRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSelectorRequestMultiError) AllErrors() []error { return m }

// CreateSelectorRequestValidationError is the validation error returned by
// CreateSelectorRequest.Validate if the designated constraints aren't met.
type CreateSelectorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSelectorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSelectorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSelectorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSelectorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSelectorRequestValidationError) ErrorName() string {
	return "CreateSelectorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSelectorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSelectorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSelectorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSelectorRequestValidationError{}

// Validate checks the field values on Selector with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Selector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Selector with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SelectorMultiError, or nil
// if none found.
func (m *Selector) ValidateAll() error {
	return m.validate(true)
}

func (m *Selector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for ResourceType

	if all {
		switch v := interface{}(m.GetBelongingTo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SelectorValidationError{
					field:  "BelongingTo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SelectorValidationError{
					field:  "BelongingTo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBelongingTo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SelectorValidationError{
				field:  "BelongingTo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for When

	if len(errors) > 0 {
		return SelectorMultiError(errors)
	}

	return nil
}

// SelectorMultiError is an error wrapping multiple validation errors returned
// by Selector.ValidateAll() if the designated constraints aren't met.
type SelectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SelectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SelectorMultiError) AllErrors() []error { return m }

// SelectorValidationError is the validation error returned by
// Selector.Validate if the designated constraints aren't met.
type SelectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SelectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SelectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SelectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SelectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SelectorValidationError) ErrorName() string { return "SelectorValidationError" }

// Error satisfies the builtin error interface
func (e SelectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSelector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SelectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SelectorValidationError{}

// Validate checks the field values on CreateSelectorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSelectorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSelectorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSelectorResponseMultiError, or nil if none found.
func (m *CreateSelectorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSelectorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSelector()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateSelectorResponseValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateSelectorResponseValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSelector()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSelectorResponseValidationError{
				field:  "Selector",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetDiagnostics() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateSelectorResponseValidationError{
						field:  fmt.Sprintf("Diagnostics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateSelectorResponseValidationError{
						field:  fmt.Sprintf("Diagnostics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateSelectorResponseValidationError{
					field:  fmt.Sprintf("Diagnostics[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateSelectorResponseMultiError(errors)
	}

	return nil
}

// CreateSelectorResponseMultiError is an error wrapping multiple validation
// errors returned by CreateSelectorResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateSelectorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSelectorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSelectorResponseMultiError) AllErrors() []error { return m }

// CreateSelectorResponseValidationError is the validation error returned by
// CreateSelectorResponse.Validate if the designated constraints aren't met.
type CreateSelectorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSelectorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSelectorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSelectorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSelectorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSelectorResponseValidationError) ErrorName() string {
	return "CreateSelectorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSelectorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSelectorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSelectorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSelectorResponseValidationError{}

// Validate checks the field values on GetSelectorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSelectorRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSelectorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSelectorRequestMultiError, or nil if none found.
func (m *GetSelectorRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSelectorRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetSelectorRequestMultiError(errors)
	}

	return nil
}

// GetSelectorRequestMultiError is an error wrapping multiple validation errors
// returned by GetSelectorRequest.ValidateAll() if the designated constraints
// aren't met.
type GetSelectorRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSelectorRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSelectorRequestMultiError) AllErrors() []error { return m }

// GetSelectorRequestValidationError is the validation error returned by
// GetSelectorRequest.Validate if the designated constraints aren't met.
type GetSelectorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSelectorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSelectorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSelectorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSelectorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSelectorRequestValidationError) ErrorName() string {
	return "GetSelectorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSelectorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSelectorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSelectorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSelectorRequestValidationError{}

// Validate checks the field values on GetSelectorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSelectorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSelectorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSelectorResponseMultiError, or nil if none found.
func (m *GetSelectorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSelectorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSelector()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetSelectorResponseValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetSelectorResponseValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSelector()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetSelectorResponseValidationError{
				field:  "Selector",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetSelectorResponseMultiError(errors)
	}

	return nil
}

// GetSelectorResponseMultiError is an error wrapping multiple validation
// errors returned by GetSelectorResponse.ValidateAll() if the designated
// constraints aren't met.
type GetSelectorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSelectorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSelectorResponseMultiError) AllErrors() []error { return m }

// GetSelectorResponseValidationError is the validation error returned by
// GetSelectorResponse.Validate if the designated constraints aren't met.
type GetSelectorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSelectorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSelectorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSelectorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSelectorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSelectorResponseValidationError) ErrorName() string {
	return "GetSelectorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSelectorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSelectorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSelectorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSelectorResponseValidationError{}

// Validate checks the field values on UpdateSelectorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSelectorRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSelectorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSelectorRequestMultiError, or nil if none found.
func (m *UpdateSelectorRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSelectorRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSelector()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSelectorRequestValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSelectorRequestValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSelector()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSelectorRequestValidationError{
				field:  "Selector",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateSelectorRequestMultiError(errors)
	}

	return nil
}

// UpdateSelectorRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateSelectorRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateSelectorRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSelectorRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSelectorRequestMultiError) AllErrors() []error { return m }

// UpdateSelectorRequestValidationError is the validation error returned by
// UpdateSelectorRequest.Validate if the designated constraints aren't met.
type UpdateSelectorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSelectorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSelectorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSelectorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSelectorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSelectorRequestValidationError) ErrorName() string {
	return "UpdateSelectorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSelectorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSelectorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSelectorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSelectorRequestValidationError{}

// Validate checks the field values on UpdateSelectorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSelectorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSelectorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSelectorResponseMultiError, or nil if none found.
func (m *UpdateSelectorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSelectorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSelector()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSelectorResponseValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSelectorResponseValidationError{
					field:  "Selector",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSelector()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSelectorResponseValidationError{
				field:  "Selector",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetDiagnostics() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpdateSelectorResponseValidationError{
						field:  fmt.Sprintf("Diagnostics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpdateSelectorResponseValidationError{
						field:  fmt.Sprintf("Diagnostics[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateSelectorResponseValidationError{
					field:  fmt.Sprintf("Diagnostics[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UpdateSelectorResponseMultiError(errors)
	}

	return nil
}

// UpdateSelectorResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateSelectorResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateSelectorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSelectorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSelectorResponseMultiError) AllErrors() []error { return m }

// UpdateSelectorResponseValidationError is the validation error returned by
// UpdateSelectorResponse.Validate if the designated constraints aren't met.
type UpdateSelectorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSelectorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSelectorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSelectorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSelectorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSelectorResponseValidationError) ErrorName() string {
	return "UpdateSelectorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSelectorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSelectorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSelectorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSelectorResponseValidationError{}

// Validate checks the field values on DeleteSelectorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSelectorRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSelectorRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSelectorRequestMultiError, or nil if none found.
func (m *DeleteSelectorRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSelectorRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteSelectorRequestMultiError(errors)
	}

	return nil
}

// DeleteSelectorRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteSelectorRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteSelectorRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSelectorRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSelectorRequestMultiError) AllErrors() []error { return m }

// DeleteSelectorRequestValidationError is the validation error returned by
// DeleteSelectorRequest.Validate if the designated constraints aren't met.
type DeleteSelectorRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSelectorRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSelectorRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSelectorRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSelectorRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSelectorRequestValidationError) ErrorName() string {
	return "DeleteSelectorRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSelectorRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSelectorRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSelectorRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSelectorRequestValidationError{}

// Validate checks the field values on DeleteSelectorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSelectorResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSelectorResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSelectorResponseMultiError, or nil if none found.
func (m *DeleteSelectorResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSelectorResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteSelectorResponseMultiError(errors)
	}

	return nil
}

// DeleteSelectorResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteSelectorResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteSelectorResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSelectorResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSelectorResponseMultiError) AllErrors() []error { return m }

// DeleteSelectorResponseValidationError is the validation error returned by
// DeleteSelectorResponse.Validate if the designated constraints aren't met.
type DeleteSelectorResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSelectorResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSelectorResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSelectorResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSelectorResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSelectorResponseValidationError) ErrorName() string {
	return "DeleteSelectorResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSelectorResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSelectorResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSelectorResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSelectorResponseValidationError{}
