package Record

import (
	recordRequests "getir_case/api/requests/Record"
	"getir_case/api/responses"
	recordResponses "getir_case/api/responses/Record"
	"getir_case/api/validator"
	"getir_case/pkg/services/Record"
	"getir_case/utils"
	"net/http"
)

type Handler struct {
	Service Record.ServiceInterface
}

func NewRecordHandler(recordService *Record.Service) *Handler {
	return &Handler{Service: recordService}
}

func (h *Handler) Filter(w http.ResponseWriter, r *http.Request) error {
	var input recordRequests.FilterInput
	if err := utils.ParseJsonBody(r.Body, &input); err != nil {
		errResponse := recordResponses.CreateFilterResponse(err.Code, err.Message, nil)
		return utils.WriteResponse(w, errResponse)
	}

	if err := validator.Validate(input); err != nil {
		validationError := validator.GetValidationErrors(err)
		responseError := utils.WriteValidationError(w, validationError)
		if responseError != nil {
			return responseError
		}
	}

	results, err := h.Service.FilterRecords(input.StartDate, input.EndDate, input.MinCount, input.MaxCount)
	if err != nil {
		errResponse := recordResponses.CreateFilterResponse(responses.Success, responses.Success.Message(), results)
		return utils.WriteResponse(w, errResponse)
	}
	return nil
}
