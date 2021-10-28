package queue

import "container/list"

type Queue interface {
	Front() interface{}
	Push(interface{})
	Pop()
	IsEmpty() bool
}

type listQueue struct{
	list *list.List
}

func NewListQueue() Queue {
	l := list.New()
	return &listQueue{list: l}
}

func (l *listQueue) Front() interface{} {
	return l.list.Front().Value
}

func (l *listQueue) Push(i interface{}) {
	l.list.PushBack(i)
}

func (l *listQueue) Pop() {
	l.list.Remove(l.list.Front())
}

func (l *listQueue) IsEmpty() bool {
	return l.list.Len() == 0
}

type sliceQueue struct {
	array []interface{}
}

func (s *sliceQueue) Front() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.array[0]
}

func (s *sliceQueue) Push(i interface{}) {
	s.array = append(s.array, i)
}

func (s *sliceQueue) Pop() {
	if len(s.array) == 0 {
		return
	}
	s.array = s.array[1:]
}

func (s *sliceQueue) IsEmpty() bool {
	return len(s.array) == 0
}

func NewSliceQueue() Queue {
	return &sliceQueue{array: make([]interface{}, 0)}
}