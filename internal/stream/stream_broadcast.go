package stream

import "log"

// Broadcast sends the given data to all connected clients.
func (sm *StreamManager) Broadcast(data string) {
	sm.clientsMutex.Lock()
	defer sm.clientsMutex.Unlock()
	for client := range sm.clients {
		select {
		case client <- data: // Send data to the client channel
		default: // Skip clients that are too slow to keep up
			log.Println("Client is slow, skipping...")
		}
	}
}
