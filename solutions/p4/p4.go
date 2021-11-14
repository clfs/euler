package p4

import (
	"fmt"
	"sort"
	"strconv"
)

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// Solve solves problem 4.
func Solve() (int, error) {
	var products []int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			products = append(products, i*j)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(products)))

	for _, p := range products {
		if isPalindrome(p) {
			return p, nil
		}
	}

	return 0, fmt.Errorf("no solution found")
}
