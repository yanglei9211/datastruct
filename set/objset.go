package set

import (
	"gopkg.in/mgo.v2/bson"
)

type ObjSet map[bson.ObjectId]struct{}

func NewObjSet(objs []bson.ObjectId) ObjSet {
	set := make(ObjSet, len(objs))
	set.AddList(objs)
	return set
}

func (set ObjSet) Add(s bson.ObjectId) ObjSet {
	set[s] = struct{}{}
	return set
}

func (set ObjSet) Del(s bson.ObjectId) ObjSet {
	delete(set, s)
	return set
}

func (set ObjSet) Has(s bson.ObjectId) bool {
	_, has := set[s]
	return has
}

func (set ObjSet) AddList(ss []bson.ObjectId) ObjSet {
	for _, s := range ss {
		set.Add(s)
	}
	return set
}

func (set ObjSet) ToList() []bson.ObjectId {
	ss := make([]bson.ObjectId, 0, len(set))
	for k := range set {
		ss = append(ss, k)
	}
	return ss
}

func (set ObjSet) Size() int {
	return len(set)
}

func (set ObjSet) Include(b ObjSet) bool {
	for key := range b {
		if _, found := set[key]; !found {
			return false
		}
	}
	return true
}
func (set ObjSet) Difference(b ObjSet) int {
	count := 0
	for key := range set {
		if _, found := b[key]; !found {
			count++
		}
	}
	return count
}

func (set ObjSet) Differences(b ObjSet) int {
	return set.Difference(b) + b.Difference(set)
}

func (set ObjSet) Union(b ObjSet) ObjSet {
	for k := range b {
		set.Add(k)
	}
	return set
}

func (set ObjSet) Inter(b ObjSet) ObjSet {
	res := ObjSet{}
	for k := range set {
		if _, found := b[k]; found {
			res.Add(k)
		}
	}
	return res
}
