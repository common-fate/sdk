// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/oauth/v1alpha1/oauth.proto

package oauthv1alpha1

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

// Validate checks the field values on CreatePagerDutyIntegrationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CreatePagerDutyIntegrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePagerDutyIntegrationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CreatePagerDutyIntegrationRequestMultiError, or nil if none found.
func (m *CreatePagerDutyIntegrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePagerDutyIntegrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreatePagerDutyIntegrationRequestMultiError(errors)
	}

	return nil
}

// CreatePagerDutyIntegrationRequestMultiError is an error wrapping multiple
// validation errors returned by
// CreatePagerDutyIntegrationRequest.ValidateAll() if the designated
// constraints aren't met.
type CreatePagerDutyIntegrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePagerDutyIntegrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePagerDutyIntegrationRequestMultiError) AllErrors() []error { return m }

// CreatePagerDutyIntegrationRequestValidationError is the validation error
// returned by CreatePagerDutyIntegrationRequest.Validate if the designated
// constraints aren't met.
type CreatePagerDutyIntegrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePagerDutyIntegrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePagerDutyIntegrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePagerDutyIntegrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePagerDutyIntegrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePagerDutyIntegrationRequestValidationError) ErrorName() string {
	return "CreatePagerDutyIntegrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePagerDutyIntegrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePagerDutyIntegrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePagerDutyIntegrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePagerDutyIntegrationRequestValidationError{}

// Validate checks the field values on CreatePagerDutyIntegrationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CreatePagerDutyIntegrationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePagerDutyIntegrationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// CreatePagerDutyIntegrationResponseMultiError, or nil if none found.
func (m *CreatePagerDutyIntegrationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePagerDutyIntegrationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uri

	if len(errors) > 0 {
		return CreatePagerDutyIntegrationResponseMultiError(errors)
	}

	return nil
}

// CreatePagerDutyIntegrationResponseMultiError is an error wrapping multiple
// validation errors returned by
// CreatePagerDutyIntegrationResponse.ValidateAll() if the designated
// constraints aren't met.
type CreatePagerDutyIntegrationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePagerDutyIntegrationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePagerDutyIntegrationResponseMultiError) AllErrors() []error { return m }

// CreatePagerDutyIntegrationResponseValidationError is the validation error
// returned by CreatePagerDutyIntegrationResponse.Validate if the designated
// constraints aren't met.
type CreatePagerDutyIntegrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePagerDutyIntegrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePagerDutyIntegrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePagerDutyIntegrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePagerDutyIntegrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePagerDutyIntegrationResponseValidationError) ErrorName() string {
	return "CreatePagerDutyIntegrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePagerDutyIntegrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePagerDutyIntegrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePagerDutyIntegrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePagerDutyIntegrationResponseValidationError{}

// Validate checks the field values on GetPagerDutyIntegrationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPagerDutyIntegrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPagerDutyIntegrationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetPagerDutyIntegrationRequestMultiError, or nil if none found.
func (m *GetPagerDutyIntegrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPagerDutyIntegrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetPagerDutyIntegrationRequestMultiError(errors)
	}

	return nil
}

// GetPagerDutyIntegrationRequestMultiError is an error wrapping multiple
// validation errors returned by GetPagerDutyIntegrationRequest.ValidateAll()
// if the designated constraints aren't met.
type GetPagerDutyIntegrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPagerDutyIntegrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPagerDutyIntegrationRequestMultiError) AllErrors() []error { return m }

// GetPagerDutyIntegrationRequestValidationError is the validation error
// returned by GetPagerDutyIntegrationRequest.Validate if the designated
// constraints aren't met.
type GetPagerDutyIntegrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPagerDutyIntegrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPagerDutyIntegrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPagerDutyIntegrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPagerDutyIntegrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPagerDutyIntegrationRequestValidationError) ErrorName() string {
	return "GetPagerDutyIntegrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPagerDutyIntegrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPagerDutyIntegrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPagerDutyIntegrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPagerDutyIntegrationRequestValidationError{}

// Validate checks the field values on GetPagerDutyIntegrationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPagerDutyIntegrationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPagerDutyIntegrationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetPagerDutyIntegrationResponseMultiError, or nil if none found.
func (m *GetPagerDutyIntegrationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPagerDutyIntegrationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Active

	if len(errors) > 0 {
		return GetPagerDutyIntegrationResponseMultiError(errors)
	}

	return nil
}

// GetPagerDutyIntegrationResponseMultiError is an error wrapping multiple
// validation errors returned by GetPagerDutyIntegrationResponse.ValidateAll()
// if the designated constraints aren't met.
type GetPagerDutyIntegrationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPagerDutyIntegrationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPagerDutyIntegrationResponseMultiError) AllErrors() []error { return m }

// GetPagerDutyIntegrationResponseValidationError is the validation error
// returned by GetPagerDutyIntegrationResponse.Validate if the designated
// constraints aren't met.
type GetPagerDutyIntegrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPagerDutyIntegrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPagerDutyIntegrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPagerDutyIntegrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPagerDutyIntegrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPagerDutyIntegrationResponseValidationError) ErrorName() string {
	return "GetPagerDutyIntegrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPagerDutyIntegrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPagerDutyIntegrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPagerDutyIntegrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPagerDutyIntegrationResponseValidationError{}

// Validate checks the field values on RemovePagerDutyIntegrationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *RemovePagerDutyIntegrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemovePagerDutyIntegrationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// RemovePagerDutyIntegrationRequestMultiError, or nil if none found.
func (m *RemovePagerDutyIntegrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemovePagerDutyIntegrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RemovePagerDutyIntegrationRequestMultiError(errors)
	}

	return nil
}

// RemovePagerDutyIntegrationRequestMultiError is an error wrapping multiple
// validation errors returned by
// RemovePagerDutyIntegrationRequest.ValidateAll() if the designated
// constraints aren't met.
type RemovePagerDutyIntegrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemovePagerDutyIntegrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemovePagerDutyIntegrationRequestMultiError) AllErrors() []error { return m }

// RemovePagerDutyIntegrationRequestValidationError is the validation error
// returned by RemovePagerDutyIntegrationRequest.Validate if the designated
// constraints aren't met.
type RemovePagerDutyIntegrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemovePagerDutyIntegrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemovePagerDutyIntegrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemovePagerDutyIntegrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemovePagerDutyIntegrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemovePagerDutyIntegrationRequestValidationError) ErrorName() string {
	return "RemovePagerDutyIntegrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemovePagerDutyIntegrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemovePagerDutyIntegrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemovePagerDutyIntegrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemovePagerDutyIntegrationRequestValidationError{}

// Validate checks the field values on RemovePagerDutyIntegrationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *RemovePagerDutyIntegrationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemovePagerDutyIntegrationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// RemovePagerDutyIntegrationResponseMultiError, or nil if none found.
func (m *RemovePagerDutyIntegrationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RemovePagerDutyIntegrationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return RemovePagerDutyIntegrationResponseMultiError(errors)
	}

	return nil
}

// RemovePagerDutyIntegrationResponseMultiError is an error wrapping multiple
// validation errors returned by
// RemovePagerDutyIntegrationResponse.ValidateAll() if the designated
// constraints aren't met.
type RemovePagerDutyIntegrationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemovePagerDutyIntegrationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemovePagerDutyIntegrationResponseMultiError) AllErrors() []error { return m }

// RemovePagerDutyIntegrationResponseValidationError is the validation error
// returned by RemovePagerDutyIntegrationResponse.Validate if the designated
// constraints aren't met.
type RemovePagerDutyIntegrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemovePagerDutyIntegrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemovePagerDutyIntegrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemovePagerDutyIntegrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemovePagerDutyIntegrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemovePagerDutyIntegrationResponseValidationError) ErrorName() string {
	return "RemovePagerDutyIntegrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemovePagerDutyIntegrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemovePagerDutyIntegrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemovePagerDutyIntegrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemovePagerDutyIntegrationResponseValidationError{}

// Validate checks the field values on CreateSlackIntegrationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSlackIntegrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSlackIntegrationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreateSlackIntegrationRequestMultiError, or nil if none found.
func (m *CreateSlackIntegrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSlackIntegrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateSlackIntegrationRequestMultiError(errors)
	}

	return nil
}

// CreateSlackIntegrationRequestMultiError is an error wrapping multiple
// validation errors returned by CreateSlackIntegrationRequest.ValidateAll()
// if the designated constraints aren't met.
type CreateSlackIntegrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSlackIntegrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSlackIntegrationRequestMultiError) AllErrors() []error { return m }

// CreateSlackIntegrationRequestValidationError is the validation error
// returned by CreateSlackIntegrationRequest.Validate if the designated
// constraints aren't met.
type CreateSlackIntegrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSlackIntegrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSlackIntegrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSlackIntegrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSlackIntegrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSlackIntegrationRequestValidationError) ErrorName() string {
	return "CreateSlackIntegrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSlackIntegrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSlackIntegrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSlackIntegrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSlackIntegrationRequestValidationError{}

// Validate checks the field values on CreateSlackIntegrationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSlackIntegrationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSlackIntegrationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreateSlackIntegrationResponseMultiError, or nil if none found.
func (m *CreateSlackIntegrationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSlackIntegrationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uri

	if len(errors) > 0 {
		return CreateSlackIntegrationResponseMultiError(errors)
	}

	return nil
}

// CreateSlackIntegrationResponseMultiError is an error wrapping multiple
// validation errors returned by CreateSlackIntegrationResponse.ValidateAll()
// if the designated constraints aren't met.
type CreateSlackIntegrationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSlackIntegrationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSlackIntegrationResponseMultiError) AllErrors() []error { return m }

// CreateSlackIntegrationResponseValidationError is the validation error
// returned by CreateSlackIntegrationResponse.Validate if the designated
// constraints aren't met.
type CreateSlackIntegrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSlackIntegrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSlackIntegrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSlackIntegrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSlackIntegrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSlackIntegrationResponseValidationError) ErrorName() string {
	return "CreateSlackIntegrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSlackIntegrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSlackIntegrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSlackIntegrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSlackIntegrationResponseValidationError{}

// Validate checks the field values on GetSlackIntegrationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSlackIntegrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSlackIntegrationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSlackIntegrationRequestMultiError, or nil if none found.
func (m *GetSlackIntegrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSlackIntegrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetSlackIntegrationRequestMultiError(errors)
	}

	return nil
}

// GetSlackIntegrationRequestMultiError is an error wrapping multiple
// validation errors returned by GetSlackIntegrationRequest.ValidateAll() if
// the designated constraints aren't met.
type GetSlackIntegrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSlackIntegrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSlackIntegrationRequestMultiError) AllErrors() []error { return m }

// GetSlackIntegrationRequestValidationError is the validation error returned
// by GetSlackIntegrationRequest.Validate if the designated constraints aren't met.
type GetSlackIntegrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSlackIntegrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSlackIntegrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSlackIntegrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSlackIntegrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSlackIntegrationRequestValidationError) ErrorName() string {
	return "GetSlackIntegrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSlackIntegrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSlackIntegrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSlackIntegrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSlackIntegrationRequestValidationError{}

// Validate checks the field values on GetSlackIntegrationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSlackIntegrationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSlackIntegrationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSlackIntegrationResponseMultiError, or nil if none found.
func (m *GetSlackIntegrationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSlackIntegrationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Active

	if len(errors) > 0 {
		return GetSlackIntegrationResponseMultiError(errors)
	}

	return nil
}

// GetSlackIntegrationResponseMultiError is an error wrapping multiple
// validation errors returned by GetSlackIntegrationResponse.ValidateAll() if
// the designated constraints aren't met.
type GetSlackIntegrationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSlackIntegrationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSlackIntegrationResponseMultiError) AllErrors() []error { return m }

// GetSlackIntegrationResponseValidationError is the validation error returned
// by GetSlackIntegrationResponse.Validate if the designated constraints
// aren't met.
type GetSlackIntegrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSlackIntegrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSlackIntegrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSlackIntegrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSlackIntegrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSlackIntegrationResponseValidationError) ErrorName() string {
	return "GetSlackIntegrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSlackIntegrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSlackIntegrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSlackIntegrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSlackIntegrationResponseValidationError{}

// Validate checks the field values on RemoveSlackIntegrationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveSlackIntegrationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveSlackIntegrationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// RemoveSlackIntegrationRequestMultiError, or nil if none found.
func (m *RemoveSlackIntegrationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveSlackIntegrationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RemoveSlackIntegrationRequestMultiError(errors)
	}

	return nil
}

// RemoveSlackIntegrationRequestMultiError is an error wrapping multiple
// validation errors returned by RemoveSlackIntegrationRequest.ValidateAll()
// if the designated constraints aren't met.
type RemoveSlackIntegrationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveSlackIntegrationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveSlackIntegrationRequestMultiError) AllErrors() []error { return m }

// RemoveSlackIntegrationRequestValidationError is the validation error
// returned by RemoveSlackIntegrationRequest.Validate if the designated
// constraints aren't met.
type RemoveSlackIntegrationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSlackIntegrationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSlackIntegrationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSlackIntegrationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSlackIntegrationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSlackIntegrationRequestValidationError) ErrorName() string {
	return "RemoveSlackIntegrationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSlackIntegrationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSlackIntegrationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSlackIntegrationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSlackIntegrationRequestValidationError{}

// Validate checks the field values on RemoveSlackIntegrationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RemoveSlackIntegrationResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RemoveSlackIntegrationResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// RemoveSlackIntegrationResponseMultiError, or nil if none found.
func (m *RemoveSlackIntegrationResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RemoveSlackIntegrationResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return RemoveSlackIntegrationResponseMultiError(errors)
	}

	return nil
}

// RemoveSlackIntegrationResponseMultiError is an error wrapping multiple
// validation errors returned by RemoveSlackIntegrationResponse.ValidateAll()
// if the designated constraints aren't met.
type RemoveSlackIntegrationResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RemoveSlackIntegrationResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RemoveSlackIntegrationResponseMultiError) AllErrors() []error { return m }

// RemoveSlackIntegrationResponseValidationError is the validation error
// returned by RemoveSlackIntegrationResponse.Validate if the designated
// constraints aren't met.
type RemoveSlackIntegrationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSlackIntegrationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSlackIntegrationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSlackIntegrationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSlackIntegrationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSlackIntegrationResponseValidationError) ErrorName() string {
	return "RemoveSlackIntegrationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSlackIntegrationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSlackIntegrationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSlackIntegrationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSlackIntegrationResponseValidationError{}
