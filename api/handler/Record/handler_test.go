package Record

import (
	"encoding/json"
	mainHandler "getir_case/api/handler"
	recordResponses "getir_case/api/responses/Record"
	"getir_case/drivers/database"
	"getir_case/pkg/models"
	"getir_case/pkg/repository/Record"
	RecordService "getir_case/pkg/services/Record"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMemory(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		getRequest, _ := http.NewRequest(http.MethodGet, "/api/filter_records", nil)
		getResponse := httptest.NewRecorder()

		db := database.SetupDatabase()

		repository := &Record.Repository{DB: db}

		service := &RecordService.Service{Repository: repository}

		recordHandler := NewRecordHandler(service)

		endpointHandler := mainHandler.Handler(recordHandler.Filter)

		endpointHandler.ServeHTTP(getResponse, getRequest)

		gotResponse := recordResponses.FilterResponse{}

		err := json.Unmarshal(getResponse.Body.Bytes(), &gotResponse)
		if err != nil {
			t.Errorf("error on response")
		}

		wantedResponse := recordResponses.FilterResponse{
			Code: 0,
			Msg:  "success",
			Records: []*models.RecordFilterModel{
				{Key: "TAKwGc6Jr4i8Z487", CreatedAt: "2017-01-28 04:22:14.398 +0300 +03", TotalCount: 2800},
				{Key: "NAeQ8eX7e5TEg7oH", CreatedAt: "2017-01-27 11:19:14.135 +0300 +03", TotalCount: 2900},
			},
		}

		if wantedResponse.Code != gotResponse.Code {
			t.Errorf("got %q, want %q", gotResponse.Code, wantedResponse.Code)
		}

		if wantedResponse.Msg != gotResponse.Msg {
			t.Errorf("got %q, want %q", gotResponse.Msg, wantedResponse.Msg)
		}

		if len(wantedResponse.Records) != len(gotResponse.Records) {
			t.Errorf("record count not equal got %q, want %q", len(gotResponse.Records), len(wantedResponse.Records))
		}

		for i, response := range wantedResponse.Records {
			if response.Key != gotResponse.Records[i].Key || response.CreatedAt != gotResponse.Records[i].CreatedAt || response.TotalCount != gotResponse.Records[i].TotalCount {
				t.Errorf("got %q, want %q", response, wantedResponse)
			}
		}
	})
}
