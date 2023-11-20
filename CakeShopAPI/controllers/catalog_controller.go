package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/models"
	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CatalogController struct {
	CakeCollection *services.CakeCollection
	Client         *mongo.Client
}

type Owner struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
}

func (c *CatalogController) CatalogRoutes(router *mux.Router) {
	router.HandleFunc("/catalog/", c.GetAll).Methods("GET")
	router.HandleFunc("/catalog/top5", c.GetTopFive).Methods("GET")
	router.HandleFunc("/catalog/deals", c.GetOnOffer).Methods("GET")
	router.HandleFunc("/catalog/search", c.Search).Methods("GET")
	//TODO Modify myCakes
	router.HandleFunc("/catalog/myCakes/{ownerID}", c.GetAllByOwner).Methods("GET")
	router.HandleFunc("/catalog/create", c.Create).Methods("POST")
	router.HandleFunc("/catalog/{id}", c.GetDetails).Methods("GET")
	router.HandleFunc("/catalog/{id}/update", c.Update).Methods("PUT")
	router.HandleFunc("/catalog/{id}/delete", c.Delete).Methods("DELETE")
}

func (c *CatalogController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cakeCollection := c.CakeCollection

	services.GetAll(context.TODO(), w, r, cakeCollection)
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog GetAll route"})
}
func (c *CatalogController) GetTopFive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cakeCollection := c.CakeCollection

	services.GetTopFive(context.TODO(), w, r, cakeCollection)
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog Top5 route"})
}
func (c *CatalogController) GetOnOffer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cakeCollection := c.CakeCollection

	services.GetCakesOnOffer(context.TODO(), w, r, cakeCollection)
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog OnOffer route"})
}
func (c *CatalogController) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Catalog Search route"})
}

// TODO - fix Get all by owner
func (c *CatalogController) GetAllByOwner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cakeCollection := c.CakeCollection
	ownerID := mux.Vars(r)["ownerID"]

	cakeList, err := services.GetAllCakesByOwner(context.TODO(), cakeCollection, ownerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting cakes by owner: %v", err)
		return
	}

	if cakeList == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "No cakes found for this owner")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cakeList)
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

	cakeCollection := c.CakeCollection

	cakeID := mux.Vars(r)["id"]

	objectID, err := primitive.ObjectIDFromHex(cakeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error converting cakeID to ObjectID: %v", err)
		return
	}

	cakeCursor := cakeCollection.Collection.FindOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader((http.StatusInternalServerError))
		fmt.Fprintf(w, "Error getting cake details: %v", err)
		return
	}

	var cake models.Cake
	if err := cakeCursor.Decode(&cake); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding cake document: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cake)
}

func (c *CatalogController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Parse the cake ID from the URL using Gorilla Mux
	params := mux.Vars(r)
	cakeID := params["id"]

	fmt.Printf("CakeID %s", cakeID)

	// Decode the updated cake details from the request body
	var updatedCake models.Cake
	err := json.NewDecoder(r.Body).Decode(&updatedCake)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Call the service function to update the cake
	err = services.UpdateCake(r.Context(), c.CakeCollection, cakeID, updatedCake)
	if err != nil {
		http.Error(w, "Failed to update the cake", http.StatusInternalServerError)
		return
	}

	// Send a success response if the update was successful
	json.NewEncoder(w).Encode(map[string]string{"status": "Cake updated successfully"})
}

func (c *CatalogController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the cake ID from the request URL
	params := mux.Vars(r)
	cakeID := params["id"]

	// Perform the deletion
	err := services.DeleteCake(r.Context(), c.CakeCollection, cakeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// If successful, respond with success status
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cake deleted successfully"})
}
