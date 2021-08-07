package database

import (
	"deneme-structHandler/database"
	"deneme-structHandler/handlers"
	"testing"
)

func TestGet(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	address := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "Address",
	}
	data, err := database.Get("_id", "", address.Col)
	if err != nil {
		t.Errorf("expected nil, received err %v", err)
	}
	if data == nil {
		t.Errorf("expected not nil, received nil %v", data)
	}
}
