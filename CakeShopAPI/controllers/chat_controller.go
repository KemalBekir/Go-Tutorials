package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/middleware"
	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/services"
	"github.com/gorilla/mux"
)

type ChatController struct {
	ChatCollection *services.ChatCollection
}

func (c *ChatController) ChatRoutes(router *mux.Router) {
	router.HandleFunc("/chat/", c.HandleChat)

}

func (c *ChatController) HandleChat(w http.ResponseWriter, r *http.Request) {
	chatCollection := c.ChatCollection
	token := r.Header.Get("x-authorization")
	if r.Method == "POST" {
		userID, err := middleware.ExtractUserIDFromToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Unauthorized access"})
			return
		}

		var requestBody struct {
			OwnerID string `json:"ownerId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid request payload"})
			return
		}

		chat, err := services.AccessChat(userID, requestBody.OwnerID, chatCollection.Collection)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "Error accessing chat"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(chat)
	}

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "Chat GetChats route"})
	}
}
