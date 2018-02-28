package heap

import (
	"fmt"
	"testing"
)

func Test_heap(t *testing.T) {
	a := []int{1, 5, 12, 32, 9, 7, 4, 35, 12, 89}
	s := []Item{}
	for idx, x := range a {
		s = append(s, Item{Score: float64(x), Value: x, Index: idx})
	}
	h := Heap{}
	h.Init(s)
	for h.Len() > 0 {
		res := h.Pop()
		fmt.Println(res)

	}
	//for i := 0; i < h.Len(); i++ {
	//	fmt.Println(h.ItemList[i])
	//}
	//h.Push(Item{Value:123, Score:123})
	//h.Push(Item{Value:2, Score:2})
	//res = h.Pop()
	//fmt.Println(res)
}
