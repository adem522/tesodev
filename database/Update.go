package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(data map[string]interface{}, col *mongo.Collection, name string) bool {
	data, err := updateDetermine(data, name)
	if err != nil {
		return false
	}
	result, err := col.UpdateOne(context.TODO(),
		bson.M{"_id": data["_id"]},
		bson.D{{Key: "$set", Value: data}},
	)
	return err == nil && result.ModifiedCount != 0
}

func updateDetermine(request map[string]interface{}, name string) (bson.M, error) {
	if name != "" {
		if name == "Address" {
			toInt(request, "cityCode") //if came double its return to int
		} else if name == "Customer" || name == "Order" {
			request["updatedAt"] = time.Now().Add(3 * time.Hour)
			if name == "Order" {
				toInt(request, "quantity") //if came double its return to int
			}
			address := request["address"].(map[string]interface{}) //address is a interface
			toInt(address, "cityCode")                             //if came double its return to int
		} else if name == "Product" {
			//product just need id
		} else {
			return nil, fmt.Errorf(" collection name not found") //if not nil and not at all these
		}
	} else {
		return nil, fmt.Errorf(" collection name must") //if name nil
	}
	return request, nil
}
