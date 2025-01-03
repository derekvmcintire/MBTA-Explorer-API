package core

import (
	"context"
)

type StreamManager interface {
	StartStreaming(ctx context.Context, apiKey string)
	AddClient(client chan string)
	RemoveClient(client chan string)
	Stop()
}
