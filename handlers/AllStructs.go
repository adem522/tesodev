package handlers

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Address struct {
	Collection  *mongo.Collection `bson:"collection"`
	AddressLine string            `bson:"addressLine"`
	City        string            `bson:"city,omitempty"`
	Country     string            `bson:"country,omitempty"`
	CityCode    int32             `bson:"cityCode,omitempty"`
}

type Product struct {
	Collection *mongo.Collection `bson:"collection"`
	Id         string            `bson:"_id,omitempty" json:"_id"`
	ImageUrl   string            `bson:"imageUrl,omitempty"`
	Name       string            `bson:"name,omitempty"`
}

type Customer struct {
	Collection *mongo.Collection `bson:"collection"`
	Id         string            `bson:"_id,omitempty" json:"_id"`
	Name       string            `bson:"name,omitempty"`
	Email      string            `bson:"email,omitempty"`
	CreatedAt  time.Time         `bson:"createdAt,omitempty"`
	UpdatedAt  time.Time         `bson:"updatedAt,omitempty"`
	Address    Address
}

type Order struct {
	Collection *mongo.Collection `bson:"collection"`
	Id         string            `bson:"_id,omitempty" json:"_id"`
	CustomerId string            `bson:"customerId,omitempty"`
	Quantity   int               `bson:"quantity,omitempty"`
	Price      float64           `bson:"price,omitempty"`
	Status     string            `bson:"status,omitempty"`
	CreatedAt  time.Time         `bson:"createdAt,omitempty"`
	UpdatedAt  time.Time         `bson:"updatedAt,omitempty"`
	Address    Address
	Product    Product
}
