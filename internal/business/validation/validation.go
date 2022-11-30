package validation

import (
	"fmt"
	"unicode/utf8"

	"github.com/relipocere/cafebackend/internal/model"
)

func ValidatePagination(page model.Pagination) error {
	if page.Page < 1 {
		return model.Error{
			Code: model.ErrorCodeBadRequest,
			Message: fmt.Sprintf("Page number can't be less than 1. Provided value: %d", page.Page),
		}
	}

	if page.ItemsPerPage < 1{
		return model.Error{
			Code: model.ErrorCodeBadRequest,
			Message: fmt.Sprintf("Number of items per page can't be less than 1. Provided value: %d", page.Page),
		}
	}

	return nil
}

// ValidateUsername validates username.
func ValidateUsername(username string) error {
	return validateLength(username, 20, "username")
}

func ValidateUsernames(usernames []string) error{
	for _, username := range usernames{
		err := ValidateUsername(username)
		if err != nil {
			return err
		}
	}

	return nil
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

func ValidateAffordability(affordability model.Affordability) error {
	if !affordability.IsValid() {
		return model.Error{
			Message: fmt.Sprintf("Invalid affordability: %s", string(affordability)),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	return nil
}

func ValidateSliceOfAffordability(aa []model.Affordability) error {
	for _, a := range aa {
		err := ValidateAffordability(a)
		if err != nil {
			return err
		}
	}

	return nil
}

func ValidateCuisine(cuisine model.Cuisine) error {
	if !cuisine.IsValid() {
		return model.Error{
			Message: fmt.Sprintf("Invalid cuisine: %s", string(cuisine)),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	return nil
}

func ValidateCuisines(cuisines []model.Cuisine) error {
	for _, cuisine := range cuisines {
		err := ValidateCuisine(cuisine)
		if err != nil {
			return err
		}
	}

	return nil
}
