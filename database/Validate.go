package database

import (
	"go.mongodb.org/mongo-driver/bson"
)

func Validate(id *string, collectionName string) bool {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	data2 := bson.M{}
	collection := client.Database("tesodev").Collection(collectionName)
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&data2)
	return err == nil
}
