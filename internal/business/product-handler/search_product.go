package producthandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/model"
)

type SearchProductsRequest struct {
	Page       model.Pagination
	StoreIDs   []int64
	PriceCents *model.IntRange
	Calories   *model.IntRange
}

func (h *Handler) SearchProducts(ctx context.Context, req SearchProductsRequest) ([]model.Product, error) {
	products, err := h.productRepo.Search(ctx, h.db, req.Page, model.ProductFilter{
		StoreIDs:   req.StoreIDs,
		PriceCents: req.PriceCents,
		Calories:   req.Calories,
	})
	if err != nil {
		return nil, fmt.Errorf("searching products: %w", err)
	}

	return products, nil
}
