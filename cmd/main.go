package main

import (
	"github.com/adem522/tesodev/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	create := e.Group("/create")

	defineDatabase(create)
	//create order with models.order
	create.POST("/order", handlers.CreateOrder)
	//insert customer with models.customer
	create.POST("/customer", handlers.CreateCustomer)

	update := e.Group("/update")
	update.POST("/order", handlers.UpdateOrder)

	get := e.Group("/get")
	//return all order when not take id
	//return order when take order id
	get.POST("/order", handlers.GetOrder)
	//return all customer when not take id
	//return customer when take customer id
	get.POST("/customer", handlers.GetCustomer)
	//return order when take customer id
	get.POST("/customerOrder", handlers.GetCustomerOrder)

	e.Logger.Fatal(e.Start(":8080"))
}

func defineDatabase(create *echo.Group) {
	//create empty collections with validator
	create.GET("/", handlers.CreateCollections)
	//insert address with models.address
	create.POST("/address", handlers.CreateAddress)
	//insert product with models.product
	create.POST("/product", handlers.CreateProduct)
}
