package main

import (
	"explorer/internal/adapters/data"    // Import data package for interacting with the MBTA API
	"explorer/internal/config/apiConfig" // Import config package for loading and handling configuration (API keys)
	"explorer/internal/config/memoryConfig"
	"explorer/internal/middleware"
	"explorer/internal/routes"
	"explorer/internal/stream"
	"explorer/internal/usecases" // Import usecases package for business logic
	"log"                        // Import log package for logging messages
	"net/http"                   // Import net/http package to start the web server and handle HTTP requests

	"github.com/gorilla/mux"   // Import the Gorilla Mux router for routing HTTP requests
	"github.com/joho/godotenv" // Import the godotenv package to load environment variables from .env file
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		// If loading the .env file fails, log an error and terminate the program
		log.Fatal("Error loading .env file")
	}

	// Retrieve the API key from the environment using the config package
	key := apiConfig.GetAPIKey()

	// Initialize the memcached client
	cache := memoryConfig.MemcachedConfig()

	// Initialize streaming of MBTA data
	cancelStream := stream.InitializeStream(key)
	defer cancelStream() // Ensure the stream is canceled on application exit

	// Initialize the use case layer by creating an mbtaApiHelper instance with the MBTA client
	mbtaApiHelper := usecases.NewMbtaApiHelper(data.NewMBTAClient(key), cache)

	// Initialize a new Gorilla Mux router
	r := mux.NewRouter()

	// Register the routes with the router
	routes.RegisterRoutes(r, mbtaApiHelper)

	// Configure CORS
	corsHandler := middleware.SetCorsHandler(r)

	// Start the server on port 8080 and log any errors that occur
	log.Printf("Server is listening on port %s...\n", "8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
