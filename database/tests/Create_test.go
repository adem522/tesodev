package database

import (
	"deneme-structHandler/database"
	"deneme-structHandler/handlers"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateCollection(t *testing.T) {
	err := database.CreateCollections()
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}

func TestCreateNil(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	address := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "Address",
	}
	data := bson.M{}
	_, err := database.Create(data, address.Col, "Address")
	if err == nil {
		t.Errorf("Expected error, received %v", err)
	}
}

func TestCreateAddress(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	collection := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "Address",
	}
	data := bson.M{
		"addressLine": "example addressline",
		"city":        "example cit",
		"country":     "example country",
		"cityCode":    1.0,
	}
	_, err := database.Create(data, collection.Col, collection.Name)
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}

func TestCreateProduct(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	collection := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Product"),
		Name: "Product",
	}
	data := bson.M{
		"_id":      uuid.NewV4().String(),
		"imageUrl": "example imageUrl",
		"name":     "example name",
	}
	_, err := database.Create(data, collection.Col, collection.Name)
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}

func TestCreateCustomer(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	collection := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Customer"),
		Name: "Customer",
	}
	data := bson.M{
		"_id":   uuid.NewV4().String(),
		"name":  "example customer",
		"email": "example@hotmail.com",
		"address": bson.M{
			"addressLine": "",
			"city":        "example city",
			"country":     "example country",
			"cityCode":    1.0,
		},
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	}
	_, err := database.Create(data, collection.Col, collection.Name)
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}
func TestCreateOrder(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	collection := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Order"),
		Name: "Order",
	}
	data := bson.M{
		"_id":        uuid.NewV4().String(),
		"customerId": uuid.NewV4().String(),
		"quantity":   1.0,
		"price":      2.4,
		"status":     "available",
		"address": bson.M{
			"addressLine": "",
			"city":        "example city",
			"country":     "example country",
			"cityCode":    1.0,
		},
		"product": bson.M{
			"_id":      uuid.NewV4().String(),
			"imageUrl": "example imageUrl",
			"name":     "example name",
		},
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	}

	_, err := database.Create(data, collection.Col, collection.Name)
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}
