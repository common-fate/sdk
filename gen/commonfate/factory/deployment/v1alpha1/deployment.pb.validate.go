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

// Validate checks the field values on DNSRecord with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DNSRecord) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DNSRecord with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DNSRecordMultiError, or nil
// if none found.
func (m *DNSRecord) ValidateAll() error {
	return m.validate(true)
}

func (m *DNSRecord) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Type

	if len(errors) > 0 {
		return DNSRecordMultiError(errors)
	}

	return nil
}

// DNSRecordMultiError is an error wrapping multiple validation errors returned
// by DNSRecord.ValidateAll() if the designated constraints aren't met.
type DNSRecordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DNSRecordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DNSRecordMultiError) AllErrors() []error { return m }

// DNSRecordValidationError is the validation error returned by
// DNSRecord.Validate if the designated constraints aren't met.
type DNSRecordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DNSRecordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DNSRecordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DNSRecordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DNSRecordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DNSRecordValidationError) ErrorName() string { return "DNSRecordValidationError" }

// Error satisfies the builtin error interface
func (e DNSRecordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDNSRecord.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DNSRecordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DNSRecordValidationError{}

// Validate checks the field values on CreateDNSRecordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateDNSRecordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateDNSRecordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateDNSRecordRequestMultiError, or nil if none found.
func (m *CreateDNSRecordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateDNSRecordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRecord()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateDNSRecordRequestValidationError{
					field:  "Record",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateDNSRecordRequestValidationError{
					field:  "Record",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRecord()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateDNSRecordRequestValidationError{
				field:  "Record",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateDNSRecordRequestMultiError(errors)
	}

	return nil
}

// CreateDNSRecordRequestMultiError is an error wrapping multiple validation
// errors returned by CreateDNSRecordRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateDNSRecordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateDNSRecordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateDNSRecordRequestMultiError) AllErrors() []error { return m }

// CreateDNSRecordRequestValidationError is the validation error returned by
// CreateDNSRecordRequest.Validate if the designated constraints aren't met.
type CreateDNSRecordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDNSRecordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDNSRecordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDNSRecordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDNSRecordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDNSRecordRequestValidationError) ErrorName() string {
	return "CreateDNSRecordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDNSRecordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDNSRecordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDNSRecordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDNSRecordRequestValidationError{}

// Validate checks the field values on CreateDNSRecordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateDNSRecordResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateDNSRecordResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateDNSRecordResponseMultiError, or nil if none found.
func (m *CreateDNSRecordResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateDNSRecordResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateDNSRecordResponseMultiError(errors)
	}

	return nil
}

// CreateDNSRecordResponseMultiError is an error wrapping multiple validation
// errors returned by CreateDNSRecordResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateDNSRecordResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateDNSRecordResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateDNSRecordResponseMultiError) AllErrors() []error { return m }

// CreateDNSRecordResponseValidationError is the validation error returned by
// CreateDNSRecordResponse.Validate if the designated constraints aren't met.
type CreateDNSRecordResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDNSRecordResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDNSRecordResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDNSRecordResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDNSRecordResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDNSRecordResponseValidationError) ErrorName() string {
	return "CreateDNSRecordResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDNSRecordResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDNSRecordResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDNSRecordResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDNSRecordResponseValidationError{}

// Validate checks the field values on GetDNSRecordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDNSRecordRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDNSRecordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDNSRecordRequestMultiError, or nil if none found.
func (m *GetDNSRecordRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDNSRecordRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return GetDNSRecordRequestMultiError(errors)
	}

	return nil
}

// GetDNSRecordRequestMultiError is an error wrapping multiple validation
// errors returned by GetDNSRecordRequest.ValidateAll() if the designated
// constraints aren't met.
type GetDNSRecordRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDNSRecordRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDNSRecordRequestMultiError) AllErrors() []error { return m }

// GetDNSRecordRequestValidationError is the validation error returned by
// GetDNSRecordRequest.Validate if the designated constraints aren't met.
type GetDNSRecordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDNSRecordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDNSRecordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDNSRecordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDNSRecordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDNSRecordRequestValidationError) ErrorName() string {
	return "GetDNSRecordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDNSRecordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDNSRecordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDNSRecordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDNSRecordRequestValidationError{}

// Validate checks the field values on GetDNSRecordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDNSRecordResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDNSRecordResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDNSRecordResponseMultiError, or nil if none found.
func (m *GetDNSRecordResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDNSRecordResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRecord()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetDNSRecordResponseValidationError{
					field:  "Record",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetDNSRecordResponseValidationError{
					field:  "Record",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRecord()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetDNSRecordResponseValidationError{
				field:  "Record",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetDNSRecordResponseMultiError(errors)
	}

	return nil
}

// GetDNSRecordResponseMultiError is an error wrapping multiple validation
// errors returned by GetDNSRecordResponse.ValidateAll() if the designated
// constraints aren't met.
type GetDNSRecordResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDNSRecordResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDNSRecordResponseMultiError) AllErrors() []error { return m }

// GetDNSRecordResponseValidationError is the validation error returned by
// GetDNSRecordResponse.Validate if the designated constraints aren't met.
type GetDNSRecordResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDNSRecordResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDNSRecordResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDNSRecordResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDNSRecordResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDNSRecordResponseValidationError) ErrorName() string {
	return "GetDNSRecordResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDNSRecordResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDNSRecordResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDNSRecordResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDNSRecordResponseValidationError{}

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

	// no validation rules for DefaultSubdomainComponent

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
