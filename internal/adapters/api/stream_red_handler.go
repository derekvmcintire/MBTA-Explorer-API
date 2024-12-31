package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// streamVehiclesHandler proxies the MBTA streaming data to the frontend
func StreamVehiclesHandler(w http.ResponseWriter, _ *http.Request) {
	apiKey := os.Getenv("MBTA_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key not configured", http.StatusInternalServerError)
		return
	}

	// Set the required headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Forward the stream from the MBTA API to the client
	url := fmt.Sprintf("https://api-v3.mbta.com/vehicles?filter[route]=Red&api_key=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error connecting to MBTA API: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Ensure the MBTA API response is an event stream
	if resp.Header.Get("Content-Type") != "text/event-stream" {
		http.Error(w, "Invalid response from MBTA API", http.StatusInternalServerError)
		return
	}

	// Stream the MBTA data to the client
	_, err = w.Write([]byte(": Streaming MBTA data\n\n")) // Initial comment to keep connection alive
	if err != nil {
		log.Printf("Error writing to client: %v", err)
		return
	}

	buffer := make([]byte, 1024)
	for {
		// Read from MBTA API response
		n, readErr := resp.Body.Read(buffer)
		if n > 0 {
			// Write data to the frontend client
			_, writeErr := w.Write(buffer[:n])
			if writeErr != nil {
				log.Printf("Error writing to client: %v", writeErr)
				break
			}
			w.(http.Flusher).Flush() // Ensure data is sent immediately
		}
		if readErr != nil {
			log.Printf("Error reading from MBTA API: %v", readErr)
			break
		}
	}
}
