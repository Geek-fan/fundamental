package LRU

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewLRUCache(t *testing.T) {
	lruCache := NewLRUCache(5)
	keys := []int{1, 2, 3, 4, 5, 2, 1, 8}
	values := []int{2, 6, 7, 4, 5, 3, 5, 1}

	for i, key := range keys {
		lruCache.Put(key, values[i])
		pairs := lruCache.GetAll()
		t.Logf("Put [%d, %d]\tCache: %v", key, values[i], pairs)
	}

	rand.Seed(int64(time.Now().Second()))
	for i := 0; i < 5; i++ {
		key := rand.Uint32() % 5
		val := -1
		v := lruCache.Get(int(key))
		if v != nil {
			val = v.(int)
		} else {
			t.Logf("Key %d not found", key)
			continue
		}
		pairs := lruCache.GetAll()
		t.Logf("Query %d Get %d\tCache: %v", key, val, pairs)
	}
}
