package cache

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var s = []string{
	"string0", "string1",
	"string2", "string3",
	"string4", "string5",
	"string0", "string1",
	"string2", "string3",
	"string4", "string5",
}

func incomingString() <-chan string {
	ch := make(chan string)
	go func() {
		for _, st := range s {
			ch <- st
		}
		close(ch)
	}()
	return ch
}

func TestCacheConcurrent(t *testing.T) {
	lru := New[string](len(s))

	var wg sync.WaitGroup
	for val := range incomingString() {
		wg.Add(1)

		go func(str string) {
			defer wg.Done()
			start := time.Now()

			// add to cache
			value, ok := lru.Get(str)
			if !ok {
				lru.Add(str, "")
			}

			fmt.Printf("%s, %s, %d string\n", str, time.Since(start), len(value))
		}(val)
	}
	wg.Wait()
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
