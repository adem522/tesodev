package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(id *string, data interface{}, col *mongo.Collection) error {
	_, err := col.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.D{{Key: "$set", Value: data}},
	)
	if err != nil {
		fmt.Println("update error", err)
		return err
	}
	return nil
}
