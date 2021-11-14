package p7

import (
	"math/big"
)

func Solve() int64 {
	var (
		one  = big.NewInt(1)
		seen int
	)

	for x := big.NewInt(2); ; x.Add(x, one) {
		if x.ProbablyPrime(0) {
			seen++
		}
		if seen == 10001 {
			return x.Int64()
		}
	}
}
