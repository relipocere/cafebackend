package store

import (
	"context"
	"fmt"

	storehandler "github.com/relipocere/cafebackend/internal/business/store-handler"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/graph/mapping"
)

type storeHandler interface {
	CreateStore(ctx context.Context, req storehandler.CreateStoreRequest) (storehandler.CreateStoreResponse, error)
}

// App is a store related handler's app.
type App struct {
	storeHandler storeHandler
}

// NewApp creates new App.
func NewApp(storeHandler storeHandler) *App {
	return &App{storeHandler: storeHandler}
}

// CreateStore creates new store.
func (a *App) CreateStore(ctx context.Context, input graphmodel.CreateStoreInput) (graphmodel.Store, error) {
	request := mapCreateStoreRequest(input)

	response, err := a.storeHandler.CreateStore(ctx, request)
	if err != nil {
		return graphmodel.Store{}, fmt.Errorf("business handler: %w", err)
	}

	return mapCreateStoreResponse(response), nil
}

func mapCreateStoreRequest(input graphmodel.CreateStoreInput) storehandler.CreateStoreRequest {
	return storehandler.CreateStoreRequest{
		Title:         input.Title,
		Affordability: mapping.MapToAffordability(input.Affordability),
		Cuisine:       mapping.MapToCuisine(input.CuisineType),
		ImageID:       input.ImageID,
	}
}

func mapCreateStoreResponse(response storehandler.CreateStoreResponse) graphmodel.Store {
	store := response.Store

	return graphmodel.Store{
		ID:            store.ID,
		Title:         store.Title,
		Affordability: mapping.MapAffordability(store.Affordability),
		CuisineType:   mapping.MapCuisine(store.Cuisine),
		ImageID:       store.ImageID,
		CreatedAt:     store.CreatedAt,
		UpdatedAt:     store.UpdatedAt,
	}
}
