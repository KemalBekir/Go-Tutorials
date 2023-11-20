package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

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
// TODO adjust to be for authenticated users only
func GetAllCakesByOwner(ctx context.Context, cakeCollection *CakeCollection, ownerID string) ([]models.Cake, error) {
	ownerObjectID, err := primitive.ObjectIDFromHex(ownerID)
	if err != nil {
		log.Printf("Error converting ownerID to ObjectID: %s", err.Error())
		return nil, err
	}

	cakesCursor, err := cakeCollection.Collection.Find(ctx, bson.M{"owner": ownerObjectID})
	if err != nil {
		log.Printf("Error finding cakes for owner %s: %s", ownerID, err.Error())
		return nil, err
	}

	defer func() {
		if err := cakesCursor.Close(ctx); err != nil {
			log.Printf("Error closing cursor: %s", err.Error())
		}
	}()

	var cakeList []models.Cake
	for cakesCursor.Next(ctx) {
		var cake models.Cake
		if err := cakesCursor.Decode(&cake); err != nil {
			log.Printf("Error decoding cake for owner %s: %s", ownerID, err.Error())
			return nil, err
		}
		log.Printf("Retrieved cake: %+v", cake) // Logging each cake
		cakeList = append(cakeList, cake)
	}

	if err := cakesCursor.Err(); err != nil {
		log.Printf("Error iterating cursor: %s", err.Error())
		return nil, err
	}

	log.Printf("Retrieved %d cakes for owner %s", len(cakeList), ownerID)
	return cakeList, nil
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

func GetDetails(ctx context.Context, w http.ResponseWriter, r *http.Request, cakeCollection *CakeCollection) {
	cakeId := mux.Vars(r)["cakeID"]

	objectID, err := primitive.ObjectIDFromHex(cakeId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error converting cakeID to ObjectID: %v", err)
		return
	}

	cakeCursor := cakeCollection.Collection.FindOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Cake with ID '%s' not found", cakeId)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
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

func UpdateCake(ctx context.Context, cakeCollection *CakeCollection, cakeID string, updatedCake models.Cake) error {
	objectID, err := primitive.ObjectIDFromHex(cakeID)
	if err != nil {
		return err
	}

	fmt.Printf("CakeID is %s: ", objectID)

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"cakeName":    updatedCake.CakeName,
		"description": updatedCake.Description,
		"price":       updatedCake.Price,
		"type":        updatedCake.Type,
		"images":      updatedCake.Images,
		"likes":       updatedCake.Likes,
		"onOffer":     updatedCake.OnOffer,
		"discount":    updatedCake.Discount,
		"created_at":  updatedCake.CreatedAt,
		"updated_at":  time.Now(), // Update the timestamp
	}}

	_, err = cakeCollection.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCake(ctx context.Context, cakeCollection *CakeCollection, cakeID string) error {
	objectId, err := primitive.ObjectIDFromHex(cakeID)
	if err != nil {
		return err
	}

	res, err := cakeCollection.Collection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no matching cake found to delete")
	}
	return nil
}
