package models

import (
	"time"
)

type Customer struct {
	Id        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	AddressId string    `bson:"addressId"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
