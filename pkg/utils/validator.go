package utils

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"server.com/pkg/types"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()
	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
// Return IsValid, FieldError
func ValidatorErrors(structure interface{}) (bool, []types.ValidateFieldError) {
	validate := NewValidator()

	err := validate.Struct(structure)
	if err == nil {
		return true, []types.ValidateFieldError{}
	}

	// Make error message for each invalid field.
	errors := Map(err.(validator.ValidationErrors), func(item validator.FieldError, _ int) types.ValidateFieldError {
		Field := item.Field()
		field := Field[:0] + strings.ToLower(Field[:1]) + Field[0+1:]

		return types.ValidateFieldError{
			Code:    10013,
			Field:   field,
			Message: item.Tag(),
		}
	})

	return false, errors
}

// verify strings if strings do not have special character
func VerifyMatchStr(str string) bool {
	// check if string has not alphabet and return false
	if !IsHasAlphabet(str) {
		return false
	}
	
	isMatch := regexp.MustCompile(`^[A-Za-z0-9_-]*$`).MatchString(str)
	return isMatch
}

// check string if has alphaber from a to z
func IsHasAlphabet(str string) bool {
	f := func(r rune) bool {
		return r > 'a' || r > 'z' || r == 'a'
	}
	return strings.IndexFunc(strings.ToLower(str), f) != -1
}
