package model

import (
	"fmt"
)

// ErrorCode is the type of error.
type ErrorCode int64

const (
	// ErrorCodeUnauthenticated no credentials
	ErrorCodeUnauthenticated ErrorCode = iota + 1
	// ErrorCodeUnauthorized no rights to perform operation
	ErrorCodeUnauthorized
	// ErrorCodeBadRequest invalid request data
	ErrorCodeBadRequest
	// ErrorCodeNotFound resource is not found
	ErrorCodeNotFound
	// ErrorCodeFailedPrecondition preconditions are not met
	ErrorCodeFailedPrecondition
)

// Error is the user-handler readable error.
type Error struct {
	Message string
	Code    ErrorCode
}

// Error implements error interface.
func (e Error) Error() string {
	return e.Message
}

// ErrMessageMissingFieldRequired is the generic "field is not provided" message.
func ErrMessageMissingFieldRequired(field string) string {
	return fmt.Sprintf("Field '%s' is required but not provided", field)
}

// ErrMessageSurpassedMaxLength is the generic "field is not provided" message.
func ErrMessageMaxLengthExceeded(field string, length int64) string {
	return fmt.Sprintf("Field '%s' exceeds maximum limit of %d characters", field, length)
}

// ErrMessageInvalidID is the generic "bad id" message.
func ErrMessageInvalidID [T any](field string, value T) string{
	return fmt.Sprintf("Field '%s' has invalid identifier(s): '%v'", field, value)
}

// ErrMessageUnauthenticated is the generic "unauthenticated" message.
func ErrMessageUnauthenticated() string {
	return "Unauthenticated"
}
