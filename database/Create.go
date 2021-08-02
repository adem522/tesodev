package database

import (
	"fmt"

	"github.com/adem522/tesodev/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(v interface{}, collectionName string) (interface{}, error) {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	data, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}
	collection := client.Database("tesodev").Collection(collectionName)
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func CreateCollections() error {
	client, ctx, cancel := Connect()
	defer Close(client, ctx, cancel)
	client.Database("tesodev").Drop(ctx) // if exist
	opt := options.CreateCollection().SetValidator(models.AddressValidator)
	err := client.Database("tesodev").CreateCollection(ctx, "Address", opt)
	if err != nil {
		return err
	}
	fmt.Println("Address Collection Successfully created")

	opt = options.CreateCollection().SetValidator(models.ProductValidator)
	err = client.Database("tesodev").CreateCollection(ctx, "Product", opt)
	if err != nil {
		return err
	}
	fmt.Println("Product Collection Successfully created")

	opt = options.CreateCollection().SetValidator(models.CustomerValidator)
	err = client.Database("tesodev").CreateCollection(ctx, "Customer", opt)
	if err != nil {
		return err
	}
	fmt.Println("Customer Collection Successfully created")

	opt = options.CreateCollection().SetValidator(models.OrderValidator)
	err = client.Database("tesodev").CreateCollection(ctx, "Order", opt)
	if err != nil {
		return err
	}
	fmt.Println("Order Collection Successfully created")
	return nil
}
