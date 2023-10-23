package MemoryMap

import (
	"errors"
	"getir_case/drivers/database"
)

type Repository struct {
	DB *database.Database
}

func (r *Repository) GetValueByKey(key string) (any, error) {
	val, ok := r.DB.InMemoryDatabase.Get(key)
	if !ok {
		return "", errors.New("value not found by key:" + key)
	}
	return val, nil
}

func (r *Repository) SetValueByKey(key, value string) {
	r.DB.InMemoryDatabase.Set(key, value)
}
