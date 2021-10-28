package queue

import (
	"testing"
)

func TestListQueue(t *testing.T) {
	q := NewListQueue()
	cases := []int{1, 2, 3, 4, 5, 6, 7}
	for _, c := range cases {
		q.Push(c)
	}
	for !q.IsEmpty() {
		t.Log(q.Front().(int))
		q.Pop()
	}
}

func TestSliceQueue(t *testing.T) {
	q := NewSliceQueue()
	cases := []int{1, 2, 3, 4, 5, 6, 7}
	for _, c := range cases {
		q.Push(c)
	}
	for !q.IsEmpty() {
		t.Log(q.Front().(int))
		q.Pop()
	}
}