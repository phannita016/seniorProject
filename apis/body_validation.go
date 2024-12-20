package apis

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var val = validator.New()

func Validation(i any) []*ErrorValidation {
	var errors []*ErrorValidation
	if err := val.Struct(i); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			var element ErrorValidation
			element.Tag = e.Tag()
			element.Field = strings.ToLower(e.Field())
			element.Message = "Error:Field validation for '" + e.Field() + "' failed on the '" + e.Tag() + "' tag"
			errors = append(errors, &element)
		}
	}

	return errors
}
