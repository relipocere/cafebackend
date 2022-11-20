package mapping

import (
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/model"
)

// MapToUserKind maps to business entity UserKind.
func MapToUserKind(userKind graphmodel.UserKindEnum) model.UserKind {
	switch userKind {
	case graphmodel.UserKindEnumBusiness:
		return model.UserKindBusiness
	case graphmodel.UserKindEnumConsumer:
		return model.UserKindConsumer
	}

	return model.UserKind(userKind.String())
}

// MapToAffordability maps affordability to business entity Affordability.
func MapToAffordability(affordability graphmodel.Affordability) model.Affordability {
	switch affordability {
	case graphmodel.AffordabilityCheap:
		return model.AffordabilityCheap
	case graphmodel.AffordabilityAffordable:
		return model.AffordabilityAffordable
	case graphmodel.AffordabilityExpensive:
		return model.AffordabilityExpensive
	}

	return model.Affordability(affordability.String())
}

// MapToCuisine maps cuisine type to business entity.
func MapToCuisine(cuisine graphmodel.CuisineType) model.Cuisine {
	switch cuisine {
	case graphmodel.CuisineTypeAmerican:
		return model.CuisineAmerican
	case graphmodel.CuisineTypeAsian:
		return model.CuisineAsian
	case graphmodel.CuisineTypeEuropean:
		return model.CuisineEuropean
	}

	return model.Cuisine(cuisine.String())
}
