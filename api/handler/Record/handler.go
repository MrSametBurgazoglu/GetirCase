package Record

import (
	"getir_case/api/requests/RecordRequests"
	"getir_case/api/responses"
	recordResponses "getir_case/api/responses/RecordResponses"
	"getir_case/api/validator"
	"getir_case/pkg/services/Record"
	"getir_case/utils"
	"log"
	"net/http"
)

type Handler struct {
	Service Record.ServiceInterface
}

func NewRecordHandler(recordService *Record.Service) *Handler {
	return &Handler{Service: recordService}
}

// Filter godoc
//
//	@Summary		Filter Records
//	@Description	filter records by date and total count range
//	@Tags			filters
//	@Accept			json
//	@Produce		json
//	@Param			body	body	RecordRequests.FilterInput	true	"Date and Total Count Range"
//	@Success 200 {object} RecordResponses.FilterResponse
//	@Router			/api/filter_records [post]
func (h *Handler) Filter(w http.ResponseWriter, r *http.Request) error {
	var input RecordRequests.FilterInput
	allowedMethods := []string{"POST"}

	utils.SetJsonHeader(w)

	if !utils.CheckMethod(allowedMethods, r) {
		return utils.WriteErrorResponse(w, "404 NOT FOUND")
	}

	if err := utils.ParseJsonBody(r.Body, &input); err != nil {
		errResponse := recordResponses.CreateFilterResponse(err.Code, err.Message, nil)
		return utils.WriteJsonResponse(w, errResponse)
	}

	if err := validator.Validate(input); err != nil {
		validationError := validator.GetValidationErrors(err)
		responseError := utils.WriteValidationError(w, validationError)
		if responseError != nil {
			return responseError
		}
	}

	results, err := h.Service.FilterRecords(input.StartDate, input.EndDate, input.MinCount, input.MaxCount)
	if err == nil {
		response := recordResponses.CreateFilterResponse(responses.Success, responses.Success.Message(), results)
		return utils.WriteJsonResponse(w, response)
	} else {
		log.Println(err.Error())
	}

	return nil
}
