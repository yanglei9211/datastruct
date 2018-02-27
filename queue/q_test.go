package queue

import (
	"fmt"
	"testing"
)

func Test_q(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	res, _ := q.Pop()
	fmt.Println(res)
	if res.(int) != 1 {
		t.Error("error!!!")
	}
	res, _ = q.Pop()
	fmt.Println(res)
	if res.(int) != 2 {
		t.Error("error!!!")
	}
	q.Push(5)
	q.Push(6)
	for !q.Empty() {
		res, _ = q.Pop()
		fmt.Println(res)
	}

}
