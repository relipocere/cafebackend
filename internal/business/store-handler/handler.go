package storehandler

import (
	"context"
	"time"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

type storeRepo interface {
	Create(ctx context.Context, q database.Queryable, store model.StoreCreate) (string, error)
}

// Handler handles user related scenarios.
type Handler struct {
	edge      database.Edge
	storeRepo storeRepo
	now       func() time.Time
}

// NewHandler creates new Handler.
func NewHandler(edge database.Edge, storeRepo storeRepo) *Handler {
	return &Handler{
		edge:      edge,
		storeRepo: storeRepo,
		now:       func() time.Time { return time.Now().UTC() },
	}
}
