package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/services"
	"github.com/gorilla/mux"
)

type CatalogController struct {
	CakeCollection *services.CakeCollection
}

func (c *CatalogController) CatalogRoutes(router *mux.Router) {
	router.HandleFunc("/catalog/", c.GetAll).Methods("GET")
	router.HandleFunc("/catalog/top5", c.GetTopFive).Methods("GET")
	router.HandleFunc("/catalog/deals", c.GetOnOffer).Methods("GET")
	router.HandleFunc("/catalog/search", c.Search).Methods("GET")
	router.HandleFunc("/catalog/myCakes", c.GetAllByOwner).Methods("GET")
	router.HandleFunc("/catalog/create", c.Create).Methods("POST")
	router.HandleFunc("/catalog/{id}", c.GetDetails).Methods("GET")
	router.HandleFunc("/cataog/{id}/update", c.Update).Methods("PUT")
	router.HandleFunc("/cataog/{id}/delete", c.Delete).Methods("DELETE")
}

func (c *CatalogController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog GetAll route"})
}
func (c *CatalogController) GetTopFive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog Top5 route"})
}
func (c *CatalogController) GetOnOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog OnOffer route"})
}
func (c *CatalogController) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog Search route"})
}
func (c *CatalogController) GetAllByOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog All By Owner route"})
}
func (c *CatalogController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Call the Create function from CakeCollection
	services.Create(w, r, c.CakeCollection)

	// Respond with a success message
	json.NewEncoder(w).Encode(map[string]string{"status": "Cake created successfully"})
}
func (c *CatalogController) GetDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog Details route"})
}
func (c *CatalogController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog Update route"})
}
func (c *CatalogController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog Delete route"})
}
