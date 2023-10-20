package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	mainHandler "getir_case/api/handler"
	memoryMapHandler "getir_case/api/handler/MemoryMap"
	memoryRequests "getir_case/api/requests/MemoryRequests"
	"getir_case/api/responses/MemoryMapResponses"
	"getir_case/config"
	"getir_case/drivers/database"
	memoryRepository "getir_case/pkg/repository/MemoryMap"
	memoryServices "getir_case/pkg/services/MemoryMap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMemory(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {

		key := "ben"
		value := "samet"

		requestBody := memoryRequests.KeyValue{
			Key:   key,
			Value: value,
		}

		requestBodyJson, err := json.Marshal(requestBody)

		if err != nil {
			t.Errorf(err.Error())
		}

		setRequest, _ := http.NewRequest(http.MethodPost, "/api/set_in_memory", bytes.NewBuffer(requestBodyJson))
		getRequest, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/get_in_memory?key=%s", key), nil)
		setResponse := httptest.NewRecorder()
		getResponse := httptest.NewRecorder()

		config.InitConfig()

		db := database.SetupDatabase()

		repository := &memoryRepository.Repository{DB: db}

		service := &memoryServices.Service{Repository: repository}

		newMemoryMapHandler := memoryMapHandler.NewMemoryMapHandler(service)

		endpointHandler := mainHandler.Handler(newMemoryMapHandler.SetValueByKey)

		endpointHandler2 := mainHandler.Handler(newMemoryMapHandler.GetValueByKey)

		endpointHandler.ServeHTTP(setResponse, setRequest)

		endpointHandler2.ServeHTTP(getResponse, getRequest)

		setKeyValueResponse := MemoryMapResponses.KeyValue{}

		err = json.Unmarshal(getResponse.Body.Bytes(), &setKeyValueResponse)
		if err != nil {
			t.Errorf("error on response")
		}

		if setKeyValueResponse.Key != key || setKeyValueResponse.Value != value {
			t.Errorf("got %s,%s, want %s,%s", setKeyValueResponse.Key, setKeyValueResponse.Value, key, value)
		}
	})
}
