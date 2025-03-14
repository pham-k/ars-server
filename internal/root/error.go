package root

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
)

// Application error codes.
//
// NOTE: These are meant to be generic and they map well to HTTP error codes.
// Different applications can have very different error code requirements so
// these should be expanded as needed (or introduce subcodes).
const (
	ErrConflict        = "conflict"
	ErrInternal        = "internal"
	ErrInvalid         = "invalid"
	ErrNotFound        = "not_found"
	ErrUnauthenticated = "unauthenticated"
	ErrUnauthorized    = "unauthorized"
)

// Error represents an application-specific error. Application errors can be
// unwrapped by the caller to extract out the code & message.
//
// Any non-application error (such as a disk error) should be reported as an
// ErrInternal error and the human user should only see "Internal error" as the
// message. These low-level internal error details should only be logged and
// reported to the operator of the application (not the end user).
type Error struct {
	// Machine-readable error code.
	Code string

	// Human-readable message.
	Message string

	// Logical operation and nested error.
	Op  string
	Err error
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		_, err := fmt.Fprintf(&buf, "%s: ", e.Op)
		if err != nil {
			slog.Error("Fprintf", err)
		}
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			_, err := fmt.Fprintf(&buf, "<%s> ", e.Code)
			if err != nil {
				slog.Error("Fprintf", err)
			}
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

// ErrorCode returns the code of the root error, if available. Otherwise, returns ErrInternal.
func ErrorCode(err error) string {
	if err == nil {
		return ""
	}
	var e *Error
	if errors.As(err, &e) && e.Code != "" {
		return e.Code
	}
	return ErrInternal
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise, returns a generic error message.
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	var e *Error
	if errors.As(err, &e) && e.Message != "" {
		return e.Message
	}
	return "An internal error has occurred"
}
