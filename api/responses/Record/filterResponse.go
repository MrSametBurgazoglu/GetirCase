package Record

import (
	"getir_case/api/responses"
	"getir_case/pkg/models"
)

type FilterResponse struct {
	Code    responses.ResponseTypes
	Msg     string
	Records []*models.RecordFilterModel
}

func CreateFilterResponse(code responses.ResponseTypes, message string, records []*models.RecordFilterModel) FilterResponse {
	return FilterResponse{Code: code, Msg: message, Records: records}
}
