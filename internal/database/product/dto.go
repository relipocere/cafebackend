package product

import (
	"time"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

var baseSelectQuery = database.PSQL.
	Select(
		"id",
		"name",
		"store_id",
		"price_cents",
		"ingredients",
		"calories",
		"image_id",
		"created_at",
		"updated_at",
	).
	From(database.TableProduct)

type productDTO struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	StoreID     int64     `db:"store_id"`
	PriceCents  int64     `db:"price_cents"`
	Ingerdients []string  `db:"ingredients"`
	Calories    int64     `db:"calories"`
	ImageID     string    `db:"image_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func mapToProducts(dtos []productDTO) []model.Product {
	products := make([]model.Product, 0, len(dtos))

	for _, dto := range dtos {
		products = append(products, model.Product{
			ID: dto.ID,
			ProductCreate: model.ProductCreate{
				Name:        dto.Name,
				StoreID:     dto.StoreID,
				PriceCents:  dto.PriceCents,
				Ingredients: dto.Ingerdients,
				Calories:    dto.Calories,
				ImageID:     dto.ImageID,
				CreatedAt:   dto.CreatedAt,
				UpdatedAt:   dto.UpdatedAt,
			},
		})
	}

	return products
}
