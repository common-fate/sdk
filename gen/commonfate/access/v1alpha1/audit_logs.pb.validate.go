// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/access/v1alpha1/audit_logs.proto

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

// Validate checks the field values on AuditLogFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuditLogFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuditLogFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuditLogFilterMultiError,
// or nil if none found.
func (m *AuditLogFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *AuditLogFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Filter.(type) {
	case *AuditLogFilter_OccurredAt:
		if v == nil {
			err := AuditLogFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetOccurredAt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AuditLogFilterValidationError{
						field:  "OccurredAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AuditLogFilterValidationError{
						field:  "OccurredAt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetOccurredAt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuditLogFilterValidationError{
					field:  "OccurredAt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AuditLogFilter_Actor:
		if v == nil {
			err := AuditLogFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetActor()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AuditLogFilterValidationError{
						field:  "Actor",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AuditLogFilterValidationError{
						field:  "Actor",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetActor()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuditLogFilterValidationError{
					field:  "Actor",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AuditLogFilter_ActionType:
		if v == nil {
			err := AuditLogFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetActionType()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AuditLogFilterValidationError{
						field:  "ActionType",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AuditLogFilterValidationError{
						field:  "ActionType",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetActionType()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuditLogFilterValidationError{
					field:  "ActionType",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AuditLogFilter_Target:
		if v == nil {
			err := AuditLogFilterValidationError{
				field:  "Filter",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetTarget()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AuditLogFilterValidationError{
						field:  "Target",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AuditLogFilterValidationError{
						field:  "Target",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTarget()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuditLogFilterValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return AuditLogFilterMultiError(errors)
	}

	return nil
}

// AuditLogFilterMultiError is an error wrapping multiple validation errors
// returned by AuditLogFilter.ValidateAll() if the designated constraints
// aren't met.
type AuditLogFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuditLogFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuditLogFilterMultiError) AllErrors() []error { return m }

// AuditLogFilterValidationError is the validation error returned by
// AuditLogFilter.Validate if the designated constraints aren't met.
type AuditLogFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuditLogFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuditLogFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuditLogFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuditLogFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuditLogFilterValidationError) ErrorName() string { return "AuditLogFilterValidationError" }

// Error satisfies the builtin error interface
func (e AuditLogFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuditLogFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuditLogFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuditLogFilterValidationError{}

// Validate checks the field values on QueryAuditLogsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QueryAuditLogsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryAuditLogsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryAuditLogsRequestMultiError, or nil if none found.
func (m *QueryAuditLogsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryAuditLogsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTarget()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QueryAuditLogsRequestValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QueryAuditLogsRequestValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTarget()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QueryAuditLogsRequestValidationError{
				field:  "Target",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QueryAuditLogsRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QueryAuditLogsRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryAuditLogsRequestValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for PageToken

	if m.Order != nil {
		// no validation rules for Order
	}

	if len(errors) > 0 {
		return QueryAuditLogsRequestMultiError(errors)
	}

	return nil
}

// QueryAuditLogsRequestMultiError is an error wrapping multiple validation
// errors returned by QueryAuditLogsRequest.ValidateAll() if the designated
// constraints aren't met.
type QueryAuditLogsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryAuditLogsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryAuditLogsRequestMultiError) AllErrors() []error { return m }

// QueryAuditLogsRequestValidationError is the validation error returned by
// QueryAuditLogsRequest.Validate if the designated constraints aren't met.
type QueryAuditLogsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryAuditLogsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryAuditLogsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryAuditLogsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryAuditLogsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryAuditLogsRequestValidationError) ErrorName() string {
	return "QueryAuditLogsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QueryAuditLogsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryAuditLogsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryAuditLogsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryAuditLogsRequestValidationError{}

// Validate checks the field values on QueryAuditLogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QueryAuditLogsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryAuditLogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryAuditLogsResponseMultiError, or nil if none found.
func (m *QueryAuditLogsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryAuditLogsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAuditLogs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QueryAuditLogsResponseValidationError{
						field:  fmt.Sprintf("AuditLogs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QueryAuditLogsResponseValidationError{
						field:  fmt.Sprintf("AuditLogs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QueryAuditLogsResponseValidationError{
					field:  fmt.Sprintf("AuditLogs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for NextPageToken

	if len(errors) > 0 {
		return QueryAuditLogsResponseMultiError(errors)
	}

	return nil
}

// QueryAuditLogsResponseMultiError is an error wrapping multiple validation
// errors returned by QueryAuditLogsResponse.ValidateAll() if the designated
// constraints aren't met.
type QueryAuditLogsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryAuditLogsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryAuditLogsResponseMultiError) AllErrors() []error { return m }

// QueryAuditLogsResponseValidationError is the validation error returned by
// QueryAuditLogsResponse.Validate if the designated constraints aren't met.
type QueryAuditLogsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryAuditLogsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryAuditLogsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryAuditLogsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryAuditLogsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryAuditLogsResponseValidationError) ErrorName() string {
	return "QueryAuditLogsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e QueryAuditLogsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryAuditLogsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryAuditLogsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryAuditLogsResponseValidationError{}

// Validate checks the field values on AuditLogPreview with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuditLogPreview) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuditLogPreview with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuditLogPreviewMultiError, or nil if none found.
func (m *AuditLogPreview) ValidateAll() error {
	return m.validate(true)
}

func (m *AuditLogPreview) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLogs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AuditLogPreviewValidationError{
						field:  fmt.Sprintf("Logs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AuditLogPreviewValidationError{
						field:  fmt.Sprintf("Logs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuditLogPreviewValidationError{
					field:  fmt.Sprintf("Logs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for LogCount

	if len(errors) > 0 {
		return AuditLogPreviewMultiError(errors)
	}

	return nil
}

// AuditLogPreviewMultiError is an error wrapping multiple validation errors
// returned by AuditLogPreview.ValidateAll() if the designated constraints
// aren't met.
type AuditLogPreviewMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuditLogPreviewMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuditLogPreviewMultiError) AllErrors() []error { return m }

// AuditLogPreviewValidationError is the validation error returned by
// AuditLogPreview.Validate if the designated constraints aren't met.
type AuditLogPreviewValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuditLogPreviewValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuditLogPreviewValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuditLogPreviewValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuditLogPreviewValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuditLogPreviewValidationError) ErrorName() string { return "AuditLogPreviewValidationError" }

// Error satisfies the builtin error interface
func (e AuditLogPreviewValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuditLogPreview.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuditLogPreviewValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuditLogPreviewValidationError{}

// Validate checks the field values on AuditLog with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuditLog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuditLog with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuditLogMultiError, or nil
// if none found.
func (m *AuditLog) ValidateAll() error {
	return m.validate(true)
}

func (m *AuditLog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Action

	if all {
		switch v := interface{}(m.GetActor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AuditLogValidationError{
					field:  "Actor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AuditLogValidationError{
					field:  "Actor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetActor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuditLogValidationError{
				field:  "Actor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOccurredAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AuditLogValidationError{
					field:  "OccurredAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AuditLogValidationError{
					field:  "OccurredAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOccurredAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuditLogValidationError{
				field:  "OccurredAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTargets() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AuditLogValidationError{
						field:  fmt.Sprintf("Targets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AuditLogValidationError{
						field:  fmt.Sprintf("Targets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuditLogValidationError{
					field:  fmt.Sprintf("Targets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetContext()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AuditLogValidationError{
					field:  "Context",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AuditLogValidationError{
					field:  "Context",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuditLogValidationError{
				field:  "Context",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetCallerIdentityChain() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AuditLogValidationError{
						field:  fmt.Sprintf("CallerIdentityChain[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AuditLogValidationError{
						field:  fmt.Sprintf("CallerIdentityChain[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuditLogValidationError{
					field:  fmt.Sprintf("CallerIdentityChain[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.RequestId != nil {
		// no validation rules for RequestId
	}

	if m.GrantId != nil {
		// no validation rules for GrantId
	}

	if m.Entitlement != nil {

		if all {
			switch v := interface{}(m.GetEntitlement()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AuditLogValidationError{
						field:  "Entitlement",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AuditLogValidationError{
						field:  "Entitlement",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEntitlement()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuditLogValidationError{
					field:  "Entitlement",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return AuditLogMultiError(errors)
	}

	return nil
}

// AuditLogMultiError is an error wrapping multiple validation errors returned
// by AuditLog.ValidateAll() if the designated constraints aren't met.
type AuditLogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuditLogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuditLogMultiError) AllErrors() []error { return m }

// AuditLogValidationError is the validation error returned by
// AuditLog.Validate if the designated constraints aren't met.
type AuditLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuditLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuditLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuditLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuditLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuditLogValidationError) ErrorName() string { return "AuditLogValidationError" }

// Error satisfies the builtin error interface
func (e AuditLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuditLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuditLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuditLogValidationError{}

// Validate checks the field values on AuditLogEntitlement with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuditLogEntitlement) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuditLogEntitlement with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuditLogEntitlementMultiError, or nil if none found.
func (m *AuditLogEntitlement) ValidateAll() error {
	return m.validate(true)
}

func (m *AuditLogEntitlement) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTarget()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AuditLogEntitlementValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AuditLogEntitlementValidationError{
					field:  "Target",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTarget()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuditLogEntitlementValidationError{
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
				errors = append(errors, AuditLogEntitlementValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AuditLogEntitlementValidationError{
					field:  "Role",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRole()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuditLogEntitlementValidationError{
				field:  "Role",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AuditLogEntitlementMultiError(errors)
	}

	return nil
}

// AuditLogEntitlementMultiError is an error wrapping multiple validation
// errors returned by AuditLogEntitlement.ValidateAll() if the designated
// constraints aren't met.
type AuditLogEntitlementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuditLogEntitlementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuditLogEntitlementMultiError) AllErrors() []error { return m }

// AuditLogEntitlementValidationError is the validation error returned by
// AuditLogEntitlement.Validate if the designated constraints aren't met.
type AuditLogEntitlementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuditLogEntitlementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuditLogEntitlementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuditLogEntitlementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuditLogEntitlementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuditLogEntitlementValidationError) ErrorName() string {
	return "AuditLogEntitlementValidationError"
}

// Error satisfies the builtin error interface
func (e AuditLogEntitlementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuditLogEntitlement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuditLogEntitlementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuditLogEntitlementValidationError{}
