package producthandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/auth"
	"github.com/relipocere/cafebackend/internal/database"
	storedb "github.com/relipocere/cafebackend/internal/database/store"
	"github.com/relipocere/cafebackend/internal/model"
)

// DeleteProduct handles the scenario of product deletion.
func (h *Handler) DeleteProduct(ctx context.Context, productID int64) error {
	tx, err := h.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer database.RollbackTx(ctx, tx, "DeleteProduct")

	user, ok := ctx.Value(auth.User).(model.User)
	if !ok {
		return fmt.Errorf("no user in the context")
	}

	stores, err := h.storeRepo.Get(ctx, tx, storedb.GetByProductIDs([]int64{productID}))
	if err != nil {
		return fmt.Errorf("getting store by product %d: %w", productID, err)
	}

	if len(stores) == 0 {
		return model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: fmt.Sprintf("Store containing product %d doesn't exist", productID),
		}
	}

	store := stores[0]
	if store.OwnerUsername != user.Username {
		return model.Error{
			Code:    model.ErrorCodeFailedPrecondition,
			Message: "You're not the owner of the store this product listed in",
		}
	}

	err = h.productRepo.Delete(ctx, tx, []int64{productID})
	if err != nil {
		return fmt.Errorf("deleting product %d: %w", productID, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}
