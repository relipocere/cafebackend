package review

import (
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

type Repo struct{}

func NewRepo() *Repo {
	return &Repo{}
}

var baseSelectQuery = database.PSQL.
	Select(
		"id",
		"author_username",
		"store_id",
		"rating",
		"commentary",
	).
	From(database.TableStoreReview)

type reviewDto struct {
	ID             int64  `db:"id"`
	AuthorUsername string `db:"author_username"`
	StoreID        int64  `db:"store_id"`
	Rating         int64  `db:"rating"`
	Commentary     string `db:"commentary"`
}

func mapToReviews(dtos []reviewDto) []model.Review {
	reviews := make([]model.Review, 0, len(dtos))

	for _, dto := range dtos {
		reviews = append(reviews, mapToReview(dto))
	}

	return reviews
}

func mapToReview(dto reviewDto) model.Review {
	return model.Review{
		ID: dto.ID,
		ReviewToCreate: model.ReviewToCreate{
			AuthorUsername: dto.AuthorUsername,
			StoreID:        dto.StoreID,
			Rating:         dto.Rating,
			Commentary:     dto.Commentary,
		},
	}
}
