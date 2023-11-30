package services

import (
	"context"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllMessagesByChatID(chatID string, messageCollection *mongo.Collection) ([]models.Message, error) {
	ctx := context.TODO() // Replace with a proper context in your application

	// Define a filter to find messages by chatID
	filter := bson.M{"chat": chatID}

	// Create a slice to store retrieved messages
	var messages []models.Message

	// Query the database
	cursor, err := messageCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate through the results and decode messages
	for cursor.Next(ctx) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func SendMessage(content string, chatID string, userID string, messageCollection *mongo.Collection, chatCollection *mongo.Collection) (*models.Message, error) {
	ctx := context.TODO() // Replace with a proper context in your application

	// Convert the userID string to ObjectID
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	chatObjectId, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return nil, err
	}

	// Construct a new message
	newMessage := models.Message{
		Sender:  userObjectID,
		Content: content,
		Chat:    chatObjectId,
	}

	// Insert the message into the message collection
	insertResult, err := messageCollection.InsertOne(ctx, newMessage)
	if err != nil {
		return nil, err
	}

	// Retrieve the inserted message ID
	insertedID := insertResult.InsertedID.(primitive.ObjectID)

	// Retrieve the inserted message using the ID
	filter := bson.M{"_id": insertedID}
	var insertedMessage models.Message
	if err := messageCollection.FindOne(ctx, filter).Decode(&insertedMessage); err != nil {
		return nil, err
	}

	// Update latestMessage field in the chat
	update := bson.M{"$set": bson.M{"latestMessage": insertedMessage}}
	if _, err := chatCollection.UpdateByID(ctx, chatID, update); err != nil {
		return nil, err
	}

	return &insertedMessage, nil
}
