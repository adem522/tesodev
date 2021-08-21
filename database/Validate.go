package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Validate(id *string, col *mongo.Collection) bool {
	data := bson.M{}
	err := col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
	return err == nil
}
