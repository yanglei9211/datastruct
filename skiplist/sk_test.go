package skiplist

import (
	"fmt"
	"testing"
)

func calcScore(x int) float64 {
	return float64(x) * 1.1
}

func Test_skip(t *testing.T) {
	sp := Skiplist{}
	sp.Init()
	datas := []Data{
		Data{X: 1},
		Data{X: 4},
		Data{X: 8},
		Data{X: 12},
		Data{X: 19},
		Data{X: 33},
		Data{X: 45},
		Data{X: 62},
		Data{X: 69},
		Data{X: 123},
		Data{X: 150},
		Data{X: 189},
		Data{X: 216},
		Data{X: 246},
		Data{X: 286},
		Data{X: 323},
		Data{X: 366},
		Data{X: 423},
		Data{X: 532},
		Data{X: 587},
	}
	scores := make([]float64, 0, len(datas))
	for _, dt := range datas {
		scores = append(scores, calcScore(dt.X))
	}

	for i := range datas {
		sp.Insert(scores[i], &datas[i])
	}

	sp.show()
	//qRange := ScoreRange{Left: 234, Right:444}
	//res := sp.FirstInRange(&qRange)
	//fmt.Println(res.Obj, res.Score)
	//res = sp.LastInRange(&qRange)
	//fmt.Println(res.Obj, res.Score)

	delData := Data{X: 246}
	s := calcScore(delData.X)
	r := sp.DeleteData(s, &delData)
	fmt.Println(r)
	sp.show()
}
