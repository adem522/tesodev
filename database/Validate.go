package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Validate(id *string, col *mongo.Collection) bool {
	return col.FindOne(context.TODO(), bson.M{"_id": id}) != nil
}
