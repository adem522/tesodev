package handlers

import (
	"net/http"
	"tesodev/database"

	"github.com/labstack/echo"
)

func (col *Collect) ChangeStatus(c echo.Context) error {
	data := struct {
		Id     string `bson:"_id,omitempty" json:"_id"`
		Status string `bson:"status,omitempty"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	if database.ChangeStatus(data.Id, data.Status, col.Col) {
		return c.JSON(http.StatusOK, true)
	}
	return c.JSON(http.StatusBadRequest, false)
}
