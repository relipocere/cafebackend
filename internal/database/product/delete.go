package product

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
)

// Delete deletes record of a product.
func (*Repo) Delete(ctx context.Context, q database.Queryable, productIDs []int64) error {
	qb := database.PSQL.Delete(database.TableProduct).
		Where(squirrel.Eq{"id": productIDs})

	_, err := q.Exec(ctx, qb)
	if err != nil {
		return fmt.Errorf("delete product: %w", err)
	}

	return nil
}
