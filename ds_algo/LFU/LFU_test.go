package LFU

import (
	"math/rand"
	"testing"
	"time"
)

func TestNewLFUCache(t *testing.T) {
	var (
		keys []int
		values []int
	)
	lfuCache := NewLFUCache(5)

	rand.Seed(int64(time.Now().Second()))
	for i := 0; i < 20; i++ {
		keys = append(keys, rand.Int() % 10)
		values = append(values, rand.Int() % 10)
	}

	for i, key := range keys {
		lfuCache.Put(key, values[i])
		pairs := lfuCache.Print()
		t.Logf("Put [%d, %d]\tCache: %s", key, values[i], pairs)
	}

	for i := 0; i < 10; i++ {
		key := rand.Uint32() % 10
		val := -1
		v := lfuCache.Get(int(key))
		if v != nil {
			val = v.(int)
		} else {
			t.Logf("Key %d not found", key)
			continue
		}
		pairs := lfuCache.Print()
		t.Logf("Query %d Get %d\tCache: %s", key, val, pairs)
	}
}
