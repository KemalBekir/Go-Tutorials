package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CakeCollection struct {
	Collection *mongo.Collection
}

func NewCakeCollection(db *mongo.Database) *CakeCollection {
	return &CakeCollection{
		Collection: db.Collection("cakes"), // replace "cakes" with your actual collection name
	}
}

// Get all cakes
func GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request, cakeCollection *CakeCollection) {
	cakesCursor, err := cakeCollection.Collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting top 5 cakes: %v", err)
		return
	}

	var cakeList []models.Cake // Replace Cake with your actual struct type
	if err := cakesCursor.All(ctx, &cakeList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding cakes: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cakeList)
}

// Get top 5 cakes
func GetTopFive(ctx context.Context, w http.ResponseWriter, r *http.Request, cakeCollection *CakeCollection) {
	cakesCursor, err := cakeCollection.Collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting top 5 cakes: %v", err)
		return
	}

	var cakeList []models.Cake // Replace YourCakeStruct with your actual struct type
	if err := cakesCursor.All(ctx, &cakeList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding cakes: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cakeList)
}

// Get cakes on offer
func GetCakesOnOffer(ctx context.Context, w http.ResponseWriter, r *http.Request, cakeCollection *CakeCollection) {
	cakesCursor, err := cakeCollection.Collection.Find(ctx, bson.M{"onOffer": true})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting cakes on offer: %v", err)
		return
	}

	var cakeList []models.Cake // Replace Cake with your actual struct type
	if err := cakesCursor.All(ctx, &cakeList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding cakes: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cakeList)
}

// Get all cakes by owner
func GetAllCakesByOwner(ctx context.Context, w http.ResponseWriter, r *http.Request, cakeCollection *CakeCollection) {
	ownerID := mux.Vars(r)["ownerID"]

	cakesCursor, err := cakeCollection.Collection.Find(ctx, bson.M{"owner": ownerID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting cakes by owner: %v", err)
		return
	}

	var cakeList []models.Cake // Replace YourCakeStruct with your actual struct type
	if err := cakesCursor.All(ctx, &cakeList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding cakes: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cakeList)
}

func Create(ctx context.Context, w http.ResponseWriter, r *http.Request, cakeCollection *CakeCollection) {
	var newCake models.Cake
	err := json.NewDecoder(r.Body).Decode(&newCake)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request body: %v", err)
		return
	}

	// Ensure the ID is set to a new unique ObjectID
	newCake.ID = primitive.NewObjectID()

	// Set default values or perform any additional validation if needed

	// Insert the new cake into the collection
	_, err = cakeCollection.Collection.InsertOne(ctx, newCake)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating new cake: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Cake created successfully")
}
