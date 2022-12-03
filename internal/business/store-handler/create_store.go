package storehandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/auth"
	"github.com/relipocere/cafebackend/internal/business/validation"
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
	user, ok := ctx.Value(auth.User).(model.User)
	if !ok {
		return resp, fmt.Errorf("no user in the context")
	}

	err := validateCreateStoreRequest(req, user)
	if err != nil {
		return resp, fmt.Errorf("validation: %w", err)
	}

	storeCreate := model.StoreCreate{
		Title:           req.Title,
		Affordability:   req.Affordability,
		Cuisine:         req.Cuisine,
		OwnerUsername:   user.Username,
		ImageID:         req.ImageID,
		AverageRating:   0,
		NumberOfReviews: 0,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	storeID, err := h.storeRepo.Create(ctx, h.db, storeCreate)
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

	err := validation.ValidateName(req.Title, "Title")
	if err != nil {
		return err
	}

	err = validation.ValidateAffordability(req.Affordability)
	if err != nil {
		return err
	}

	err = validation.ValidateCuisine(req.Cuisine)
	if err != nil {
		return err
	}

	return nil
}
