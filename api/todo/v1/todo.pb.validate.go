// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: todo/v1/todo.proto

package todov1

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
)

// Validate checks the field values on Todo with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Todo) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Description

	// no validation rules for Completed

	if v, ok := interface{}(m.GetModifiedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TodoValidationError{
				field:  "ModifiedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TodoValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TodoValidationError is the validation error returned by Todo.Validate if the
// designated constraints aren't met.
type TodoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TodoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TodoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TodoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TodoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TodoValidationError) ErrorName() string { return "TodoValidationError" }

// Error satisfies the builtin error interface
func (e TodoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTodo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TodoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TodoValidationError{}
