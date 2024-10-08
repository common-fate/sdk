// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/config/v1alpha1/eks_cluster.proto

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

// Validate checks the field values on CreateEKSClusterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateEKSClusterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateEKSClusterRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateEKSClusterRequestMultiError, or nil if none found.
func (m *CreateEKSClusterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateEKSClusterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Region

	// no validation rules for AwsAccountId

	// no validation rules for IntegrationId

	if len(errors) > 0 {
		return CreateEKSClusterRequestMultiError(errors)
	}

	return nil
}

// CreateEKSClusterRequestMultiError is an error wrapping multiple validation
// errors returned by CreateEKSClusterRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateEKSClusterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateEKSClusterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateEKSClusterRequestMultiError) AllErrors() []error { return m }

// CreateEKSClusterRequestValidationError is the validation error returned by
// CreateEKSClusterRequest.Validate if the designated constraints aren't met.
type CreateEKSClusterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEKSClusterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEKSClusterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEKSClusterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEKSClusterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEKSClusterRequestValidationError) ErrorName() string {
	return "CreateEKSClusterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEKSClusterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEKSClusterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEKSClusterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEKSClusterRequestValidationError{}

// Validate checks the field values on EKSCluster with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EKSCluster) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EKSCluster with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EKSClusterMultiError, or
// nil if none found.
func (m *EKSCluster) ValidateAll() error {
	return m.validate(true)
}

func (m *EKSCluster) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Arn

	// no validation rules for Name

	// no validation rules for Region

	// no validation rules for AwsAccountId

	// no validation rules for IntegrationId

	if len(errors) > 0 {
		return EKSClusterMultiError(errors)
	}

	return nil
}

// EKSClusterMultiError is an error wrapping multiple validation errors
// returned by EKSCluster.ValidateAll() if the designated constraints aren't met.
type EKSClusterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EKSClusterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EKSClusterMultiError) AllErrors() []error { return m }

// EKSClusterValidationError is the validation error returned by
// EKSCluster.Validate if the designated constraints aren't met.
type EKSClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EKSClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EKSClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EKSClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EKSClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EKSClusterValidationError) ErrorName() string { return "EKSClusterValidationError" }

// Error satisfies the builtin error interface
func (e EKSClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEKSCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EKSClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EKSClusterValidationError{}

// Validate checks the field values on CreateEKSClusterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateEKSClusterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateEKSClusterResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateEKSClusterResponseMultiError, or nil if none found.
func (m *CreateEKSClusterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateEKSClusterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCluster()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateEKSClusterResponseValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateEKSClusterResponseValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCluster()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateEKSClusterResponseValidationError{
				field:  "Cluster",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateEKSClusterResponseMultiError(errors)
	}

	return nil
}

// CreateEKSClusterResponseMultiError is an error wrapping multiple validation
// errors returned by CreateEKSClusterResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateEKSClusterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateEKSClusterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateEKSClusterResponseMultiError) AllErrors() []error { return m }

// CreateEKSClusterResponseValidationError is the validation error returned by
// CreateEKSClusterResponse.Validate if the designated constraints aren't met.
type CreateEKSClusterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateEKSClusterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateEKSClusterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateEKSClusterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateEKSClusterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateEKSClusterResponseValidationError) ErrorName() string {
	return "CreateEKSClusterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateEKSClusterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateEKSClusterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateEKSClusterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateEKSClusterResponseValidationError{}

// Validate checks the field values on GetEKSClusterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetEKSClusterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEKSClusterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEKSClusterRequestMultiError, or nil if none found.
func (m *GetEKSClusterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEKSClusterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Arn

	if len(errors) > 0 {
		return GetEKSClusterRequestMultiError(errors)
	}

	return nil
}

// GetEKSClusterRequestMultiError is an error wrapping multiple validation
// errors returned by GetEKSClusterRequest.ValidateAll() if the designated
// constraints aren't met.
type GetEKSClusterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEKSClusterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEKSClusterRequestMultiError) AllErrors() []error { return m }

// GetEKSClusterRequestValidationError is the validation error returned by
// GetEKSClusterRequest.Validate if the designated constraints aren't met.
type GetEKSClusterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEKSClusterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEKSClusterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEKSClusterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEKSClusterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEKSClusterRequestValidationError) ErrorName() string {
	return "GetEKSClusterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEKSClusterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEKSClusterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEKSClusterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEKSClusterRequestValidationError{}

// Validate checks the field values on GetEKSClusterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetEKSClusterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEKSClusterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEKSClusterResponseMultiError, or nil if none found.
func (m *GetEKSClusterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEKSClusterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCluster()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetEKSClusterResponseValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetEKSClusterResponseValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCluster()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetEKSClusterResponseValidationError{
				field:  "Cluster",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetEKSClusterResponseMultiError(errors)
	}

	return nil
}

// GetEKSClusterResponseMultiError is an error wrapping multiple validation
// errors returned by GetEKSClusterResponse.ValidateAll() if the designated
// constraints aren't met.
type GetEKSClusterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEKSClusterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEKSClusterResponseMultiError) AllErrors() []error { return m }

// GetEKSClusterResponseValidationError is the validation error returned by
// GetEKSClusterResponse.Validate if the designated constraints aren't met.
type GetEKSClusterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEKSClusterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEKSClusterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEKSClusterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEKSClusterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEKSClusterResponseValidationError) ErrorName() string {
	return "GetEKSClusterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEKSClusterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEKSClusterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEKSClusterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEKSClusterResponseValidationError{}

// Validate checks the field values on UpdateEKSClusterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateEKSClusterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateEKSClusterRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateEKSClusterRequestMultiError, or nil if none found.
func (m *UpdateEKSClusterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateEKSClusterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Arn

	// no validation rules for IntegrationId

	if len(errors) > 0 {
		return UpdateEKSClusterRequestMultiError(errors)
	}

	return nil
}

// UpdateEKSClusterRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateEKSClusterRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateEKSClusterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateEKSClusterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateEKSClusterRequestMultiError) AllErrors() []error { return m }

// UpdateEKSClusterRequestValidationError is the validation error returned by
// UpdateEKSClusterRequest.Validate if the designated constraints aren't met.
type UpdateEKSClusterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateEKSClusterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateEKSClusterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateEKSClusterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateEKSClusterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateEKSClusterRequestValidationError) ErrorName() string {
	return "UpdateEKSClusterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateEKSClusterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateEKSClusterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateEKSClusterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateEKSClusterRequestValidationError{}

// Validate checks the field values on UpdateEKSClusterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateEKSClusterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateEKSClusterResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateEKSClusterResponseMultiError, or nil if none found.
func (m *UpdateEKSClusterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateEKSClusterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCluster()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateEKSClusterResponseValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateEKSClusterResponseValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCluster()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateEKSClusterResponseValidationError{
				field:  "Cluster",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateEKSClusterResponseMultiError(errors)
	}

	return nil
}

// UpdateEKSClusterResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateEKSClusterResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateEKSClusterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateEKSClusterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateEKSClusterResponseMultiError) AllErrors() []error { return m }

// UpdateEKSClusterResponseValidationError is the validation error returned by
// UpdateEKSClusterResponse.Validate if the designated constraints aren't met.
type UpdateEKSClusterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateEKSClusterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateEKSClusterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateEKSClusterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateEKSClusterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateEKSClusterResponseValidationError) ErrorName() string {
	return "UpdateEKSClusterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateEKSClusterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateEKSClusterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateEKSClusterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateEKSClusterResponseValidationError{}

// Validate checks the field values on DeleteEKSClusterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteEKSClusterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteEKSClusterRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteEKSClusterRequestMultiError, or nil if none found.
func (m *DeleteEKSClusterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteEKSClusterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Arn

	if len(errors) > 0 {
		return DeleteEKSClusterRequestMultiError(errors)
	}

	return nil
}

// DeleteEKSClusterRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteEKSClusterRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteEKSClusterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteEKSClusterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteEKSClusterRequestMultiError) AllErrors() []error { return m }

// DeleteEKSClusterRequestValidationError is the validation error returned by
// DeleteEKSClusterRequest.Validate if the designated constraints aren't met.
type DeleteEKSClusterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEKSClusterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEKSClusterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEKSClusterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEKSClusterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEKSClusterRequestValidationError) ErrorName() string {
	return "DeleteEKSClusterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEKSClusterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEKSClusterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEKSClusterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEKSClusterRequestValidationError{}

// Validate checks the field values on DeleteEKSClusterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteEKSClusterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteEKSClusterResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteEKSClusterResponseMultiError, or nil if none found.
func (m *DeleteEKSClusterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteEKSClusterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Arn

	if len(errors) > 0 {
		return DeleteEKSClusterResponseMultiError(errors)
	}

	return nil
}

// DeleteEKSClusterResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteEKSClusterResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteEKSClusterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteEKSClusterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteEKSClusterResponseMultiError) AllErrors() []error { return m }

// DeleteEKSClusterResponseValidationError is the validation error returned by
// DeleteEKSClusterResponse.Validate if the designated constraints aren't met.
type DeleteEKSClusterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteEKSClusterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteEKSClusterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteEKSClusterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteEKSClusterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteEKSClusterResponseValidationError) ErrorName() string {
	return "DeleteEKSClusterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteEKSClusterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteEKSClusterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteEKSClusterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteEKSClusterResponseValidationError{}
