// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: granted/registry/aws/v1alpha1/aws.proto

package awsv1alpha1

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

// Validate checks the field values on GetProfileForAccountAndRoleRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetProfileForAccountAndRoleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileForAccountAndRoleRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetProfileForAccountAndRoleRequestMultiError, or nil if none found.
func (m *GetProfileForAccountAndRoleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileForAccountAndRoleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccountId

	// no validation rules for AccountName

	// no validation rules for RoleName

	if len(errors) > 0 {
		return GetProfileForAccountAndRoleRequestMultiError(errors)
	}

	return nil
}

// GetProfileForAccountAndRoleRequestMultiError is an error wrapping multiple
// validation errors returned by
// GetProfileForAccountAndRoleRequest.ValidateAll() if the designated
// constraints aren't met.
type GetProfileForAccountAndRoleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileForAccountAndRoleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileForAccountAndRoleRequestMultiError) AllErrors() []error { return m }

// GetProfileForAccountAndRoleRequestValidationError is the validation error
// returned by GetProfileForAccountAndRoleRequest.Validate if the designated
// constraints aren't met.
type GetProfileForAccountAndRoleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileForAccountAndRoleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileForAccountAndRoleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileForAccountAndRoleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileForAccountAndRoleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileForAccountAndRoleRequestValidationError) ErrorName() string {
	return "GetProfileForAccountAndRoleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileForAccountAndRoleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileForAccountAndRoleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileForAccountAndRoleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileForAccountAndRoleRequestValidationError{}

// Validate checks the field values on GetProfileForAccountAndRoleResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetProfileForAccountAndRoleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProfileForAccountAndRoleResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetProfileForAccountAndRoleResponseMultiError, or nil if none found.
func (m *GetProfileForAccountAndRoleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProfileForAccountAndRoleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProfile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProfileForAccountAndRoleResponseValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProfileForAccountAndRoleResponseValidationError{
					field:  "Profile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProfile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProfileForAccountAndRoleResponseValidationError{
				field:  "Profile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetProfileForAccountAndRoleResponseMultiError(errors)
	}

	return nil
}

// GetProfileForAccountAndRoleResponseMultiError is an error wrapping multiple
// validation errors returned by
// GetProfileForAccountAndRoleResponse.ValidateAll() if the designated
// constraints aren't met.
type GetProfileForAccountAndRoleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProfileForAccountAndRoleResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProfileForAccountAndRoleResponseMultiError) AllErrors() []error { return m }

// GetProfileForAccountAndRoleResponseValidationError is the validation error
// returned by GetProfileForAccountAndRoleResponse.Validate if the designated
// constraints aren't met.
type GetProfileForAccountAndRoleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProfileForAccountAndRoleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProfileForAccountAndRoleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProfileForAccountAndRoleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProfileForAccountAndRoleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProfileForAccountAndRoleResponseValidationError) ErrorName() string {
	return "GetProfileForAccountAndRoleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProfileForAccountAndRoleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProfileForAccountAndRoleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProfileForAccountAndRoleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProfileForAccountAndRoleResponseValidationError{}

// Validate checks the field values on ListProfilesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProfilesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProfilesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProfilesRequestMultiError, or nil if none found.
func (m *ListProfilesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProfilesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageToken

	if len(errors) > 0 {
		return ListProfilesRequestMultiError(errors)
	}

	return nil
}

// ListProfilesRequestMultiError is an error wrapping multiple validation
// errors returned by ListProfilesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListProfilesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProfilesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProfilesRequestMultiError) AllErrors() []error { return m }

// ListProfilesRequestValidationError is the validation error returned by
// ListProfilesRequest.Validate if the designated constraints aren't met.
type ListProfilesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProfilesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProfilesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProfilesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProfilesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProfilesRequestValidationError) ErrorName() string {
	return "ListProfilesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProfilesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProfilesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProfilesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProfilesRequestValidationError{}

// Validate checks the field values on ListProfilesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListProfilesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListProfilesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListProfilesResponseMultiError, or nil if none found.
func (m *ListProfilesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListProfilesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetProfiles() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListProfilesResponseValidationError{
						field:  fmt.Sprintf("Profiles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListProfilesResponseValidationError{
						field:  fmt.Sprintf("Profiles[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProfilesResponseValidationError{
					field:  fmt.Sprintf("Profiles[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return ListProfilesResponseMultiError(errors)
	}

	return nil
}

// ListProfilesResponseMultiError is an error wrapping multiple validation
// errors returned by ListProfilesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListProfilesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListProfilesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListProfilesResponseMultiError) AllErrors() []error { return m }

// ListProfilesResponseValidationError is the validation error returned by
// ListProfilesResponse.Validate if the designated constraints aren't met.
type ListProfilesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProfilesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProfilesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProfilesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProfilesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProfilesResponseValidationError) ErrorName() string {
	return "ListProfilesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListProfilesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProfilesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProfilesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProfilesResponseValidationError{}

// Validate checks the field values on Profile with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Profile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Profile with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProfileMultiError, or nil if none found.
func (m *Profile) ValidateAll() error {
	return m.validate(true)
}

func (m *Profile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	for idx, item := range m.GetAttributes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProfileValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProfileValidationError{
						field:  fmt.Sprintf("Attributes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProfileValidationError{
					field:  fmt.Sprintf("Attributes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ProfileMultiError(errors)
	}

	return nil
}

// ProfileMultiError is an error wrapping multiple validation errors returned
// by Profile.ValidateAll() if the designated constraints aren't met.
type ProfileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProfileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProfileMultiError) AllErrors() []error { return m }

// ProfileValidationError is the validation error returned by Profile.Validate
// if the designated constraints aren't met.
type ProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileValidationError) ErrorName() string { return "ProfileValidationError" }

// Error satisfies the builtin error interface
func (e ProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileValidationError{}

// Validate checks the field values on ProfileAttributes with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProfileAttributes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProfileAttributes with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProfileAttributesMultiError, or nil if none found.
func (m *ProfileAttributes) ValidateAll() error {
	return m.validate(true)
}

func (m *ProfileAttributes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	if len(errors) > 0 {
		return ProfileAttributesMultiError(errors)
	}

	return nil
}

// ProfileAttributesMultiError is an error wrapping multiple validation errors
// returned by ProfileAttributes.ValidateAll() if the designated constraints
// aren't met.
type ProfileAttributesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProfileAttributesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProfileAttributesMultiError) AllErrors() []error { return m }

// ProfileAttributesValidationError is the validation error returned by
// ProfileAttributes.Validate if the designated constraints aren't met.
type ProfileAttributesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProfileAttributesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProfileAttributesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProfileAttributesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProfileAttributesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProfileAttributesValidationError) ErrorName() string {
	return "ProfileAttributesValidationError"
}

// Error satisfies the builtin error interface
func (e ProfileAttributesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProfileAttributes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProfileAttributesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProfileAttributesValidationError{}