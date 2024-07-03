// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commonfate/control/integration/v1alpha1/aws_idc.proto

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

// Validate checks the field values on AWSIDC with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AWSIDC) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AWSIDC with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AWSIDCMultiError, or nil if none found.
func (m *AWSIDC) ValidateAll() error {
	return m.validate(true)
}

func (m *AWSIDC) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SsoInstanceArn

	// no validation rules for IdentityStoreId

	// no validation rules for SsoRegion

	// no validation rules for ReaderRoleArn

	// no validation rules for AuditRoleName

	// no validation rules for SsoAccessPortalUrl

	// no validation rules for ProvisionerRoleArn

	if len(errors) > 0 {
		return AWSIDCMultiError(errors)
	}

	return nil
}

// AWSIDCMultiError is an error wrapping multiple validation errors returned by
// AWSIDC.ValidateAll() if the designated constraints aren't met.
type AWSIDCMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AWSIDCMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AWSIDCMultiError) AllErrors() []error { return m }

// AWSIDCValidationError is the validation error returned by AWSIDC.Validate if
// the designated constraints aren't met.
type AWSIDCValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AWSIDCValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AWSIDCValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AWSIDCValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AWSIDCValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AWSIDCValidationError) ErrorName() string { return "AWSIDCValidationError" }

// Error satisfies the builtin error interface
func (e AWSIDCValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAWSIDC.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AWSIDCValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AWSIDCValidationError{}
