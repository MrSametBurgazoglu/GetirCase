package errors

import "getir_case/api/responses"

type ParseJsonError struct {
	Code    responses.ResponseTypes
	Message string
}

func (receiver ParseJsonError) Error() string {
	return receiver.Message
}
