package product

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

// Creates inserts new product into db.
func (*Repo) Create(ctx context.Context, q database.Queryable, product model.ProductCreate) (int64, error) {
	qb := database.PSQL.
		Insert(database.TableProduct).
		Columns(
			"name",
			"store_id",
			"ingredients",
			"calories",
			"image_id",
			"created_at",
			"updated_at",
		).
		Values(
			product.Name,
			product.StoreID,
			database.PreventNullSlice(product.Ingredients),
			product.Calories,
			product.ImageID,
			product.CreatedAt,
			product.UpdatedAt,
		).Suffix("returning id")

	var id int64
	err := q.Get(ctx, &id, qb)
	if err != nil {
		return 0, fmt.Errorf("insert product: %w", err)
	}

	return id, nil
}
