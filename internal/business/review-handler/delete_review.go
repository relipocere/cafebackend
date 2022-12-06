package reviewhandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/auth"
	"github.com/relipocere/cafebackend/internal/database"
	storedb "github.com/relipocere/cafebackend/internal/database/store"
	"github.com/relipocere/cafebackend/internal/model"
)

func (h *Handler) DeleteReview(ctx context.Context, reviewID int64) error {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer database.RollbackTx(ctx, tx, "CreateReview")

	now := h.now()

	user, ok := ctx.Value(auth.User).(model.User)
	if !ok {
		return fmt.Errorf("no user in the context")
	}

	reviews, err := h.reviewRepo.Get(ctx, tx, []int64{reviewID})
	if err != nil {
		return fmt.Errorf("getting review %d: %w", reviewID, err)
	}

	if len(reviews) == 0 {
		return model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: fmt.Sprintf("Review %d doesn't exist", reviewID),
		}
	}
	review := reviews[0]

	if review.AuthorUsername != user.Username {
		return model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: "You're not the author this review",
		}
	}

	stores, err := h.storeRepo.Get(ctx, tx, storedb.GetByIDs([]int64{review.StoreID}))
	if err != nil {
		return fmt.Errorf("getting store %d: %w", review.StoreID, err)
	}

	if len(stores) == 0 {
		return fmt.Errorf("store %d doesn't exist", review.StoreID)
	}
	store := stores[0]

	newAverage := subFromAverage(store.AverageRating, store.NumberOfReviews, review.Rating)

	err = h.storeRepo.Update(
		ctx,
		tx,
		review.StoreID,
		now,
		storedb.SetAverageRating(int64(newAverage)),
		storedb.SetNumberOfReviews(store.NumberOfReviews-1),
	)
	if err != nil {
		return fmt.Errorf("update store %d: %w", review.StoreID, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
