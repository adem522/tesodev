package database

import (
	"deneme-structHandler/database"
	"deneme-structHandler/handlers"
	"testing"
)

func TestDeleteEmptyArg(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	address := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "Address",
	}
	err := database.Delete(nil, address.Col)
	if err == nil {
		t.Errorf("expected error, received err %v", err)
	}
}
