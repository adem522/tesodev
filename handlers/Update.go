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
	if database.Update(data, col.Col, col.Name) {
		return c.JSON(http.StatusOK, true)
	}
	return c.JSON(http.StatusBadRequest, false)
}
