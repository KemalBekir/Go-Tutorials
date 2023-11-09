package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type MessageController struct{}

func (c *MessageController) MessageRoutes(router *mux.Router) {
	router.HandleFunc("/message/{id}", c.GetAllMessages).Methods("GET")
	router.HandleFunc("/message/{id}/create", c.SendMessage).Methods("POST")
}

func (c *MessageController) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Message GetAllMessages route"})
}

func (c *MessageController) SendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Message SendMessage route"})
}
