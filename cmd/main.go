package main

import (
	"deneme-structHandler/database"
	"deneme-structHandler/handlers"

	"github.com/labstack/echo"
)

func main() {
	client := database.Connect()
	defer database.Close(client)
	product := handlers.Product{
		Collection: client.Database("tesodev").Collection("Product"),
	}
	address := handlers.Address{
		Collection: client.Database("tesodev").Collection("Address"),
	}
	order := handlers.Order{
		Collection: client.Database("tesodev").Collection("Order"),
	}
	customer := handlers.Customer{
		Collection: client.Database("tesodev").Collection("Customer"),
	}
	e := echo.New()
	create := e.Group("/create")
	create.GET("/", handlers.CreateCollections) //create empty collections with validator
	create.POST("/product", product.Create)     //insert address with models.address
	create.POST("/address", address.Create)     //insert product with models.product
	create.POST("/order", order.Create)         //create order with models.order
	create.POST("/customer", customer.Create)   //insert customer with models.customer

	get := e.Group("/get")
	get.POST("/order", order.Get)                      //return all order when not take id
	get.POST("/customer", customer.Get)                //return all customer when not take id
	get.POST("/customerOrder", order.GetCustomerOrder) //return all customer order

	e.POST("/changeStatus", order.ChangeStatus) //change status with status and order id
	e.POST("/validate", customer.Validate)      //check validate with customer id

	update := e.Group("/update")
	update.POST("/order", order.Update)       //update order with models.order
	update.POST("/customer", customer.Update) //update customer with models.customer

	delete := e.Group("/delete")
	delete.POST("/order", order.Delete)       //delete order with order id
	delete.POST("/customer", customer.Delete) //update customer with customer id
	e.Logger.Fatal(e.Start(":8080"))
}
