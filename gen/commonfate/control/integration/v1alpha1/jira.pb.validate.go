// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/integration/v1alpha1/jira.proto

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

// Validate checks the field values on Jira with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Jira) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Jira with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in JiraMultiError, or nil if none found.
func (m *Jira) ValidateAll() error {
	return m.validate(true)
}

func (m *Jira) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientId

	// no validation rules for ClientSecretSecretPath

	// no validation rules for CloudId

	// no validation rules for JiraUrl

	if len(errors) > 0 {
		return JiraMultiError(errors)
	}

	return nil
}

// JiraMultiError is an error wrapping multiple validation errors returned by
// Jira.ValidateAll() if the designated constraints aren't met.
type JiraMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JiraMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JiraMultiError) AllErrors() []error { return m }

// JiraValidationError is the validation error returned by Jira.Validate if the
// designated constraints aren't met.
type JiraValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JiraValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JiraValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JiraValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JiraValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JiraValidationError) ErrorName() string { return "JiraValidationError" }

// Error satisfies the builtin error interface
func (e JiraValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJira.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JiraValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JiraValidationError{}
