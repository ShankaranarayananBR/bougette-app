package handlars

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type ValidationError struct {
	Error     string `json:"error"`
	Key       string `json:"key"`
	Condition string `json:"condition"`
}

func (h *Handler) ValidateBodyRequest(c echo.Context, payload interface{}) []*ValidationError {
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())
	var errors []*ValidationError
	err := validate.Struct(payload)
	validationerrors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, validation := range validationerrors {
			fmt.Println(validation.Field())
			fmt.Println(validation.Tag())       // what failed
			fmt.Println(validation.ActualTag()) // field in the json payload that failed
			currentValidationError := &ValidationError{
				Error:     "e",
				Key:       "k",
				Condition: "required",
			}
			errors = append(errors, currentValidationError)
		}
	}

	return errors
}
