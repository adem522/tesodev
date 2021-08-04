package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(id *string, col *mongo.Collection) error {
	result, err := col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		return fmt.Errorf("deleted count 0 or %w ", err)
	}
	return nil
}
