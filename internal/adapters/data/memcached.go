package data

import "github.com/bradfitz/gomemcache/memcache"

// CacheClient wraps the memcache.Client and provides higher-level methods
// to interact with a Memcached server.
type CacheClient struct {
	client *memcache.Client // Underlying Memcached client
}

// NewCacheClient creates and initializes a new CacheClient instance.
//
// Parameters:
// - server: Address of the Memcached server (e.g., "localhost:11211").
//
// Returns:
// - A pointer to a new CacheClient instance.
func NewCacheClient(server string) *CacheClient {
	return &CacheClient{
		client: memcache.New(server), // Initialize the Memcached client with the server address
	}
}

// Get retrieves a value from the Memcached server for the given key.
//
// Parameters:
// - key: The key to look up in the cache.
//
// Returns:
//   - The value associated with the key as a string, or an error if the key does not exist
//     or an issue occurred during retrieval.
func (c *CacheClient) Get(key string) (string, error) {
	item, err := c.client.Get(key) // Fetch the item from Memcached
	if err != nil {
		return "", err // Return an empty string and propagate the error
	}
	return string(item.Value), nil // Convert the byte array to a string and return
}

// Set stores a key-value pair in the Memcached server.
//
// Parameters:
// - key: The key under which the value will be stored.
// - value: The value to store in the cache.
//
// Returns:
// - An error if the key-value pair could not be stored, or nil on success.
func (c *CacheClient) Set(key, value string) error {
	return c.client.Set(&memcache.Item{
		Key:   key,           // Set the key
		Value: []byte(value), // Convert the string value to a byte array
	})
}
