package database

import (
	"context"
	"deneme-structHandler/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(data interface{}, col *mongo.Collection) (interface{}, error) {
	data, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}
	result, err := col.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func CreateCollections() error {
	client := Connect()
	defer Close(client)
	client.Database("tesodev").Drop(context.TODO()) // if exist

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
		err := client.Database("tesodev").CreateCollection(context.TODO(), table.data, table.validator)
		if err != nil {
			return err
		}
		fmt.Println(table.data, "Collection Successfully created")
	}
	return nil
}
