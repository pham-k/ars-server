package error

import (
	"fmt"
	"net/http"
)

type Code string

const (
	CodeConflict Code = "conflict"
	CodeInternal Code = "internal"
	CodeInvalid  Code = "invalid"
	CodeNotFound Code = "not_found"
)

type CoreError interface {
	error
	AddOp(op string)
	GetOp() string
	ToHttpStatus() int
	Unwrap() error
}

type coreError struct {
	// Machine-readable error code.
	Code Code

	// Human-readable message.
	Message string
	Op      string

	Err error
}

type Option func(*coreError)

func New(message string, opts ...Option) CoreError {
	coreErr := &coreError{
		Message: message,
		Code:    CodeInternal,
	}

	for _, opt := range opts {
		opt(coreErr)
	}

	return coreErr
}

func WithCode(code Code) Option {
	return func(e *coreError) {
		e.Code = code
	}
}

func WithErr(err error) Option {
	return func(e *coreError) {
		e.Err = err
	}
}

func (e *coreError) AddOp(op string) {
	if e.Op == "" {
		e.Op = op
	} else {
		e.Op = fmt.Sprintf("%s > %s", e.Op, op)
	}
}

func (e *coreError) GetOp() string {
	return e.Op
}

func (e *coreError) Unwrap() error {
	return e.Err
}

func (e *coreError) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}

func (e *coreError) ToHttpStatus() int {
	switch e.Code {
	case CodeInvalid:
		return http.StatusBadRequest
	case CodeNotFound:
		return http.StatusNotFound
	case CodeConflict:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
