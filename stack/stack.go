package stack

import (
	"container/list"
	"errors"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	ls := list.New()
	return &Stack{ls}
}

func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack empty")
	} else {
		e := s.list.Back()
		if e != nil {
			s.list.Remove(e)
			return e.Value, nil
		}
		return nil, errors.New("inter error")
	}
}

func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack empty")
	} else {
		e := s.list.Back()
		if e != nil {
			return e.Value, nil
		}
		return nil, errors.New("inter error")
	}
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}
