package MemoryMap

import (
	"encoding/json"
	mainHandler "getir_case/api/handler"
	"getir_case/api/responses/MemoryMap"
	"getir_case/drivers/database"
	memoryRepository "getir_case/pkg/repository/MemoryMap"
	memoryServices "getir_case/pkg/services/MemoryMap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMemory(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		setRequest, _ := http.NewRequest(http.MethodPost, "/api/set_in_memory?key=ben&value=samet", nil)
		getRequest, _ := http.NewRequest(http.MethodGet, "/api/set_in_memory?key=ben", nil)
		setResponse := httptest.NewRecorder()
		getResponse := httptest.NewRecorder()

		db := database.SetupDatabase()

		repository := &memoryRepository.Repository{DB: db}

		service := &memoryServices.Service{Repository: repository}

		memoryMapHandler := NewMemoryMapHandler(service)

		endpointHandler := mainHandler.Handler(memoryMapHandler.GetValueByKey)

		endpointHandler2 := mainHandler.Handler(memoryMapHandler.GetValueByKey)

		endpointHandler.ServeHTTP(setResponse, setRequest)

		endpointHandler2.ServeHTTP(getResponse, getRequest)

		setKeyValueResponse := MemoryMap.KeyValue{}

		err := json.Unmarshal(getResponse.Body.Bytes(), &setKeyValueResponse)
		if err != nil {
			t.Errorf("error on response")
		}

		key := "ben"
		value := "samet"

		if setKeyValueResponse.Key != key || setKeyValueResponse.Value != value {
			t.Errorf("got %s,%s, want %s,%s", setKeyValueResponse.Key, setKeyValueResponse.Value, key, value)
		}
	})
}
