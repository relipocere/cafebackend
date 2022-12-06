package reviewhandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/auth"
	"github.com/relipocere/cafebackend/internal/database"
	storedb "github.com/relipocere/cafebackend/internal/database/store"
	"github.com/relipocere/cafebackend/internal/model"
)

type CreateReviewRequest struct {
	StoreID    int64
	Rating     int64
	Commentary string
}

func (h *Handler) CreateReview(ctx context.Context, req CreateReviewRequest) (model.Review, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return model.Review{}, fmt.Errorf("begin tx: %w", err)
	}
	defer database.RollbackTx(ctx, tx, "CreateReview")

	err = validateCreateReviewRequest(req)
	if err != nil {
		return model.Review{}, fmt.Errorf("request validation: %w", err)
	}

	now := h.now()

	user, ok := ctx.Value(auth.User).(model.User)
	if !ok {
		return model.Review{}, fmt.Errorf("no user in the context")
	}

	stores, err := h.storeRepo.Get(ctx, tx, storedb.GetByIDs([]int64{req.StoreID}))
	if err != nil {
		return model.Review{}, fmt.Errorf("getting store %d: %w", req.StoreID, err)
	}

	if len(stores) == 0 {
		return model.Review{}, model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: fmt.Sprintf("Store %d doesn't exist", req.StoreID),
		}
	}
	store := stores[0]

	reviewToCreate := model.ReviewToCreate{
		StoreID:        req.StoreID,
		Rating:         req.Rating,
		AuthorUsername: user.Username,
		Commentary:     req.Commentary,
	}

	reviewID, err := h.reviewRepo.Create(ctx, tx, reviewToCreate)
	if err != nil {
		return model.Review{}, fmt.Errorf("create review: %w", err)
	}

	newAverage := addToAverage(store.AverageRating, store.NumberOfReviews, req.Rating)

	err = h.storeRepo.Update(
		ctx,
		tx,
		req.StoreID,
		now,
		storedb.SetAverageRating(newAverage),
		storedb.SetNumberOfReviews(store.NumberOfReviews+1),
	)
	if err != nil {
		return model.Review{}, fmt.Errorf("update store %d: %w", req.StoreID, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return model.Review{}, fmt.Errorf("commit tx: %w", err)
	}

	return model.Review{
		ID:             reviewID,
		ReviewToCreate: reviewToCreate,
	}, nil
}

func validateCreateReviewRequest(req CreateReviewRequest) error {
	if req.Rating < 1 || req.Rating > 5 {
		return model.Error{
			Code:    model.ErrorCodeBadRequest,
			Message: "Rating must be a value between 1 and 5",
		}
	}

	return nil
}
