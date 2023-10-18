package Record

import (
	"context"
	"getir_case/drivers/database"
	"getir_case/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

type Repository struct {
	DB *database.Database
}

func (r *Repository) FilterRecordsByDateRangeAndCountRange(ctx context.Context, startDate, endDate time.Time, minCount, maxCount int) ([]*models.RecordFilterModel, error) {
	groupStage := bson.A{
		bson.D{
			{"$match", bson.D{
				{"createdAt", bson.D{
					{"$lte", endDate}, {"$gte", startDate}},
				},
			},
			},
			{Key: "$group", Value: bson.D{
				{"totalCount", bson.D{{"$sum", "$counts"}}},
				{"key", "$key"},
				{"createdAt", "$createdAt"},
			},
			},
			{Key: "$match", Value: bson.D{
				{"totalCount", bson.D{{"$lte", maxCount}, {"$gte", minCount}}},
			},
			},
		},
	}

	aggregate, err := r.DB.RecordCollection.Aggregate(ctx, groupStage)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var results []bson.M
	if err = aggregate.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	var resultModels []*models.RecordFilterModel
	for _, result := range results {
		model := new(models.RecordFilterModel)
		model.Key = result["key"].(string)
		model.CreatedAt = result["createdAt"].(string)
		model.TotalCount = result["totalCount"].(int)
		resultModels = append(resultModels, model)
	}
	return resultModels, nil
}
