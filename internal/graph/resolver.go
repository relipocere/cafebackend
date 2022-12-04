package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	producthandler "github.com/relipocere/cafebackend/internal/business/product-handler"
	storehandler "github.com/relipocere/cafebackend/internal/business/store-handler"
	userhandler "github.com/relipocere/cafebackend/internal/business/user-handler"
	"github.com/relipocere/cafebackend/internal/database"
	"github.com/relipocere/cafebackend/internal/graph/directive"
	"github.com/relipocere/cafebackend/internal/graph/generated"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/graph/image"
	"github.com/relipocere/cafebackend/internal/graph/product"
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

type productHandler interface {
	CreateProdcut(ctx context.Context, req producthandler.CreateProductRequest) (model.Product, error)
	DeleteProduct(ctx context.Context, productID int64) error
	SearchProducts(ctx context.Context, req producthandler.SearchProductsRequest) ([]model.Product, error)
}

type imageRepo interface {
	Create(ctx context.Context, q database.Queryable, image model.ImageMeta) error
	Get(ctx context.Context, q database.Queryable, imageID string) (*model.ImageMeta, error)
}

func NewResolver(
	filesDir string,
	db database.PGX,
	imageRepo imageRepo,
	userHandler userHandler,
	storeHandler storeHandler,
	productHandler productHandler,
) generated.Config {
	userApp := user.NewApp(userHandler)
	storeApp := store.NewApp(storeHandler)
	directiveApp := directive.NewApp()
	imageApp := image.NewApp(filesDir, db, imageRepo)
	productApp := product.NewApp(productHandler)

	cfg := generated.Config{
		Resolvers: &Resolver{
			mutationResolver: &mutationResolver{
				user:    userApp,
				store:   storeApp,
				image:   imageApp,
				product: productApp,
			},

			queryResolver: &queryResolver{
				user:    userApp,
				store:   storeApp,
				product: productApp,
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
	user    *user.App
	store   *store.App
	image   *image.App
	product *product.App
}

type queryResolver struct {
	user    *user.App
	store   *store.App
	product *product.App
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return r.mutationResolver
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

func (m *mutationResolver) UploadImage(ctx context.Context, image graphql.Upload) (string, error) {
	return m.image.UploadImage(ctx, image)
}

func (m *mutationResolver) CreateProduct(ctx context.Context, input graphmodel.CreateProductInput) (graphmodel.Product, error) {
	return m.product.CreateProduct(ctx, input)
}
func (m *mutationResolver) DeleteProduct(ctx context.Context, productID int64) (bool, error) {
	return m.product.DeleteProduct(ctx, productID)
}

func (r *Resolver) Query() generated.QueryResolver {
	return r.queryResolver
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

func (q *queryResolver) SearchProducts(ctx context.Context, input graphmodel.SearchProductsInput) ([]graphmodel.Product, error) {
	return q.product.SearchProducts(ctx, input)
}
