package kangtuo

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	kt := KT{}
	kt.init(10)
	//fmt.Println(kt.fac)
	//fmt.Println(kt.kangtuo(4, []int{2,1,4,3}))
	res := kt.reverse_kangtuo(4, 2)
	fmt.Println(res)

}
