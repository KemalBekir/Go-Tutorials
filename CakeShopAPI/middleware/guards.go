package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/KemalBekir/Go-Tutorials/CakeShopAPI/models"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.User)
		if user.ID != "" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Please log in"})
		}
	})
}

func IsGuest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.User)
		if user.ID == "" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "You are already signed in"})
		}
	})
}

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.User)
		if user.Role == "admin" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "You cannot create a new record"})
		}
	})
}

func IsOwner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(models.User)
		cakeOwnerID := r.Context().Value("cake.owner._id").(string) // Assuming this is the way you're accessing the cake owner's ID

		if user.ID == cakeOwnerID && user.Role == "admin" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"message": "You cannot modify this record"})
		}
	})
}
