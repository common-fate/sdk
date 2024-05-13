// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/factory/deployment/v1alpha1/deployment.proto

package deploymentv1alpha1

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

// Validate checks the field values on GetDeploymentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeploymentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeploymentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeploymentRequestMultiError, or nil if none found.
func (m *GetDeploymentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeploymentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetDeploymentRequestMultiError(errors)
	}

	return nil
}

// GetDeploymentRequestMultiError is an error wrapping multiple validation
// errors returned by GetDeploymentRequest.ValidateAll() if the designated
// constraints aren't met.
type GetDeploymentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeploymentRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeploymentRequestMultiError) AllErrors() []error { return m }

// GetDeploymentRequestValidationError is the validation error returned by
// GetDeploymentRequest.Validate if the designated constraints aren't met.
type GetDeploymentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeploymentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeploymentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeploymentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeploymentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeploymentRequestValidationError) ErrorName() string {
	return "GetDeploymentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeploymentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeploymentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeploymentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeploymentRequestValidationError{}

// Validate checks the field values on GetDeploymentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeploymentResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeploymentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeploymentResponseMultiError, or nil if none found.
func (m *GetDeploymentResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeploymentResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDeployment()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetDeploymentResponseValidationError{
					field:  "Deployment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetDeploymentResponseValidationError{
					field:  "Deployment",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeployment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDeploymentResponseValidationError{
				field:  "Deployment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetDeploymentResponseMultiError(errors)
	}

	return nil
}

// GetDeploymentResponseMultiError is an error wrapping multiple validation
// errors returned by GetDeploymentResponse.ValidateAll() if the designated
// constraints aren't met.
type GetDeploymentResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeploymentResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeploymentResponseMultiError) AllErrors() []error { return m }

// GetDeploymentResponseValidationError is the validation error returned by
// GetDeploymentResponse.Validate if the designated constraints aren't met.
type GetDeploymentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeploymentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeploymentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeploymentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeploymentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeploymentResponseValidationError) ErrorName() string {
	return "GetDeploymentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeploymentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeploymentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeploymentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeploymentResponseValidationError{}

// Validate checks the field values on Deployment with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Deployment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Deployment with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeploymentMultiError, or
// nil if none found.
func (m *Deployment) ValidateAll() error {
	return m.validate(true)
}

func (m *Deployment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for DefaultDomain

	if len(errors) > 0 {
		return DeploymentMultiError(errors)
	}

	return nil
}

// DeploymentMultiError is an error wrapping multiple validation errors
// returned by Deployment.ValidateAll() if the designated constraints aren't met.
type DeploymentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeploymentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeploymentMultiError) AllErrors() []error { return m }

// DeploymentValidationError is the validation error returned by
// Deployment.Validate if the designated constraints aren't met.
type DeploymentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentValidationError) ErrorName() string { return "DeploymentValidationError" }

// Error satisfies the builtin error interface
func (e DeploymentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeployment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentValidationError{}

// Validate checks the field values on RegisterDeploymentNameserversRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *RegisterDeploymentNameserversRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterDeploymentNameserversRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// RegisterDeploymentNameserversRequestMultiError, or nil if none found.
func (m *RegisterDeploymentNameserversRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterDeploymentNameserversRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RegisterDeploymentNameserversRequestMultiError(errors)
	}

	return nil
}

// RegisterDeploymentNameserversRequestMultiError is an error wrapping multiple
// validation errors returned by
// RegisterDeploymentNameserversRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterDeploymentNameserversRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterDeploymentNameserversRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterDeploymentNameserversRequestMultiError) AllErrors() []error { return m }

// RegisterDeploymentNameserversRequestValidationError is the validation error
// returned by RegisterDeploymentNameserversRequest.Validate if the designated
// constraints aren't met.
type RegisterDeploymentNameserversRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterDeploymentNameserversRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterDeploymentNameserversRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterDeploymentNameserversRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterDeploymentNameserversRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterDeploymentNameserversRequestValidationError) ErrorName() string {
	return "RegisterDeploymentNameserversRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterDeploymentNameserversRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterDeploymentNameserversRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterDeploymentNameserversRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterDeploymentNameserversRequestValidationError{}

// Validate checks the field values on RegisterDeploymentNameserversResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *RegisterDeploymentNameserversResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterDeploymentNameserversResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// RegisterDeploymentNameserversResponseMultiError, or nil if none found.
func (m *RegisterDeploymentNameserversResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterDeploymentNameserversResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RegisterDeploymentNameserversResponseMultiError(errors)
	}

	return nil
}

// RegisterDeploymentNameserversResponseMultiError is an error wrapping
// multiple validation errors returned by
// RegisterDeploymentNameserversResponse.ValidateAll() if the designated
// constraints aren't met.
type RegisterDeploymentNameserversResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterDeploymentNameserversResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterDeploymentNameserversResponseMultiError) AllErrors() []error { return m }

// RegisterDeploymentNameserversResponseValidationError is the validation error
// returned by RegisterDeploymentNameserversResponse.Validate if the
// designated constraints aren't met.
type RegisterDeploymentNameserversResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterDeploymentNameserversResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterDeploymentNameserversResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterDeploymentNameserversResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterDeploymentNameserversResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterDeploymentNameserversResponseValidationError) ErrorName() string {
	return "RegisterDeploymentNameserversResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterDeploymentNameserversResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterDeploymentNameserversResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterDeploymentNameserversResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterDeploymentNameserversResponseValidationError{}
