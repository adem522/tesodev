package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ChangeStatus(id, status *string, col *mongo.Collection) bool {
	_, err := col.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "status", Value: status},
			}},
		},
	)
	return err == nil
}
