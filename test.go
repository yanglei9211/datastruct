package main

import (
	"container/heap"
	"datastruct/priorityqueue"
	"datastruct/skiplist"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func calcScore(x int) float64 {
	return float64(x) * 1.1
}

func test_priority_queue() {
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

func test_skiplist() {
	sp := skiplist.Skiplist{}
	sp.Init()
	datas := []skiplist.Data{
		skiplist.Data{X: 1},
		skiplist.Data{X: 4},
		skiplist.Data{X: 8},
		skiplist.Data{X: 12},
		skiplist.Data{X: 19},
		skiplist.Data{X: 33},
		skiplist.Data{X: 45},
		skiplist.Data{X: 62},
		skiplist.Data{X: 69},
		skiplist.Data{X: 123},
		skiplist.Data{X: 150},
		skiplist.Data{X: 189},
		skiplist.Data{X: 216},
		skiplist.Data{X: 246},
		skiplist.Data{X: 286},
		skiplist.Data{X: 323},
		skiplist.Data{X: 366},
		skiplist.Data{X: 423},
		skiplist.Data{X: 532},
		skiplist.Data{X: 587},
	}
	scores := make([]float64, 0, len(datas))
	for _, dt := range datas {
		scores = append(scores, calcScore(dt.X))
	}
	for i := range datas {
		sp.Insert(scores[i], &datas[i])
	}
	var x *skiplist.SkiplistNode
	for i := 0; i < skiplist.SKIPLIST_MAX_LEVEL; i++ {
		if sp.Head.Level[i].Forward != nil {
			fmt.Printf("%d: ", i+1)
			x = sp.Head.Level[i].Forward
			for x != nil {
				fmt.Printf("%d --- ", x.Obj.X)
				x = x.Level[i].Forward
			}
			fmt.Println()
		}
	}
	qRange := skiplist.ScoreRange{Left: 234, Right: 444}
	res := sp.FirstInRange(&qRange)
	fmt.Println(res.Obj, res.Score)
	res = sp.LastInRange(&qRange)
	fmt.Println(res.Obj, res.Score)
	//head := skiplist.Skiplist{}
	//dt := []int{3,2,1}
	//for _, x := range dt {
	//    node := skiplist.Skiplist{}
	//    node.Next = head.Next
	//    node.Score = x
	//    head.Next = &node
	//}
	//tp := head.Next
	//fmt.Println(head.Score)
	//fmt.Println(tp)
	//for tp!= nil {
	//   fmt.Println(tp.Score)
	//   tp = tp.Next
	//}
}

func main() {
	// test_priority_queue()
	test_skiplist()
}
