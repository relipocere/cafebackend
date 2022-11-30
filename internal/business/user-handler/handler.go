package userhandler

import (
	"context"
	"time"

	"github.com/relipocere/cafebackend/internal/database"
	userdb "github.com/relipocere/cafebackend/internal/database/user"
	"github.com/relipocere/cafebackend/internal/model"
)

type userRepo interface {
	Create(ctx context.Context, q database.Queryable, user model.User) (string, error)
	Get(ctx context.Context, q database.Queryable, filter userdb.GetFilter) (*model.User, error)
	SetSession(ctx context.Context, q database.Queryable, username string, session model.Session) error
}

// Handler handles user related scenarios
type Handler struct {
	db       database.PGX
	userRepo userRepo
	now      func() time.Time
}

// NewHandler creates Handler.
func NewHandler(db database.PGX, userRepo userRepo) *Handler {
	return &Handler{
		db:       db,
		userRepo: userRepo,
		now:      func() time.Time { return time.Now().UTC() },
	}
}
