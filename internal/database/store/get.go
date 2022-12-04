package store

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// GetPredicateFn is function that sets filter when getting store.
type GetPredicateFn func(sq sq.SelectBuilder) sq.SelectBuilder

func GetByIDs(ids []int64) GetPredicateFn {
	return func(qb sq.SelectBuilder) sq.SelectBuilder {
		return qb.Where(sq.Eq{"id": ids})
	}
}

func GetByProductIDs(productIDs []int64) GetPredicateFn {
	return func(qb sq.SelectBuilder) sq.SelectBuilder {
		return qb.Join(database.TableProduct + " p on p.store_id=s.id").
			Where(sq.Eq{"p.id": productIDs})
	}
}

// Get gets the store by id.
func (r *Repo) Get(ctx context.Context, q database.Queryable, predicateFn GetPredicateFn) ([]model.Store, error) {
	qb := predicateFn(baseSelectQuery)

	var dtos []storeDto
	err := q.Select(ctx, &dtos, qb)

	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	stores := mapToStores(dtos)
	return stores, nil
}
