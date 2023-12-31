package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
  lock sync.RWMutex
  data map[string][]byte
}

var _ ICacher = (*Cache)(nil)

func NewInMemoryCache() *Cache {
  return &Cache{
    data: make(map[string][]byte),
  }
}

func (c *Cache) Delete(key []byte) error {
  c.lock.Lock()
  defer c.lock.Unlock()

  delete(c.data, string(key))

  return nil
}

func (c *Cache) Has(key []byte) bool {
  c.lock.RLock()
  defer c.lock.RUnlock()

  _, ok := c.data[string(key)] 

  return ok

}

func (c *Cache) Get(key []byte) ([]byte, error) {
  c.lock.RLock()
  defer c.lock.RUnlock()

  strKey := string(key)

  val, ok := c.data[strKey] 
  if !ok {
    return nil, fmt.Errorf("key (%s) not found", strKey)
  }

  return val, nil
}

func (c *Cache) Set(key, value []byte, ttl time.Duration) error {
  c.lock.Lock()
  defer c.lock.Unlock()

  c.data[string(key)] = value

  return nil
}

