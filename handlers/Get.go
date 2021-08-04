package handlers

import (
	"deneme-structHandler/database"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func (col *Collect) Get(c echo.Context) error {
	data := struct {
		Id         string `bson:"_id" json:"_id"`
		CustomerId string `bson:"customerId" json:"customerId"`
	}{}
	data2 := []bson.M{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if col.Name == "order" && data.CustomerId != "" {
		data2, err = database.Get("customerId", data.CustomerId, col.Col)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, data2)
	}
	data2, err = database.Get("_id", data.Id, col.Col)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}
