package queue

import (
	"container/list"
	"errors"
)

type Queue struct {
	list *list.List
}

func NewQueue() *Queue {
	ls := list.New()
	return &Queue{ls}
}

func (q *Queue) Push(value interface{}) {
	q.list.PushBack(value)
}

func (q *Queue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("queue empty")
	} else {
		e := q.list.Front()
		if e != nil {
			q.list.Remove(e)
			return e.Value, nil
		}
		return nil, errors.New("inter error")
	}
}

func (q *Queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("queue empty")
	} else {
		e := q.list.Front()
		if e != nil {
			return e.Value, nil
		} else {
			return nil, errors.New("inter error")
		}
	}
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}
