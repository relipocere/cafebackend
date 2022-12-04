package product

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

func (*Repo) Get(ctx context.Context, q database.Queryable, productIDs []int64) ([]model.Product, error) {
	qb := baseSelectQuery.
		Where(squirrel.Eq{"product_id": productIDs})

	var dtos []productDTO
	err := q.Select(ctx, &dtos, qb)
	if err != nil {
		return nil, fmt.Errorf("select products: %w", err)
	}

	return mapToProducts(dtos), nil
}
