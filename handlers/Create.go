package handlers

import (
	"tesodev/database"

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
	var data map[string]interface{}
	err = c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := database.Create(data, col.Col, col.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	insertedId = result.InsertedID.(string)
	return c.JSON(http.StatusOK, insertedId)
}
