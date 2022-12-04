package mapping

import (
	"fmt"

	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/model"
)

// MapProducts maps business products to graphql products.
func MapProducts(pp []model.Product) []graphmodel.Product {
	products := make([]graphmodel.Product, 0, len(pp))

	for _, p := range pp {
		products = append(products, MapProduct(p))
	}

	return products
}

// MapProduct maps business product to graphql product.
func MapProduct(p model.Product) graphmodel.Product {
	return graphmodel.Product{
		ID:          p.ID,
		Name:        p.Name,
		StoreID:     p.StoreID,
		PriceCents:  p.PriceCents,
		Ingredients: p.Ingredients,
		Calories:    p.Calories,
		ImageID:     p.ImageID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

// MapUser maps business user to graphql user.
func MapUser(u model.User) (graphmodel.User, error) {
	kind, err := mapUserKind(u.Kind)
	if err != nil {
		return graphmodel.User{}, err
	}

	return graphmodel.User{
		Username: u.Username,
		Kind:     kind,
		FullName: u.FullName,
	}, nil
}

func mapUserKind(k model.UserKind) (graphmodel.UserKindEnum, error) {
	switch k {
	case model.UserKindBusiness:
		return graphmodel.UserKindEnumBusiness, nil
	case model.UserKindConsumer:
		return graphmodel.UserKindEnumConsumer, nil
	}

	return "", fmt.Errorf("unknown user kind: %s", k)
}

func MapStores(ss []model.Store) []graphmodel.Store {
	stores := make([]graphmodel.Store, 0, len(ss))

	for _, s := range ss {
		stores = append(stores, MapStore(s))
	}

	return stores
}

// MapStore maps store to gql entity Store.
func MapStore(s model.Store) graphmodel.Store {
	return graphmodel.Store{
		ID:            s.ID,
		Title:         s.Title,
		Affordability: MapAffordability(s.Affordability),
		CuisineType:   MapCuisine(s.Cuisine),
		OwnerUsername: s.OwnerUsername,
		ImageID:       s.ImageID,
		AverageRating: s.AverageRating,
		CreatedAt:     s.CreatedAt,
		UpdatedAt:     s.UpdatedAt,
	}
}

// MapAffordability maps affordability to gql entity Affordability.
func MapAffordability(a model.Affordability) graphmodel.Affordability {
	switch a {
	case model.AffordabilityCheap:
		return graphmodel.AffordabilityCheap
	case model.AffordabilityAffordable:
		return graphmodel.AffordabilityAffordable
	case model.AffordabilityExpensive:
		return graphmodel.AffordabilityExpensive
	}

	return graphmodel.Affordability(string(a))
}

// MapCuisine maps cuisine type to gql entity.
func MapCuisine(c model.Cuisine) graphmodel.CuisineType {
	switch c {
	case model.CuisineAmerican:
		return graphmodel.CuisineTypeAmerican
	case model.CuisineAsian:
		return graphmodel.CuisineTypeAsian
	case model.CuisineEuropean:
		return graphmodel.CuisineTypeEuropean
	}

	return graphmodel.CuisineType(string(c))
}
