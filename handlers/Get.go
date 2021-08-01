package handlers

import (
	"fmt"
	"net/http"

	"github.com/adem522/tesodev/database"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

type gecici struct {
	Id uuid.UUID `bson:"_id"`
}

func GetOrder(c echo.Context) error {
	data := gecici{}
	err := c.Bind(&data)
	fmt.Println(data.Id)
	if err != nil {
		fmt.Println("bind hatası")
		return c.JSON(http.StatusBadRequest, err)
	}
	err = data.Id.UnmarshalBinary(data.Id[:])
	if err != nil {
		fmt.Println("unmarshal hatası")
		return c.JSON(http.StatusBadRequest, err)
	}
	data2, err := database.GetOrder(&data.Id, "Order")
	if err != nil {
		fmt.Println("update hatası")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data2)
}
