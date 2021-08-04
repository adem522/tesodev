package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(data bson.M, col *mongo.Collection) error {
	data["updatedAt"] = time.Now().Add(3 * time.Hour)
	_, err := col.UpdateOne(context.TODO(),
		bson.M{"_id": data["_id"]},
		bson.D{{Key: "$set", Value: data}},
	)

	if err != nil {
		return err
	}
	return nil
}
