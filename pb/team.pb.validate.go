// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: team.proto

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

// Validate checks the field values on GetTeamByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTeamByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTeamByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTeamByIDRequestMultiError, or nil if none found.
func (m *GetTeamByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTeamByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TeamID

	if len(errors) > 0 {
		return GetTeamByIDRequestMultiError(errors)
	}

	return nil
}

// GetTeamByIDRequestMultiError is an error wrapping multiple validation errors
// returned by GetTeamByIDRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTeamByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTeamByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTeamByIDRequestMultiError) AllErrors() []error { return m }

// GetTeamByIDRequestValidationError is the validation error returned by
// GetTeamByIDRequest.Validate if the designated constraints aren't met.
type GetTeamByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTeamByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTeamByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTeamByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTeamByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTeamByIDRequestValidationError) ErrorName() string {
	return "GetTeamByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTeamByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTeamByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTeamByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTeamByIDRequestValidationError{}

// Validate checks the field values on GetTeamByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTeamByIDResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTeamByIDResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTeamByIDResponseMultiError, or nil if none found.
func (m *GetTeamByIDResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTeamByIDResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTeamByIDResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTeamByIDResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTeamByIDResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTeamByIDResponseMultiError(errors)
	}

	return nil
}

// GetTeamByIDResponseMultiError is an error wrapping multiple validation
// errors returned by GetTeamByIDResponse.ValidateAll() if the designated
// constraints aren't met.
type GetTeamByIDResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTeamByIDResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTeamByIDResponseMultiError) AllErrors() []error { return m }

// GetTeamByIDResponseValidationError is the validation error returned by
// GetTeamByIDResponse.Validate if the designated constraints aren't met.
type GetTeamByIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTeamByIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTeamByIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTeamByIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTeamByIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTeamByIDResponseValidationError) ErrorName() string {
	return "GetTeamByIDResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTeamByIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTeamByIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTeamByIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTeamByIDResponseValidationError{}

// Validate checks the field values on Team with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Team) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Team with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TeamMultiError, or nil if none found.
func (m *Team) ValidateAll() error {
	return m.validate(true)
}

func (m *Team) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Type

	// no validation rules for LocationId

	// no validation rules for HubId

	if len(errors) > 0 {
		return TeamMultiError(errors)
	}

	return nil
}

// TeamMultiError is an error wrapping multiple validation errors returned by
// Team.ValidateAll() if the designated constraints aren't met.
type TeamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TeamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TeamMultiError) AllErrors() []error { return m }

// TeamValidationError is the validation error returned by Team.Validate if the
// designated constraints aren't met.
type TeamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TeamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TeamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TeamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TeamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TeamValidationError) ErrorName() string { return "TeamValidationError" }

// Error satisfies the builtin error interface
func (e TeamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTeam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TeamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TeamValidationError{}

// Validate checks the field values on CreateTeamRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateTeamRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTeamRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTeamRequestMultiError, or nil if none found.
func (m *CreateTeamRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTeamRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTeam()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateTeamRequestValidationError{
					field:  "Team",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateTeamRequestValidationError{
					field:  "Team",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTeam()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateTeamRequestValidationError{
				field:  "Team",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateTeamRequestMultiError(errors)
	}

	return nil
}

// CreateTeamRequestMultiError is an error wrapping multiple validation errors
// returned by CreateTeamRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateTeamRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTeamRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTeamRequestMultiError) AllErrors() []error { return m }

// CreateTeamRequestValidationError is the validation error returned by
// CreateTeamRequest.Validate if the designated constraints aren't met.
type CreateTeamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTeamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTeamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTeamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTeamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTeamRequestValidationError) ErrorName() string {
	return "CreateTeamRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTeamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTeamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTeamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTeamRequestValidationError{}

// Validate checks the field values on CreateTeamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateTeamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateTeamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateTeamResponseMultiError, or nil if none found.
func (m *CreateTeamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateTeamResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return CreateTeamResponseMultiError(errors)
	}

	return nil
}

// CreateTeamResponseMultiError is an error wrapping multiple validation errors
// returned by CreateTeamResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateTeamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateTeamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateTeamResponseMultiError) AllErrors() []error { return m }

// CreateTeamResponseValidationError is the validation error returned by
// CreateTeamResponse.Validate if the designated constraints aren't met.
type CreateTeamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateTeamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateTeamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateTeamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateTeamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateTeamResponseValidationError) ErrorName() string {
	return "CreateTeamResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateTeamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateTeamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateTeamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateTeamResponseValidationError{}

// Validate checks the field values on GetListTeamRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListTeamRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListTeamRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListTeamRequestMultiError, or nil if none found.
func (m *GetListTeamRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListTeamRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Offset

	// no validation rules for Limit

	if len(errors) > 0 {
		return GetListTeamRequestMultiError(errors)
	}

	return nil
}

// GetListTeamRequestMultiError is an error wrapping multiple validation errors
// returned by GetListTeamRequest.ValidateAll() if the designated constraints
// aren't met.
type GetListTeamRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListTeamRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListTeamRequestMultiError) AllErrors() []error { return m }

// GetListTeamRequestValidationError is the validation error returned by
// GetListTeamRequest.Validate if the designated constraints aren't met.
type GetListTeamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListTeamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListTeamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListTeamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListTeamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListTeamRequestValidationError) ErrorName() string {
	return "GetListTeamRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetListTeamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListTeamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListTeamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListTeamRequestValidationError{}

// Validate checks the field values on GetListTeamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetListTeamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetListTeamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetListTeamResponseMultiError, or nil if none found.
func (m *GetListTeamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetListTeamResponse) validate(all bool) error {
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
					errors = append(errors, GetListTeamResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetListTeamResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetListTeamResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return GetListTeamResponseMultiError(errors)
	}

	return nil
}

// GetListTeamResponseMultiError is an error wrapping multiple validation
// errors returned by GetListTeamResponse.ValidateAll() if the designated
// constraints aren't met.
type GetListTeamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetListTeamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetListTeamResponseMultiError) AllErrors() []error { return m }

// GetListTeamResponseValidationError is the validation error returned by
// GetListTeamResponse.Validate if the designated constraints aren't met.
type GetListTeamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetListTeamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetListTeamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetListTeamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetListTeamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetListTeamResponseValidationError) ErrorName() string {
	return "GetListTeamResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetListTeamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetListTeamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetListTeamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetListTeamResponseValidationError{}

// Validate checks the field values on UpdateTeamRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateTeamRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTeamRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTeamRequestMultiError, or nil if none found.
func (m *UpdateTeamRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTeamRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTeam()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateTeamRequestValidationError{
					field:  "Team",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateTeamRequestValidationError{
					field:  "Team",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTeam()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateTeamRequestValidationError{
				field:  "Team",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateTeamRequestMultiError(errors)
	}

	return nil
}

// UpdateTeamRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateTeamRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateTeamRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTeamRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTeamRequestMultiError) AllErrors() []error { return m }

// UpdateTeamRequestValidationError is the validation error returned by
// UpdateTeamRequest.Validate if the designated constraints aren't met.
type UpdateTeamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTeamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTeamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTeamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTeamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTeamRequestValidationError) ErrorName() string {
	return "UpdateTeamRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTeamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTeamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTeamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTeamRequestValidationError{}

// Validate checks the field values on UpdateTeamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTeamResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTeamResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTeamResponseMultiError, or nil if none found.
func (m *UpdateTeamResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTeamResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return UpdateTeamResponseMultiError(errors)
	}

	return nil
}

// UpdateTeamResponseMultiError is an error wrapping multiple validation errors
// returned by UpdateTeamResponse.ValidateAll() if the designated constraints
// aren't met.
type UpdateTeamResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTeamResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTeamResponseMultiError) AllErrors() []error { return m }

// UpdateTeamResponseValidationError is the validation error returned by
// UpdateTeamResponse.Validate if the designated constraints aren't met.
type UpdateTeamResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTeamResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTeamResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTeamResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTeamResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTeamResponseValidationError) ErrorName() string {
	return "UpdateTeamResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTeamResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTeamResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTeamResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTeamResponseValidationError{}
