// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/access/v1alpha1/attachments.proto

package accessv1alpha1

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

// Validate checks the field values on JiraIssue with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *JiraIssue) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JiraIssue with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in JiraIssueMultiError, or nil
// if none found.
func (m *JiraIssue) ValidateAll() error {
	return m.validate(true)
}

func (m *JiraIssue) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Summary

	// no validation rules for Url

	if len(errors) > 0 {
		return JiraIssueMultiError(errors)
	}

	return nil
}

// JiraIssueMultiError is an error wrapping multiple validation errors returned
// by JiraIssue.ValidateAll() if the designated constraints aren't met.
type JiraIssueMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JiraIssueMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JiraIssueMultiError) AllErrors() []error { return m }

// JiraIssueValidationError is the validation error returned by
// JiraIssue.Validate if the designated constraints aren't met.
type JiraIssueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JiraIssueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JiraIssueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JiraIssueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JiraIssueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JiraIssueValidationError) ErrorName() string { return "JiraIssueValidationError" }

// Error satisfies the builtin error interface
func (e JiraIssueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJiraIssue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JiraIssueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JiraIssueValidationError{}

// Validate checks the field values on QueryJiraIssuesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QueryJiraIssuesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryJiraIssuesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryJiraIssuesRequestMultiError, or nil if none found.
func (m *QueryJiraIssuesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryJiraIssuesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.SummaryFilter != nil {
		// no validation rules for SummaryFilter
	}

	if len(errors) > 0 {
		return QueryJiraIssuesRequestMultiError(errors)
	}

	return nil
}

// QueryJiraIssuesRequestMultiError is an error wrapping multiple validation
// errors returned by QueryJiraIssuesRequest.ValidateAll() if the designated
// constraints aren't met.
type QueryJiraIssuesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryJiraIssuesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryJiraIssuesRequestMultiError) AllErrors() []error { return m }

// QueryJiraIssuesRequestValidationError is the validation error returned by
// QueryJiraIssuesRequest.Validate if the designated constraints aren't met.
type QueryJiraIssuesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryJiraIssuesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryJiraIssuesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryJiraIssuesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryJiraIssuesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryJiraIssuesRequestValidationError) ErrorName() string {
	return "QueryJiraIssuesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QueryJiraIssuesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryJiraIssuesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryJiraIssuesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryJiraIssuesRequestValidationError{}

// Validate checks the field values on QueryJiraIssuesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QueryJiraIssuesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryJiraIssuesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryJiraIssuesResponseMultiError, or nil if none found.
func (m *QueryJiraIssuesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryJiraIssuesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetIssues() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QueryJiraIssuesResponseValidationError{
						field:  fmt.Sprintf("Issues[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QueryJiraIssuesResponseValidationError{
						field:  fmt.Sprintf("Issues[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryJiraIssuesResponseValidationError{
					field:  fmt.Sprintf("Issues[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QueryJiraIssuesResponseMultiError(errors)
	}

	return nil
}

// QueryJiraIssuesResponseMultiError is an error wrapping multiple validation
// errors returned by QueryJiraIssuesResponse.ValidateAll() if the designated
// constraints aren't met.
type QueryJiraIssuesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryJiraIssuesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryJiraIssuesResponseMultiError) AllErrors() []error { return m }

// QueryJiraIssuesResponseValidationError is the validation error returned by
// QueryJiraIssuesResponse.Validate if the designated constraints aren't met.
type QueryJiraIssuesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryJiraIssuesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryJiraIssuesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryJiraIssuesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryJiraIssuesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryJiraIssuesResponseValidationError) ErrorName() string {
	return "QueryJiraIssuesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e QueryJiraIssuesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryJiraIssuesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryJiraIssuesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryJiraIssuesResponseValidationError{}

// Validate checks the field values on AttachRequestCommentToJiraIssuesRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *AttachRequestCommentToJiraIssuesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// AttachRequestCommentToJiraIssuesRequest with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// AttachRequestCommentToJiraIssuesRequestMultiError, or nil if none found.
func (m *AttachRequestCommentToJiraIssuesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AttachRequestCommentToJiraIssuesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAccessRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttachRequestCommentToJiraIssuesRequestValidationError{
					field:  "AccessRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttachRequestCommentToJiraIssuesRequestValidationError{
					field:  "AccessRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccessRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttachRequestCommentToJiraIssuesRequestValidationError{
				field:  "AccessRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AttachRequestCommentToJiraIssuesRequestMultiError(errors)
	}

	return nil
}

// AttachRequestCommentToJiraIssuesRequestMultiError is an error wrapping
// multiple validation errors returned by
// AttachRequestCommentToJiraIssuesRequest.ValidateAll() if the designated
// constraints aren't met.
type AttachRequestCommentToJiraIssuesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttachRequestCommentToJiraIssuesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttachRequestCommentToJiraIssuesRequestMultiError) AllErrors() []error { return m }

// AttachRequestCommentToJiraIssuesRequestValidationError is the validation
// error returned by AttachRequestCommentToJiraIssuesRequest.Validate if the
// designated constraints aren't met.
type AttachRequestCommentToJiraIssuesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttachRequestCommentToJiraIssuesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttachRequestCommentToJiraIssuesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttachRequestCommentToJiraIssuesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttachRequestCommentToJiraIssuesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttachRequestCommentToJiraIssuesRequestValidationError) ErrorName() string {
	return "AttachRequestCommentToJiraIssuesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AttachRequestCommentToJiraIssuesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttachRequestCommentToJiraIssuesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttachRequestCommentToJiraIssuesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttachRequestCommentToJiraIssuesRequestValidationError{}

// Validate checks the field values on AttachRequestCommentToJiraIssuesResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *AttachRequestCommentToJiraIssuesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// AttachRequestCommentToJiraIssuesResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// AttachRequestCommentToJiraIssuesResponseMultiError, or nil if none found.
func (m *AttachRequestCommentToJiraIssuesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AttachRequestCommentToJiraIssuesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AttachRequestCommentToJiraIssuesResponseMultiError(errors)
	}

	return nil
}

// AttachRequestCommentToJiraIssuesResponseMultiError is an error wrapping
// multiple validation errors returned by
// AttachRequestCommentToJiraIssuesResponse.ValidateAll() if the designated
// constraints aren't met.
type AttachRequestCommentToJiraIssuesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttachRequestCommentToJiraIssuesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttachRequestCommentToJiraIssuesResponseMultiError) AllErrors() []error { return m }

// AttachRequestCommentToJiraIssuesResponseValidationError is the validation
// error returned by AttachRequestCommentToJiraIssuesResponse.Validate if the
// designated constraints aren't met.
type AttachRequestCommentToJiraIssuesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttachRequestCommentToJiraIssuesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttachRequestCommentToJiraIssuesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttachRequestCommentToJiraIssuesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttachRequestCommentToJiraIssuesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttachRequestCommentToJiraIssuesResponseValidationError) ErrorName() string {
	return "AttachRequestCommentToJiraIssuesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AttachRequestCommentToJiraIssuesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttachRequestCommentToJiraIssuesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttachRequestCommentToJiraIssuesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttachRequestCommentToJiraIssuesResponseValidationError{}
