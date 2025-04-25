package ratelimiter

import "time"

type RateLimiter struct {
	FastBucket *TokenBucket
	SlowBucket *TokenBucket
}

func (rl *RateLimiter) ConsumeToken() {
	rl.FastBucket.consumeToken()
	rl.SlowBucket.consumeToken()
}

func (rl *RateLimiter) IsLimited() bool {
	return rl.FastBucket.IsEmpty() || rl.SlowBucket.IsEmpty()
}

func NewRateLimiter() *RateLimiter {
	fb := NewTokenBucket(19, 19, time.Second*1)
	sb := NewTokenBucket(100, 100, time.Minute*2)

	return &RateLimiter{
		FastBucket: fb,
		SlowBucket: sb,
	}
}
