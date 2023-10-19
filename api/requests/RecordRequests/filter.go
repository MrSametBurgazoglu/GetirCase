package RecordRequests

// FilterInput godoc
type FilterInput struct {
	StartDate string `json:"startDate" binding:"required" example:"2016-01-26"`
	EndDate   string `json:"endDate" binding:"required" example:"2018-02-02"`
	MinCount  int    `json:"minCount" binding:"required" example:"2700"`
	MaxCount  int    `json:"maxCount" binding:"required" example:"3000"`
}
