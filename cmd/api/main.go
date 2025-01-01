package main

import (
	"context"
	"explorer/internal/adapters/api"     // Import API package for route handlers
	"explorer/internal/adapters/data"    // Import data package for interacting with the MBTA API
	"explorer/internal/config/apiConfig" // Import config package for loading and handling configuration (API keys)
	"explorer/internal/config/memoryConfig"
	"explorer/internal/stream"
	"explorer/internal/usecases" // Import usecases package for business logic
	"log"                        // Import log package for logging messages
	"net/http"                   // Import net/http package to start the web server and handle HTTP requests
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"   // Import the Gorilla Mux router for routing HTTP requests
	"github.com/joho/godotenv" // Import the godotenv package to load environment variables from .env file
	"github.com/rs/cors"
)

func main() {
	// Log that the server is starting
	log.Println("Starting server...")

	// Initialize a new Gorilla Mux router
	r := mux.NewRouter()

	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		// If loading the .env file fails, log an error and terminate the program
		log.Fatal("Error loading .env file")
	}

	// Retrieve the API key from the environment using the config package
	key := apiConfig.GetAPIKey()

	cache := memoryConfig.MemcachedConfig()

	// Initialize the use case layer by creating a FetchData instance with the MBTA client
	fetchData := usecases.NewFetchData(data.NewMBTAClient(key), cache)

	sm := stream.MBTAStreamManager

	// Context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start streaming MBTA data
	go func() {
		if key == "" {
			log.Fatal("MBTA_API_KEY environment variable not set")
		}

		url := "https://api-v3.mbta.com/vehicles?filter[route]=Red,Orange,Blue,Green-B,Green-C,Green-D,Green-E,Mattapan"
		sm.StartStreaming(ctx, url, key)
	}()

	// Handle shutdown signals
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		cancel()
		time.Sleep(1 * time.Second) // Give goroutines time to clean up
		os.Exit(0)
	}()

	// Set up API routes and map them to corresponding handler functions
	r.HandleFunc("/api/stops", api.FetchRouteStops(fetchData)).Methods("GET")   // Route for fetching stops
	r.HandleFunc("/api/shapes", api.FetchRouteShapes(fetchData)).Methods("GET") // Route for fetching shapes
	r.HandleFunc("/api/live", api.UpdateLivePosition(fetchData)).Methods("GET") // Route for fetching live data (vehicle locations)
	r.HandleFunc("/api/routes", api.FetchRoutes(fetchData)).Methods("GET")
	r.HandleFunc("/stream/vehicles", api.StreamVehiclesHandler)

	// Configure CORS options
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},         // Allow frontend's origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},  // Allowed HTTP methods
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,                                      // Allow credentials (e.g., cookies, authorization headers)
	})

	// Apply CORS middleware to the router
	handler := corsOptions.Handler(r)

	// Start the server on port 8080 and log any errors that occur
	log.Printf("Server is listening on port %s...\n", "8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
