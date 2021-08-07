package database

import (
	"context"
	"reflect"
	"time"

	"deneme-structHandler/models"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(request map[string]interface{}, col *mongo.Collection, name string) (result *mongo.InsertOneResult, err error) {
	request, err = determine(request, name)
	fmt.Println("create request", request)
	if err != nil {
		return nil, fmt.Errorf("error from database/create %w", err)
	}
	result, err = col.InsertOne(context.TODO(), request)
	if err != nil {
		fmt.Println("create database error", err)
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

func determine(request map[string]interface{}, name string) (bson.M, error) {
	fmt.Println("determine request", request)
	if name == "" {
		return nil, fmt.Errorf(" collection name must")
	} else if name == "Address" {
		request["_id"] = uuid.NewV4().String()
		toInt(request, "cityCode")
	} else if name == "Product" {
		request["_id"] = uuid.NewV4().String()
	} else if name == "Customer" || name == "Order" {
		request["_id"] = uuid.NewV4().String()
		request["createdAt"] = time.Now().Add(3 * time.Hour)
		request["updatedAt"] = time.Now().Add(3 * time.Hour)
		toInt(request, "quantity")
		address := request["address"].(map[string]interface{})
		toInt(address, "cityCode")
	} else {
		return nil, fmt.Errorf(" collection name not found")
	}
	return request, nil
}

func toInt(request map[string]interface{}, name string) {
	if reflect.TypeOf(request[name]) == reflect.TypeOf(1.0) {
		request[name] = int(request[name].(float64))
	}
}
