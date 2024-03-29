package ratelimit

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// Default: 5 req per sec
const (
	RedisDefaultLimit      int64 = 5
	RedisDefaultWindowSize int64 = 1
)

type Option func(*RateConfig) error

// singleton db connection
var (
	rdb  *redis.Client
	once sync.Once
)

// Desired Rate = limit / windowSize (eg 5 requests per second)
type RateConfig struct {
	Id         string // api key
	Limit      int64  // requests allowed in window
	WindowSize int64  // length of window in milliseconds
}

func WithLimit(limit int64) func(*RateConfig) error {
	return func(rc *RateConfig) error {
		rc.Limit = limit
		return nil
	}
}

func Init(address string) error {
	opts, err := redis.ParseURL(address)
	if err != nil {
		return err
	}
	once.Do(func() {
		rdb = redis.NewClient(opts)
	})
	return err
}

func WithWindowSize(windowSize int64) func(*RateConfig) error {
	return func(rc *RateConfig) error {
		rc.WindowSize = windowSize
		return nil
	}
}

func NewRateLimterConfig(id string, opts ...Option) (RateConfig, error) {
	rc := RateConfig{
		Id:         id,
		Limit:      RedisDefaultLimit,
		WindowSize: RedisDefaultWindowSize,
	}
	for _, opt := range opts {
		if err := opt(&rc); err != nil {
			return RateConfig{}, fmt.Errorf("option failed %w", err)
		}
	}
	return rc, nil
}

// Sliding window log
func (rc *RateConfig) RateLimit(ctx context.Context) bool {
	key := fmt.Sprintf("rate_limit:%s", rc.Id)
	now := time.Now().UnixMilli()

	// purge reqs before window
	rdb.ZRemRangeByScore(ctx, key, "-inf", fmt.Sprintf("%d", now-(1000*rc.WindowSize)))

	// count number of reqs in curr window
	current, err := rdb.ZCard(ctx, key).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return true
	}
	if current >= rc.Limit {
		// rate limited
		return true
	}

	// req successful, add it to curr window
	rdb.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now),
		Member: now,
	})
	// whole set should expire after window size
	rdb.Expire(ctx, key, time.Duration(rc.WindowSize)*time.Second)
	return false
}
