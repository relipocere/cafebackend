package product

import (
	"context"
	"fmt"

	producthandler "github.com/relipocere/cafebackend/internal/business/product-handler"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/graph/mapping"
	"github.com/relipocere/cafebackend/internal/model"
)

type productHandler interface {
	CreateProdcut(ctx context.Context, req producthandler.CreateProductRequest) (model.Product, error)
	DeleteProduct(ctx context.Context, productID int64) error
}

type App struct {
	productHandler productHandler
}

func NewApp(productHandler productHandler) *App {
	return &App{
		productHandler: productHandler,
	}
}

func (a *App) CreateProduct(ctx context.Context, input graphmodel.CreateProductInput) (graphmodel.Product, error) {
	product, err := a.productHandler.CreateProdcut(ctx, producthandler.CreateProductRequest{
		Name:        input.Name,
		StoreID:     input.StoreID,
		PriceCents:  input.PriceCents,
		Ingerdients: input.Ingredients,
		Calories:    input.Calories,
		ImageID:     input.ImageID,
	})
	if err != nil {
		return graphmodel.Product{}, fmt.Errorf("business handler: %w", err)
	}

	return mapping.MapProduct(product), nil
}

func (a *App) DeleteProduct(ctx context.Context, productID int64) (bool, error) {
	err := a.productHandler.DeleteProduct(ctx, productID)
	if err != nil {
		return false, fmt.Errorf("business handler: %w", err)
	}

	return true, nil
}
