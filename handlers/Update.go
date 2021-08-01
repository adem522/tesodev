package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adem522/tesodev/database"
	"github.com/adem522/tesodev/models"
	"github.com/labstack/echo"
)

func UpdateOrder(c echo.Context) error {
	data := models.Order{
		UpdatedAt: time.Now(),
	}
	err := c.Bind(&data)
	if err != nil {
		fmt.Println("bind hatası")
		return c.JSON(http.StatusBadRequest, err)
	}
	err = database.UpdateOrder(&data, "Order")
	if err != nil {
		fmt.Println("update hatası")
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, data.Id.String())
}
