package database

import (
	"fmt"
	"time"

	"github.com/adem522/tesodev/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateOrder(v *models.Order, collectionName string) error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	collection := client.Database("tesodev").Collection(collectionName)
	_, err := collection.UpdateOne(ctx,
		bson.M{"_id": v.Id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "updatedAt", Value: time.Now().Add(3 * time.Hour)},
				{Key: "address", Value: v.Address},
				{Key: "product", Value: v.Product},
				{Key: "quantity", Value: v.Quantity},
				{Key: "price", Value: v.Price},
				{Key: "status", Value: v.Status},
				{Key: "customerId", Value: v.CustomerId},
				{Key: "createdAt", Value: v.CreatedAt},
			}},
		},
	)
	if err != nil {
		fmt.Println("update error", err)
		return err
	}
	return nil
}

func UpdateCustomer(v *models.Customer, collectionName string) error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	collection := client.Database("tesodev").Collection(collectionName)
	_, err := collection.UpdateOne(ctx,
		bson.M{"_id": v.Id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: v.Name},
				{Key: "email", Value: v.Email},
				{Key: "address", Value: v.Address},
				{Key: "updatedAt", Value: time.Now().Add(3 * time.Hour)},
				{Key: "createdAt", Value: v.CreatedAt},
			}},
		},
	)
	if err != nil {
		fmt.Println("update error", err)
		return err
	}
	return nil
}
