package services

import (
	"getir_case/drivers/database"
	"getir_case/pkg/repository"
	"getir_case/pkg/services/Record"
)

type Services struct {
	RecordService *Record.Service
}

func SetupServices(db *database.Database) *Services {
	repositories := repository.NewRepositories(db)
	recordService := &Record.Service{Repository: repositories.Record}

	return &Services{
		RecordService: recordService,
	}
}
