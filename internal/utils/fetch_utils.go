package utils

import (
	"fmt"
	"net/http"
	"time"
)

func FetchDataWithRetry(url string, maxRetries int) ([]byte, error) {
	var err error
	for i := 0; i < maxRetries; i++ {
		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			return io.ReadAll(resp.Body)
		}
		time.Sleep(2 * time.Second) // retry after a delay
	}
	return nil, fmt.Errorf("failed to fetch data after %d retries", maxRetries)
}
