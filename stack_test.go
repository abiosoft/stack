package stack

import (
	"testing"
)

func Test(t *testing.T) {
	//Push 1, Peek, Size, Pop, Size

	s := NewStack()
	s.Push(1)

	if i, ok := s.Peek().(int); !ok || i != 1 {
		t.Errorf("expected %v found %v", 1, i)
	}

	s.Push(5)

	if i, ok := s.Peek().(int); !ok || i != 5 {
		t.Errorf("expected %v found %v", 5, i)
	}

	if s.Size() != 2 {
		t.Errorf("expected %v found %v", 2, s.Size())
	}

	val := s.Pop()
	if i, ok := val.(int); !ok || i != 5 {
		t.Errorf("expected %v found %v", 5, i)
	}

	if s.Size() != 1 {
		t.Errorf("expected %v found %v", 1, s.Size())
	}
}
