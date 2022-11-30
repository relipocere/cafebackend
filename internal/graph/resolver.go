package graph

import (
	"context"

	storehandler "github.com/relipocere/cafebackend/internal/business/store-handler"
	userhandler "github.com/relipocere/cafebackend/internal/business/user-handler"
	"github.com/relipocere/cafebackend/internal/graph/directive"
	"github.com/relipocere/cafebackend/internal/graph/generated"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/graph/store"
	"github.com/relipocere/cafebackend/internal/graph/user"
	"github.com/relipocere/cafebackend/internal/model"
)

type userHandler interface {
	CreateUser(ctx context.Context, req userhandler.CreateUserRequest) error
	LogIn(ctx context.Context, req userhandler.LogInRequest) (userhandler.LogInResponse, error)
}

type storeHandler interface {
	CreateStore(ctx context.Context, req storehandler.CreateStoreRequest) (storehandler.CreateStoreResponse, error)
	DeleteStore(ctx context.Context, req storehandler.DeleteStoreRequest) error
	SearchStores(ctx context.Context, req storehandler.SearchStoresRequest) ([]model.Store, error)
}

func NewResolver(
	userHandler userHandler,
	storeHandler storeHandler,
) generated.Config {
	userApp := user.NewApp(userHandler)
	storeApp := store.NewApp(storeHandler)
	directiveApp := directive.NewApp()

	cfg := generated.Config{
		Resolvers: &Resolver{
			mutationResolver: &mutationResolver{
				user:  userApp,
				store: storeApp,
			},

			queryResolver: &queryResolver{
				user:  userApp,
				store: storeApp,
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
	user  *user.App
	store *store.App
}

type queryResolver struct {
	user  *user.App
	store *store.App
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

func (m *mutationResolver) CreateStore(ctx context.Context, input graphmodel.CreateStoreInput) (graphmodel.Store, error) {
	return m.store.CreateStore(ctx, input)
}

func (m *mutationResolver) DeleteStore(ctx context.Context, input graphmodel.DeleteStoreInput) (bool, error) {
	return m.store.DeleteStore(ctx, input)
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

func (q *queryResolver) SearchStores(ctx context.Context, input graphmodel.SearchStoresInput) ([]graphmodel.Store, error) {
	return q.store.SearchStores(ctx, input)
}
