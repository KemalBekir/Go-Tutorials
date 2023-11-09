package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ChatController struct{}

func (c *ChatController) ChatRoutes(router *mux.Router) {
	router.HandleFunc("/chat/", c.HandleChat)

}

func (c *ChatController) HandleChat(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "Chat AccessChat route"})
	}

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "Chat GetChats route"})
	}
}
