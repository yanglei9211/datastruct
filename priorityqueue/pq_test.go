package priorityqueue

import (
	"fmt"
	"testing"
)

type DataList struct {
	ItemList
}

func (s DataList) Less(i, j int) bool {
	return s.ItemList[i].Score > s.ItemList[j].Score
}

func NewDataList() DataList {
	dl := DataList{}
	dl.ItemList = make(ItemList, 0)
	return dl
}

func Test_pq(t *testing.T) {
	dl := NewDataList()
	pq := NewPrioirtyQueue(&dl)
	pq.PushItem(&Item{Score: 12, Value: 12})
	pq.PushItem(&Item{Score: 22, Value: 22})
	pq.PushItem(&Item{Score: 2, Value: 2})
	pq.PushItem(&Item{Score: 13, Value: 13})

	pq.PushItem(&Item{Score: 13, Value: 13})
	pq.PushItem(&Item{Score: 33, Value: 33})
	pq.PushItem(&Item{Score: 1, Value: 1})
	res := pq.PopItem()
	fmt.Println(res.Index, res.Value, res.Score)
}
