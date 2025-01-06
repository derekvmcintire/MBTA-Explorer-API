package core

import (
	"context"
)

type MbtaStreamManager interface {
	Start(ctx context.Context, apiKey string)
	AddClient(client chan string)
	RemoveClient(client chan string)
	Stop()
	cancelFunc() context.CancelFunc
}
