package handlers

import (
	"net/http"
	"time"

	"github.com/adem522/tesodev/database"
	"github.com/adem522/tesodev/models"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func CreateAddress(c echo.Context) error {
	data := models.Address{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = database.Create(&data, "Address")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "ok moruk")
}

func CreateProduct(c echo.Context) error {
	data := models.Product{
		Id: uuid.NewV4().String(),
	}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = database.Create(&data, "Product")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data.Id)
}

func CreateOrder(c echo.Context) error {
	data := models.Order{
		Id:        uuid.NewV4().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = database.CreateOrder(&data, "Order")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data.Id)
}

func CreateCustomer(c echo.Context) error {
	data := models.Customer{
		Id:        uuid.NewV4().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = database.Create(&data, "Customer")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data.Id)
}
