package codeWriter

import "fmt"

type StackItem struct {
	funcName string
	counter  int
}

type Stack struct {
	items []StackItem
}

func (s *Stack) Push(entry StackItem) {
	s.items = append(s.items, entry)
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Top() (StackItem, error) {
	if s.IsEmpty() {
		return StackItem{}, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}
