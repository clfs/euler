package p15

import "math/big"

func Solve() int64 {
	// There are $\binom{40}{20}$ ways to arrange 20 downs and 20 rights. Easy!
	return new(big.Int).Binomial(40, 20).Int64()
}
