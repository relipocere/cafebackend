package userhandler

import (
	"context"
	"fmt"

	"github.com/relipocere/cafebackend/internal/business/validation"
	"github.com/relipocere/cafebackend/internal/model"
	"github.com/relipocere/cafebackend/internal/service/security"
)

// CreateUserRequest is the request for.
type CreateUserRequest struct {
	Username string
	FullName string
	Kind     model.UserKind
	Password string
}

// CreateUser handles user creation scenario.
func (h *Handler) CreateUser(ctx context.Context, req CreateUserRequest) error {
	validationErr := validateCreateUserRequest(req)
	if validationErr != nil {
		return validationErr
	}

	now := h.now()
	salt, err := security.GenerateSalt()
	if err != nil {
		return fmt.Errorf("salt generation: %w", err)
	}
	passwordHash := security.Hash(req.Password, salt)

	_, err = h.userRepo.Create(ctx, h.db, model.User{
		Username:     req.Username,
		FullName:     req.FullName,
		Kind:         req.Kind,
		PasswordHash: passwordHash,
		Salt:         salt,
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if err != nil {
		return fmt.Errorf("user creation: %w", err)
	}

	return nil
}

func validateCreateUserRequest(req CreateUserRequest) error {
	if req.Username == "" {
		return model.Error{
			Message: model.ErrMessageMissingFieldRequired("Username"),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	if req.Password == "" {
		return model.Error{
			Message: model.ErrMessageMissingFieldRequired("Password"),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	if req.FullName == "" {
		return model.Error{
			Message: model.ErrMessageMissingFieldRequired("Full name"),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	if !req.Kind.IsValid() {
		return model.Error{
			Message: fmt.Sprintf("Invalid account kind: %s", string(req.Kind)),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	err := validation.ValidateUsername(req.Username)
	if err != nil {
		return err
	}

	return nil
}
