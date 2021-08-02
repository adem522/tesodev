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
	//v.UpdatedAt = time.Now()
	collection := client.Database("tesodev").Collection(collectionName)
	_, err := collection.UpdateOne(ctx,
		bson.M{"_id": v.Id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "updatedAt", Value: time.Now()},
			}},
		},
	)
	if err != nil {
		fmt.Println("update hatası", err)
		return err
	}
	return nil
}
