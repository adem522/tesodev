package handlers

import (
	"fmt"
	"net/http"
	"tesodev/database"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func (col *Database) Get(c echo.Context) error {
	temp := struct {
		Id         string `bson:"_id" json:"_id"`
		CustomerId string `bson:"customerId" json:"customerId"`
	}{}
	data := []bson.M{}
	err := c.Bind(&temp)
	name := c.Request().URL.Path[1:]
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println("request name", name)
	if name == "order" && temp.CustomerId != "" { //		 //all customer order with costumerId in Order Collection
		data, err = database.Get(bson.M{"customerId": temp.CustomerId}, col.Database.Collection(name))
	} else if temp.Id != "" { //						 		// every data with _id in came collection
		data, err = database.Get(bson.M{"_id": temp.Id}, col.Database.Collection(name))
	} else { //										 			//every data in came collection
		data, err = database.Get(nil, col.Database.Collection(name))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data)
}
