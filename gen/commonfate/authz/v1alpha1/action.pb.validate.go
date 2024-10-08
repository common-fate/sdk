// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/authz/v1alpha1/action.proto

package authzv1alpha1

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

// Validate checks the field values on Action with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Action) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Action with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ActionMultiError, or nil if none found.
func (m *Action) ValidateAll() error {
	return m.validate(true)
}

func (m *Action) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Namespace

	// no validation rules for Id

	if len(errors) > 0 {
		return ActionMultiError(errors)
	}

	return nil
}

// ActionMultiError is an error wrapping multiple validation errors returned by
// Action.ValidateAll() if the designated constraints aren't met.
type ActionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActionMultiError) AllErrors() []error { return m }

// ActionValidationError is the validation error returned by Action.Validate if
// the designated constraints aren't met.
type ActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActionValidationError) ErrorName() string { return "ActionValidationError" }

// Error satisfies the builtin error interface
func (e ActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActionValidationError{}

// Validate checks the field values on ActionOptions with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ActionOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ActionOptions with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ActionOptionsMultiError, or
// nil if none found.
func (m *ActionOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *ActionOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ReadOnly

	for idx, item := range m.GetActionGroups() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ActionOptionsValidationError{
						field:  fmt.Sprintf("ActionGroups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ActionOptionsValidationError{
						field:  fmt.Sprintf("ActionGroups[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ActionOptionsValidationError{
					field:  fmt.Sprintf("ActionGroups[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.ServiceNamespace != nil {
		// no validation rules for ServiceNamespace
	}

	if len(errors) > 0 {
		return ActionOptionsMultiError(errors)
	}

	return nil
}

// ActionOptionsMultiError is an error wrapping multiple validation errors
// returned by ActionOptions.ValidateAll() if the designated constraints
// aren't met.
type ActionOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ActionOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ActionOptionsMultiError) AllErrors() []error { return m }

// ActionOptionsValidationError is the validation error returned by
// ActionOptions.Validate if the designated constraints aren't met.
type ActionOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ActionOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ActionOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ActionOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ActionOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ActionOptionsValidationError) ErrorName() string { return "ActionOptionsValidationError" }

// Error satisfies the builtin error interface
func (e ActionOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sActionOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ActionOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ActionOptionsValidationError{}
