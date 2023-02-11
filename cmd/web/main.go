package main

//https://ayaacodes.hashnode.dev/creating-a-scalable-api-with-go-gin-and-mongodb-atlas
import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api-mongo/driver"
	"go-api-mongo/modules/config"
	"log"
	"os"
)

var app config.GoAppTools

func main() {
	InfoLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)

	app.InfoLogger = *InfoLogger
	app.ErrorLogger = *ErrorLogger

	err := godotenv.Load()
	if err != nil {
		app.ErrorLogger.Fatal("No .env file available")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		app.ErrorLogger.Fatalln("mongodb uri string not found : ")
	}

	//connection to the database
	Client := driver.Connection(uri)
	defer func() {
		if err = Client.Disconnect(context.TODO()); err != nil {
			app.ErrorLogger.Fatal(err)
			return
		}

	}()

	appRouter := gin.New()
	appRouter.GET("/", func(ctx *gin.Context) {
		app.InfoLogger.Println("Creating a scalable web application with Gin")
	})

	err = appRouter.Run()
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}
}
