package services

import (
	"getir_case/drivers/database"
	"getir_case/pkg/repository"
	"getir_case/pkg/services/MemoryMap"
	"getir_case/pkg/services/Record"
)

type Services struct {
	RecordService    *Record.Service
	MemoryMapService *MemoryMap.Service
}

func SetupServices(db *database.Database) *Services {
	repositories := repository.NewRepositories(db)
	recordService := &Record.Service{Repository: repositories.Record}
	memoryMapService := &MemoryMap.Service{Repository: repositories.MemoryMap}

	return &Services{
		RecordService:    recordService,
		MemoryMapService: memoryMapService,
	}
}
