package store

import (
	"context"
	"fmt"

	storehandler "github.com/relipocere/cafebackend/internal/business/store-handler"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/graph/mapping"
	"github.com/relipocere/cafebackend/internal/model"
)

func (a *App) SearchStores(ctx context.Context, input graphmodel.SearchStoresInput) ([]graphmodel.Store, error) {
	req := mapSearchStoresRequest(input)

	stores, err := a.storeHandler.SearchStores(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("business handler: %w", err)
	}

	return mapping.MapStores(stores), nil
}

func mapSearchStoresRequest(input graphmodel.SearchStoresInput) storehandler.SearchStoresRequest {
	page := mapping.MapToPagination(input.Page)

	filter := model.StoreFilter{
		TitleQuery:     input.TitleQuery,
		AverageRating:  mapping.MapToIntRange(input.Rating),
		OwnerUsernames: input.OwnerUsernames,
		Affordability:  mapping.MapToAffordabilitySlice(input.Affordability),
		Cuisines:       mapping.MapToCuisines(input.Cuisines),
	}

	return storehandler.SearchStoresRequest{
		Page:   page,
		Filter: filter,
	}
}
