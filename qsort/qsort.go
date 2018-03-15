package qsort

import "fmt"

type Qsort struct {
	d []int
}

func (s *Qsort) qs(l, r int) {
	rl := l
	rr := r
	if l == r {
		return
	}
	mid := s.d[(l+r)>>1]
	fmt.Println(l, r, mid)
	fmt.Println(s.d)
	for l <= r {
		fmt.Println(l, r)
		for s.d[l] < mid {
			l++
		}
		for s.d[r] > mid {
			r--
		}
		if l <= r {
			s.d[l], s.d[r] = s.d[r], s.d[l]
			l++
			r--
		}
	}
	fmt.Println("----------")
	fmt.Println(s.d)
	if l < rr {
		s.qs(l, rr)
	}
	if rl < r {
		s.qs(rl, r)
	}
}
