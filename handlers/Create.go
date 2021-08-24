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

func (db *Collect) Create(c echo.Context) (err error) {
	insertedId := ""
	var data map[string]interface{}
	err = c.Bind(&data)
	request := c.Request().URL.Path[1:] //it came with /
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := database.Create(data, db.Database, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	insertedId = result.InsertedID.(string)
	return c.JSON(http.StatusOK, insertedId)
}
