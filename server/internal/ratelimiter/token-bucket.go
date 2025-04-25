package ratelimiter

import (
	"math"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens     int
	capacity   int
	ticker     *time.Ticker
	lastRefill time.Time

	mu sync.Mutex
}

func (tb *TokenBucket) refill() {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.lastRefill = time.Now()
	tb.tokens = tb.capacity
}

func (tb *TokenBucket) consumeToken() {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.tokens = int(math.Max(float64(tb.tokens-1), 0))
}

func (tb *TokenBucket) manage() {
	for {
		select {
		case <-tb.ticker.C:
			tb.refill()
		}
	}
}

func (tb *TokenBucket) IsEmpty() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	return tb.tokens == 0
}

func NewTokenBucket(capacity, tokens int, refillRate time.Duration) *TokenBucket {
	tb := &TokenBucket{
		tokens:   tokens,
		capacity: capacity,
		ticker:   time.NewTicker(refillRate),
	}

	go tb.manage()

	return tb
}
