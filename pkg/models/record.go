package models

type RecordFilterModel struct {
	Key        string `json:"key" bson:"key"`
	CreatedAt  string `json:"createdAt" bson:"createdAt"`
	TotalCount int    `json:"totalCount" bson:"totalCount"`
}
