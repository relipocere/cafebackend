package producthandler

import (
	"context"
	"time"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

type productRepo interface {
	Create(ctx context.Context, q database.Queryable, product model.ProductCreate) (int64, error)
}

type storeRepo interface {
	Get(ctx context.Context, q database.Queryable, ids []int64) ([]model.Store, error)
}

// Handler handles product related scenarios.
type Handler struct {
	db          database.PGX
	productRepo productRepo
	storeRepo   storeRepo
	now         func() time.Time
}

func NewHandler(db database.PGX, productRepo productRepo, storeRepo storeRepo) *Handler {
	return &Handler{
		db:          db,
		productRepo: productRepo,
		storeRepo:   storeRepo,
		now:         func() time.Time { return time.Now().UTC() },
	}
}
