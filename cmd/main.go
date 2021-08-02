package main

import (
	"github.com/adem522/tesodev/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	create := e.Group("/create")
	defineDatabase(create)
	create.POST("/order", handlers.CreateOrder)       //create order with models.order
	create.POST("/customer", handlers.CreateCustomer) //insert customer with models.customer

	get := e.Group("/get")
	get.POST("/order", handlers.GetOrder)                 //return all order when not take id
	get.POST("/customer", handlers.GetCustomer)           //return all customer when not take id
	get.POST("/customerOrder", handlers.GetCustomerOrder) //return order when take customer id

	delete := e.Group("/delete")
	delete.POST("/order", handlers.DeleteOrder)       //delete order with order id
	delete.POST("/customer", handlers.DeleteCustomer) //update customer with customer id

	e.POST("/changeStatus", handlers.ChangeStatus) //change status with status and order id
	e.POST("/validate", handlers.Validate)         //check validate with customer id

	update := e.Group("/update")
	update.POST("/order", handlers.UpdateOrder)       //update order with models.order
	update.POST("/customer", handlers.UpdateCustomer) //update customer with models.customer

	e.Logger.Fatal(e.Start(":8080"))
}

func defineDatabase(create *echo.Group) {
	create.GET("/", handlers.CreateCollections)     //create empty collections with validator
	create.POST("/address", handlers.CreateAddress) //insert address with models.address
	create.POST("/product", handlers.CreateProduct) //insert product with models.product
}
