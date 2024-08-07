// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/leastprivilege/v1alpha1/report.proto

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

// Validate checks the field values on Report with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Report) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Report with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ReportMultiError, or nil if none found.
func (m *Report) ValidateAll() error {
	return m.validate(true)
}

func (m *Report) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsageSummaries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReportValidationError{
						field:  fmt.Sprintf("UsageSummaries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReportValidationError{
						field:  fmt.Sprintf("UsageSummaries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReportValidationError{
					field:  fmt.Sprintf("UsageSummaries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReportValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReportValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReportValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReportValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReportValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReportValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUsageCutoffTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReportValidationError{
					field:  "UsageCutoffTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReportValidationError{
					field:  "UsageCutoffTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUsageCutoffTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReportValidationError{
				field:  "UsageCutoffTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ReportMultiError(errors)
	}

	return nil
}

// ReportMultiError is an error wrapping multiple validation errors returned by
// Report.ValidateAll() if the designated constraints aren't met.
type ReportMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReportMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReportMultiError) AllErrors() []error { return m }

// ReportValidationError is the validation error returned by Report.Validate if
// the designated constraints aren't met.
type ReportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReportValidationError) ErrorName() string { return "ReportValidationError" }

// Error satisfies the builtin error interface
func (e ReportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReportValidationError{}

// Validate checks the field values on UsageSummary with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UsageSummary) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UsageSummary with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UsageSummaryMultiError, or
// nil if none found.
func (m *UsageSummary) ValidateAll() error {
	return m.validate(true)
}

func (m *UsageSummary) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Label

	// no validation rules for InUseCount

	// no validation rules for UnusedCount

	// no validation rules for IndeterminateCount

	if len(errors) > 0 {
		return UsageSummaryMultiError(errors)
	}

	return nil
}

// UsageSummaryMultiError is an error wrapping multiple validation errors
// returned by UsageSummary.ValidateAll() if the designated constraints aren't met.
type UsageSummaryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UsageSummaryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UsageSummaryMultiError) AllErrors() []error { return m }

// UsageSummaryValidationError is the validation error returned by
// UsageSummary.Validate if the designated constraints aren't met.
type UsageSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UsageSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UsageSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UsageSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UsageSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UsageSummaryValidationError) ErrorName() string { return "UsageSummaryValidationError" }

// Error satisfies the builtin error interface
func (e UsageSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUsageSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UsageSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UsageSummaryValidationError{}
