package handlers

import (
	"deneme-structHandler/database"
	"fmt"
	"time"

	"net/http"

	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func CreateCollections(c echo.Context) error {
	err := database.CreateCollections()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't created user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "All collection successfully created.")
}

func (col *Address) Create(c echo.Context) error {
	data := Address{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := database.Create(data, col.Collection)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}

func (col *Product) Create(c echo.Context) error {
	data := Product{
		Id: uuid.NewV4().String(),
	}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := database.Create(data, col.Collection)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}

func (col *Customer) Create(c echo.Context) error {
	data := Customer{
		Id:        uuid.NewV4().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := database.Create(data, col.Collection)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}

func (col *Order) Create(c echo.Context) error {
	data := Order{
		Id:        uuid.NewV4().String(),
		CreatedAt: time.Now().Add(3 * time.Hour),
		UpdatedAt: time.Now().Add(3 * time.Hour),
	}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := database.Create(data, col.Collection)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, result)
}
