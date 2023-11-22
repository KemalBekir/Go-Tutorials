package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Chat struct {
	ChatName      string    `json:"chatName"`
	IsGroupChat   bool      `json:"isGroupChat"`
	Users         []string  `json:"users"`
	LatestMessage string    `json:"latestMessage"`
	GroupAdmin    string    `json:"groupAdmin"`
	CreatedAt     time.Time `json:"createdAt"`
}

func accessChat(userId, ownerId string, chatCollection *mongo.Collection) (*Chat, error) {
	// Perform createChat operation asynchronously using goroutine
	chatChan := make(chan *Chat)
	errChan := make(chan error)

	go func() {
		result, err := createChat(userId, ownerId, chatCollection)
		if err != nil {
			errChan <- err
			return
		}
		chatChan <- &result
	}()

	// Wait for the asynchronous operation to complete using select
	select {
	case result := <-chatChan:
		// In this case, the chat is already created within the createChat function,
		// so we don't need to perform any save operation here.

		// Return the result
		return result, nil
	case err := <-errChan:
		return nil, err
	}
}

func createChat(userId string, ownerId string, chatCollection *mongo.Collection) (Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var isChat Chat
	filter := bson.M{
		"isGroupChat": false,
		"$and": []bson.M{
			{"users": bson.M{"$elemMatch": bson.M{"$eq": userId}}},
			{"users": bson.M{"$elemMatch": bson.M{"$eq": ownerId}}},
		},
	}

	err := chatCollection.FindOne(ctx, filter).Decode(&isChat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Create a new chat since one doesn't exist
			chatData := Chat{
				ChatName:    "sender",
				IsGroupChat: false,
				Users:       []string{userId, ownerId},
				CreatedAt:   time.Now(),
			}

			_, err := chatCollection.InsertOne(ctx, chatData)
			if err != nil {
				return Chat{}, err
			}

			// Retrieve the created chat
			err = chatCollection.FindOne(ctx, filter).Decode(&isChat)
			if err != nil {
				return Chat{}, err
			}
		} else {
			return Chat{}, err
		}
	}

	// Perform necessary population of fields
	// Logic for populating fields goes here

	return isChat, nil
}
