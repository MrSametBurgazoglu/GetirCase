package utils

import (
	"encoding/json"
	"getir_case/api/responses"
	"getir_case/pkg/errors"
	"io"
)

func ParseJsonBody(closer io.ReadCloser, model any) *errors.ParseJsonError {
	err := json.NewDecoder(closer).Decode(model)
	if err != nil {
		return &errors.ParseJsonError{Code: responses.ParseJsonError, Message: responses.ParseJsonError.Message()}
	}
	return nil
}
