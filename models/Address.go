package models

type Address struct {
	AddressLine string `bson:"addressLine"`
	City        string `bson:"city"`
	Country     string `bson:"country"`
	CityCode    int    `bson:"cityCode"`
}
