package handlers

import "go.mongodb.org/mongo-driver/mongo"

type Collect struct {
	Col  *mongo.Collection
	Name string
}
