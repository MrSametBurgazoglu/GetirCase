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
	var startDateISO string
	var endDateISO string

	ctx := context.TODO()

	sourceLayout := "2006-01-02"
	destlayout := "2006-01-02T15:04:05Z"

	startDateTime, err := time.Parse(sourceLayout, startDate)
	startDateISO = startDateTime.Format(destlayout)
	if err != nil {
		return nil, err
	}

	endDateTime, err := time.Parse(sourceLayout, startDate)
	endDateISO = endDateTime.Format(destlayout)
	if err != nil {
		return nil, err
	}

	return s.Repository.FilterRecordsByDateRangeAndCountRange(ctx, startDateISO, endDateISO, minCount, maxCount)

}
