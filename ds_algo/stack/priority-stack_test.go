package stack

import (
	"reflect"
	"testing"
)

func TestNewPriorityStack(t *testing.T) {
	s := NewPriorityStack()
	arrays := []int{1, 1, 1, 2, 2, 3}
	for _, num := range arrays {
		s.Push(num)
	}
	tests := []struct {
		name string
		want int
	}{
		{"1", 1},
		{"2", 2},
		{"3", 1},
		{"4", 3},
		{"5", 2},
		{"6", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Pop().(int); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
