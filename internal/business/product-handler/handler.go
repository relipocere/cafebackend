package producthandler

import (
	"context"
	"time"

	"github.com/relipocere/cafebackend/internal/database"
	storedb "github.com/relipocere/cafebackend/internal/database/store"
	"github.com/relipocere/cafebackend/internal/model"
)

type productRepo interface {
	Create(ctx context.Context, q database.Queryable, product model.ProductCreate) (int64, error)
	Get(ctx context.Context, q database.Queryable, productIDs []int64) ([]model.Product, error)
	Delete(ctx context.Context, q database.Queryable, productIDs []int64) error
	Search(ctx context.Context, q database.Queryable, page model.Pagination, filter model.ProductFilter) ([]model.Product, error)
}

type storeRepo interface {
	Get(ctx context.Context, q database.Queryable, predicateFn storedb.GetPredicateFn) ([]model.Store, error)
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
