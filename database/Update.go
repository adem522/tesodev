package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(data bson.M, col *mongo.Collection) bool {
	data["updatedAt"] = time.Now().Add(3 * time.Hour)
	result, err := col.UpdateOne(context.TODO(),
		bson.M{"_id": data["_id"]},
		bson.D{{Key: "$set", Value: data}},
	)
	return err == nil && result.ModifiedCount != 0
}
