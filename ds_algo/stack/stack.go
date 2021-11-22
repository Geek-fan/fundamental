package stack

import "container/list"

type Stack interface {
	Push(interface{})
	Pop() interface{}
	IsEmpty() bool
}

type sliceStack struct {
	array []interface{}
}

func (s *sliceStack) Push(elem interface{}) {
	s.array = append(s.array, elem)
}

func (s *sliceStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	length := len(s.array)
	elem := s.array[length - 1]
	s.array = s.array[:length - 1]
	return elem
}

func (s *sliceStack) IsEmpty() bool {
	return len(s.array) == 0
}

func NewSliceStack() Stack {
	return &sliceStack{[]interface{}{}}
}

type listStack struct {
	list *list.List
}

func (l *listStack) Push(i interface{}) {
	l.list.PushBack(i)
}

func (l *listStack) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}
	elem := l.list.Back()
	l.list.Remove(elem)
	return elem.Value
}

func (l *listStack) IsEmpty() bool {
	return l.list.Len() == 0
}

func NewListStack() Stack {
	return &listStack{list.New()}
}


