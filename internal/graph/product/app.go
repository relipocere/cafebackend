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
	SearchProducts(ctx context.Context, req producthandler.SearchProductsRequest) ([]model.Product, error)
	GetProdcuts(ctx context.Context, productIDs []int64) ([]model.Product, error)
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

func (a *App) SearchProducts(ctx context.Context, input graphmodel.SearchProductsInput) ([]graphmodel.Product, error) {
	page := mapping.MapToPagination(input.Page)

	products, err := a.productHandler.SearchProducts(ctx, producthandler.SearchProductsRequest{
		Page:       page,
		StoreIDs:   input.StoreIDs,
		PriceCents: mapping.MapToIntRange(input.PriceCents),
		Calories:   mapping.MapToIntRange(input.Calories),
	})
	if err != nil {
		return nil, fmt.Errorf("business handler: %w", err)
	}

	return mapping.MapProducts(products), nil
}

func (a *App) GetProducts(ctx context.Context, productIDs []int64) ([]graphmodel.Product, error) {
	products, err := a.productHandler.GetProdcuts(ctx, productIDs)
	if err != nil {
		return nil, fmt.Errorf("business handler: %w", err)
	}

	return mapping.MapProducts(products), nil
}
