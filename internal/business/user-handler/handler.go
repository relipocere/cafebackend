package userhandler

import (
	"context"
	"time"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

type userRepo interface {
	Create(ctx context.Context, q database.Queryable, user model.UserCreate) (string, error)
	Get(ctx context.Context, q database.Queryable, username string) (*model.User, error)
	CreateSession(ctx context.Context, q database.Queryable, session model.Session) error
	GetSession(ctx context.Context, q database.Queryable, username string) (*model.Session, error)
}

// Handler handles user related scenarios
type Handler struct {
	edge     database.Edge
	userRepo userRepo
	now      func() time.Time
}

// NewHandler creates Handler.
func NewHandler(edge database.Edge, userRepo userRepo) *Handler {
	return &Handler{
		edge:     edge,
		userRepo: userRepo,
		now:      func() time.Time { return time.Now().UTC() },
	}
}
