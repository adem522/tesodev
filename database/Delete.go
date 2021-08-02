package database

import "go.mongodb.org/mongo-driver/bson"

func Delete(id *string, collectionName string) error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)

	collection := client.Database("tesodev").Collection(collectionName)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil

}
