package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Desired Rate = limit / windowSize (eg 5 requests per second)
type RateConfig struct {
	id         string
	limit      int64 // requests allowed in window
	windowSize int64 // length of window in milliseconds
}

func NewRateLimterConfig(id string, limit int64, windowSize int64) *RateConfig {
	return &RateConfig{
		id:         id, // unique identifier for now
		limit:      limit,
		windowSize: windowSize,
	}
}

// Sliding window log
func (rc RateConfig) rateLimit(ctx context.Context, rdb *redis.Client) bool {
	key := fmt.Sprintf("rate_limit:%s", rc.id)
	now := time.Now().UnixMilli()

	// purge reqs before window
	rdb.ZRemRangeByScore(ctx, key, "-inf", fmt.Sprintf("%d", now-(1000*rc.windowSize)))

	// count number of reqs in curr window
	current, err := rdb.ZCard(ctx, key).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	if current >= rc.limit {
		// rate limited
		return false
	}

	// req successful, add it to curr window
	rdb.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now),
		Member: now,
	})
	// whole set should expire after window size
	rdb.Expire(ctx, key, time.Duration(rc.windowSize)*time.Second)
	return true
}
