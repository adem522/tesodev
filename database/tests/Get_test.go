package database

import (
	"tesodev/database"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestGetWithEmptyArg(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)

	data, err := database.Get(bson.M{}, client.Database("tesodev").Collection("order"))
	if err != nil {
		t.Errorf("expected nil, received err %v", err)
	}
	if data == nil {
		t.Errorf("expected not nil, received nil %v", data)
	}
}
func TestGetWithArg(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	order := exampleOrder()
	data, err := database.Get(bson.M{"_id": order.InsertedID}, client.Database("tesodev").Collection("order"))
	if err != nil {
		t.Errorf("expected nil, received err %v", err)
	}
	if data == nil {
		t.Errorf("expected not nil, received nil %v", data)
	}
	collectionName := "order"
	inserted := order.InsertedID.(string)
	database.Delete(&inserted, &collectionName, client.Database("tesodev"))
}
