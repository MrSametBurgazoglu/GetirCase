package MemoryMap

import (
	"context"
	"errors"
	"getir_case/drivers/database"
)

type Repository struct {
	DB *database.Database
}

func (r *Repository) GetValueByKey(ctx context.Context, key string) (string, error) {
	val, err := r.DB.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", errors.New("value not found by key:" + key)
	}
	return val, nil
}

func (r *Repository) SetValueByKey(ctx context.Context, key, value string) error {
	err := r.DB.RedisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		return errors.New("couldn't set for key:" + key)
	}
	return nil
}
