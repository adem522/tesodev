package models

type Product struct {
	Id       string `bson:"_id"`
	ImageUrl string `bson:"imageUrl"`
	Name     string `bson:"name"`
}
