package mapping

import (
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/model"
)

// MapToUserKind maps to business entity UserKind.
func MapToUserKind(enum graphmodel.UserKindEnum) model.UserKind {
	switch enum {
	case graphmodel.UserKindEnumBusiness:
		return model.UserKindBusiness
	case graphmodel.UserKindEnumConsumer:
		return model.UserKindConsumer
	}

	return model.UserKind(enum.String())
}
