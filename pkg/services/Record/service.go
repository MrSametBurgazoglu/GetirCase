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
	var startDateISO time.Time
	var endDateISO time.Time

	ctx := context.TODO()

	layout := "2006-01-02T15:04:05Z"

	startDateISO, err = time.Parse(layout, startDate)
	if err != nil {
		return nil, err
	}

	endDateISO, err = time.Parse(layout, endDate)
	if err != nil {
		return nil, err
	}

	return s.Repository.FilterRecordsByDateRangeAndCountRange(ctx, startDateISO, endDateISO, minCount, maxCount)

}
