package LRU

import (
	"container/list"
	"fmt"
)

type LRUCache interface {
	Get(interface{}) interface{}
	Put(interface{}, interface{})
	GetAll() []*Pair
}

type hashList struct {
	hmap map[interface{}]*list.Element
	list     *list.List
	capacity int
}

type Pair struct {
	Key   interface{}
	Value interface{}
}

func (p *Pair) String() string {
	return fmt.Sprintf("[%v: %v]", p.Key, p.Value)
}

func NewLRUCache(capacity int) LRUCache {
	return &hashList{
		hmap:     map[interface{}]*list.Element{},
		list:     list.New(),
		capacity: capacity,
	}
}

func getKey(elem *list.Element) interface{} {
	return elem.Value.(*Pair).Key
}

func getValue(elem *list.Element) interface{} {
	return elem.Value.(*Pair).Value
}

func setValue(elem *list.Element, val interface{}) {
	elem.Value.(*Pair).Value = val
}

func (h *hashList) makeRecent(key interface{}) {
	h.list.MoveToFront(h.hmap[key])
}

func (h *hashList) Get(key interface{}) interface{} {
	if elem, ok := h.hmap[key]; !ok {
		return nil
	} else {
		h.makeRecent(key)
		return getValue(elem)
	}
}

func (h *hashList) Put(key interface{}, val interface{}) {
	// key exists
	if elem, ok := h.hmap[key]; ok {
		h.makeRecent(key)
		setValue(elem, val)
	} else {
		// full, remote the last one
		if len(h.hmap) == h.capacity {
			back := h.list.Back()
			h.list.Remove(back)
			delete(h.hmap, getKey(back))
		}
		e := h.list.PushFront(&Pair{Key: key, Value: val})
		h.hmap[key] = e
	}
}

func (h *hashList) GetAll() (pairs []*Pair) {
	for elem := h.list.Front(); elem != nil; elem = elem.Next() {
		pairs = append(pairs, &Pair{Key: getKey(elem), Value: getValue(elem)})
	}
	return
}


