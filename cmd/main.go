package main

import (
	"github.com/adem522/tesodev/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	create := e.Group("/create")
	create.GET("/", handlers.CreateCollections)
	create.POST("/address", handlers.CreateAddress)
	create.POST("/product", handlers.CreateProduct)
	create.POST("/order", handlers.CreateOrder)
	create.POST("/customer", handlers.CreateCustomer)

	update := e.Group("/update")
	update.POST("/order", handlers.UpdateOrder)

	get := e.Group("/get")
	get.POST("/order", handlers.GetOrder)

	e.Logger.Fatal(e.Start(":8080"))
}
