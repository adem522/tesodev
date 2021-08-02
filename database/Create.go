package database

import (
	"fmt"

	"github.com/adem522/tesodev/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(v interface{}, collectionName string) (interface{}, error) {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}
	collection := client.Database("tesodev").Collection(collectionName)
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func CreateCollections() error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	client.Database("tesodev").Drop(ctx) // if exist

	tables := []struct {
		data      string
		validator *options.CreateCollectionOptions
	}{
		{data: "Address", validator: options.CreateCollection().SetValidator(models.AddressValidator)},
		{data: "Product", validator: options.CreateCollection().SetValidator(models.ProductValidator)},
		{data: "Customer", validator: options.CreateCollection().SetValidator(models.CustomerValidator)},
		{data: "Order", validator: options.CreateCollection().SetValidator(models.OrderValidator)},
	}

	for _, table := range tables {
		err := client.Database("tesodev").CreateCollection(ctx, table.data, table.validator)
		if err != nil {
			return err
		}
		fmt.Println(table.data, "Collection Successfully created")
	}
	return nil
}
