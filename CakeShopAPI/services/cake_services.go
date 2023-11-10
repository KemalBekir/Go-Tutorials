package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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
func GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cakes, err := cakeCollection.Collection.Find(ctx, bson.M{}).Sort(bson.D{{"createdAt", -1}}).Limit(5).All(ctx, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting top 5 cakes: %v", err)
		return
	}

	// Encode the cakes to JSON and write them to the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cakes)
}

// Get top 5 cakes
// Get top 5 cakes
func GetTopFive(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cakeCollection := NewCakeCollection( /* pass your mongo database instance here */ )

	cakes, err := cakeCollection.Collection.Find(ctx, bson.M{}).Sort(bson.D{{"createdAt", -1}}).Limit(5)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting top 5 cakes: %v", err)
		return
	}

	var cakeList []CakeCollection // Replace YourCakeStruct with your actual struct type
	if err := cakes.All(ctx, &cakeList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding cakes: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cakeList)
}

// Get cakes on offer
func GetCakesOnOffer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cakeCollection := NewCakeCollection( /* pass your mongo database instance here */ )

	cakes, err := cakeCollection.Collection.Find(ctx, bson.M{"onOffer": true})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting cakes on offer: %v", err)
		return
	}

	var cakeList []CakeCollection // Replace CakeCollection with your actual struct type
	if err := cakes.All(ctx, &cakeList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding cakes: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cakeList)
}

// Get all cakes by owner
func GetAllCakesByOwner(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cakeCollection := NewCakeCollection( /* pass your mongo database instance here */ )
	ownerID := mux.Vars(r)["ownerID"]

	cakes, err := cakeCollection.Collection.Find(ctx, bson.M{"owner": ownerID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting cakes by owner: %v", err)
		return
	}

	var cakeList []CakeCollection // Replace YourCakeStruct with your actual struct type
	if err := cakes.All(ctx, &cakeList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error decoding cakes: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cakeList)
}
