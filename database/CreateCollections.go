package database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateCollections() error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	client.Database("tesodev").Drop(ctx) // if exist
	opt := options.CreateCollection().SetValidator(addressValidator)
	err := client.Database("tesodev").CreateCollection(ctx, "Address", opt)
	if err != nil {
		return err
	}
	fmt.Println("Address Collection Successfully created")

	opt = options.CreateCollection().SetValidator(productValidator)
	err = client.Database("tesodev").CreateCollection(ctx, "Product", opt)
	if err != nil {
		return err
	}
	fmt.Println("Product Collection Successfully created")

	opt = options.CreateCollection().SetValidator(customerValidator)
	err = client.Database("tesodev").CreateCollection(ctx, "Customer", opt)
	if err != nil {
		return err
	}
	fmt.Println("Customer Collection Successfully created")

	opt = options.CreateCollection().SetValidator(orderValidator)
	err = client.Database("tesodev").CreateCollection(ctx, "Order", opt)
	if err != nil {
		return err
	}
	fmt.Println("Order Collection Successfully created")
	return nil
}
