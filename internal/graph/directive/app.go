package directive

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/relipocere/cafebackend/internal/model"
)

// App contains directives implementations.
type App struct {
}

// NewApp creates App.
func NewApp() *App {
	return &App{}
}

// IsAuthenticated implements IsAuthenticated directive.
func (a *App) IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	unauthenticatedErr := model.Error{
		Message: model.ErrMessageUnauthenticated(),
		Code:    model.ErrorCodeUnauthenticated,
	}

	userVal := ctx.Value("user")
	if userVal == nil {
		return nil, unauthenticatedErr
	}

	user, ok := userVal.(*model.User)
	if !ok || user == nil {
		return nil, unauthenticatedErr
	}

	return next(ctx)
}
