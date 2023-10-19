package MemoryMap

import (
	"context"
	"getir_case/drivers/database"
)

type Repository struct {
	DB *database.Database
}

func (r *Repository) GetValueByKey(ctx context.Context, key string) (string, error) {
	val, err := r.DB.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *Repository) SetValueByKey(ctx context.Context, key, value string) error {
	return r.DB.RedisClient.Set(ctx, key, value, 0).Err()
}
