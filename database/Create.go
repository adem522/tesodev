package database

import (
	"context"
	"reflect"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(request map[string]interface{}, db *mongo.Database, name string) (result *mongo.InsertOneResult, err error) {
	data, err := determine(request, db, name)
	if err != nil {
		return nil, fmt.Errorf("error from database/create %w", err)
	}
	fmt.Println("request", request)
	result, err = db.Collection(name).InsertOne(context.TODO(), data)
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
		{request: "address"},
		{request: "product"},
		{request: "customer"},
		{request: "order"},
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
	fmt.Println("içerdeyim")
	switch name {
	case "customer", "order":
		{
			fmt.Println("name", name)
			if name == "order" {
				request["product"], err = returnData(db.Collection("product"), request["productId"].(string))
				if err != nil {
					return nil, fmt.Errorf(" not found product for productId %w", err)
				}
				delete(request, "productId") //delete productId in request
			}
			request["address"], err = returnData(db.Collection("address"), request["AddressId"].(string))
			if err != nil {
				return nil, fmt.Errorf(" not found address for addressId %w", err)
			}
			delete(request, "addressId")
		}
	}
	fmt.Println("çıktım")
	return request, nil
}

func toInt(request map[string]interface{}, name string) {
	if reflect.TypeOf(request[name]).Key().Kind() == reflect.Float64 {
		request[name] = int(request[name].(float64)) //request is a pointer
	} else if reflect.TypeOf(request[name]).Key().Kind() == reflect.Float32 {
		request[name] = int(request[name].(float32)) //request is a pointer
	}
}

func returnData(collection *mongo.Collection, id string) (temp map[string]interface{}, err error) {
	data, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf(" find error %w", err)
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": data}).Decode(&temp)
	if err != nil {
		return nil, fmt.Errorf(" find error %w", err)
	}
	return temp, nil
}
