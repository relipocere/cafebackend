package store

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Repo is the repository for working with users.
type Repo struct {
}

// NewRepo creates Repo.
func NewRepo() *Repo {
	return &Repo{}
}

// Create creates new store.
func (*Repo) Create(ctx context.Context, q database.Queryable, store model.StoreCreate) (int64, error) {
	qb := database.PSQL.
		Insert(database.TableStore).
		Columns(
			"title",
			"affordability",
			"cuisine",
			"owner_username",
			"image_id",
			"avg_rating",
			"number_of_reviews",
			"created_at",
			"updated_at",
		).
		Values(
			store.Title,
			store.Affordability,
			store.Cuisine,
			store.OwnerUsername,
			store.ImageID,
			store.AverageRating,
			store.NumberOfReviews,
			store.CreatedAt,
			store.UpdatedAt,
		).Suffix("returning id")

	var id int64
	err := q.Get(ctx, &id, qb)
	if err != nil {
		return 0, fmt.Errorf("insert store: %w", err)
	}

	return id, nil
}
