package LFU

import (
	"fmt"
	"sort"
)

type LFU interface {
	Get(interface{}) interface{}
	Put(interface{}, interface{})
	Print() string
}

type LFUCache struct {
	keyToVal    map[interface{}]interface{}
	keyToFreq 	map[interface{}]uint
	freqToLh  	map[uint]*linkedHashset
	minFreq   	uint
	capacity  	int
}

func NewLFUCache(capacity int) LFU {
	return &LFUCache{
		keyToVal:  map[interface{}]interface{}{},
		keyToFreq: map[interface{}]uint{},
		freqToLh:  map[uint]*linkedHashset{},
		minFreq:   0,
		capacity:  capacity,
	}
}

func (l *LFUCache) increaseFreq(key interface{}) {
	freq := l.keyToFreq[key]
	l.keyToFreq[key] = freq + 1

	lh := l.freqToLh[freq]
	lh.delete(key)
	if lh.empty() {
		delete(l.freqToLh, freq)
		if freq == l.minFreq {
			l.minFreq++
		}
	}

	if newLh, ok := l.freqToLh[freq + 1]; !ok {
		newLh = newLinkedHashset()
		l.freqToLh[freq + 1] = newLh
		newLh.insert(key)
	} else {
		newLh.insert(key)
	}
}

func (l *LFUCache) removeLeastFrequency() {
	lh := l.freqToLh[l.minFreq]
	key := lh.removeLast()
	delete(l.keyToFreq, key)
	delete(l.keyToVal, key)
}

func (l *LFUCache) Get(key interface{}) interface{} {
	if val, ok := l.keyToVal[key]; ok {
		l.increaseFreq(key)
		return val
	} else {
		return nil
	}
}

func (l *LFUCache) Put(key interface{}, value interface{}) {
	if _, ok := l.keyToVal[key]; ok {
		l.increaseFreq(key)
		l.keyToVal[key] = value
	} else {
		if len(l.keyToFreq) == l.capacity {
			l.removeLeastFrequency()
		}
		if _, ok := l.freqToLh[1]; !ok {
			l.freqToLh[1] = newLinkedHashset()
		}
		l.freqToLh[1].insert(key)
		l.keyToFreq[key] = 1
		l.keyToVal[key] = value
		l.minFreq = 1
	}
}

type pairsAtFreq struct {
	freq uint
	pairs string
}

type allPairs []*pairsAtFreq

func (a allPairs) Len() int { return len(a) }
func (a allPairs) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a allPairs) Less(i, j int) bool { return a[i].freq < a[j].freq }

func (l *LFUCache) Print() string {
	var (
		pairs  allPairs
		result string
	)
	for freq, lh := range l.freqToLh {
		pairsString := ""
		keys := lh.traverse()
		for _, key := range keys {
			pairsString += fmt.Sprintf("[%v: %v]", key, l.keyToVal[key])
		}
		pairs = append(pairs, &pairsAtFreq{freq: freq, pairs: pairsString})
	}
	sort.Sort(pairs)
	for _, pairs := range pairs {
		result += fmt.Sprintf("\nfrequency %d: %s", pairs.freq, pairs.pairs)
	}

	return result
}