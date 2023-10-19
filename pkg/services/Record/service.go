package Record

import (
	"context"
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

	startDateTime, err := time.Parse(sourceLayout, startDate)
	if err != nil {
		return nil, err
	}

	endDateTime, err := time.Parse(sourceLayout, endDate)
	if err != nil {
		return nil, err
	}

	return s.Repository.FilterRecordsByDateRangeAndCountRange(ctx, startDateTime, endDateTime, minCount, maxCount)

}
