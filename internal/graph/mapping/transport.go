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
