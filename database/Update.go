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
	switch name {
	case "address":
		{
			toInt(request, "cityCode") //if came double its return to int
		}
	case "product":
	case "customer", "order":
		{
			request["updatedAt"] = time.Now().Add(3 * time.Hour)
			if name == "order" {
				toInt(request, "quantity") //if came double its return to int
			}
			if request["address"] != nil {
				toInt(request["address"].(map[string]interface{}), "cityCode") //if came double its return to int
			}
		}
	default:
		{
			return nil, fmt.Errorf(" collection name not found") //if name not found
		}
	}
	return request, nil
}
