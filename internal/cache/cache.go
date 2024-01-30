package cache

import (
	"sync"
	"time"

	"github.com/brunocapri/movie-engine/internal/domain"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	movies    []domain.Movie
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache:    make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c Cache) Add(key string, movies []domain.Movie) {
	c.mu.Lock()
	c.cache[key] = cacheEntry{
		movies:    movies,
		createdAt: time.Now(),
	}
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]domain.Movie, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.cache[key]
	return val.movies, ok
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for k, val := range c.cache {
			if time.Now().After(val.createdAt.Add(c.interval)) {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
	}
}
