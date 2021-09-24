package models

import (
	"time"
)

type Order struct {
	CreatedAt  time.Time `bson:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
	Id         string    `bson:"_id"`
	CustomerId string    `bson:"customerId"`
	AddressId  string    `bson:"addressId"`
	ProductId  string    `bson:"productId"`
	Status     string    `bson:"status"`
	Price      float64   `bson:"price"`
	Quantity   int       `bson:"quantity"`
}
