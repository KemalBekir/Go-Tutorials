package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/services"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func ExtractUserIDFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("my jwt secret"), nil
	})
	if err != nil || !token.Valid {
		return "", errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to parse claims")
	}

	userID, ok := claims["_id"].(string)
	if !ok {
		return "", errors.New("user ID not found in token claims")
	}

	return userID, nil
}

func AuthMiddleware(userCollection *mongo.Collection) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("x-authorization")

			if token != "" {
				userData, err := services.VerifySession(token, userCollection)
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					json.NewEncoder(w).Encode(map[string]string{"message": "Invalid access token. Please sign in"})
					return
				}
				// Attach user data to request context
				ctx := context.WithValue(r.Context(), "user", userData)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"message": "Token missing"})
			}
		})
	}
}
