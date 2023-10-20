package RecordRequests

// FilterInput godoc
type FilterInput struct {
	StartDate string `json:"startDate" validate:"required" example:"2016-01-26"`
	EndDate   string `json:"endDate" validate:"required" example:"2018-02-02"`
	MinCount  int    `json:"minCount" validate:"required" example:"2700"`
	MaxCount  int    `json:"maxCount" validate:"required" example:"3000"`
}
