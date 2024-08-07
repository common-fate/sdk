// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/leastprivilege/v1alpha1/leastprivilege.proto

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

// Validate checks the field values on GetLatestReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLatestReportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLatestReportRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLatestReportRequestMultiError, or nil if none found.
func (m *GetLatestReportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLatestReportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetLatestReportRequestMultiError(errors)
	}

	return nil
}

// GetLatestReportRequestMultiError is an error wrapping multiple validation
// errors returned by GetLatestReportRequest.ValidateAll() if the designated
// constraints aren't met.
type GetLatestReportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLatestReportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLatestReportRequestMultiError) AllErrors() []error { return m }

// GetLatestReportRequestValidationError is the validation error returned by
// GetLatestReportRequest.Validate if the designated constraints aren't met.
type GetLatestReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLatestReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLatestReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLatestReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLatestReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLatestReportRequestValidationError) ErrorName() string {
	return "GetLatestReportRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetLatestReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLatestReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLatestReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLatestReportRequestValidationError{}

// Validate checks the field values on GetLatestReportResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetLatestReportResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetLatestReportResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetLatestReportResponseMultiError, or nil if none found.
func (m *GetLatestReportResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetLatestReportResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetReport()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetLatestReportResponseValidationError{
					field:  "Report",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetLatestReportResponseValidationError{
					field:  "Report",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReport()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetLatestReportResponseValidationError{
				field:  "Report",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetLatestReportResponseMultiError(errors)
	}

	return nil
}

// GetLatestReportResponseMultiError is an error wrapping multiple validation
// errors returned by GetLatestReportResponse.ValidateAll() if the designated
// constraints aren't met.
type GetLatestReportResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetLatestReportResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetLatestReportResponseMultiError) AllErrors() []error { return m }

// GetLatestReportResponseValidationError is the validation error returned by
// GetLatestReportResponse.Validate if the designated constraints aren't met.
type GetLatestReportResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetLatestReportResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetLatestReportResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetLatestReportResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetLatestReportResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetLatestReportResponseValidationError) ErrorName() string {
	return "GetLatestReportResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetLatestReportResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetLatestReportResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetLatestReportResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetLatestReportResponseValidationError{}

// Validate checks the field values on GetHistoricalReportsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetHistoricalReportsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHistoricalReportsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetHistoricalReportsRequestMultiError, or nil if none found.
func (m *GetHistoricalReportsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHistoricalReportsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetHistoricalReportsRequestMultiError(errors)
	}

	return nil
}

// GetHistoricalReportsRequestMultiError is an error wrapping multiple
// validation errors returned by GetHistoricalReportsRequest.ValidateAll() if
// the designated constraints aren't met.
type GetHistoricalReportsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHistoricalReportsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHistoricalReportsRequestMultiError) AllErrors() []error { return m }

// GetHistoricalReportsRequestValidationError is the validation error returned
// by GetHistoricalReportsRequest.Validate if the designated constraints
// aren't met.
type GetHistoricalReportsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHistoricalReportsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHistoricalReportsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHistoricalReportsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHistoricalReportsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHistoricalReportsRequestValidationError) ErrorName() string {
	return "GetHistoricalReportsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetHistoricalReportsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHistoricalReportsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHistoricalReportsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHistoricalReportsRequestValidationError{}

// Validate checks the field values on GetHistoricalReportsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetHistoricalReportsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHistoricalReportsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetHistoricalReportsResponseMultiError, or nil if none found.
func (m *GetHistoricalReportsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHistoricalReportsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetReports() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetHistoricalReportsResponseValidationError{
						field:  fmt.Sprintf("Reports[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetHistoricalReportsResponseValidationError{
						field:  fmt.Sprintf("Reports[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetHistoricalReportsResponseValidationError{
					field:  fmt.Sprintf("Reports[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetHistoricalReportsResponseMultiError(errors)
	}

	return nil
}

// GetHistoricalReportsResponseMultiError is an error wrapping multiple
// validation errors returned by GetHistoricalReportsResponse.ValidateAll() if
// the designated constraints aren't met.
type GetHistoricalReportsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHistoricalReportsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHistoricalReportsResponseMultiError) AllErrors() []error { return m }

// GetHistoricalReportsResponseValidationError is the validation error returned
// by GetHistoricalReportsResponse.Validate if the designated constraints
// aren't met.
type GetHistoricalReportsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHistoricalReportsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHistoricalReportsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHistoricalReportsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHistoricalReportsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHistoricalReportsResponseValidationError) ErrorName() string {
	return "GetHistoricalReportsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetHistoricalReportsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHistoricalReportsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHistoricalReportsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHistoricalReportsResponseValidationError{}

// Validate checks the field values on DownloadEntitlementUsageReportRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *DownloadEntitlementUsageReportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadEntitlementUsageReportRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// DownloadEntitlementUsageReportRequestMultiError, or nil if none found.
func (m *DownloadEntitlementUsageReportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadEntitlementUsageReportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DownloadEntitlementUsageReportRequestMultiError(errors)
	}

	return nil
}

// DownloadEntitlementUsageReportRequestMultiError is an error wrapping
// multiple validation errors returned by
// DownloadEntitlementUsageReportRequest.ValidateAll() if the designated
// constraints aren't met.
type DownloadEntitlementUsageReportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadEntitlementUsageReportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadEntitlementUsageReportRequestMultiError) AllErrors() []error { return m }

// DownloadEntitlementUsageReportRequestValidationError is the validation error
// returned by DownloadEntitlementUsageReportRequest.Validate if the
// designated constraints aren't met.
type DownloadEntitlementUsageReportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadEntitlementUsageReportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadEntitlementUsageReportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadEntitlementUsageReportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadEntitlementUsageReportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadEntitlementUsageReportRequestValidationError) ErrorName() string {
	return "DownloadEntitlementUsageReportRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadEntitlementUsageReportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadEntitlementUsageReportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadEntitlementUsageReportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadEntitlementUsageReportRequestValidationError{}

// Validate checks the field values on DownloadEntitlementUsageReportResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *DownloadEntitlementUsageReportResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// DownloadEntitlementUsageReportResponse with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// DownloadEntitlementUsageReportResponseMultiError, or nil if none found.
func (m *DownloadEntitlementUsageReportResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadEntitlementUsageReportResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DownloadUrl

	if len(errors) > 0 {
		return DownloadEntitlementUsageReportResponseMultiError(errors)
	}

	return nil
}

// DownloadEntitlementUsageReportResponseMultiError is an error wrapping
// multiple validation errors returned by
// DownloadEntitlementUsageReportResponse.ValidateAll() if the designated
// constraints aren't met.
type DownloadEntitlementUsageReportResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadEntitlementUsageReportResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadEntitlementUsageReportResponseMultiError) AllErrors() []error { return m }

// DownloadEntitlementUsageReportResponseValidationError is the validation
// error returned by DownloadEntitlementUsageReportResponse.Validate if the
// designated constraints aren't met.
type DownloadEntitlementUsageReportResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadEntitlementUsageReportResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadEntitlementUsageReportResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadEntitlementUsageReportResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadEntitlementUsageReportResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadEntitlementUsageReportResponseValidationError) ErrorName() string {
	return "DownloadEntitlementUsageReportResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadEntitlementUsageReportResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadEntitlementUsageReportResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadEntitlementUsageReportResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadEntitlementUsageReportResponseValidationError{}
