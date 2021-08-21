package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(filter bson.M, col *mongo.Collection) (data []bson.M, err error) {
	filterCursor := new(mongo.Cursor)
	if filter != nil {
		filterCursor, err = col.Find(context.TODO(), filter)
	} else {
		filterCursor, err = col.Find(context.TODO(), bson.M{})
	}
	if err != nil {
		return nil, err
	}
	if err = filterCursor.All(context.TODO(), &data); err != nil {
		return nil, err
	}
	return data, nil
}
