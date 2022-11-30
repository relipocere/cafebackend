package store

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Search searches stores by the provided criteria.
func (*Repo) Search(ctx context.Context, q database.Queryable, page model.Pagination, filter model.StoreFilter) ([]model.Store, error) {
	qb := baseSelectQuery.
		Limit(page.Limit()).
		Offset(page.Offset())

	qb = applySearchFilter(qb, filter)

	var dtos []storeDto
	err := q.Select(ctx, &dtos, qb)

	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	stores := mapToStores(dtos)
	return stores, nil
}

func applySearchFilter(qb sq.SelectBuilder, filter model.StoreFilter) sq.SelectBuilder {
	if filter.TitleQuery != nil {
		predicate := "%" + database.SanitizeLikeQuery(*filter.TitleQuery) + "%s"
		qb = qb.Where(sq.ILike{"title": predicate})
	}

	if filter.AverageRating != nil {
		qb = database.ApplyIntFilter(qb, "avg_rating", *filter.AverageRating)
	}

	if len(filter.OwnerUsernames) != 0 {
		qb = qb.Where(sq.Eq{"owner_username": filter.OwnerUsernames})
	}

	if len(filter.Affordability) != 0 {
		var values []string
		for _, afforaffordability := range filter.Affordability {
			values = append(values, string(afforaffordability))
		}

		qb = qb.Where(sq.Eq{"affordability": values})
	}

	if len(filter.Cuisines) != 0 {
		var values []string
		for _, cuisine := range filter.Cuisines {
			values = append(values, string(cuisine))
		}

		qb = qb.Where(sq.Eq{"cuisine": values})
	}

	return qb
}
