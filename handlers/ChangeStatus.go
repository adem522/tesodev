package handlers

import (
	"net/http"

	"github.com/adem522/tesodev/database"
	"github.com/labstack/echo"
)

func ChangeStatus(c echo.Context) error {
	data := struct {
		Id     string `bson:"_id,omitempty" json:"_id"`
		Status string `bson:"status,omitempty"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	err = database.ChangeStatus(data.Id, data.Status, "Order")
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	return c.JSON(http.StatusOK, true)
}
