// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/employee.proto

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

// Validate checks the field values on EmployeeCreateCommand with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EmployeeCreateCommand) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmployeeCreateCommand with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EmployeeCreateCommandMultiError, or nil if none found.
func (m *EmployeeCreateCommand) ValidateAll() error {
	return m.validate(true)
}

func (m *EmployeeCreateCommand) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 3 {
		err := EmployeeCreateCommandValidationError{
			field:  "Name",
			reason: "value length must be at least 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return EmployeeCreateCommandMultiError(errors)
	}

	return nil
}

// EmployeeCreateCommandMultiError is an error wrapping multiple validation
// errors returned by EmployeeCreateCommand.ValidateAll() if the designated
// constraints aren't met.
type EmployeeCreateCommandMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmployeeCreateCommandMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmployeeCreateCommandMultiError) AllErrors() []error { return m }

// EmployeeCreateCommandValidationError is the validation error returned by
// EmployeeCreateCommand.Validate if the designated constraints aren't met.
type EmployeeCreateCommandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmployeeCreateCommandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmployeeCreateCommandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmployeeCreateCommandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmployeeCreateCommandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmployeeCreateCommandValidationError) ErrorName() string {
	return "EmployeeCreateCommandValidationError"
}

// Error satisfies the builtin error interface
func (e EmployeeCreateCommandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmployeeCreateCommand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmployeeCreateCommandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmployeeCreateCommandValidationError{}

// Validate checks the field values on EmployeeCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EmployeeCreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmployeeCreateResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EmployeeCreateResponseMultiError, or nil if none found.
func (m *EmployeeCreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EmployeeCreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if len(errors) > 0 {
		return EmployeeCreateResponseMultiError(errors)
	}

	return nil
}

// EmployeeCreateResponseMultiError is an error wrapping multiple validation
// errors returned by EmployeeCreateResponse.ValidateAll() if the designated
// constraints aren't met.
type EmployeeCreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmployeeCreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmployeeCreateResponseMultiError) AllErrors() []error { return m }

// EmployeeCreateResponseValidationError is the validation error returned by
// EmployeeCreateResponse.Validate if the designated constraints aren't met.
type EmployeeCreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmployeeCreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmployeeCreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmployeeCreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmployeeCreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmployeeCreateResponseValidationError) ErrorName() string {
	return "EmployeeCreateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EmployeeCreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmployeeCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmployeeCreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmployeeCreateResponseValidationError{}

// Validate checks the field values on EmployeeDeleteCommand with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EmployeeDeleteCommand) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmployeeDeleteCommand with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EmployeeDeleteCommandMultiError, or nil if none found.
func (m *EmployeeDeleteCommand) ValidateAll() error {
	return m.validate(true)
}

func (m *EmployeeDeleteCommand) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return EmployeeDeleteCommandMultiError(errors)
	}

	return nil
}

// EmployeeDeleteCommandMultiError is an error wrapping multiple validation
// errors returned by EmployeeDeleteCommand.ValidateAll() if the designated
// constraints aren't met.
type EmployeeDeleteCommandMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmployeeDeleteCommandMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmployeeDeleteCommandMultiError) AllErrors() []error { return m }

// EmployeeDeleteCommandValidationError is the validation error returned by
// EmployeeDeleteCommand.Validate if the designated constraints aren't met.
type EmployeeDeleteCommandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmployeeDeleteCommandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmployeeDeleteCommandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmployeeDeleteCommandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmployeeDeleteCommandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmployeeDeleteCommandValidationError) ErrorName() string {
	return "EmployeeDeleteCommandValidationError"
}

// Error satisfies the builtin error interface
func (e EmployeeDeleteCommandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmployeeDeleteCommand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmployeeDeleteCommandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmployeeDeleteCommandValidationError{}

// Validate checks the field values on EmployeeAssignProjectCommand with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EmployeeAssignProjectCommand) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmployeeAssignProjectCommand with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EmployeeAssignProjectCommandMultiError, or nil if none found.
func (m *EmployeeAssignProjectCommand) ValidateAll() error {
	return m.validate(true)
}

func (m *EmployeeAssignProjectCommand) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ProjectId

	if len(errors) > 0 {
		return EmployeeAssignProjectCommandMultiError(errors)
	}

	return nil
}

// EmployeeAssignProjectCommandMultiError is an error wrapping multiple
// validation errors returned by EmployeeAssignProjectCommand.ValidateAll() if
// the designated constraints aren't met.
type EmployeeAssignProjectCommandMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmployeeAssignProjectCommandMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmployeeAssignProjectCommandMultiError) AllErrors() []error { return m }

// EmployeeAssignProjectCommandValidationError is the validation error returned
// by EmployeeAssignProjectCommand.Validate if the designated constraints
// aren't met.
type EmployeeAssignProjectCommandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmployeeAssignProjectCommandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmployeeAssignProjectCommandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmployeeAssignProjectCommandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmployeeAssignProjectCommandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmployeeAssignProjectCommandValidationError) ErrorName() string {
	return "EmployeeAssignProjectCommandValidationError"
}

// Error satisfies the builtin error interface
func (e EmployeeAssignProjectCommandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmployeeAssignProjectCommand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmployeeAssignProjectCommandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmployeeAssignProjectCommandValidationError{}

// Validate checks the field values on EmployeeUnassignProjectCommand with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EmployeeUnassignProjectCommand) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmployeeUnassignProjectCommand with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// EmployeeUnassignProjectCommandMultiError, or nil if none found.
func (m *EmployeeUnassignProjectCommand) ValidateAll() error {
	return m.validate(true)
}

func (m *EmployeeUnassignProjectCommand) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return EmployeeUnassignProjectCommandMultiError(errors)
	}

	return nil
}

// EmployeeUnassignProjectCommandMultiError is an error wrapping multiple
// validation errors returned by EmployeeUnassignProjectCommand.ValidateAll()
// if the designated constraints aren't met.
type EmployeeUnassignProjectCommandMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmployeeUnassignProjectCommandMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmployeeUnassignProjectCommandMultiError) AllErrors() []error { return m }

// EmployeeUnassignProjectCommandValidationError is the validation error
// returned by EmployeeUnassignProjectCommand.Validate if the designated
// constraints aren't met.
type EmployeeUnassignProjectCommandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmployeeUnassignProjectCommandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmployeeUnassignProjectCommandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmployeeUnassignProjectCommandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmployeeUnassignProjectCommandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmployeeUnassignProjectCommandValidationError) ErrorName() string {
	return "EmployeeUnassignProjectCommandValidationError"
}

// Error satisfies the builtin error interface
func (e EmployeeUnassignProjectCommandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmployeeUnassignProjectCommand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmployeeUnassignProjectCommandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmployeeUnassignProjectCommandValidationError{}

// Validate checks the field values on Feedback with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Feedback) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Feedback with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FeedbackMultiError, or nil
// if none found.
func (m *Feedback) ValidateAll() error {
	return m.validate(true)
}

func (m *Feedback) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FeedbackType

	// no validation rules for EmployeeId

	// no validation rules for ItemId

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FeedbackValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FeedbackValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedbackValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FeedbackMultiError(errors)
	}

	return nil
}

// FeedbackMultiError is an error wrapping multiple validation errors returned
// by Feedback.ValidateAll() if the designated constraints aren't met.
type FeedbackMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeedbackMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeedbackMultiError) AllErrors() []error { return m }

// FeedbackValidationError is the validation error returned by
// Feedback.Validate if the designated constraints aren't met.
type FeedbackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedbackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedbackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedbackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedbackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedbackValidationError) ErrorName() string { return "FeedbackValidationError" }

// Error satisfies the builtin error interface
func (e FeedbackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedback.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedbackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedbackValidationError{}

// Validate checks the field values on EmployeeInsertFeedbackCommand with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EmployeeInsertFeedbackCommand) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EmployeeInsertFeedbackCommand with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// EmployeeInsertFeedbackCommandMultiError, or nil if none found.
func (m *EmployeeInsertFeedbackCommand) ValidateAll() error {
	return m.validate(true)
}

func (m *EmployeeInsertFeedbackCommand) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetFeedback()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EmployeeInsertFeedbackCommandValidationError{
					field:  "Feedback",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EmployeeInsertFeedbackCommandValidationError{
					field:  "Feedback",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFeedback()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EmployeeInsertFeedbackCommandValidationError{
				field:  "Feedback",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EmployeeInsertFeedbackCommandMultiError(errors)
	}

	return nil
}

// EmployeeInsertFeedbackCommandMultiError is an error wrapping multiple
// validation errors returned by EmployeeInsertFeedbackCommand.ValidateAll()
// if the designated constraints aren't met.
type EmployeeInsertFeedbackCommandMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmployeeInsertFeedbackCommandMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmployeeInsertFeedbackCommandMultiError) AllErrors() []error { return m }

// EmployeeInsertFeedbackCommandValidationError is the validation error
// returned by EmployeeInsertFeedbackCommand.Validate if the designated
// constraints aren't met.
type EmployeeInsertFeedbackCommandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmployeeInsertFeedbackCommandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmployeeInsertFeedbackCommandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmployeeInsertFeedbackCommandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmployeeInsertFeedbackCommandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmployeeInsertFeedbackCommandValidationError) ErrorName() string {
	return "EmployeeInsertFeedbackCommandValidationError"
}

// Error satisfies the builtin error interface
func (e EmployeeInsertFeedbackCommandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmployeeInsertFeedbackCommand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmployeeInsertFeedbackCommandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmployeeInsertFeedbackCommandValidationError{}
