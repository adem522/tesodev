package handlers

import (
	"net/http"

	"github.com/adem522/tesodev/database"
	"github.com/labstack/echo"
)

func GetOrder(c echo.Context) error {
	data := struct {
		Id string `bson:"_id" json:"_id"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.Get("_id", data.Id, "Order")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}

func GetCustomer(c echo.Context) error {
	data := struct {
		Id string `bson:"_id" json:"_id"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.Get("_id", data.Id, "Customer")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}

func GetCustomerOrder(c echo.Context) error {
	data := struct {
		CustomerId string `bson:"customerId" json:"customerId"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.Get("customerId", data.CustomerId, "Order")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}
