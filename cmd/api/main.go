package main

import (
	"explorer/internal/adapters/data"
	apiHttp "explorer/internal/adapters/http"
	"explorer/internal/infrastructure/config"
	"explorer/internal/infrastructure/middleware"
	"explorer/internal/infrastructure/stream"
	"explorer/internal/usecases"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		// If loading the .env file fails, log an error and terminate the program
		log.Fatal("Error loading .env file")
	}

	// Retrieve the API key from the environment using the config package
	key := config.GetAPIKey()

	// Initialize the memcached client
	cache := config.MemcachedConfig()

	// Initialize the use case layer by creating an mbtaApiHelper instance with the MBTA client
	mbtaApiHelper := usecases.NewMbtaApiHelper(data.NewMBTAClient(key), cache)

	// Initialize a new Gorilla Mux router
	r := mux.NewRouter()

	sm := stream.NewStreamManager()

	// Register the routes with the router
	apiHttp.RegisterRoutes(r, mbtaApiHelper, sm)

	// Configure CORS
	corsHandler := middleware.SetCorsHandler(r)

	// Start the server on port 8080 and log any errors that occur
	log.Printf("Server is listening on port %s...\n", "8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
