package stream

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type StreamManager struct {
	clients      map[chan string]struct{} // Map to store connected client channels
	clientsMutex sync.Mutex               // Mutex to safely update the clients map
	stop         chan struct{}            // Channel to signal when to stop streaming
	cancelFunc   context.CancelFunc
}

var streamOnce sync.Once

func (sm *StreamManager) EnsureStreaming(url, apiKey string) {
	streamOnce.Do(func() {
		log.Println("Ensuring streaming is started...")
		// Create a new context with cancellation support
		ctx, cancel := context.WithCancel(context.Background())
		sm.cancelFunc = cancel // Store the cancel function

		// Start streaming MBTA data in a separate goroutine
		go func() {
			// Check if the API key is provided, if not, log an error and stop
			if apiKey == "" {
				log.Fatal("MBTA_API_KEY environment variable not set")
			}

			// Start the MBTA stream with the provided URL and API key
			sm.Start(ctx, url, apiKey)
		}()

		// Handle system shutdown signals (e.g., SIGINT, SIGTERM) in a separate goroutine
		go func() {
			// Create a channel to receive shutdown signals
			sigChan := make(chan os.Signal, 1)
			// Notify the channel for interrupt (Ctrl+C) or termination signals
			signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
			// Wait for a signal to shut down the stream
			<-sigChan
			// Cancel the context to stop the streaming process
			cancel()
			// Allow a brief moment for goroutines to clean up
			time.Sleep(1 * time.Second)
			// Exit the program
			os.Exit(0)
		}()
	})
}
