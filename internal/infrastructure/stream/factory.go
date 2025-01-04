package stream

// NewStreamManager initializes and returns a new StreamManager instance.
func NewStreamManager() *StreamManager {
	return &StreamManager{
		clients: make(map[chan string]struct{}),
		stop:    make(chan struct{}),
	}
}

// Global instance of MBTAStreamManager
var MBTAStreamManager = NewStreamManager()
