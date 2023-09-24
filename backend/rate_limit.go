package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB:   0,
})

var ctx = context.Background()

type Request struct {
	Kind string
}

type Config struct {
	WindowSize int64 // length of window in milliseconds
	Limit      int64 // requests allowed in window
}

func rateLimiter(requests chan Request, configs chan Config, responses chan bool) {
	currentConfig := Config{
		WindowSize: 500,
		Limit:      2,
	}
	for {
		select {
		case req := <-requests:
			success := rateLimit(rdb, req.Kind, currentConfig.WindowSize, currentConfig.Limit)
			responses <- success
		case config := <-configs:
			currentConfig = config
		}
	}
}

func rateLimit(rdb *redis.Client, ip string, windowSize int64, limit int64) bool {
	key := fmt.Sprintf("rate_limit:%s", ip)
	now := time.Now().UnixNano() / int64(time.Millisecond)

	rdb.ZRemRangeByScore(ctx, key, "-inf", fmt.Sprintf("%d", now-windowSize))

	current, err := rdb.ZCard(ctx, key).Result()
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	if current >= limit {
		return false
	}

	rdb.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now),
		Member: now,
	})
	rdb.Expire(ctx, key, time.Duration(windowSize)*time.Second)
	return true
}

func test() {
	requests := make(chan Request)
	configs := make(chan Config)
	responses := make(chan bool)

	go rateLimiter(requests, configs, responses)

	for i := 0; i < 10; i++ {
		r := Request{Kind: "glorb"}
		requests <- r

		result := <-responses
		canProceed := result

		if canProceed {
			fmt.Println("You can proceed")
		} else {
			fmt.Println("You cannot proceed")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
