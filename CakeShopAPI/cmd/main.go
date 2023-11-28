package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/controllers"
	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/middleware"
	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Replace "<connection string>" with your actual Atlas connection string
	const uri = "mongodb://localhost:27017"

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

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	collection := client.Database("cakeShop").Collection("cakes")
	authDB := client.Database("cakeShop").Collection("users")
	chatDB := client.Database("cakeShop").Collection("chats")
	userCollection := services.NewUserCollection(authDB) // Replace "test" with your actual database name
	chatCollection := services.NewChatCollection(chatDB)
	cakeCollection := services.NewCakeCollection(collection)

	var result bson.M
	if err := client.Database("cakeShop").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	r := mux.NewRouter()
	r.Use(middleware.CorsMiddleware)

	userController := &controllers.UserController{
		UserCollection: userCollection,
	}
	catalogController := &controllers.CatalogController{
		CakeCollection: cakeCollection,
	}
	chatController := &controllers.ChatController{
		ChatCollection: chatCollection,
		UserCollection: userCollection,
	}
	messageController := &controllers.MessageController{}

	userController.RegisterRoutes(r)
	catalogController.CatalogRoutes(r)
	chatController.ChatRoutes(r)
	messageController.MessageRoutes(r)

	cNames, err := client.Database("cakeShop").ListCollectionNames(ctx, bson.D{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cNames)

	server := &http.Server{Addr: ":8080", Handler: r}

	log.Println("REST API running on port: ", server.Addr)
	if err := http.ListenAndServe(server.Addr, r); err != nil {
		log.Fatal(err)
	}
}
