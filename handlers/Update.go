package handlers

import (
	"net/http"
	"tesodev/database"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func (col *Collect) Update(c echo.Context) error {
	data := bson.M{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	if database.Update(data, col.Col) {
		return c.JSON(http.StatusOK, true)
	}
	return c.JSON(http.StatusBadRequest, false)
}
