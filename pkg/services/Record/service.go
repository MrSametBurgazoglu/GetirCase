package Record

import (
	"context"
	"errors"
	"getir_case/pkg/models"
	"getir_case/pkg/repository/Record"
	"time"
)

type ServiceInterface interface {
	FilterRecords(startDate, endDate string, minCount, maxCount int) ([]*models.RecordFilterModel, error)
}

type Service struct {
	Repository *Record.Repository
}

func (s *Service) FilterRecords(startDate, endDate string, minCount, maxCount int) ([]*models.RecordFilterModel, error) {
	var err error

	ctx := context.TODO()

	sourceLayout := "2006-01-02"

	startDateTime, err := time.Parse(sourceLayout, startDate) // convert to go time
	if err != nil {
		return nil, errors.New("cannot parse startDate please use 2006-01-02 format")
	}

	endDateTime, err := time.Parse(sourceLayout, endDate) // convert to go time
	if err != nil {
		return nil, errors.New("cannot parse endDate please use 2006-01-02 format")
	}

	return s.Repository.FilterRecordsByDateRangeAndCountRange(ctx, startDateTime, endDateTime, minCount, maxCount)

}
