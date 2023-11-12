package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/controllers"
	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Replace "<connection string>" with your actual Atlas connection string
	const uri = "mongodb://localhost:27017/"

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection to MongoDB closed.")
	}()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	collection := client.Database("cakeShop").Collection("cakes") // Replace "test" with your actual database name
	cakeCollection := services.NewCakeCollection(collection)

	r := mux.NewRouter()

	userController := &controllers.UserController{}
	catalogController := &controllers.CatalogController{CakeCollection: cakeCollection}
	chatController := &controllers.ChatController{}
	messageController := &controllers.MessageController{}

	userController.RegisterRoutes(r)
	catalogController.CatalogRoutes(r)
	chatController.ChatRoutes(r)
	messageController.MessageRoutes(r)

	server := &http.Server{Addr: ":8080", Handler: r}

	log.Println("REST API running on port: ", server.Addr)
	if err := http.ListenAndServe(server.Addr, r); err != nil {
		log.Fatal(err)
	}
}
