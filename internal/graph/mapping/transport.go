package mapping

import (
	"fmt"

	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/model"
)

// MapUser maps business user to graphql user.
func MapUser(user model.User) (graphmodel.User, error) {
	kind, err := mapUserKind(user.Kind)
	if err != nil {
		return graphmodel.User{}, err
	}

	return graphmodel.User{
		Username: user.Username,
		Kind:     kind,
		FullName: user.FullName,
	}, nil
}

func mapUserKind(kind model.UserKind) (graphmodel.UserKindEnum, error) {
	switch kind {
	case model.UserKindBusiness:
		return graphmodel.UserKindEnumBusiness, nil
	case model.UserKindConsumer:
		return graphmodel.UserKindEnumConsumer, nil
	}

	return "", fmt.Errorf("unknown user kind: %s", kind)
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
func MapAffordability(affordability model.Affordability) graphmodel.Affordability {
	switch affordability {
	case model.AffordabilityCheap:
		return graphmodel.AffordabilityCheap
	case model.AffordabilityAffordable:
		return graphmodel.AffordabilityAffordable
	case model.AffordabilityExpensive:
		return graphmodel.AffordabilityExpensive
	}

	return graphmodel.Affordability(string(affordability))
}

// MapCuisine maps cuisine type to gql entity.
func MapCuisine(cuisine model.Cuisine) graphmodel.CuisineType {
	switch cuisine {
	case model.CuisineAmerican:
		return graphmodel.CuisineTypeAmerican
	case model.CuisineAsian:
		return graphmodel.CuisineTypeAsian
	case model.CuisineEuropean:
		return graphmodel.CuisineTypeEuropean
	}

	return graphmodel.CuisineType(string(cuisine))
}
