package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(id, collectionName *string, col *mongo.Database) bool {
	result, err := col.Collection(*collectionName).DeleteOne(context.TODO(), bson.M{"_id": *id})
	return err == nil && result.DeletedCount != 0
}
