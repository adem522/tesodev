package handlers

import "go.mongodb.org/mongo-driver/mongo"

type Database struct {
	Database *mongo.Database
}
