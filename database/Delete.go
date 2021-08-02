package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Delete(id *string, collectionName string) error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)

	collection := client.Database("tesodev").Collection(collectionName)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println(result.DeletedCount)
	if err != nil || result.DeletedCount == 0 {
		return fmt.Errorf("deleted count 0 or %w ", err)
	}
	return nil
}
