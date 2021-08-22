package database

import (
	"context"
	"reflect"
	"time"

	"fmt"
	"tesodev/models"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(request map[string]interface{}, col *mongo.Collection, name string) (result *mongo.InsertOneResult, err error) {
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

func CreateCollections() (err error) {
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
		err = client.Database("tesodev").CreateCollection(context.TODO(), table.request, table.validator)
		if err != nil {
			return fmt.Errorf("error from database/createCollections - %w", err)
		}
		fmt.Println(table.request, "Collection Successfully created")
	}
	return nil
}

func determine(request map[string]interface{}, name string) (bson.M, error) {
	if name != "" {
		request["_id"] = uuid.NewV4().String() //every collection need id
		if name == "Address" {
			toInt(request, "cityCode") //if came double its return to int
		} else if name == "Customer" || name == "Order" {
			request["createdAt"] = time.Now().Add(3 * time.Hour)
			request["updatedAt"] = time.Now().Add(3 * time.Hour)
			if name == "Order" {
				toInt(request, "quantity") //if came double its return to int
			}
			address := request["address"].(map[string]interface{}) //address is a interface
			toInt(address, "cityCode")                             //if came double its return to int
		} else if name == "Product" {
			//product just need id
		} else {
			return nil, fmt.Errorf(" collection name not found") //if not nil and not at all these
		}
	} else {
		return nil, fmt.Errorf(" collection name must") //if name nil
	}
	return request, nil
}

func toInt(request map[string]interface{}, name string) {
	if reflect.TypeOf(request[name]) == reflect.TypeOf(1.0) {
		request[name] = int(request[name].(float64))
	}
}
