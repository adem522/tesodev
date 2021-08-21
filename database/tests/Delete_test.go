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
	err := database.Delete(nil, address.Col)
	if err == nil {
		t.Errorf("expected error, received nil %v", err)
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
	err := database.Delete(&result2, address.Col)
	if err != nil {
		t.Errorf("expected nil, received err %v", err)
	}
}
