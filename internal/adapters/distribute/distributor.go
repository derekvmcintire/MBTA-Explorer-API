package distribute

import (
	"log"
	"sync"
)

type ClientDistributor struct {
	clients      map[chan string]struct{}
	clientsMutex sync.Mutex
	stop         chan struct{} // Channel to signal when to stop streaming
}

func NewClientDistributor() *ClientDistributor {
	return &ClientDistributor{
		clients: make(map[chan string]struct{}),
		stop:    make(chan struct{}),
	}
}

// Broadcast sends the given data to all connected clients.
// This method ensures thread-safe access to the client map and skips clients
// that are unable to keep up with the data flow.
func (cd *ClientDistributor) Broadcast(data string) {
	// Acquire the mutex lock to safely access the client map.
	cd.clientsMutex.Lock()
	defer cd.clientsMutex.Unlock() // Ensure the mutex is unlocked after the operation.

	// Iterate over all registered clients.
	for client := range cd.clients {
		select {
		case client <- data:
			// Attempt to send data to the client's channel.
			// If the client is ready to receive, this operation succeeds immediately.
		default:
			// Skip clients whose channels are full or unresponsive.
			// Log a message to indicate the client was skipped due to being slow.
			log.Println("Stream Manager Client is slow, skipping...")
		}
	}
}

// AddClient adds a new client channel to the manager to receive data updates.
// It locks the client list to ensure thread safety during modification.
func (cd *ClientDistributor) AddClient(client chan string) {
	cd.clientsMutex.Lock()          // Lock to ensure safe access to the clients map
	defer cd.clientsMutex.Unlock()  // Unlock once the operation is done
	cd.clients[client] = struct{}{} // Add the client channel to the map of active clients
}

// RemoveClient removes a client channel when they disconnect.
// It locks the client list to ensure thread safety during modification.
func (cd *ClientDistributor) RemoveClient(client chan string) {
	cd.clientsMutex.Lock()         // Lock to ensure safe access to the clients map
	defer cd.clientsMutex.Unlock() // Unlock once the operation is done
	// Check if the client exists in the map
	if _, ok := cd.clients[client]; ok {
		// Remove the client from the map and close the channel to signal disconnection
		delete(cd.clients, client)
		close(client)
	}
}

// Stop stops the stream manager and signals all processes to stop.
// It closes the stop channel to initiate the shutdown process.
func (cd *ClientDistributor) Stop() {
	// Close the stop channel to signal all goroutines to stop
	close(cd.stop)
}
