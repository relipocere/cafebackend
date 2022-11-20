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

	// if can't cast to user or no value is present
	_, ok := ctx.Value("user").(model.User)
	if !ok {
		return nil, unauthenticatedErr
	}

	return next(ctx)
}
