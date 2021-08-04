package database

import (
	"context"
	"deneme-structHandler/models"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(data bson.M, col *mongo.Collection, name string) (result *mongo.InsertOneResult, err error) {
	if name == "address" {
		result, err = col.InsertOne(context.TODO(), data)
	}
	if name == "product" {
		data["_id"] = uuid.NewV4().String()
		result, err = col.InsertOne(context.TODO(), data)
	}
	if name == "customer" {
		data["_id"] = uuid.NewV4().String()
		data["createdAt"] = time.Now().Add(3 * time.Hour)
		data["updatedAt"] = time.Now().Add(3 * time.Hour)
		result, err = col.InsertOne(context.TODO(), data)
	}
	if name == "order" {
		data["_id"] = uuid.NewV4().String()
		data["createdAt"] = time.Now().Add(3 * time.Hour)
		data["updatedAt"] = time.Now().Add(3 * time.Hour)
		result, err = col.InsertOne(context.TODO(), data)
	}
	if err != nil {
		return nil, err
	}
	return result, nil
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
