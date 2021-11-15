package p12

import "math"

func NDivisors(x int) int {
	var res int
	bound := int(math.Sqrt(float64(x))) // rounding isn't an issue for numbers this small
	for i := 1; i <= bound; i++ {
		if x%i == 0 {
			if i*i == x {
				res += 1
			} else {
				res += 2
			}
		}
	}
	return res
}

func Solve() int {
	for n := 1; ; n++ {
		t := n * (n + 1) / 2
		if NDivisors(t) > 500 {
			return t
		}
	}
}
