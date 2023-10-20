package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RecordFilterModel struct {
	Key        string             `json:"key" bson:"key"`
	CreatedAt  primitive.DateTime `json:"createdAt" bson:"createdAt"`
	TotalCount int                `json:"totalCount" bson:"totalCount"`
}
