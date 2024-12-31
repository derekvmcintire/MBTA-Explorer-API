package memory

import "github.com/bradfitz/gomemcache/memcache"

type CacheClient struct {
	client *memcache.Client
}

func NewCacheClient(server string) *CacheClient {
	return &CacheClient{
		client: memcache.New(server),
	}
}

func (c *CacheClient) Get(key string) (string, error) {
	item, err := c.client.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

func (c *CacheClient) Set(key, value string) error {
	return c.client.Set(&memcache.Item{
		Key:   key,
		Value: []byte(value),
	})
}
