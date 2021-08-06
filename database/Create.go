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

func Create(request bson.M, col *mongo.Collection, name string) (result *mongo.InsertOneResult, err error) {
	request, err = determine(request, name)
	if err != nil {
		return nil, fmt.Errorf("error from database/create %w", err)
	}
	result, err = col.InsertOne(context.TODO(), request)
	if err != nil {
		return nil, fmt.Errorf("error from database/create %w", err)
	}
	return result, nil
}

func CreateCollections() error {
	client := Connect()
	defer Close(client)
	client.Database("tesodev").Drop(context.TODO()) // if exist

	tables := []struct {
		request   string
		validator *options.CreateCollectionOptions
	}{
		{request: "Address", validator: options.CreateCollection().SetValidator(models.AddressValidator)},
		{request: "Product", validator: options.CreateCollection().SetValidator(models.ProductValidator)},
		{request: "Customer", validator: options.CreateCollection().SetValidator(models.CustomerValidator)},
		{request: "Order", validator: options.CreateCollection().SetValidator(models.OrderValidator)},
	}

	for _, table := range tables {
		err := client.Database("tesodev").CreateCollection(context.TODO(), table.request, table.validator)
		if err != nil {
			return fmt.Errorf("error from database/createCollections - %w", err)
		}
		fmt.Println(table.request, "Collection Successfully created")
	}
	return nil
}

func determine(request bson.M, name string) (bson.M, error) {
	if name == "" {
		return nil, fmt.Errorf(" collection name must")
	} else if name == "Address" {
		request["cityCode"] = toInt(request, "cityCode")
	} else if name == "Product" {
		request["_id"] = uuid.NewV4().String()
	} else if name == "Customer" || name == "Order" {
		address := request["address"].(map[string]interface{})
		request["address"] = bson.M{
			"city":     address["city"],
			"country":  address["country"],
			"cityCode": toInt(address, "cityCode"),
		}
		request["_id"] = uuid.NewV4().String()
		if name == "Order" {
			request["quantity"] = toInt(request, "quantity")
		}
		request["createdAt"] = time.Now().Add(3 * time.Hour)
		request["updatedAt"] = time.Now().Add(3 * time.Hour)
	} else {
		return nil, fmt.Errorf(" collection name not found")
	}
	return request, nil
}

func toInt(request bson.M, name string) int {
	return int(request[name].(float64)) //postman send double but database validator is int
}
