package main

import (
	"context"
	"fmt"
	"time"

	database "github.com/adem522/tesodev/database"
	. "github.com/adem522/tesodev/models"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, ctx, cancel := database.Connect()
	defer database.Close(client, ctx, cancel)
	/*e := echo.New()
	create := e.Group("/create")
	create.GET("/", handlers.Create)*/

	//client.Database("tesodev").Drop(ctx)
	database.CreateCollections(client, ctx)

	//fmt.Println(uuid.NewV4())
	//fn0(client, ctx)
	//fn2(client, ctx)
	//fn3(client, ctx)
}

func fn0(client *mongo.Client, ctx context.Context) {
	add := Address{
		AddressLine: "şirinevler mh.",
		City:        "bursa",
		Country:     "turkey",
		CityCode:    16,
	}
	data := Customer{
		Id:        uuid.NewV4(),
		Name:      "adem",
		Email:     "adem@hotmail.com",
		Address:   add,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := client.Database("tesodev").Collection("Customer").InsertOne(ctx, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("dataid", data.Id, result)
}

func fn2(client *mongo.Client, ctx context.Context) {
	data := Product{
		Id:       uuid.NewV4(),
		ImageUrl: "aaa",
		Name:     "aaa",
	}
	result, err := client.Database("tesodev").Collection("Product").InsertOne(ctx, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func fn3(client *mongo.Client, ctx context.Context) {
	order := Order{
		Id:         uuid.NewV4(),
		CustomerId: uuid.NewV4(),
		Quantity:   "4",
		Price:      2.2,
		Status:     "available",
		Address: Address{
			AddressLine: "asaa",
			City:        "bursa",
			Country:     "turkey",
			CityCode:    16,
		},
		Product: Product{
			Id:       uuid.NewV4(),
			ImageUrl: "imageurl",
			Name:     "product name",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := client.Database("tesodev").Collection("Order").InsertOne(ctx, order)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order.CustomerId, result)
}
