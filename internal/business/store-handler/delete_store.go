package storehandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/model"
)

// DeleteStoreRequest id of the store to delete.
type DeleteStoreRequest struct {
	StoreID string
}

// DeleteStore deletes the store.
func (h *Handler) DeleteStore(ctx context.Context, req DeleteStoreRequest) error {
	user, ok := ctx.Value("user").(model.User)
	if !ok {
		return fmt.Errorf("no user in the context")
	}

	err := validateDeleteStoreRequest(req)
	if err != nil {
		return fmt.Errorf("request validation: %w", err)
	}

	store, err := h.storeRepo.Get(ctx, h.edge, req.StoreID)
	if err != nil {
		return fmt.Errorf("get store %s: %w", req.StoreID, err)
	}

	if store == nil {
		return model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: fmt.Sprintf("Store with id %s doesn't exist", req.StoreID),
		}
	}

	if store.OwnerUsername != user.Username {
		return model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: "You're not the owner of the store",
		}
	}

	// no tx here
	err = h.storeRepo.Delete(ctx, h.edge, []string{req.StoreID})
	if err != nil {
		return fmt.Errorf("delete store %s: %w", req.StoreID, err)
	}

	return nil
}

func validateDeleteStoreRequest(req DeleteStoreRequest) error {
	if req.StoreID == "" {
		return model.Error{
			Code:    model.ErrorCodeBadRequest,
			Message: model.ErrMessageMissingFieldRequired("Store id"),
		}
	}

	return nil
}
