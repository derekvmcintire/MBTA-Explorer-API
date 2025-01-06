package stream

import "log"

// Broadcast sends the given data to all connected clients.
// This method ensures thread-safe access to the client map and skips clients
// that are unable to keep up with the data flow.
func (sm *StreamManager) Broadcast(data string) {
	log.Println("Broadcast has been called")
	// Acquire the mutex lock to safely access the client map.
	sm.clientsMutex.Lock()
	defer sm.clientsMutex.Unlock() // Ensure the mutex is unlocked after the operation.

	// Iterate over all registered clients.
	for client := range sm.clients {
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
