package stack

import (
	"fmt"
	"testing"
)

func Test_stack(t *testing.T) {
	a := []int{1, 3, 5, 4, 2}
	st := NewStack()
	for _, x := range a {
		st.Push(x)
	}
	for !st.Empty() {
		r, _ := st.Pop()
		fmt.Println(r.(int))
	}

	st = NewStack()
	st.Push(1)
	st.Push(4)
	st.Push(3)
	fmt.Println(st.Pop())
	st.Push(2)
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
	fmt.Println(st.Pop())
}
