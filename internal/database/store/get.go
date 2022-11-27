package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/edgedb/edgedb-go"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Get gets the store by id.
func (r *Repo) Get(ctx context.Context, q database.Queryable, id string) (*model.Store, error) {
	query := `select Store{
		id,
		title,
		affordability,
		cuisine_type,
		owner_username := .owner.username,
		image_id, 
		created_at,
		updated_at
	} filter .id=<uuid>$0)`

	var dto storeDto
	var dbError edgedb.Error
	err := q.QuerySingle(ctx, query, &dto, id)
	if errors.As(err, &dbError) && dbError.Category(edgedb.NoDataError) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("query single: %w", err)
	}

	store := mapToStore(dto)
	return &store, nil
}
