package storehandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/auth"
	storedb "github.com/relipocere/cafebackend/internal/database/store"
	"github.com/relipocere/cafebackend/internal/model"
)

// DeleteStoreRequest id of the store to delete.
type DeleteStoreRequest struct {
	StoreID int64
}

// DeleteStore deletes the store.
func (h *Handler) DeleteStore(ctx context.Context, req DeleteStoreRequest) error {
	user, ok := ctx.Value(auth.User).(model.User)
	if !ok {
		return fmt.Errorf("no user in the context")
	}

	err := validateDeleteStoreRequest(req)
	if err != nil {
		return fmt.Errorf("request validation: %w", err)
	}

	stores, err := h.storeRepo.Get(ctx, h.db, storedb.GetByIDs([]int64{req.StoreID}))
	if err != nil {
		return fmt.Errorf("get store %d: %w", req.StoreID, err)
	}

	if len(stores) == 0 {
		return model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: fmt.Sprintf("Store with id %d doesn't exist", req.StoreID),
		}
	}
	store := stores[0]

	if store.OwnerUsername != user.Username {
		return model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: "Only owner of the store can delete it",
		}
	}

	err = h.storeRepo.Delete(ctx, h.db, []int64{store.ID})
	if err != nil {
		return fmt.Errorf("delete store %d: %w", req.StoreID, err)
	}

	return nil
}

func validateDeleteStoreRequest(req DeleteStoreRequest) error {
	if req.StoreID == 0 {
		return model.Error{
			Code:    model.ErrorCodeBadRequest,
			Message: model.ErrMessageMissingFieldRequired("Store ID"),
		}
	}

	if req.StoreID < 0 {
		return model.Error{
			Code:    model.ErrorCodeBadRequest,
			Message: model.ErrMessageInvalidID("Store ID", req.StoreID),
		}
	}

	return nil
}
