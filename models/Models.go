package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Address struct {
	AddressLine string `bson:"addressLine"`
	City        string `bson:"city,omitempty"`
	Country     string `bson:"country,omitempty"`
	CityCode    int32  `bson:"cityCode,omitempty"`
}

type Customer struct {
	Id        uuid.UUID `bson:"_id,omitempty"`
	Name      string    `bson:"name,omitempty"`
	Email     string    `bson:"email,omitempty"`
	Address   Address
	CreatedAt time.Time `bson:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}

type Product struct {
	Id       uuid.UUID `bson:"_id,omitempty"`
	ImageUrl string    `bson:"imageUrl,omitempty"`
	Name     string    `bson:"name,omitempty"`
}

type Order struct {
	Id         uuid.UUID `bson:"_id,omitempty"`
	CustomerId uuid.UUID `bson:"customerId,omitempty"`
	Quantity   string    `bson:"quantity,omitempty"`
	Price      float64   `bson:"price,omitempty"`
	Status     string    `bson:"status,omitempty"`
	Address    Address
	Product    Product
	CreatedAt  time.Time `bson:"createdAt,omitempty"`
	UpdatedAt  time.Time `bson:"updatedAt,omitempty"`
}
