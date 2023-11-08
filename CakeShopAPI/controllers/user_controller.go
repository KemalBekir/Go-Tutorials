package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct{}

func (c *UserController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", c.Login).Methods("POST")
	router.HandleFunc("/register", c.Register).Methods("POST")
	router.HandleFunc("/logout", c.Logout).Methods("POST")
	router.HandleFunc("/profile", c.Profile).Methods("GET")
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "User logged in successfully",
		"token":   "your-auth-token",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Register route"})
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Logout route"})
}

func (c *UserController) Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Profile route"})
}
