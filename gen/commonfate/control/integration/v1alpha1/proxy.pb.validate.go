// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/integration/v1alpha1/proxy.proto

package integrationv1alpha1

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

// Validate checks the field values on RegisterProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterProxyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterProxyRequestMultiError, or nil if none found.
func (m *RegisterProxyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterProxyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IntegrationId

	for idx, item := range m.GetResources() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegisterProxyRequestValidationError{
						field:  fmt.Sprintf("Resources[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegisterProxyRequestValidationError{
						field:  fmt.Sprintf("Resources[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterProxyRequestValidationError{
					field:  fmt.Sprintf("Resources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch v := m.InstanceConfig.(type) {
	case *RegisterProxyRequest_AwsEcsProxyInstanceConfig:
		if v == nil {
			err := RegisterProxyRequestValidationError{
				field:  "InstanceConfig",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetAwsEcsProxyInstanceConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegisterProxyRequestValidationError{
						field:  "AwsEcsProxyInstanceConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegisterProxyRequestValidationError{
						field:  "AwsEcsProxyInstanceConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAwsEcsProxyInstanceConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegisterProxyRequestValidationError{
					field:  "AwsEcsProxyInstanceConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return RegisterProxyRequestMultiError(errors)
	}

	return nil
}

// RegisterProxyRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterProxyRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterProxyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterProxyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterProxyRequestMultiError) AllErrors() []error { return m }

// RegisterProxyRequestValidationError is the validation error returned by
// RegisterProxyRequest.Validate if the designated constraints aren't met.
type RegisterProxyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterProxyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterProxyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterProxyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterProxyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterProxyRequestValidationError) ErrorName() string {
	return "RegisterProxyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterProxyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterProxyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterProxyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterProxyRequestValidationError{}

// Validate checks the field values on RegisterProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterProxyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterProxyResponseMultiError, or nil if none found.
func (m *RegisterProxyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterProxyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return RegisterProxyResponseMultiError(errors)
	}

	return nil
}

// RegisterProxyResponseMultiError is an error wrapping multiple validation
// errors returned by RegisterProxyResponse.ValidateAll() if the designated
// constraints aren't met.
type RegisterProxyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterProxyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterProxyResponseMultiError) AllErrors() []error { return m }

// RegisterProxyResponseValidationError is the validation error returned by
// RegisterProxyResponse.Validate if the designated constraints aren't met.
type RegisterProxyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterProxyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterProxyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterProxyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterProxyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterProxyResponseValidationError) ErrorName() string {
	return "RegisterProxyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterProxyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterProxyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterProxyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterProxyResponseValidationError{}

// Validate checks the field values on GetProxyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProxyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProxyRequestMultiError, or nil if none found.
func (m *GetProxyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProxyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetProxyRequestMultiError(errors)
	}

	return nil
}

// GetProxyRequestMultiError is an error wrapping multiple validation errors
// returned by GetProxyRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProxyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProxyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProxyRequestMultiError) AllErrors() []error { return m }

// GetProxyRequestValidationError is the validation error returned by
// GetProxyRequest.Validate if the designated constraints aren't met.
type GetProxyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProxyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProxyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProxyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProxyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProxyRequestValidationError) ErrorName() string { return "GetProxyRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetProxyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProxyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProxyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProxyRequestValidationError{}

// Validate checks the field values on GetProxyResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProxyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProxyResponseMultiError, or nil if none found.
func (m *GetProxyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProxyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for AwsRegion

	// no validation rules for AwsAccountId

	// no validation rules for CommonFateAwsAccountId

	// no validation rules for AwsPartition

	// no validation rules for SubnetIds

	// no validation rules for VpcId

	// no validation rules for EcsClusterId

	// no validation rules for EcsClusterName

	// no validation rules for AuthIssuer

	// no validation rules for ProxyServiceClientId

	// no validation rules for ProxyServiceClientSecret

	if len(errors) > 0 {
		return GetProxyResponseMultiError(errors)
	}

	return nil
}

// GetProxyResponseMultiError is an error wrapping multiple validation errors
// returned by GetProxyResponse.ValidateAll() if the designated constraints
// aren't met.
type GetProxyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProxyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProxyResponseMultiError) AllErrors() []error { return m }

// GetProxyResponseValidationError is the validation error returned by
// GetProxyResponse.Validate if the designated constraints aren't met.
type GetProxyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProxyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProxyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProxyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProxyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProxyResponseValidationError) ErrorName() string { return "GetProxyResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetProxyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProxyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProxyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProxyResponseValidationError{}

// Validate checks the field values on UpdateProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProxyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProxyRequestMultiError, or nil if none found.
func (m *UpdateProxyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProxyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for AwsRegion

	// no validation rules for AwsAccountId

	// no validation rules for CommonFateAwsAccountId

	// no validation rules for AwsPartition

	// no validation rules for SubnetIds

	// no validation rules for VpcId

	// no validation rules for EcsClusterId

	// no validation rules for EcsClusterName

	// no validation rules for AuthIssuer

	// no validation rules for ProxyServiceClientId

	// no validation rules for ProxyServiceClientSecret

	if len(errors) > 0 {
		return UpdateProxyRequestMultiError(errors)
	}

	return nil
}

// UpdateProxyRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateProxyRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateProxyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProxyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProxyRequestMultiError) AllErrors() []error { return m }

// UpdateProxyRequestValidationError is the validation error returned by
// UpdateProxyRequest.Validate if the designated constraints aren't met.
type UpdateProxyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProxyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProxyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProxyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProxyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProxyRequestValidationError) ErrorName() string {
	return "UpdateProxyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProxyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProxyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProxyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProxyRequestValidationError{}

// Validate checks the field values on UpdateProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateProxyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateProxyResponseMultiError, or nil if none found.
func (m *UpdateProxyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateProxyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for IntegrationId

	for idx, item := range m.GetResources() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpdateProxyResponseValidationError{
						field:  fmt.Sprintf("Resources[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpdateProxyResponseValidationError{
						field:  fmt.Sprintf("Resources[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateProxyResponseValidationError{
					field:  fmt.Sprintf("Resources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch v := m.InstanceConfig.(type) {
	case *UpdateProxyResponse_AwsEcsProxyInstanceConfig:
		if v == nil {
			err := UpdateProxyResponseValidationError{
				field:  "InstanceConfig",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetAwsEcsProxyInstanceConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpdateProxyResponseValidationError{
						field:  "AwsEcsProxyInstanceConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpdateProxyResponseValidationError{
						field:  "AwsEcsProxyInstanceConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAwsEcsProxyInstanceConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpdateProxyResponseValidationError{
					field:  "AwsEcsProxyInstanceConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return UpdateProxyResponseMultiError(errors)
	}

	return nil
}

// UpdateProxyResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateProxyResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateProxyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateProxyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateProxyResponseMultiError) AllErrors() []error { return m }

// UpdateProxyResponseValidationError is the validation error returned by
// UpdateProxyResponse.Validate if the designated constraints aren't met.
type UpdateProxyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProxyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProxyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProxyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProxyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProxyResponseValidationError) ErrorName() string {
	return "UpdateProxyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProxyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProxyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProxyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProxyResponseValidationError{}

// Validate checks the field values on DeleteProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProxyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProxyRequestMultiError, or nil if none found.
func (m *DeleteProxyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProxyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteProxyRequestMultiError(errors)
	}

	return nil
}

// DeleteProxyRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteProxyRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteProxyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProxyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProxyRequestMultiError) AllErrors() []error { return m }

// DeleteProxyRequestValidationError is the validation error returned by
// DeleteProxyRequest.Validate if the designated constraints aren't met.
type DeleteProxyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProxyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProxyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProxyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProxyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProxyRequestValidationError) ErrorName() string {
	return "DeleteProxyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProxyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProxyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProxyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProxyRequestValidationError{}

// Validate checks the field values on DeleteProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteProxyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteProxyResponseMultiError, or nil if none found.
func (m *DeleteProxyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteProxyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteProxyResponseMultiError(errors)
	}

	return nil
}

// DeleteProxyResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteProxyResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteProxyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteProxyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteProxyResponseMultiError) AllErrors() []error { return m }

// DeleteProxyResponseValidationError is the validation error returned by
// DeleteProxyResponse.Validate if the designated constraints aren't met.
type DeleteProxyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProxyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProxyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProxyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProxyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProxyResponseValidationError) ErrorName() string {
	return "DeleteProxyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProxyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProxyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProxyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProxyResponseValidationError{}
