// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/integration/v1alpha1/snowflake.proto

package integrationv1alpha1

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

// Validate checks the field values on Snowflake with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Snowflake) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Snowflake with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SnowflakeMultiError, or nil
// if none found.
func (m *Snowflake) ValidateAll() error {
	return m.validate(true)
}

func (m *Snowflake) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccountId

	// no validation rules for Region

	// no validation rules for Username

	// no validation rules for PasswordSecretPath

	if len(errors) > 0 {
		return SnowflakeMultiError(errors)
	}

	return nil
}

// SnowflakeMultiError is an error wrapping multiple validation errors returned
// by Snowflake.ValidateAll() if the designated constraints aren't met.
type SnowflakeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SnowflakeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SnowflakeMultiError) AllErrors() []error { return m }

// SnowflakeValidationError is the validation error returned by
// Snowflake.Validate if the designated constraints aren't met.
type SnowflakeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SnowflakeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SnowflakeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SnowflakeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SnowflakeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SnowflakeValidationError) ErrorName() string { return "SnowflakeValidationError" }

// Error satisfies the builtin error interface
func (e SnowflakeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSnowflake.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SnowflakeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SnowflakeValidationError{}
