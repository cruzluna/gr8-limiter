package cache

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
