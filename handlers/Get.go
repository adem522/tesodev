package handlers

import (
	"deneme-structHandler/database"
	"net/http"

	"github.com/labstack/echo"
)

func (col *Order) Get(c echo.Context) error {
	data := struct {
		Id string `bson:"_id" json:"_id"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.Get("_id", data.Id, col.Collection)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}

func (col *Customer) Get(c echo.Context) error {
	data := struct {
		Id string `bson:"_id" json:"_id"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.Get("_id", data.Id, col.Collection)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}

func (col *Order) GetCustomerOrder(c echo.Context) error {
	data := struct {
		CustomerId string `bson:"customerId" json:"customerId"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.Get("customerId", data.CustomerId, col.Collection)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}
