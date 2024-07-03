// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/factory/monitoring/v1alpha1/validate.proto

package monitoringv1alpha1

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

// Validate checks the field values on ValidateWriteTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ValidateWriteTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ValidateWriteTokenRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ValidateWriteTokenRequestMultiError, or nil if none found.
func (m *ValidateWriteTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ValidateWriteTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WriteToken

	if len(errors) > 0 {
		return ValidateWriteTokenRequestMultiError(errors)
	}

	return nil
}

// ValidateWriteTokenRequestMultiError is an error wrapping multiple validation
// errors returned by ValidateWriteTokenRequest.ValidateAll() if the
// designated constraints aren't met.
type ValidateWriteTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValidateWriteTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValidateWriteTokenRequestMultiError) AllErrors() []error { return m }

// ValidateWriteTokenRequestValidationError is the validation error returned by
// ValidateWriteTokenRequest.Validate if the designated constraints aren't met.
type ValidateWriteTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateWriteTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateWriteTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateWriteTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateWriteTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateWriteTokenRequestValidationError) ErrorName() string {
	return "ValidateWriteTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ValidateWriteTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateWriteTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateWriteTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateWriteTokenRequestValidationError{}

// Validate checks the field values on ValidateWriteTokenResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ValidateWriteTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ValidateWriteTokenResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ValidateWriteTokenResponseMultiError, or nil if none found.
func (m *ValidateWriteTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ValidateWriteTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DeploymentId

	// no validation rules for AccountId

	if len(errors) > 0 {
		return ValidateWriteTokenResponseMultiError(errors)
	}

	return nil
}

// ValidateWriteTokenResponseMultiError is an error wrapping multiple
// validation errors returned by ValidateWriteTokenResponse.ValidateAll() if
// the designated constraints aren't met.
type ValidateWriteTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValidateWriteTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValidateWriteTokenResponseMultiError) AllErrors() []error { return m }

// ValidateWriteTokenResponseValidationError is the validation error returned
// by ValidateWriteTokenResponse.Validate if the designated constraints aren't met.
type ValidateWriteTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateWriteTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateWriteTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateWriteTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateWriteTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateWriteTokenResponseValidationError) ErrorName() string {
	return "ValidateWriteTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ValidateWriteTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateWriteTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateWriteTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateWriteTokenResponseValidationError{}
