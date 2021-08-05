package database

import (
	"deneme-structHandler/database"
	"deneme-structHandler/handlers"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateCollection(t *testing.T) {
	err := database.CreateCollections()
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}

func TestCreateNil(t *testing.T) {
	_, err := database.Create(nil, &mongo.Collection{}, "deneme")
	if err == nil {
		t.Errorf("Expected error, received %v", err)
	}
}

/*
func TestCreateAddressNil(t *testing.T) {
	data := struct{}{}
	err := database.Create(data, "Address")
	if err == nil {
		t.Errorf("Expected error, received %v", err)
	}
}*/

func TestCreateAddress(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	address := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "address",
	}
	data := bson.M{
		"addressLine": "example addressline",
		"city":        "example cit",
		"country":     "example country",
		"cityCode":    1,
	}
	_, err := database.Create(data, address.Col, address.Name)
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}

/*
func TestCreateProduct(t *testing.T) {
	data := struct {
		Id       string `bson:"_id,omitempty" json:"_id"`
		ImageUrl string `bson:"imageUrl,omitempty"`
		Name     string `bson:"name,omitempty"`
	}{
		Id:       uuid.NewV4().String(),
		ImageUrl: "example imageUrl",
		Name:     "example name",
	}
	err := database.Create(data, "Product")
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}
func TestCreateProductNil(t *testing.T) {
	data := struct{}{}
	err := database.Create(data, "Product")
	if err == nil {
		t.Errorf("Expected error, received %v", err)
	}
}
func TestCreateCustomer(t *testing.T) {
	data := models.Customer{
		Id:    uuid.NewV4().String(),
		Name:  "example customer",
		Email: "example@hotmail.com",
		Address: models.Address{
			AddressLine: "",
			City:        "example city",
			Country:     "example country",
			CityCode:    1,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := database.Create(data, "Customer")
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}
func TestCreateCustomerNil(t *testing.T) {
	data := struct{}{}
	err := database.Create(data, "Customer")
	if err == nil {
		t.Errorf("Expected error, received %v", err)
	}
}
func TestCreateOrder(t *testing.T) {
	data := models.Order{
		Id:         uuid.NewV4().String(),
		CustomerId: uuid.NewV4().String(),
		Quantity:   1,
		Price:      2.4,
		Status:     "available",
		Address: models.Address{
			AddressLine: "",
			City:        "example city",
			Country:     "example country",
			CityCode:    1},
		Product: models.Product{
			Id:       uuid.NewV4().String(),
			ImageUrl: "example imageUrl",
			Name:     "example name",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := database.Create(data, "Order")
	if err != nil {
		t.Errorf("Expected nil, received %v", err)
	}
}
func TestCreateOrderNil(t *testing.T) {
	data := struct{}{}
	err := database.Create(data, "Order")
	if err == nil {
		t.Errorf("Expected error, received %v", err)
	}
}
*/
