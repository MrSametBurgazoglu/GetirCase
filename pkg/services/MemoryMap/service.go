package MemoryMap

import (
	"context"
	"getir_case/pkg/repository/MemoryMap"
)

type ServiceInterface interface {
	GetValueByKey(key string) (string, error)
	SetValueByKey(key, value string) error
}

type Service struct {
	Repository *MemoryMap.Repository
}

func (s *Service) GetValueByKey(key string) (string, error) {
	ctx := context.TODO()

	return s.Repository.GetValueByKey(ctx, key)

}

func (s *Service) SetValueByKey(key, value string) error {
	ctx := context.TODO()

	return s.Repository.SetValueByKey(ctx, key, value)
}
