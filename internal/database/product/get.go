package product

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

func (*Repo) Get(ctx context.Context, q database.Queryable, productIDs []int64) ([]model.Product, error) {
	if len(productIDs) == 0{
		return nil, nil
	}

	qb := baseSelectQuery.
		Where(squirrel.Eq{"id": productIDs})

	var dtos []productDTO
	err := q.Select(ctx, &dtos, qb)
	if err != nil {
		return nil, fmt.Errorf("select products: %w", err)
	}

	return mapToProducts(dtos), nil
}
