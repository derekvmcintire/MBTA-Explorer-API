package memoryConfig

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func MemcachedConfig() *memcache.Client {
	// Connect to the Memcached server
	mc := memcache.New("127.0.0.1:11211")

	// Test connection
	err := mc.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to Memcached: %v", err)
	}

	log.Println("Connected to Memcached!")
	return mc
}
