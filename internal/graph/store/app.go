package store

import (
	"context"

	storehandler "github.com/relipocere/cafebackend/internal/business/store-handler"
	"github.com/relipocere/cafebackend/internal/model"
)

type storeHandler interface {
	CreateStore(ctx context.Context, req storehandler.CreateStoreRequest) (storehandler.CreateStoreResponse, error)
	DeleteStore(ctx context.Context, req storehandler.DeleteStoreRequest) error
	SearchStores(ctx context.Context, req storehandler.SearchStoresRequest) ([]model.Store, error)
}

// App is a store related handler's app.
type App struct {
	storeHandler storeHandler
}

// NewApp creates new App.
func NewApp(storeHandler storeHandler) *App {
	return &App{storeHandler: storeHandler}
}
