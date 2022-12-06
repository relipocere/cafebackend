package review

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Creates inserts a review into the review table.
func (*Repo) Create(ctx context.Context, q database.Queryable, review model.ReviewToCreate) (int64, error) {
	qb := database.PSQL.
		Insert(database.TableStoreReview).
		Columns(
			"author_username",
			"store_id",
			"rating",
			"commentary",
		).
		Values(
			review.AuthorUsername,
			review.StoreID,
			review.Rating,
			review.Commentary,
		).Suffix("returning id")

	var id int64
	err := q.Get(ctx, &id, qb)
	if err != nil {
		return 0, fmt.Errorf("insert review: %w", err)
	}

	return id, nil
}
