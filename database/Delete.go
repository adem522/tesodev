package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(id *string, col *mongo.Collection) bool {
	result, err := col.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err == nil && result.DeletedCount != 0
}
