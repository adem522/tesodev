package handlers

import (
	"net/http"
	"tesodev/database"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func (col *Collect) Validate(c echo.Context) error {
	data := struct {
		Id string `bson:"_id" json:"_id"`
	}{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	if database.Validate(bson.M{"_id": data.Id}, col.Col) {
		return c.JSON(http.StatusOK, true)
	}
	return c.JSON(http.StatusBadRequest, false)
}
