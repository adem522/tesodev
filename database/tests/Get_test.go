package database

import (
	"tesodev/database"
	"tesodev/handlers"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestGet(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	address := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "Address",
	}
	data, err := database.Get(bson.M{"_id": ""}, address.Col)
	if err != nil {
		t.Errorf("expected nil, received err %v", err)
	}
	if data == nil {
		t.Errorf("expected not nil, received nil %v", data)
	}
}
