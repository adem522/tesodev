package handlers

import (
	"fmt"
	"net/http"

	"github.com/adem522/tesodev/database"
	"github.com/adem522/tesodev/models"
	"github.com/labstack/echo"
)

func UpdateOrder(c echo.Context) error {
	data := models.Order{}
	err := c.Bind(&data)
	if err != nil {
		fmt.Println("bind error")
		return c.JSON(http.StatusBadRequest, err)
	}
	err = database.UpdateOrder(&data, "Order")
	if err != nil {
		fmt.Println("update error")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data.Id)
}
func UpdateCustomer(c echo.Context) error {
	data := models.Customer{}
	err := c.Bind(&data)
	if err != nil {
		fmt.Println("bind error")
		return c.JSON(http.StatusBadRequest, err)
	}
	err = database.UpdateCustomer(&data, "Customer")
	if err != nil {
		fmt.Println("update error")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data.Id)
}
