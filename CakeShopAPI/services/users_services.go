package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserCollection struct {
	Collection *mongo.Collection
}

type User struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username       string             `json:"username,omitempty" bson:"username,omitempty"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty"`
	HashedPassword string             `json:"hashedPassword,omitempty" bson:"hashedPassword,omitempty"`
	Role           string             `json:"role,omitempty" bson:"role,omitempty"`
}

var blacklist []string

func NewUserCollection(collection *mongo.Collection) *UserCollection {
	return &UserCollection{
		Collection: collection,
	}
}

func (s *UserCollection) Register(username, email, password string) (map[string]interface{}, error) {
	existingUser := User{}
	err := s.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&existingUser)
	if err == nil {
		// Email already exists
		return nil, errors.New("email already exists")
	} else if err != nil && err != mongo.ErrNoDocuments {
		// Handle other errors
		log.Printf("Error finding existing user: %s", err)
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating hashed password: %s", err)
		return nil, err
	}

	newUser := User{
		Username:       username,
		Email:          email,
		HashedPassword: string(hashedPassword),
	}

	result, err := s.Collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Printf("Error inserting new user: %s", err)
		return nil, err
	}

	userID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("Failed to get user ID")
		return nil, errors.New("failed to get user ID")
	}

	user := &User{
		ID:    userID,
		Email: email,
		Role:  "user",
	}

	return CreateSession(user), nil
}

func (s *UserCollection) Login(email, password string) (map[string]interface{}, error) {
	user := User{}
	err := s.Collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.New("incorrect email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return nil, errors.New("incorrect email or password")
	}

	return CreateSession(&user), nil
}

func Logout(token string) {
	blacklist = append(blacklist, token)
}

func CreateSession(user *User) map[string]interface{} {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"_id":   user.ID.Hex(),
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("secret"))

	return map[string]interface{}{
		"email":       user.Email,
		"_id":         user.ID.Hex(),
		"accessToken": tokenString,
	}
}

func VerifySession(token string, userCollection *mongo.Collection) (map[string]interface{}, error) {
	if containsToken(token) {
		return nil, errors.New("token is invalidated")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil // Replace "your-secret-key" with your actual secret key
	})

	if err != nil || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}

	return map[string]interface{}{
		"email": claims["email"],
		"_id":   claims["_id"],
		"role":  claims["role"],
		"token": token,
	}, nil
}

func containsToken(token string) bool {
	for _, t := range blacklist {
		if t == token {
			return true
		}
	}
	return false
}
