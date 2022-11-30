package mapping

import (
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/model"
)

// MapToUserKind maps to business entity UserKind.
func MapToUserKind(k graphmodel.UserKindEnum) model.UserKind {
	switch k {
	case graphmodel.UserKindEnumBusiness:
		return model.UserKindBusiness
	case graphmodel.UserKindEnumConsumer:
		return model.UserKindConsumer
	}

	return model.UserKind(k.String())
}

func MapToAffordabilitySlice(aa []graphmodel.Affordability) []model.Affordability {
	var affordability []model.Affordability

	for _, a := range aa {
		affordability = append(affordability, MapToAffordability(a))
	}

	return affordability
}

// MapToAffordability maps affordability to business entity Affordability.
func MapToAffordability(a graphmodel.Affordability) model.Affordability {
	switch a {
	case graphmodel.AffordabilityCheap:
		return model.AffordabilityCheap
	case graphmodel.AffordabilityAffordable:
		return model.AffordabilityAffordable
	case graphmodel.AffordabilityExpensive:
		return model.AffordabilityExpensive
	}

	return model.Affordability(a.String())
}

func MapToCuisines(cc []graphmodel.CuisineType ) []model.Cuisine{
	var cuisines []model.Cuisine

	for _, c := range cc {
		cuisines = append(cuisines, MapToCuisine(c))
	}

	return cuisines 
}

// MapToCuisine maps cuisine type to business entity.
func MapToCuisine(c graphmodel.CuisineType) model.Cuisine {
	switch c {
	case graphmodel.CuisineTypeAmerican:
		return model.CuisineAmerican
	case graphmodel.CuisineTypeAsian:
		return model.CuisineAsian
	case graphmodel.CuisineTypeEuropean:
		return model.CuisineEuropean
	}

	return model.Cuisine(c.String())
}

func MapToPagination(p graphmodel.Pagination) model.Pagination {
	return model.Pagination{
		Page:         p.Page,
		ItemsPerPage: p.Limit,
	}
}

func MapToIntRange(r *graphmodel.IntRange) *model.IntRange {
	if r == nil{
		return nil
	}

	return &model.IntRange{
		Start:          r.Start,
		End:            r.End,
		StartExclusive: r.StartExclusive,
		EndExclusive:   r.EndExclusive,
	}
}
