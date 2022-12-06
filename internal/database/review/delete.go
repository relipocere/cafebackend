package review

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
)

// Delete removes a review from the review table.
func (*Repo) Delete(ctx context.Context, q database.Queryable, ids []int64) error {
	qb := database.PSQL.
		Delete(database.TableStoreReview).
		Where(sq.Eq{"id": ids})

	_, err := q.Exec(ctx, qb)
	if err != nil {
		return fmt.Errorf("delete review: %w", err)
	}

	return nil
}
