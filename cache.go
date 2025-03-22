package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type CacheData struct {
	cache map[string][]Data
	size  int
	mu    sync.Mutex
}
type CacheResult struct {
	err     error
	success bool
}

func createCache() *CacheData {
	return &CacheData{
		cache: make(map[string][]Data),
		size:  0,
	}
}
func (c *CacheData) addData(url string, data []Data, resultChan chan<- CacheResult) {
	go func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		if _, ok := c.cache[url]; ok {
			resultChan <- CacheResult{err: errors.New("url already exists in cache"), success: false}
		}
		c.cache[url] = data
		c.size += 1
		resultChan <- CacheResult{err: nil, success: true}
	}()

}
func (c *CacheData) deleteData() {

	go func() {
		time.Sleep(5 * time.Millisecond)
		c.mu.Lock()
		c.cache = make(map[string][]Data)
		c.mu.Unlock()
		fmt.Println("Cache cleared after 5 minutes.")
	}()
}
func (c *CacheData) printSize() {
	fmt.Println(c.size)
}

var my_cache = createCache()
