package stream

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitializeStream(apiKey string) context.CancelFunc {
	// Initialize the stream manager
	sm := MBTAStreamManager

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())

	// Start streaming MBTA data
	go func() {
		if apiKey == "" {
			log.Fatal("MBTA_API_KEY environment variable not set")
		}

		sm.StartStreaming(ctx, apiKey)
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

	// Return the cancel function to allow manual shutdown if needed
	return cancel
}
