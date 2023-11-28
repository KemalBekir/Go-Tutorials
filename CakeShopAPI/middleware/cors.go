package middleware

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set allowed origin (replace this with your specific origin)
		w.Header().Set("Access-Control-Allow-Origin", "https://cakeshop-ej4v.onrender.com")
		// Set allowed methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, HEAD, PUT, OPTIONS")
		// Set allowed headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
