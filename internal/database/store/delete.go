package store

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/relipocere/cafebackend/internal/database"
)

// Delete deletes stores by ids.
func (r *Repo) Delete(ctx context.Context, q database.Queryable, ids []int64) error {
	qb := database.PSQL.
		Delete(database.TableStore).
		Where(sq.Eq{"id": ids})

	_, err := q.Exec(ctx, qb)
	if err != nil {
		return fmt.Errorf("exec delete: %w", err)
	}

	return nil
}
