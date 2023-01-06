// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: hub.proto

package pb

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

// Validate checks the field values on GetHubByIDRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetHubByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHubByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetHubByIDRequestMultiError, or nil if none found.
func (m *GetHubByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHubByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for HubID

	if len(errors) > 0 {
		return GetHubByIDRequestMultiError(errors)
	}

	return nil
}

// GetHubByIDRequestMultiError is an error wrapping multiple validation errors
// returned by GetHubByIDRequest.ValidateAll() if the designated constraints
// aren't met.
type GetHubByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHubByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHubByIDRequestMultiError) AllErrors() []error { return m }

// GetHubByIDRequestValidationError is the validation error returned by
// GetHubByIDRequest.Validate if the designated constraints aren't met.
type GetHubByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHubByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHubByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHubByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHubByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHubByIDRequestValidationError) ErrorName() string {
	return "GetHubByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetHubByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHubByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHubByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHubByIDRequestValidationError{}

// Validate checks the field values on GetHubByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetHubByIDResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHubByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetHubByIDResponseMultiError, or nil if none found.
func (m *GetHubByIDResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHubByIDResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetHubByIDResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetHubByIDResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetHubByIDResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetHubByIDResponseMultiError(errors)
	}

	return nil
}

// GetHubByIDResponseMultiError is an error wrapping multiple validation errors
// returned by GetHubByIDResponse.ValidateAll() if the designated constraints
// aren't met.
type GetHubByIDResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHubByIDResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHubByIDResponseMultiError) AllErrors() []error { return m }

// GetHubByIDResponseValidationError is the validation error returned by
// GetHubByIDResponse.Validate if the designated constraints aren't met.
type GetHubByIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHubByIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHubByIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHubByIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHubByIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHubByIDResponseValidationError) ErrorName() string {
	return "GetHubByIDResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetHubByIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHubByIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHubByIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHubByIDResponseValidationError{}

// Validate checks the field values on Hub with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Hub) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Hub with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HubMultiError, or nil if none found.
func (m *Hub) ValidateAll() error {
	return m.validate(true)
}

func (m *Hub) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for LocationId

	if len(errors) > 0 {
		return HubMultiError(errors)
	}

	return nil
}

// HubMultiError is an error wrapping multiple validation errors returned by
// Hub.ValidateAll() if the designated constraints aren't met.
type HubMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HubMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HubMultiError) AllErrors() []error { return m }

// HubValidationError is the validation error returned by Hub.Validate if the
// designated constraints aren't met.
type HubValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HubValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HubValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HubValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HubValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HubValidationError) ErrorName() string { return "HubValidationError" }

// Error satisfies the builtin error interface
func (e HubValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHub.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HubValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HubValidationError{}

// Validate checks the field values on CreateHubRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateHubRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateHubRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateHubRequestMultiError, or nil if none found.
func (m *CreateHubRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateHubRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHub()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateHubRequestValidationError{
					field:  "Hub",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateHubRequestValidationError{
					field:  "Hub",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHub()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateHubRequestValidationError{
				field:  "Hub",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateHubRequestMultiError(errors)
	}

	return nil
}

// CreateHubRequestMultiError is an error wrapping multiple validation errors
// returned by CreateHubRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateHubRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateHubRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateHubRequestMultiError) AllErrors() []error { return m }

// CreateHubRequestValidationError is the validation error returned by
// CreateHubRequest.Validate if the designated constraints aren't met.
type CreateHubRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateHubRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateHubRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateHubRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateHubRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateHubRequestValidationError) ErrorName() string { return "CreateHubRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateHubRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateHubRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateHubRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateHubRequestValidationError{}

// Validate checks the field values on CreateHubResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateHubResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateHubResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateHubResponseMultiError, or nil if none found.
func (m *CreateHubResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateHubResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return CreateHubResponseMultiError(errors)
	}

	return nil
}

// CreateHubResponseMultiError is an error wrapping multiple validation errors
// returned by CreateHubResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateHubResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateHubResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateHubResponseMultiError) AllErrors() []error { return m }

// CreateHubResponseValidationError is the validation error returned by
// CreateHubResponse.Validate if the designated constraints aren't met.
type CreateHubResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateHubResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateHubResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateHubResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateHubResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateHubResponseValidationError) ErrorName() string {
	return "CreateHubResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateHubResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateHubResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateHubResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateHubResponseValidationError{}

// Validate checks the field values on GetListHubRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetListHubRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListHubRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListHubRequestMultiError, or nil if none found.
func (m *GetListHubRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListHubRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Offset

	// no validation rules for Limit

	if len(errors) > 0 {
		return GetListHubRequestMultiError(errors)
	}

	return nil
}

// GetListHubRequestMultiError is an error wrapping multiple validation errors
// returned by GetListHubRequest.ValidateAll() if the designated constraints
// aren't met.
type GetListHubRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListHubRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListHubRequestMultiError) AllErrors() []error { return m }

// GetListHubRequestValidationError is the validation error returned by
// GetListHubRequest.Validate if the designated constraints aren't met.
type GetListHubRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListHubRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListHubRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListHubRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListHubRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListHubRequestValidationError) ErrorName() string {
	return "GetListHubRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetListHubRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListHubRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListHubRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListHubRequestValidationError{}

// Validate checks the field values on GetListHubResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListHubResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListHubResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListHubResponseMultiError, or nil if none found.
func (m *GetListHubResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListHubResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetListHubResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetListHubResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetListHubResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return GetListHubResponseMultiError(errors)
	}

	return nil
}

// GetListHubResponseMultiError is an error wrapping multiple validation errors
// returned by GetListHubResponse.ValidateAll() if the designated constraints
// aren't met.
type GetListHubResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListHubResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListHubResponseMultiError) AllErrors() []error { return m }

// GetListHubResponseValidationError is the validation error returned by
// GetListHubResponse.Validate if the designated constraints aren't met.
type GetListHubResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListHubResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListHubResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListHubResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListHubResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListHubResponseValidationError) ErrorName() string {
	return "GetListHubResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetListHubResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListHubResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListHubResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListHubResponseValidationError{}

// Validate checks the field values on UpdateHubRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateHubRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateHubRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateHubRequestMultiError, or nil if none found.
func (m *UpdateHubRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateHubRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHub()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateHubRequestValidationError{
					field:  "Hub",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateHubRequestValidationError{
					field:  "Hub",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHub()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateHubRequestValidationError{
				field:  "Hub",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateHubRequestMultiError(errors)
	}

	return nil
}

// UpdateHubRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateHubRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateHubRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateHubRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateHubRequestMultiError) AllErrors() []error { return m }

// UpdateHubRequestValidationError is the validation error returned by
// UpdateHubRequest.Validate if the designated constraints aren't met.
type UpdateHubRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateHubRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateHubRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateHubRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateHubRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateHubRequestValidationError) ErrorName() string { return "UpdateHubRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateHubRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateHubRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateHubRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateHubRequestValidationError{}

// Validate checks the field values on UpdateHubResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateHubResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateHubResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateHubResponseMultiError, or nil if none found.
func (m *UpdateHubResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateHubResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return UpdateHubResponseMultiError(errors)
	}

	return nil
}

// UpdateHubResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateHubResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateHubResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateHubResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateHubResponseMultiError) AllErrors() []error { return m }

// UpdateHubResponseValidationError is the validation error returned by
// UpdateHubResponse.Validate if the designated constraints aren't met.
type UpdateHubResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateHubResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateHubResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateHubResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateHubResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateHubResponseValidationError) ErrorName() string {
	return "UpdateHubResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateHubResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateHubResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateHubResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateHubResponseValidationError{}