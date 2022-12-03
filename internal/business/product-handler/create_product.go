package producthandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/auth"
	"github.com/relipocere/cafebackend/internal/business/validation"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/model"
)

type CreateProductRequest struct {
	Name        string
	StoreID     int64
	Ingerdients []string
	Calories    int64
	ImageID     string
}

// CreateProduct handles scenario of product creation.
func (h *Handler) CreateProdcut(ctx context.Context, req CreateProductRequest) (model.Product, error) {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return model.Product{}, fmt.Errorf("begin tx: %w", err)
	}
	defer database.RollbackTx(ctx, tx, "CreateProduct")

	err = validateCreateProdcutRequest(req)
	if err != nil {
		return model.Product{}, fmt.Errorf("request validation: %w", err)
	}

	now := h.now()

	user, ok := ctx.Value(auth.User).(model.User)
	if !ok {
		return model.Product{}, fmt.Errorf("no user in the context")
	}

	stores, err := h.storeRepo.Get(ctx, tx, []int64{req.StoreID})
	if err != nil {
		return model.Product{}, fmt.Errorf("getting store %d: %w", req.StoreID, err)
	}

	if len(stores) == 0 {
		return model.Product{}, model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: fmt.Sprintf("Store %d doesn't exist", req.StoreID),
		}
	}

	store := stores[0]
	if store.OwnerUsername != user.Username {
		return model.Product{}, model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: "You're not the owner of the store",
		}
	}

	productToCreate := model.ProductCreate{
		Name:        req.Name,
		StoreID:     req.StoreID,
		Ingredients: req.Ingerdients,
		Calories:    req.Calories,
		ImageID:     req.ImageID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	productID, err := h.productRepo.Create(ctx, tx, productToCreate)
	if err != nil {
		return model.Product{}, fmt.Errorf("create product record: %w", err)
	}

	return model.Product{
		ID:            productID,
		ProductCreate: productToCreate,
	}, nil
}

func validateCreateProdcutRequest(req CreateProductRequest) error {
	err := validation.ValidateName(req.Name, "Product name")
	if err != nil {
		return err
	}

	if req.Calories < 0 {
		return model.Error{
			Code:    model.ErrorCodeBadRequest,
			Message: "Number of calories must be positive",
		}
	}

	return nil
}
