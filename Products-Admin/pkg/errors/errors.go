// Package errors implements functions to manipulate errors.
package errors

import (
	"errors"
	"net/http"
)

// Application errors
var (
	ErrInvalid           = errors.New("validation failed")
	ErrUnauthorized      = errors.New("access denied")
	ErrForbidden         = errors.New("forbidden")
	ErrNotFound          = errors.New("not found")
	ErrInternal          = errors.New("internal system error")
	ErrTemporaryDisabled = errors.New("temporary disabled")
	ErrTimeout           = errors.New("timeout")
)

type AppError struct {
	err       error
	errorCode int
}

// New returns new app error that formats as the given text.
func New(message string, errCode int) *AppError {
	return newAppError(errors.New(message), errCode)
}

// Wrap returns new app error wrapping target error.
// If passed value is nil will fallback to internal
func Wrap(err error) *AppError {
	return newAppError(err, http.StatusInternalServerError)
}

func newAppError(err error, errCode int) *AppError {
	if err == nil {
		err = ErrInternal
	}

	return &AppError{
		err:       err,
		errorCode: errCode,
	}
}

// Error returns the string representation of the error message.
func (e *AppError) Error() string {
	return e.err.Error()
}

// UnWraps returns the actual err of type error
func (e *AppError) Unwrap() error {
	return e.err
}

func (e *AppError) GetErrCode() int {
	return e.errorCode
}
