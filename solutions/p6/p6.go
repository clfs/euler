package p6

func Solve() int {
	var sumOfSquares int

	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
	}

	var squareOfSum int

	for i := 1; i <= 100; i++ {
		squareOfSum += i
	}
	squareOfSum *= squareOfSum

	return absInt(sumOfSquares - squareOfSum)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
