package models

import "time"

type Address struct {
	Id          string `bson:"_id,omitempty"`
	AddressLine string `bson:"addressLine"`
	City        string `bson:"city,omitempty"`
	Country     string `bson:"country,omitempty"`
	CityCode    int    `bson:"cityCode,omitempty"`
}

type Customer struct {
	Id        string `bson:"_id,omitempty"`
	Name      string `bson:"name,omitempty"`
	Email     string `bson:"email,omitempty"`
	Address   Address
	CreatedAt time.Time `bson:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}

type Product struct {
	Id       string `bson:"_id,omitempty"`
	ImageUrl string `bson:"imageUrl,omitempty"`
	Name     string `bson:"name,omitempty"`
}

type Order struct {
	Id         string  `bson:"_id,omitempty"`
	CustomerId string  `bson:"customerId,omitempty"`
	Quantity   int     `bson:"quantity,omitempty"`
	Price      float64 `bson:"price,omitempty"`
	Status     string  `bson:"status,omitempty"`
	Address    Address
	Product    Product
	CreatedAt  time.Time `bson:"createdAt,omitempty"`
	UpdatedAt  time.Time `bson:"updatedAt,omitempty"`
}
