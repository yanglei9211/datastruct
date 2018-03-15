package merge

func less(x, y int) bool {
	return x < y
}

func merge(l, r []int) []int {
	i := 0
	j := 0
	res := make([]int, 0, len(l)+len(r))
	for i < len(l) && j < len(r) {
		if less(l[i], r[j]) {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	for i < len(l) {
		res = append(res, l[i])
		i++
	}
	for j < len(r) {
		res = append(res, r[j])
		j++
	}
	return res
}

func ms(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) >> 1
	left := ms(a[:mid])
	right := ms(a[mid:])
	res := merge(left, right)
	return res
}
