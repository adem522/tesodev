package database

import (
	"fmt"

	"github.com/adem522/tesodev/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateOrder(v *models.Order, collectionName string) error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)

	data, err := bson.Marshal(v)
	if err != nil {
		fmt.Println("marshal hatası")
		return err
	}
	collection := client.Database("tesodev").Collection(collectionName)
	_, err = collection.UpdateOne(ctx, v.Id.String(), data)
	if err != nil {
		fmt.Println("update hatası")
		return err
	}
	return nil
}
