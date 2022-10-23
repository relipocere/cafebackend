package graph

import (
	"context"

	userhandler "github.com/relipocere/cafebackend/internal/business/user-handler"
	"github.com/relipocere/cafebackend/internal/graph/directive"
	"github.com/relipocere/cafebackend/internal/graph/generated"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/graph/user"
)

type userHandler interface {
	CreateUser(ctx context.Context, req userhandler.CreateUserRequest) error
	LogIn(ctx context.Context, req userhandler.LogInRequest) (userhandler.LogInResponse, error)
}

func NewResolver(
	userHandler userHandler,
) generated.Config {
	userApp := user.NewApp(userHandler)
	directiveApp := directive.NewApp()

	cfg := generated.Config{
		Resolvers: &Resolver{
			mutationResolver: &mutationResolver{
				user: userApp,
			},

			queryResolver: &queryResolver{
				user: userApp,
			},
		},
	}

	cfg.Directives.IsAuthenticated = directiveApp.IsAuthenticated
	return cfg
}

type Resolver struct {
	mutationResolver *mutationResolver
	queryResolver    *queryResolver
}

type mutationResolver struct {
	user *user.App
}

type queryResolver struct {
	user *user.App
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return r.mutationResolver
}

func (r *Resolver) Query() generated.QueryResolver {
	return r.queryResolver
}

func (m *mutationResolver) CreateUser(ctx context.Context, input graphmodel.CreateUserInput) (bool, error) {
	return m.user.CreateUser(ctx, input)
}

func (q *queryResolver) GetAuthToken(
	ctx context.Context,
	input graphmodel.GetAuthTokenInput,
) (graphmodel.GetAuthTokenPayload, error) {
	return q.user.GetAuthToken(ctx, input)
}

func (q *queryResolver) Me(ctx context.Context) (graphmodel.User, error) {
	return q.user.Me(ctx)
}
