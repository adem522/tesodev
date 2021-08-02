package handlers

import (
	"fmt"
	"net/http"

	"github.com/adem522/tesodev/database"
	"github.com/labstack/echo"
)

func DeleteOrder(c echo.Context) error {
	data := struct {
		Id string `bson:"_id" json:"_id"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		fmt.Println("bind error")
		return c.JSON(http.StatusBadRequest, false)
	}
	err = database.Delete(&data.Id, "Order")
	if err != nil {
		fmt.Println("update error")
		return c.JSON(http.StatusBadRequest, false)
	}
	return c.JSON(http.StatusOK, true)
}

func DeleteCustomer(c echo.Context) error {
	data := struct {
		Id string `bson:"_id" json:"_id"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		fmt.Println("bind error")
		return c.JSON(http.StatusBadRequest, false)
	}
	err = database.Delete(&data.Id, "Customer")
	if err != nil {
		fmt.Println("update error")
		return c.JSON(http.StatusBadRequest, false)
	}
	return c.JSON(http.StatusOK, true)
}
