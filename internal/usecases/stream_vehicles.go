package usecases

import (
	"context"
	"explorer/internal/infrastructure/stream"
)

type StreamVehiclesUseCase struct {
	streamManager *stream.StreamManager
}

func NewStreamVehiclesUseCase(sm *stream.StreamManager) *StreamVehiclesUseCase {
	return &StreamVehiclesUseCase{
		streamManager: sm,
	}
}

// StreamSetup initializes the stream and returns a client channel
func (uc *StreamVehiclesUseCase) StreamSetup(url, apiKey string) chan string {
	// Ensure the stream is running
	uc.streamManager.EnsureStreaming(url, apiKey)

	// Create and register client channel
	clientChan := make(chan string, 100)
	uc.streamManager.AddClient(clientChan)

	return clientChan
}

// HandleDisconnect sets up disconnection handling for a client
func (uc *StreamVehiclesUseCase) HandleDisconnect(ctx context.Context, clientChan chan string) {
	go func() {
		<-ctx.Done()
		uc.streamManager.RemoveClient(clientChan)
	}()
}
