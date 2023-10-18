package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "e164":
		return "Invalid phone number"
	}
	return ""
}

func GetValidationErrors(err error) string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fe := range ve {
			if fe != nil {
				errorCause := fmt.Sprintf("%s: %s", fe.Field(), msgForTag(fe.Tag()))
				return errorCause
			}
		}
	}
	return ""
}

func Validate(model any) error {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		return err
	}
	return nil
}
