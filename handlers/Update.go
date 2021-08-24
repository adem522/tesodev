package handlers

import (
	"net/http"
	"tesodev/database"

	"github.com/labstack/echo"
)

func (col *Collect) Update(c echo.Context) error {
	var data map[string]interface{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	name := c.Request().URL.Path[1:]
	if database.Update(data, col.Database.Collection(name), name) {
		return c.JSON(http.StatusOK, true)
	}
	return c.JSON(http.StatusBadRequest, false)
}
