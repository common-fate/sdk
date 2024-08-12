// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: granted/ingest/aws/v1alpha1/ingest.proto

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

// Validate checks the field values on BatchWriteEventsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BatchWriteEventsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BatchWriteEventsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BatchWriteEventsRequestMultiError, or nil if none found.
func (m *BatchWriteEventsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BatchWriteEventsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BatchWriteEventsRequestValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BatchWriteEventsRequestValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BatchWriteEventsRequestValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BatchWriteEventsRequestMultiError(errors)
	}

	return nil
}

// BatchWriteEventsRequestMultiError is an error wrapping multiple validation
// errors returned by BatchWriteEventsRequest.ValidateAll() if the designated
// constraints aren't met.
type BatchWriteEventsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BatchWriteEventsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BatchWriteEventsRequestMultiError) AllErrors() []error { return m }

// BatchWriteEventsRequestValidationError is the validation error returned by
// BatchWriteEventsRequest.Validate if the designated constraints aren't met.
type BatchWriteEventsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchWriteEventsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchWriteEventsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchWriteEventsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchWriteEventsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchWriteEventsRequestValidationError) ErrorName() string {
	return "BatchWriteEventsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BatchWriteEventsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchWriteEventsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchWriteEventsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchWriteEventsRequestValidationError{}

// Validate checks the field values on BatchWriteEventsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BatchWriteEventsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BatchWriteEventsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BatchWriteEventsResponseMultiError, or nil if none found.
func (m *BatchWriteEventsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BatchWriteEventsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BatchWriteEventsResponseMultiError(errors)
	}

	return nil
}

// BatchWriteEventsResponseMultiError is an error wrapping multiple validation
// errors returned by BatchWriteEventsResponse.ValidateAll() if the designated
// constraints aren't met.
type BatchWriteEventsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BatchWriteEventsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BatchWriteEventsResponseMultiError) AllErrors() []error { return m }

// BatchWriteEventsResponseValidationError is the validation error returned by
// BatchWriteEventsResponse.Validate if the designated constraints aren't met.
type BatchWriteEventsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchWriteEventsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchWriteEventsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchWriteEventsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchWriteEventsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchWriteEventsResponseValidationError) ErrorName() string {
	return "BatchWriteEventsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e BatchWriteEventsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchWriteEventsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchWriteEventsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchWriteEventsResponseValidationError{}

// Validate checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Event) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Event with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EventMultiError, or nil if none found.
func (m *Event) ValidateAll() error {
	return m.validate(true)
}

func (m *Event) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EventValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EventValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch v := m.Details.(type) {
	case *Event_AssumeRole:
		if v == nil {
			err := EventValidationError{
				field:  "Details",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetAssumeRole()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "AssumeRole",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "AssumeRole",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAssumeRole()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					field:  "AssumeRole",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Event_ApiCallAttempt:
		if v == nil {
			err := EventValidationError{
				field:  "Details",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetApiCallAttempt()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "ApiCallAttempt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "ApiCallAttempt",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetApiCallAttempt()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					field:  "ApiCallAttempt",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Event_ApiCall:
		if v == nil {
			err := EventValidationError{
				field:  "Details",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetApiCall()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "ApiCall",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EventValidationError{
						field:  "ApiCall",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetApiCall()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EventValidationError{
					field:  "ApiCall",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return EventMultiError(errors)
	}

	return nil
}

// EventMultiError is an error wrapping multiple validation errors returned by
// Event.ValidateAll() if the designated constraints aren't met.
type EventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EventMultiError) AllErrors() []error { return m }

// EventValidationError is the validation error returned by Event.Validate if
// the designated constraints aren't met.
type EventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EventValidationError) ErrorName() string { return "EventValidationError" }

// Error satisfies the builtin error interface
func (e EventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EventValidationError{}

// Validate checks the field values on AssumeRoleEvent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AssumeRoleEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AssumeRoleEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AssumeRoleEventMultiError, or nil if none found.
func (m *AssumeRoleEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *AssumeRoleEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccountId

	// no validation rules for RoleName

	// no validation rules for AccessKeyId

	if len(errors) > 0 {
		return AssumeRoleEventMultiError(errors)
	}

	return nil
}

// AssumeRoleEventMultiError is an error wrapping multiple validation errors
// returned by AssumeRoleEvent.ValidateAll() if the designated constraints
// aren't met.
type AssumeRoleEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AssumeRoleEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AssumeRoleEventMultiError) AllErrors() []error { return m }

// AssumeRoleEventValidationError is the validation error returned by
// AssumeRoleEvent.Validate if the designated constraints aren't met.
type AssumeRoleEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AssumeRoleEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AssumeRoleEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AssumeRoleEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AssumeRoleEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AssumeRoleEventValidationError) ErrorName() string { return "AssumeRoleEventValidationError" }

// Error satisfies the builtin error interface
func (e AssumeRoleEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAssumeRoleEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AssumeRoleEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AssumeRoleEventValidationError{}

// Validate checks the field values on APICallAttemptEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *APICallAttemptEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on APICallAttemptEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// APICallAttemptEventMultiError, or nil if none found.
func (m *APICallAttemptEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *APICallAttemptEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientId

	// no validation rules for Service

	// no validation rules for Api

	// no validation rules for AttemptLatency

	// no validation rules for Fqdn

	// no validation rules for UserAgent

	// no validation rules for AccessKey

	// no validation rules for Region

	// no validation rules for HttpStatusCode

	// no validation rules for XAmzRequestId

	// no validation rules for XAmzId_2

	if len(errors) > 0 {
		return APICallAttemptEventMultiError(errors)
	}

	return nil
}

// APICallAttemptEventMultiError is an error wrapping multiple validation
// errors returned by APICallAttemptEvent.ValidateAll() if the designated
// constraints aren't met.
type APICallAttemptEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m APICallAttemptEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m APICallAttemptEventMultiError) AllErrors() []error { return m }

// APICallAttemptEventValidationError is the validation error returned by
// APICallAttemptEvent.Validate if the designated constraints aren't met.
type APICallAttemptEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e APICallAttemptEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e APICallAttemptEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e APICallAttemptEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e APICallAttemptEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e APICallAttemptEventValidationError) ErrorName() string {
	return "APICallAttemptEventValidationError"
}

// Error satisfies the builtin error interface
func (e APICallAttemptEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAPICallAttemptEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = APICallAttemptEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = APICallAttemptEventValidationError{}

// Validate checks the field values on APICallEvent with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *APICallEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on APICallEvent with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in APICallEventMultiError, or
// nil if none found.
func (m *APICallEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *APICallEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientId

	// no validation rules for Service

	// no validation rules for Api

	// no validation rules for AttemptCount

	// no validation rules for Region

	// no validation rules for UserAgent

	// no validation rules for FinalHttpStatusCode

	// no validation rules for AttemptLatency

	// no validation rules for MaxRetriesExceeded

	if len(errors) > 0 {
		return APICallEventMultiError(errors)
	}

	return nil
}

// APICallEventMultiError is an error wrapping multiple validation errors
// returned by APICallEvent.ValidateAll() if the designated constraints aren't met.
type APICallEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m APICallEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m APICallEventMultiError) AllErrors() []error { return m }

// APICallEventValidationError is the validation error returned by
// APICallEvent.Validate if the designated constraints aren't met.
type APICallEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e APICallEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e APICallEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e APICallEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e APICallEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e APICallEventValidationError) ErrorName() string { return "APICallEventValidationError" }

// Error satisfies the builtin error interface
func (e APICallEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAPICallEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = APICallEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = APICallEventValidationError{}
