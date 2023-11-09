package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/controllers"
	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/utils"
	"github.com/gorilla/mux"
)

func main() {
	client, err := utils.ConnectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	r := mux.NewRouter()

	userController := &controllers.UserController{}
	catalogController := &controllers.CatalogController{}
	chatController := &controllers.ChatController{}
	messageController := &controllers.MessageController{}
	userController.RegisterRoutes(r)
	catalogController.CatalogRoutes(r)
	chatController.ChatRoutes(r)
	messageController.MessageRoutes(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "Rest service operational"})
	})

	server := &http.Server{Addr: ":8080", Handler: r}

	log.Println("REST API running on port: ", server.Addr)
	http.ListenAndServe(server.Addr, r)
}
