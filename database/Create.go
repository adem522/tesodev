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

func Create(request map[string]interface{}, db *mongo.Database, name string) (result *mongo.InsertOneResult, err error) {
	request, err = determine(request, db, name)
	if err != nil {
		return nil, fmt.Errorf("error from database/create %w", err)
	}
	result, err = db.Collection(name).InsertOne(context.TODO(), request)
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
		{request: "address", validator: options.CreateCollection().SetValidator(models.AddressValidator)},
		{request: "product", validator: options.CreateCollection().SetValidator(models.ProductValidator)},
		{request: "customer", validator: options.CreateCollection().SetValidator(models.CustomerValidator)},
		{request: "order", validator: options.CreateCollection().SetValidator(models.OrderValidator)},
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

func determine(request map[string]interface{}, db *mongo.Database, name string) (bson.M, error) {
	var err error
	request["_id"] = uuid.NewV4().String() //every collection need id
	switch name {
	case "address":
		{
			toInt(request, "cityCode") //if came double its return to int
		}
	case "product":
	case "customer", "order":
		{
			request["createdAt"] = time.Now().Add(3 * time.Hour)
			request["updatedAt"] = time.Now().Add(3 * time.Hour)
			if name == "order" {
				toInt(request, "quantity") //if came double its return to int
				request["product"], err = returnData(db.Collection("product"), request["productId"].(string))
				if err != nil {
					return nil, fmt.Errorf(" not found product for productId %w", err)
				}
				delete(request, "productId") //delete productId in request
			}
			request["address"], err = returnData(db.Collection("address"), request["addressId"].(string))
			if err != nil {
				return nil, fmt.Errorf(" not found address for addressId %w", err)
			}
			delete(request, "addressId")
		}
	default:
		{
			return nil, fmt.Errorf(" collection name not found") //if name not found
		}
	}
	return request, nil
}

func returnData(collection *mongo.Collection, id string) (temp map[string]interface{}, err error) {
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&temp)
	if err != nil {
		return nil, fmt.Errorf(" find error")
	}
	return temp, nil
}

func toInt(request map[string]interface{}, name string) {
	if reflect.TypeOf(request[name]) == reflect.TypeOf(1.0) {
		request[name] = int(request[name].(float64)) //request is a pointer
	}
}
