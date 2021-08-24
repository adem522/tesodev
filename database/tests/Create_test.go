package database

import (
	"tesodev/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
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
	data := map[string]interface{}{}
	_, err := database.Create(data, client.Database("tesodev"), "Address")
	if err == nil {
		t.Errorf("Expected error, received %v", err)
	}
}

func TestCreateAddress(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	data := map[string]interface{}{
		"addressLine": "example addressline",
		"city":        "example cit",
		"country":     "example country",
		"cityCode":    1,
	}
	address, err := database.Create(data, client.Database("tesodev"), "address")
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
	collectionName := "address"
	inserted := address.InsertedID.(string)
	database.Delete(&inserted, &collectionName, client.Database("tesodev"))
}

func TestCreateProduct(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	data := map[string]interface{}{
		"_id":      uuid.NewV4().String(),
		"imageUrl": "example imageUrl",
		"name":     "example name",
	}
	product, err := database.Create(data, client.Database("tesodev"), "product")
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
	collectionName := "address"
	inserted := product.InsertedID.(string)
	database.Delete(&inserted, &collectionName, client.Database("tesodev"))
}

func TestCreateCustomer(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	data := map[string]interface{}{
		"addressLine": "example addressline",
		"city":        "example cit",
		"country":     "example country",
		"cityCode":    1,
	}
	address, _ := database.Create(data, client.Database("tesodev"), "address")

	data2 := map[string]interface{}{
		"_id":       uuid.NewV4().String(),
		"name":      "example customer",
		"email":     "example@hotmail.com",
		"addressId": address.InsertedID,
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
	}
	customer, err := database.Create(data2, client.Database("tesodev"), "customer")
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
	//delet test items
	collectionName := "customer"
	inserted := customer.InsertedID.(string)
	database.Delete(&inserted, &collectionName, client.Database("tesodev"))
	collectionName = "address"
	inserted = address.InsertedID.(string)
	database.Delete(&inserted, &collectionName, client.Database("tesodev"))
}

func TestCreateOrder(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)

	data := map[string]interface{}{
		"addressLine": "example addressline",
		"city":        "example cit",
		"country":     "example country",
		"cityCode":    1,
	}
	address, _ := database.Create(data, client.Database("tesodev"), "address")
	data2 := map[string]interface{}{
		"_id":      uuid.NewV4().String(),
		"imageUrl": "example imageUrl",
		"name":     "example name",
	}
	product, _ := database.Create(data2, client.Database("tesodev"), "product")

	data3 := map[string]interface{}{
		"_id":        uuid.NewV4().String(),
		"customerId": uuid.NewV4().String(),
		"quantity":   1,
		"price":      2.4,
		"status":     "available",
		"addressId":  address.InsertedID,
		"productId":  product.InsertedID,
		"createdAt":  time.Now(),
		"updatedAt":  time.Now(),
	}
	order, err := database.Create(data3, client.Database("tesodev"), "order")
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
	collectionName := "customer"
	inserted := product.InsertedID.(string)
	database.Delete(&inserted, &collectionName, client.Database("tesodev"))
	collectionName = "address"
	inserted = address.InsertedID.(string)
	database.Delete(&inserted, &collectionName, client.Database("tesodev"))
	collectionName = "order"
	inserted = order.InsertedID.(string)
	database.Delete(&inserted, &collectionName, client.Database("tesodev"))
}

func exampleOrder() (order *mongo.InsertOneResult) {
	client := database.Connect()
	defer database.Close(client)

	data := map[string]interface{}{
		"addressLine": "example addressline",
		"city":        "example cit",
		"country":     "example country",
		"cityCode":    1,
	}
	address, _ := database.Create(data, client.Database("tesodev"), "address")
	data2 := map[string]interface{}{
		"_id":      uuid.NewV4().String(),
		"imageUrl": "example imageUrl",
		"name":     "example name",
	}
	product, _ := database.Create(data2, client.Database("tesodev"), "product")

	data3 := map[string]interface{}{
		"_id":        uuid.NewV4().String(),
		"customerId": uuid.NewV4().String(),
		"quantity":   1,
		"price":      2.4,
		"status":     "available",
		"addressId":  address.InsertedID,
		"productId":  product.InsertedID,
		"createdAt":  time.Now(),
		"updatedAt":  time.Now(),
	}
	order, _ = database.Create(data3, client.Database("tesodev"), "order")
	return order
}
