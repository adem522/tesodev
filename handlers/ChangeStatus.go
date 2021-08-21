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
	data2 := database.ChangeStatus(data.Id, data.Status, col.Col)
	if data2 {
		return c.JSON(http.StatusOK, data2)
	}
	return c.JSON(http.StatusBadRequest, data2)
}
