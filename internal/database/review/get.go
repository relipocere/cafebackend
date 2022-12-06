package review

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

func (*Repo) Get(ctx context.Context, q database.Queryable, ids []int64) ([]model.Review, error) {
	qb := baseSelectQuery.Where(squirrel.Eq{"id": ids})

	var dtos []reviewDto
	err := q.Select(ctx, &dtos, qb)
	if err != nil {
		return nil, fmt.Errorf("select reviews: %w", err)
	}

	return mapToReviews(dtos), nil
}
