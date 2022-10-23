package user

import (
	"fmt"
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/relipocere/cafebackend/internal/model"
)

type userDto struct {
	ID           edgedb.UUID `edgedb:"id"`
	Username     string      `edgedb:"username"`
	FullName     string      `edgedb:"full_name"`
	Kind         string      `edgedb:"kind"`
	PasswordHash string      `edgedb:"password_hash"`
	Salt         string      `edgedb:"salt"`
	CreatedAt    time.Time   `edgedb:"created_at"`
	UpdatedAt    time.Time   `edgedb:"updated_at"`
	Session      sessionDto  `edgedb:"session"`
}

type sessionDto struct {
	edgedb.Optional
	SessionID string    `edgedb:"session_id"`
	ExpiresAt time.Time `edgedb:"expires_at"`
}

func mapToUser(d userDto) (model.User, error) {
	userKind := model.UserKind(d.Kind)
	if !userKind.IsValid() {
		return model.User{}, fmt.Errorf("invalid user-handler kind %s", d.Kind)
	}

	return model.User{
		ID: d.ID.String(),
		UserCreate: model.UserCreate{
			Username:     d.Username,
			FullName:     d.FullName,
			Kind:         userKind,
			PasswordHash: d.PasswordHash,
			Salt:         d.Salt,
			CreatedAt:    d.CreatedAt,
			UpdatedAt:    d.UpdatedAt,
		},
	}, nil
}
