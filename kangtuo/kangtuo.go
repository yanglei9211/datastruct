package kangtuo

import "fmt"

type KT struct {
	fac []int
	N   int
}

func (s *KT) init(n int) {
	s.N = n
	fmt.Println(n)
	s.fac = make([]int, 10)
	s.fac[0] = 1
	for i := 1; i < s.N; i++ {
		s.fac[i] = s.fac[i-1] * i
	}
	fmt.Println(s.fac)
}

func (s *KT) kangtuo(n int, a []int) int {
	sum := 0
	for i := 0; i < n; i++ {
		t := 0
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				t++
			}
		}
		sum += t * s.fac[n-i-1]
	}
	return sum + 1
}

func (s *KT) reverse_kangtuo(n, k int) []int {
	vis := make([]bool, s.N)
	res := make([]int, n)
	var i, j int
	for i = 0; i < n; i++ {
		t := k / s.fac[n-i-1]
		for j = 1; j < n; j++ {
			if !vis[j] {
				if t == 0 {
					break
				}
				t--
			}
		}
		res[i] = j
		vis[j] = true
		k %= s.fac[n-i-1]
	}
	return res
}

func T() {
	//x := 2143
}
