package database

import (
	"deneme-structHandler/database"
	"deneme-structHandler/handlers"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
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
	data := map[string]interface{}{}
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
	data := map[string]interface{}{
		"addressLine": "example addressline",
		"city":        "example cit",
		"country":     "example country",
		"cityCode":    1,
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
	data := map[string]interface{}{
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
	data := map[string]interface{}{
		"_id":   uuid.NewV4().String(),
		"name":  "example customer",
		"email": "example@hotmail.com",
		"address": map[string]interface{}{
			"addressLine": "",
			"city":        "example city",
			"country":     "example country",
			"cityCode":    1,
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

	data := map[string]interface{}{
		"_id":        uuid.NewV4().String(),
		"customerId": uuid.NewV4().String(),
		"quantity":   1,
		"price":      2.4,
		"status":     "available",
		"address": map[string]interface{}{
			"addressLine": "",
			"city":        "example city",
			"country":     "example country",
			"cityCode":    1,
		},
		"product": map[string]interface{}{
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
