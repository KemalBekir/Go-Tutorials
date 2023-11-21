package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/services"
	"github.com/gorilla/mux"
)

type UserController struct {
	UserCollection *services.UserCollection
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *UserController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", c.Login).Methods("POST")
	router.HandleFunc("/register", c.Register).Methods("POST")
	router.HandleFunc("/logout", c.Logout).Methods("POST")
	router.HandleFunc("/profile", c.Profile).Methods("GET")
}

func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request data
	var request LoginReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	// Call UserCollection Login method
	token, err := c.UserCollection.Login(request.Email, request.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Respond with token
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User logged in successfully",
		"token":   token,
	})
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse request data
	var request RegisterReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	//Call UserCollection Register method
	tokenData, err := c.UserCollection.Register(request.Username, request.Email, request.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMsg := "Registration failed"

		if strings.Contains(err.Error(), "email already exists") {
			statusCode = http.StatusBadRequest
			errorMsg = "Email already exists"
		}

		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(map[string]string{"error": errorMsg})
		// Log the error for internal debugging
		log.Printf("Registration failed: %s", err.Error())
		return
	}

	// Respond with token
	json.NewEncoder(w).Encode(tokenData)
}

func (c *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Logout route"})
}

func (c *UserController) Profile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Profile route"})
}
