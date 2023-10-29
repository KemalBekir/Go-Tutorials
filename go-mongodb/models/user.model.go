package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignUpInput struct {
	Name            string    `json:"name" "bson:"name" binding="required""`
	Email           string    `json:"email" "bson:"email" binding="required""`
	Password        string    `json:"password" "bson:"password" binding="required, min=8""`
	PasswordConfirm string    `json:"passwordConfirm" "bson:"passwordConfigm, omitempty" binding="required""`
	Role            string    `json:"role" "bson:"role" binding="required""`
	Verified        bool      `json:"verified" "bson:"verified""`
	CreatedAt       time.Time `json:"created_at" "bson:"created_at""`
	UpdatedAt       time.Time `json:"updated_at" "bson:"updated_at""`
}

type SignInInput struct {
	Email    string `json:"email bson:"email" binding:"requered'"`
	Password string `json:"password bson:"password" binding:"requered'"`
}

type DBResponse struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Name            string             `json:"name" bson:"name"`
	Email           string             `json:"email" bson:"email"`
	Password        string             `json:"password" bson:"password"`
	PasswordConfirm string             `json:"passwordConfirm,omitempty" bson:"passwordConfirm,omitempty"`
	Role            string             `json:"role" bson:"role"`
	Verified        bool               `json:"verified" bson:"verified"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
}
