package merge

import (
	"fmt"
	"testing"
)

func Test_m(t *testing.T) {
	a := []int{1, 9, 32, 34, 123, 3, 2, 67, 1, 32, 90, 123, 0, -12, -34, 44312, 123, 34, 32}
	fmt.Println(a)
	b := ms(a)
	fmt.Println(b)
}
