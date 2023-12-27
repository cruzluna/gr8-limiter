package cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheConcurrent(t *testing.T) {
	strs := []string{
		"string0", "string1",
		"string2", "string3",
		"string4", "string5",
		"string0", "string1",
		"string2", "string3",
		"string4", "string5",
	}
	lru := New[string](len(strs))

	ch := make(chan string)
	done := make(chan string)
	go func() {
		defer close(done)
		for _, st := range strs {
			ch <- st
		}
	}()

	for {
		select {
		case s := <-ch:
			// do work
			go func(str string) {
				start := time.Now()
				// add to cache
				value, ok := lru.Get(str)
				if !ok {
					lru.Add(str, "")
				}
				fmt.Printf("%s, %s, %d string\n", str, time.Since(start), len(value))
			}(s)
		case <-done:
			return
		}
	}
}

func TestCreateLRUCache(t *testing.T) {
	lru := New[string](10)
	assert.Equal(t, 10, lru.capacity, 10)
	lru.Add("dummyInsert", "dummyValue")
	assert.Equal(t, "dummyValue", lru.cache["dummyInsert"].Value.value)

	elem, _ := lru.cache["dummyInsert"]

	assert.NotEmpty(t, elem)
	assert.Equal(t, "dummyValue", lru.list.Front().Value.value)
	lru.Add("dummyInsert2", "dummyValue2")

	random := "abdasdjfladjfdkafsldjfa"
	for i := 0; i < 11; i++ {
		lru.Add(random[0:i], fmt.Sprint(i*i*i))
	}
	assert.Equal(t, 10, lru.list.Len())
	assert.Equal(t, 10, len(lru.cache))

	elem, ok := lru.cache["dummyInsert"]
	assert.Equal(t, false, ok)
	assert.Empty(t, elem)
}
