package handlers

import (
	"net/http"
	"tesodev/database"

	"github.com/labstack/echo"
)

func (col *Collect) Update(c echo.Context) error {
	data := col.define()
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	if database.Update(data.(map[string]interface{}), col.Col, col.Name) {
		return c.JSON(http.StatusOK, true)
	}
	return c.JSON(http.StatusBadRequest, false)
}
