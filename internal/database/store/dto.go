package store

import (
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/relipocere/cafebackend/internal/model"
)

type storeDto struct {
	ID            edgedb.UUID `edgedb:"id"`
	Title         string      `edgedb:"title"`
	Affordability string      `edgedb:"affordability"`
	CuisineType   string      `edgedb:"cuisine_type"`
	OwnerUsername string      `edgedb:"owner_username"`
	ImageID       string      `edgedb:"image_id"`
	CreatedAt     time.Time   `edgedb:"created_at"`
	UpdatedAt     time.Time   `edgedb:"updated_At"`
}

func mapToStore(dto storeDto) model.Store {
	return model.Store{
		ID: dto.ID.String(),
		StoreCreate: model.StoreCreate{
			Title:         dto.Title,
			Affordability: model.Affordability(dto.Affordability),
			Cuisine:       model.Cuisine(dto.CuisineType),
			OwnerUsername: dto.OwnerUsername,
			ImageID:       dto.ImageID,
			CreatedAt:     dto.CreatedAt,
			UpdatedAt:     dto.UpdatedAt,
		},
	}
}
