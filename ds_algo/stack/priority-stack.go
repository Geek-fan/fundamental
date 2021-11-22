package stack

/*
	Priority stack is a special stack that pops element having the highest frequency first,
	elements having the same frequency will pop the last pushed one.
 */
type PriorityStack interface {
	Push(interface{})
	Pop() interface{}
	IsEmpty() bool
}

type customStack struct {
	stacks []Stack
	count map[interface{}]int
}

func NewPriorityStack() PriorityStack {
	return &customStack{count: make(map[interface{}]int)}
}

func (s *customStack) Push(elem interface{}) {
	s.count[elem]++
	if s.count[elem] > len(s.stacks) {
		s.stacks = append(s.stacks, &sliceStack{})
	}
	s.stacks[s.count[elem] - 1].Push(elem)
}

func (s *customStack) Pop() interface{} {
	numStacks := len(s.stacks)
	if numStacks == 0 {
		return nil
	}
	stack := s.stacks[numStacks - 1]
	elem := stack.Pop()
	s.count[elem]--
	if stack.IsEmpty() {
		s.stacks = s.stacks[:numStacks - 1]
	}
	return elem
}

func (s *customStack) IsEmpty() bool {
	return len(s.stacks) == 0
}