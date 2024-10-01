// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/integration/v1alpha1/proxy_resource.proto

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

// Validate checks the field values on Resource with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Resource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Resource with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResourceMultiError, or nil
// if none found.
func (m *Resource) ValidateAll() error {
	return m.validate(true)
}

func (m *Resource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Resource.(type) {
	case *Resource_AwsRdsDatabase:
		if v == nil {
			err := ResourceValidationError{
				field:  "Resource",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetAwsRdsDatabase()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResourceValidationError{
						field:  "AwsRdsDatabase",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResourceValidationError{
						field:  "AwsRdsDatabase",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAwsRdsDatabase()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceValidationError{
					field:  "AwsRdsDatabase",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ResourceMultiError(errors)
	}

	return nil
}

// ResourceMultiError is an error wrapping multiple validation errors returned
// by Resource.ValidateAll() if the designated constraints aren't met.
type ResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResourceMultiError) AllErrors() []error { return m }

// ResourceValidationError is the validation error returned by
// Resource.Validate if the designated constraints aren't met.
type ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceValidationError) ErrorName() string { return "ResourceValidationError" }

// Error satisfies the builtin error interface
func (e ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceValidationError{}

// Validate checks the field values on AWSRDSDatabase with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AWSRDSDatabase) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AWSRDSDatabase with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AWSRDSDatabaseMultiError,
// or nil if none found.
func (m *AWSRDSDatabase) ValidateAll() error {
	return m.validate(true)
}

func (m *AWSRDSDatabase) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Engine

	// no validation rules for Region

	// no validation rules for Account

	// no validation rules for InstanceId

	// no validation rules for Database

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AWSRDSDatabaseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AWSRDSDatabaseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AWSRDSDatabaseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Endpoint

	if len(errors) > 0 {
		return AWSRDSDatabaseMultiError(errors)
	}

	return nil
}

// AWSRDSDatabaseMultiError is an error wrapping multiple validation errors
// returned by AWSRDSDatabase.ValidateAll() if the designated constraints
// aren't met.
type AWSRDSDatabaseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AWSRDSDatabaseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AWSRDSDatabaseMultiError) AllErrors() []error { return m }

// AWSRDSDatabaseValidationError is the validation error returned by
// AWSRDSDatabase.Validate if the designated constraints aren't met.
type AWSRDSDatabaseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AWSRDSDatabaseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AWSRDSDatabaseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AWSRDSDatabaseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AWSRDSDatabaseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AWSRDSDatabaseValidationError) ErrorName() string { return "AWSRDSDatabaseValidationError" }

// Error satisfies the builtin error interface
func (e AWSRDSDatabaseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAWSRDSDatabase.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AWSRDSDatabaseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AWSRDSDatabaseValidationError{}

// Validate checks the field values on AWSRDSDatabaseUser with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AWSRDSDatabaseUser) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AWSRDSDatabaseUser with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AWSRDSDatabaseUserMultiError, or nil if none found.
func (m *AWSRDSDatabaseUser) ValidateAll() error {
	return m.validate(true)
}

func (m *AWSRDSDatabaseUser) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Username

	// no validation rules for PasswordSecretsManagerArn

	if m.Endpoint != nil {
		// no validation rules for Endpoint
	}

	if len(errors) > 0 {
		return AWSRDSDatabaseUserMultiError(errors)
	}

	return nil
}

// AWSRDSDatabaseUserMultiError is an error wrapping multiple validation errors
// returned by AWSRDSDatabaseUser.ValidateAll() if the designated constraints
// aren't met.
type AWSRDSDatabaseUserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AWSRDSDatabaseUserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AWSRDSDatabaseUserMultiError) AllErrors() []error { return m }

// AWSRDSDatabaseUserValidationError is the validation error returned by
// AWSRDSDatabaseUser.Validate if the designated constraints aren't met.
type AWSRDSDatabaseUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AWSRDSDatabaseUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AWSRDSDatabaseUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AWSRDSDatabaseUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AWSRDSDatabaseUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AWSRDSDatabaseUserValidationError) ErrorName() string {
	return "AWSRDSDatabaseUserValidationError"
}

// Error satisfies the builtin error interface
func (e AWSRDSDatabaseUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAWSRDSDatabaseUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AWSRDSDatabaseUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AWSRDSDatabaseUserValidationError{}
