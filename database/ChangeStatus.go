package database

import (
	"go.mongodb.org/mongo-driver/bson"
)

func ChangeStatus(id, status, collectionName string) error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)

	collection := client.Database("tesodev").Collection(collectionName)
	_, err := collection.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "status", Value: status},
			}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
