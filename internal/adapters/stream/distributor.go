// adapters/stream/distributor.go
package stream

import "sync"

type ClientDistributor struct {
	clients      map[chan string]struct{}
	clientsMutex sync.Mutex
}

func NewClientDistributor() *ClientDistributor {
	return &ClientDistributor{
		clients: make(map[chan string]struct{}),
	}
}
