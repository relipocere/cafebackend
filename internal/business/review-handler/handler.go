package reviewhandler

import (
	"context"
	"time"

	"github.com/relipocere/cafebackend/internal/database"
	storedb "github.com/relipocere/cafebackend/internal/database/store"
	"github.com/relipocere/cafebackend/internal/model"
)

type storeRepo interface {
	Get(ctx context.Context, q database.Queryable, predicateFn storedb.GetPredicateFn) ([]model.Store, error)
	Update(ctx context.Context, q database.Queryable, id int64, now time.Time, setters ...storedb.SetFieldFn) error
}

type reviewRepo interface {
	Search(ctx context.Context, q database.Queryable, page model.Pagination, filter model.ReviewFilter) ([]model.Review, error)
	Delete(ctx context.Context, q database.Queryable, ids []int64) error
	Create(ctx context.Context, q database.Queryable, review model.ReviewToCreate) (int64, error)
	Get(ctx context.Context, q database.Queryable, ids []int64) ([]model.Review, error)
}

// Handler handles review related scenarios.
type Handler struct {
	db         database.PGX
	storeRepo  storeRepo
	reviewRepo reviewRepo
	now        func() time.Time
}

func NewHandler(db database.PGX, storeRepo storeRepo, reviewRepo reviewRepo) *Handler {
	return &Handler{
		db:         db,
		storeRepo:  storeRepo,
		reviewRepo: reviewRepo,
		now:        func() time.Time { return time.Now().UTC() },
	}
}

func addToAverage(average float64, size int64, value int64) float64 {
	n := float64(size)
	v := float64(value)

	return average + (v)/(n+1)
}

func subFromAverage(average float64, size int64, value int64) float64 {
	if size == 1 {
		return 0
	}

	n := float64(size)
	v := float64(value)

	return (n*average - v) / (n - 1)
}
