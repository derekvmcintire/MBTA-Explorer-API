package usecases

import (
	"context"
)

type StreamVehiclesUseCase struct {
	streamManager *StreamManagerUseCase
}

func NewStreamVehiclesUseCase(sm *StreamManagerUseCase) *StreamVehiclesUseCase {
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
	uc.streamManager.Distributor.AddClient(clientChan)

	return clientChan
}

// HandleDisconnect sets up disconnection handling for a client
func (uc *StreamVehiclesUseCase) HandleDisconnect(ctx context.Context, clientChan chan string) {
	go func() {
		<-ctx.Done()
		uc.streamManager.Distributor.RemoveClient(clientChan)
	}()
}
