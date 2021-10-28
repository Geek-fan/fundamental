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
		{"1st", 1},
		{"2nd", 2},
		{"3rd", 1},
		{"4th", 3},
		{"5th", 2},
		{"6th", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Pop().(int); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
