package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateCollections(client *mongo.Client, ctx context.Context) {
	opt := options.CreateCollection().SetValidator(addressValidator)
	err := client.Database("tesodev").CreateCollection(ctx, "Address", opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Address Collection Successfully created")

	opt = options.CreateCollection().SetValidator(productValidator)
	err = client.Database("tesodev").CreateCollection(ctx, "Product", opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Product Collection Successfully created")

	opt = options.CreateCollection().SetValidator(customerValidator)
	err = client.Database("tesodev").CreateCollection(ctx, "Customer", opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Customer Collection Successfully created")

	opt = options.CreateCollection().SetValidator(orderValidator)
	err = client.Database("tesodev").CreateCollection(ctx, "Order", opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Order Collection Successfully created")
}
