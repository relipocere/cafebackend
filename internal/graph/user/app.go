package user

import (
	"context"
	"fmt"

	userhandler "github.com/relipocere/cafebackend/internal/business/user-handler"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/graph/mapping"
	"github.com/relipocere/cafebackend/internal/model"
)

type userHandler interface {
	CreateUser(ctx context.Context, req userhandler.CreateUserRequest) error
	LogIn(ctx context.Context, req userhandler.LogInRequest) (userhandler.LogInResponse, error)
}

// App is the user-handler related handlers' app.
type App struct {
	userHandler userHandler
}

// NewApp creates App.
func NewApp(userHandler userHandler) *App {
	return &App{userHandler: userHandler}
}

// CreateUser implements graphql mutation CreateUser.
func (a *App) CreateUser(ctx context.Context, input graphmodel.CreateUserInput) (bool, error) {
	err := a.userHandler.CreateUser(ctx, userhandler.CreateUserRequest{
		Username: input.Username,
		FullName: input.FullName,
		Kind:     mapping.MapToUserKind(input.Kind),
		Password: input.Password,
	})
	if err != nil {
		return false, fmt.Errorf("business handler: %w", err)
	}

	return true, nil
}

// GetAuthToken implements graphql mutation GetAuthToken.
func (a *App) GetAuthToken(
	ctx context.Context,
	input graphmodel.GetAuthTokenInput,
) (graphmodel.GetAuthTokenPayload, error) {
	var resp graphmodel.GetAuthTokenPayload

	handlerResp, err := a.userHandler.LogIn(ctx, userhandler.LogInRequest{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		return resp, fmt.Errorf("business handler: %w", err)
	}

	resp.Token = handlerResp.Token
	return resp, nil
}

func (a *App) Me(ctx context.Context) (graphmodel.User, error) {
	userVal := ctx.Value("user")
	if userVal == nil {
		return graphmodel.User{}, nil
	}

	user, ok := userVal.(*model.User)
	if !ok || user == nil {
		return graphmodel.User{}, nil
	}

	resp, err := mapping.MapUser(*user)
	if err != nil {
		return graphmodel.User{}, fmt.Errorf("mapping user: %w", err)
	}

	return resp, nil
}
