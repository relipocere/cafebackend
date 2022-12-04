package product

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

func (*Repo) Search(ctx context.Context, q database.Queryable, page model.Pagination, filter model.ProductFilter) ([]model.Product, error) {
	qb := baseSelectQuery.
		Offset(page.Offset()).
		Limit(page.Limit())

	qb = applySearchFilter(qb, filter)

	var dtos []productDTO
	err := q.Select(ctx, &dtos, qb)
	if err != nil {
		return nil, fmt.Errorf("select products: %w", err)
	}

	return mapToProducts(dtos), nil
}

func applySearchFilter(qb sq.SelectBuilder, filter model.ProductFilter) sq.SelectBuilder {
	if len(filter.StoreIDs) != 0 {
		qb = qb.Where(sq.Eq{"store_id": filter.StoreIDs})
	}

	if filter.PriceCents != nil {
		qb = database.ApplyIntFilter(qb, "price_cents", *filter.PriceCents)
	}

	if filter.Calories != nil {
		qb = database.ApplyIntFilter(qb, "calories", *filter.Calories)
	}

	return qb
}
