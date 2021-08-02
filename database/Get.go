package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(keyWord, id string, collectionName string) (interface{}, error) {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	collection := client.Database("tesodev").Collection(collectionName)
	filterCursor := new(mongo.Cursor)
	var err error
	if id != "" {
		filterCursor, err = collection.Find(ctx, bson.M{keyWord: id})
		if err != nil {
			return nil, err
		}
	} else {
		filterCursor, err = collection.Find(ctx, bson.M{})
		if err != nil {
			return nil, err
		}
	}
	data := []bson.M{}
	if err = filterCursor.All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}
