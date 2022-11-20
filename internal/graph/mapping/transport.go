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
		UUID:     user.ID,
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

	return "", fmt.Errorf("unkown user kind: %s", kind)
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
