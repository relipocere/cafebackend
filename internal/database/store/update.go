package store

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
)

type SetFieldFn func(qb sq.UpdateBuilder) sq.UpdateBuilder

func SetAverageRating(rating float64) SetFieldFn {
	return func(qb sq.UpdateBuilder) sq.UpdateBuilder {
		return qb.Set("avg_rating", rating)
	}
}

func SetNumberOfReviews(numberOfReviews int64) SetFieldFn {
	return func(qb sq.UpdateBuilder) sq.UpdateBuilder {
		return qb.Set("number_of_reviews", numberOfReviews)
	}
}

func (*Repo) Update(ctx context.Context, q database.Queryable, id int64, now time.Time, setters ...SetFieldFn) error {
	if len(setters) == 0 {
		return nil
	}

	qb := database.PSQL.
		Update(database.TableStore).
		Set("updated_at", now).
		Where(sq.Eq{"id": id})

	for _, setter := range setters {
		qb = setter(qb)
	}

	_, err := q.Exec(ctx, qb)
	if err != nil {
		return fmt.Errorf("update store: %w", err)
	}

	return nil
}
