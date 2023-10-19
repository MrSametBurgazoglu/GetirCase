package utils

import (
	"encoding/json"
	"getir_case/api/responses"
	"getir_case/pkg/errors"
	"io"
	"net/http"
)

func SetJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func WriteJsonResponse(writer io.Writer, response any) error {
	return json.NewEncoder(writer).Encode(response)
}

func WriteErrorResponse(writer io.Writer, errorMessage string) error {
	standardError := responses.ErrorResponse{Error: errorMessage}
	return WriteJsonResponse(writer, standardError)
}

func WriteValidationError(writer io.Writer, message string) error {
	validationError := errors.ValidationError{Code: responses.ValidationError, Message: message}
	return json.NewEncoder(writer).Encode(validationError)
}
