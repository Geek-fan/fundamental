package LFU

import (
	"container/list"
)

type linkedHashset struct {
	hashset map[interface{}]*list.Element
	list    *list.List
}

func newLinkedHashset() *linkedHashset {
	return &linkedHashset{
		hashset: map[interface{}]*list.Element{},
		list:    list.New(),
	}
}

func (l *linkedHashset) traverse() (res []interface{}) {
	for elem := l.list.Front(); elem != nil; elem = elem.Next() {
		res = append(res, elem.Value)
	}
	return
}

func (l *linkedHashset) empty() bool {
	return len(l.hashset) == 0
}

func (l *linkedHashset) removeLast() interface{} {
	key := l.list.Remove(l.list.Back())
	delete(l.hashset, key)
	return key
}

func (l *linkedHashset) insert(key interface{}) {
	elem := l.list.PushFront(key)
	l.hashset[key] = elem
}

func (l *linkedHashset) delete(key interface{}) {
	elem := l.hashset[key]
	l.list.Remove(elem)
	delete(l.hashset, key)
}