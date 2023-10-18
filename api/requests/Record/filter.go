package Record

type FilterInput struct {
	StartDate string `json:"startDate" binding:"required"`
	EndDate   string `json:"endDate" binding:"required"`
	MinCount  int    `json:"minCount" binding:"required"`
	MaxCount  int    `json:"maxCount" binding:"required"`
}
