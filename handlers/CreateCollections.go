package handlers

import (
	"net/http"

	"github.com/adem522/tesodev/database"
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
