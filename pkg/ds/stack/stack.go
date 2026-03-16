package stack

import "errors"

type Stack[T any] struct {
	// Смысл есть реализовывать стек на слайсе,
	// а не на нодах. Ноды разбросаны по куче - траблы с cache miss
	// и несмотря на O(1) на Push в реализации через односвязный список,
	// на практике эта реализация будет медленней.
	data []T
}

var ErrEmptyStack = errors.New("stack is empty")

// isEmpty, Pop, Push, Clean, NewStack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}
func (s *Stack[T]) Push(elem T) {
	s.data = append(s.data, elem)
}
func (s *Stack[T]) Top() (T, error) {
	if len(s.data) == 0 {
		var zero T
		return zero, ErrEmptyStack
	}
	return s.data[len(s.data)-1], nil
}
func (s *Stack[T]) Pop() (T, error) {
	if len(s.data) == 0 {
		var zero T
		return zero, ErrEmptyStack
	}
	item := s.data[len(s.data)-1]
	var zero T
	s.data[len(s.data)-1] = zero
	s.data = s.data[:len(s.data)-1]
	return item, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
func (s *Stack[T]) Clean() {
	for !s.IsEmpty() {
		s.Pop()
	}
}
