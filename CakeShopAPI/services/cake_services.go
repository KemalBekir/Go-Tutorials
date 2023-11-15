package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CakeCollection struct {
	Collection *mongo.Collection
}

func NewCakeCollection(collection *mongo.Collection) *CakeCollection {
	return &CakeCollection{
		Collection: collection, // replace "cakes" with your actual collection name
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
	// Limit the number of cakes to 5
	cakesCursor, err := cakeCollection.Collection.Find(ctx, bson.M{}, options.Find().SetLimit(5))
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
func GetAllCakesByOwner(ctx context.Context, w http.ResponseWriter, r *http.Request, cakeCollection *CakeCollection, ownerId string) {
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

func Create(w http.ResponseWriter, r *http.Request, cakeCollection *CakeCollection) {
	// Decode the request body into a models.Cake object
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
	result, err := cakeCollection.Collection.InsertOne(context.TODO(), newCake)
	if err != nil {
		log.Printf("Error creating new cake: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating new cake: %v", err)
		return
	}

	// Log the inserted document ID
	log.Printf("Inserted document ID: %v", result.InsertedID)

	// Encode the new cake to JSON and write it to the response
	json.NewEncoder(w).Encode(newCake)
}
