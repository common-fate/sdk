// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/config/v1alpha1/deployment.proto

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

// Validate checks the field values on GetDeploymentConfigRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeploymentConfigRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeploymentConfigRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeploymentConfigRequestMultiError, or nil if none found.
func (m *GetDeploymentConfigRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeploymentConfigRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetDeploymentConfigRequestMultiError(errors)
	}

	return nil
}

// GetDeploymentConfigRequestMultiError is an error wrapping multiple
// validation errors returned by GetDeploymentConfigRequest.ValidateAll() if
// the designated constraints aren't met.
type GetDeploymentConfigRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeploymentConfigRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeploymentConfigRequestMultiError) AllErrors() []error { return m }

// GetDeploymentConfigRequestValidationError is the validation error returned
// by GetDeploymentConfigRequest.Validate if the designated constraints aren't met.
type GetDeploymentConfigRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeploymentConfigRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeploymentConfigRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeploymentConfigRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeploymentConfigRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeploymentConfigRequestValidationError) ErrorName() string {
	return "GetDeploymentConfigRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeploymentConfigRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeploymentConfigRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeploymentConfigRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeploymentConfigRequestValidationError{}

// Validate checks the field values on GetDeploymentConfigResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeploymentConfigResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeploymentConfigResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeploymentConfigResponseMultiError, or nil if none found.
func (m *GetDeploymentConfigResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeploymentConfigResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDeploymentConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetDeploymentConfigResponseValidationError{
					field:  "DeploymentConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetDeploymentConfigResponseValidationError{
					field:  "DeploymentConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeploymentConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDeploymentConfigResponseValidationError{
				field:  "DeploymentConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetDeploymentConfigResponseMultiError(errors)
	}

	return nil
}

// GetDeploymentConfigResponseMultiError is an error wrapping multiple
// validation errors returned by GetDeploymentConfigResponse.ValidateAll() if
// the designated constraints aren't met.
type GetDeploymentConfigResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeploymentConfigResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeploymentConfigResponseMultiError) AllErrors() []error { return m }

// GetDeploymentConfigResponseValidationError is the validation error returned
// by GetDeploymentConfigResponse.Validate if the designated constraints
// aren't met.
type GetDeploymentConfigResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeploymentConfigResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeploymentConfigResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeploymentConfigResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeploymentConfigResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeploymentConfigResponseValidationError) ErrorName() string {
	return "GetDeploymentConfigResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeploymentConfigResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeploymentConfigResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeploymentConfigResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeploymentConfigResponseValidationError{}

// Validate checks the field values on DeploymentConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeploymentConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeploymentConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeploymentConfigMultiError, or nil if none found.
func (m *DeploymentConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DeploymentConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SamlSsoAcsUrl

	// no validation rules for SamlSsoEntityId

	// no validation rules for TerraformOidcClientId

	// no validation rules for ReadOnlyOidcClientId

	// no validation rules for ProvisionerOidcClientId

	// no validation rules for OidcIssuer

	if len(errors) > 0 {
		return DeploymentConfigMultiError(errors)
	}

	return nil
}

// DeploymentConfigMultiError is an error wrapping multiple validation errors
// returned by DeploymentConfig.ValidateAll() if the designated constraints
// aren't met.
type DeploymentConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeploymentConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeploymentConfigMultiError) AllErrors() []error { return m }

// DeploymentConfigValidationError is the validation error returned by
// DeploymentConfig.Validate if the designated constraints aren't met.
type DeploymentConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentConfigValidationError) ErrorName() string { return "DeploymentConfigValidationError" }

// Error satisfies the builtin error interface
func (e DeploymentConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeploymentConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentConfigValidationError{}

// Validate checks the field values on GetDeploymentSecretRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeploymentSecretRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeploymentSecretRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeploymentSecretRequestMultiError, or nil if none found.
func (m *GetDeploymentSecretRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeploymentSecretRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Secret

	if len(errors) > 0 {
		return GetDeploymentSecretRequestMultiError(errors)
	}

	return nil
}

// GetDeploymentSecretRequestMultiError is an error wrapping multiple
// validation errors returned by GetDeploymentSecretRequest.ValidateAll() if
// the designated constraints aren't met.
type GetDeploymentSecretRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeploymentSecretRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeploymentSecretRequestMultiError) AllErrors() []error { return m }

// GetDeploymentSecretRequestValidationError is the validation error returned
// by GetDeploymentSecretRequest.Validate if the designated constraints aren't met.
type GetDeploymentSecretRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeploymentSecretRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeploymentSecretRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeploymentSecretRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeploymentSecretRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeploymentSecretRequestValidationError) ErrorName() string {
	return "GetDeploymentSecretRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeploymentSecretRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeploymentSecretRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeploymentSecretRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeploymentSecretRequestValidationError{}

// Validate checks the field values on GetDeploymentSecretResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeploymentSecretResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeploymentSecretResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeploymentSecretResponseMultiError, or nil if none found.
func (m *GetDeploymentSecretResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeploymentSecretResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Secret

	if len(errors) > 0 {
		return GetDeploymentSecretResponseMultiError(errors)
	}

	return nil
}

// GetDeploymentSecretResponseMultiError is an error wrapping multiple
// validation errors returned by GetDeploymentSecretResponse.ValidateAll() if
// the designated constraints aren't met.
type GetDeploymentSecretResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeploymentSecretResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeploymentSecretResponseMultiError) AllErrors() []error { return m }

// GetDeploymentSecretResponseValidationError is the validation error returned
// by GetDeploymentSecretResponse.Validate if the designated constraints
// aren't met.
type GetDeploymentSecretResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeploymentSecretResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeploymentSecretResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeploymentSecretResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeploymentSecretResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeploymentSecretResponseValidationError) ErrorName() string {
	return "GetDeploymentSecretResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeploymentSecretResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeploymentSecretResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeploymentSecretResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeploymentSecretResponseValidationError{}
