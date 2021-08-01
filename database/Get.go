package database

import (
	"fmt"

	"github.com/adem522/tesodev/models"
	uuid "github.com/satori/go.uuid"
)

func GetOrder(id *uuid.UUID, collectionName string) (models.Order, error) {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)

	deneme := models.Order{}
	collection := client.Database("tesodev").Collection(collectionName)
	err := collection.FindOne(ctx, id).Decode(&deneme)
	if err != nil {
		fmt.Println("update hatası")
		return deneme, err
	}
	return deneme, err
}
