package review

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

func (*Repo) Search(ctx context.Context, q database.Queryable, page model.Pagination, filter model.ReviewFilter) ([]model.Review, error) {
	qb := baseSelectQuery.
		Limit(page.Limit()).
		Offset(page.Offset())

	qb = applySearchFilter(qb, filter)

	var dtos []reviewDto
	err := q.Select(ctx, &dtos, qb)
	if err != nil {
		return nil, fmt.Errorf("select reviews: %w", err)
	}

	return mapToReviews(dtos), nil
}

func applySearchFilter(qb sq.SelectBuilder, filter model.ReviewFilter) sq.SelectBuilder {
	if len(filter.StoreIDs) > 0 {
		qb = qb.Where(sq.Eq{"store_id": filter.StoreIDs})
	}

	if len(filter.AuthorUsernames) > 0 {
		qb = qb.Where(sq.Eq{"author_username": filter.AuthorUsernames})
	}

	if filter.Rating != nil {
		qb = database.ApplyIntFilter(qb, "rating", *filter.Rating)
	}

	return qb
}
