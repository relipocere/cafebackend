package store

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/database"
)

// Delete deletes stores by ids.
func (r *Repo) Delete(ctx context.Context, q database.Queryable, ids []string) error {
	query := `delete Store filter .id=array<uuid>$0`

	err := q.Execute(ctx, query, ids)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}
