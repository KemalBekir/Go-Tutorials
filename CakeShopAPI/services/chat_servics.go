package services

import (
	"context"
	"time"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Chat struct {
	ChatName      string          `json:"chatName"`
	IsGroupChat   bool            `json:"isGroupChat"`
	Users         []string        `json:"users"`
	LatestMessage *models.Message `json:"latestMessage"`
	GroupAdmin    string          `json:"groupAdmin"`
	CreatedAt     time.Time       `json:"createdAt"`
}

type ChatCollection struct {
	Collection *mongo.Collection
}

func NewChatCollection(collection *mongo.Collection) *ChatCollection {
	return &ChatCollection{
		Collection: collection,
	}
}

func AccessChat(userId, ownerId string, chatCollection *mongo.Collection) (*Chat, error) {
	// Perform createChat operation asynchronously using goroutine
	chatChan := make(chan *Chat)
	errChan := make(chan error)

	go func() {
		result, err := CreateChat(userId, ownerId, chatCollection)
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

func CreateChat(userId string, ownerId string, chatCollection *mongo.Collection) (Chat, error) {
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

	return isChat, nil
}

func GetChats(userId string, chatCollection *mongo.Collection, userCollection *mongo.Collection) ([]Chat, error) {
	filter := bson.M{
		"users": bson.M{
			"$elemMatch": bson.M{"$eq": userId},
		},
	}

	options := options.Find().
		SetSort(bson.D{{"updatedAt", -1}}).
		SetProjection(bson.M{
			"users":         0,
			"groupAdmin":    bson.M{"hashedPassword": 0, "roel": 0, "myAds": 0},
			"latestMessage": -1,
		}).
		SetMaxTime(2 * time.Second)

	cursor, err := chatCollection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var chats []Chat
	if err := cursor.All(context.Background(), &chats); err != nil {
		return nil, err
	}

	for i, chat := range chats {
		if chat.LatestMessage != nil {
			sender, err := GetUser(chat.LatestMessage.Sender, userCollection)
			if err != nil {
				return nil, err
			}
			chats[i].LatestMessage.Sender = sender.ID
		}
	}

	return chats, nil

}

func GetUser(userID primitive.ObjectID, userCollection *mongo.Collection) (*User, error) {
	filter := bson.M{"_id": userID}
	var user User
	if err := userCollection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
