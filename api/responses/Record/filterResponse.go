package Record

import (
	"getir_case/api/responses"
	"getir_case/pkg/models"
	"io"
)

type FilterResponse struct {
	Code    responses.ResponseTypes
	Msg     string
	Records []*models.RecordFilterModel
}

func CreateFilterResponse(code responses.ResponseTypes, message string, records []*models.RecordFilterModel) *FilterResponse {
	filterResponse := new(FilterResponse)
	filterResponse.Code = code
	filterResponse.Msg = message
	filterResponse.Records = records
	return filterResponse
}

func (receiver *FilterResponse) WriteResponse(writer io.Writer) {

}
