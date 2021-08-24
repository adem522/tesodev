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
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	collectionName := c.Request().URL.Path[1:]
	if database.Delete(&data.Id, &collectionName, col.Database) {
		return c.JSON(http.StatusOK, true)
	}
	return c.JSON(http.StatusBadRequest, false)
}
