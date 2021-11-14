package p9

func Solve() int {
	var res int
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				res = a * b * c
				break
			}
		}
	}
	return res
}
