package store

import (
	"time"

	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

var baseSelectQuery = database.PSQL.
	Select(
		"id",
		"title",
		"affordability",
		"cuisine",
		"owner_username",
		"image_id",
		"avg_rating",
		"number_of_reviews",
		"created_at",
		"updated_at",
	).
	From(database.TableStore)

type storeDto struct {
	ID              int64     `db:"id"`
	Title           string    `db:"title"`
	Affordability   string    `db:"affordability"`
	CuisineType     string    `db:"cuisine"`
	OwnerUsername   string    `db:"owner_username"`
	ImageID         string    `db:"image_id"`
	AverageRating   int64     `db:"avg_rating"`
	NumberOfReviews int64     `db:"number_of_reviews"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

func mapToStores(dtos []storeDto) []model.Store {
	stores := make([]model.Store, 0, len(dtos))

	for _, dto := range dtos {
		stores = append(stores, mapToStore(dto))
	}

	return stores
}

func mapToStore(dto storeDto) model.Store {
	return model.Store{
		ID: dto.ID,
		StoreCreate: model.StoreCreate{
			Title:           dto.Title,
			Affordability:   model.Affordability(dto.Affordability),
			Cuisine:         model.Cuisine(dto.CuisineType),
			OwnerUsername:   dto.OwnerUsername,
			ImageID:         dto.ImageID,
			AverageRating:   dto.AverageRating,
			NumberOfReviews: dto.NumberOfReviews,
			CreatedAt:       dto.CreatedAt,
			UpdatedAt:       dto.UpdatedAt,
		},
	}
}
