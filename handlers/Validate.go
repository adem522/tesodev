package handlers

import (
	"net/http"

	"github.com/adem522/tesodev/database"
	"github.com/labstack/echo"
)

func Validate(c echo.Context) error {
	data := struct {
		Id string `bson:"_id" json:"_id"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	data2 := database.Validate(&data.Id, "Customer")
	if data2 {
		return c.JSON(http.StatusOK, data2)
	}
	return c.JSON(http.StatusBadRequest, data2)
}