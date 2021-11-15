package p14

func Solve() int {
	var (
		bestStart  int
		bestLength int
	)
	for i := 1; i < 1_000_000; i++ {
		length := 1
		for n := i; n != 1; length++ {
			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
		}
		if length > bestLength {
			bestLength = length
			bestStart = i
		}
	}
	return bestStart
}
