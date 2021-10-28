package LRU

type LRUCache interface {
	Get() interface{}
	Put(interface{})
}
