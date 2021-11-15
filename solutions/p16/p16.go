package p16

import "math/big"

func digitSum(n *big.Int) int {
	var res int
	for _, d := range n.String() {
		res += int(d - '0')
	}
	return res
}

func Solve() int {
	n := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil) // 2^1000
	return digitSum(n)
}
