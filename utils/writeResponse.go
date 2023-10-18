package utils

import (
	"encoding/json"
	"getir_case/api/responses"
	"getir_case/pkg/errors"
	"io"
)

func WriteResponse(writer io.Writer, response any) error {
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		return err
	}
	return nil
}

func SetHeaderForJson(writer io.Writer) {
}

func WriteValidationError(writer io.Writer, message string) error {
	validationError := errors.ValidationError{Code: responses.ValidationError, Message: message}
	err := json.NewEncoder(writer).Encode(validationError)
	if err != nil {
		return err
	}
	return nil
}
