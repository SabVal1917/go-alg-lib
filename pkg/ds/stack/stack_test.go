package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack[int]()
	if !s.IsEmpty() {
		t.Errorf("stack should be empty after creating")
	}

	s.Push(10)
	s.Push(20)
	if check, _ := s.Top(); check != 20 {
		t.Errorf("stack top should be 20")
	}

	if check, _ := s.Pop(); check != 20 {
		t.Errorf("stack top should be 20")
	}
	s.Pop()
	if _, err := s.Pop(); err != ErrEmptyStack {
		t.Errorf("expected ErrEmptyStack, got %v", err)
	}

}
