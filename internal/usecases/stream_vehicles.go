package usecases

import (
	"context"
	"errors"
	"explorer/internal/infrastructure"
	"io"
	"log"
	"net/http"
)

func StreamVehicles(ctx context.Context, apiKey string, w http.ResponseWriter) error {
	// Build the MBTA API request
	url := "https://api-v3.mbta.com/vehicles?filter[route]=Red,Orange,Blue,Green-B,Green-C,Green-D,Green-E,Mattapan"
	resp, err := infrastructure.FetchStream(ctx, url, apiKey)
	if err != nil {
		log.Printf("Error fetching stream: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Validate the content type
	if resp.Header.Get("Content-Type") != "text/event-stream" {
		log.Println("Content-Type is: ", resp.Header.Get("Content-Type"))
		return errors.New("Invalid response content type")
	}

	log.Println("******** Streaming data to client ***********")

	// Stream data to the client
	flusher := w.(http.Flusher)
	flusher.Flush()

	if _, err := io.Copy(w, resp.Body); err != nil {
		if err == io.EOF {
			log.Println("Stream ended normally.")
		} else {
			log.Printf("Error streaming data: %v", err)
		}
		return err
	}

	log.Println("******** Stream ended ***********")
	return nil
}
