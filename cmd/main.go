package main

import (
	"deneme-structHandler/database"
	"deneme-structHandler/handlers"

	"github.com/labstack/echo"
)

func main() {
	client := database.Connect()
	defer database.Close(client)
	order := handlers.Order{
		Collection: client.Database("tesodev").Collection("Order"),
	}
	customer := handlers.Customer{
		Collection: client.Database("tesodev").Collection("Customer"),
	}
	address2 := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Address"),
		Name: "address",
	}
	product2 := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Product"),
		Name: "product",
	}
	order2 := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Order"),
		Name: "order",
	}
	customer2 := handlers.Collect{
		Col:  client.Database("tesodev").Collection("Customer"),
		Name: "customer",
	}
	e := echo.New()
	create := e.Group("/create")
	create.GET("/", handlers.CreateCollections) //create empty collections with validator
	create.POST("/product", product2.Create)    //insert address with models.address
	create.POST("/address", address2.Create)    //insert product with models.product
	create.POST("/order", order2.Create)        //create order with models.order
	create.POST("/customer", customer2.Create)  //insert customer with models.customer

	get := e.Group("/get")
	get.POST("/order", order.Get)                      //return all order when not take id
	get.POST("/customer", customer.Get)                //return all customer when not take id
	get.POST("/customerOrder", order.GetCustomerOrder) //return all customer order

	update := e.Group("/update")
	update.POST("/order", order.Update)       //update order with models.order
	update.POST("/customer", customer.Update) //update customer with models.customer

	delete := e.Group("/delete")
	delete.POST("/order", order.Delete)       //delete order with order id
	delete.POST("/customer", customer.Delete) //update customer with customer id

	e.POST("/changeStatus", order.ChangeStatus) //change status with status and order id
	e.POST("/validate", customer.Validate)      //check validate with customer id

	e.Logger.Fatal(e.Start(":8080"))
}
