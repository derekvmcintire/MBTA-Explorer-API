package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// CORS Middleware Configuration
func SetCorsHandler(router *mux.Router) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)
}
