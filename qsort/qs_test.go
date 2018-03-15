package qsort

import (
	"fmt"
	"testing"
)

func Test_qs(t *testing.T) {
	a := []int{1, 9, 32, 34, 123, 3, 2, 67, 1, 32, 90, 123, 0, -12, -34, 44312, 123, 34, 32}
	q := Qsort{}
	q.d = a
	q.qs(0, len(a)-1)
	fmt.Println(q.d)
}
