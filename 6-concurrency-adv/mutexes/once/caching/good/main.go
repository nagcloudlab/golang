package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var numberOfRequests int32

// test the program using the -race flag
// cd into caching/good
// go run -race main.go
func main() {
	c := client{
		cache: map[string]*cacheEntry{},
		mu:    new(sync.Mutex),
	}

	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			delay := time.Duration(r.Intn(500)) * time.Millisecond
			c.httpCall("req1", delay)
		}()
	}

	wg.Wait()
	fmt.Println("number of requests:", numberOfRequests)
}

type cacheEntry struct {
	data []byte
	once *sync.Once
}

type client struct {
	cache map[string]*cacheEntry
	mu    *sync.Mutex
}

func (c *client) httpCall(key string, delay time.Duration) []byte {
	c.mu.Lock()
	entry, found := c.cache[key]
	if !found {
		entry = &cacheEntry{once: new(sync.Once)}
		c.cache[key] = entry
	}
	c.mu.Unlock()

	entry.once.Do(func() {
		entry.data = externalAPI(key, delay)
	})
	return entry.data
}

func externalAPI(key string, delay time.Duration) []byte {
	time.Sleep(delay)
	atomic.AddInt32(&numberOfRequests, 1)
	// simulate 5KB data
	s := strings.Repeat("a", 1024)
	_ = []string{s, s, s, s, s}
	return []byte("response: " + key)
}
