package kangtuo

type KT struct {
	fac []int
	N int
}

func (s *KT) init(n int) {
	s.N = n
	s.fac = make([]int, 10)
	for i := 1; i < s.N; i++ {
		s.fac[i] = s.fac[i-1] * i
	}
}

//func (s *KT) kangtuo(n int, s []int) int {
//	sum := 0
//	for i := 0; i < n; i ++ {
//		t := 0
//		for j := i+1; j < n; j++ {
//			if
//		}
//	}
//}

func main() {
	//x := 2143
	kt := KT{}
	kt.init(10)
}
