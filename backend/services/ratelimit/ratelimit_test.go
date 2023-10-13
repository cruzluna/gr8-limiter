package ratelimit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultRateConfig(t *testing.T) {
	dummyId := "dummyId"
	rc, _ := NewRateLimterConfig(dummyId)

	expect := RateConfig{
		Id:         dummyId,
		Limit:      RedisDefaultLimit,
		WindowSize: RedisDefaultWindowSize,
	}

	assert.Equal(t, rc, expect)

	rc, _ = NewRateLimterConfig(dummyId, WithLimit(10))

	assert.Equal(t, rc, RateConfig{
		Id:         dummyId,
		Limit:      10,
		WindowSize: RedisDefaultWindowSize,
	})

	rc, _ = NewRateLimterConfig(dummyId, WithLimit(10), WithWindowSize(20))

	assert.Equal(t, rc, RateConfig{
		Id:         dummyId,
		Limit:      10,
		WindowSize: 20,
	})
}
