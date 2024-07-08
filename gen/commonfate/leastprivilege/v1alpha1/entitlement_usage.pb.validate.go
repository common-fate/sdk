// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/leastprivilege/v1alpha1/entitlement_usage.proto

package leastprivilegev1alpha1

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

// Validate checks the field values on EntitlementUsage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *EntitlementUsage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EntitlementUsage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EntitlementUsageMultiError, or nil if none found.
func (m *EntitlementUsage) ValidateAll() error {
	return m.validate(true)
}

func (m *EntitlementUsage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTarget()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntitlementUsageValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntitlementUsageValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTarget()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntitlementUsageValidationError{
				field:  "Target",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRole()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntitlementUsageValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntitlementUsageValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntitlementUsageValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPrincipal()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EntitlementUsageValidationError{
					field:  "Principal",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EntitlementUsageValidationError{
					field:  "Principal",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrincipal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EntitlementUsageValidationError{
				field:  "Principal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Usage

	// no validation rules for UsageReason

	for idx, item := range m.GetAccessPaths() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntitlementUsageValidationError{
						field:  fmt.Sprintf("AccessPaths[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntitlementUsageValidationError{
						field:  fmt.Sprintf("AccessPaths[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntitlementUsageValidationError{
					field:  fmt.Sprintf("AccessPaths[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.LastUsed != nil {

		if all {
			switch v := interface{}(m.GetLastUsed()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntitlementUsageValidationError{
						field:  "LastUsed",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntitlementUsageValidationError{
						field:  "LastUsed",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLastUsed()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntitlementUsageValidationError{
					field:  "LastUsed",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EntitlementUsageMultiError(errors)
	}

	return nil
}

// EntitlementUsageMultiError is an error wrapping multiple validation errors
// returned by EntitlementUsage.ValidateAll() if the designated constraints
// aren't met.
type EntitlementUsageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntitlementUsageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntitlementUsageMultiError) AllErrors() []error { return m }

// EntitlementUsageValidationError is the validation error returned by
// EntitlementUsage.Validate if the designated constraints aren't met.
type EntitlementUsageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntitlementUsageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntitlementUsageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntitlementUsageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntitlementUsageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntitlementUsageValidationError) ErrorName() string { return "EntitlementUsageValidationError" }

// Error satisfies the builtin error interface
func (e EntitlementUsageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntitlementUsage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntitlementUsageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntitlementUsageValidationError{}

// Validate checks the field values on AccessPath with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AccessPath) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AccessPath with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AccessPathMultiError, or
// nil if none found.
func (m *AccessPath) ValidateAll() error {
	return m.validate(true)
}

func (m *AccessPath) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPath() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AccessPathValidationError{
						field:  fmt.Sprintf("Path[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AccessPathValidationError{
						field:  fmt.Sprintf("Path[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccessPathValidationError{
					field:  fmt.Sprintf("Path[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AccessPathMultiError(errors)
	}

	return nil
}

// AccessPathMultiError is an error wrapping multiple validation errors
// returned by AccessPath.ValidateAll() if the designated constraints aren't met.
type AccessPathMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AccessPathMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AccessPathMultiError) AllErrors() []error { return m }

// AccessPathValidationError is the validation error returned by
// AccessPath.Validate if the designated constraints aren't met.
type AccessPathValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccessPathValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccessPathValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccessPathValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccessPathValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccessPathValidationError) ErrorName() string { return "AccessPathValidationError" }

// Error satisfies the builtin error interface
func (e AccessPathValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessPath.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccessPathValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccessPathValidationError{}
