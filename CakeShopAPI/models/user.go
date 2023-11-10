package models

type User struct {
	ID             string   `json:"_id" bson:"_id"`
	Email          string   `json:"email" bson:"email"`
	Username       string   `json:"username" bson:"username"`
	HashedPassword string   `json:"hashedPassword" bson:"hashedPassword"`
	Role           string   `json:"role" bson:"role"`
	MyAds          []string `json:"myAds" bson:"myAds"`
}
