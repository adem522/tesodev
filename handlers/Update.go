package handlers

import (
	"deneme-structHandler/database"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

func (col *Collect) Update(c echo.Context) error {
	data := bson.M{}
	err := c.Bind(&data)
	fmt.Println(data["_id"])
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	err = database.Update(data, col.Col)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	return c.JSON(http.StatusOK, true)
}
