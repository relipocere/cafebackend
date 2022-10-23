// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphmodel

import (
	"fmt"
	"io"
	"strconv"
)

type CreateUserInput struct {
	Username string       `json:"username"`
	Password string       `json:"password"`
	Kind     UserKindEnum `json:"kind"`
	FullName string       `json:"fullName"`
}

type GetAuthTokenInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetAuthTokenPayload struct {
	Token string `json:"token"`
}

type User struct {
	UUID     string       `json:"uuid"`
	Username string       `json:"username"`
	Kind     UserKindEnum `json:"kind"`
	FullName string       `json:"fullName"`
}

type ErrorCode string

const (
	ErrorCodeInternal           ErrorCode = "INTERNAL"
	ErrorCodeUnauthenticated    ErrorCode = "UNAUTHENTICATED"
	ErrorCodeUnauthorized       ErrorCode = "UNAUTHORIZED"
	ErrorCodeBadRequest         ErrorCode = "BAD_REQUEST"
	ErrorCodeNotFound           ErrorCode = "NOT_FOUND"
	ErrorCodeFailedPrecondition ErrorCode = "FAILED_PRECONDITION"
)

var AllErrorCode = []ErrorCode{
	ErrorCodeInternal,
	ErrorCodeUnauthenticated,
	ErrorCodeUnauthorized,
	ErrorCodeBadRequest,
	ErrorCodeNotFound,
	ErrorCodeFailedPrecondition,
}

func (e ErrorCode) IsValid() bool {
	switch e {
	case ErrorCodeInternal, ErrorCodeUnauthenticated, ErrorCodeUnauthorized, ErrorCodeBadRequest, ErrorCodeNotFound, ErrorCodeFailedPrecondition:
		return true
	}
	return false
}

func (e ErrorCode) String() string {
	return string(e)
}

func (e *ErrorCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ErrorCode", str)
	}
	return nil
}

func (e ErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserKindEnum string

const (
	UserKindEnumConsumer UserKindEnum = "CONSUMER"
	UserKindEnumBusiness UserKindEnum = "BUSINESS"
)

var AllUserKindEnum = []UserKindEnum{
	UserKindEnumConsumer,
	UserKindEnumBusiness,
}

func (e UserKindEnum) IsValid() bool {
	switch e {
	case UserKindEnumConsumer, UserKindEnumBusiness:
		return true
	}
	return false
}

func (e UserKindEnum) String() string {
	return string(e)
}

func (e *UserKindEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserKindEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserKindEnum", str)
	}
	return nil
}

func (e UserKindEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
