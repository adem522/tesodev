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
	address := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "Address",
	}
	product := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Product"),
		Name: "Product",
	}
	order := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Order"),
		Name: "Order",
	}
	customer := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Customer"),
		Name: "Customer",
	}

	e := echo.New()
	e.Use(utils.Logger()) //Uri and Status logger

	create := e.Group("")                                  //endpoint group
	create.GET("/collections", handlers.CreateCollections) //create empty collections with validator
	create.POST("/product", product.Create)                //insert product with models.product
	create.POST("/address", address.Create)                //insert address with models.address
	create.POST("/order", order.Create)                    //create order with models.order
	create.POST("/customer", customer.Create)              //insert customer with models.customer

	get := e.Group("")                 //endpoint group
	get.PUT("/customer", customer.Get) //return all customer when not take id
	get.PUT("/order", order.Get)       //return all order when not take order id
	//							       //or return customer order when take customer id

	update := e.Group("")                      //endpoint group
	update.PATCH("/customer", customer.Update) //update order with models.order
	update.PATCH("/order", order.Update)       //update customer with models.customer

	delete := e.Group("")                       //endpoint group
	delete.DELETE("/order", order.Delete)       //delete order with order id
	delete.DELETE("/customer", customer.Delete) //delete customer with customer id

	e.POST("/changeStatus", order.ChangeStatus) //change status with status and order id
	e.POST("/validate", customer.Validate)      //check validate with customer id

	e.Logger.Fatal(e.Start(":8080")) //server start in localhost:/8080
}
