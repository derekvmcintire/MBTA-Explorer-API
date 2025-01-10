// ports/streaming.go
package ports

import "context"

// StreamSource defines how to interact with an external streaming data source
type StreamSource interface {
	Start(ctx context.Context, url, apiKey string)
}

// StreamDistributor defines how to manage client connections and data distribution
type StreamDistributor interface {
	AddClient(client chan string)
	RemoveClient(client chan string)
	Broadcast(data string)
	Stop()
}

// StreamManager combines both source and distribution capabilities
type StreamManager interface {
	StreamSource
	StreamDistributor
}
