package stream

import "sync"

type StreamManager struct {
	clients      map[chan string]struct{} // Map to store connected client channels
	clientsMutex sync.Mutex               // Mutex to safely update the clients map
	stop         chan struct{}            // Channel to signal when to stop streaming
}

// NewStreamManager initializes and returns a new StreamManager instance.
func NewStreamManager() *StreamManager {
	return &StreamManager{
		clients: make(map[chan string]struct{}),
		stop:    make(chan struct{}),
	}
}

// Global instance of MBTAStreamManager
var MBTAStreamManager = NewStreamManager()

// "https://api-v3.mbta.com/vehicles?filter[route]=Red,Orange,Blue,Green-B,Green-C,Green-D,Green-E,Mattapan"
