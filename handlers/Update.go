package handlers

import (
	"deneme-structHandler/database"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func (col *Order) Update(c echo.Context) error {
	col.UpdatedAt = time.Now().Add(3 * time.Hour)
	err := c.Bind(&col)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	err = database.Update(&col.Id, col, col.Collection)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	return c.JSON(http.StatusOK, true)
}

func (col *Customer) Update(c echo.Context) error {
	err := c.Bind(&col)
	col.UpdatedAt = time.Now().Add(3 * time.Hour)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	err = database.Update(&col.Id, col, col.Collection)
	if err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}
	return c.JSON(http.StatusOK, true)
}
