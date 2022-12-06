package producthandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/model"
)

func (h *Handler) GetProdcuts(ctx context.Context, productIDs []int64) ([]model.Product, error) {
	prodcuts, err := h.productRepo.Get(ctx, h.db, productIDs)
	if err != nil {
		return nil, fmt.Errorf("getting products with ids %+v: %w", productIDs, err)
	}

	return prodcuts, nil
}
