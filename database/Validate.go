package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Validate(id *string, col *mongo.Collection) bool {
	data2 := bson.M{}
	err := col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data2)
	return err == nil
}
