package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(keyWord, id string, col *mongo.Collection) ([]bson.M, error) {
	filterCursor := new(mongo.Cursor)
	var err error
	if id != "" {
		filterCursor, err = col.Find(context.TODO(), bson.M{keyWord: id})
		if err != nil {
			return nil, err
		}
	} else {
		filterCursor, err = col.Find(context.TODO(), bson.M{})
		if err != nil {
			return nil, err
		}
	}
	data := []bson.M{}
	if err = filterCursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return data, nil
}
