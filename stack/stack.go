package main

import "fmt"

// Stack represents a stack data structure
type Stack struct {
	items []int
}

// Push adds an item to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // Return false if the stack is empty
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

// Peek returns the top item without removing it
func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // Return false if the stack is empty
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack after pushes:", stack.items)

	top, _ := stack.Pop()
	fmt.Println("Popped:", top)
	fmt.Println("Stack after pop:", stack.items)

	peek, _ := stack.Peek()
	fmt.Println("Peek:", peek)

	fmt.Println("Is stack empty?", stack.IsEmpty())
}
