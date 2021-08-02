package handlers

import (
	"net/http"

	"github.com/adem522/tesodev/database"
	"github.com/labstack/echo"
)

type temp struct {
	Id string `bson:"_id,omitempty" json:"_id"`
}

func GetOrder(c echo.Context) error {
	data := new(temp)
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.Get(data.Id, "Order")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}

func GetCustomer(c echo.Context) error {
	data := new(temp)
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.Get(data.Id, "Customer")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}

type temp2 struct {
	CustomerId string `bson:"customerId,omitempty" json:"customerId"`
}

func GetCustomerOrder(c echo.Context) error {
	data := new(temp2)
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.GetCustomerOrder(data.CustomerId, "Order")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}
