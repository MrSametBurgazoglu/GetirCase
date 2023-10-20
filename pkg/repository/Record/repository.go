package Record

import (
	"context"
	"getir_case/drivers/database"
	"getir_case/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Repository struct {
	DB *database.Database
}

func (r *Repository) FilterRecordsByDateRangeAndCountRange(ctx context.Context, startDate, endDate time.Time, minCount, maxCount int) ([]*models.RecordFilterModel, error) {
	matchStage1 := bson.D{ //first match for range date time
		{"$match", bson.D{
			{"createdAt", bson.D{
				{"$lte", primitive.NewDateTimeFromTime(endDate)}, {"$gte", primitive.NewDateTimeFromTime(startDate)}},
			},
		},
		},
	}

	projectStage := bson.D{ //sum counts as totalCount
		{Key: "$project", Value: bson.D{
			{"totalCount", bson.D{{"$sum", "$counts"}}},
			{"_id", 1},
			{"key", 1},
			{"createdAt", 1},
		},
		},
	}

	matchStage2 := bson.D{ //second match stage for total count
		{Key: "$match", Value: bson.D{
			{"totalCount", bson.D{{"$lte", maxCount}, {"$gte", minCount}}},
		},
		},
	}

	mongoPipeLine := mongo.Pipeline{matchStage1, projectStage, matchStage2}

	aggregate, err := r.DB.RecordCollection.Aggregate(ctx, mongoPipeLine)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var results []*models.RecordFilterModel
	if err = aggregate.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results, nil
}
