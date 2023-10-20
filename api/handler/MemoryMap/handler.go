package MemoryMap

import (
	"errors"
	memoryRequests "getir_case/api/requests/MemoryRequests"
	memoryMapResponses "getir_case/api/responses/MemoryMapResponses"
	"getir_case/api/validator"
	"getir_case/pkg/services/MemoryMap"
	"getir_case/utils"
	"net/http"
)

type Handler struct {
	Service MemoryMap.ServiceInterface
}

func NewMemoryMapHandler(memoryMapService *MemoryMap.Service) *Handler {
	return &Handler{Service: memoryMapService}
}

// GetValueByKey godoc
//
//	@Summary		Get Value
//	@Description	get value by key
//	@Tags			in_memory
//	@Accept			json
//	@Produce		json
//	@Param key query string true "key"
//	@Success 200 {object} MemoryMapResponses.KeyValue
//	@Router			/api/get_in_memory [get]
func (h *Handler) GetValueByKey(w http.ResponseWriter, r *http.Request) error {
	allowedMethods := []string{"GET"}

	utils.SetJsonHeader(w)

	if !utils.CheckMethod(allowedMethods, r) {
		return utils.WriteErrorResponse(w, "404 NOT FOUND")
	}

	key, err := h.GetKey(r)
	if err != nil {
		return err
	}

	value, err := h.Service.GetValueByKey(key)
	if err != nil {
		return utils.WriteErrorResponse(w, err.Error())
	} else {
		Response := memoryMapResponses.CreateKeyValueResponse(key, value)
		return utils.WriteJsonResponse(w, Response)
	}
}

// SetValueByKey godoc
//
//	@Summary		Set Value
//	@Description	set value by key
//	@Tags			in_memory
//	@Accept			json
//	@Produce		json
//	@Param			body	body	MemoryRequests.KeyValue	true	"Set value by key in memory database"
//	@Success 200 {object} MemoryMapResponses.KeyValue
//	@Router			/api/set_in_memory [post]
func (h *Handler) SetValueByKey(w http.ResponseWriter, r *http.Request) error {
	var input memoryRequests.KeyValue
	allowedMethods := []string{"POST"}

	utils.SetJsonHeader(w)

	if !utils.CheckMethod(allowedMethods, r) {
		return utils.WriteErrorResponse(w, "404 NOT FOUND")
	}

	if err := utils.ParseJsonBody(r.Body, &input); err != nil {
		return utils.WriteErrorResponse(w, err.Message)
	}

	if err := validator.Validate(input); err != nil {
		validationError := validator.GetValidationErrors(err)
		return utils.WriteValidationError(w, validationError)
	}

	err := h.Service.SetValueByKey(input.Key, input.Value)
	if err != nil {
		return utils.WriteErrorResponse(w, err.Error())
	} else {
		Response := memoryMapResponses.CreateKeyValueResponse(input.Key, input.Value)
		return utils.WriteJsonResponse(w, Response)
	}
}

func (h *Handler) GetKey(r *http.Request) (string, error) {
	keyParam := r.URL.Query().Get("key")

	if keyParam == "" {
		return "", errors.New("wrong param key")
	}
	return keyParam, nil
}

func (h *Handler) GetValue(r *http.Request) (string, error) {
	valueParam := r.URL.Query().Get("value")

	if valueParam == "" {
		return "", errors.New("wrong param value")
	}
	return valueParam, nil
}
