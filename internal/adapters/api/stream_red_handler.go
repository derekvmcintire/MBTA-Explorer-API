package api

import (
	"io"
	"log"
	"net/http"
	"os"
)

func StreamVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("******** Starting stream ***********")

	apiKey := os.Getenv("MBTA_API_KEY")
	if apiKey == "" {
		http.Error(w, "API key not configured", http.StatusInternalServerError)
		return
	}

	// Set the required headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Build the MBTA API request
	url := "https://api-v3.mbta.com/vehicles?filter[route]=Red"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "Failed to create request to MBTA API", http.StatusInternalServerError)
		log.Printf("Failed to create request: %v", err)
		return
	}

	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("X-API-Key", apiKey)

	// Use an HTTP client to perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to connect to MBTA API", http.StatusInternalServerError)
		log.Printf("Failed to connect to MBTA API: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.Header.Get("Content-Type") != "text/event-stream" {
		log.Println("Content-Type is: ", resp.Header.Get("Content-Type"))
		log.Println("******** MBTA Response not a stream ***********")
		http.Error(w, "Invalid response from MBTA API", http.StatusInternalServerError)
		return
	}

	log.Println("******** Streaming data to client ***********")

	// Use http.CloseNotifier to handle client disconnects
	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		log.Println("Client closed connection")
		resp.Body.Close()
	}()

	flusher := w.(http.Flusher)
	flusher.Flush()

	// Stream data from the MBTA API to the client
	if _, err := io.Copy(w, resp.Body); err != nil {
		if err == io.EOF {
			log.Println("Stream ended normally.")
		} else {
			log.Printf("Error streaming data: %v", err)
		}
	}

	log.Println("******** Stream ended ***********")
}
