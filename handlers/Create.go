package handlers

import (
	"fmt"
	"reflect"
	"tesodev/database"
	"tesodev/models"
	"time"

	"net/http"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func CreateCollections(c echo.Context) error {
	err := database.CreateCollections()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Can't created user because of ": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "All collection successfully created.")
}

func (db *Database) Create(c echo.Context) (err error) {
	request := c.Request().URL.Path[1:] //it came with /
	data, err := define(request, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	result, err := database.Create(data, db.Database, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, result.InsertedID)
}

func define(request string, c echo.Context) (map[string]interface{}, error) {
	switch request {
	case "address":
		{
			temp := models.Address{}
			if err := c.Bind(&temp); err != nil {
				return nil, fmt.Errorf("bind error")
			}
			if temp.AddressLine == "" {
				temp.AddressLine = "Empty addressline"
			}
			inter := structs.Map(temp)
			fmt.Println("temp", temp)
			v := reflect.ValueOf(temp)
			for i := 0; i < v.NumField(); i++ {
				if v.Field(i).Interface() == "" {
					return nil, fmt.Errorf("empty value not accepted")
				}
			}
			fmt.Println("v", v)
			return inter, nil
		} /*
			case "product":
				{
					temp := models.Product{}
					if err := c.Bind(&temp); err != nil {
						return nil, fmt.Errorf("bind error")
					}
					temp.Id = uuid.NewV4().String()
					v := reflect.ValueOf(temp)
					for i := 0; i < v.NumField(); i++ {
						if v.Field(i).Interface() == "" {
							return nil, fmt.Errorf("empty value not accepted")
						}
					}
					return temp, nil
				}*/
	case "customer":
		{
			temp := models.Customer{}
			if err := c.Bind(&temp); err != nil {
				return nil, fmt.Errorf("bind error")
			}
			temp.Id = uuid.NewV4().String()
			temp.CreatedAt = time.Now().Add(time.Hour * 3)
			temp.UpdatedAt = time.Now().Add(time.Hour * 3)
			fmt.Println("temp", temp)
			inter := structs.Map(temp)
			fmt.Println("inter", inter)
			customer := reflect.ValueOf(temp)
			for i := 0; i < customer.NumField(); i++ {
				if customer.Field(i).Interface() == "" {
					return nil, fmt.Errorf("empty value not accepted in customer")
				}
			}
			return inter, nil
		}
		/*
			case "order":
				{
					temp := models.Order{}
					if err := c.Bind(&temp); err != nil {
						return nil, fmt.Errorf("bind error")
					}
					temp.Id = uuid.NewV4().String()
					temp.CreatedAt = time.Now().Add(time.Hour * 3)
					temp.UpdatedAt = time.Now().Add(time.Hour * 3)
					order := reflect.ValueOf(temp)
					for i := 0; i < order.NumField(); i++ {
						if order.Field(i).Interface() == "" {
							return nil, fmt.Errorf("empty value not accepted in order")
						}
					}
					return temp, nil
				}*/
	default:
		{
			return nil, nil
		}
	}
}
