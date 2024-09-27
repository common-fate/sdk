// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/filters/v1alpha1/filters.proto

package filtersv1alpha1

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

	authzv1alpha1 "github.com/common-fate/sdk/gen/commonfate/authz/v1alpha1"
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

	_ = authzv1alpha1.Decision(0)
)

// Validate checks the field values on DecisionFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DecisionFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DecisionFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DecisionFilterMultiError,
// or nil if none found.
func (m *DecisionFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *DecisionFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Decision

	if len(errors) > 0 {
		return DecisionFilterMultiError(errors)
	}

	return nil
}

// DecisionFilterMultiError is an error wrapping multiple validation errors
// returned by DecisionFilter.ValidateAll() if the designated constraints
// aren't met.
type DecisionFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DecisionFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DecisionFilterMultiError) AllErrors() []error { return m }

// DecisionFilterValidationError is the validation error returned by
// DecisionFilter.Validate if the designated constraints aren't met.
type DecisionFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DecisionFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DecisionFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DecisionFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DecisionFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DecisionFilterValidationError) ErrorName() string { return "DecisionFilterValidationError" }

// Error satisfies the builtin error interface
func (e DecisionFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDecisionFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DecisionFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DecisionFilterValidationError{}

// Validate checks the field values on TagFilter with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TagFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TagFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TagFilterMultiError, or nil
// if none found.
func (m *TagFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *TagFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Value

	// no validation rules for Comparison

	if len(errors) > 0 {
		return TagFilterMultiError(errors)
	}

	return nil
}

// TagFilterMultiError is an error wrapping multiple validation errors returned
// by TagFilter.ValidateAll() if the designated constraints aren't met.
type TagFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TagFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TagFilterMultiError) AllErrors() []error { return m }

// TagFilterValidationError is the validation error returned by
// TagFilter.Validate if the designated constraints aren't met.
type TagFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagFilterValidationError) ErrorName() string { return "TagFilterValidationError" }

// Error satisfies the builtin error interface
func (e TagFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTagFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagFilterValidationError{}

// Validate checks the field values on EntityFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EntityFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EntityFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EntityFilterMultiError, or
// nil if none found.
func (m *EntityFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *EntityFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetIds() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EntityFilterValidationError{
						field:  fmt.Sprintf("Ids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EntityFilterValidationError{
						field:  fmt.Sprintf("Ids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EntityFilterValidationError{
					field:  fmt.Sprintf("Ids[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Comparison

	if len(errors) > 0 {
		return EntityFilterMultiError(errors)
	}

	return nil
}

// EntityFilterMultiError is an error wrapping multiple validation errors
// returned by EntityFilter.ValidateAll() if the designated constraints aren't met.
type EntityFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntityFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntityFilterMultiError) AllErrors() []error { return m }

// EntityFilterValidationError is the validation error returned by
// EntityFilter.Validate if the designated constraints aren't met.
type EntityFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityFilterValidationError) ErrorName() string { return "EntityFilterValidationError" }

// Error satisfies the builtin error interface
func (e EntityFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntityFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityFilterValidationError{}

// Validate checks the field values on EntityTypeFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *EntityTypeFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EntityTypeFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EntityTypeFilterMultiError, or nil if none found.
func (m *EntityTypeFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *EntityTypeFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Comparison

	if len(errors) > 0 {
		return EntityTypeFilterMultiError(errors)
	}

	return nil
}

// EntityTypeFilterMultiError is an error wrapping multiple validation errors
// returned by EntityTypeFilter.ValidateAll() if the designated constraints
// aren't met.
type EntityTypeFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EntityTypeFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EntityTypeFilterMultiError) AllErrors() []error { return m }

// EntityTypeFilterValidationError is the validation error returned by
// EntityTypeFilter.Validate if the designated constraints aren't met.
type EntityTypeFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityTypeFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityTypeFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityTypeFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityTypeFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityTypeFilterValidationError) ErrorName() string { return "EntityTypeFilterValidationError" }

// Error satisfies the builtin error interface
func (e EntityTypeFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntityTypeFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityTypeFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityTypeFilterValidationError{}

// Validate checks the field values on OccurredAtFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OccurredAtFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OccurredAtFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OccurredAtFilterMultiError, or nil if none found.
func (m *OccurredAtFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *OccurredAtFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OccurredAtFilterValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OccurredAtFilterValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OccurredAtFilterValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Comparison

	if len(errors) > 0 {
		return OccurredAtFilterMultiError(errors)
	}

	return nil
}

// OccurredAtFilterMultiError is an error wrapping multiple validation errors
// returned by OccurredAtFilter.ValidateAll() if the designated constraints
// aren't met.
type OccurredAtFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OccurredAtFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OccurredAtFilterMultiError) AllErrors() []error { return m }

// OccurredAtFilterValidationError is the validation error returned by
// OccurredAtFilter.Validate if the designated constraints aren't met.
type OccurredAtFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OccurredAtFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OccurredAtFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OccurredAtFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OccurredAtFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OccurredAtFilterValidationError) ErrorName() string { return "OccurredAtFilterValidationError" }

// Error satisfies the builtin error interface
func (e OccurredAtFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOccurredAtFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OccurredAtFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OccurredAtFilterValidationError{}

// Validate checks the field values on IntegerFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IntegerFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IntegerFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IntegerFilterMultiError, or
// nil if none found.
func (m *IntegerFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *IntegerFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	// no validation rules for Comparison

	if len(errors) > 0 {
		return IntegerFilterMultiError(errors)
	}

	return nil
}

// IntegerFilterMultiError is an error wrapping multiple validation errors
// returned by IntegerFilter.ValidateAll() if the designated constraints
// aren't met.
type IntegerFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IntegerFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IntegerFilterMultiError) AllErrors() []error { return m }

// IntegerFilterValidationError is the validation error returned by
// IntegerFilter.Validate if the designated constraints aren't met.
type IntegerFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegerFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegerFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegerFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegerFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegerFilterValidationError) ErrorName() string { return "IntegerFilterValidationError" }

// Error satisfies the builtin error interface
func (e IntegerFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegerFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegerFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegerFilterValidationError{}

// Validate checks the field values on BoolFilter with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BoolFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BoolFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BoolFilterMultiError, or
// nil if none found.
func (m *BoolFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *BoolFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Comparison

	if len(errors) > 0 {
		return BoolFilterMultiError(errors)
	}

	return nil
}

// BoolFilterMultiError is an error wrapping multiple validation errors
// returned by BoolFilter.ValidateAll() if the designated constraints aren't met.
type BoolFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BoolFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BoolFilterMultiError) AllErrors() []error { return m }

// BoolFilterValidationError is the validation error returned by
// BoolFilter.Validate if the designated constraints aren't met.
type BoolFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BoolFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BoolFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BoolFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BoolFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BoolFilterValidationError) ErrorName() string { return "BoolFilterValidationError" }

// Error satisfies the builtin error interface
func (e BoolFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBoolFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BoolFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BoolFilterValidationError{}

// Validate checks the field values on RelativeRange with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RelativeRange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RelativeRange with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RelativeRangeMultiError, or
// nil if none found.
func (m *RelativeRange) ValidateAll() error {
	return m.validate(true)
}

func (m *RelativeRange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RelativeRangeValidationError{
					field:  "Duration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RelativeRangeValidationError{
					field:  "Duration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RelativeRangeValidationError{
				field:  "Duration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RelativeRangeMultiError(errors)
	}

	return nil
}

// RelativeRangeMultiError is an error wrapping multiple validation errors
// returned by RelativeRange.ValidateAll() if the designated constraints
// aren't met.
type RelativeRangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RelativeRangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RelativeRangeMultiError) AllErrors() []error { return m }

// RelativeRangeValidationError is the validation error returned by
// RelativeRange.Validate if the designated constraints aren't met.
type RelativeRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RelativeRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RelativeRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RelativeRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RelativeRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RelativeRangeValidationError) ErrorName() string { return "RelativeRangeValidationError" }

// Error satisfies the builtin error interface
func (e RelativeRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRelativeRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RelativeRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RelativeRangeValidationError{}

// Validate checks the field values on AbsoluteRange with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AbsoluteRange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AbsoluteRange with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AbsoluteRangeMultiError, or
// nil if none found.
func (m *AbsoluteRange) ValidateAll() error {
	return m.validate(true)
}

func (m *AbsoluteRange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AbsoluteRangeValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AbsoluteRangeValidationError{
					field:  "Time",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AbsoluteRangeValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AbsoluteRangeMultiError(errors)
	}

	return nil
}

// AbsoluteRangeMultiError is an error wrapping multiple validation errors
// returned by AbsoluteRange.ValidateAll() if the designated constraints
// aren't met.
type AbsoluteRangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AbsoluteRangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AbsoluteRangeMultiError) AllErrors() []error { return m }

// AbsoluteRangeValidationError is the validation error returned by
// AbsoluteRange.Validate if the designated constraints aren't met.
type AbsoluteRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AbsoluteRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AbsoluteRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AbsoluteRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AbsoluteRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AbsoluteRangeValidationError) ErrorName() string { return "AbsoluteRangeValidationError" }

// Error satisfies the builtin error interface
func (e AbsoluteRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAbsoluteRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AbsoluteRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AbsoluteRangeValidationError{}

// Validate checks the field values on TimeRange with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TimeRange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TimeRange with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TimeRangeMultiError, or nil
// if none found.
func (m *TimeRange) ValidateAll() error {
	return m.validate(true)
}

func (m *TimeRange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Range.(type) {
	case *TimeRange_Relative:
		if v == nil {
			err := TimeRangeValidationError{
				field:  "Range",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetRelative()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TimeRangeValidationError{
						field:  "Relative",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TimeRangeValidationError{
						field:  "Relative",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRelative()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TimeRangeValidationError{
					field:  "Relative",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TimeRange_Absolute:
		if v == nil {
			err := TimeRangeValidationError{
				field:  "Range",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetAbsolute()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TimeRangeValidationError{
						field:  "Absolute",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TimeRangeValidationError{
						field:  "Absolute",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAbsolute()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TimeRangeValidationError{
					field:  "Absolute",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return TimeRangeMultiError(errors)
	}

	return nil
}

// TimeRangeMultiError is an error wrapping multiple validation errors returned
// by TimeRange.ValidateAll() if the designated constraints aren't met.
type TimeRangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TimeRangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TimeRangeMultiError) AllErrors() []error { return m }

// TimeRangeValidationError is the validation error returned by
// TimeRange.Validate if the designated constraints aren't met.
type TimeRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimeRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimeRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimeRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimeRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimeRangeValidationError) ErrorName() string { return "TimeRangeValidationError" }

// Error satisfies the builtin error interface
func (e TimeRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimeRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimeRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimeRangeValidationError{}

// Validate checks the field values on TimeRangeFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TimeRangeFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TimeRangeFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TimeRangeFilterMultiError, or nil if none found.
func (m *TimeRangeFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *TimeRangeFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Start != nil {

		if all {
			switch v := interface{}(m.GetStart()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TimeRangeFilterValidationError{
						field:  "Start",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TimeRangeFilterValidationError{
						field:  "Start",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStart()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TimeRangeFilterValidationError{
					field:  "Start",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.End != nil {

		if all {
			switch v := interface{}(m.GetEnd()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TimeRangeFilterValidationError{
						field:  "End",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TimeRangeFilterValidationError{
						field:  "End",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEnd()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TimeRangeFilterValidationError{
					field:  "End",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TimeRangeFilterMultiError(errors)
	}

	return nil
}

// TimeRangeFilterMultiError is an error wrapping multiple validation errors
// returned by TimeRangeFilter.ValidateAll() if the designated constraints
// aren't met.
type TimeRangeFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TimeRangeFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TimeRangeFilterMultiError) AllErrors() []error { return m }

// TimeRangeFilterValidationError is the validation error returned by
// TimeRangeFilter.Validate if the designated constraints aren't met.
type TimeRangeFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TimeRangeFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TimeRangeFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TimeRangeFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TimeRangeFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TimeRangeFilterValidationError) ErrorName() string { return "TimeRangeFilterValidationError" }

// Error satisfies the builtin error interface
func (e TimeRangeFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTimeRangeFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TimeRangeFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TimeRangeFilterValidationError{}
