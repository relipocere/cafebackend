package store

import (
	"context"
	"fmt"

	storehandler "github.com/relipocere/cafebackend/internal/business/store-handler"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/graph/mapping"
)

// CreateStore creates new store.
func (a *App) CreateStore(ctx context.Context, input graphmodel.CreateStoreInput) (graphmodel.Store, error) {
	request := mapCreateStoreRequest(input)

	response, err := a.storeHandler.CreateStore(ctx, request)
	if err != nil {
		return graphmodel.Store{}, fmt.Errorf("business handler: %w", err)
	}

	return mapping.MapStore(response.Store), nil
}

func mapCreateStoreRequest(input graphmodel.CreateStoreInput) storehandler.CreateStoreRequest {
	return storehandler.CreateStoreRequest{
		Title:         input.Title,
		Affordability: mapping.MapToAffordability(input.Affordability),
		Cuisine:       mapping.MapToCuisine(input.CuisineType),
		ImageID:       input.ImageID,
	}
}
