package user

import (
	"fmt"
	"time"

	"github.com/relipocere/cafebackend/internal/model"
)

type userDto struct {
	ID               int64      `db:"id"`
	Username         string     `db:"username"`
	FullName         string     `db:"full_name"`
	Kind             string     `db:"kind"`
	PasswordHash     string     `db:"password_hash"`
	Salt             string     `db:"salt"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at"`
	SessionID        *string    `db:"session_id"`
	SessionExpiresAt *time.Time `db:"session_expires_at"`
}

func mapToUser(d userDto) (model.User, error) {
	userKind := model.UserKind(d.Kind)
	if !userKind.IsValid() {
		return model.User{}, fmt.Errorf("invalid user kind %s", d.Kind)
	}

	var session *model.Session
	if d.SessionID != nil && d.SessionExpiresAt != nil{
		session = &model.Session{
			ID: *d.SessionID,
			ExpiresAt: *d.SessionExpiresAt,
		}
	}
	return model.User{
		Username:         d.Username,
		FullName:         d.FullName,
		Kind:             userKind,
		PasswordHash:     d.PasswordHash,
		Salt:             d.Salt,
		CreatedAt:        d.CreatedAt,
		UpdatedAt:        d.UpdatedAt,
		Session: session,
	}, nil
}
