package handlers

import (
	"net/http"
	"tesodev/database"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func (col *Collect) Get(c echo.Context) error {
	temp := struct {
		Id         string `bson:"_id" json:"_id"`
		CustomerId string `bson:"customerId" json:"customerId"`
	}{}
	data := []bson.M{}
	err := c.Bind(&temp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if col.Name == "Order" && temp.CustomerId != "" { //		 //all customer order with costumerId in Order Collection
		data, err = database.Get(bson.M{"customerId": temp.CustomerId}, col.Col)
	} else if temp.Id != "" { //						 		// every data with _id in came collection
		data, err = database.Get(bson.M{"_id": temp.Id}, col.Col)
	} else { //										 			//every data in came collection
		data, err = database.Get(nil, col.Col)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data)
}
