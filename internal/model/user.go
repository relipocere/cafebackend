package model

import (
	"time"
)

type UserKind string

// All possible kinds of users.
const (
	UserKindConsumer = "Consumer"
	UserKindBusiness = "Business"
)

var validUserKinds = map[UserKind]struct{}{
	UserKindConsumer: {},
	UserKindBusiness: {},
}

func (k UserKind) IsValid() bool {
	_, ok := validUserKinds[k]
	return ok
}

// User is the user-handler model.
type User struct {
	ID string
	UserCreate
}

// UserCreate is the struct for user-handler creation.
type UserCreate struct {
	Username     string
	FullName     string
	Kind         UserKind
	PasswordHash string
	Salt         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
