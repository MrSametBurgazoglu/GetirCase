package handler

import (
	"getir_case/api/handler/MemoryMap"
	"getir_case/api/handler/Record"
	"getir_case/pkg/services"
	"log"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		log.Println(err.Error())
	}
}

func SetupHandlers(apiServices *services.Services) {

	recordHandler := Record.NewRecordHandler(apiServices.RecordService)
	http.Handle("api/filter_records", Handler(recordHandler.Filter))

	memoryMapHandler := MemoryMap.NewMemoryMapHandler(apiServices.MemoryMapService)
	http.Handle("api/get_in_memory", Handler(memoryMapHandler.GetValueByKey))
	http.Handle("api/set_in_memory", Handler(memoryMapHandler.SetValueByKey))
}
