package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
	done    chan struct{}
}

func (cache *Cache) reapLoop(interval time.Duration) {
	// create a ticker that will "tick" every interval duration
	// uses channels under the hood
	ticker := time.NewTicker(interval)

	// start a go routine that will run concurrently with the
	// rest of the program
	go func() {
		// std inf loop
		for {
			select {
			// wait for the ticker channle to recieve a value
			// happens every tick
			case <-ticker.C:
				// lock the mutex to safely access teh map
				// prevents a race
				cache.mu.Lock()
				// look at all entries in the cache
				for mkey, entry := range cache.entries {
					// remove stale ones
					if time.Since(entry.createdAt) > interval {
						delete(cache.entries, mkey)
					}
				}
				cache.mu.Unlock()

			case <-cache.done:
				ticker.Stop()
				return
			}
		}
	}()
}

func (cache *Cache) Add(key string, val []byte) {
	createdAt := time.Now()

	//lock the mutex
	cache.mu.Lock()
	defer cache.mu.Unlock()

	// make the change
	cache.entries[key] = cacheEntry{
		createdAt: createdAt,
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, exists := cache.entries[key]

	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (cache *Cache) Close() {
	close(cache.done)
}

func NewCache(interval time.Duration) *Cache {

	cache := Cache{
		entries: make(map[string]cacheEntry),
		done:    make(chan struct{}),
	}

	cache.reapLoop(interval)
	return &cache
}
