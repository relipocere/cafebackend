package review

import (
	"context"
	"fmt"

	reviewhandler "github.com/relipocere/cafebackend/internal/business/review-handler"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/graph/mapping"
	"github.com/relipocere/cafebackend/internal/model"
)

type reviewHandler interface {
	CreateReview(ctx context.Context, req reviewhandler.CreateReviewRequest) (model.Review, error)
	DeleteReview(ctx context.Context, reviewID int64) error
	SearchReviews(ctx context.Context, req reviewhandler.SearchReviewsRequest) ([]model.Review, error)
}

type App struct {
	reviewHandler reviewHandler
}

func NewApp(reviewHandler reviewHandler) *App {
	return &App{
		reviewHandler: reviewHandler,
	}
}

func (a *App) CreateReview(ctx context.Context, input graphmodel.CreateReviewInput) (graphmodel.Review, error) {
	review, err := a.reviewHandler.CreateReview(ctx, reviewhandler.CreateReviewRequest{
		StoreID:    input.StoreID,
		Rating:     input.Rating,
		Commentary: input.Commentary,
	})
	if err != nil {
		return graphmodel.Review{}, fmt.Errorf("business handler: %w", err)
	}

	return mapping.MapReview(review), nil
}

func (a *App) DeleteReview(ctx context.Context, reviewID int64) (bool, error) {
	err := a.reviewHandler.DeleteReview(ctx, reviewID)
	if err != nil {
		return false, fmt.Errorf("business handler: %w", err)
	}

	return true, nil
}

func (a *App) SearchReviews(ctx context.Context, input graphmodel.SearchReviewsInput) ([]graphmodel.Review, error) {
	page := mapping.MapToPagination(input.Page)

	reviews, err := a.reviewHandler.SearchReviews(ctx, reviewhandler.SearchReviewsRequest{
		Page: page,
		Filter: model.ReviewFilter{
			StoreIDs:        input.StoreIDs,
			AuthorUsernames: input.AuthorUsernames,
			Rating:          mapping.MapToIntRange(input.Rating),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("business handler: %w", err)
	}

	return mapping.MapReviews(reviews), nil
}
