package p3

func Solve() int {
	var res int

	n := 600851475143

	for f, np := 2, n; f*f <= n; f++ {
		for np%f == 0 {
			res = f
			np /= f
		}
	}

	return res
}
