// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/meta_protocol_proxy/filters/global_ratelimit/v1alpha/global_ratelimit.proto

package aeraki_meta_protocol_proxy_filters_ratelimit_v1alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on RateLimit with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RateLimit) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitValidationError{
				field:  "Match",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetDomain()) < 1 {
		return RateLimitValidationError{
			field:  "Domain",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FailureModeDeny

	if m.GetRateLimitService() == nil {
		return RateLimitValidationError{
			field:  "RateLimitService",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRateLimitService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitValidationError{
				field:  "RateLimitService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetDescriptors()) < 1 {
		return RateLimitValidationError{
			field:  "Descriptors",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetDescriptors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitValidationError{
					field:  fmt.Sprintf("Descriptors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RateLimitValidationError is the validation error returned by
// RateLimit.Validate if the designated constraints aren't met.
type RateLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitValidationError) ErrorName() string { return "RateLimitValidationError" }

// Error satisfies the builtin error interface
func (e RateLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitValidationError{}

// Validate checks the field values on Descriptor with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Descriptor) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetProperty()) < 1 {
		return DescriptorValidationError{
			field:  "Property",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDescriptorKey()) < 1 {
		return DescriptorValidationError{
			field:  "DescriptorKey",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// DescriptorValidationError is the validation error returned by
// Descriptor.Validate if the designated constraints aren't met.
type DescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescriptorValidationError) ErrorName() string { return "DescriptorValidationError" }

// Error satisfies the builtin error interface
func (e DescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescriptorValidationError{}
