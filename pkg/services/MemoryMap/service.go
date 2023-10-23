package MemoryMap

import (
	"errors"
	"getir_case/pkg/repository/MemoryMap"
)

type ServiceInterface interface {
	GetValueByKey(key string) (string, error)
	SetValueByKey(key, value string)
}

type Service struct {
	Repository *MemoryMap.Repository
}

func (s *Service) GetValueByKey(key string) (string, error) {
	val, err := s.Repository.GetValueByKey(key)
	if err != nil {
		return "", err
	}
	value, ok := val.(string)
	if !ok {
		return "", errors.New("value cannot convert to string")
	}
	return value, nil

}

func (s *Service) SetValueByKey(key, value string) {
	s.Repository.SetValueByKey(key, value)
}
