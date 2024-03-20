package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//create the router
	router := mux.NewRouter()

	// Register the route
	home := homeHandler{}

	router.HandleFunc("/", home.ServerHTTP)

	// Start the server
	http.ListenAndServe(":8010", router)
}

type homeHandler struct{}

func (h *homeHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
