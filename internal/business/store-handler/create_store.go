package storehandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/model"
)

// CreateStoreRequest is the request for CreateStore handler.
type CreateStoreRequest struct {
	Title         string
	Affordability model.Affordability
	Cuisine       model.Cuisine
	ImageID       string
}

// CreateStoreResponse is the response of CreateStore handler.
type CreateStoreResponse struct {
	Store model.Store
}

// CreateStore handles store creation scenario.
func (h *Handler) CreateStore(ctx context.Context, req CreateStoreRequest) (CreateStoreResponse, error) {
	resp := CreateStoreResponse{}

	now := h.now()
	user, ok := ctx.Value("user").(model.User)
	if !ok {
		return resp, fmt.Errorf("no user in the context")
	}

	err := validateCreateStoreRequest(req, user)
	if err != nil {
		return resp, fmt.Errorf("validation: %w", err)
	}

	storeCreate := model.StoreCreate{
		Title:         req.Title,
		Affordability: req.Affordability,
		Cuisine:       req.Cuisine,
		OwnerUsername: user.Username,
		ImageID:       req.ImageID,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	storeID, err := h.storeRepo.Create(ctx, h.edge, storeCreate)
	if err != nil {
		return resp, fmt.Errorf("store creation: %w", err)
	}

	return CreateStoreResponse{
		Store: model.Store{
			ID:          storeID,
			StoreCreate: storeCreate,
		},
	}, nil
}

func validateCreateStoreRequest(req CreateStoreRequest, user model.User) error {
	if user.Kind != model.UserKindBusiness {
		return model.Error{
			Message: "User have account of type 'Business' to create a store",
			Code:    model.ErrorCodeFailedPrecondition,
		}
	}

	if !req.Affordability.IsValid() {
		return model.Error{
			Message: fmt.Sprintf("Invalid affordability: %s", string(req.Affordability)),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	if !req.Cuisine.IsValid() {
		return model.Error{
			Message: fmt.Sprintf("Invalid cuisine: %s", string(req.Cuisine)),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	return nil
}
