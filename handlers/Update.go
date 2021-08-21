package handlers

import (
	"fmt"
	"net/http"
	"tesodev/database"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func (col *Collect) Update(c echo.Context) error {
	data := bson.M{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	err = database.Update(data, col.Col)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, false)
	}
	return c.JSON(http.StatusOK, true)
}
