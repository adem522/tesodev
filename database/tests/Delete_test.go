package database

import (
	"tesodev/database"
	"tesodev/handlers"
	"testing"
)

func TestDeleteEmptyArg(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	address := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "Address",
	}
	if database.Delete(nil, address.Col) {
		t.Errorf("expected error, received nil")
	}
}
func TestDeleteWithArg(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	address := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "Address",
	}
	data := map[string]interface{}{
		"addressLine": "example addressline",
		"city":        "example cit",
		"country":     "example country",
		"cityCode":    1,
	}
	result2 := ""
	result, _ := database.Create(data, address.Col, address.Col.Name())
	result2 = result.InsertedID.(string)
	if database.Delete(&result2, address.Col) {
		t.Errorf("expected nil, received err")
	}
}
