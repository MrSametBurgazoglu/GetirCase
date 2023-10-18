package repository

import (
	"getir_case/drivers/database"
	"getir_case/pkg/repository/Record"
)

type Repositories struct {
	Record *Record.Repository
}

func NewRepositories(db *database.Database) *Repositories {
	return &Repositories{
		Record: &Record.Repository{DB: db},
	}
}
