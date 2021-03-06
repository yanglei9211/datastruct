package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const SKIPLIST_MAX_LEVEL = 32
const SKIPLIST_P = 0.25

type Data struct {
	X int
}

type ScoreRange struct {
	Left, Right float64
}

func (u *Data) Compare(d *Data) int {
	return u.X - (*d).X
}

type SkiplistLevel struct {
	Forward *SkiplistNode
	Span    int
}

type SkiplistNode struct {
	Obj      *Data
	Score    float64
	Next     *SkiplistNode
	Backward *SkiplistNode
	Level    []SkiplistLevel
}

type Skiplist struct {
	Head   *SkiplistNode
	Tail   *SkiplistNode
	Length int
	Level  int // 最大层数
}

func CreateNode(level int, score float64, obj *Data) *SkiplistNode {
	node := SkiplistNode{
		Level: make([]SkiplistLevel, level),
		Score: score,
		Obj:   obj,
	}
	return &node
}

func (s *Skiplist) Init() {
	rand.Seed(time.Now().Unix())
	s.Level = 1
	s.Length = 0
	s.Head = CreateNode(SKIPLIST_MAX_LEVEL, 0, nil)
	for i := 0; i < SKIPLIST_MAX_LEVEL; i++ {
		s.Head.Level[i].Forward = nil
		s.Head.Level[i].Span = 0
	}
	s.Head.Backward = nil
	s.Tail = nil
}

func GenRandomLevel() int {
	level := 1
	for (float64(rand.Int()&0xffff) < SKIPLIST_P*0xffff) && (level < SKIPLIST_MAX_LEVEL-1) {
		level++
	}
	return level
}

func (s *Skiplist) Insert(score float64, obj *Data) *SkiplistNode {
	var x *SkiplistNode
	update := make([]*SkiplistNode, SKIPLIST_MAX_LEVEL)
	rank := make([]int, SKIPLIST_MAX_LEVEL)
	x = s.Head
	for i := s.Level - 1; i >= 0; i-- {
		if s.Level-1 == i {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.Level[i].Forward != nil && (x.Level[i].Forward.Score < score ||
			(x.Level[i].Forward.Score == score && (*x.Level[i].Forward.Obj).Compare(obj) < 0)) {
			rank[i] += x.Level[i].Span
			x = x.Level[i].Forward
		}
		update[i] = x
	}
	level := GenRandomLevel()
	if level > s.Level {
		for i := s.Level; i < level; i++ {
			rank[i] = 0
			update[i] = s.Head
			update[i].Level[i].Span = s.Length
		}
		s.Level = level
	}
	x = CreateNode(level, score, obj)

	for i := 0; i < level; i++ {
		x.Level[i].Forward = update[i].Level[i].Forward
		update[i].Level[i].Forward = x
		x.Level[i].Span = (rank[0] - rank[i]) + 1
	}

	for i := level; i < s.Level; i++ {
		update[i].Level[i].Span++
	}

	if update[0] == s.Head {
		x.Backward = nil
	} else {
		x.Backward = update[0]
	}
	if x.Level[0].Forward != nil {
		x.Level[0].Forward.Backward = x
	} else {
		s.Tail = x
	}
	s.Level++
	return x
}

func (s *Skiplist) DeleteNode(x *SkiplistNode, update []*SkiplistNode) {
	for i := 0; i < s.Level; i++ {
		if update[i].Level[i].Forward == x {
			update[i].Level[i].Span += x.Level[i].Span - 1
			update[i].Level[i].Forward = x.Level[i].Forward
		} else {
			update[i].Level[i].Span -= 1
		}
	}
	if x.Level[0].Forward != nil {
		x.Level[0].Forward.Backward = x.Backward
	} else {
		s.Tail = x.Backward
	}
	for s.Level > 1 && s.Head.Level[s.Level-1].Forward == nil {
		s.Level--
	}
	s.Length--
}

func (s *Skiplist) DeleteData(score float64, obj *Data) bool {
	var x *SkiplistNode
	update := make([]*SkiplistNode, SKIPLIST_MAX_LEVEL)
	x = s.Head
	for i := s.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil && (x.Level[i].Forward.Score < score ||
			(x.Level[i].Forward.Score == score && (*x.Level[i].Forward.Obj).Compare(obj) < 0)) {
			x = x.Level[i].Forward
		}
		update[i] = x
	}
	if x != nil && x.Level[0].Forward != nil {
		x = x.Level[0].Forward
	}
	if obj.Compare(x.Obj) != 0 {
		return false
	}
	s.DeleteNode(x, update)
	return true
}

func (s *Skiplist) sIsInRange(r *ScoreRange) bool {
	var x *SkiplistNode
	if r.Left > r.Right {
		return false
	}
	x = s.Tail
	if x == nil || r.Left > x.Score {
		return false
	}
	x = s.Head.Level[0].Forward
	if x == nil || r.Right < x.Score {
		return false
	}
	return true
}

func (s *Skiplist) FirstInRange(r *ScoreRange) *SkiplistNode {
	var x *SkiplistNode
	if !s.sIsInRange(r) {
		return nil
	}
	x = s.Head
	for i := s.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil && x.Level[i].Forward.Score < r.Left {
			x = x.Level[i].Forward
		}
	}
	x = x.Level[0].Forward
	if x.Score < r.Left {
		return nil
	}
	return x
}

func (s *Skiplist) LastInRange(r *ScoreRange) *SkiplistNode {
	var x *SkiplistNode
	if !s.sIsInRange(r) {
		return nil
	}
	x = s.Head
	for i := s.Level - 1; i >= 0; i-- {
		for x.Level[i].Forward != nil && x.Level[i].Forward.Score < r.Right {
			x = x.Level[i].Forward
		}
	}
	if x.Score >= r.Left {
		return x
	} else {
		return nil
	}
}

func (s *Skiplist) show() {
	var x *SkiplistNode
	for i := 0; i < SKIPLIST_MAX_LEVEL; i++ {
		if s.Head.Level[i].Forward != nil {
			fmt.Printf("%d: ", i+1)
			x = s.Head.Level[i].Forward
			for x != nil {
				fmt.Printf("%d, step: %d ------->", x.Obj.X, x.Level[i].Span)
				x = x.Level[i].Forward
			}
			fmt.Println()
		}
	}
}
