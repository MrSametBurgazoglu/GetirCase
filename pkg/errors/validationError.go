package errors

import "getir_case/api/responses"

type ValidationError struct {
	Code    responses.ResponseTypes
	Message string
}

func (receiver ValidationError) Error() string {
	return receiver.Message
}
