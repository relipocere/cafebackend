package reviewhandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/business/validation"
	"github.com/relipocere/cafebackend/internal/model"
)

type SearchReviewsRequest struct {
	Page   model.Pagination
	Filter model.ReviewFilter
}

func (h *Handler) SearchReviews(ctx context.Context, req SearchReviewsRequest) ([]model.Review, error) {
	err := validateSearchReviewsRequest(req)
	if err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	reviews, err := h.reviewRepo.Search(ctx, h.db, req.Page, req.Filter)
	if err != nil {
		return nil, fmt.Errorf("search of reviews: %w", err)
	}

	return reviews, nil
}

func validateSearchReviewsRequest(req SearchReviewsRequest) error {
	err := validation.ValidatePagination(req.Page)
	if err != nil {
		return err
	}

	err = validation.ValidateUsernames(req.Filter.AuthorUsernames)
	if err != nil {
		return err
	}

	return nil
}
