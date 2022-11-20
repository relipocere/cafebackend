package store

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Repo is the repository for working with users.
type Repo struct {
}

// NewRepo creates Repo.
func NewRepo() *Repo {
	return &Repo{}
}

// Create creates new store.
func (*Repo) Create(ctx context.Context, q database.Queryable, store model.StoreCreate) (string, error) {
	query := `insert Store{
			title := <str>$0,
			affordability := <Affordability>$1,
			cuisine_type := <Cuisine>$2,
			owner := (select User filter .username = <str>$3),
			image_id := <str>$4, 
			created_at := <datetime>$5,
			updated_at := <datetime>$6
		}`

	var dto storeDto
	err := q.QuerySingle(ctx, query, &dto,
		store.Title,
		string(store.Affordability),
		string(store.Cuisine),
		store.OwnerUsername,
		store.ImageID,
		store.CreatedAt,
		store.UpdatedAt,
	)
	if err != nil {
		return "", fmt.Errorf("insert: %w", err)
	}

	return dto.ID.String(), nil
}
