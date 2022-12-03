// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphmodel

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CreateProductInput struct {
	Name        string   `json:"name"`
	StoreID     int64    `json:"storeID"`
	Ingredients []string `json:"ingredients"`
	Calories    int64    `json:"calories"`
	ImageID     string   `json:"imageID"`
}

type CreateStoreInput struct {
	Title         string        `json:"title"`
	Affordability Affordability `json:"affordability"`
	CuisineType   CuisineType   `json:"cuisineType"`
	ImageID       string        `json:"imageID"`
}

type CreateUserInput struct {
	Username string       `json:"username"`
	Password string       `json:"password"`
	Kind     UserKindEnum `json:"kind"`
	FullName string       `json:"fullName"`
}

type DeleteStoreInput struct {
	ID int64 `json:"id"`
}

type GetAuthTokenInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetAuthTokenPayload struct {
	Token string `json:"token"`
}

type IntRange struct {
	Start          *int64 `json:"start"`
	End            *int64 `json:"end"`
	EndExclusive   bool   `json:"endExclusive"`
	StartExclusive bool   `json:"startExclusive"`
}

type Pagination struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	StoreID     int64     `json:"storeID"`
	Ingredients []string  `json:"ingredients"`
	Calories    int64     `json:"calories"`
	ImageID     string    `json:"imageID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type SearchStoresInput struct {
	Page           Pagination      `json:"page"`
	TitleQuery     *string         `json:"titleQuery"`
	Rating         *IntRange       `json:"rating"`
	OwnerUsernames []string        `json:"ownerUsernames"`
	Affordability  []Affordability `json:"affordability"`
	Cuisines       []CuisineType   `json:"cuisines"`
}

type Store struct {
	ID            int64         `json:"id"`
	Title         string        `json:"title"`
	Affordability Affordability `json:"affordability"`
	CuisineType   CuisineType   `json:"cuisineType"`
	OwnerUsername string        `json:"ownerUsername"`
	ImageID       string        `json:"imageID"`
	AverageRating int64         `json:"averageRating"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}

type User struct {
	Username string       `json:"username"`
	Kind     UserKindEnum `json:"kind"`
	FullName string       `json:"fullName"`
}

type Affordability string

const (
	AffordabilityCheap      Affordability = "CHEAP"
	AffordabilityAffordable Affordability = "AFFORDABLE"
	AffordabilityExpensive  Affordability = "EXPENSIVE"
)

var AllAffordability = []Affordability{
	AffordabilityCheap,
	AffordabilityAffordable,
	AffordabilityExpensive,
}

func (e Affordability) IsValid() bool {
	switch e {
	case AffordabilityCheap, AffordabilityAffordable, AffordabilityExpensive:
		return true
	}
	return false
}

func (e Affordability) String() string {
	return string(e)
}

func (e *Affordability) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Affordability(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Affordability", str)
	}
	return nil
}

func (e Affordability) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CuisineType string

const (
	CuisineTypeAmerican CuisineType = "AMERICAN"
	CuisineTypeAsian    CuisineType = "ASIAN"
	CuisineTypeEuropean CuisineType = "EUROPEAN"
)

var AllCuisineType = []CuisineType{
	CuisineTypeAmerican,
	CuisineTypeAsian,
	CuisineTypeEuropean,
}

func (e CuisineType) IsValid() bool {
	switch e {
	case CuisineTypeAmerican, CuisineTypeAsian, CuisineTypeEuropean:
		return true
	}
	return false
}

func (e CuisineType) String() string {
	return string(e)
}

func (e *CuisineType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CuisineType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CuisineType", str)
	}
	return nil
}

func (e CuisineType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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
