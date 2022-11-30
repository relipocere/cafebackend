package store

import (
	"context"
	"fmt"

	storehandler "github.com/relipocere/cafebackend/internal/business/store-handler"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
)

// DeleteStore implements graphql DeleteStore method.
func (a *App) DeleteStore(ctx context.Context, input graphmodel.DeleteStoreInput) (bool, error) {
	err := a.storeHandler.DeleteStore(ctx, storehandler.DeleteStoreRequest{
		StoreID: input.ID,
	})
	if err != nil {
		return false, fmt.Errorf("business handler: %w", err)
	}

	return true, nil
}
