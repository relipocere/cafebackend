package model

import (
	"time"
)

// Session is user-handler authentication session.
type Session struct {
	SessionID string
	Username  string
	ExpiresAt time.Time
}
