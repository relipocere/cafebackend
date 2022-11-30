package storehandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/business/validation"
	"github.com/relipocere/cafebackend/internal/model"
)

// SearchStoresRequest filters to search stores with.
type SearchStoresRequest struct {
	Page   model.Pagination
	Filter model.StoreFilter
}

// SearchStores performs stores search.
func (h *Handler) SearchStores(ctx context.Context, req SearchStoresRequest) ([]model.Store, error) {
	err := validateSearchStoresRequest(req)
	if err != nil {
		return nil, fmt.Errorf("validation: %w", err)
	}

	stores, err := h.storeRepo.Search(ctx, h.db, req.Page, req.Filter)
	if err != nil {
		return nil, fmt.Errorf("stores search: %w", err)
	}

	return stores, nil
}

func validateSearchStoresRequest(req SearchStoresRequest) error {
	err := validation.ValidatePagination(req.Page)
	if err != nil {
		return err
	}

	err = validation.ValidateUsernames(req.Filter.OwnerUsernames)
	if err != nil {
		return err
	}

	err = validation.ValidateCuisines(req.Filter.Cuisines)
	if err != nil {
		return err
	}

	err = validation.ValidateSliceOfAffordability(req.Filter.Affordability)
	if err != nil {
		return err
	}

	return nil
}
