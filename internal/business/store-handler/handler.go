package storehandler

import (
	"context"
	"time"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

type storeRepo interface {
	Create(ctx context.Context, q database.Queryable, store model.StoreCreate) (int64, error)
	Get(ctx context.Context, q database.Queryable, ids []int64) ([]model.Store, error)
	Search(ctx context.Context, q database.Queryable, page model.Pagination, filter model.StoreFilter) ([]model.Store, error)
	Delete(ctx context.Context, q database.Queryable, ids []int64) error
}

// Handler handles user related scenarios.
type Handler struct {
	db        database.PGX
	storeRepo storeRepo
	now       func() time.Time
}

// NewHandler creates new Handler.
func NewHandler(db database.PGX, storeRepo storeRepo) *Handler {
	return &Handler{
		db:        db,
		storeRepo: storeRepo,
		now:       func() time.Time { return time.Now().UTC() },
	}
}
