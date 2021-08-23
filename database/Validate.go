package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Validate(filter bson.M, col *mongo.Collection) bool {
	return col.FindOne(context.TODO(), filter).Err() == nil
}
