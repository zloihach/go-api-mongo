package driver

import (
	"context"
	"go-api-mongo/modules/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var app config.GoAppTools

func Connection(URL string) *mongo.Client {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		app.ErrorLogger.Fatalln(err)
	}
	return client
}
