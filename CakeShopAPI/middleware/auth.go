package middleware

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
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
