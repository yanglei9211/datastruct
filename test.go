package main

import (
	"container/heap"
	"datastruct/priorityqueue"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func test() {
	x := []bson.ObjectId{
		bson.ObjectIdHex("50e227000000000000002101"),
		bson.ObjectIdHex("50e227000000000000002102"),
		bson.ObjectIdHex("50e227000000000000002103"),
		bson.ObjectIdHex("50e227000000000000002104"),
		bson.ObjectIdHex("50e227000000000000002105"),
		bson.ObjectIdHex("50e227000000000000002106"),
		bson.ObjectIdHex("50e227000000000000002107"),
		bson.ObjectIdHex("50e227000000000000002108"),
		bson.ObjectIdHex("50e227000000000000002109"),
	}
	y := []float64{132, 341, 3124, 123, 4321, 4, 321, 4312, 4123}
	pq := make(priorityqueue.PriorityQueue, len(x))
	for i := 0; i < len(x); i++ {
		pq[i] = &priorityqueue.Item{
			Value: x[i],
			Score: y[i],
			Index: i,
		}
	}
	heap.Init(&pq)
	for pq.Len() > 0 {
		res := heap.Pop(&pq).(*priorityqueue.Item)
		fmt.Println(res.Value, res.Score)
	}
}

func main() {
	test()
}
