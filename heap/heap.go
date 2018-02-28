package heap

import "fmt"

type Item struct {
	Value interface{}
	Score float64
	Index int
}

//type ItemList []Item

type Heap struct {
	ItemList []Item
}

func (h Heap) Len() int {
	return len(h.ItemList)
}

func (h Heap) Less(i, j int) bool {
	return h.ItemList[i].Score < h.ItemList[j].Score
}

func (h *Heap) Swap(i, j int) {
	h.ItemList[i], h.ItemList[j] = h.ItemList[j], h.ItemList[i]
	h.ItemList[i].Index = i
	h.ItemList[j].Index = j
}

func (h *Heap) Init(its []Item) {
	h.ItemList = its
	n := len(its)
	for i := (n >> 1) - 1; i >= 0; i-- {
		h.down(i, n)
	}
	//for i := n-1; i > 0; i-- {
	//	h.up(i)
	//}
}

func (h *Heap) up(j int) {
	for j > 0 {
		i := (j - 1) >> 1
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func (h *Heap) down(i0, n int) bool {
	i := i0
	for {
		j1 := (i << 1) + 1
		if j1 >= n || j1 < 0 { //overflow
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && !h.Less(j1, j2) {
			j = j2
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

func (h *Heap) Push(it Item) {
	h.ItemList = append(h.ItemList, it)
	h.up(h.Len() - 1)
}

func (h *Heap) Pop() Item {
	res := h.ItemList[0]
	n := h.Len() - 1
	h.Swap(0, n)
	h.down(0, n)
	h.ItemList = h.ItemList[0:n]
	return res
}

func (h Heap) ShowList() {
	fmt.Println("--------------------")
	for i := 0; i < h.Len(); i++ {
		fmt.Println(h.ItemList[i])
	}
	fmt.Println("--------------------")
}
