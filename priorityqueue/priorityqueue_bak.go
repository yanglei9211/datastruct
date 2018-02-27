package priorityqueue

//
//import (
//	"container/heap"
//)
//
//type Item struct {
//	Value interface{}
//	Score float64
//	Index int // 更新的时候用
//}
//
//type PriorityQueue []*Item
//
//func (q PriorityQueue) Len() int { return len(q) }
//
//func (q PriorityQueue) Less(i, j int) bool {
//	return q[i].Score > q[j].Score
//}
//
//func (q PriorityQueue) Swap(i, j int) {
//	q[i], q[j] = q[j], q[i]
//	q[i].Index = i
//	q[j].Index = j
//}
//
//func (q *PriorityQueue) Push(x interface{}) {
//	n := len(*q)
//	item := x.(*Item)
//	item.Index = n
//	*q = append(*q, item)
//}
//
//func (q *PriorityQueue) Pop() interface{} {
//	old := *q
//	n := len(old)
//	item := old[n-1]
//	item.Index = -1
//	*q = old[0 : n-1]
//	return item
//}
//
//func (q *PriorityQueue) Update(item *Item, value interface{}, score float64) {
//	item.Value = value
//	item.Score = score
//	heap.Fix(q, item.Index)
//}
