package p5

func Solve() int {
	// Primes less than 20: 2, 3, 5, 7, 11, 13, 17, 19.
	// The answer's prime factorization must contain each prime at least once.

	step := 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19

	for i := step; ; i += step {
		for j := 2; j < 20; j++ { // skip 1 and 20
			if i%j != 0 {
				break
			}
			if j == 19 {
				return i
			}
		}
	}
}
