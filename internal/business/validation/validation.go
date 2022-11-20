package validation

import (
	"unicode/utf8"

	"github.com/relipocere/cafebackend/internal/model"
)

// ValidateUsername validates username.
func ValidateUsername(username string) error {
	return validateLength(username, 20, "username")
}

func validateLength(s string, maxLen int64, fieldName string) error {
	if int64(utf8.RuneCountInString(s)) > maxLen {
		return model.Error{
			Message: model.ErrMessageMaxLengthExceeded(fieldName, maxLen),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	return nil
}
