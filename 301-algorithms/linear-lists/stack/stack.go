package main

type MyItem int

type Stack struct {
	capacity uint64
	depth    uint64
	list     []MyItem
}

func NewStack(capacity uint64) *Stack {
	return &Stack{
		capacity: capacity,
		list:     make([]MyItem, capacity),
		depth:    0,
	}
}

func (s *Stack) Push(value MyItem) {
	if s.depth >= s.capacity {
		panic("out of cap")
	}
	s.list[s.depth] = value
	s.depth++
}

func (s *Stack) Pop() MyItem {
	if s.depth > 0 {
		s.depth--
		return s.list[s.depth]
	}
	return -1
}

func main() {
	stack := NewStack(4)
	stack.Push(1)
	stack.Push(2)
	println(stack.Pop())
	stack.Push(3)
	stack.Push(4)
	println(stack.Pop())
	println(stack.Pop())
	println(stack.Pop())
}
