package handlers

import (
	"tesodev/database"
	"tesodev/models"

	"net/http"

	"github.com/labstack/echo"
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

func (col *Collect) Create(c echo.Context) (err error) {
	insertedId := ""
	data := col.define()
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := database.Create(data.(map[string]interface{}), col.Col, col.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	insertedId = result.InsertedID.(string)
	return c.JSON(http.StatusOK, insertedId)

}

func (col *Collect) define() (data interface{}) {
	if col.Name == "Address" {
		return models.Address{}
	} else if col.Name == "Product" {
		return models.Product{}
	} else if col.Name == "Customer" {
		return models.Customer{}
	} else if col.Name == "Order" {
		return models.Order{}
	} else {
		return nil
	}
}
