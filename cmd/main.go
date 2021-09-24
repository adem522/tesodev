package main

import (
	"tesodev/database"
	"tesodev/handlers"
	"tesodev/utils"

	"github.com/labstack/echo"
)

func main() {
	client := database.Connect() //database client connect opening and return client
	defer database.Close(client) //database client connect closing
	database := handlers.Database{
		Database: client.Database("tesodev"),
	}

	e := echo.New()
	e.Use(utils.Logger()) //Uri and Status logger

	create := e.Group("")                                  //endpoint group
	create.GET("/collections", handlers.CreateCollections) //create empty collections with validator
	create.POST("/product", database.Create)               //insert product with models.product
	create.POST("/address", database.Create)               //insert address with models.address
	create.POST("/order", database.Create)                 //create order with models.order
	create.POST("/customer", database.Create)              //insert customer with models.customer

	get := e.Group("")                 //endpoint group
	get.PUT("/customer", database.Get) //return all customer when not take id
	get.PUT("/order", database.Get)    //return all order when not take order id
	//							       //or return customer order when take customer id

	update := e.Group("")                      //endpoint group
	update.PATCH("/customer", database.Update) //update order with models.order
	update.PATCH("/order", database.Update)    //update customer with models.customer

	delete := e.Group("")                       //endpoint group
	delete.DELETE("/order", database.Delete)    //delete order with order id
	delete.DELETE("/customer", database.Delete) //delete customer with customer id

	e.PATCH("/changeStatus", database.ChangeStatus) //change status with status and order id
	e.PUT("/validate", database.Validate)           //check validate with customer id

	e.Logger.Fatal(e.Start(":8080")) //server start in localhost:/8080
}
