package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ChatName      string               `bson:"chatName,omitempty"`
	IsGroupChat   bool                 `bson:"isGroupChat,omitempty"`
	Users         []primitive.ObjectID `bson:"users,omitempty"`
	LatestMessage primitive.ObjectID   `bson:"latestMessage,omitempty"`
	GroupAdmin    primitive.ObjectID   `bson:"groupAdmin,omitempty"`
	CreatedAt     time.Time            `bson:"createdAt,omitempty"`
	UpdatedAt     time.Time            `bson:"updatedAt,omitempty"`
}
