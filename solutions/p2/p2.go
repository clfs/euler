package p2

// Solve solves problem 2.
func Solve() int {
	var res int

	a, b := 1, 1
	for b < 4000000 {
		if b%2 == 0 {
			res += b
		}
		a, b = b, a+b
	}

	return res
}
