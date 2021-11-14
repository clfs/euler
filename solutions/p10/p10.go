package p10

import "math/big"

func Solve() int64 {
	tmp := new(big.Int)

	var res int64
	for x := int64(2); x < 2000000; x += 1 {
		if tmp.SetInt64(x).ProbablyPrime(0) {
			res += x
		}
	}
	return res
}
