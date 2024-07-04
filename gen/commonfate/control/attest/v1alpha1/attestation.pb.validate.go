// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/attest/v1alpha1/attestation.proto

package attestv1alpha1

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

// Validate checks the field values on Header with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Header) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Header with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HeaderMultiError, or nil if none found.
func (m *Header) ValidateAll() error {
	return m.validate(true)
}

func (m *Header) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Version

	// no validation rules for Timestamp

	// no validation rules for Type

	// no validation rules for ContentDigest

	// no validation rules for Kid

	if len(errors) > 0 {
		return HeaderMultiError(errors)
	}

	return nil
}

// HeaderMultiError is an error wrapping multiple validation errors returned by
// Header.ValidateAll() if the designated constraints aren't met.
type HeaderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HeaderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HeaderMultiError) AllErrors() []error { return m }

// HeaderValidationError is the validation error returned by Header.Validate if
// the designated constraints aren't met.
type HeaderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderValidationError) ErrorName() string { return "HeaderValidationError" }

// Error satisfies the builtin error interface
func (e HeaderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeader.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderValidationError{}

// Validate checks the field values on Device with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Device) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Device with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DeviceMultiError, or nil if none found.
func (m *Device) ValidateAll() error {
	return m.validate(true)
}

func (m *Device) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SerialNumber

	// no validation rules for Uuid

	// no validation rules for Platform

	// no validation rules for Architecture

	// no validation rules for Hostname

	// no validation rules for KernelRelease

	// no validation rules for KernelVersion

	if len(errors) > 0 {
		return DeviceMultiError(errors)
	}

	return nil
}

// DeviceMultiError is an error wrapping multiple validation errors returned by
// Device.ValidateAll() if the designated constraints aren't met.
type DeviceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeviceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeviceMultiError) AllErrors() []error { return m }

// DeviceValidationError is the validation error returned by Device.Validate if
// the designated constraints aren't met.
type DeviceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceValidationError) ErrorName() string { return "DeviceValidationError" }

// Error satisfies the builtin error interface
func (e DeviceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDevice.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceValidationError{}

// Validate checks the field values on Attestation with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Attestation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Attestation with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AttestationMultiError, or
// nil if none found.
func (m *Attestation) ValidateAll() error {
	return m.validate(true)
}

func (m *Attestation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHeader()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AttestationValidationError{
					field:  "Header",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AttestationValidationError{
					field:  "Header",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHeader()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttestationValidationError{
				field:  "Header",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Signature

	if len(errors) > 0 {
		return AttestationMultiError(errors)
	}

	return nil
}

// AttestationMultiError is an error wrapping multiple validation errors
// returned by Attestation.ValidateAll() if the designated constraints aren't met.
type AttestationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AttestationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AttestationMultiError) AllErrors() []error { return m }

// AttestationValidationError is the validation error returned by
// Attestation.Validate if the designated constraints aren't met.
type AttestationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttestationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttestationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttestationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttestationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttestationValidationError) ErrorName() string { return "AttestationValidationError" }

// Error satisfies the builtin error interface
func (e AttestationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttestation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttestationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttestationValidationError{}