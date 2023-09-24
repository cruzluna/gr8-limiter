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

func rateLimit(rdb *redis.Client, ip string, limit int64, windowSize int64) bool {
	key := fmt.Sprintf("rate_limit:%s", ip)
	now := time.Now().UnixNano() / int64(time.Millisecond)

	rdb.ZRemRangeByScore(ctx, key, "-inf", fmt.Sprintf("%d", now-(1000*windowSize)))

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
	ip := "123.45.67.89"
	windowSize := int64(1) // seconds
	limit := int64(5)      // per window

	for i := 1; i <= 7; i++ {
		allowed := rateLimit(rdb, ip, limit, windowSize)
		if allowed {
			fmt.Printf("Request %d allowed.\n", i)
		} else {
			fmt.Printf("Request %d blocked.\n", i)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
