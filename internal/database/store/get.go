package store

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Get gets the store by id.
func (r *Repo) Get(ctx context.Context, q database.Queryable, ids []int64) ([]model.Store, error) {
	qb := baseSelectQuery.Where(sq.Eq{"id": ids})

	var dtos []storeDto
	err := q.Select(ctx, &dtos, qb)

	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	stores := mapToStores(dtos)
	return stores, nil
}
