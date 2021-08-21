package handlers

import (
	"net/http"
	"tesodev/database"

	"github.com/labstack/echo"
)

func (col *Collect) Delete(c echo.Context) error {
	data := struct {
		Id string `bson:"_id" json:"_id"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	err = database.Delete(&data.Id, col.Col)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	return c.JSON(http.StatusOK, true)
}
