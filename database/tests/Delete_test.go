package database

import (
	"tesodev/database"
	"testing"
)

func TestDeleteEmptyArg(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	collectionName := ""
	if database.Delete(nil, &collectionName, client.Database("tesodev")) {
		t.Errorf("expected error, received nil")
	}
}
func TestDeleteWithArg(t *testing.T) {
	client := database.Connect()
	defer database.Close(client)
	data := map[string]interface{}{
		"addressLine": "example addressline",
		"city":        "example city",
		"country":     "example country",
		"cityCode":    1,
	}
	collectionName := "address"
	result2 := ""
	result, _ := database.Create(data, client.Database("tesodev"), collectionName)
	result2 = result.InsertedID.(string)
	if !database.Delete(&result2, &collectionName, client.Database("tesodev")) { //if delete it return true
		t.Errorf("expected nil, received err")
	}
}
