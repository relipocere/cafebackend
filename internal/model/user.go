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

// User is struct representing user.
type User struct {
	Username     string
	FullName     string
	Kind         UserKind
	PasswordHash string
	Salt         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Session      *Session
}

// Session user authentication session.
type Session struct {
	ID        string
	ExpiresAt time.Time
}
